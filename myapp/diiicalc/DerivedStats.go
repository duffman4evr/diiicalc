package diiicalc

type DerivedStats struct {
	BaseStats         BaseStats
	SkillChoices      []SkillChoice
	Dexterity         float64
	Vitality          float64
	Armor             float64
	Life              float64
	LifeOnHit         float64
	LifeRegen         float64
	ResistArcane      float64
	ResistFire        float64
	ResistLightning   float64
	ResistPoison      float64
	ResistCold        float64
	ResistPhysical    float64
	MitigationSources map[string]float64
}

func NewDerivedStats(baseStats *BaseStats, skillChoices []SkillChoice) *DerivedStats {

	self := new(DerivedStats)

	self.BaseStats = *baseStats
	self.SkillChoices = skillChoices

	self.Dexterity = baseStats.Dexterity
	self.Vitality = baseStats.Vitality
	self.Armor = baseStats.Armor
	self.Life = baseStats.Life
	self.LifeOnHit = baseStats.LifeOnHit
	self.LifeRegen = baseStats.LifeRegen
	self.ResistArcane = baseStats.ResistArcane
	self.ResistFire = baseStats.ResistFire
	self.ResistLightning = baseStats.ResistLightning
	self.ResistPoison = baseStats.ResistPoison
	self.ResistCold = baseStats.ResistCold
	self.ResistPhysical = baseStats.ResistPhysical

	for i := 0; i < len(skillChoices); i++ {
		skillChoices[i].ModifyDerivedStats(self)
	}

	return self
}
