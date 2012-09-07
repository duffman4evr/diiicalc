package diiicalc

import (
	"fmt"
	"net/http"
)

const (
	standardUrlValueOff = "off"
	standardUrlValueOn  = "on"

	deadlyReachUrlKey    = "deadlyReach"
	deadlyReachSkillSlug = "deadly-reach"

	energyArmorUrlKey    = "energyArmor"
	energyArmorSkillSlug = "energy-armor"

	fistsOfThunderUrlKey    = "fistsOfThunder"
	fistsOfThunderSkillSlug = "fists-of-thunder"

	glassCannonUrlKey    = "glassCannon"
	glassCannonSkillSlug = "glass-cannon"

	guardiansPathUrlKey    = "guardiansPath"
	guardiansPathSkillSlug = "the-guardians-path"

	horrifyUrlKey    = "horrify"
	horrifySkillSlug = "horrify"

	ignorePainUrlKey    = "ignorePain"
	ignorePainSkillSlug = "ignore-pain"

	jungleFortitudeUrlKey    = "jungleFortitude"
	jungleFortitudeSkillSlug = "jungle-fortitude"

	leapUrlKey    = "leap"
	leapSkillSlug = "leap"

	mantraOfEvasionUrlKey    = "mantraOfEvasion"
	mantraOfEvasionSkillSlug = "mantra-of-evasion"

	mantraOfHealingUrlKey    = "mantraOfHealing"
	mantraOfHealingSkillSlug = "mantra-of-healing"

	nervesOfSteelUrlKey    = "nervesOfSteel"
	nervesOfSteelSkillSlug = "nerves-of-steel"

	oneWithEverythingUrlKey    = "oneWithEverything"
	oneWithEverythingSkillSlug = "one-with-everything"

	poweredArmorUrlKey    = "poweredArmor"
	poweredArmorSkillSlug = ""

	seizeTheInitiativeUrlKey    = "seizeTheInitiative"
	seizeTheInitiativeSkillSlug = "seize-the-initiative"

	toughAsNailsUrlKey    = "toughAsNails"
	toughAsNailsSkillSlug = "tough-as-nails"

	warCryUrlKey    = "warCry"
	warCrySkillSlug = "war-cry"
)

var (
	emptyRuneSlugs           = []string{}
	deadlyReachRuneSlugs     = []string{"deadly-reach-e"}
	energyArmorRuneSlugs     = []string{"energy-armor-a"}
	fistsOfThunderRuneSlugs  = []string{"fists-of-thunder-e"}
	horrifyRuneSlugs         = []string{"horrify-a"}
	leapRuneSlugs            = []string{"leap-d"}
	mantraOfHealingRuneSlugs = []string{"mantra-of-healing-e"}
	mantraOfEvasionRuneSlugs = []string{"spam", "mantra-of-evasion-e", "hard-target-spam"}
	warCryRuneSlugs          = []string{"war-cry-a", "war-cry-c"}
)

type SkillChoice interface {
	ModifyDerivedStats(derivedStats *DerivedStats)
	PrintHtml(w http.ResponseWriter)
	GetValue() string
	SetValue(value string)
	GetUrlKey() string
	GetSkillSlug() string
	GetRuneSlugs() []string
}

func InitializeSkillChoice(skillChoice SkillChoice, r *http.Request) {

	// If there is already a value for this SkillChoice's URL key, then continue using it!
	urlValue := r.FormValue(skillChoice.GetUrlKey())

	if urlValue != "" {
		skillChoice.SetValue(urlValue)
		return
	}

	// If not, default the skill to off.
	skillChoice.SetValue(standardUrlValueOff)

	// If there is no value for this SkillChoice's URL key, then we must parse the 
	// skills from the user's build to see if this skill is being used or not.

	// First, check passive skills.
	for i := 0; i < len(urlKeysPassiveSkills); i++ {
		if skillChoice.GetSkillSlug() == r.FormValue(urlKeysPassiveSkills[i]) {
			skillChoice.SetValue(standardUrlValueOn)
			return
		}
	}

	// Now, check active skills and runes.
	for i := 0; i < len(urlKeysActiveSkills); i++ {

		if skillChoice.GetSkillSlug() != r.FormValue(urlKeysActiveSkills[i]) {
			continue
		}

		// If we got to this point, that means we matched up a user's active skill choice
		// in their build to a SkillChoice we recognize for our calculator.

		// Start by defaulting the skill to 'on'.
		skillChoice.SetValue(standardUrlValueOn)

		// Furthermore, check if we have a match on any of the runes we care about.
		userRuneSlug := r.FormValue(urlKeysActiveRunes[i])
		supportedRuneSlugs := skillChoice.GetRuneSlugs()

		for j := 0; j < len(supportedRuneSlugs); j++ {
			// If we have a match, set the URL value for this SkillChoice to match
			// the given rune slug. 
			if userRuneSlug == supportedRuneSlugs[j] {
				skillChoice.SetValue(userRuneSlug)
				return
			}
		}
	}
}

