CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,
  login TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  name TEXT,
  age INTEGER
  );