/*==============================================================*/
/* DBMS name:      ORACLE Version 11g                           */
/* Created on:     2019/6/11 19:31:01                           */
/*==============================================================*/


alter table "Class"
   drop constraint FK_CLASS_TEACHERMA_TEACHER;

alter table "Course"
   drop constraint FK_COURSE_COURSEOPE_SEMESTER;

alter table "Course"
   drop constraint FK_COURSE_LEADINGTE_TEACHER;

alter table "CourseProgram"
   drop constraint FK_COURSEPR_COURSEHAS_COURSE;

alter table "CourseProgram"
   drop constraint FK_COURSEPR_COURSETAK_CLASSROO;

alter table "Semester"
   drop constraint FK_SEMESTER_ACADEMICY_ACADEMIC;

alter table "Student"
   drop constraint FK_STUDENT_HUMANINHE_HUMAN;

alter table "Student"
   drop constraint FK_STUDENT_STUDENTBE_CLASS;

alter table "Student"
   drop constraint FK_STUDENT_STUDENTBE_COLLEGE;

alter table "Student"
   drop constraint FK_STUDENT_STUDENTHA_SPECIALT;

alter table "StudentAttendCourse"
   drop constraint FK_STUDENTA_STUDENTAT_COURSE;

alter table "StudentAttendCourse"
   drop constraint FK_STUDENTA_STUDENTAT_STUDENT;

alter table "Teacher"
   drop constraint FK_TEACHER_HUMANINHE_HUMAN;

alter table "Teacher"
   drop constraint FK_TEACHER_TEACHERBE_COLLEGE;

alter table "Teacher"
   drop constraint FK_TEACHER_TEACHERHA_SPECIALT;

alter table "TeacherTeachsCourse"
   drop constraint FK_TEACHERT_TEACHERTE_TEACHER;

alter table "TeacherTeachsCourse"
   drop constraint FK_TEACHERT_TEACHERTE_COURSE;

drop table "AcademicYear" cascade constraints;

drop table "Administrator" cascade constraints;

drop index "TeacherMangeClass_FK";

drop table "Class" cascade constraints;

drop table "ClassRoom" cascade constraints;

drop table "College" cascade constraints;

drop index "CourseOpensAtSemester_FK";

drop index "LeadingTeacherLeadsACourse_FK";

drop table "Course" cascade constraints;

drop index "CoureTakesPlaceInClassroom_FK";

drop index "CourseHasCourseProgram_FK";

drop table "CourseProgram" cascade constraints;

drop table "Human" cascade constraints;

drop index "AcademicYearHasSemestes_FK";

drop table "Semester" cascade constraints;

drop table "Specialty" cascade constraints;

drop index "StudentBelongsToClass_FK";

drop index "StudentHasProfessional_FK";

drop index "StudentBelongsToCollege_FK";

drop table "Student" cascade constraints;

drop index "StudentAttendCourse2_FK";

drop index "StudentAttendCourse_FK";

drop table "StudentAttendCourse" cascade constraints;

drop index "TeacherHasProfessional_FK";

drop index "TeacherBelongsToAcademy_FK";

drop table "Teacher" cascade constraints;

drop index "TeacherTeachCourse2_FK";

drop index "TeacherTeachCourse_FK";

drop table "TeacherTeachsCourse" cascade constraints;

/*==============================================================*/
/* Table: "AcademicYear"                                        */
/*==============================================================*/
create table "AcademicYear" 
(
   "AcademicYear"       INTEGER              not null,
   constraint PK_ACADEMICYEAR primary key ("AcademicYear")
);

/*==============================================================*/
/* Table: "Administrator"                                       */
/*==============================================================*/
create table "Administrator" 
(
   "AdminLoginName"     NVARCHAR2(32)        not null,
   "AdminPassHash"      VARCHAR2(1024),
   constraint PK_ADMINISTRATOR primary key ("AdminLoginName")
);

/*==============================================================*/
/* Table: "Class"                                               */
/*==============================================================*/
create table "Class" 
(
   "ClassID"            INTEGER              not null,
   "HeadTeacherNumber"  INTEGER,
   constraint PK_CLASS primary key ("ClassID")
);

/*==============================================================*/
/* Index: "TeacherMangeClass_FK"                                */
/*==============================================================*/
create index "TeacherMangeClass_FK" on "Class" (
   "HeadTeacherNumber" ASC
);

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
);

