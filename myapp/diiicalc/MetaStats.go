package diiicalc

import (
	"fmt"
)

// Some useful types
type MetaStats struct {
	DerivedStats            DerivedStats
	MinResistance           float64
	AverageBlockAmount      float64
	EffectiveLifeMultiplier float64
	EffectiveLife           float64
	EffectiveLifeNoShield   float64
	EffectiveLifeOnHit      float64
	EffectiveLifeRegen      float64
	TotalMitigation         float64
	MitigationSources       map[string]float64
}

type StatChangeEffect struct {
	Name  string
	Value string
	Color string
}

const (
	MitigationSourceArmor       = "Armor"
	MitigationSourceResistances = "Resistance"
	MitigationSourceMeleeClass  = "Monk/Barb Bonus"
	MitigationSourceDodge       = "Dodge"
)

func NewMetaStats(derivedStats *DerivedStats) *MetaStats {

	self := new(MetaStats)

	self.MitigationSources = make(map[string]float64)

	// Copy in some data.
	self.DerivedStats = *derivedStats

	// Find the minimum resistance that the user has.
	self.MinResistance = findMin(self.DerivedStats.ResistArcane, self.DerivedStats.ResistFire, self.DerivedStats.ResistLightning, self.DerivedStats.ResistPoison, self.DerivedStats.ResistCold, self.DerivedStats.ResistPhysical)

	// Calculate our average block amount.
	self.AverageBlockAmount = self.DerivedStats.BaseStats.BlockChance * ((self.DerivedStats.BaseStats.BlockAmountMin + self.DerivedStats.BaseStats.BlockAmountMax) / 2)

	// Calculate our different reduction sources.
	reductionFromArmor := self.DerivedStats.Armor / ((50.0 * self.DerivedStats.BaseStats.Level) + self.DerivedStats.Armor)
	reductionFromResistances := self.MinResistance / ((5.0 * self.DerivedStats.BaseStats.Level) + self.MinResistance)

	self.MitigationSources[MitigationSourceArmor] = reductionFromArmor
	self.MitigationSources[MitigationSourceResistances] = reductionFromResistances

	// Special case: Melee classes should have a reduction of 30%.
	if self.DerivedStats.BaseStats.HeroClass == urlValueHeroClassBarbarian || self.DerivedStats.BaseStats.HeroClass == urlValueHeroClassMonk {
		self.MitigationSources[MitigationSourceMeleeClass] = 0.30
	}

	// Add any mitigation sources that might have been added by DerivedStats (Skills like Ignore Pain)
	for key := range self.DerivedStats.MitigationSources {
		self.MitigationSources[key] = self.DerivedStats.MitigationSources[key]
	}

	// Icorporate dodge from dexterity into the mitigation.
	// Be careful to *multiplicatively* stack this with existing dodge from abilities, if any.
	addDodge(getDodgeChanceFromDexterity(self.DerivedStats.Dexterity), &self.MitigationSources)

	var totalNonMitigation float64 = 1.0

	// Multiplicatively stack all mitigation sources to get our total mitigation.
	for key := range self.MitigationSources {
		totalNonMitigation *= 1 - self.MitigationSources[key]
	}

	self.TotalMitigation = 1 - totalNonMitigation
	self.EffectiveLifeNoShield = self.DerivedStats.Life * (1 / (1 - self.TotalMitigation))
	self.EffectiveLifeMultiplier = (1 / (1 - self.TotalMitigation)) * (1 + (self.AverageBlockAmount / self.DerivedStats.Life))
	self.EffectiveLife = self.DerivedStats.Life * self.EffectiveLifeMultiplier
	self.EffectiveLifeOnHit = self.DerivedStats.LifeOnHit * self.EffectiveLifeMultiplier
	self.EffectiveLifeRegen = self.DerivedStats.LifeRegen * self.EffectiveLifeMultiplier

	return self
}

func (self *MetaStats) ComputeEffectiveLifeChangeForVitChange(vitChange float64) (effectiveLifeGain float64) {

	baseLifeChange := getLifeFromVitality(vitChange, self.DerivedStats.BaseStats.Level)
	actualLifeChange := baseLifeChange * (1 + self.DerivedStats.BaseStats.LifePercent)

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Life += actualLifeChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.EffectiveLife - self.EffectiveLife
}

func (self *MetaStats) ComputeEffectiveLifeChangeForPercentLifeChange(percentLifeChange float64) (effectiveLifeGain float64) {

	newPercentLifeBonus := 1 + self.DerivedStats.BaseStats.LifePercent + (percentLifeChange / 100.0)
	newLife := (getLifeFromVitality(self.DerivedStats.BaseStats.Vitality, self.DerivedStats.BaseStats.Level) * newPercentLifeBonus)

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Life = newLife

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.EffectiveLife - self.EffectiveLife
}

func (self *MetaStats) ComputeEffectiveLifeChangeForDexterityChange(dexterityChange float64) (effectiveLifeGain float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Dexterity += dexterityChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	return modifiedMetaStats.EffectiveLife - self.EffectiveLife
}

