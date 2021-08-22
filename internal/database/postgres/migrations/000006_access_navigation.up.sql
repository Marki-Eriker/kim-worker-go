CREATE TABLE access_navigation
(
    access_id BIGINT REFERENCES access (id) ON DELETE CASCADE,
    navigation_id BIGINT REFERENCES navigation (id) ON DELETE CASCADE
);
