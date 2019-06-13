-- 创建数据库模式
drop user dbcd cascade;
create user dbcd identified by dbcdpwd;
grant connect, resource, dba to dbcd;
