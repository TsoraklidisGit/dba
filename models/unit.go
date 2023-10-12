package models

// Unit represents a military unit.
type Unit struct {
	Name          string
	CombatVsFoot  int
	CombatVsMtn   int
	GoodGoingMove int
	BadGoingMove  int
	SpecialAttr   string
	ElementOrder  int
}
