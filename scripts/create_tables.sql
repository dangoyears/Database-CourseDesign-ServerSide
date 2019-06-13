/*==============================================================*/
/* DBMS name:      ORACLE Version 11g                           */
/* Created on:     2019/6/13 14:35:14                           */
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

drop trigger "CompoundDeleteTrigger_classroo"
/

drop trigger "CompoundInsertTrigger_classroo"
/

drop trigger "CompoundUpdateTrigger_classroo"
/

drop trigger "tib_classroom"
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

drop trigger "CompoundDeleteTrigger_coursepr"
/

drop trigger "CompoundInsertTrigger_coursepr"
/

drop trigger "CompoundUpdateTrigger_coursepr"
/

drop trigger "tib_courseprogram"
/

drop trigger "CompoundDeleteTrigger_human"
/

drop trigger "CompoundInsertTrigger_human"
/

drop trigger "CompoundUpdateTrigger_human"
/

drop trigger "tib_human"
/

drop trigger "CompoundDeleteTrigger_semester"
/

drop trigger "CompoundInsertTrigger_semester"
/

drop trigger "CompoundUpdateTrigger_semester"
/

drop trigger "tib_semester"
/

drop trigger "CompoundDeleteTrigger_specialt"
/

drop trigger "CompoundInsertTrigger_specialt"
/

drop trigger "CompoundUpdateTrigger_specialt"
/

drop trigger "tib_specialty"
/

drop trigger "CompoundDeleteTrigger_student"
/

drop trigger "CompoundInsertTrigger_student"
/

drop trigger "CompoundUpdateTrigger_student"
/

drop trigger "tib_student"
/

drop trigger "CompoundDeleteTrigger_teacher"
/

drop trigger "CompoundInsertTrigger_teacher"
/

drop trigger "CompoundUpdateTrigger_teacher"
/

drop trigger "tib_teacher"
/

alter table "Class"
   drop constraint FK_CLASS_CLASSBELO_COLLEGE
/

alter table "Class"
   drop constraint FK_CLASS_TEACHERMA_TEACHER
/

alter table "Course"
   drop constraint FK_COURSE_COURSEOPE_SEMESTER
/

alter table "Course"
   drop constraint FK_COURSE_LEADINGTE_TEACHER
/

alter table "CourseProgram"
   drop constraint FK_COURSEPR_COURSEHAS_COURSE
/

alter table "CourseProgram"
   drop constraint FK_COURSEPR_COURSETAK_CLASSROO
/

alter table "Semester"
   drop constraint FK_SEMESTER_ACADEMICY_ACADEMIC
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

alter table "Student"
   drop constraint FK_STUDENT_STUDENTBE_COLLEGE
/

alter table "Student"
   drop constraint FK_STUDENT_STUDENTHA_SPECIALT
/

alter table "StudentAttendCourse"
   drop constraint FK_STUDENTA_STUDENTAT_COURSE
/

alter table "StudentAttendCourse"
   drop constraint FK_STUDENTA_STUDENTAT_STUDENT
/

alter table "Teacher"
   drop constraint FK_TEACHER_HUMANINHE_HUMAN
/

alter table "Teacher"
   drop constraint FK_TEACHER_TEACHERBE_COLLEGE
/

alter table "Teacher"
   drop constraint FK_TEACHER_TEACHERHA_SPECIALT
/

alter table "TeacherTeachsCourse"
   drop constraint FK_TEACHERT_TEACHERTE_TEACHER
/

alter table "TeacherTeachsCourse"
   drop constraint FK_TEACHERT_TEACHERTE_COURSE
/

drop table "AcademicYear" cascade constraints
/

drop table "Administrator" cascade constraints
/

drop index "ClassBelongsToCollege_FK"
/

drop index "TeacherMangeClass_FK"
/

drop table "Class" cascade constraints
/

drop table "ClassRoom" cascade constraints
/

drop table "College" cascade constraints
/

drop index "CourseOpensAtSemester_FK"
/

drop index "LeadingTeacherLeadsACourse_FK"
/

drop table "Course" cascade constraints
/

drop index "CourseTakesPlaceInClassroom_FK"
/

drop index "CourseHasCourseProgram_FK"
/

drop table "CourseProgram" cascade constraints
/

drop table "Human" cascade constraints
/

drop index "AcademicYearHasSemestes_FK"
/

drop table "Semester" cascade constraints
/

drop index "SpecialtyBelongsToCollege_FK"
/

drop table "Specialty" cascade constraints
/

drop index "StudentBelongsToClass_FK"
/

