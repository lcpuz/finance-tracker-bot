CREATE TABLE IF NOT EXISTS tbl_users (
    user_id SERIAL PRIMARY KEY,
    user_chat_id BIGINT NOT NULL UNIQUE,
    user_name VARCHAR(255),
    user_nickname VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tbl_categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
            REFERENCES tbl_users(user_id)
                ON DELETE CASCADE
);