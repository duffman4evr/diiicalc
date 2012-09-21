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

func ParseOffensiveSkillChoices(r *http.Request) (skillChoices []OffensiveSkillChoice) {

	skillChoices = make([]OffensiveSkillChoice, 0, 10)

	heroClass := r.FormValue(util.UrlKeyHeroClass)

	if heroClass == util.UrlValueHeroClassBarbarian {

		skillChoices = append(skillChoices, new(BashSkillChoice))
		skillChoices = append(skillChoices, new(BattleRageSkillChoice))
		skillChoices = append(skillChoices, new(BerzerkerRageSkillChoice))
		skillChoices = append(skillChoices, new(BrawlerSkillChoice))
		skillChoices = append(skillChoices, new(FrenzySkillChoice))
		skillChoices = append(skillChoices, new(OverpowerSkillChoice))
		skillChoices = append(skillChoices, new(RevengeSkillChoice))
		skillChoices = append(skillChoices, new(RuthlessSkillChoice))
		skillChoices = append(skillChoices, new(WeaponsMasterSkillChoice))
		skillChoices = append(skillChoices, new(WrathOfTheBerzerkerSkillChoice))

	} else if heroClass == util.UrlValueHeroClassMonk {

		skillChoices = append(skillChoices, new(BreathOfHeavenSkillChoice))
		skillChoices = append(skillChoices, new(CombinationStrikeSkillChoice))
		skillChoices = append(skillChoices, new(CripplingWaveSkillChoice))
		skillChoices = append(skillChoices, new(DeadlyReachSkillChoice))
		skillChoices = append(skillChoices, new(ExplodingPalmSkillChoice))
		skillChoices = append(skillChoices, new(InnerSanctuarySkillChoice))
		skillChoices = append(skillChoices, new(MantraOfConvictionSkillChoice))
		skillChoices = append(skillChoices, new(MantraOfRetributionSkillChoice))
		skillChoices = append(skillChoices, new(WayOfTheHundredFistsSkillChoice))

	} else if heroClass == util.UrlValueHeroClassWizard {

	} else if heroClass == util.UrlValueHeroClassDemonHunter {

		skillChoices = append(skillChoices, new(ArcherySkillChoice))
		skillChoices = append(skillChoices, new(CaltropsSkillChoice))
		skillChoices = append(skillChoices, new(CullTheWeakSkillChoice))
		skillChoices = append(skillChoices, new(MarkedForDeathSkillChoice))
		skillChoices = append(skillChoices, new(SteadyAimSkillChoice))

	} else if heroClass == util.UrlValueHeroClassWitchDoctor {

	}

	for _, skillChoice := range skillChoices {
		InitializeOffensiveSkillChoice(skillChoice, r)
	}

	return skillChoices
}

