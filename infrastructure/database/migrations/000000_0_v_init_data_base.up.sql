CREATE TABLE IF NOT EXISTS "user" (
    id uuid NOT NULL,
    names character varying(250) NOT NULL,
    last_names character varying(250) NOT NULL,
    email character varying(150) NOT NULL UNIQUE,
    password character varying(300) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id)
);
ALTER TABLE "user" OWNER to postgres;