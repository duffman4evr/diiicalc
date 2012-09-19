package diiicalc

import (
	"diiicalc/defensive"
	"diiicalc/offensive"
	"diiicalc/util"
	"fmt"
	"net/http"
)

const (
	standardUrlValueOff = "off"
	standardUrlValueOn  = "on"

	bashUrlKey                   = "bash"
	bashSkillSlug                = "bash"
	battleRageUrlKey             = "battleRage"
	battleRageSkillSlug          = "battle-rage"
	berzerkerRageUrlKey          = "berzerkerRage"
	berzerkerRageSkillSlug       = "berzerker-rage"
	brawlerUrlKey                = "brawler"
	brawlerSkillSlug             = "brawler"
	deadlyReachUrlKey            = "deadlyReach"
	deadlyReachSkillSlug         = "deadly-reach"
	energyArmorUrlKey            = "energyArmor"
	energyArmorSkillSlug         = "energy-armor"
	fistsOfThunderUrlKey         = "fistsOfThunder"
	fistsOfThunderSkillSlug      = "fists-of-thunder"
	frenzyUrlKey                 = "frenzy"
	frenzySkillSlug              = "frenzy"
	glassCannonUrlKey            = "glassCannon"
	glassCannonSkillSlug         = "glass-cannon"
	guardiansPathUrlKey          = "guardiansPath"
	guardiansPathSkillSlug       = "the-guardians-path"
	horrifyUrlKey                = "horrify"
	horrifySkillSlug             = "horrify"
	ignorePainUrlKey             = "ignorePain"
	ignorePainSkillSlug          = "ignore-pain"
	jungleFortitudeUrlKey        = "jungleFortitude"
	jungleFortitudeSkillSlug     = "jungle-fortitude"
	leapUrlKey                   = "leap"
	leapSkillSlug                = "leap"
	mantraOfEvasionUrlKey        = "mantraOfEvasion"
	mantraOfEvasionSkillSlug     = "mantra-of-evasion"
	mantraOfHealingUrlKey        = "mantraOfHealing"
	mantraOfHealingSkillSlug     = "mantra-of-healing"
	nervesOfSteelUrlKey          = "nervesOfSteel"
	nervesOfSteelSkillSlug       = "nerves-of-steel"
	oneWithEverythingUrlKey      = "oneWithEverything"
	oneWithEverythingSkillSlug   = "one-with-everything"
	overpowerUrlKey              = "overpower"
	overpowerSkillSlug           = "overpower"
	poweredArmorUrlKey           = "poweredArmor"
	poweredArmorSkillSlug        = ""
	revengeUrlKey                = "revenge"
	revengeSkillSlug             = "revenge"
	ruthlessUrlKey               = "ruthless"
	ruthlessSkillSlug            = "ruthless"
	seizeTheInitiativeUrlKey     = "seizeTheInitiative"
	seizeTheInitiativeSkillSlug  = "seize-the-initiative"
	toughAsNailsUrlKey           = "toughAsNails"
	toughAsNailsSkillSlug        = "tough-as-nails"
	warCryUrlKey                 = "warCry"
	warCrySkillSlug              = "war-cry"
	weaponsMasterUrlKey          = "weaponsMaster"
	weaponsMasterSkillSlug       = "weapons-master"
	wrathOfTheBerzerkerUrlKey    = "wrathOfTheBerzerker"
	wrathOfTheBerzerkerSkillSlug = "wrath-of-the-berzerker"
)

var emptyRuneSlugs = []string{}

type SkillChoice interface {
	GetValue() string
	SetValue(value string)
	GetUrlKey() string
	GetSkillSlug() string
}

type DefensiveSkillChoice interface {
	ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats)
	PrintDefensiveHtml(w http.ResponseWriter)
	GetDefensiveRuneSlugs() []string
	GetValue() string
	SetValue(value string)
	GetUrlKey() string
	GetSkillSlug() string
}

type OffensiveSkillChoice interface {
	ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats)
	PrintOffensiveHtml(w http.ResponseWriter)
	GetOffensiveRuneSlugs() []string
	GetValue() string
	SetValue(value string)
	GetUrlKey() string
	GetSkillSlug() string
}

