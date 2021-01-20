CREATE TABLE tasks (
  id UUID,
  text VARCHAR,
  created_at TIMESTAMPTZ,
  created_by_user_id UUID,
  PRIMARY KEY(id),
  CONSTRAINT fk_user FOREIGN KEY(created_by_user_id) REFERENCES users(id)
);