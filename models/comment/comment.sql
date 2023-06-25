CREATE TABLE comment (
                         id SERIAL PRIMARY KEY,
                         summary_id INTEGER NOT NULL REFERENCES summary(id) ON DELETE CASCADE,
                         user_id INTEGER NOT NULL REFERENCES users(id),
                         content TEXT NOT NULL,
                         time TIMESTAMP NOT NULL
);