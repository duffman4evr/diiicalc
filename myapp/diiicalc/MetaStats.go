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
	MitigationSources       []MitigationSource
}

type MitigationSource struct {
	Name  string
	Value float64
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

	// Copy in some data.
	self.DerivedStats = *derivedStats

	// Find the minimum resistance that the user has.
	self.MinResistance = findMin(self.DerivedStats.ResistArcane, self.DerivedStats.ResistFire, self.DerivedStats.ResistLightning, self.DerivedStats.ResistPoison, self.DerivedStats.ResistCold, self.DerivedStats.ResistPhysical)

	// Calculate our average block amount.
	self.AverageBlockAmount = self.DerivedStats.BaseStats.BlockChance * ((self.DerivedStats.BaseStats.BlockAmountMin + self.DerivedStats.BaseStats.BlockAmountMax) / 2)

	// Calculate our different reduction sources.
	reductionFromArmor := self.DerivedStats.Armor / ((50.0 * self.DerivedStats.BaseStats.Level) + self.DerivedStats.Armor)
	reductionFromResistances := self.MinResistance / ((5.0 * self.DerivedStats.BaseStats.Level) + self.MinResistance)

	self.MitigationSources = append(self.MitigationSources, MitigationSource{MitigationSourceArmor, reductionFromArmor})
	self.MitigationSources = append(self.MitigationSources, MitigationSource{MitigationSourceResistances, reductionFromResistances})

	// Special case: Melee classes should have a reduction of 30%.
	if self.DerivedStats.BaseStats.HeroClass == urlValueHeroClassBarbarian || self.DerivedStats.BaseStats.HeroClass == urlValueHeroClassMonk {
		self.MitigationSources = append(self.MitigationSources, MitigationSource{MitigationSourceMeleeClass, 0.30})
	}

	// Special case: Demon Hunters and Monks should incorporate dodge into their mitigation.
	if self.DerivedStats.BaseStats.HeroClass == urlValueHeroClassDemonHunter || self.DerivedStats.BaseStats.HeroClass == urlValueHeroClassMonk {
		self.MitigationSources = append(self.MitigationSources, MitigationSource{MitigationSourceDodge, getDodgeChanceFromDexterity(self.DerivedStats.Dexterity)})
	}

	// Add any mitigation sources that might have been added by DerivedStats (Skills like Ignore Pain)
	for i := 0; i < len(self.DerivedStats.MitigationSources); i++ {
		self.MitigationSources = append(self.MitigationSources, self.DerivedStats.MitigationSources[i])
	}

	var totalNonMitigation float64 = 1.0

	for i := 0; i < len(self.MitigationSources); i++ {
		totalNonMitigation *= 1 - self.MitigationSources[i].Value
	}

	self.TotalMitigation = 1 - totalNonMitigation
	self.EffectiveLifeNoShield = self.DerivedStats.Life * (1 / (1 - self.TotalMitigation))
	self.EffectiveLifeMultiplier = (1 / (1 - self.TotalMitigation)) * (1 + (self.AverageBlockAmount / self.DerivedStats.Life))
	self.EffectiveLife = self.DerivedStats.Life * self.EffectiveLifeMultiplier
	self.EffectiveLifeOnHit = self.DerivedStats.LifeOnHit * self.EffectiveLifeMultiplier
	self.EffectiveLifeRegen = self.DerivedStats.LifeRegen * self.EffectiveLifeMultiplier

	return self
}

func (self *MetaStats) GetMitigationSource(name string) (mitigationSource *MitigationSource) {
	for i := 0; i < len(self.MitigationSources); i++ {
		if self.MitigationSources[i].Name == name {
			return &self.MitigationSources[i]
		}
	}
	return nil
}

func (self *MetaStats) GetMitigationSources() (mitigationSources *[]MitigationSource) {
	return &self.MitigationSources
}

func (self *MetaStats) ComputeEffectiveLifeChangeForVitChange(vitGain float64) (effectiveLifeGain float64) {

	var lifePerVit float64

	if self.DerivedStats.BaseStats.Level < 35 {
		lifePerVit = 10
	} else {
		lifePerVit = self.DerivedStats.BaseStats.Level - 25
	}

	modifiedDerivedStats := self.DerivedStats

	modifiedDerivedStats.Life += (lifePerVit * vitGain)

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

	selfReductionFromArmor := self.GetMitigationSource(MitigationSourceArmor).Value
	selfReductionFromResistances := self.GetMitigationSource(MitigationSourceResistances).Value

	modifiedReductionFromResistances := modifiedMetaStats.GetMitigationSource(MitigationSourceResistances).Value

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
