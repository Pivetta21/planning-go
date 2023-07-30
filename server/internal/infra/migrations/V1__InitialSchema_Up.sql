CREATE DATABASE planning_go;

ALTER DATABASE planning_go SET TIME ZONE 'UTC';

CREATE TABLE IF NOT EXISTS public.users(
	id BIGINT GENERATED ALWAYS AS IDENTITY,
	username VARCHAR(25) NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	CONSTRAINT pk_users PRIMARY KEY (id),
	CONSTRAINT uq_users_username UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS public.user_sessions(
	id BIGINT GENERATED ALWAYS AS IDENTITY,
	user_id BIGINT NOT NULL,
	identifier VARCHAR(25) NOT NULL,
	opaque_token UUID NOT NULL CONSTRAINT uq_user_sessions_opaque_token UNIQUE,
	origin INTEGER NOT NULL,
	expires_at TIMESTAMPTZ NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	CONSTRAINT pk_user_sessions PRIMARY KEY (id),
	CONSTRAINT fk_user_sessions_users FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE,
	CONSTRAINT uq_user_sessions_user_id_identifier UNIQUE (user_id, identifier)
);

