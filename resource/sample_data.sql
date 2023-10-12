-- Sample data for the 'units' table
INSERT INTO units (name, combat_vs_foot, combat_vs_mtn, good_going_move, bad_going_move, special_attr) VALUES
('Infantry', 3, 1, 4, 2, 'Charge'),
('Cavalry', 4, 2, 6, 3, 'Charge'),
('Artillery', 2, 0, 2, 1, 'Long-range');

-- Sample data for the 'armies' table
INSERT INTO armies (name) VALUES
('Army 1'),
('Army 2');

-- Sample data for the 'army_elements' table
-- Add 12 units to Army 1
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 1, 1); -- 1st unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 1, 2); -- 2nd unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 1, 3); -- 3rd unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 1, 4); -- 4th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 2, 5); -- 5th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 2, 6); -- 6th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 2, 7); -- 7th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 2, 8); -- 8th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 3, 9); -- 9th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 3, 10); -- 10th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 3, 11); -- 11th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (1, 3, 12); -- 12th unit

-- Add 12 units to Army 2
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 1, 1); -- 1st unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 1, 2); -- 2nd unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 1, 3); -- 3rd unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 1, 4); -- 4th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 2, 5); -- 5th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 2, 6); -- 6th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 2, 7); -- 7th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 2, 8); -- 8th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 3, 9); -- 9th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 3, 10); -- 10th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 3, 11); -- 11th unit
INSERT INTO army_elements (army_id, unit_id, element_order) VALUES (2, 3, 12); -- 12th unit