/*==============================================================*/
/* Table: "College"                                             */
/*==============================================================*/
create table "College" 
(
   "CollegeID"          INTEGER              not null,
   "CollegeName"        NVARCHAR2(32),
   constraint PK_COLLEGE primary key ("CollegeID")
);

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
);

comment on column "Course"."CourseProperty" is
'1: 专业必修
2: 专业选修
3: 通识性选修
4: 体育选修';

/*==============================================================*/
/* Index: "LeadingTeacherLeadsACourse_FK"                       */
/*==============================================================*/
create index "LeadingTeacherLeadsACourse_FK" on "Course" (
   "LeadTeacherNumber" ASC
);

/*==============================================================*/
/* Index: "CourseOpensAtSemester_FK"                            */
/*==============================================================*/
create index "CourseOpensAtSemester_FK" on "Course" (
   "SmesterID" ASC
);

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
);

/*==============================================================*/
/* Index: "CourseHasCourseProgram_FK"                           */
/*==============================================================*/
create index "CourseHasCourseProgram_FK" on "CourseProgram" (
   "CourseID" ASC
);

/*==============================================================*/
/* Index: "CoureTakesPlaceInClassroom_FK"                       */
/*==============================================================*/
create index "CoureTakesPlaceInClassroom_FK" on "CourseProgram" (
   "ClassroomID" ASC
);

/*==============================================================*/
/* Table: "Human"                                               */
/*==============================================================*/
create table "Human" 
(
   "HumanID"            INTEGER              not null,
   "Name"               NVARCHAR2(32)        not null,
   "Sex"                NCHAR(1)             not null
      constraint CKC_SEX_HUMAN check ("Sex" in ('男','女')),
   "Birthday"           DATE,
   "Identity"           CHAR(18),
   "Notes"              CLOB,
   "PasswordHash"       VARCHAR2(1024),
   constraint PK_HUMAN primary key ("HumanID")
);

/*==============================================================*/
/* Table: "Semester"                                            */
/*==============================================================*/
create table "Semester" 
(
   "SmesterID"          INTEGER              not null,
   "AcademicYear"       INTEGER              not null,
   "SmesterCode"        INTEGER,
   constraint PK_SEMESTER primary key ("SmesterID")
);

comment on column "Semester"."SmesterCode" is
'1: 春季学期
2: 秋季学期';

/*==============================================================*/
/* Index: "AcademicYearHasSemestes_FK"                          */
/*==============================================================*/
create index "AcademicYearHasSemestes_FK" on "Semester" (
   "AcademicYear" ASC
);

/*==============================================================*/
/* Table: "Specialty"                                           */
/*==============================================================*/
create table "Specialty" 
(
   "SpecialtyID"        INTEGER              not null,
   "SpecialtyName"      NVARCHAR2(32),
   constraint PK_SPECIALTY primary key ("SpecialtyID")
);

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
   constraint PK_STUDENT primary key ("HumanID"),
   constraint AK_STUDENTNUMBER_STUDENT unique ("StudentNumber")
);

/*==============================================================*/
/* Index: "StudentBelongsToCollege_FK"                          */
/*==============================================================*/
create index "StudentBelongsToCollege_FK" on "Student" (
   "CollegeID" ASC
);

/*==============================================================*/
/* Index: "StudentHasProfessional_FK"                           */
/*==============================================================*/
create index "StudentHasProfessional_FK" on "Student" (
   "SpecialtyID" ASC
);

/*==============================================================*/
/* Index: "StudentBelongsToClass_FK"                            */
/*==============================================================*/
create index "StudentBelongsToClass_FK" on "Student" (
   "ClassID" ASC
);

/*==============================================================*/
/* Table: "StudentAttendCourse"                                 */
/*==============================================================*/
create table "StudentAttendCourse" 
(
   "CourseID"           INTEGER              not null,
   "StudentNumber"      INTEGER              not null,
   "Score"              INTEGER,
   constraint PK_STUDENTATTENDCOURSE primary key ("StudentNumber", "CourseID")
);

/*==============================================================*/
/* Index: "StudentAttendCourse_FK"                              */
/*==============================================================*/
create index "StudentAttendCourse_FK" on "StudentAttendCourse" (
   "CourseID" ASC
);

/*==============================================================*/
/* Index: "StudentAttendCourse2_FK"                             */
/*==============================================================*/
create index "StudentAttendCourse2_FK" on "StudentAttendCourse" (
   "StudentNumber" ASC
);

