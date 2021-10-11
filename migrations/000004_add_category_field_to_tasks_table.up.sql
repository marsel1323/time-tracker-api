ALTER TABLE task
    ADD COLUMN IF NOT EXISTS category_id bigserial references categories;

