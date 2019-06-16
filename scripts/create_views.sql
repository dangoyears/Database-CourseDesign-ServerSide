create view "ClassInfo" as
select "Class"."ClassID", "CollegeName", "SpecialtyName", "Grade", "ClassCode", (
    select count(*) from "Student" where "Student"."ClassID"="Class"."ClassID"
    ) "TotalStudentCount" 
from "College", "Specialty", "Class"
where "Specialty"."CollegeID"="College"."CollegeID" and "Class"."SpecialtyID"="Specialty"."SpecialtyID"
with check option;
