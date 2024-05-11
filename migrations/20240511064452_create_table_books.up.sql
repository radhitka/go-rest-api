CREATE TABLE books
(
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(225)  NULL,
    total_pages INT NULL,
    cover VARCHAR(225)  NULL,
    author VARCHAR(225)  NULL,
    publisher VARCHAR(225)  NULL,
    is_published TINYINT NULL,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB;