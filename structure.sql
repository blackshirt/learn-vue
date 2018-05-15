create table `users` (
	`id` integer primary key,
	`username` varchar(50) not null,
	`firstname` varchar(50),
	`lastname` varchar(50),
	`hashedPwd` varchar(200)	
);

create table `roles` (
	`id` integer primary key,
	`name` varchar(50) not null,
	`description` varchar(100)	
);

create table `permissions` (
	`id` integer primary key,
	`name` varchar(50) not null,
	`description` varchar(50)
);

create table `role_perm` (
	`role_id` integer unsigned not null,
	`perm_id` integer unsigned not null,
	
	foreign key(`role_id`) references roles(`id`),
	foreign key(`perm_id`) references permissions(`id`)
);

create table `user_role` (
	`user_id` integer unsigned not null,
	`role_id` integer unsigned not null,
	
	foreign key(`user_id`) references users(`id`),
	foreign key(`role_id`) references roles(`id`)
);