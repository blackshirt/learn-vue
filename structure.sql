PRAGMA foreign_keys=ON;


/* ------------------------------------------------------------------------- */
/* ---------------  users table -------------------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `users` (
	`uid` INTEGER PRIMARY KEY,
	`username` VARCHAR(50) NOT NULL UNIQUE,
	`hashedPwd` VARCHAR(200) NOT NULL	
);

/* ------------------------------------------------------------------------- */
/* ---------------  insert users sample ----------------------------------- */
/* ------------------------------------------------------------------------- */
BEGIN TRANSACTION;
INSERT INTO `users`(uid,username,hashedPwd) VALUES(0,"guest","0");
INSERT INTO `users`(uid,username,hashedPwd) VALUES(1,"blackshirt","1");
INSERT INTO `users`(uid,username,hashedPwd) VALUES(2,"greenshirt","2");
INSERT INTO `users`(uid,username,hashedPwd) VALUES(3,"redshirt","3");
INSERT INTO `users`(uid,username,hashedPwd) VALUES(4,"pinkshirt","4");
INSERT INTO `users`(uid,username,hashedPwd) VALUES(5,"blueshirt","5");
INSERT INTO `users`(uid,username,hashedPwd) VALUES(6,"brownshirt","6");
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  roles table -------------------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `roles` (
	`rid` INTEGER NOT NULL PRIMARY KEY,
	`name` VARCHAR(50) NOT NULL UNIQUE,
	`description` VARCHAR(100)
);

/* ------------------------------------------------------------------------- */
/* ---------------  insert roles sample ----------------------------------- */
/* ------------------------------------------------------------------------- */
BEGIN TRANSACTION;
INSERT INTO `roles`(rid,name,description) VALUES(0,"anonymouse", "anonymouse roles");
INSERT INTO `roles`(rid,name,description) VALUES(1,"root", "super user roles");
INSERT INTO `roles`(rid,name,description) VALUES(2,"admin", "administrative roles");
INSERT INTO `roles`(rid,name,description) VALUES(3,"regular", "regular roles");
INSERT INTO `roles`(rid,name,description) VALUES(4,"special", "special roles");
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  resources table ------------------------------------------ */
/* ------------------------------------------------------------------------- */
CREATE TABLE `resources` (
	`resid` INTEGER NOT NULL PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL UNIQUE,
	`locked` INTEGER NOT NULL DEFAULT 0
);

/* ------------------------------------------------------------------------- */
/* ---------------  insert resources sample ----------------------------------- */
/* ------------------------------------------------------------------------- */
BEGIN TRANSACTION;
INSERT INTO `resources`(resid,name) VALUES(0,"Users");
INSERT INTO `resources`(resid,name) VALUES(1,"Roles");
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  operations table --------------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `permissions` (
	`opid` INTEGER NOT NULL PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL,
	`locked` INTEGER NOT NULL DEFAULT 0
);

BEGIN TRANSACTION;
INSERT INTO `permissions`(opid, name) VALUES(0,"create");
INSERT INTO `permissions`(opid, name) VALUES(1,"read");
INSERT INTO `permissions`(opid, name) VALUES(2,"update");
INSERT INTO `permissions`(opid, name) VALUES(3,"delete");
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  permissions table -------------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `perm_resource` (
    `pid` INTEGER NOT NULL UNIQUE, -- to prevent foreign key mismatch when referenced
	`name` VARCHAR(100) NOT NULL,
	`description` VARCHAR(100),
	`perm_id` INTEGER NOT NULL,
    `resource_id` INTEGER NOT NULL,
    
	PRIMARY KEY(`pid`, `perm_id`, `resource_id`),
	FOREIGN KEY(`resource_id`) REFERENCES `resources`(`resid`),
	FOREIGN KEY(`perm_id`) REFERENCES `permissions`(`opid`)
);

BEGIN TRANSACTION;
INSERT INTO `perm_resource`(pid,name,perm_id,resource_id) VALUES(0,"createUsers",0,0);
INSERT INTO `perm_resource`(pid,name,perm_id,resource_id) VALUES(1,"readUsers",1,0);
INSERT INTO `perm_resource`(pid,name,perm_id,resource_id) VALUES(2,"updateUsers",2,0);
INSERT INTO `perm_resource`(pid,name,perm_id,resource_id) VALUES(3,"deleteUsers",3,0);
INSERT INTO `perm_resource`(pid,name,perm_id,resource_id) VALUES(4,"createRoles",0,1);
INSERT INTO `perm_resource`(pid,name,perm_id,resource_id) VALUES(5,"readRoles",1,1);
INSERT INTO `perm_resource`(pid,name,perm_id,resource_id) VALUES(6,"updateRoles",2,1);
INSERT INTO `perm_resource`(pid,name,perm_id,resource_id) VALUES(7,"deleteRoles",3,1);
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  role perm_resource table --------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `role_perm`(
	`role_id` INTEGER NOT NULL,
	`perm_id` INTEGER NOT NULL,
	
	PRIMARY KEY(`role_id`, `perm_id`),
	FOREIGN KEY(`role_id`) REFERENCES `roles`(rid),
	FOREIGN KEY(`perm_id`) REFERENCES `perm_resource`(pid)
);

BEGIN TRANSACTION;
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,0); -- root -> createUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,1); -- root -> readUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,2); -- root -> updateUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,3); -- root -> deleteUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,4); -- root -> createRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,5); -- root -> readRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,6); -- root -> updateRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,7); -- root -> deleteRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,1); -- admin -> readUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,2); -- admin -> updateUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,3); -- admin -> createRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,4); -- admin -> readRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,5); -- admin -> updateRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(3,1); -- regular -> readUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(3,4); -- regular -> readRoles
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  user role table ---------------------------------------- */
/* ----------Assign single user a role --------------------------------------*/
/* ------------------------------------------------------------------------- */
CREATE TABLE `user_role` (
	`user_id` INTEGER UNSIGNED NOT NULL,
	`role_id` INTEGER UNSIGNED NOT NULL,
	
	PRIMARY KEY(`user_id`, `role_id`),
	FOREIGN KEY(`user_id`) REFERENCES `users`(`uid`),
	FOREIGN KEY(`role_id`) REFERENCES `roles`(`rid`)
);

BEGIN TRANSACTION;
INSERT INTO `user_role`(user_id,role_id) VALUES(0,0); -- guest -> anonymouse
INSERT INTO `user_role`(user_id,role_id) VALUES(1,1); -- blackshirt -> root
INSERT INTO `user_role`(user_id,role_id) VALUES(2,2); -- greenshirt -> admin
INSERT INTO `user_role`(user_id,role_id) VALUES(3,3); -- redshirt -> regular
INSERT INTO `user_role`(user_id,role_id) VALUES(4,4); -- pinkshirt -> special
INSERT INTO `user_role`(user_id,role_id) VALUES(5,3); -- blueshirt -> regular
INSERT INTO `user_role`(user_id,role_id) VALUES(6,3); -- brownshirt -> regular

COMMIT TRANSACTION;