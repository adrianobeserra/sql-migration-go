CREATE SCHEMA db_schema;
GRANT ALL ON SCHEMA db_schema TO postgres;

CREATE TABLE db_schema.users (
   id BIGSERIAL primary key,
   first_name VARCHAR not null,
   last_name VARCHAR,
   created_at TIMESTAMP default now()
);

CREATE TABLE db_schema.other_table (
   id BIGSERIAL primary key,
   description VARCHAR not null,
   created_at TIMESTAMP default now()
);