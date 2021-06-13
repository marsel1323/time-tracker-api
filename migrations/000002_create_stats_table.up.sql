CREATE TABLE IF NOT EXISTS stats
(
    id           bigserial PRIMARY KEY,
    milliseconds bigint                      NOT NULL,
    task_id      bigint                      NOT NULL REFERENCES task ON DELETE CASCADE,
    created_at   timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at   timestamp(0) with time zone NOT NULL DEFAULT NOW()
);