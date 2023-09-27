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
    category_name VARCHAR(255) NOT NULL UNIQUE,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
            REFERENCES tbl_users(user_id)
                ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tbl_user_states (
    user_state_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    income_state INT,
    expense_state INT,   
    income_category_state INT,
    expense_category_state INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
            REFERENCES tbl_users(user_id)
                ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tbl_transactions (
    transaction_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    category_id INT NOT NULL,
    transaction_type INT NOT NULL, -- 0: income, 1: expense, 2: uncategorized 
    transaction_amount INT NOT NULL,
    transaction_description VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
            REFERENCES tbl_users(user_id)
                ON DELETE CASCADE,

    CONSTRAINT fk_category_id
        FOREIGN KEY(category_id)
            REFERENCES tbl_categories(category_id)
                ON DELETE CASCADE
);