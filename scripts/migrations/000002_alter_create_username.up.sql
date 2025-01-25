ALTER TABLE fastcampus.users
    ADD username VARCHAR(255) NOT NULL;

ALTER TABLE fastcampus.users
    ADD CONSTRAINT users_username_unique UNIQUE (username);