/*==============================================================*/
/* Table: "Teacher"                                             */
/*==============================================================*/
create table "Teacher" 
(
   "HumanID"            INTEGER              not null,
   "SpecialtyID"        INTEGER,
   "CollegeID"          INTEGER              not null,
   "TeacherNumber"      INTEGER              not null,
   constraint PK_TEACHER primary key ("HumanID"),
   constraint AK_TEACHERNUMBER_TEACHER unique ("TeacherNumber")
);

/*==============================================================*/
/* Index: "TeacherBelongsToAcademy_FK"                          */
/*==============================================================*/
create index "TeacherBelongsToAcademy_FK" on "Teacher" (
   "CollegeID" ASC
);

/*==============================================================*/
/* Index: "TeacherHasProfessional_FK"                           */
/*==============================================================*/
create index "TeacherHasProfessional_FK" on "Teacher" (
   "SpecialtyID" ASC
);

/*==============================================================*/
/* Table: "TeacherTeachsCourse"                                 */
/*==============================================================*/
create table "TeacherTeachsCourse" 
(
   "CourseID"           INTEGER              not null,
   "TeacherNumber"      INTEGER              not null,
   constraint PK_TEACHERTEACHSCOURSE primary key ("TeacherNumber", "CourseID")
);

/*==============================================================*/
/* Index: "TeacherTeachCourse_FK"                               */
/*==============================================================*/
create index "TeacherTeachCourse_FK" on "TeacherTeachsCourse" (
   "TeacherNumber" ASC
);

/*==============================================================*/
/* Index: "TeacherTeachCourse2_FK"                              */
/*==============================================================*/
create index "TeacherTeachCourse2_FK" on "TeacherTeachsCourse" (
   "CourseID" ASC
);

alter table "Class"
   add constraint FK_CLASS_TEACHERMA_TEACHER foreign key ("HeadTeacherNumber")
      references "Teacher" ("TeacherNumber");

alter table "Course"
   add constraint FK_COURSE_COURSEOPE_SEMESTER foreign key ("SmesterID")
      references "Semester" ("SmesterID");

alter table "Course"
   add constraint FK_COURSE_LEADINGTE_TEACHER foreign key ("LeadTeacherNumber")
      references "Teacher" ("TeacherNumber");

alter table "CourseProgram"
   add constraint FK_COURSEPR_COURSEHAS_COURSE foreign key ("CourseID")
      references "Course" ("CourseID");

alter table "CourseProgram"
   add constraint FK_COURSEPR_COURSETAK_CLASSROO foreign key ("ClassroomID")
      references "ClassRoom" ("ClassroomID");

alter table "Semester"
   add constraint FK_SEMESTER_ACADEMICY_ACADEMIC foreign key ("AcademicYear")
      references "AcademicYear" ("AcademicYear");

alter table "Student"
   add constraint FK_STUDENT_HUMANINHE_HUMAN foreign key ("HumanID")
      references "Human" ("HumanID");

alter table "Student"
   add constraint FK_STUDENT_STUDENTBE_CLASS foreign key ("ClassID")
      references "Class" ("ClassID");

alter table "Student"
   add constraint FK_STUDENT_STUDENTBE_COLLEGE foreign key ("CollegeID")
      references "College" ("CollegeID");

alter table "Student"
   add constraint FK_STUDENT_STUDENTHA_SPECIALT foreign key ("SpecialtyID")
      references "Specialty" ("SpecialtyID");

alter table "StudentAttendCourse"
   add constraint FK_STUDENTA_STUDENTAT_COURSE foreign key ("CourseID")
      references "Course" ("CourseID");

alter table "StudentAttendCourse"
   add constraint FK_STUDENTA_STUDENTAT_STUDENT foreign key ("StudentNumber")
      references "Student" ("StudentNumber");

alter table "Teacher"
   add constraint FK_TEACHER_HUMANINHE_HUMAN foreign key ("HumanID")
      references "Human" ("HumanID");

alter table "Teacher"
   add constraint FK_TEACHER_TEACHERBE_COLLEGE foreign key ("CollegeID")
      references "College" ("CollegeID");

alter table "Teacher"
   add constraint FK_TEACHER_TEACHERHA_SPECIALT foreign key ("SpecialtyID")
      references "Specialty" ("SpecialtyID");

alter table "TeacherTeachsCourse"
   add constraint FK_TEACHERT_TEACHERTE_TEACHER foreign key ("TeacherNumber")
      references "Teacher" ("TeacherNumber");

alter table "TeacherTeachsCourse"
   add constraint FK_TEACHERT_TEACHERTE_COURSE foreign key ("CourseID")
      references "Course" ("CourseID");

