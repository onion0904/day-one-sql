CREATE TABLE diary (
    user_name VARCHAR NOT NULL,
    diary_date DATE NOT NULL,
    diary_content VARCHAR NOT NULL,
    PRIMARY KEY (user_name, diary_date),
    FOREIGN KEY (user_name) REFERENCES users(user_name) ON DELETE CASCADE
);