CREATE TABLE summary (
                         id SERIAL PRIMARY KEY,
                         user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                         date DATE NOT NULL,
                         title VARCHAR(100) NOT NULL,
                         content TEXT NOT NULL,
                         is_shared BOOLEAN DEFAULT false
);