package diiicalc

import (
	"diiicalc/offensive"
	"diiicalc/util"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func offensivePage(w http.ResponseWriter, r *http.Request) {

	// First, check if the user switched heroes on the drop-down.
	var (
		pageHeroId         = r.FormValue(util.UrlKeyHeroId)
		dropDownHeroId     = r.FormValue(util.UrlKeyHeroIdUser)
		dashStyleBattleTag = r.FormValue(util.UrlKeyBattleTagSystem)
		realm              = r.FormValue(util.UrlKeyRealm)
	)

	if dropDownHeroId != "" && pageHeroId != dropDownHeroId {
		// If they switched heroes, do a redirect.
		redirectToOffensivePage(dropDownHeroId, dashStyleBattleTag, realm, w, r)
		return
	}

	// Parse information out of the request.
	baseStats := offensive.NewBaseStats(r)
	skillChoices := ParseOffensiveSkillChoices(r)

	// Derive actual stats using the skills chosen by user or profile.
	derivedStats := offensive.NewDerivedStats(baseStats)

	for i := 0; i < len(skillChoices); i++ {
		skillChoices[i].ModifyOffensiveDerivedStats(derivedStats)
	}

	// Populate metadata about the stats.
	metaStats := offensive.NewMetaStats(derivedStats)

	var (
		leftCompareDpsChange   = metaStats.CalculateDpsChange(r.FormValue(util.UrlKeyLeftCompareType), r.FormValue(util.UrlKeyLeftCompareValue))
		centerCompareDpsChange = metaStats.CalculateDpsChange(r.FormValue(util.UrlKeyCenterCompareType), r.FormValue(util.UrlKeyCenterCompareValue))
		rightCompareDpsChange  = metaStats.CalculateDpsChange(r.FormValue(util.UrlKeyRightCompareType), r.FormValue(util.UrlKeyRightCompareValue))
	)

	// Print HTML Stuff.
	printHtmlIntro(w)

	fmt.Fprintln(w, `<form id="defensiveForm" method="GET" autocomplete="off">`)

	// Stuff all of your URL params into the form as hidden elements.
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyBattleTagSystem, r.FormValue(util.UrlKeyBattleTagSystem), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyHeroId, r.FormValue(util.UrlKeyHeroId), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyHeroes, r.FormValue(util.UrlKeyHeroes), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyRealm, r.FormValue(util.UrlKeyRealm), "\n")

	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyHeroName, r.FormValue(util.UrlKeyHeroName), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyHeroClass, r.FormValue(util.UrlKeyHeroClass), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyLevel, r.FormValue(util.UrlKeyLevel), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyStrength, r.FormValue(util.UrlKeyStrength), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyIntelligence, r.FormValue(util.UrlKeyIntelligence), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyDexterity, r.FormValue(util.UrlKeyDexterity), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyVitality, r.FormValue(util.UrlKeyVitality), "\n")

	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyCritChance, r.FormValue(util.UrlKeyCritChance), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyCritDamage, r.FormValue(util.UrlKeyCritDamage), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyAttackSpeedBonus, r.FormValue(util.UrlKeyAttackSpeedBonus), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyAverageDamageBonus, r.FormValue(util.UrlKeyAverageDamageBonus), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyMainWeaponAverageDamage, r.FormValue(util.UrlKeyMainWeaponAverageDamage), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyMainWeaponAttackSpeedBase, r.FormValue(util.UrlKeyMainWeaponAttackSpeedBase), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyMainWeaponAttackSpeedBonus, r.FormValue(util.UrlKeyMainWeaponAttackSpeedBonus), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyMainWeaponType, r.FormValue(util.UrlKeyMainWeaponType), "\n")

	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeySkill1, r.FormValue(util.UrlKeySkill1), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeySkill2, r.FormValue(util.UrlKeySkill2), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeySkill3, r.FormValue(util.UrlKeySkill3), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeySkill4, r.FormValue(util.UrlKeySkill4), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeySkill5, r.FormValue(util.UrlKeySkill5), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeySkill6, r.FormValue(util.UrlKeySkill6), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyPassive1, r.FormValue(util.UrlKeyPassive1), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyPassive2, r.FormValue(util.UrlKeyPassive2), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyPassive3, r.FormValue(util.UrlKeyPassive3), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyRune1, r.FormValue(util.UrlKeyRune1), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyRune2, r.FormValue(util.UrlKeyRune2), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyRune3, r.FormValue(util.UrlKeyRune3), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyRune4, r.FormValue(util.UrlKeyRune4), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyRune5, r.FormValue(util.UrlKeyRune5), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyRune6, r.FormValue(util.UrlKeyRune6), "\n")

	// Header.
	fmt.Fprintln(w, `<table>`)
	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td style="width: 100px;">`)
	fmt.Fprintln(w, `<a href="CharacterFind" style="text-decoration: none;">&lt; Search</a>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td style="width: 700px;">`)
	fmt.Fprintln(w, `<div style="font-size: 24px; margin: 10px;"><span>Offensive Stat Summary for </span>`)
	printHeroSelect(w, r.FormValue(util.UrlKeyHeroes), r.FormValue(util.UrlKeyHeroId))
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td style="width: 100px;">`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)
	fmt.Fprintln(w, `</table>`)

	// Main summary.
	fmt.Fprintln(w, `<div class="roundedBorder" style="float: left; width: 360px;">`)
	fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">DPS: <span style="font-weight: bold;">%s</span></div>%s`, util.GenerateCommaLadenValue(metaStats.Dps), "\n")
	fmt.Fprintln(w, `</div>`)

	// Skill selection.
	fmt.Fprintln(w, `<div class="roundedBorder centerText" style="float: right; width: 523px; height:65px; background-color: #B2D1B2; display: table;">`)
	fmt.Fprintln(w, `<div style="display: table-cell; vertical-align: middle;">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)
	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td>`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	for i := 0; i < len(skillChoices); i += 2 {
		skillChoices[i].PrintOffensiveHtml(w)
	}

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td>`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	for i := 1; i < len(skillChoices); i += 2 {
		skillChoices[i].PrintOffensiveHtml(w)
	}

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)
	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</div>`)

	// Mitigation Sources and Stat Equivalencies.
	fmt.Fprintln(w, `<div class="roundedBorder" style="width: 442px; clear: both; float: left">`)
	fmt.Fprintln(w, `<table class="fullWidth">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2" class="centerText" style="text-decoration: underline; font-size: 20px;">Key Stats</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">Crit Chance: </td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight" style="font-weight: bold;">%.1f%%</td>%s`, metaStats.DerivedStats.CritChance*100, "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">Crit Damage: </td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight" style="font-weight: bold;">%.0f%%</td>%s`, (metaStats.DerivedStats.CritDamage-1)*100, "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">Attack Speed: </td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight" style="font-weight: bold;">%.0f%%</td>%s`, metaStats.DerivedStats.AttackSpeedBonus*100, "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)

	fmt.Fprintln(w, `<div class="roundedBorder" style="width: 442px; float: right;">`)
	fmt.Fprintln(w, `<table class="fullWidth">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2" class="centerText" style="text-decoration: underline; font-size: 20px;">Stat Equivalencies</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">1% Crit Chance =</td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight">%.1f%% Crit Damage</td>%s`, metaStats.ComputeCritDamageEquivalentForCritChanceChange(0.01)*100, "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">1% Crit Chance =</td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight">%.0f DPS</td>%s`, metaStats.ComputeDpsChangeForCritChanceChange(0.01), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">1% Attack Speed =</td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight">%.0f DPS</td>%s`, metaStats.ComputeDpsChangeForAttackSpeedChange(0.01), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)

	// Print stat comparison utility.
	fmt.Fprintln(w, `<div id="statCompare" style="clear: both;">`)
	fmt.Fprintln(w, `<div class="centerText" style="margin-top: 10px; margin-bottom: 10px; font-size: 24px;">Compare Stats:</div>`)

	fmt.Fprintln(w, `<table class="fullWidth" style="margin-top:10px">`)

	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%s" onkeyup="showUpdateButton();" />%s`, util.UrlKeyLeftCompareValue, r.FormValue(util.UrlKeyLeftCompareValue), "\n")
	printOffensiveComparisonSelect(w, util.UrlKeyLeftCompareType, r.FormValue(util.UrlKeyLeftCompareType), baseStats.HeroClass)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%s" onkeyup="showUpdateButton();" />%s`, util.UrlKeyCenterCompareValue, r.FormValue(util.UrlKeyCenterCompareValue), "\n")
	printOffensiveComparisonSelect(w, util.UrlKeyCenterCompareType, r.FormValue(util.UrlKeyCenterCompareType), baseStats.HeroClass)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%s" onkeyup="showUpdateButton();" />%s`, util.UrlKeyRightCompareValue, r.FormValue(util.UrlKeyRightCompareValue), "\n")
	printOffensiveComparisonSelect(w, util.UrlKeyRightCompareType, r.FormValue(util.UrlKeyRightCompareType), baseStats.HeroClass)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="thirdWidth rightBorder">`)
	fmt.Fprintln(w, `<div class="fullWidth centerText">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintf(w, `<td class="tableLeft" style="color: %s;">%+.0f</td>%s`, util.GetColorForValue(leftCompareDpsChange), leftCompareDpsChange, "\n")
	fmt.Fprintln(w, `<td class="tableRight">DPS</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="thirdWidth rightBorder">`)
	fmt.Fprintln(w, `<div class="fullWidth centerText">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintf(w, `<td class="tableLeft" style="color: %s;">%+.0f</td>%s`, util.GetColorForValue(centerCompareDpsChange), centerCompareDpsChange, "\n")
	fmt.Fprintln(w, `<td class="tableRight">DPS</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="thirdWidth">`)
	fmt.Fprintln(w, `<div class="fullWidth centerText">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintf(w, `<td class="tableLeft" style="color: %s;">%+.0f</td>%s`,  util.GetColorForValue(rightCompareDpsChange), rightCompareDpsChange, "\n")
	fmt.Fprintln(w, `<td class="tableRight">DPS</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)

	fmt.Fprintln(w, `<div class="centerText" style="height: 26px;"><input type="submit" id="updateButton" value="Update!" style="font-size: 24px; display: none;" /></div>`)
	fmt.Fprintln(w, `</div>`)

	fmt.Fprintln(w, `</form>`)

	printHtmlOutro(w)
}

func redirectToOffensivePage(heroId string, dashStyleBattleTag string, realm string, w http.ResponseWriter, r *http.Request) {

	hero, heroLookupError := util.LookUpHero(dashStyleBattleTag, heroId, realm, r)

	if heroLookupError != nil {
		return
	}

	doneChannel := make(chan int, 5)

	// Look up each item in a goroutine. 
	// Send a signal when done.
	go func() {
		util.LookUpItem(&hero.Items.MainHand, realm, r)
		doneChannel <- 1
	}()
	go func() {
		util.LookUpItem(&hero.Items.OffHand, realm, r)
		doneChannel <- 1
	}()
	go func() {
		util.LookUpItem(&hero.Items.LeftFinger, realm, r)
		doneChannel <- 1
	}()
	go func() {
		util.LookUpItem(&hero.Items.RightFinger, realm, r)
		doneChannel <- 1
	}()
	go func() {
		util.LookUpItem(&hero.Items.Neck, realm, r)
		doneChannel <- 1
	}()

	for i := 0; i < 5; i++ {
		<-doneChannel
	}

	var (
		twoHandedWeapon            = hero.Items.MainHand.IsTwoHandedWeapon()
		mainWeaponType             = hero.Items.MainHand.GetWeaponType()
		mainWeaponAverageDamage    = hero.Items.MainHand.CalculateAverageWeaponDamage()
		mainWeaponAttackSpeedBase  = hero.Items.MainHand.Attributes.AttacksPerSecondBase.Min
		mainWeaponAttackSpeedBonus = hero.Items.MainHand.Attributes.AttacksPerSecondBonus.Min
		attackSpeedBonus           = (hero.Stats.AttackSpeed / mainWeaponAttackSpeedBase) - 1 - mainWeaponAttackSpeedBonus
		averageDamageBonus         = 0.0
	)

	averageDamageBonus += hero.Items.OffHand.CalculateAverageDamageBonus()
	averageDamageBonus += hero.Items.LeftFinger.CalculateAverageDamageBonus()
	averageDamageBonus += hero.Items.RightFinger.CalculateAverageDamageBonus()
	averageDamageBonus += hero.Items.Neck.CalculateAverageDamageBonus()

	// Build the actual URL.
	urlValues := url.Values{}

	urlValues.Set(util.UrlKeyBattleTagSystem, dashStyleBattleTag)
	urlValues.Set(util.UrlKeyRealm, realm)
	urlValues.Set(util.UrlKeyHeroId, strconv.FormatFloat(hero.Id, 'f', 0, 64))
	urlValues.Set(util.UrlKeyHeroes, r.FormValue(util.UrlKeyHeroes))

	urlValues.Set(util.UrlKeyHeroName, hero.Name)
	urlValues.Set(util.UrlKeyHeroClass, hero.Class)
	urlValues.Set(util.UrlKeyLevel, strconv.FormatFloat(hero.Level, 'f', 0, 64))
	urlValues.Set(util.UrlKeyStrength, strconv.FormatFloat(hero.Stats.Strength, 'f', 0, 64))
	urlValues.Set(util.UrlKeyIntelligence, strconv.FormatFloat(hero.Stats.Intelligence, 'f', 0, 64))
	urlValues.Set(util.UrlKeyDexterity, strconv.FormatFloat(hero.Stats.Dexterity, 'f', 0, 64))
	urlValues.Set(util.UrlKeyVitality, strconv.FormatFloat(hero.Stats.Vitality, 'f', 0, 64))

	urlValues.Set(util.UrlKeyTwoHandedWeapon, strconv.FormatBool(twoHandedWeapon))
	urlValues.Set(util.UrlKeyCritChance, strconv.FormatFloat(hero.Stats.CritChance, 'f', 3, 64))
	urlValues.Set(util.UrlKeyCritDamage, strconv.FormatFloat(hero.Stats.CritDamage, 'f', 2, 64))
	urlValues.Set(util.UrlKeyAttackSpeedBonus, strconv.FormatFloat(attackSpeedBonus, 'f', 6, 64))
	urlValues.Set(util.UrlKeyAverageDamageBonus, strconv.FormatFloat(averageDamageBonus, 'f', 6, 64))
	urlValues.Set(util.UrlKeyMainWeaponAverageDamage, strconv.FormatFloat(mainWeaponAverageDamage, 'f', 6, 64))
	urlValues.Set(util.UrlKeyMainWeaponAttackSpeedBase, strconv.FormatFloat(mainWeaponAttackSpeedBase, 'f', 6, 64))
	urlValues.Set(util.UrlKeyMainWeaponAttackSpeedBonus, strconv.FormatFloat(mainWeaponAttackSpeedBonus, 'f', 6, 64))
	urlValues.Set(util.UrlKeyMainWeaponType, mainWeaponType)

	urlValues.Set(util.UrlKeySkill1, hero.Skills.Active[0].Skill.Slug)
	urlValues.Set(util.UrlKeySkill2, hero.Skills.Active[1].Skill.Slug)
	urlValues.Set(util.UrlKeySkill3, hero.Skills.Active[2].Skill.Slug)
	urlValues.Set(util.UrlKeySkill4, hero.Skills.Active[3].Skill.Slug)
	urlValues.Set(util.UrlKeySkill5, hero.Skills.Active[4].Skill.Slug)
	urlValues.Set(util.UrlKeySkill6, hero.Skills.Active[5].Skill.Slug)

	urlValues.Set(util.UrlKeyRune1, hero.Skills.Active[0].Rune.Slug)
	urlValues.Set(util.UrlKeyRune2, hero.Skills.Active[1].Rune.Slug)
	urlValues.Set(util.UrlKeyRune3, hero.Skills.Active[2].Rune.Slug)
	urlValues.Set(util.UrlKeyRune4, hero.Skills.Active[3].Rune.Slug)
	urlValues.Set(util.UrlKeyRune5, hero.Skills.Active[4].Rune.Slug)
	urlValues.Set(util.UrlKeyRune6, hero.Skills.Active[5].Rune.Slug)

	urlValues.Set(util.UrlKeyPassive1, hero.Skills.Passive[0].Skill.Slug)
	urlValues.Set(util.UrlKeyPassive2, hero.Skills.Passive[1].Skill.Slug)
	urlValues.Set(util.UrlKeyPassive3, hero.Skills.Passive[2].Skill.Slug)

	var (
		leftCompareValue   = "4.5"
		centerCompareValue = "35"
		rightCompareValue  = "100"

		leftCompareType   = util.UrlValueCompareTypeCritChance
		centerCompareType = util.UrlValueCompareTypeCritDamage
		rightCompareType  = util.UrlValueCompareTypeMainStat
	)

	if r.FormValue(util.UrlKeyLeftCompareValue) != "" {
		leftCompareValue = r.FormValue(util.UrlKeyLeftCompareValue)
	}
	if r.FormValue(util.UrlKeyCenterCompareValue) != "" {
		centerCompareValue = r.FormValue(util.UrlKeyCenterCompareValue)
	}
	if r.FormValue(util.UrlKeyRightCompareValue) != "" {
		rightCompareValue = r.FormValue(util.UrlKeyRightCompareValue)
	}

	urlValues.Set(util.UrlKeyLeftCompareValue, leftCompareValue)
	urlValues.Set(util.UrlKeyCenterCompareValue, centerCompareValue)
	urlValues.Set(util.UrlKeyRightCompareValue, rightCompareValue)

	if r.FormValue(util.UrlKeyLeftCompareType) != "" {
		leftCompareType = r.FormValue(util.UrlKeyLeftCompareType)
	}
	if r.FormValue(util.UrlKeyCenterCompareType) != "" {
		centerCompareType = r.FormValue(util.UrlKeyCenterCompareType)
	}
	if r.FormValue(util.UrlKeyCenterCompareType) != "" {
		rightCompareType = r.FormValue(util.UrlKeyRightCompareType)
	}

	urlValues.Set(util.UrlKeyLeftCompareType, leftCompareType)
	urlValues.Set(util.UrlKeyCenterCompareType, centerCompareType)
	urlValues.Set(util.UrlKeyRightCompareType, rightCompareType)

	http.Redirect(w, r, "Offensive?"+urlValues.Encode(), 301)
}

func printOffensiveComparisonSelect(w http.ResponseWriter, urlKey string, selectedCompareType string, heroClass string) {

	var (
		mainStatSelected    = ""
		attackSpeedSelected = ""
		critChanceSelected  = ""
		critDamageSelected  = ""
		userVisibleMainStat = ""
	)

	switch {
	case selectedCompareType == util.UrlValueCompareTypeMainStat:
		mainStatSelected = "selected"
	case selectedCompareType == util.UrlValueCompareTypeAttackSpeed:
		attackSpeedSelected = "selected"
	case selectedCompareType == util.UrlValueCompareTypeCritChance:
		critChanceSelected = "selected"
	case selectedCompareType == util.UrlValueCompareTypeCritDamage:
		critDamageSelected = "selected"
	}

	switch {
	case heroClass == util.UrlValueHeroClassMonk:
		fallthrough
	case heroClass == util.UrlValueHeroClassDemonHunter:
		userVisibleMainStat = "Dexterity"
	case heroClass == util.UrlValueHeroClassWizard:
		fallthrough
	case heroClass == util.UrlValueHeroClassWitchDoctor:
		userVisibleMainStat = "Intelligence"
	case heroClass == util.UrlValueHeroClassBarbarian:
		userVisibleMainStat = "Strength"
	}

	fmt.Fprintf(w, `<select name="%s" onchange="showUpdateButton();">%s`, urlKey, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypeMainStat, mainStatSelected, userVisibleMainStat, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypeAttackSpeed, attackSpeedSelected, util.CompareTypeMap[util.UrlValueCompareTypeAttackSpeed], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypeCritChance, critChanceSelected, util.CompareTypeMap[util.UrlValueCompareTypeCritChance], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypeCritDamage, critDamageSelected, util.CompareTypeMap[util.UrlValueCompareTypeCritDamage], "\n")
	fmt.Fprintln(w, `</select>`)

}
