ALTER TABLE task
    ADD COLUMN IF NOT EXISTS category_id int references categories;

