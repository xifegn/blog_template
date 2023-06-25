CREATE TABLE liked (
                       id SERIAL PRIMARY KEY,
                       summary_id INTEGER NOT NULL REFERENCES summary(id) ON DELETE CASCADE,
                       user_id INTEGER NOT NULL REFERENCES users(id)
);