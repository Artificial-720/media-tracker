
create table users (
  id INT PRIMARY KEY,
  username VARCHAR(50),
  password_hash VARCHAR(50),
  created_at DATE
);

create table media_items (
  id INT PRIMARY KEY,
  title VARCHAR(50),
  type VARCHAR(7),
  source VARCHAR(10),
  external_id VARCHAR(50)
);

create table user_media (
  id INT PRIMARY KEY,
  user_id INT NOT NULL,
  media_id INT NOT NULL,
  status VARCHAR(11),
  note TEXT,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(media_id) REFERENCES media_items(id)
);
