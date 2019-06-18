create or replace view "ClassInfo" as
select "College"."CollegeID", "Specialty"."SpecialtyID", "Class"."ClassID", 
"CollegeName", "SpecialtyName", "Grade", "ClassCode", (
    select count(*) from "Student" where "Student"."ClassID"="Class"."ClassID"
) "TotalStudentCount" 
from "College", "Specialty", "Class"
where "Specialty"."CollegeID"="College"."CollegeID" and "Class"."SpecialtyID"="Specialty"."SpecialtyID"
with check option;

create or replace view "StudentInfo" as
select "Human"."HumanID", "College"."CollegeID", "Specialty"."SpecialtyID", "Class"."ClassID", 
"CollegeName", "SpecialtyName", "Grade", "ClassCode", 
"Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash",
"StudentNumber", "AdmissionDate", "GraduationDate", "StudentDegree", "YearOfSchool", "Status"
from "Human", "Student", "College", "Specialty", "Class"
where "Student"."HumanID"="Human"."HumanID" 
and "Student"."ClassID"="Class"."ClassID"
and "Specialty"."SpecialtyID"="Class"."SpecialtyID" 
and "Specialty"."CollegeID"="College"."CollegeID"
with check option;

create or replace view "TeacherInfo" as
select "Human"."HumanID", "College"."CollegeID", "CollegeName", 
"Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash",
"TeacherNumber", "GraduationSchool", "Position", "TeacherDegree"
from "Teacher", "Human", "College"
where "Teacher"."HumanID"="Human"."HumanID" and "Teacher"."CollegeID"="College"."CollegeID"
with check option;