func InitializeDefensiveSkillChoice(skillChoice DefensiveSkillChoice, r *http.Request) {
	InitializeSkillChoice(skillChoice, skillChoice.GetDefensiveRuneSlugs(), r)
}

func InitializeOffensiveSkillChoice(skillChoice OffensiveSkillChoice, r *http.Request) {
	InitializeSkillChoice(skillChoice, skillChoice.GetOffensiveRuneSlugs(), r)
}

func InitializeSkillChoice(skillChoice SkillChoice, supportedRuneSlugs []string, r *http.Request) {
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
	for i := 0; i < len(util.UrlKeysPassiveSkills); i++ {
		if skillChoice.GetSkillSlug() == r.FormValue(util.UrlKeysPassiveSkills[i]) {
			skillChoice.SetValue(standardUrlValueOn)
			return
		}
	}

	// Now, check active skills and runes.
	for i := 0; i < len(util.UrlKeysActiveSkills); i++ {

		if skillChoice.GetSkillSlug() != r.FormValue(util.UrlKeysActiveSkills[i]) {
			continue
		}

		// If we got to this point, that means we matched up a user's active skill choice
		// in their build to a SkillChoice we recognize for our calculator.

		// Start by defaulting the skill to 'on'.
		skillChoice.SetValue(standardUrlValueOn)

		// Furthermore, check if we have a match on any of the runes we care about.
		userRuneSlug := r.FormValue(util.UrlKeysActiveRunes[i])

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

func ParseOffensiveSkillChoices(r *http.Request) (offensiveSkillChoices []OffensiveSkillChoice) {

	offensiveSkillChoices = make([]OffensiveSkillChoice, 0, 10)

	heroClass := r.FormValue(util.UrlKeyHeroClass)

	if heroClass == util.UrlValueHeroClassBarbarian {

		offensiveSkillChoices = append(offensiveSkillChoices, new(BashSkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(BattleRageSkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(BerzerkerRageSkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(BrawlerSkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(FrenzySkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(OverpowerSkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(RevengeSkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(RuthlessSkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(WeaponsMasterSkillChoice))
		offensiveSkillChoices = append(offensiveSkillChoices, new(WrathOfTheBerzerkerSkillChoice))

	} else if heroClass == util.UrlValueHeroClassMonk {

	} else if heroClass == util.UrlValueHeroClassWizard {

	} else if heroClass == util.UrlValueHeroClassDemonHunter {

	} else if heroClass == util.UrlValueHeroClassWitchDoctor {

	}

	for _, skillChoice := range offensiveSkillChoices {
		InitializeOffensiveSkillChoice(skillChoice, r)
	}

	return offensiveSkillChoices
}

// TODO do same thing as ^ down here.
func ParseDefensiveSkillChoices(r *http.Request) (defensiveSkillChoices []DefensiveSkillChoice) {

	defensiveSkillChoices = make([]DefensiveSkillChoice, 0, 10)

	heroClass := r.FormValue(util.UrlKeyHeroClass)

	if heroClass == util.UrlValueHeroClassBarbarian {

		var (
			toughAsNailsSkillChoice  = new(ToughAsNailsSkillChoice)
			nervesOfSteelSkillChoice = new(NervesOfSteelSkillChoice)
			leapSkillChoice          = new(LeapSkillChoice)
			ignorePainSkillChoice    = new(IgnorePainSkillChoice)
		)

		InitializeDefensiveSkillChoice(toughAsNailsSkillChoice, r)
		InitializeDefensiveSkillChoice(nervesOfSteelSkillChoice, r)
		InitializeDefensiveSkillChoice(leapSkillChoice, r)
		InitializeDefensiveSkillChoice(ignorePainSkillChoice, r)

		defensiveSkillChoices = append(defensiveSkillChoices, toughAsNailsSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, nervesOfSteelSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, leapSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, ignorePainSkillChoice)

	} else if heroClass == util.UrlValueHeroClassMonk {

		var (
			seizeTheInitiativeSkillChoice = new(SeizeTheInitiativeSkillChoice)
			oneWithEverythingSkillChoice  = new(OneWithEverythingSkillChoice)
			deadlyReachSkillChoice        = new(DeadlyReachSkillChoice)
			mantraOfHealingSkillChoice    = new(MantraOfHealingSkillChoice)
			fistsOfThunderSkillChoice     = new(FistsOfThunderSkillChoice)
			guardiansPathSkillChoice      = new(GuardiansPathSkillChoice)
		)

		InitializeDefensiveSkillChoice(seizeTheInitiativeSkillChoice, r)
		InitializeDefensiveSkillChoice(oneWithEverythingSkillChoice, r)
		InitializeDefensiveSkillChoice(deadlyReachSkillChoice, r)
		InitializeDefensiveSkillChoice(mantraOfHealingSkillChoice, r)
		InitializeDefensiveSkillChoice(fistsOfThunderSkillChoice, r)
		InitializeDefensiveSkillChoice(guardiansPathSkillChoice, r)

		defensiveSkillChoices = append(defensiveSkillChoices, seizeTheInitiativeSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, oneWithEverythingSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, deadlyReachSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, mantraOfHealingSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, fistsOfThunderSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, guardiansPathSkillChoice)

	} else if heroClass == util.UrlValueHeroClassWizard {

		var (
			energyArmorSkillChoice = new(EnergyArmorSkillChoice)
			glassCannonSkillChoice = new(GlassCannonSkillChoice)
		)

		InitializeDefensiveSkillChoice(energyArmorSkillChoice, r)
		InitializeDefensiveSkillChoice(glassCannonSkillChoice, r)

		defensiveSkillChoices = append(defensiveSkillChoices, energyArmorSkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, glassCannonSkillChoice)

	} else if heroClass == util.UrlValueHeroClassDemonHunter {

		// lol Demon Hunter...

	} else if heroClass == util.UrlValueHeroClassWitchDoctor {

		var (
			horrifySkillChoice         = new(HorrifySkillChoice)
			jungleFortitudeSkillChoice = new(JungleFortitudeSkillChoice)
		)

		InitializeDefensiveSkillChoice(horrifySkillChoice, r)
		InitializeDefensiveSkillChoice(jungleFortitudeSkillChoice, r)

		defensiveSkillChoices = append(defensiveSkillChoices, horrifySkillChoice)
		defensiveSkillChoices = append(defensiveSkillChoices, jungleFortitudeSkillChoice)

	}

	var (
		warCrySkillChoice          = new(WarCrySkillChoice)
		mantraOfEvasionSkillChoice = new(MantraOfEvasionSkillChoice)
		poweredArmorSkillChoice    = new(PoweredArmorSkillChoice)
	)

	InitializeDefensiveSkillChoice(warCrySkillChoice, r)
	InitializeDefensiveSkillChoice(mantraOfEvasionSkillChoice, r)
	InitializeDefensiveSkillChoice(poweredArmorSkillChoice, r)

	defensiveSkillChoices = append(defensiveSkillChoices, warCrySkillChoice)
	defensiveSkillChoices = append(defensiveSkillChoices, mantraOfEvasionSkillChoice)
	defensiveSkillChoices = append(defensiveSkillChoices, poweredArmorSkillChoice)

	return defensiveSkillChoices
}

func printSimpleOnOffHtml(sc SkillChoice, title string, w http.ResponseWriter) {
	printAgnosticHtml(sc, title, true, []string{}, []string{}, w)
}

func printDefensiveHtml(sc DefensiveSkillChoice, title string, hasSimpleOn bool, runeChoiceNames []string, w http.ResponseWriter) {
	printAgnosticHtml(sc, title, hasSimpleOn, runeChoiceNames, sc.GetDefensiveRuneSlugs(), w)
}

func printOffensiveHtml(sc OffensiveSkillChoice, title string, hasSimpleOn bool, runeChoiceNames []string, w http.ResponseWriter) {
	printAgnosticHtml(sc, title, hasSimpleOn, runeChoiceNames, sc.GetOffensiveRuneSlugs(), w)
}

func printAgnosticHtml(sc SkillChoice, title string, hasSimpleOn bool, runeChoiceNames []string, runeChoiceSlugs []string, w http.ResponseWriter) {
	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintf(w, `<td class="tableLeft">%s:</td>%s`, title, "\n")
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('defensiveForm').submit();">%s`, sc.GetUrlKey(), "\n")
	fmt.Fprintf(w, `<option value="%s" %s >Off</option>%s`, standardUrlValueOff, GetSelected(sc, standardUrlValueOff), "\n")

	if hasSimpleOn {
		fmt.Fprintf(w, `<option value="%s" %s >On</option>%s`, standardUrlValueOn, GetSelected(sc, standardUrlValueOn), "\n")
	}

	for i := 0; i < len(runeChoiceNames); i++ {
		fmt.Fprintf(w, `<option value="%s" %s >%s</option>%s`, runeChoiceSlugs[i], GetSelected(sc, runeChoiceSlugs[i]), runeChoiceNames[i], "\n")
	}

	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)
}

type BashSkillChoice struct {
	Value string
}
type BattleRageSkillChoice struct {
	Value string
}
type BerzerkerRageSkillChoice struct {
	Value string
}
type BrawlerSkillChoice struct {
	Value string
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
type FrenzySkillChoice struct {
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
type OverpowerSkillChoice struct {
	Value string
}
type PoweredArmorSkillChoice struct {
	Value string
}
type RevengeSkillChoice struct {
	Value string
}
type RuthlessSkillChoice struct {
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
type WeaponsMasterSkillChoice struct {
	Value string
}
type WrathOfTheBerzerkerSkillChoice struct {
	Value string
}

func (self *BashSkillChoice) GetValue() string {
	return self.Value
}
func (self *BattleRageSkillChoice) GetValue() string {
	return self.Value
}
func (self *BerzerkerRageSkillChoice) GetValue() string {
	return self.Value
}
func (self *BrawlerSkillChoice) GetValue() string {
	return self.Value
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
func (self *FrenzySkillChoice) GetValue() string {
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
func (self *OverpowerSkillChoice) GetValue() string {
	return self.Value
}
func (self *PoweredArmorSkillChoice) GetValue() string {
	return self.Value
}
func (self *RevengeSkillChoice) GetValue() string {
	return self.Value
}
func (self *RuthlessSkillChoice) GetValue() string {
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
func (self *WeaponsMasterSkillChoice) GetValue() string {
	return self.Value
}
func (self *WrathOfTheBerzerkerSkillChoice) GetValue() string {
	return self.Value
}

func (self *BashSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *BattleRageSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *BerzerkerRageSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *BrawlerSkillChoice) SetValue(value string) {
	self.Value = value
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
func (self *FrenzySkillChoice) SetValue(value string) {
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
func (self *OverpowerSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *PoweredArmorSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *RevengeSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *RuthlessSkillChoice) SetValue(value string) {
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
func (self *WeaponsMasterSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *WrathOfTheBerzerkerSkillChoice) SetValue(value string) {
	self.Value = value
}

func (self *BashSkillChoice) GetUrlKey() string {
	return bashUrlKey
}
func (self *BattleRageSkillChoice) GetUrlKey() string {
	return battleRageUrlKey
}
func (self *BerzerkerRageSkillChoice) GetUrlKey() string {
	return berzerkerRageUrlKey
}
func (self *BrawlerSkillChoice) GetUrlKey() string {
	return brawlerUrlKey
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
func (self *FrenzySkillChoice) GetUrlKey() string {
	return frenzyUrlKey
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
func (self *OverpowerSkillChoice) GetUrlKey() string {
	return overpowerUrlKey
}
func (self *PoweredArmorSkillChoice) GetUrlKey() string {
	return poweredArmorUrlKey
}
func (self *RevengeSkillChoice) GetUrlKey() string {
	return revengeUrlKey
}
func (self *RuthlessSkillChoice) GetUrlKey() string {
	return ruthlessUrlKey
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
func (self *WeaponsMasterSkillChoice) GetUrlKey() string {
	return weaponsMasterUrlKey
}
func (self *WrathOfTheBerzerkerSkillChoice) GetUrlKey() string {
	return wrathOfTheBerzerkerUrlKey
}

func (self *BashSkillChoice) GetSkillSlug() string {
	return bashSkillSlug
}
func (self *BattleRageSkillChoice) GetSkillSlug() string {
	return battleRageSkillSlug
}
func (self *BerzerkerRageSkillChoice) GetSkillSlug() string {
	return berzerkerRageSkillSlug
}
func (self *BrawlerSkillChoice) GetSkillSlug() string {
	return brawlerSkillSlug
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
func (self *FrenzySkillChoice) GetSkillSlug() string {
	return frenzySkillSlug
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
func (self *OverpowerSkillChoice) GetSkillSlug() string {
	return overpowerSkillSlug
}
func (self *PoweredArmorSkillChoice) GetSkillSlug() string {
	return poweredArmorSkillSlug
}
func (self *RevengeSkillChoice) GetSkillSlug() string {
	return revengeSkillSlug
}
func (self *RuthlessSkillChoice) GetSkillSlug() string {
	return ruthlessSkillSlug
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
func (self *WeaponsMasterSkillChoice) GetSkillSlug() string {
	return weaponsMasterSkillSlug
}
func (self *WrathOfTheBerzerkerSkillChoice) GetSkillSlug() string {
	return wrathOfTheBerzerkerSkillSlug
}

func (self *BashSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"bash-a"}
}
func (self *BattleRageSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"battle-rage-a", "battle-rage-e"}
}
func (self *BerzerkerRageSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *BrawlerSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *DeadlyReachSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"deadly-reach-e"}
}
func (self *DeadlyReachSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"deadly-reach-f"}
}
func (self *EnergyArmorSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"energy-armor-a"}
}
func (self *FistsOfThunderSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"fists-of-thunder-e"}
}
func (self *FrenzySkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"frenzy-a"}
}
func (self *GlassCannonSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *GuardiansPathSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *HorrifySkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"horrify-a"}
}
func (self *IgnorePainSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *JungleFortitudeSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *LeapSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"leap-z123"}
}
func (self *MantraOfEvasionSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"spam", "mantra-of-evasion-e", "hard-target-spam"}
}
func (self *MantraOfHealingSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"mantra-of-healing-e"}
}
func (self *NervesOfSteelSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *OneWithEverythingSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *OverpowerSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"overpower-a"}
}
func (self *PoweredArmorSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *RevengeSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"revenge-a"}
}
func (self *RuthlessSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *SeizeTheInitiativeSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *ToughAsNailsSkillChoice) GetDefensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *WarCrySkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"war-cry-a", "war-cry-c"}
}
func (self *WeaponsMasterSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *WrathOfTheBerzerkerSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"wrath-of-the-berzerker-a", "wrath-of-the-berzerker-c"}
}

// TODO switch all of these to switch statemetns

func (self *BashSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Punish
		derivedStats.SkillDamageBonus += 0.24
	}
}
func (self *BattleRageSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	switch {
	case self.Value == standardUrlValueOn:
		derivedStats.SkillDamageBonus += 0.15
		derivedStats.CritChance += 0.03
	case self.Value == runeSlugs[0]: // Marauder's Rage
		derivedStats.SkillDamageBonus += 0.30
		derivedStats.CritChance += 0.03
	}
}
func (self *BerzerkerRageSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.SkillDamageBonus += 0.25
	}
}
func (self *BrawlerSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.SkillDamageBonus += 0.30
	}
}
func (self *FrenzySkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Maniac //TODO split this and bash into stacks.
		derivedStats.SkillDamageBonus += 0.20
	}
}
func (self *OverpowerSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Killing Spree
		derivedStats.CritChance += 0.10
	}
}
func (self *RevengeSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Best Served Cold
		derivedStats.CritChance += 0.10
	}
}
func (self *RuthlessSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.CritChance += 0.05
		derivedStats.CritDamage += 0.50
	}
}
func (self *WeaponsMasterSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		switch t := derivedStats.BaseStats.MainWeaponType; {
		case t == util.UrlValueWeaponTypeSword || t == util.UrlValueWeaponTypeDagger:
			derivedStats.SkillDamageBonus += 0.15
		case t == util.UrlValueWeaponTypeMace || t == util.UrlValueWeaponTypeAxe:
			derivedStats.CritChance += 0.10
		case t == util.UrlValueWeaponTypePolearm || t == util.UrlValueWeaponTypeSpear:
			derivedStats.AttackSpeedBonus += 0.10
		}
	}
}
func (self *WrathOfTheBerzerkerSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	switch {
	case self.Value == standardUrlValueOn:
		derivedStats.CritChance += 0.10
		derivedStats.AttackSpeedBonus += 0.25
	case self.Value == runeSlugs[0]: // Insanity
		derivedStats.SkillDamageBonus += 1.00
	}
}

