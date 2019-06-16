create or replace view "ClassInfo" as
select "Class"."ClassID", "CollegeName", "SpecialtyName", "Grade", "ClassCode", (
    select count(*) from "Student" where "Student"."ClassID"="Class"."ClassID"
    ) "TotalStudentCount" 
from "College", "Specialty", "Class"
where "Specialty"."CollegeID"="College"."CollegeID" and "Class"."SpecialtyID"="Specialty"."SpecialtyID"
with check option;

create or replace view "StudentInfo" as
select "Human"."HumanID", "Class"."ClassID", "StudentNumber","CollegeName", "SpecialtyName", "Grade", "ClassCode",
"Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash", "YearOfSchool"
from "Human", "Student", "College", "Specialty", "Class"
where "Student"."HumanID"="Human"."HumanID" and "Student"."ClassID"="Class"."ClassID"
and "Specialty"."SpecialtyID"="Class"."SpecialtyID" and "Specialty"."CollegeID"="College"."CollegeID"
with check option;

create or replace view "TeacherInfo" as
select "Teacher"."HumanID", "College"."CollegeID", "TeacherNumber", "CollegeName",
"Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash"
from "Teacher", "Human", "College"
where "Teacher"."HumanID"="Human"."HumanID" and "Teacher"."CollegeID"="College"."CollegeID"
with check option;
