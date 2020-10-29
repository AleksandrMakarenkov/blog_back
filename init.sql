-- public."account" definition
CREATE TABLE account (
	email varchar NOT NULL,
	"role" varchar NOT NULL,
	"password" varchar NOT NULL,
	id serial NOT NULL,
	CONSTRAINT account_pk PRIMARY KEY (id),
	CONSTRAINT account_un UNIQUE (email)
);