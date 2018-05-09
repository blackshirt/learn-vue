create table user(
	id integer primary key autoincrement,
	username varchar(200) not null,
	firstname varchar(200),
	lastname varchar(200)
);

create table role(
	id integer primary key autoincrement,
	name varchar(200) not null,
	description varchar(200)
);

insert into role(name, description) values("anonymouse", "anonymouse user");
insert into role(name, description) values("reguler", "reguler user");
insert into role(name, description) values("operator", "operator user");
insert into role(name, description) values("root", "administrative user");


create table user_role(
	id integer primary key autoincrement,
	user integer not null,
	role integer not null,
	foreign key(user) references user(id) on update restrict on delete restrict;
	foreign key(role) references role(id) on update restrict on delete restrict;
);