func ParseDefensiveSkillChoices(r *http.Request) (skillChoices []DefensiveSkillChoice) {

	skillChoices = make([]DefensiveSkillChoice, 0, 10)

	heroClass := r.FormValue(util.UrlKeyHeroClass)

	if heroClass == util.UrlValueHeroClassBarbarian {

		skillChoices = append(skillChoices, new(ToughAsNailsSkillChoice))
		skillChoices = append(skillChoices, new(NervesOfSteelSkillChoice))
		skillChoices = append(skillChoices, new(LeapSkillChoice))
		skillChoices = append(skillChoices, new(IgnorePainSkillChoice))

	} else if heroClass == util.UrlValueHeroClassMonk {

		skillChoices = append(skillChoices, new(SeizeTheInitiativeSkillChoice))
		skillChoices = append(skillChoices, new(OneWithEverythingSkillChoice))
		skillChoices = append(skillChoices, new(DeadlyReachSkillChoice))
		skillChoices = append(skillChoices, new(MantraOfHealingSkillChoice))
		skillChoices = append(skillChoices, new(FistsOfThunderSkillChoice))
		skillChoices = append(skillChoices, new(GuardiansPathSkillChoice))

	} else if heroClass == util.UrlValueHeroClassWizard {

		skillChoices = append(skillChoices, new(EnergyArmorSkillChoice))
		skillChoices = append(skillChoices, new(GlassCannonSkillChoice))

	} else if heroClass == util.UrlValueHeroClassDemonHunter {

		// lol Demon Hunter...

	} else if heroClass == util.UrlValueHeroClassWitchDoctor {

		skillChoices = append(skillChoices, new(HorrifySkillChoice))
		skillChoices = append(skillChoices, new(JungleFortitudeSkillChoice))

	}

	skillChoices = append(skillChoices, new(WarCrySkillChoice))
	skillChoices = append(skillChoices, new(MantraOfEvasionSkillChoice))
	skillChoices = append(skillChoices, new(PoweredArmorSkillChoice))

	for _, skillChoice := range skillChoices {
		InitializeDefensiveSkillChoice(skillChoice, r)
	}

	return skillChoices
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

type ArcherySkillChoice struct {
	Value string
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
type BreathOfHeavenSkillChoice struct {
	Value string
}
type CaltropsSkillChoice struct {
	Value string
}
type CombinationStrikeSkillChoice struct {
	Value string
}
type CripplingWaveSkillChoice struct {
	Value string
}
type CullTheWeakSkillChoice struct {
	Value string
}
type DeadlyReachSkillChoice struct {
	Value string
}
type EnergyArmorSkillChoice struct {
	Value string
}
type ExplodingPalmSkillChoice struct {
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
type InnerSanctuarySkillChoice struct {
	Value string
}
type JungleFortitudeSkillChoice struct {
	Value string
}
type LeapSkillChoice struct {
	Value string
}
type MantraOfConvictionSkillChoice struct {
	Value string
}
type MantraOfEvasionSkillChoice struct {
	Value string
}
type MantraOfHealingSkillChoice struct {
	Value string
}
type MantraOfRetributionSkillChoice struct {
	Value string
}
type MarkedForDeathSkillChoice struct {
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
type SteadyAimSkillChoice struct {
	Value string
}
type ToughAsNailsSkillChoice struct {
	Value string
}
type WarCrySkillChoice struct {
	Value string
}
type WayOfTheHundredFistsSkillChoice struct {
	Value string
}
type WeaponsMasterSkillChoice struct {
	Value string
}
type WrathOfTheBerzerkerSkillChoice struct {
	Value string
}

func (self *ArcherySkillChoice) GetValue() string {
	return self.Value
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
func (self *BreathOfHeavenSkillChoice) GetValue() string {
	return self.Value
}
func (self *CaltropsSkillChoice) GetValue() string {
	return self.Value
}
func (self *CombinationStrikeSkillChoice) GetValue() string {
	return self.Value
}
func (self *CripplingWaveSkillChoice) GetValue() string {
	return self.Value
}
func (self *CullTheWeakSkillChoice) GetValue() string {
	return self.Value
}
func (self *DeadlyReachSkillChoice) GetValue() string {
	return self.Value
}
func (self *EnergyArmorSkillChoice) GetValue() string {
	return self.Value
}
func (self *ExplodingPalmSkillChoice) GetValue() string {
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
func (self *InnerSanctuarySkillChoice) GetValue() string {
	return self.Value
}
func (self *JungleFortitudeSkillChoice) GetValue() string {
	return self.Value
}
func (self *LeapSkillChoice) GetValue() string {
	return self.Value
}
func (self *MantraOfConvictionSkillChoice) GetValue() string {
	return self.Value
}
func (self *MantraOfEvasionSkillChoice) GetValue() string {
	return self.Value
}
func (self *MantraOfHealingSkillChoice) GetValue() string {
	return self.Value
}
func (self *MantraOfRetributionSkillChoice) GetValue() string {
	return self.Value
}
func (self *MarkedForDeathSkillChoice) GetValue() string {
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
func (self *SteadyAimSkillChoice) GetValue() string {
	return self.Value
}
func (self *ToughAsNailsSkillChoice) GetValue() string {
	return self.Value
}
func (self *WarCrySkillChoice) GetValue() string {
	return self.Value
}
func (self *WayOfTheHundredFistsSkillChoice) GetValue() string {
	return self.Value
}
func (self *WeaponsMasterSkillChoice) GetValue() string {
	return self.Value
}
func (self *WrathOfTheBerzerkerSkillChoice) GetValue() string {
	return self.Value
}

func (self *ArcherySkillChoice) SetValue(value string) {
	self.Value = value
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
func (self *BreathOfHeavenSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *CaltropsSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *CombinationStrikeSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *CripplingWaveSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *CullTheWeakSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *DeadlyReachSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *EnergyArmorSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *ExplodingPalmSkillChoice) SetValue(value string) {
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
func (self *InnerSanctuarySkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *JungleFortitudeSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *LeapSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *MantraOfConvictionSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *MantraOfEvasionSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *MantraOfHealingSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *MantraOfRetributionSkillChoice) SetValue(value string) {
	self.Value = value
}
func (self *MarkedForDeathSkillChoice) SetValue(value string) {
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
func (self *SteadyAimSkillChoice) SetValue(value string) {
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
func (self *WayOfTheHundredFistsSkillChoice) SetValue(value string) {
	self.Value = value
}

// TODO shorten these when done.
func (self *ArcherySkillChoice) GetUrlKey() string {
	return "archery"
}
func (self *BashSkillChoice) GetUrlKey() string {
	return "bash"
}
func (self *BattleRageSkillChoice) GetUrlKey() string {
	return "battleRage"
}
func (self *BerzerkerRageSkillChoice) GetUrlKey() string {
	return "berzerkerRage"
}
func (self *BrawlerSkillChoice) GetUrlKey() string {
	return "brawler"
}
func (self *BreathOfHeavenSkillChoice) GetUrlKey() string {
	return "breathOfHeaven"
}
func (self *CaltropsSkillChoice) GetUrlKey() string {
	return "caltrops"
}
func (self *CombinationStrikeSkillChoice) GetUrlKey() string {
	return "combinationStrike"
}
func (self *CripplingWaveSkillChoice) GetUrlKey() string {
	return "cripplingWave"
}
func (self *CullTheWeakSkillChoice) GetUrlKey() string {
	return "cullTheWeak"
}
func (self *DeadlyReachSkillChoice) GetUrlKey() string {
	return "deadlyReach"
}
func (self *EnergyArmorSkillChoice) GetUrlKey() string {
	return "energyArmor"
}
func (self *ExplodingPalmSkillChoice) GetUrlKey() string {
	return "explodingPalm"
}
func (self *FistsOfThunderSkillChoice) GetUrlKey() string {
	return "fistsOfThunder"
}
func (self *FrenzySkillChoice) GetUrlKey() string {
	return "frenzy"
}
func (self *GlassCannonSkillChoice) GetUrlKey() string {
	return "glassCannon"
}
func (self *GuardiansPathSkillChoice) GetUrlKey() string {
	return "guardiansPath"
}
func (self *HorrifySkillChoice) GetUrlKey() string {
	return "horrify"
}
func (self *IgnorePainSkillChoice) GetUrlKey() string {
	return "ignorePain"
}
func (self *InnerSanctuarySkillChoice) GetUrlKey() string {
	return "innerSanctuary"
}
func (self *JungleFortitudeSkillChoice) GetUrlKey() string {
	return "jungleFortitude"
}
func (self *LeapSkillChoice) GetUrlKey() string {
	return "leap"
}
func (self *MantraOfConvictionSkillChoice) GetUrlKey() string {
	return "moc"
}
func (self *MantraOfEvasionSkillChoice) GetUrlKey() string {
	return "moe"
}
func (self *MantraOfHealingSkillChoice) GetUrlKey() string {
	return "moh"
}
func (self *MantraOfRetributionSkillChoice) GetUrlKey() string {
	return "mor"
}
func (self *MarkedForDeathSkillChoice) GetUrlKey() string {
	return "mfd"
}
func (self *NervesOfSteelSkillChoice) GetUrlKey() string {
	return "nervesOfSteel"
}
func (self *OneWithEverythingSkillChoice) GetUrlKey() string {
	return "oneWithEverything"
}
func (self *OverpowerSkillChoice) GetUrlKey() string {
	return "overpower"
}
func (self *PoweredArmorSkillChoice) GetUrlKey() string {
	return "poweredArmor"
}
func (self *RevengeSkillChoice) GetUrlKey() string {
	return "revenge"
}
func (self *RuthlessSkillChoice) GetUrlKey() string {
	return "ruthless"
}
func (self *SeizeTheInitiativeSkillChoice) GetUrlKey() string {
	return "seizeTheInitiative"
}
func (self *SteadyAimSkillChoice) GetUrlKey() string {
	return "steadyAim"
}
func (self *ToughAsNailsSkillChoice) GetUrlKey() string {
	return "toughAsNails"
}
func (self *WarCrySkillChoice) GetUrlKey() string {
	return "warCry"
}
func (self *WayOfTheHundredFistsSkillChoice) GetUrlKey() string {
	return "wothf"
}
func (self *WeaponsMasterSkillChoice) GetUrlKey() string {
	return "weaponsMaster"
}
func (self *WrathOfTheBerzerkerSkillChoice) GetUrlKey() string {
	return "wotb"
}

func (self *ArcherySkillChoice) GetSkillSlug() string {
	return "archery"
}
func (self *BashSkillChoice) GetSkillSlug() string {
	return "bash"
}
func (self *BattleRageSkillChoice) GetSkillSlug() string {
	return "battle-rage"
}
func (self *BerzerkerRageSkillChoice) GetSkillSlug() string {
	return "berzerker-rage"
}
func (self *BrawlerSkillChoice) GetSkillSlug() string {
	return "brawler"
}
func (self *BreathOfHeavenSkillChoice) GetSkillSlug() string {
	return "breath-of-heaven"
}
func (self *CaltropsSkillChoice) GetSkillSlug() string {
	return "caltrops"
}
func (self *CombinationStrikeSkillChoice) GetSkillSlug() string {
	return "combination-strike"
}
func (self *CripplingWaveSkillChoice) GetSkillSlug() string {
	return "crippling-wave"
}
func (self *CullTheWeakSkillChoice) GetSkillSlug() string {
	return "cull-the-weak"
}
func (self *DeadlyReachSkillChoice) GetSkillSlug() string {
	return "deadly-reach"
}
func (self *EnergyArmorSkillChoice) GetSkillSlug() string {
	return "energy-armor"
}
func (self *ExplodingPalmSkillChoice) GetSkillSlug() string {
	return "exploding-palm"
}
func (self *FistsOfThunderSkillChoice) GetSkillSlug() string {
	return "fists-of-thunder"
}
func (self *FrenzySkillChoice) GetSkillSlug() string {
	return "frenzy"
}
func (self *GlassCannonSkillChoice) GetSkillSlug() string {
	return "glass-cannon"
}
func (self *GuardiansPathSkillChoice) GetSkillSlug() string {
	return "the-guardians-path"
}
func (self *HorrifySkillChoice) GetSkillSlug() string {
	return "horrify"
}
func (self *IgnorePainSkillChoice) GetSkillSlug() string {
	return "ignore-pain"
}
func (self *InnerSanctuarySkillChoice) GetSkillSlug() string {
	return "inner-sanctuary"
}
func (self *JungleFortitudeSkillChoice) GetSkillSlug() string {
	return "jungle-fortitude"
}
func (self *LeapSkillChoice) GetSkillSlug() string {
	return "leap"
}
func (self *MantraOfConvictionSkillChoice) GetSkillSlug() string {
	return "mantra-of-conviction"
}
func (self *MantraOfEvasionSkillChoice) GetSkillSlug() string {
	return "mantra-of-evasion"
}
func (self *MantraOfHealingSkillChoice) GetSkillSlug() string {
	return "mantra-of-healing"
}
func (self *MantraOfRetributionSkillChoice) GetSkillSlug() string {
	return "mantra-of-retribution"
}
func (self *MarkedForDeathSkillChoice) GetSkillSlug() string {
	return "marked-for-death"
}
func (self *NervesOfSteelSkillChoice) GetSkillSlug() string {
	return "nerves-of-steel"
}
func (self *OneWithEverythingSkillChoice) GetSkillSlug() string {
	return "one-with-everything"
}
func (self *OverpowerSkillChoice) GetSkillSlug() string {
	return "overpower"
}
func (self *PoweredArmorSkillChoice) GetSkillSlug() string {
	return ""
}
func (self *RevengeSkillChoice) GetSkillSlug() string {
	return "revenge"
}
func (self *RuthlessSkillChoice) GetSkillSlug() string {
	return "ruthless"
}
func (self *SeizeTheInitiativeSkillChoice) GetSkillSlug() string {
	return "seize-the-initiative"
}
func (self *SteadyAimSkillChoice) GetSkillSlug() string {
	return "steady-aim"
}
func (self *ToughAsNailsSkillChoice) GetSkillSlug() string {
	return "tough-as-nails"
}
func (self *WarCrySkillChoice) GetSkillSlug() string {
	return "war-cry"
}
func (self *WayOfTheHundredFistsSkillChoice) GetSkillSlug() string {
	return "way-of-the-hundred-fists"
}
func (self *WeaponsMasterSkillChoice) GetSkillSlug() string {
	return "weapons-master"
}
func (self *WrathOfTheBerzerkerSkillChoice) GetSkillSlug() string {
	return "wrath-of-the-berzerker"
}

func (self *ArcherySkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
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
func (self *BreathOfHeavenSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"breath-of-heaven-f"}
}
func (self *CaltropsSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"caltrops-a"}
}
func (self *CombinationStrikeSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"1-stack", "2-stack", "combination-strike-f", "4-stack", "5-stack", "6-stack"}
}
func (self *CripplingWaveSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"crippling-wave-a"}
}
func (self *CullTheWeakSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *DeadlyReachSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"deadly-reach-f"}
}
func (self *ExplodingPalmSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"exploding-palm-f"}
}
func (self *FrenzySkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"frenzy-a"}
}
func (self *InnerSanctuarySkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"inner-sanctuary-a"}
}
func (self *MantraOfConvictionSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"spam", "mantra-of-conviction-a", "Overawe Spam"}
}
func (self *MantraOfRetributionSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"mantra-of-retribution-a"}
}
func (self *MarkedForDeathSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *OverpowerSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"overpower-a"}
}
func (self *RevengeSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"revenge-a"}
}
func (self *RuthlessSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *SteadyAimSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *WayOfTheHundredFistsSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"1-stack", "2-stack", "way-of-the-hundred-fists-a"}
}
func (self *WeaponsMasterSkillChoice) GetOffensiveRuneSlugs() []string {
	return emptyRuneSlugs
}
func (self *WrathOfTheBerzerkerSkillChoice) GetOffensiveRuneSlugs() []string {
	return []string{"wrath-of-the-berzerker-a", "wrath-of-the-berzerker-c"}
}

