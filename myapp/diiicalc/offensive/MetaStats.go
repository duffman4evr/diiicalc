package offensive

import (
	"diiicalc/util"
	_"fmt"
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

	baseStats := &self.DerivedStats.BaseStats

	// Do some special setup in case of dual wielding. Got to make those two weapons look like one.
	var (
		averagedWeaponDamage      = baseStats.MainWeaponAverageDamage
		averagedWeaponAttackSpeed = baseStats.MainWeaponAttackSpeedBase * (1 + baseStats.MainWeaponAttackSpeedBonus + derivedStats.AttackSpeedBonus)
	)

	// TODO sometimes in-game damage does not match the damage you see when browsing blizz's own profiles...
	if baseStats.WeaponSetup == util.UrlValueWeaponSetupDualWield {

		averagedWeaponDamage += baseStats.OffWeaponAverageDamage
		averagedWeaponDamage *= 0.5
		averagedWeaponAttackSpeed += baseStats.OffWeaponAttackSpeedBase * (1 + baseStats.OffWeaponAttackSpeedBonus + derivedStats.AttackSpeedBonus)
		averagedWeaponAttackSpeed *= 0.5

	}

	// Looks like the equation is this:
	//       Weapon Term: [((AverageMainHandWeaponDamage + AverageOffHandWeaponDamage) / 2) + DamageBonusFromOtherItems] * 
	// Attack Speed Term: [WeaponBaseAttackSpeed * (1 + WeaponAttackSpeedBonus + OtherAttackSpeedBonuses)] *
	//    Main Stat Term: [1 + (MainStat * .01)] *
	//         Crit Term: [1 + (CritChance * CritDamageBonus)] *
	//        Skill Term: [1 + SkillBonus]
	var (
		weaponTerm      = averagedWeaponDamage + baseStats.AverageDamageBonus
		attackSpeedTerm = averagedWeaponAttackSpeed
		mainStatTerm    = 1 + (mainStat * .01)
		critTerm        = 1 + (derivedStats.CritChance * derivedStats.CritDamageBonus)
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

	modifiedDerivedStats.CritDamageBonus += critDamageChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeDpsChangeForAttackSpeedChange(attackSpeedChange float64) (dps float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.AttackSpeedBonus += attackSpeedChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeDpsChangeForMainStatChange(mainStatChange float64) (dps float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Strength += mainStatChange
	modifiedDerivedStats.Intelligence += mainStatChange
	modifiedDerivedStats.Dexterity += mainStatChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.Dps - self.Dps

}

func (self *MetaStats) ComputeCritDamageEquivalentForCritChanceChange(critChanceChange float64) (dps float64) {

	var (
		cc  = self.DerivedStats.CritChance
		cca = cc + critChanceChange
		cdb = self.DerivedStats.CritDamageBonus
	)

	return (cca * cdb / cc) - cdb
}

func (self *MetaStats) CalculateDpsChange(changeType string, changeValue string) (dpsChange float64) {

	changeValueFloat, _ := strconv.ParseFloat(changeValue, 64)

	switch {
	case changeType == util.UrlValueCompareTypeMainStat:
		dpsChange = self.ComputeDpsChangeForMainStatChange(changeValueFloat)
	case changeType == util.UrlValueCompareTypeAttackSpeed:
		dpsChange = self.ComputeDpsChangeForAttackSpeedChange(changeValueFloat * 0.01)
	case changeType == util.UrlValueCompareTypeCritChance:
		dpsChange = self.ComputeDpsChangeForCritChanceChange(changeValueFloat * 0.01)
	case changeType == util.UrlValueCompareTypeCritDamage:
		dpsChange = self.ComputeDpsChangeForCritDamageChange(changeValueFloat * 0.01)
	}

	return dpsChange
}
