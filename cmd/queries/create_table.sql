CREATE TABLE snippets (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets (
  title, content, created, expires
  ) VALUES ( 
  'An old silen pond',
  'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsuo Basho',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (
  title, content, created, expires
  ) VALUES ( 
  'Over the intry forest',
  'Over the wintry\nforest, winds holw in rage\nwith no leaves to blow.\n\n- Natsume Soseki',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);


INSERT INTO snippets (
  title, content, created, expires
  ) VALUES ( 
  'In jails',
  'Inside the jail, no rum nor flower\nGlorious nightsight is not ignoreable\nwith no leaves to blow.\n\n- Natsume Soseki',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';

-- CREATE USER TABLE

CREATE TABLE users (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(225) NOT NULL,
  email VARCHAR(225) NOT NULL,
  hashed_password VARCHAR(225) NOT NULL,
  created DATETIME NOT NULL
);

ALTER TABLE 
ALTER TABLE users
   ADD CONSTRAINT users_uc_email UNIQUE (email);
