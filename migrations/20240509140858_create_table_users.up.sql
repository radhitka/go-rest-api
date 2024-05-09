CREATE TABLE users
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(225)  NULL,
    username VARCHAR(225)  NULL,
    password VARCHAR(225)  NULL,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB;