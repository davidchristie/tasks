CREATE TABLE users (
  id UUID,
  name VARCHAR,
  email VARCHAR,
  github_id INTEGER UNIQUE,
  PRIMARY KEY(id)
);