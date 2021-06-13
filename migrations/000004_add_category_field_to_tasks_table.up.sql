ALTER TABLE task
    ADD COLUMN category_id bigserial references categories;

