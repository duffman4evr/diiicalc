package offensive

type DerivedStats struct {
	BaseStats        BaseStats
	Strength         float64
	Dexterity        float64
	Intelligence     float64
	CritChance       float64
	CritDamage       float64
	AttackSpeedBonus float64
	SkillDamageBonus float64
}

func NewDerivedStats(baseStats *BaseStats) *DerivedStats {

	self := new(DerivedStats)

	self.BaseStats = *baseStats

	self.Strength = baseStats.Strength
	self.Dexterity = baseStats.Dexterity
	self.Intelligence = baseStats.Intelligence
	self.CritChance = baseStats.CritChance
	self.CritDamage = baseStats.CritDamage
	self.AttackSpeedBonus = baseStats.AttackSpeedBonus

	return self
}
