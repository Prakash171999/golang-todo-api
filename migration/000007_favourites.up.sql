CREATE TABLE IF NOT EXISTS favourites (
  id INT NOT NULL AUTO_INCREMENT,
  userId int,
  todoId int,
  FOREIGN KEY (userId) REFERENCES users(id),
  FOREIGN KEY (todoId) REFERENCES todos(id),
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  deleted_at DATETIME NULL,
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