func (self *DeadlyReachSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	runeSlugs := self.GetDefensiveRuneSlugs()
	if self.Value == runeSlugs[0] {
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.5
	}
}
func (self *EnergyArmorSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	runeSlugs := self.GetDefensiveRuneSlugs()
	switch {
	case self.Value == standardUrlValueOn:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.65

	case self.Value == runeSlugs[0]:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.65
		derivedStats.ResistArcane += derivedStats.BaseStats.ResistArcane * 0.4
		derivedStats.ResistFire += derivedStats.BaseStats.ResistFire * 0.4
		derivedStats.ResistLightning += derivedStats.BaseStats.ResistLightning * 0.4
		derivedStats.ResistPoison += derivedStats.BaseStats.ResistPoison * 0.4
		derivedStats.ResistCold += derivedStats.BaseStats.ResistCold * 0.4
		derivedStats.ResistPhysical += derivedStats.BaseStats.ResistPhysical * 0.4
	}
}
func (self *FistsOfThunderSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	runeSlugs := self.GetDefensiveRuneSlugs()
	if self.Value == runeSlugs[0] {
		util.AddDodge(0.16, &derivedStats.MitigationSources)
	}
}
func (self *GlassCannonSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
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
func (self *GuardiansPathSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		util.AddDodge(0.15, &derivedStats.MitigationSources)
	}
}
func (self *HorrifySkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	runeSlugs := self.GetDefensiveRuneSlugs()
	if self.Value == runeSlugs[0] {
		derivedStats.Armor += derivedStats.BaseStats.Armor
	}
}
func (self *IgnorePainSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.MitigationSources["Ignore Pain"] = 0.65
	}
}
func (self *LeapSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	runeSlugs := self.GetDefensiveRuneSlugs()
	if self.Value == runeSlugs[0] {
		derivedStats.Armor += derivedStats.BaseStats.Armor * 3.0
	}
}
func (self *MantraOfEvasionSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	runeSlugs := self.GetDefensiveRuneSlugs()
	switch {
	case self.Value == standardUrlValueOn:
		util.AddDodge(0.15, &derivedStats.MitigationSources)

	case self.Value == runeSlugs[0]: // Spam
		util.AddDodge(0.30, &derivedStats.MitigationSources)

	case self.Value == runeSlugs[1]: // Hard Target
		util.AddDodge(0.15, &derivedStats.MitigationSources)
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.2

	case self.Value == runeSlugs[2]: // Hard Target Spam
		util.AddDodge(0.30, &derivedStats.MitigationSources)
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.2
	}
}
func (self *MantraOfHealingSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	runeSlugs := self.GetDefensiveRuneSlugs()
	if self.Value == runeSlugs[0] {
		derivedStats.ResistArcane += derivedStats.BaseStats.ResistArcane * 0.2
		derivedStats.ResistFire += derivedStats.BaseStats.ResistFire * 0.2
		derivedStats.ResistLightning += derivedStats.BaseStats.ResistLightning * 0.2
		derivedStats.ResistPoison += derivedStats.BaseStats.ResistPoison * 0.2
		derivedStats.ResistCold += derivedStats.BaseStats.ResistCold * 0.2
		derivedStats.ResistPhysical += derivedStats.BaseStats.ResistPhysical * 0.2
	}
}
func (self *NervesOfSteelSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	if self.Value == standardUrlValueOn {

		multiplier := derivedStats.Armor / derivedStats.BaseStats.Armor

		derivedStats.BaseStats.Armor += float64(derivedStats.BaseStats.Vitality)

		derivedStats.Armor = derivedStats.BaseStats.Armor * multiplier
	}
}
func (self *OneWithEverythingSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		max := util.FindMax(derivedStats.ResistArcane, derivedStats.ResistFire, derivedStats.ResistLightning, derivedStats.ResistPoison, derivedStats.ResistCold, derivedStats.ResistPhysical)
		derivedStats.ResistArcane = max
		derivedStats.ResistFire = max
		derivedStats.ResistLightning = max
		derivedStats.ResistPoison = max
		derivedStats.ResistCold = max
		derivedStats.ResistPhysical = max
	}
}
func (self *PoweredArmorSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.15
	}
}
func (self *SeizeTheInitiativeSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.Armor += float64(derivedStats.Dexterity)
	}
}
func (self *ToughAsNailsSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.25
	}
}
func (self *JungleFortitudeSkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.MitigationSources["Jungle Fortitude"] = 0.20
	}
}
func (self *WarCrySkillChoice) ModifyDefensiveDerivedStats(derivedStats *defensive.DerivedStats) {
	runeSlugs := self.GetDefensiveRuneSlugs()
	switch {
	case self.Value == standardUrlValueOn:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.2

	case self.Value == runeSlugs[0]:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.4

	case self.Value == runeSlugs[1]:
		derivedStats.Armor += derivedStats.BaseStats.Armor * 0.2
		derivedStats.ResistFire += derivedStats.BaseStats.ResistFire * 0.5
		derivedStats.ResistLightning += derivedStats.BaseStats.ResistLightning * 0.5
		derivedStats.ResistPoison += derivedStats.BaseStats.ResistPoison * 0.5
		derivedStats.ResistCold += derivedStats.BaseStats.ResistCold * 0.5
		derivedStats.ResistArcane += derivedStats.BaseStats.ResistArcane * 0.5
		derivedStats.ResistPhysical += derivedStats.BaseStats.ResistPhysical * 0.5

	}
}

