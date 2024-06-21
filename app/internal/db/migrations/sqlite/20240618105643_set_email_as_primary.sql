-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS usersB (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) UNIQUE NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    is_email_verified BOOLEAN DEFAULT true NOT NULL,
    status VARCHAR(255) DEFAULT "UNVERIFIED" NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);
INSERT INTO usersB (
        email,
        password_hash,
        last_name,
        first_name,
        created_at,
        updated_at
    ) 
SELECT
        email,
        password_hash,
        last_name,
        first_name,
        created_at,
        updated_at
    
FROM users
WHERE true;
DROP TABLE users;
ALTER TABLE usersB
    RENAME TO users;
DROP TABLE usersB;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd