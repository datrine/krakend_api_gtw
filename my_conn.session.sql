/*
CREATE TABLE IF NOT EXISTS users (
    email VARCHAR(100) UNIQUE NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);
*/

/**/
PRAGMA table_info('users');

/*SELECT * FROM users;*/