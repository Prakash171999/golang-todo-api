CREATE TABLE IF NOT EXISTS todos (
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(45) NULL,
  description VARCHAR(100) NULL,
  image VARCHAR(300) NULL,
  userId int,
  statusId int,
  priorityId int,
  categoryId int,
  FOREIGN KEY (userId) REFERENCES users(id),
  FOREIGN KEY (statusId) REFERENCES status(id),
  FOREIGN KEY (priorityId) REFERENCES priorities(id),
  FOREIGN KEY (categoryId) REFERENCES categories(id),
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  deleted_at DATETIME NULL,
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

