/*==============================================================*/
/* DBMS name:      ORACLE Version 11g                           */
/* Created on:     2019/6/24 21:33:30                           */
/*==============================================================*/



-- Type package declaration
create or replace package PDTypes  
as
    TYPE ref_cursor IS REF CURSOR;
end;
/

-- Integrity package declaration
create or replace package IntegrityPackage AS
 procedure InitNestLevel;
 function GetNestLevel return number;
 procedure NextNestLevel;
 procedure PreviousNestLevel;
 end IntegrityPackage;
/

-- Integrity package definition
create or replace package body IntegrityPackage AS
 NestLevel number;

-- Procedure to initialize the trigger nest level
 procedure InitNestLevel is
 begin
 NestLevel := 0;
 end;


-- Function to return the trigger nest level
 function GetNestLevel return number is
 begin
 if NestLevel is null then
     NestLevel := 0;
 end if;
 return(NestLevel);
 end;

-- Procedure to increase the trigger nest level
 procedure NextNestLevel is
 begin
 if NestLevel is null then
     NestLevel := 0;
 end if;
 NestLevel := NestLevel + 1;
 end;

-- Procedure to decrease the trigger nest level
 procedure PreviousNestLevel is
 begin
 NestLevel := NestLevel - 1;
 end;

 end IntegrityPackage;
/


drop trigger "CompoundDeleteTrigger_class"
/

drop trigger "CompoundInsertTrigger_class"
/

drop trigger "CompoundUpdateTrigger_class"
/

drop trigger "tib_class"
/

drop trigger "CompoundDeleteTrigger_college"
/

drop trigger "CompoundInsertTrigger_college"
/

drop trigger "CompoundUpdateTrigger_college"
/

drop trigger "tib_college"
/

drop trigger "CompoundDeleteTrigger_course"
/

drop trigger "CompoundInsertTrigger_course"
/

drop trigger "CompoundUpdateTrigger_course"
/

drop trigger "tib_course"
/

drop trigger "CompoundDeleteTrigger_human"
/

drop trigger "CompoundInsertTrigger_human"
/

drop trigger "CompoundUpdateTrigger_human"
/

drop trigger "tib_human"
/

drop trigger "CompoundDeleteTrigger_specialt"
/

drop trigger "CompoundInsertTrigger_specialt"
/

drop trigger "CompoundUpdateTrigger_specialt"
/

drop trigger "tib_specialty"
/

alter table "Class"
   drop constraint FK_CLASS_CLASSBELO_SPECIALT
/

alter table "Class"
   drop constraint FK_CLASS_TEACHERMA_TEACHER
/

alter table "Course"
   drop constraint FK_COURSE_LEADINGTE_TEACHER
/

alter table "Specialty"
   drop constraint FK_SPECIALT_SPECIALTY_COLLEGE
/

alter table "Student"
   drop constraint FK_STUDENT_HUMANINHE_HUMAN
/

alter table "Student"
   drop constraint FK_STUDENT_STUDENTBE_CLASS
/

alter table "StudentAttendsCourse"
   drop constraint FK_STUDENTA_STUDENTAT_COURSE
/

alter table "StudentAttendsCourse"
   drop constraint FK_STUDENTA_STUDENTAT_STUDENT
/

alter table "Teacher"
   drop constraint FK_TEACHER_HUMANINHE_HUMAN
/

alter table "Teacher"
   drop constraint FK_TEACHER_TEACHERBE_COLLEGE
/

alter table "TeacherTeachsCourse"
   drop constraint FK_TEACHERT_TEACHERTE_TEACHER
/

alter table "TeacherTeachsCourse"
   drop constraint FK_TEACHERT_TEACHERTE_COURSE
/

drop table "Administrator" cascade constraints
/

drop index "ClassBelongsToSpecialty_FK"
/

drop index "TeacherMangeClass_FK"
/

drop table "Class" cascade constraints
/

drop table "College" cascade constraints
/

drop index "LeadingTeacherLeadsACourse_FK"
/

drop table "Course" cascade constraints
/

drop index "Identity_UK"
/

drop table "Human" cascade constraints
/

drop index "SpecialtyBelongsToCollege_FK"
/

drop table "Specialty" cascade constraints
/

drop index "StudentBelongsToClass_FK"
/

drop table "Student" cascade constraints
/

drop index "StudentAttendCourse2_FK"
/

drop index "StudentAttendCourse_FK"
/

drop table "StudentAttendsCourse" cascade constraints
/

drop index "TeacherBelongsToCollege_FK"
/

drop table "Teacher" cascade constraints
/

drop index "TeacherTeachsCourse2_FK"
/

drop index "TeacherTeachsCourse_FK"
/

drop table "TeacherTeachsCourse" cascade constraints
/

