CREATE TYPE priority AS ENUM ('low', 'medium', 'high');

CREATE TABLE tasks (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT (now()) NOT NULL,
    priority priority NOT NULL DEFAULT('low'),
    deadline TIMESTAMP WITH TIME ZONE
);