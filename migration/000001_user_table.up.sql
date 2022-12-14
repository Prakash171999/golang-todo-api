CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    full_name VARCHAR(255) NULL,
    email   VARCHAR(255) NOT NULL,
    phone_number   VARCHAR(50) NULL,
    password        VARCHAR(300) NOT NULL,
    user_role VARCHAR(50) NULL,
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
