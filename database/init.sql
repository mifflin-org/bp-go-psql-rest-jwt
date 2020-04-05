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
