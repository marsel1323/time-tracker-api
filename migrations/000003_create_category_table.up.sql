CREATE TABLE IF NOT EXISTS categories
(
    id         bigserial PRIMARY KEY,
    name       text UNIQUE                 NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

INSERT INTO categories (name)
VALUES ('all'),
       ('development'),
       ('job'),
       ('sport'),
       ('gaming'),
       ('english'),
       ('reading'),
       ('drawing'),
       ('3d'),
       ('bad habits'),
       ('house cleaning'),
       ('guitar')
ON CONFLICT DO NOTHING;