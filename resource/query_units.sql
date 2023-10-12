-- query_units.sql
SELECT units.name, units.combat_vs_foot, units.combat_vs_mtn, units.good_going_move,
       units.bad_going_move, units.special_attr, army_elements.element_order
FROM army_elements
JOIN units ON army_elements.unit_id = units.id
WHERE army_elements.army_id = ?
ORDER BY army_elements.element_order;