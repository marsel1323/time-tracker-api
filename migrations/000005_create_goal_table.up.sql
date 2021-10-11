CREATE TABLE IF NOT EXISTS goals
(
    id             bigserial PRIMARY KEY,
    name           text UNIQUE                 NOT NULL,
    time           bigint                      NOT NULL default 0,
    category_id    bigint                      NOT NULL REFERENCES categories ON DELETE CASCADE,
    active         bool                        NOT NULL default true,
    less_is_better bool                        NOT NULL default false,
    created_at     timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at     timestamp(0) with time zone NOT NULL DEFAULT NOW()
);