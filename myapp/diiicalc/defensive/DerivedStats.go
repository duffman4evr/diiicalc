package defensive

import "diiicalc/util"

type DerivedStats struct {
	BaseStats         BaseStats
	Dexterity         float64
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

func NewDerivedStats(baseStats *BaseStats) *DerivedStats {

	self := new(DerivedStats)

	self.BaseStats = *baseStats
	self.Dexterity = baseStats.Dexterity
	self.Armor = baseStats.Armor
	self.Life = util.DeriveLifeFromVitality(baseStats.Vitality, baseStats.Level) * (1 + baseStats.LifePercent)
	self.LifeOnHit = baseStats.LifeOnHit
	self.LifeRegen = baseStats.LifeRegen
	self.ResistArcane = baseStats.ResistArcane
	self.ResistFire = baseStats.ResistFire
	self.ResistLightning = baseStats.ResistLightning
	self.ResistPoison = baseStats.ResistPoison
	self.ResistCold = baseStats.ResistCold
	self.ResistPhysical = baseStats.ResistPhysical

	self.MitigationSources = make(map[string]float64)

	return self
}
