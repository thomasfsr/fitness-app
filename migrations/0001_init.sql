-- Basic tables for users, workouts, messages (for reference)
CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  first_name char(20),
  last_name char(50),
  whatsapp bigint,
  active boolean DEFAULT true,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);

CREATE TYPE message_role AS ENUM ('user','assistant');

CREATE TABLE IF NOT EXISTS workout_sets (
  id bigserial PRIMARY KEY,
  user_id integer REFERENCES users(id),
  exercise char(100),
  weight integer,
  reps smallint,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS messages (
  id bigserial PRIMARY KEY,
  user_id integer REFERENCES users(id),
  role message_role,
  message char(100),
  created_at timestamptz DEFAULT now()
);