drop index "StudentHasSpecialty_FK"
/

drop index "StudentBelongsToCollege_FK"
/

drop table "Student" cascade constraints
/

drop index "StudentAttendCourse2_FK"
/

drop index "StudentAttendCourse_FK"
/

drop table "StudentAttendCourse" cascade constraints
/

drop index "TeacherHasSpecialty_FK"
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
start with 0
 nomaxvalue
 nominvalue
/

/*==============================================================*/
/* Table: "AcademicYear"                                        */
/*==============================================================*/
create table "AcademicYear" 
(
   "AcademicYear"       INTEGER              not null,
   constraint PK_ACADEMICYEAR primary key ("AcademicYear")
)
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
   "CollegeID"          INTEGER              not null,
   "MasterTeacherNumber" INTEGER,
   "Grade"              INTEGER              not null,
   "Class"              INTEGER              not null,
   constraint PK_CLASS primary key ("ClassID")
)
/

/*==============================================================*/
/* Index: "TeacherMangeClass_FK"                                */
/*==============================================================*/
create index "TeacherMangeClass_FK" on "Class" (
   "MasterTeacherNumber" ASC
)
/

/*==============================================================*/
/* Index: "ClassBelongsToCollege_FK"                            */
/*==============================================================*/
create index "ClassBelongsToCollege_FK" on "Class" (
   "CollegeID" ASC
)
/

/*==============================================================*/
/* Table: "ClassRoom"                                           */
/*==============================================================*/
create table "ClassRoom" 
(
   "ClassroomID"        INTEGER              not null,
   "Location"           NVARCHAR2(64)        not null,
   "Capacity"           INTEGER,
   constraint PK_CLASSROOM primary key ("ClassroomID"),
   constraint AK_LOCATION_CLASSROO unique ("Location")
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
   "SmesterID"          INTEGER              not null,
   "LeadTeacherNumber"  INTEGER,
   "CourseName"         NVARCHAR2(32),
   "Credits"            NUMBER(1,1),
   "CourseProperty"     INTEGER             
      constraint CKC_COURSEPROPERTY_COURSE check ("CourseProperty" is null or ("CourseProperty" between 1 and 4)),
   constraint PK_COURSE primary key ("CourseID")
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
   "LeadTeacherNumber" ASC
)
/

/*==============================================================*/
/* Index: "CourseOpensAtSemester_FK"                            */
/*==============================================================*/
create index "CourseOpensAtSemester_FK" on "Course" (
   "SmesterID" ASC
)
/

/*==============================================================*/
/* Table: "CourseProgram"                                       */
/*==============================================================*/
create table "CourseProgram" 
(
   "CourseProgramID"    INTEGER              not null,
   "CourseID"           INTEGER,
   "ClassroomID"        INTEGER              not null,
   "Week"               INTEGER,
   "Weekfay"            INTEGER,
   "Section"            INTEGER,
   constraint PK_COURSEPROGRAM primary key ("CourseProgramID")
)
/

/*==============================================================*/
/* Index: "CourseHasCourseProgram_FK"                           */
/*==============================================================*/
create index "CourseHasCourseProgram_FK" on "CourseProgram" (
   "CourseID" ASC
)
/

/*==============================================================*/
/* Index: "CourseTakesPlaceInClassroom_FK"                      */
/*==============================================================*/
create index "CourseTakesPlaceInClassroom_FK" on "CourseProgram" (
   "ClassroomID" ASC
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
   "Identity"           CHAR(18),
   "Notes"              CLOB,
   "PasswordHash"       VARCHAR2(1024),
   constraint PK_HUMAN primary key ("HumanID")
)
/

/*==============================================================*/
/* Table: "Semester"                                            */
/*==============================================================*/
create table "Semester" 
(
   "SmesterID"          INTEGER              not null,
   "AcademicYear"       INTEGER              not null,
   "SmesterCode"        INTEGER,
   constraint PK_SEMESTER primary key ("SmesterID")
)
/

comment on column "Semester"."SmesterCode" is
'1: 春季学期
2: 秋季学期'
/

/*==============================================================*/
/* Index: "AcademicYearHasSemestes_FK"                          */
/*==============================================================*/
create index "AcademicYearHasSemestes_FK" on "Semester" (
   "AcademicYear" ASC
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
   "CollegeID"          INTEGER              not null,
   "SpecialtyID"        INTEGER              not null,
   "StudentNumber"      INTEGER              not null,
   "Enrollment"         DATE,
   constraint PK_STUDENT primary key ("StudentNumber"),
   constraint AK_HUMANID_STUDENT unique ("HumanID")
)
/

