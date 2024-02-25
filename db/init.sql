SELECT 'CREATE DATABASE hrga_api' 
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'hrga_api')\gexec

ALTER DATABASE hrga_api OWNER TO postgres;

\connect hrga_api

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off

CREATE TABLE IF NOT EXISTS public.employees
(
    id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT employee_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.employees
    OWNER to postgres;

CREATE TABLE IF NOT EXISTS public.users
(
    id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

INSERT INTO public.employees (id, name) VALUES ('3c900c8d-8081-4836-8a93-5f4d49eb09ac', 'staff1') ON CONFLICT DO NOTHING;


INSERT INTO public.users (id, email, password) VALUES ('23dd38da-88d4-4843-a97d-bf6e39730d36', 'admin@admin.com', '$2a$10$QCo6T/g4k7W5Xard6iP4h.6qSIGgXic3RnLJSYOSwy2vTFnizHnR6') ON CONFLICT DO NOTHING;
