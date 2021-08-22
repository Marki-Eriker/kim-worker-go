CREATE TABLE users
(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    base_role VARCHAR(50) NOT NULL,
    service_types INTEGER[] DEFAULT ARRAY[]::INTEGER[]
);