/*==============================================================*/
/* Index: "StudentBelongsToCollege_FK"                          */
/*==============================================================*/
create index "StudentBelongsToCollege_FK" on "Student" (
   "CollegeID" ASC
)
/

/*==============================================================*/
/* Index: "StudentHasSpecialty_FK"                              */
/*==============================================================*/
create index "StudentHasSpecialty_FK" on "Student" (
   "SpecialtyID" ASC
)
/

/*==============================================================*/
/* Index: "StudentBelongsToClass_FK"                            */
/*==============================================================*/
create index "StudentBelongsToClass_FK" on "Student" (
   "ClassID" ASC
)
/

/*==============================================================*/
/* Table: "StudentAttendCourse"                                 */
/*==============================================================*/
create table "StudentAttendCourse" 
(
   "CourseID"           INTEGER              not null,
   "StudentNumber"      INTEGER              not null,
   "Score"              INTEGER,
   constraint PK_STUDENTATTENDCOURSE primary key ("StudentNumber", "CourseID")
)
/

/*==============================================================*/
/* Index: "StudentAttendCourse_FK"                              */
/*==============================================================*/
create index "StudentAttendCourse_FK" on "StudentAttendCourse" (
   "CourseID" ASC
)
/

/*==============================================================*/
/* Index: "StudentAttendCourse2_FK"                             */
/*==============================================================*/
create index "StudentAttendCourse2_FK" on "StudentAttendCourse" (
   "StudentNumber" ASC
)
/

