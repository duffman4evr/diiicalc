package offensive

import (
	"diiicalc/util"
	"strconv"
)

type MetaStats struct {
	DerivedStats DerivedStats
	Dps          float64
}

type StatChangeEffect struct {
	Name  string
	Value string
	Color string
}

func NewMetaStats(derivedStats *DerivedStats) *MetaStats {

	self := new(MetaStats)

	self.DerivedStats = *derivedStats

	// Get the main stat.
	var mainStat float64

	switch heroClass := self.DerivedStats.BaseStats.HeroClass; {
	case heroClass == util.UrlValueHeroClassBarbarian:
		mainStat = derivedStats.Strength
	case heroClass == util.UrlValueHeroClassMonk || heroClass == util.UrlValueHeroClassDemonHunter:
		mainStat = derivedStats.Dexterity
	case heroClass == util.UrlValueHeroClassWizard || heroClass == util.UrlValueHeroClassWitchDoctor:
		mainStat = derivedStats.Intelligence
	}

	// TODO Question: are damage bonuses on rings and crap multiplied by the physical damage bonus on the weapon?

	// Looks like the equation is this
	//       Weapon Term: [AverageWeaponDamage] * 
	// Attack Speed Term: [WeaponBaseAttackSpeed * (1 + WeaponAttackSpeedBonus + OtherAttackSpeedBonuses)] *
	//    Main Stat Term: [1 + (MainStat * .01)] *
	//         Crit Term: [1 + (CritChance * CritDamageBonus)] *
	//        Skill Term: [1 + SkillBonus]
	var (
		baseStats = &self.DerivedStats.BaseStats

		weaponTerm      = baseStats.MainWeaponAverageDamage + baseStats.AverageDamageBonus
		attackSpeedTerm = baseStats.MainWeaponAttackSpeedBase * (1 + baseStats.MainWeaponAttackSpeedBonus + derivedStats.AttackSpeedBonus)
		mainStatTerm    = 1 + (mainStat * .01)
		critTerm        = 1 + (derivedStats.CritChance * (derivedStats.CritDamage - 1))
		skillTerm       = 1 + derivedStats.SkillDamageBonus
	)

	//fmt.Printf("\n\n\n%v\n%v\n%v\n%v\n%v\n\n\n", weaponTerm, attackSpeedTerm, mainStatTerm, critTerm, skillTerm)

	self.Dps = weaponTerm * attackSpeedTerm * mainStatTerm * critTerm * skillTerm

	return self
}

func (self *MetaStats) ComputeDpsChangeForCritChanceChange(critChanceChange float64) (dps float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.CritChance += critChanceChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeDpsChangeForCritDamageChange(critDamageChange float64) (dps float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.CritDamage += critDamageChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeDpsChangeForAttackSpeedChange(attackSpeedChange float64) (dps float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.AttackSpeedBonus += attackSpeedChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeDpsChangeForIntelligenceChange(mainStatChange float64) (dps float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Intelligence += mainStatChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeDpsChangeForStrengthChange(mainStatChange float64) (dps float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Strength += mainStatChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeDpsChangeForDexterityChange(mainStatChange float64) (dps float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Dexterity += mainStatChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeCritDamageEquivalentForCritChanceChange(critChanceChange float64) (dps float64) {
	return (self.DerivedStats.CritChance + critChanceChange) * (self.DerivedStats.CritDamage - 1) / self.DerivedStats.CritChance
}

func (self *MetaStats) CalculateDpsChange(changeType string, changeValue string) (dpsChange float64) {

	changeValueFloat, _ := strconv.ParseFloat(changeValue, 64)

	// TODO make this a switch on defensive side.
	switch {
	case changeType == util.UrlValueCompareTypeIntelligence:
		dpsChange = self.ComputeDpsChangeForIntelligenceChange(changeValueFloat)
	case changeType == util.UrlValueCompareTypeStrength:
		dpsChange = self.ComputeDpsChangeForStrengthChange(changeValueFloat)
	case changeType == util.UrlValueCompareTypeDexterity:
		dpsChange = self.ComputeDpsChangeForDexterityChange(changeValueFloat)
	case changeType == util.UrlValueCompareTypeAttackSpeed:
		dpsChange = self.ComputeDpsChangeForAttackSpeedChange(changeValueFloat * 0.01)
	case changeType == util.UrlValueCompareTypeCritChance:
		dpsChange = self.ComputeDpsChangeForCritChanceChange(changeValueFloat * 0.01)
	case changeType == util.UrlValueCompareTypeCritDamage:
		dpsChange = self.ComputeDpsChangeForCritDamageChange(changeValueFloat * 0.01)
	}

	return dpsChange
}