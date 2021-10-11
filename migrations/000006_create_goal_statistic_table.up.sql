CREATE TABLE IF NOT EXISTS goal_statistics
(
    id           bigserial PRIMARY KEY,
    milliseconds bigint                      NOT NULL,
    goal_id      bigint                      NOT NULL REFERENCES goals ON DELETE CASCADE,
    created_at   timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at   timestamp(0) with time zone NOT NULL DEFAULT NOW()
);