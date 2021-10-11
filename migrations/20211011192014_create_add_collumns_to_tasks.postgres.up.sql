ALTER TABLE task
    ADD COLUMN IF NOT EXISTS done bool default false;

