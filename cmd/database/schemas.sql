CREATE TABLE IF NOT EXISTS "users" (user_name TEXT PRIMARY KEY);

CREATE TABLE IF NOT EXISTS "solutions" (
    key TEXT PRIMARY KEY,
    file_name TEXT NOT NULL UNIQUE
);