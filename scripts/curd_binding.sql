-- AcademicYear
insert into "AcademicYear" values (:1)
delete from "AcademicYear" where "AcademicYear"."AcademicYear"=:1

-- Administrator
select count(*) from "Administrator"
insert into "Administrator" ("AdminLoginName", "AdminPassHash")
values (:1, :2)

-- College
insert into "College" ("CollegeName") 
values (:1)
delete from "College" where "CollegeName"=:1
