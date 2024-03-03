CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       VARCHAR(255)        NOT NULL,
    nickname   VARCHAR(255) UNIQUE NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL
);

CREATE TABLE refresh_tokens
(
    token      UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Generate UUID v7 for token
    user_id    INT       NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);