CREATE TABLE users (
  id UUID,
  name VARCHAR(355),
  email VARCHAR(320),
  github_id INTEGER UNIQUE
);