func (self *MetaStats) ComputeStatChangesForResistChange(resistChange float64) (effectiveLifeChange float64, effectiveLifeOnHitChange float64, effectiveLifeRegenChange float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.ResistArcane += resistChange
	modifiedDerivedStats.ResistFire += resistChange
	modifiedDerivedStats.ResistLightning += resistChange
	modifiedDerivedStats.ResistPoison += resistChange
	modifiedDerivedStats.ResistCold += resistChange
	modifiedDerivedStats.ResistPhysical += resistChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	effectiveLifeChange = modifiedMetaStats.EffectiveLife - self.EffectiveLife
	effectiveLifeOnHitChange = modifiedMetaStats.EffectiveLifeOnHit - self.EffectiveLifeOnHit
	effectiveLifeRegenChange = modifiedMetaStats.EffectiveLifeRegen - self.EffectiveLifeRegen

	return
}

func (self *MetaStats) ComputeStatChangesForArmorChange(armorChange float64) (effectiveLifeChange float64, effectiveLifeOnHitChange float64, effectiveLifeRegenChange float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Armor += armorChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	effectiveLifeChange = modifiedMetaStats.EffectiveLife - self.EffectiveLife
	effectiveLifeOnHitChange = modifiedMetaStats.EffectiveLifeOnHit - self.EffectiveLifeOnHit
	effectiveLifeRegenChange = modifiedMetaStats.EffectiveLifeRegen - self.EffectiveLifeRegen

	return
}

func (self *MetaStats) ComputeArmorEquivalentForResistChange(resistChange float64) (armorEquivalent float64) {

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.ResistArcane += resistChange
	modifiedDerivedStats.ResistFire += resistChange
	modifiedDerivedStats.ResistLightning += resistChange
	modifiedDerivedStats.ResistPoison += resistChange
	modifiedDerivedStats.ResistCold += resistChange
	modifiedDerivedStats.ResistPhysical += resistChange

	modifiedMetaStats := NewMetaStats(&modifiedDerivedStats)

	selfReductionFromArmor := self.MitigationSources[MitigationSourceArmor]
	selfReductionFromResistances := self.MitigationSources[MitigationSourceResistances]

	modifiedReductionFromResistances := modifiedMetaStats.MitigationSources[MitigationSourceResistances]

	// Used some algebra here...
	unmitigatedArmor := (1.0 - selfReductionFromArmor) * (1.0 - modifiedReductionFromResistances) / (1.0 - selfReductionFromResistances)

	mitigatedArmor := 1.0 - unmitigatedArmor

	armor := getArmorFromDr(mitigatedArmor, self.DerivedStats.BaseStats.Level)

	return armor - self.DerivedStats.Armor
}

func (self *MetaStats) CalculateStatChangeEffect(changeType string, changeValue float64) []StatChangeEffect {
	statChanges := make([]StatChangeEffect, 0, 5)

	var effectiveLifeChange float64 = 0
	var effectiveLifeOnHitChange float64 = 0
	var effectiveLifeRegenChange float64 = 0

	if changeType == urlValueCompareTypeVitality {
		effectiveLifeChange = self.ComputeEffectiveLifeChangeForVitChange(changeValue)
	} else if changeType == urlValueCompareTypeResist {
		effectiveLifeChange, effectiveLifeOnHitChange, effectiveLifeRegenChange = self.ComputeStatChangesForResistChange(changeValue)
	} else if changeType == urlValueCompareTypeArmor {
		effectiveLifeChange, effectiveLifeOnHitChange, effectiveLifeRegenChange = self.ComputeStatChangesForArmorChange(changeValue)
	} else if changeType == urlValueCompareTypePercentLife {
		effectiveLifeChange = self.ComputeEffectiveLifeChangeForPercentLifeChange(changeValue)
	} else if changeType == urlValueCompareTypeDexterity {
		effectiveLifeChange = self.ComputeEffectiveLifeChangeForDexterityChange(changeValue)
	}

	if effectiveLifeChange < -0.01 || effectiveLifeChange > 0.01 {

		changeInfo := new(StatChangeEffect)

		changeInfo.Name = "Effective Life"
		changeInfo.Value = fmt.Sprintf("%s%s", getSignForValue(effectiveLifeChange), getCommaLadenValue(effectiveLifeChange))
		changeInfo.Color = getColorForValue(effectiveLifeChange)

		statChanges = append(statChanges, *changeInfo)

	}

	if effectiveLifeOnHitChange < -0.01 || effectiveLifeOnHitChange > 0.01 {

		changeInfo := new(StatChangeEffect)

		changeInfo.Name = "Effective Life On Hit"
		changeInfo.Value = fmt.Sprintf("%s%s", getSignForValue(effectiveLifeOnHitChange), getCommaLadenValue(effectiveLifeOnHitChange))
		changeInfo.Color = getColorForValue(effectiveLifeOnHitChange)

		statChanges = append(statChanges, *changeInfo)

	}

	if effectiveLifeRegenChange < -0.01 || effectiveLifeRegenChange > 0.01 {

		changeInfo := new(StatChangeEffect)

		changeInfo.Name = "Effective Life Regen"
		changeInfo.Value = fmt.Sprintf("%s%s", getSignForValue(effectiveLifeRegenChange), getCommaLadenValue(effectiveLifeRegenChange))
		changeInfo.Color = getColorForValue(effectiveLifeRegenChange)

		statChanges = append(statChanges, *changeInfo)

	}
	return statChanges
}
