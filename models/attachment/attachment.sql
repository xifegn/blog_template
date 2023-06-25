CREATE TABLE attachment (
                            id SERIAL PRIMARY KEY,
                            summary_id INTEGER NOT NULL REFERENCES summary(id) ON DELETE CASCADE,
                            url VARCHAR(200) NOT NULL,
                            type VARCHAR(50) NOT NULL
);