/* Table Administrator */
-- create
insert into "Administrator" ("AdminLoginName", "AdminPassHash") values (:1, :2);
-- read
select count(*) from "Administrator";
select count(*) from "Administrator" where "AdminLoginName"=:1;
select "AdminLoginName", "AdminPassHash" from "Administrator" where "AdminLoginName"=:1;
-- delete
delete from "Administrator" where "AdminLoginName"=:1;


/* Table Class */
-- create
insert into "Class" ("SpecialtyID", "MasterTeacherHumanID", "Grade", "ClassCode")
values (:1, :2, :3, :4);
-- read
select count(*) from "Class" where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=:3;
select ("ClassID", "SpecialtyID", "MasterTeacherHumanID", "Grade", "ClassCode") from "Class";
-- update
update "Class" set "MasterTeacherHumanID"=:4
where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=3;
-- delete
delete from "Class" where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=:3;


/* Table College */
-- create
insert into "College" ("CollegeName") values (:1);
-- read
select count(*) from "College" where "CollegeName"=:1;
select "CollegeID", "CollegeName" from "College";
select "CollegeID", "CollegeName" from "College" where "CollegeID"=:1;
select "CollegeID", "CollegeName" from "College" where "CollegeName"=:1;
-- update
update "College" set "CollegeName"=:2 where "CollegeName"=:1;
-- delete
delete from "College" where "CollegeName"=:1;


/* Table Course */
-- create
insert into "Course" ("CourseName", "CourseNumber", "Credits", "CourseProperty", "Accommodate", "Time", "Address", "RestrictClass")
values (:1, :2, :3, :4, :5, :6, :7, :8);
-- read
select "CourseID", "LeadTeacherHumanID", "CourseName", "CourseNumber", "Credits", "CourseProperty", "Accommodate", "Time", "Address", "RestrictClass" from "Course";
select count(*) from "Course" where "CourseNumber"=:1;
select "CourseID", "LeadTeacherHumanID", "CourseName", "CourseNumber", "Credits", "CourseProperty", "Accommodate", "Time", "Address", "RestrictClass" 
from "Course" where "CourseNumber"=:1;
select "CourseID" from "TeacherTeachsCourse" where "TeacherHumanID"=:1;
select "CourseID", "LeadTeacherHumanID", "CourseName", "CourseNumber", "Credits", "CourseProperty", "Accommodate", "Time", "Address", "RestrictClass" 
from "Course" where "CourseID"=:1;
select "CourseID" from "StudentAttendsCourse" where "StudentHumanID"=:1;
select "CourseID", "LeadTeacherHumanID", "CourseName", "CourseNumber", "Credits", "CourseProperty", "Accommodate", "Time", "Address", "RestrictClass" 
from "Course" where "CourseID"=:1;
select "Score" from "StudentAttendsCourse" where "StudentHumanID"=:1;
select "StudentHumanID" from "StudentAttendsCourse" where "CourseID"=:1;
select "TeacherHumanID" from "TeacherTeachsCourse" where "CourseID"=:1;
-- update
update "Course" set "CourseName"=:1, "Credits"=:2, "CourseProperty"=:3, "Accommodate"=:4, "Time"=:5, "Address"=:6, "RestrictClass"=:7 where "CourseNumber"=:8;
update "Course" set "LeadTeacherHumanID"=:1 where "CourseNumber"=:2;
-- delete
delete from "StudentAttendsCourse" where "CourseID"=:1 and "StudentHumanID"=:2;
delete from "TeacherTeachsCourse" where "CourseID"=:1;
delete from "Course" where "CourseNumber"=:1;


/* Table Human */
-- create
insert into "Human" ("Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash")
values (:1, :2, :3, :4, :5, :6);
-- read
select count(*) from "Human" where "HumanID"=:1;
select count(*) from "Human" where "Identity"=:1;
select "HumanID" from "Human" where "Identity"=:1;
select "HumanID", "Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash" from "Human" where "HumanID"=:1;
-- update
update "Human" set "Name"=:1, "Sex"=:2, "Birthday"=:3, "Identity"=:4, "Notes"=:5, "PasswordHash"=:6 where "HumanID"=:7;
-- delete
delete from "Human" where "HumanID"=:1;


/* Table Specialty */
-- create
insert into "Specialty" ("CollegeID", "SpecialtyName") values (:1, :2);
-- read
select count(*) from "Specialty";
select count(*) from "Specialty" where "SepcialtyName"=:1;
select ("SpecialtyID", "CollegeID", "SpecailtyName") from "Specialty" where "SpecailtyName"=:1;
-- delete
delete from "Specialty" where "SpecailtyName"=:1;


/* Table Student */
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
delete from "Student" where "StudentNumber"=:1;


/* Table Teacher */
-- create
insert into "Teacher" ("HumanID", "CollegeID", "TeacherNumber", "GraduationSchool", "Position", "TeacherDegree")
values (:1, :2, :3, :4, :5, :6);
-- read
select "HumanID", "CollegeID", "TeacherNumber", "GraduationSchool", "Position", "TeacherDegree" from "Teacher" where "TeacherNumber"=:1;
-- update
update "Teacher" set "CollegeID"=:1, "GraduationSchool"=:2, "Position"=:3, "TeacherDegree"=:4 where "TeacherNumber"=:5;
-- delete
delete from "Teacher" where "TeacherNumber"=:1;


/* View ClassInfo */
-- read
select "CollegeID", "SpecialtyID", "ClassID", "CollegeName", "SpecialtyName", "Grade", "ClassCode", "TotalStudentCount" from "ClassInfo";


/* View StudentInfo */
-- read
select "HumanID", "CollegeID", "SpecialtyID", "ClassID", 
"CollegeName", "SpecialtyName", "Grade", "ClassCode", 
"Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash",
"StudentNumber", "AdmissionDate", "GraduationDate", "StudentDegree", "YearOfSchool", "Status"
from "StudentInfo"
where "StudentNumber"=:1


/* View TeacherInfo */
-- read
select "HumanID", "CollegeID", "CollegeName", "Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash", 
"TeacherNumber", "GraduationSchool", "Position", "TeacherDegree" from "TeacherInfo" where "Name"=:1;