drop sequence "IDSequence"
/

create sequence "IDSequence"
increment by 1
start with 1
 nomaxvalue
nocycle
/

/*==============================================================*/
/* Table: "Administrator"                                       */
/*==============================================================*/
create table "Administrator" 
(
   "AdminLoginName"     NVARCHAR2(32)        not null,
   "AdminPassHash"      VARCHAR2(1024),
   constraint PK_ADMINISTRATOR primary key ("AdminLoginName")
)
/

/*==============================================================*/
/* Table: "Class"                                               */
/*==============================================================*/
create table "Class" 
(
   "ClassID"            INTEGER              not null,
   "SpecialtyID"        INTEGER              not null,
   "MasterTeacherHumanID" INTEGER,
   "Grade"              INTEGER              not null,
   "ClassCode"          INTEGER              not null,
   constraint PK_CLASS primary key ("ClassID")
)
/

/*==============================================================*/
/* Index: "TeacherMangeClass_FK"                                */
/*==============================================================*/
create index "TeacherMangeClass_FK" on "Class" (
   "MasterTeacherHumanID" ASC
)
/

/*==============================================================*/
/* Index: "ClassBelongsToSpecialty_FK"                          */
/*==============================================================*/
create index "ClassBelongsToSpecialty_FK" on "Class" (
   "SpecialtyID" ASC
)
/

/*==============================================================*/
/* Table: "College"                                             */
/*==============================================================*/
create table "College" 
(
   "CollegeID"          INTEGER              not null,
   "CollegeName"        NVARCHAR2(32)        not null,
   constraint PK_COLLEGE primary key ("CollegeID"),
   constraint AK_COLLEGENAME_COLLEGE unique ("CollegeName")
)
/

/*==============================================================*/
/* Table: "Course"                                              */
/*==============================================================*/
create table "Course" 
(
   "CourseID"           INTEGER              not null,
   "LeadTeacherHumanID" INTEGER,
   "CourseName"         NVARCHAR2(32),
   "CourseNumber"       INTEGER              not null,
   "Credits"            NUMBER(2,1),
   "CourseProperty"     INTEGER             
      constraint CKC_COURSEPROPERTY_COURSE check ("CourseProperty" is null or ("CourseProperty" between 1 and 4)),
   "Accommodate"        INTEGER,
   "Time"               CLOB,
   "Address"            NVARCHAR2(64),
   "RestrictClass"      CLOB,
   constraint PK_COURSE primary key ("CourseID"),
   constraint AK_COURSENUMBER_COURSE unique ("CourseNumber")
)
/

comment on column "Course"."CourseProperty" is
'1: 专业必修
2: 专业选修
3: 通识性选修
4: 体育选修'
/

/*==============================================================*/
/* Index: "LeadingTeacherLeadsACourse_FK"                       */
/*==============================================================*/
create index "LeadingTeacherLeadsACourse_FK" on "Course" (
   "LeadTeacherHumanID" ASC
)
/

/*==============================================================*/
/* Table: "Human"                                               */
/*==============================================================*/
create table "Human" 
(
   "HumanID"            INTEGER              not null,
   "Name"               NVARCHAR2(32),
   "Sex"                NCHAR(1)            
      constraint CKC_SEX_HUMAN check ("Sex" is null or ("Sex" in ('男','女'))),
   "Birthday"           DATE,
   "Identity"           CHAR(18)             not null,
   "Notes"              CLOB,
   "PasswordHash"       VARCHAR2(1024),
   constraint PK_HUMAN primary key ("HumanID"),
   constraint AK_IDENTITY_HUMAN unique ("Identity")
)
/

/*==============================================================*/
/* Index: "Identity_UK"                                         */
/*==============================================================*/
create unique index "Identity_UK" on "Human" (
   "Identity" ASC
)
/

/*==============================================================*/
/* Table: "Specialty"                                           */
/*==============================================================*/
create table "Specialty" 
(
   "SpecialtyID"        INTEGER              not null,
   "CollegeID"          INTEGER              not null,
   "SpecialtyName"      NVARCHAR2(32),
   constraint PK_SPECIALTY primary key ("SpecialtyID")
)
/

/*==============================================================*/
/* Index: "SpecialtyBelongsToCollege_FK"                        */
/*==============================================================*/
create index "SpecialtyBelongsToCollege_FK" on "Specialty" (
   "CollegeID" ASC
)
/

