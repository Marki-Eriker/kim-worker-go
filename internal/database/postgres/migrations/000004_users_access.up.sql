CREATE TABLE users_access
(
    user_id BIGINT REFERENCES users (id) ON DELETE CASCADE,
    access_id BIGINT REFERENCES access (id) ON DELETE CASCADE
);