func (self *BashSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Bash", false, []string{"Punish"}, w)
}
func (self *BattleRageSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Battle Rage", true, []string{"Marauder's Rage"}, w)
}
func (self *BerzerkerRageSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Berzerker Rage", w)
}
func (self *BrawlerSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Brawler", w)
}
func (self *DeadlyReachSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printDefensiveHtml(self, "Deadly Reach", false, []string{"Keen Eye"}, w)
}
func (self *EnergyArmorSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printDefensiveHtml(self, "Energy Armor", true, []string{"Prismatic"}, w)
}
func (self *FistsOfThunderSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printDefensiveHtml(self, "Fists of Thunder", false, []string{"Lightning Flash"}, w)
}
func (self *FrenzySkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Frenzy", false, []string{"Maniac"}, w)
}
func (self *GlassCannonSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "GlassCannon", w)
}
func (self *GuardiansPathSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Guardian's Path", w)
}
func (self *HorrifySkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printDefensiveHtml(self, "Horrify", false, []string{"Frightening Aspect"}, w)
}
func (self *IgnorePainSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Ignore Pain", w)
}
func (self *JungleFortitudeSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Jungle Fortitude", w)
}
func (self *LeapSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printDefensiveHtml(self, "Leap", false, []string{"Iron Impact"}, w)
}
func (self *MantraOfEvasionSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printDefensiveHtml(self, "Mantra of Evasion", true, []string{"Spam", "Hard Target", "Hard Target Spam"}, w)
}
func (self *MantraOfHealingSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printDefensiveHtml(self, "Mantra of Healing", false, []string{"Time of Need"}, w)
}
func (self *NervesOfSteelSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Nerves of Steel", w)
}
func (self *OneWithEverythingSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "One With Everything", w)
}
func (self *OverpowerSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Overpower", false, []string{"Killing Spree"}, w)
}
func (self *PoweredArmorSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Enchantress Armor", w)
}
func (self *RevengeSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Revenge", false, []string{"Best Served Cold"}, w)
}
func (self *RuthlessSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Ruthless", w)
}
func (self *SeizeTheInitiativeSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Seize the Initiative", w)
}
func (self *ToughAsNailsSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Tough as Nails", w)
}
func (self *WarCrySkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printDefensiveHtml(self, "War Cry", true, []string{"Hardened Wrath", "Impunity"}, w)
}
func (self *WeaponsMasterSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Weapons Master", w)
}
func (self *WrathOfTheBerzerkerSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "WotB", true, []string{"Insanity"}, w)
}
