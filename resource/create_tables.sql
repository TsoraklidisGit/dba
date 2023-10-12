-- Create the 'units' table
CREATE TABLE IF NOT EXISTS units (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT,
   combat_vs_foot INT,
   combat_vs_mtn INT,
   good_going_move INT,
   bad_going_move INT,
   special_attr TEXT
);

-- Create the 'armies' table
CREATE TABLE IF NOT EXISTS armies (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT
);

-- Create the 'army_elements' table
CREATE TABLE IF NOT EXISTS army_elements (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   army_id INTEGER,
   unit_id INTEGER,
   element_order INT,
   FOREIGN KEY (army_id) REFERENCES armies (id),
   FOREIGN KEY (unit_id) REFERENCES units (id)
);
