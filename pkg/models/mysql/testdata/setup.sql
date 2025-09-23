CREATE TABLE snippets (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL
)

CREATE INDEX idx_snippets_created ON snippets(created);

CREATE TABLE users (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(225) NOT NULL,
  email VARCHAR(225) NOT NULL,
  hashed_password VARCHAR(225) NOT NULL,
  created DATETIME NOT NULL
);

ALTER TABLE users
   ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users (
  name, email, hashed_password, created
  ) VALUES ( 
  'John Doe',
  'johndoe@example.com',
  '$2a$12$wI9e0bY8h1HjkFh8H6fUuOa8G9b1QeF5z1Z6Z6Z6Z6Z6Z6Z6Z6',
 '2024-06-01 00:00:00'
);
