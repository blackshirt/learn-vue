PRAGMA foreign_keys=ON;


/* ------------------------------------------------------------------------- */
/* ---------------  users table -------------------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `users` (
	`uid` INTEGER PRIMARY KEY,
	`username` VARCHAR(50) NOT NULL UNIQUE,
	`firstname` VARCHAR(50),
	`lastname` VARCHAR(50),
	`hashedPwd` VARCHAR(200)	
);

/* ------------------------------------------------------------------------- */
/* ---------------  insert users sample ----------------------------------- */
/* ------------------------------------------------------------------------- */
BEGIN TRANSACTION;
INSERT INTO `users`(uid,username,firstname,lastname) VALUES(0,"guest","guest","guest");
INSERT INTO `users`(uid,username,firstname,lastname) VALUES(1,"blackshirt","black","shirt");
INSERT INTO `users`(uid,username,firstname,lastname) VALUES(2,"greenshirt","green","shirt");
INSERT INTO `users`(uid,username,firstname,lastname) VALUES(3,"redshirt","red","shirt");
INSERT INTO `users`(uid,username,firstname,lastname) VALUES(4,"pinkshirt","pink","shirt");
INSERT INTO `users`(uid,username,firstname,lastname) VALUES(5,"blueshirt","blue","shirt");
INSERT INTO `users`(uid,username,firstname,lastname) VALUES(6,"brownshirt","brown","shirt");
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
/* ---------------  objects table ------------------------------------------ */
/* ------------------------------------------------------------------------- */
CREATE TABLE `objects` (
	`obid` INTEGER NOT NULL PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL UNIQUE,
	`locked` INTEGER NOT NULL DEFAULT 0
);

/* ------------------------------------------------------------------------- */
/* ---------------  insert object sample ----------------------------------- */
/* ------------------------------------------------------------------------- */
BEGIN TRANSACTION;
INSERT INTO `objects`(obid,name) VALUES(0,"Users");
INSERT INTO `objects`(obid,name) VALUES(1,"Roles");
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  operations table --------------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `operations` (
	`opid` INTEGER NOT NULL PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL,
	`locked` INTEGER NOT NULL DEFAULT 0
);

BEGIN TRANSACTION;
INSERT INTO `operations`(opid, name) VALUES(0,"preventing");
INSERT INTO `operations`(opid, name) VALUES(1,"creating");
INSERT INTO `operations`(opid, name) VALUES(2,"reading");
INSERT INTO `operations`(opid, name) VALUES(3,"updating");
INSERT INTO `operations`(opid, name) VALUES(4,"deleting");
INSERT INTO `operations`(opid, name, locked) VALUES(5,"locking", 1);
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  permissions table -------------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `permissions` (
    `pid` INTEGER NOT NULL UNIQUE, -- to prevent foreign key mismatch when referenced
	`name` VARCHAR(100) NOT NULL,
	`description` VARCHAR(100),
	`operation_id` INTEGER NOT NULL,
    `object_id` INTEGER NOT NULL,
    
	PRIMARY KEY(`pid`, `operation_id`, `object_id`),
	FOREIGN KEY(`object_id`) REFERENCES `objects`(`obid`),
	FOREIGN KEY(`operation_id`) REFERENCES `operations`(`opid`)
);

BEGIN TRANSACTION;
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(0,"noAccessUsers",0,0);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(1,"createUsers",1,0);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(2,"readUsers",2,0);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(3,"updateUsers",3,0);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(4,"deleteUsers",4,0);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(5,"noAccessRoles",0,1);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(6,"createRoles",1,1);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(7,"readRoles",2,1);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(8,"updateRoles",3,1);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(9,"deleteRoles",4,1);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(10,"lockUser",5,0);
INSERT INTO `permissions`(pid,name,operation_id,object_id) VALUES(11,"lockRole",5,1);
COMMIT TRANSACTION;



/* ------------------------------------------------------------------------- */
/* ---------------  role permissions table --------------------------------- */
/* ------------------------------------------------------------------------- */
CREATE TABLE `role_perm`(
	`role_id` INTEGER NOT NULL,
	`perm_id` INTEGER NOT NULL,
	
	PRIMARY KEY(`role_id`, `perm_id`),
	FOREIGN KEY(`role_id`) REFERENCES `roles`(rid),
	FOREIGN KEY(`perm_id`) REFERENCES `permissions`(pid)
);

BEGIN TRANSACTION;
INSERT INTO `role_perm`(role_id, perm_id) VALUES(0,0); -- anonymouse -> noAccessRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(0,5); -- anonymouse -> noAccessUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,1); -- root -> createUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,2); -- root -> readUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,3); -- root -> updateUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,4); -- root -> deleteUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,5); -- root -> noAccessUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,6); -- root -> createRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,7); -- root -> readRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,8); -- root -> updateRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,9); -- root -> deleteRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,10); -- root -> lockUser
INSERT INTO `role_perm`(role_id, perm_id) VALUES(1,11); -- root -> lockRole
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,2); -- admin -> readUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,3); -- admin -> updateUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,7); -- admin -> createRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(2,8); -- admin -> updateRoles
INSERT INTO `role_perm`(role_id, perm_id) VALUES(3,2); -- regular -> readUsers
INSERT INTO `role_perm`(role_id, perm_id) VALUES(3,7); -- regular -> readRoles
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