func (self *DeadlyReachSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"deadly-reach-e"}
}
func (self *EnergyArmorSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"energy-armor-a"}
}
func (self *FistsOfThunderSkillChoice) GetDefensiveRuneSlugs() []string {
	return []string{"fists-of-thunder-e"}
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
func (self *PoweredArmorSkillChoice) GetDefensiveRuneSlugs() []string {
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

func (self *ArcherySkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		switch t := derivedStats.BaseStats.MainWeaponType; {
		case t == util.UrlValueWeaponTypeBow:
			derivedStats.SkillDamageBonus += 0.15
		case t == util.UrlValueWeaponTypeCrossbow:
			derivedStats.CritDamage += 0.50
		case t == util.UrlValueWeaponTypeHandCrossbow:
			derivedStats.CritChance += 0.10
		}
	}
}
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
func (self *BreathOfHeavenSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.SkillDamageBonus += 0.15
	}
}
func (self *CaltropsSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Bait the Trap
		derivedStats.CritChance += 0.10
	}
}
func (self *CombinationStrikeSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	switch {
	case self.Value == runeSlugs[0]:
		derivedStats.SkillDamageBonus += 0.08
	case self.Value == runeSlugs[0]:
		derivedStats.SkillDamageBonus += 0.16
	case self.Value == runeSlugs[0]:
		derivedStats.SkillDamageBonus += 0.24
	case self.Value == runeSlugs[0]:
		derivedStats.SkillDamageBonus += 0.32
	case self.Value == runeSlugs[0]:
		derivedStats.SkillDamageBonus += 0.40
	case self.Value == runeSlugs[0]:
		derivedStats.SkillDamageBonus += 0.48
	}
}
func (self *CripplingWaveSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Breaking Wave //TODO mark all runes like this.
		derivedStats.SkillDamageBonus += 0.10
	}
}
func (self *CullTheWeakSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.SkillDamageBonus += 0.15
	}
}
func (self *DeadlyReachSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] {
		derivedStats.SkillDamageBonus += 0.18
	}
}
func (self *ExplodingPalmSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] {
		derivedStats.SkillDamageBonus += 0.12
	}
}
func (self *FrenzySkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Maniac //TODO split this and bash into stacks.
		derivedStats.SkillDamageBonus += 0.20
	}
}
func (self *InnerSanctuarySkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Forbidden Palace
		derivedStats.SkillDamageBonus += 0.10
	}
}
func (self *MantraOfConvictionSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	switch {
	case self.Value == standardUrlValueOn:
		derivedStats.SkillDamageBonus += 0.12
	case self.Value == runeSlugs[0]: // Spam
		derivedStats.SkillDamageBonus += 0.24
	case self.Value == runeSlugs[1]: // Overawe
		derivedStats.SkillDamageBonus += 0.24
	case self.Value == runeSlugs[2]: // Overawe Spam
		derivedStats.SkillDamageBonus += 0.48
	}
}
func (self *MantraOfRetributionSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	if self.Value == runeSlugs[0] { // Transgression
		derivedStats.AttackSpeedBonus += 0.08
	}
}
func (self *MarkedForDeathSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.SkillDamageBonus += 0.12
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
func (self *SteadyAimSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	if self.Value == standardUrlValueOn {
		derivedStats.SkillDamageBonus += 0.20
	}
}
func (self *WayOfTheHundredFistsSkillChoice) ModifyOffensiveDerivedStats(derivedStats *offensive.DerivedStats) {
	runeSlugs := self.GetOffensiveRuneSlugs()
	switch {
	case self.Value == runeSlugs[0]: // Blazing Fists 1 Stack
		derivedStats.AttackSpeedBonus += 0.05
	case self.Value == runeSlugs[1]: // Blazing Fists 2 Stack
		derivedStats.AttackSpeedBonus += 0.10
	case self.Value == runeSlugs[2]: // Blazing Fists 3 Stack
		derivedStats.AttackSpeedBonus += 0.15
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

func (self *ArcherySkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Archery", w)
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
func (self *BreathOfHeavenSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Breath of Heaven", false, []string{"Blazing Wrath"}, w)
}
func (self *CaltropsSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Caltrops", false, []string{"Bait the Trap"}, w)
}
func (self *CombinationStrikeSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Combination Strike", false, []string{"1 Stack", "2 Stacks", "3 Stacks", "4 Stacks", "5 Stacks", "6 Stacks"}, w)
}
func (self *CripplingWaveSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Crippling Wave", false, []string{"Breaking Wave"}, w)
}
func (self *CullTheWeakSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Cull the Weak", w)
}
func (self *DeadlyReachSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Deadly Reach", false, []string{"Foresight"}, w)
}
func (self *ExplodingPalmSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Exploding Palm", false, []string{"The Flesh Is Weak"}, w)
}
func (self *FrenzySkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Frenzy", false, []string{"Maniac"}, w)
}
func (self *InnerSanctuarySkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Inner Sanctuary", false, []string{"Forbidden Palace"}, w)
}
func (self *OverpowerSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Overpower", false, []string{"Killing Spree"}, w)
}
func (self *MantraOfConvictionSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Mantra of Conviction", true, []string{"Spam", "Overawe", "Overawe Spam"}, w)
}
func (self *MantraOfRetributionSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Mantra of Retribution", false, []string{"Transgression"}, w)
}
func (self *MarkedForDeathSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Marked for Death", w)
}
func (self *RevengeSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "Revenge", false, []string{"Best Served Cold"}, w)
}
func (self *RuthlessSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Ruthless", w)
}
func (self *SteadyAimSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Steady Aim", w)
}
func (self *WayOfTheHundredFistsSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "WotHF", false, []string{"1 Stack BF", "2 Stack BF", "3 Stack BF"}, w)
}
func (self *WeaponsMasterSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Weapons Master", w)
}
func (self *WrathOfTheBerzerkerSkillChoice) PrintOffensiveHtml(w http.ResponseWriter) {
	printOffensiveHtml(self, "WotB", true, []string{"Insanity"}, w)
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
func (self *PoweredArmorSkillChoice) PrintDefensiveHtml(w http.ResponseWriter) {
	printSimpleOnOffHtml(self, "Enchantress Armor", w)
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