/*==============================================================*/
/* Table: "Teacher"                                             */
/*==============================================================*/
create table "Teacher" 
(
   "HumanID"            INTEGER              not null,
   "SpecialtyID"        INTEGER,
   "CollegeID"          INTEGER              not null,
   "TeacherNumber"      INTEGER              not null,
   constraint PK_TEACHER primary key ("TeacherNumber"),
   constraint AK_HUMANID_TEACHER unique ("HumanID")
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
/* Index: "TeacherHasSpecialty_FK"                              */
/*==============================================================*/
create index "TeacherHasSpecialty_FK" on "Teacher" (
   "SpecialtyID" ASC
)
/

/*==============================================================*/
/* Table: "TeacherTeachsCourse"                                 */
/*==============================================================*/
create table "TeacherTeachsCourse" 
(
   "CourseID"           INTEGER              not null,
   "TeacherNumber"      INTEGER              not null,
   constraint PK_TEACHERTEACHSCOURSE primary key ("TeacherNumber", "CourseID")
)
/

/*==============================================================*/
/* Index: "TeacherTeachsCourse_FK"                              */
/*==============================================================*/
create index "TeacherTeachsCourse_FK" on "TeacherTeachsCourse" (
   "TeacherNumber" ASC
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
   add constraint FK_CLASS_CLASSBELO_COLLEGE foreign key ("CollegeID")
      references "College" ("CollegeID")
/

alter table "Class"
   add constraint FK_CLASS_TEACHERMA_TEACHER foreign key ("MasterTeacherNumber")
      references "Teacher" ("TeacherNumber")
/

alter table "Course"
   add constraint FK_COURSE_COURSEOPE_SEMESTER foreign key ("SmesterID")
      references "Semester" ("SmesterID")
/

alter table "Course"
   add constraint FK_COURSE_LEADINGTE_TEACHER foreign key ("LeadTeacherNumber")
      references "Teacher" ("TeacherNumber")
/

alter table "CourseProgram"
   add constraint FK_COURSEPR_COURSEHAS_COURSE foreign key ("CourseID")
      references "Course" ("CourseID")
/

alter table "CourseProgram"
   add constraint FK_COURSEPR_COURSETAK_CLASSROO foreign key ("ClassroomID")
      references "ClassRoom" ("ClassroomID")
/

alter table "Semester"
   add constraint FK_SEMESTER_ACADEMICY_ACADEMIC foreign key ("AcademicYear")
      references "AcademicYear" ("AcademicYear")
/

alter table "Specialty"
   add constraint FK_SPECIALT_SPECIALTY_COLLEGE foreign key ("CollegeID")
      references "College" ("CollegeID")
/

alter table "Student"
   add constraint FK_STUDENT_HUMANINHE_HUMAN foreign key ("HumanID")
      references "Human" ("HumanID")
/

alter table "Student"
   add constraint FK_STUDENT_STUDENTBE_CLASS foreign key ("ClassID")
      references "Class" ("ClassID")
/

alter table "Student"
   add constraint FK_STUDENT_STUDENTBE_COLLEGE foreign key ("CollegeID")
      references "College" ("CollegeID")
/

alter table "Student"
   add constraint FK_STUDENT_STUDENTHA_SPECIALT foreign key ("SpecialtyID")
      references "Specialty" ("SpecialtyID")
/

alter table "StudentAttendCourse"
   add constraint FK_STUDENTA_STUDENTAT_COURSE foreign key ("CourseID")
      references "Course" ("CourseID")
/

alter table "StudentAttendCourse"
   add constraint FK_STUDENTA_STUDENTAT_STUDENT foreign key ("StudentNumber")
      references "Student" ("StudentNumber")
/

alter table "Teacher"
   add constraint FK_TEACHER_HUMANINHE_HUMAN foreign key ("HumanID")
      references "Human" ("HumanID")
/

alter table "Teacher"
   add constraint FK_TEACHER_TEACHERBE_COLLEGE foreign key ("CollegeID")
      references "College" ("CollegeID")
/

alter table "Teacher"
   add constraint FK_TEACHER_TEACHERHA_SPECIALT foreign key ("SpecialtyID")
      references "Specialty" ("SpecialtyID")
/

alter table "TeacherTeachsCourse"
   add constraint FK_TEACHERT_TEACHERTE_TEACHER foreign key ("TeacherNumber")
      references "Teacher" ("TeacherNumber")
/

alter table "TeacherTeachsCourse"
   add constraint FK_TEACHERT_TEACHERTE_COURSE foreign key ("CourseID")
      references "Course" ("CourseID")
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

END
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

END
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

END
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


create or replace trigger "CompoundDeleteTrigger_classroo"
for delete on "ClassRoom" compound trigger
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

END
/


create or replace trigger "CompoundInsertTrigger_classroo"
for insert on "ClassRoom" compound trigger
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

END
/


create or replace trigger "CompoundUpdateTrigger_classroo"
for update on "ClassRoom" compound trigger
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

END
/


create trigger "tib_classroom" before insert
on "ClassRoom" for each row
declare
    integrity_error  exception;
    errno            integer;
    errmsg           char(200);
    dummy            integer;
    found            boolean;

begin
    --  Column ""ClassroomID"" uses sequence IDSequence
    select "IDSequence".NEXTVAL INTO :new."ClassroomID" from dual;

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

END
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

END
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

END
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

END
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

END
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

END
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


create or replace trigger "CompoundDeleteTrigger_coursepr"
for delete on "CourseProgram" compound trigger
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

END
/


create or replace trigger "CompoundInsertTrigger_coursepr"
for insert on "CourseProgram" compound trigger
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

END
/


create or replace trigger "CompoundUpdateTrigger_coursepr"
for update on "CourseProgram" compound trigger
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

END
/


create trigger "tib_courseprogram" before insert
on "CourseProgram" for each row
declare
    integrity_error  exception;
    errno            integer;
    errmsg           char(200);
    dummy            integer;
    found            boolean;

begin
    --  Column ""CourseProgramID"" uses sequence IDSequence
    select "IDSequence".NEXTVAL INTO :new."CourseProgramID" from dual;

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

END
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

END
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

END
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


create or replace trigger "CompoundDeleteTrigger_semester"
for delete on "Semester" compound trigger
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

END
/


create or replace trigger "CompoundInsertTrigger_semester"
for insert on "Semester" compound trigger
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

END
/


create or replace trigger "CompoundUpdateTrigger_semester"
for update on "Semester" compound trigger
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

END
/


create trigger "tib_semester" before insert
on "Semester" for each row
declare
    integrity_error  exception;
    errno            integer;
    errmsg           char(200);
    dummy            integer;
    found            boolean;

begin
    --  Column ""SmesterID"" uses sequence IDSequence
    select "IDSequence".NEXTVAL INTO :new."SmesterID" from dual;

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

END
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

END
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

END
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


create or replace trigger "CompoundDeleteTrigger_student"
for delete on "Student" compound trigger
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

END
/


create or replace trigger "CompoundInsertTrigger_student"
for insert on "Student" compound trigger
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

END
/


create or replace trigger "CompoundUpdateTrigger_student"
for update on "Student" compound trigger
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

END
/

create or replace trigger "CompoundDeleteTrigger_teacher"
for delete on "Teacher" compound trigger
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

END
/


create or replace trigger "CompoundInsertTrigger_teacher"
for insert on "Teacher" compound trigger
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

END
/


create or replace trigger "CompoundUpdateTrigger_teacher"
for update on "Teacher" compound trigger
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

END
/