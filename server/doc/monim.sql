SHOW TABLES;
ALTER TABLE user DROP COLUMN username;
ALTER TABLE user DROP COLUMN email;
ALTER TABLE user
ADD username VARCHAR(30) NOT NULL UNIQUE;
ALTER TABLE user
ADD email VARCHAR(100) NOT NULL UNIQUE;
-- 
CREATE TABLE user (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(30) NOT NULL UNIQUE,
    email varchar(100) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    profile_image VARCHAR(2000),
    PRIMARY KEY (id)
) ENGINE = InnoDB;
-- 
CREATE TABLE sosmed (
    id INT NOT NULL AUTO_INCREMENT,
    owner_id INT NOT NULL,
    instagram VARCHAR(2000),
    facebook VARCHAR(2000),
    twitter VARCHAR(2000),
    linkedin VARCHAR(2000),
    github VARCHAR(2000),
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES user(id)
) ENGINE = InnoDB;
--
CREATE TABLE skill (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;
--
CREATE TABLE mock_interview_post (
    id INT NOT NULL AUTO_INCREMENT,
    owner_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    role VARCHAR(30) NOT NULL,
    logo VARCHAR(2000),
    description TEXT,
    agreement TEXT,
    published_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES user(id)
) ENGINE = InnoDB;
--
CREATE TABLE webinar_post (
    id INT NOT NULL AUTO_INCREMENT,
    owner_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    banner VARCHAR(2000),
    description TEXT,
    published_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES user(id)
) ENGINE = InnoDB;
--
CREATE TABLE webinar_label (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;
--
CREATE TABLE user_skill (
    id INT NOT NULL AUTO_INCREMENT,
    owner_id INT NOT NULL,
    skill_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES user(id),
    FOREIGN KEY (skill_id) REFERENCES skill(id)
) ENGINE = InnoDB;
--
CREATE TABLE webinar_type (
    id INT NOT NULL AUTO_INCREMENT,
    webinar_id INT NOT NULL,
    label_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (webinar_id) REFERENCES webinar_post(id),
    FOREIGN KEY (label_id) REFERENCES webinar_label(id)
) ENGINE = InnoDB;
--