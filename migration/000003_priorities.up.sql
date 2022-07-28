CREATE TABLE IF NOT EXISTS priorities (
  id INT NOT NULL AUTO_INCREMENT,
  priority_type VARCHAR(45) NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NULL,
  deleted_at DATETIME NULL,
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