func GetSelected(skillChoice SkillChoice, urlValue string) (retVal string) {
	if skillChoice.GetValue() == urlValue {
		retVal = "selected"
	} else {
		retVal = ""
	}
	return
}

func ParseSkillChoices(r *http.Request) (skillChoices []SkillChoice) {

	skillChoices = make([]SkillChoice, 0, 10)

	heroClass := r.FormValue(urlKeyHeroClass)

	if heroClass == urlValueHeroClassBarbarian {

		var (
			toughAsNailsSkillChoice  = new(ToughAsNailsSkillChoice)
			nervesOfSteelSkillChoice = new(NervesOfSteelSkillChoice)
			warCrySkillChoice        = new(WarCrySkillChoice)
			leapSkillChoice          = new(LeapSkillChoice)
			ignorePainSkillChoice    = new(IgnorePainSkillChoice)
			poweredArmorSkillChoice  = new(PoweredArmorSkillChoice)
		)

		InitializeSkillChoice(toughAsNailsSkillChoice, r)
		InitializeSkillChoice(nervesOfSteelSkillChoice, r)
		InitializeSkillChoice(warCrySkillChoice, r)
		InitializeSkillChoice(leapSkillChoice, r)
		InitializeSkillChoice(ignorePainSkillChoice, r)
		InitializeSkillChoice(poweredArmorSkillChoice, r)

		skillChoices = append(skillChoices, toughAsNailsSkillChoice)
		skillChoices = append(skillChoices, nervesOfSteelSkillChoice)
		skillChoices = append(skillChoices, warCrySkillChoice)
		skillChoices = append(skillChoices, leapSkillChoice)
		skillChoices = append(skillChoices, ignorePainSkillChoice)
		skillChoices = append(skillChoices, poweredArmorSkillChoice)

	} else if heroClass == urlValueHeroClassMonk {

		var (
			seizeTheInitiativeSkillChoice = new(SeizeTheInitiativeSkillChoice)
			oneWithEverythingSkillChoice  = new(OneWithEverythingSkillChoice)
			deadlyReachSkillChoice        = new(DeadlyReachSkillChoice)
			mantraOfEvasionSkillChoice    = new(MantraOfEvasionSkillChoice)
			mantraOfHealingSkillChoice    = new(MantraOfHealingSkillChoice)
			fistsOfThunderSkillChoice     = new(FistsOfThunderSkillChoice)
			guardiansPathSkillChoice      = new(GuardiansPathSkillChoice)
			poweredArmorSkillChoice       = new(PoweredArmorSkillChoice)
		)

		InitializeSkillChoice(seizeTheInitiativeSkillChoice, r)
		InitializeSkillChoice(oneWithEverythingSkillChoice, r)
		InitializeSkillChoice(deadlyReachSkillChoice, r)
		InitializeSkillChoice(mantraOfEvasionSkillChoice, r)
		InitializeSkillChoice(mantraOfHealingSkillChoice, r)
		InitializeSkillChoice(fistsOfThunderSkillChoice, r)
		InitializeSkillChoice(guardiansPathSkillChoice, r)
		InitializeSkillChoice(poweredArmorSkillChoice, r)

		skillChoices = append(skillChoices, seizeTheInitiativeSkillChoice)
		skillChoices = append(skillChoices, oneWithEverythingSkillChoice)
		skillChoices = append(skillChoices, deadlyReachSkillChoice)
		skillChoices = append(skillChoices, mantraOfEvasionSkillChoice)
		skillChoices = append(skillChoices, mantraOfHealingSkillChoice)
		skillChoices = append(skillChoices, fistsOfThunderSkillChoice)
		skillChoices = append(skillChoices, guardiansPathSkillChoice)
		skillChoices = append(skillChoices, poweredArmorSkillChoice)

	} else if heroClass == urlValueHeroClassWizard {

		var (
			energyArmorSkillChoice  = new(EnergyArmorSkillChoice)
			glassCannonSkillChoice  = new(GlassCannonSkillChoice)
			poweredArmorSkillChoice = new(PoweredArmorSkillChoice)
		)

		InitializeSkillChoice(energyArmorSkillChoice, r)
		InitializeSkillChoice(glassCannonSkillChoice, r)
		InitializeSkillChoice(poweredArmorSkillChoice, r)

		skillChoices = append(skillChoices, energyArmorSkillChoice)
		skillChoices = append(skillChoices, glassCannonSkillChoice)
		skillChoices = append(skillChoices, poweredArmorSkillChoice)

	} else if heroClass == urlValueHeroClassDemonHunter {

		poweredArmorSkillChoice := new(PoweredArmorSkillChoice)

		InitializeSkillChoice(poweredArmorSkillChoice, r)

		skillChoices = append(skillChoices, poweredArmorSkillChoice)

	} else if heroClass == urlValueHeroClassWitchDoctor {

		var (
			horrifySkillChoice         = new(HorrifySkillChoice)
			jungleFortitudeSkillChoice = new(JungleFortitudeSkillChoice)
			poweredArmorSkillChoice    = new(PoweredArmorSkillChoice)
		)

		InitializeSkillChoice(horrifySkillChoice, r)
		InitializeSkillChoice(jungleFortitudeSkillChoice, r)
		InitializeSkillChoice(poweredArmorSkillChoice, r)

		skillChoices = append(skillChoices, horrifySkillChoice)
		skillChoices = append(skillChoices, jungleFortitudeSkillChoice)
		skillChoices = append(skillChoices, poweredArmorSkillChoice)

	}

	return skillChoices
}

