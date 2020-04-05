create table todos
(
	id serial,
	content varchar not null,
	created_at timestamp not null,
	completed bool default false
);

create unique index todos_id_uindex
	on todos (id);

alter table todos
	add constraint todos_pk
		primary key (id);

create table users
(
    id varchar not null,
    name varchar not null,
    email varchar not null,
    password varchar not null,
    created_at timestamp not null
);

create unique index users_email_uindex
    on users (email);

create unique index users_id_uindex
    on users (id);

alter table users
    add constraint users_pk
        primary key (id);

