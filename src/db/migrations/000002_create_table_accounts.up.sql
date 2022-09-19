CREATE TABLE db_schema.accounts (
   id BIGSERIAL PRIMARY KEY,
   owner VARCHAR NOT NULL,
   balance BIGINT NOT NULL,
   currency VARCHAR NOT NULL,
   created_at TIMESTAMP NOT NULL default now()
);