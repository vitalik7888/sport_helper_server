# TODO create indexes
CREATE TABLE persons(
  id INTEGER PRIMARY KEY,
  first_name VARCHAR(50) DEFAULT '',
  last_name VARCHAR(50) NOT NULL,
  birth_date DATE NOT NULL,
  gender CHAR(1) NOT NULL,
  height TINYINT(255) DEFAULT 0
);
CREATE TABLE training_sessions(
  id INTEGER PRIMARY KEY,
  person_id INTEGER NOT NULL,
  start DATETIME NOT NULL,
  end DATETIME NOT NULL,
  evaluation TINYINT(10) DEFAULT 0,
  notes VARCHAR(255) DEFAULT '',
  FOREIGN KEY (person_id) REFERENCES persons(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
CREATE TABLE exercises(
  id INTEGER PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(255) DEFAULT ''
);
CREATE TABLE training_exercises(
  id INTEGER PRIMARY KEY,
  session_id INTEGER NOT NULL,
  exercise_id INTEGER NOT NULL,
  total VARCHAR(20) NOT NULL,
  notes VARCHAR(255) DEFAULT '',
  FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY (session_id) REFERENCES training_sessions(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
