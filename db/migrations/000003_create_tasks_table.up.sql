CREATE TABLE tasks (
  id UUID UNIQUE NOT NULL,
  text VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  created_by_user_id UUID NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_user FOREIGN KEY(created_by_user_id) REFERENCES users(id)
);