/*==============================================================*/
/* Table: "Student"                                             */
/*==============================================================*/
create table "Student" 
(
   "HumanID"            INTEGER              not null,
   "ClassID"            INTEGER              not null,
   "StudentNumber"      INTEGER              not null,
   "AdmissionDate"      DATE,
   "GraduationDate"     DATE,
   "StudentDegree"      NVARCHAR2(8)        
      constraint CKC_STUDENTDEGREE_STUDENT check ("StudentDegree" is null or ("StudentDegree" in ('学士','硕士','博士'))),
   "YearOfSchool"       INTEGER              default 4,
   "Status"             NVARCHAR2(8)        
      constraint CKC_STATUS_STUDENT check ("Status" is null or ("Status" in ('在读','毕业'))),
   constraint PK_STUDENT primary key ("HumanID"),
   constraint AK_STUDENTNUMBER_STUDENT unique ("StudentNumber")
)
/

comment on column "Student"."GraduationDate" is
'根据学制可设置成毕业年的9月份。'
/

/*==============================================================*/
/* Index: "StudentBelongsToClass_FK"                            */
/*==============================================================*/
create index "StudentBelongsToClass_FK" on "Student" (
   "ClassID" ASC
)
/

/*==============================================================*/
/* Table: "StudentAttendsCourse"                                */
/*==============================================================*/
create table "StudentAttendsCourse" 
(
   "CourseID"           INTEGER              not null,
   "StudentHumanID"     INTEGER              not null,
   "Score"              INTEGER,
   constraint PK_STUDENTATTENDSCOURSE primary key ("CourseID", "StudentHumanID")
)
/

/*==============================================================*/
/* Index: "StudentAttendCourse_FK"                              */
/*==============================================================*/
create index "StudentAttendCourse_FK" on "StudentAttendsCourse" (
   "CourseID" ASC
)
/

/*==============================================================*/
/* Index: "StudentAttendCourse2_FK"                             */
/*==============================================================*/
create index "StudentAttendCourse2_FK" on "StudentAttendsCourse" (
   "StudentHumanID" ASC
)
/

/*==============================================================*/
/* Table: "Teacher"                                             */
/*==============================================================*/
create table "Teacher" 
(
   "HumanID"            INTEGER              not null,
   "CollegeID"          INTEGER              not null,
   "TeacherNumber"      INTEGER              not null,
   "GraduationSchool"   NVARCHAR2(32),
   "Position"           NVARCHAR2(8)        
      constraint CKC_POSITION_TEACHER check ("Position" is null or ("Position" in ('教务办主任','普通教师'))),
   "TeacherDegree"      NVARCHAR2(8)        
      constraint CKC_TEACHERDEGREE_TEACHER check ("TeacherDegree" is null or ("TeacherDegree" in ('学士','硕士','博士','博士后'))),
   constraint PK_TEACHER primary key ("HumanID"),
   constraint AK_TEACHERNUMBER_TEACHER unique ("TeacherNumber")
)
/

/*==============================================================*/
/* Index: "TeacherBelongsToCollege_FK"                          */
/*==============================================================*/
create index "TeacherBelongsToCollege_FK" on "Teacher" (
   "CollegeID" ASC
)
/

/*==============================================================*/
/* Table: "TeacherTeachsCourse"                                 */
/*==============================================================*/
create table "TeacherTeachsCourse" 
(
   "TeacherHumanID"     INTEGER              not null,
   "CourseID"           INTEGER              not null,
   constraint PK_TEACHERTEACHSCOURSE primary key ("TeacherHumanID", "CourseID")
)
/

/*==============================================================*/
/* Index: "TeacherTeachsCourse_FK"                              */
/*==============================================================*/
create index "TeacherTeachsCourse_FK" on "TeacherTeachsCourse" (
   "TeacherHumanID" ASC
)
/

/*==============================================================*/
/* Index: "TeacherTeachsCourse2_FK"                             */
/*==============================================================*/
create index "TeacherTeachsCourse2_FK" on "TeacherTeachsCourse" (
   "CourseID" ASC
)
/

alter table "Class"
   add constraint FK_CLASS_CLASSBELO_SPECIALT foreign key ("SpecialtyID")
      references "Specialty" ("SpecialtyID")
      on delete cascade
/

alter table "Class"
   add constraint FK_CLASS_TEACHERMA_TEACHER foreign key ("MasterTeacherHumanID")
      references "Teacher" ("HumanID")
      on delete cascade
/

alter table "Course"
   add constraint FK_COURSE_LEADINGTE_TEACHER foreign key ("LeadTeacherHumanID")
      references "Teacher" ("HumanID")
      on delete cascade
/

alter table "Specialty"
   add constraint FK_SPECIALT_SPECIALTY_COLLEGE foreign key ("CollegeID")
      references "College" ("CollegeID")
      on delete cascade
/

alter table "Student"
   add constraint FK_STUDENT_HUMANINHE_HUMAN foreign key ("HumanID")
      references "Human" ("HumanID")
      on delete cascade
/

alter table "Student"
   add constraint FK_STUDENT_STUDENTBE_CLASS foreign key ("ClassID")
      references "Class" ("ClassID")
      on delete cascade
