-- Table Human
insert into "Human" ("Name") 
values ('仁占位');

-- Table College
insert into "College" ("CollegeName")
values ('计算机科学与网络工程学院');
update "College" set "College"."CollegeID"=0
where "College"."CollegeName"='计算机科学与网络工程学院';

-- Table Specialty
insert into "Specialty" ("CollegeID", "SpecialtyName")
values (0, '软件工程');
update "Specialty" set "Specialty"."SpecialtyID"=0
where "Specialty"."SpecialtyName"='软件工程';

-- Table Class
insert into "Class" ("SpecialtyID", "Grade", "ClassCode")
values (0, 17, 1);