type DeadlyReachSkillChoice struct {
	Value string
}
type EnergyArmorSkillChoice struct {
	Value string
}
type FistsOfThunderSkillChoice struct {
	Value string
}
type GlassCannonSkillChoice struct {
	Value string
}
type GuardiansPathSkillChoice struct {
	Value string
}
type HorrifySkillChoice struct {
	Value string
}
type IgnorePainSkillChoice struct {
	Value string
}
type JungleFortitudeSkillChoice struct {
	Value string
}
type LeapSkillChoice struct {
	Value string
}
type MantraOfEvasionSkillChoice struct {
	Value string
}
type MantraOfHealingSkillChoice struct {
	Value string
}
type NervesOfSteelSkillChoice struct {
	Value string
}
type OneWithEverythingSkillChoice struct {
	Value string
}
type PoweredArmorSkillChoice struct {
	Value string
}
type SeizeTheInitiativeSkillChoice struct {
	Value string
}
type ToughAsNailsSkillChoice struct {
	Value string
}
type WarCrySkillChoice struct {
	Value string
}

func (self *DeadlyReachSkillChoice) GetValue() string {
	return self.Value
}
func (self *EnergyArmorSkillChoice) GetValue() string {
	return self.Value
}
func (self *FistsOfThunderSkillChoice) GetValue() string {
	return self.Value
}
func (self *GlassCannonSkillChoice) GetValue() string {
	return self.Value
}
func (self *GuardiansPathSkillChoice) GetValue() string {
	return self.Value
}
func (self *HorrifySkillChoice) GetValue() string {
	return self.Value
}
func (self *IgnorePainSkillChoice) GetValue() string {
	return self.Value
}
func (self *JungleFortitudeSkillChoice) GetValue() string {
	return self.Value
}
func (self *LeapSkillChoice) GetValue() string {
	return self.Value
}
func (self *MantraOfEvasionSkillChoice) GetValue() string {
	return self.Value
}
func (self *MantraOfHealingSkillChoice) GetValue() string {
	return self.Value
}
func (self *NervesOfSteelSkillChoice) GetValue() string {
	return self.Value
}
func (self *OneWithEverythingSkillChoice) GetValue() string {
	return self.Value
}
func (self *PoweredArmorSkillChoice) GetValue() string {
	return self.Value
}
func (self *SeizeTheInitiativeSkillChoice) GetValue() string {
	return self.Value
}
func (self *ToughAsNailsSkillChoice) GetValue() string {
	return self.Value
}
func (self *WarCrySkillChoice) GetValue() string {
	return self.Value
}