/

alter table "StudentAttendsCourse"
   add constraint FK_STUDENTA_STUDENTAT_COURSE foreign key ("CourseID")
      references "Course" ("CourseID")
      on delete cascade
/

alter table "StudentAttendsCourse"
   add constraint FK_STUDENTA_STUDENTAT_STUDENT foreign key ("StudentHumanID")
      references "Student" ("HumanID")
      on delete cascade
/

alter table "Teacher"
   add constraint FK_TEACHER_HUMANINHE_HUMAN foreign key ("HumanID")
      references "Human" ("HumanID")
      on delete cascade
/

alter table "Teacher"
   add constraint FK_TEACHER_TEACHERBE_COLLEGE foreign key ("CollegeID")
      references "College" ("CollegeID")
      on delete cascade
/

alter table "TeacherTeachsCourse"
   add constraint FK_TEACHERT_TEACHERTE_TEACHER foreign key ("TeacherHumanID")
      references "Teacher" ("HumanID")
      on delete cascade
/

alter table "TeacherTeachsCourse"
   add constraint FK_TEACHERT_TEACHERTE_COURSE foreign key ("CourseID")
      references "Course" ("CourseID")
      on delete cascade
/


create or replace trigger "CompoundDeleteTrigger_class"
for delete on "Class" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundInsertTrigger_class"
for insert on "Class" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundUpdateTrigger_class"
for update on "Class" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create trigger "tib_class" before insert
on "Class" for each row
declare
    integrity_error  exception;
    errno            integer;
    errmsg           char(200);
    dummy            integer;
    found            boolean;

begin
    --  Column ""ClassID"" uses sequence IDSequence
    select "IDSequence".NEXTVAL INTO :new."ClassID" from dual;

--  Errors handling
exception
    when integrity_error then
       raise_application_error(errno, errmsg);
end;
/


create or replace trigger "CompoundDeleteTrigger_college"
for delete on "College" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundInsertTrigger_college"
for insert on "College" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundUpdateTrigger_college"
for update on "College" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create trigger "tib_college" before insert
on "College" for each row
declare
    integrity_error  exception;
    errno            integer;
    errmsg           char(200);
    dummy            integer;
    found            boolean;

begin
    --  Column ""CollegeID"" uses sequence IDSequence
    select "IDSequence".NEXTVAL INTO :new."CollegeID" from dual;

--  Errors handling
exception
    when integrity_error then
       raise_application_error(errno, errmsg);
end;
/


create or replace trigger "CompoundDeleteTrigger_course"
for delete on "Course" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundInsertTrigger_course"
for insert on "Course" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundUpdateTrigger_course"
for update on "Course" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create trigger "tib_course" before insert
on "Course" for each row
declare
    integrity_error  exception;
    errno            integer;
    errmsg           char(200);
    dummy            integer;
    found            boolean;

begin
    --  Column ""CourseID"" uses sequence IDSequence
    select "IDSequence".NEXTVAL INTO :new."CourseID" from dual;

--  Errors handling
exception
    when integrity_error then
       raise_application_error(errno, errmsg);
end;
/


create or replace trigger "CompoundDeleteTrigger_human"
for delete on "Human" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundInsertTrigger_human"
for insert on "Human" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundUpdateTrigger_human"
for update on "Human" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create trigger "tib_human" before insert
on "Human" for each row
declare
    integrity_error  exception;
    errno            integer;
    errmsg           char(200);
    dummy            integer;
    found            boolean;

begin
    --  Column ""HumanID"" uses sequence IDSequence
    select "IDSequence".NEXTVAL INTO :new."HumanID" from dual;

--  Errors handling
exception
    when integrity_error then
       raise_application_error(errno, errmsg);
end;
/


create or replace trigger "CompoundDeleteTrigger_specialt"
for delete on "Specialty" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundInsertTrigger_specialt"
for insert on "Specialty" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create or replace trigger "CompoundUpdateTrigger_specialt"
for update on "Specialty" compound trigger
-- Declaration
-- Body
  before statement is
  begin
     NULL;
  end before statement;

  before each row is
  begin
     NULL;
  end before each row;

  after each row is
  begin
     NULL;
  end after each row;

  after statement is
  begin
     NULL;
  end after statement;
END;
/


create trigger "tib_specialty" before insert
on "Specialty" for each row
declare
    integrity_error  exception;
    errno            integer;
    errmsg           char(200);
    dummy            integer;
    found            boolean;

begin
    --  Column ""SpecialtyID"" uses sequence IDSequence
    select "IDSequence".NEXTVAL INTO :new."SpecialtyID" from dual;

--  Errors handling
exception
    when integrity_error then
       raise_application_error(errno, errmsg);
end;
/

