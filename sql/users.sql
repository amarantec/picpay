CREATE TABLE IF NOT EXISTS users (
	id serial primary key,
	first_name text not null,
	last_name text not null,
	document text not null unique,
	email text not null unique,
	password text not null,
	balance	double precision,
	user_type integer
);