func (self *DeadlyReachSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *EnergyArmorSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *FistsOfThunderSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *GlassCannonSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *GuardiansPathSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *HorrifySkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *IgnorePainSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *JungleFortitudeSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *LeapSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *MantraOfEvasionSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *MantraOfHealingSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *NervesOfSteelSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *OneWithEverythingSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *PoweredArmorSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *SeizeTheInitiativeSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *ToughAsNailsSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *WarCrySkillChoice) SetValue(value string) {
	self.Value = value
}

func (self *DeadlyReachSkillChoice) GetUrlKey() string {
	return deadlyReachUrlKey
}
func (self *EnergyArmorSkillChoice) GetUrlKey() string {
	return energyArmorUrlKey
}
func (self *FistsOfThunderSkillChoice) GetUrlKey() string {
	return fistsOfThunderUrlKey
}
func (self *GlassCannonSkillChoice) GetUrlKey() string {
	return glassCannonUrlKey
}
func (self *GuardiansPathSkillChoice) GetUrlKey() string {
	return guardiansPathUrlKey
}
func (self *HorrifySkillChoice) GetUrlKey() string {
	return horrifyUrlKey
}
func (self *IgnorePainSkillChoice) GetUrlKey() string {
	return ignorePainUrlKey
}
func (self *JungleFortitudeSkillChoice) GetUrlKey() string {
	return jungleFortitudeUrlKey
}
func (self *LeapSkillChoice) GetUrlKey() string {
	return leapUrlKey
}
func (self *MantraOfEvasionSkillChoice) GetUrlKey() string {
	return mantraOfEvasionUrlKey
}
func (self *MantraOfHealingSkillChoice) GetUrlKey() string {
	return mantraOfHealingUrlKey
}
func (self *NervesOfSteelSkillChoice) GetUrlKey() string {
	return nervesOfSteelUrlKey
}
func (self *OneWithEverythingSkillChoice) GetUrlKey() string {
	return oneWithEverythingUrlKey
}
func (self *PoweredArmorSkillChoice) GetUrlKey() string {
	return poweredArmorUrlKey
}
func (self *SeizeTheInitiativeSkillChoice) GetUrlKey() string {
	return seizeTheInitiativeUrlKey
}
func (self *ToughAsNailsSkillChoice) GetUrlKey() string {
	return toughAsNailsUrlKey
}
func (self *WarCrySkillChoice) GetUrlKey() string {
	return warCryUrlKey
}

func (self *DeadlyReachSkillChoice) GetSkillSlug() string {
	return deadlyReachSkillSlug
}
func (self *EnergyArmorSkillChoice) GetSkillSlug() string {
	return energyArmorSkillSlug
}
func (self *FistsOfThunderSkillChoice) GetSkillSlug() string {
	return fistsOfThunderSkillSlug
}
func (self *GlassCannonSkillChoice) GetSkillSlug() string {
	return glassCannonSkillSlug
}
func (self *GuardiansPathSkillChoice) GetSkillSlug() string {
	return guardiansPathSkillSlug
}
func (self *HorrifySkillChoice) GetSkillSlug() string {
	return horrifySkillSlug
}
func (self *IgnorePainSkillChoice) GetSkillSlug() string {
	return ignorePainSkillSlug
}
func (self *JungleFortitudeSkillChoice) GetSkillSlug() string {
	return jungleFortitudeSkillSlug
}
func (self *LeapSkillChoice) GetSkillSlug() string {
	return leapSkillSlug
}
func (self *MantraOfEvasionSkillChoice) GetSkillSlug() string {
	return mantraOfEvasionSkillSlug
}
func (self *MantraOfHealingSkillChoice) GetSkillSlug() string {
	return mantraOfHealingSkillSlug
}
func (self *NervesOfSteelSkillChoice) GetSkillSlug() string {
	return nervesOfSteelSkillSlug
}
func (self *OneWithEverythingSkillChoice) GetSkillSlug() string {
	return oneWithEverythingSkillSlug
}
func (self *PoweredArmorSkillChoice) GetSkillSlug() string {
	return poweredArmorSkillSlug
}
func (self *SeizeTheInitiativeSkillChoice) GetSkillSlug() string {
	return seizeTheInitiativeSkillSlug
}
func (self *ToughAsNailsSkillChoice) GetSkillSlug() string {
	return toughAsNailsSkillSlug
}
func (self *WarCrySkillChoice) GetSkillSlug() string {
	return warCrySkillSlug
}

