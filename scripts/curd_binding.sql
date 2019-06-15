/* AcademicYear */
-- create
insert into "AcademicYear" values (:1);
-- read
select count(*) from "AcademicYear";
select count(*) from "AcademicYear" where "AcademicYear"=:1;
-- update
update "AcademicYear" set "AcademicYear"=:2 where "AcademicYear"=:1;
-- delete
delete from "AcademicYear" where "AcademicYear"=:1;


/* Administrator */
-- create
insert into "Administrator" ("AdminLoginName", "AdminPassHash") values (:1, :2);
-- read
select count(*) from "Administrator";
select count(*) from "Administrator" where "AdminLoginName"=:1;


/* College */
-- create
insert into "College" ("CollegeName") values (:1);
-- read
select "CollegeID", "CollegeName" from "College";
-- update
update "College" set "CollegeName"=:2 where "CollegeName"=:1;
-- delete
delete from "College" where "CollegeName"=:1;


/* Specialty */
-- create
insert into "Sepcialty" ("CollegeID", "SepcialtyName") values (:1, :2);
-- read
select count(*) from "Specialty";
select count(*) from "Specialty" where "SepcialtyName"=:1;
select ("SepcialtyID", "CollegeID", "SepcailtyName") from "Sepcailty"
where "SepcailtyName"=:1;
-- delete
delete from "Specialty" where "SepcailtyName"=:1;


/* Class */
-- create
insert into "Class" ("SpecialtyID", "MasterTeacherHumanID", "Grade", "ClassCode")
values (:1, :2, :3, :4);
-- read
select ("ClassID", "SpecialtyID", "MasterTeacherHumanID", "Grade", "ClassCode") from "Class";
-- update
update "Class" set "MasterTeacherHumanID"=:4
where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=3;
-- delete
delete from "Class" where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=:3;


/* Student */
-- create
insert into "Student" ("HumanID", "ClassID", "StudentNumber", "AdmissionDate", "GraduationDate", "StudentDegree", "YearOfSchool", "Status")
values (:1, :2, :3, :4, :5, :6, :7, :8);
-- read
select "HumanID", "ClassID", "StudentNumber", "AdmissionDate", "GraduationDate", "StudentDegree", "YearOfSchool", "Status"
from "Student";

select "HumanID", "ClassID", "StudentNumber", "AdmissionDate", "GraduationDate", "StudentDegree", "YearOfSchool", "Status"
from "Student"
where "StudentNumber"=:1;
-- update
update "Student"
set "ClassID"=:2, "StudentNumber"=:3, "AdmissionDate"=:4, "GraduationDate"=:5, "StudentDegree"=:6, "YearOfSchool"=:7, "Status"=:8
where "StudentNumber"=:1;
-- delete


/* Teacher */
-- create
insert into "Teacher" ("HumanID", "CollegeID", "TeacherNumber", "GraduationSchool", "Position", "TeacherDegree")
values (:1, :2, :3, :4, :5, :6);
-- update
-- delete
