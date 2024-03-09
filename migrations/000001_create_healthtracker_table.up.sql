CREATE TABLE IF NOT EXISTS healthtracker (
id bigserial PRIMARY KEY,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
walking integer NOT NULL,
hydrate integer NOT NULL,
sleep integer NOT NULL,
version integer NOT NULL DEFAULT 1
);