func (self *DeadlyReachSkillChoice) GetRuneSlugs() []string {
	return deadlyReachRuneSlugs
}
func (self *EnergyArmorSkillChoice) GetRuneSlugs() []string {
	return energyArmorRuneSlugs
}
func (self *FistsOfThunderSkillChoice) GetRuneSlugs() []string {
	return fistsOfThunderRuneSlugs
}
func (self *GlassCannonSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *GuardiansPathSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *HorrifySkillChoice) GetRuneSlugs() []string {
	return horrifyRuneSlugs
}
func (self *IgnorePainSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *JungleFortitudeSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *LeapSkillChoice) GetRuneSlugs() []string {
	return leapRuneSlugs
}
func (self *MantraOfEvasionSkillChoice) GetRuneSlugs() []string {
	return mantraOfEvasionRuneSlugs
}
func (self *MantraOfHealingSkillChoice) GetRuneSlugs() []string {
	return mantraOfHealingRuneSlugs
}
func (self *NervesOfSteelSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *OneWithEverythingSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *PoweredArmorSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *SeizeTheInitiativeSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *ToughAsNailsSkillChoice) GetRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *WarCrySkillChoice) GetRuneSlugs() []string {
	return warCryRuneSlugs
}

func (self *DeadlyReachSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == deadlyReachRuneSlugs[0] {
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.5
	}
}
func (self *EnergyArmorSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	switch {
	case self.Value == standardUrlValueOn:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.65

	case self.Value == energyArmorRuneSlugs[0]:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.65
		derivedStats.ResistArcane += derivedStats.BaseStats.ResistArcane * 0.4
		derivedStats.ResistFire += derivedStats.BaseStats.ResistFire * 0.4
		derivedStats.ResistLightning += derivedStats.BaseStats.ResistLightning * 0.4
		derivedStats.ResistPoison += derivedStats.BaseStats.ResistPoison * 0.4
		derivedStats.ResistCold += derivedStats.BaseStats.ResistCold * 0.4
		derivedStats.ResistPhysical += derivedStats.BaseStats.ResistPhysical * 0.4
	}
}
func (self *FistsOfThunderSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == fistsOfThunderRuneSlugs[0] {
		addDodge(0.16, &derivedStats.MitigationSources)
	}
}
func (self *GlassCannonSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.Armor += derivedStats.BaseStats.Armor * -0.1
		derivedStats.ResistArcane += derivedStats.BaseStats.ResistArcane * -0.1
		derivedStats.ResistFire += derivedStats.BaseStats.ResistFire * -0.1
		derivedStats.ResistLightning += derivedStats.BaseStats.ResistLightning * -0.1
		derivedStats.ResistPoison += derivedStats.BaseStats.ResistPoison * -0.1
		derivedStats.ResistCold += derivedStats.BaseStats.ResistCold * -0.1
		derivedStats.ResistPhysical += derivedStats.BaseStats.ResistPhysical * -0.1
	}
}
func (self *GuardiansPathSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		addDodge(0.15, &derivedStats.MitigationSources)
	}
}
func (self *HorrifySkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == horrifyRuneSlugs[0] {
		derivedStats.Armor += derivedStats.BaseStats.Armor
	}
}
func (self *IgnorePainSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.MitigationSources["Ignore Pain"] = 0.65
	}
}
func (self *LeapSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == leapRuneSlugs[0] {
		derivedStats.Armor += derivedStats.BaseStats.Armor * 3.0
	}
}
func (self *MantraOfEvasionSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	switch {
	case self.Value == standardUrlValueOn:
		addDodge(0.15, &derivedStats.MitigationSources)

	case self.Value == mantraOfEvasionRuneSlugs[0]: // Spam
		addDodge(0.30, &derivedStats.MitigationSources)

	case self.Value == mantraOfEvasionRuneSlugs[1]: // Hard Target
		addDodge(0.15, &derivedStats.MitigationSources)
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.2

	case self.Value == mantraOfEvasionRuneSlugs[2]: // Hard Target Spam
		addDodge(0.30, &derivedStats.MitigationSources)
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.2
	}
}
func (self *MantraOfHealingSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == mantraOfHealingRuneSlugs[0] {
		derivedStats.ResistArcane += derivedStats.BaseStats.ResistArcane * 0.2
		derivedStats.ResistFire += derivedStats.BaseStats.ResistFire * 0.2
		derivedStats.ResistLightning += derivedStats.BaseStats.ResistLightning * 0.2
		derivedStats.ResistPoison += derivedStats.BaseStats.ResistPoison * 0.2
		derivedStats.ResistCold += derivedStats.BaseStats.ResistCold * 0.2
		derivedStats.ResistPhysical += derivedStats.BaseStats.ResistPhysical * 0.2
	}
}
func (self *NervesOfSteelSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.Armor += float64(derivedStats.Vitality)
	}
}
func (self *OneWithEverythingSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		max := findMax(derivedStats.ResistArcane, derivedStats.ResistFire, derivedStats.ResistLightning, derivedStats.ResistPoison, derivedStats.ResistCold, derivedStats.ResistPhysical)
		derivedStats.ResistArcane = max
		derivedStats.ResistFire = max
		derivedStats.ResistLightning = max
		derivedStats.ResistPoison = max
		derivedStats.ResistCold = max
		derivedStats.ResistPhysical = max
	}
}
func (self *PoweredArmorSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.15
	}
}
func (self *SeizeTheInitiativeSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.Armor += float64(derivedStats.Dexterity)
	}
}
func (self *ToughAsNailsSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.25
	}
}
func (self *JungleFortitudeSkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.MitigationSources["Jungle Fortitude"] = 0.20
	}
}
func (self *WarCrySkillChoice) ModifyDerivedStats(derivedStats *DerivedStats) {
	switch {
	case self.Value == standardUrlValueOn:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.2

	case self.Value == warCryRuneSlugs[0]:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.4

	case self.Value == warCryRuneSlugs[1]:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.2
		derivedStats.ResistFire += derivedStats.BaseStats.ResistFire * 0.5
		derivedStats.ResistLightning += derivedStats.BaseStats.ResistLightning * 0.5
		derivedStats.ResistPoison += derivedStats.BaseStats.ResistPoison * 0.5
		derivedStats.ResistCold += derivedStats.BaseStats.ResistCold * 0.5
		derivedStats.ResistArcane += derivedStats.BaseStats.ResistArcane * 0.5
		derivedStats.ResistPhysical += derivedStats.BaseStats.ResistPhysical * 0.5

	}
}

