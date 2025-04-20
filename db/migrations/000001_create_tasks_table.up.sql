CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    id serial PRIMARY KEY,
    status VARCHAR(255) NOT NULL,
    task VARCHAR(255) NOT NULL,
    result VARCHAR(255),
    error VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    ended_at TIMESTAMP DEFAULT NULL,
    user_id INTEGER NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);