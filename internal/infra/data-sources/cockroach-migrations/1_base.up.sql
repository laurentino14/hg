CREATE SCHEMA IF NOT EXISTS lv;

CREATE SEQUENCE IF NOT EXISTS lv.seq_roles START WITH 0 INCREMENT BY 1 MINVALUE 0 NO MAXVALUE CACHE 1;

CREATE SEQUENCE IF NOT EXISTS lv.seq_oss START WITH 1 INCREMENT BY 1 MINVALUE 1 NO MAXVALUE CACHE 1;

CREATE SEQUENCE IF NOT EXISTS lv.seq_projects START WITH 1 INCREMENT BY 1 MINVALUE 1 NO MAXVALUE CACHE 1;

CREATE SEQUENCE IF NOT EXISTS lv.seq_services START WITH 1 INCREMENT BY 1 MINVALUE 1 NO MAXVALUE CACHE 1;

CREATE SEQUENCE IF NOT EXISTS lv.seq_tasks START WITH 1 INCREMENT BY 1 MINVALUE 1 NO MAXVALUE CACHE 1;

CREATE TABLE IF NOT EXISTS lv.roles (
    "role_cod" INT NOT NULL PRIMARY KEY DEFAULT nextval('lv.seq_roles'),
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    "updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
    "deleted_at" TIMESTAMPTZ NULL DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "roles_name_uindex" ON lv.roles ("name");
CREATE UNIQUE INDEX IF NOT EXISTS "roles_cod_uindex" ON lv.roles ("role_cod");

CREATE TABLE IF NOT EXISTS lv.users (
    "user_id" VARCHAR(255) NOT NULL PRIMARY KEY UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "birthdate" TIMESTAMPTZ NOT NULL,
    "type" INT8 REFERENCES lv.roles("role_cod"),
    "doc" VARCHAR(11) NOT NULL,
    "phone" VARCHAR(14) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
    "deleted_at" TIMESTAMPTZ NULL DEFAULT NULL
);


CREATE UNIQUE INDEX IF NOT EXISTS "users_email_uindex" ON lv.users ("email");
CREATE UNIQUE INDEX IF NOT EXISTS "users_doc_uindex" ON lv.users ("doc");
CREATE UNIQUE INDEX IF NOT EXISTS "users_phone_uindex" ON lv.users ("phone");
CREATE UNIQUE INDEX IF NOT EXISTS "users_id_uindex" ON lv.users ("user_id");

CREATE TABLE IF NOT EXISTS lv.services (
    "service_cod" INT NOT NULL PRIMARY KEY DEFAULT nextval('lv.seq_services'),
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "description" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
    "deleted_at" TIMESTAMPTZ NULL DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "services_name_uindex" ON lv.services ("name");
CREATE UNIQUE INDEX IF NOT EXISTS "services_cod_uindex" ON lv.services ("service_cod");

CREATE TABLE IF NOT EXISTS lv.projects (
    "project_cod" INT NOT NULL PRIMARY KEY UNIQUE DEFAULT nextval('lv.seq_projects'),
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "status" INT8 NOT NULL DEFAULT 0,
    "type" INT REFERENCES lv.services("service_cod"),
    "owner_id" VARCHAR(255) REFERENCES lv.users("user_id"),
    "start_date" TIMESTAMPTZ NULL DEFAULT NULL,
    "end_date" TIMESTAMPTZ NULL DEFAULT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
    "deleted_at" TIMESTAMPTZ NULL DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "projects_cod_uindex" ON lv.projects ("project_cod");
CREATE UNIQUE INDEX IF NOT EXISTS "projects_owner_uindex" ON lv.projects ("owner_id");

CREATE TABLE IF NOT EXISTS lv.os (
    "os_cod" INT NOT NULL PRIMARY KEY,
    "project_cod" INT REFERENCES lv.projects("project_cod"),
    "price" INT NOT NULL,
    "status" INT NOT NULL DEFAULT 0,
    "worker_id" VARCHAR(255) REFERENCES lv.users("user_id"),
    "client_id" VARCHAR(255) REFERENCES lv.users("user_id"),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
    "deleted_at" TIMESTAMPTZ NULL DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "os_cod_uindex" ON lv.os("os_cod");
CREATE UNIQUE INDEX IF NOT EXISTS "os_worker_uindex" ON lv.os ("worker_id");
CREATE UNIQUE INDEX IF NOT EXISTS "os_client_uindex" ON lv.os ("client_id");
CREATE UNIQUE INDEX IF NOT EXISTS "os_project_uindex" ON lv.os ("project_cod");

CREATE TABLE IF NOT EXISTS lv.tasks (
    "task_cod" INT NOT NULL PRIMARY KEY DEFAULT nextval('lv.seq_tasks'),
    "name" VARCHAR(255) NOT NULL,
    "status" INT NOT NULL DEFAULT 0,
    "owner_id" VARCHAR(255) REFERENCES lv.users("user_id"),
    "project_cod" INT REFERENCES lv.projects("project_cod"),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
    "deleted_at" TIMESTAMPTZ NULL DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "tasks_cod_uindex" ON lv.tasks ("task_cod");
CREATE UNIQUE INDEX IF NOT EXISTS "tasks_owner_uindex" ON lv.tasks  ("owner_id");
CREATE UNIQUE INDEX IF NOT EXISTS "tasks_project_uindex" ON lv.tasks  ("project_cod");

INSERT INTO lv.roles ("name") VALUES ('client');
INSERT INTO lv.roles ("name") VALUES ('admin');
INSERT INTO lv.roles ("name") VALUES ('worker');