func (self *DeadlyReachSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Deadly Reach:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Keen Eye</option>%s`, deadlyReachRuneSlugs[0], GetSelected(self, deadlyReachRuneSlugs[0]), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *EnergyArmorSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Energy Armor:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Prismatic</option>%s`, energyArmorRuneSlugs[0], GetSelected(self, energyArmorRuneSlugs[0]), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *FistsOfThunderSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Fists of Thunder:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Lightning Flash</option>%s`, fistsOfThunderRuneSlugs[0], GetSelected(self, fistsOfThunderRuneSlugs[0]), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *GlassCannonSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Glass Cannon:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *GuardiansPathSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Guardian's Path:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *HorrifySkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Horrify:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Frightening Aspect</option>%s`, horrifyRuneSlugs[0], GetSelected(self, horrifyRuneSlugs[0]), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *IgnorePainSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Ignore Pain:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *JungleFortitudeSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Jungle Fortitude:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *LeapSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Leap:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Iron Impact</option>%s`, leapRuneSlugs[0], GetSelected(self, leapRuneSlugs[0]), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *MantraOfEvasionSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Mantra of Evasion:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Spam</option>%s`, mantraOfEvasionRuneSlugs[0], GetSelected(self, mantraOfEvasionRuneSlugs[0]), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Hard Target</option>%s`, mantraOfEvasionRuneSlugs[1], GetSelected(self, mantraOfEvasionRuneSlugs[1]), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Hard Target Spam</option>%s`, mantraOfEvasionRuneSlugs[2], GetSelected(self, mantraOfEvasionRuneSlugs[2]), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *MantraOfHealingSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Mantra of Healing:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Time of Need</option>%s`, mantraOfHealingRuneSlugs[0], GetSelected(self, mantraOfHealingRuneSlugs[0]), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *NervesOfSteelSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Nerves of Steel:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *OneWithEverythingSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">One With Everything:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *PoweredArmorSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Enchantress Armor:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *SeizeTheInitiativeSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Seize the Initiative:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *ToughAsNailsSkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Tough as Nails:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
func (self *WarCrySkillChoice) PrintHtml(w http.ResponseWriter) {

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">War Cry:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, self.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(self, standardUrlValueOff), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(self, standardUrlValueOn), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Hardened Wrath</option>%s`, warCryRuneSlugs[0], GetSelected(self, warCryRuneSlugs[0]), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Impunity</option>%s`, warCryRuneSlugs[1], GetSelected(self, warCryRuneSlugs[1]), "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

}
