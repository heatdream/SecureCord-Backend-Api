CREATE TABLE "users" (
  "username" varchar UNIQUE PRIMARY KEY,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE,
  "uuid" uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");
