CREATE TABLE "whisky"
(
    id         SERIAL PRIMARY KEY,
    strength   integer NOT NULL,
    size       integer NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (NOW()),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (NOW())
);
