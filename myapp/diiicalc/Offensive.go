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

	fmt.Fprintln(w, `<form id="mainForm" method="GET" autocomplete="off">`)

	// Stuff all of your URL params into the form as hidden elements.

	// TODO turn this into a for loop over offensive URL keys
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

	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyWeaponSetup, r.FormValue(util.UrlKeyWeaponSetup), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyMainWeaponType, r.FormValue(util.UrlKeyMainWeaponType), "\n")

	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyCritChance, r.FormValue(util.UrlKeyCritChance), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyCritDamageBonus, r.FormValue(util.UrlKeyCritDamageBonus), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyAttackSpeedBonus, r.FormValue(util.UrlKeyAttackSpeedBonus), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyAverageDamageBonus, r.FormValue(util.UrlKeyAverageDamageBonus), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyMainWeaponAverageDamage, r.FormValue(util.UrlKeyMainWeaponAverageDamage), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyMainWeaponAttackSpeedBase, r.FormValue(util.UrlKeyMainWeaponAttackSpeedBase), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyMainWeaponAttackSpeedBonus, r.FormValue(util.UrlKeyMainWeaponAttackSpeedBonus), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyOffWeaponAverageDamage, r.FormValue(util.UrlKeyOffWeaponAverageDamage), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyOffWeaponAttackSpeedBase, r.FormValue(util.UrlKeyOffWeaponAttackSpeedBase), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyOffWeaponAttackSpeedBonus, r.FormValue(util.UrlKeyOffWeaponAttackSpeedBonus), "\n")

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
	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td>`)
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

	// Main Summary + Skill Selection
	fmt.Fprintln(w, `<table class="fullWidth">`)
	fmt.Fprintln(w, `<tr>`)

	// Main summary.
	fmt.Fprintln(w, `<td style="width: 30%">`)
	fmt.Fprintln(w, `<div class="roundedBorder" style="width: auto;">`)
	fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">DPS: <span style="font-weight: bold;">%s</span></div>%s`, util.GenerateCommaLadenValue(metaStats.Dps), "\n")
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	// Skill selection.
	fmt.Fprintln(w, `<td style="width: 70%">`)
	fmt.Fprintln(w, `<div class="roundedBorder centerText" style="width: auto; background-color: #B2D1B2;">`)
	fmt.Fprintln(w, `<div style="vertical-align: middle;">`)
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
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)
	fmt.Fprintln(w, `</table>`)

	// Key Stats and Stat Equivalencies.
	fmt.Fprintln(w, `<table class="fullWidth">`)
	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="halfWidth">`)
	fmt.Fprintln(w, `<div class="roundedBorder" style="width: auto;">`)
	fmt.Fprintln(w, `<table class="fullWidth">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2" class="centerText" style="text-decoration: underline; font-size: 20px;">Key Stats</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">Crit Chance: </td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight" style="font-weight: bold;">%.1f%%</td>%s`, metaStats.DerivedStats.CritChance*100, "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">Crit Damage Bonus: </td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight" style="font-weight: bold;">%.0f%%</td>%s`, (metaStats.DerivedStats.CritDamageBonus)*100, "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">Attack Speed Bonus: </td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight" style="font-weight: bold;">%.0f%%</td>%s`, metaStats.DerivedStats.AttackSpeedBonus*100, "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="halfWidth">`)
	fmt.Fprintln(w, `<div class="roundedBorder" style="width: auto;">`)
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
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)
	fmt.Fprintln(w, `</table>`)

	// Print stat comparison utility.
	fmt.Fprintln(w, `<div id="statCompare" class="roundedBorder centerText" style="background-color: #B2D1B2;">`)

	fmt.Fprintln(w, `<table class="fullWidth">`)
	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="thirdWidth">`)
	fmt.Fprintln(w, `<input type="submit" id="updateButton" value="Update!" style="font-size: 24px; display: none;" />`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="thirdWidth">`)
	fmt.Fprintln(w, `<div class="centerText" style="margin-top: 10px; margin-bottom: 10px; font-size: 24px;">Compare Stats:</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="thirdWidth">`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)
	fmt.Fprintln(w, `</table>`)

	fmt.Fprintln(w, `<table class="fullWidth">`)
	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input class="smallInput" name="%s" type="text" size="10" value="%s" onkeyup="showUpdateButton();" />%s`, util.UrlKeyLeftCompareValue, r.FormValue(util.UrlKeyLeftCompareValue), "\n")
	printOffensiveComparisonSelect(w, util.UrlKeyLeftCompareType, r.FormValue(util.UrlKeyLeftCompareType), baseStats.HeroClass)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input class="smallInput" name="%s" type="text" size="10" value="%s" onkeyup="showUpdateButton();" />%s`, util.UrlKeyCenterCompareValue, r.FormValue(util.UrlKeyCenterCompareValue), "\n")
	printOffensiveComparisonSelect(w, util.UrlKeyCenterCompareType, r.FormValue(util.UrlKeyCenterCompareType), baseStats.HeroClass)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input class="smallInput" name="%s" type="text" size="10" value="%s" onkeyup="showUpdateButton();" />%s`, util.UrlKeyRightCompareValue, r.FormValue(util.UrlKeyRightCompareValue), "\n")
	printOffensiveComparisonSelect(w, util.UrlKeyRightCompareType, r.FormValue(util.UrlKeyRightCompareType), baseStats.HeroClass)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="3"><div style="height: 5px"></div></td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="thirdWidth rightBorder">`)
	fmt.Fprintln(w, `<div class="fullWidth centerText">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	fmt.Fprintln(w, `<tr">`)
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
	fmt.Fprintf(w, `<td class="tableLeft" style="color: %s;">%+.0f</td>%s`, util.GetColorForValue(rightCompareDpsChange), rightCompareDpsChange, "\n")
	fmt.Fprintln(w, `<td class="tableRight">DPS</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="3"><div style="height: 5px"></div></td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)

	fmt.Fprintln(w, `</div>`)

	// Weapon Comparison
	fmt.Fprintln(w, `<div id="weapCompare" class="roundedBorder centerText" style="background-color: #B2D1B2; margin-top: 6px;">`)
	fmt.Fprintln(w, `<table class="fullWidth">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2">`)
	fmt.Fprintln(w, `<div class="centerText" style="margin-top: 10px; margin-bottom: 10px; font-size: 24px;">Compare Weapon Upgrades:</div>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="halfWidth">`)
	printWeaponComparisonWidget(w, r, util.LeftCompareMap)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="halfWidth">`)
	printWeaponComparisonWidget(w, r, util.RightCompareMap)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<table class="fullWidth">`)
	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td style="width: 10%;"></td>`)

	fmt.Fprintln(w, `<td style="width: 30%l">`)
	fmt.Fprintln(w, `+1234 DPS`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td style="width: 20%;"><button class="skip" onclick="">Compare!</button></td>`)

	fmt.Fprintln(w, `<td style="width: 30%l">`)
	fmt.Fprintln(w, `+1234 DPS`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td style="width: 10%;"></td>`)

	fmt.Fprintln(w, `</tr>`)
	fmt.Fprintln(w, `</table>`)

	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)

	fmt.Fprintln(w, `</form>`)

	printHtmlOutro(w)
}

func redirectToOffensivePage(heroId string, dashStyleBattleTag string, realm string, w http.ResponseWriter, r *http.Request) {

	hero, heroLookupError := util.LookUpHero(dashStyleBattleTag, heroId, realm, r)

	if heroLookupError != nil {
		return
	}

	hero.Items.LookUpAll(realm, r)

	var (
		weaponSetup                = hero.DeduceWeaponSetup()
		mainWeaponType             = hero.Items.MainHand.GetWeaponType()
		mainWeaponDps  = hero.Items.MainHand.GetWeaponDps()
		mainWeaponAverageDamage    = hero.Items.MainHand.CalculateAverageWeaponDamage()
		mainWeaponAttackSpeedBase  = hero.Items.MainHand.GetWeaponBaseAttackSpeed()
		mainWeaponAttackSpeedBonus = hero.Items.MainHand.GetWeaponAttackSpeedBonus()
		mainWeaponCritDamageBonus = hero.Items.MainHand.GetCritDamageBonus()
		offWeaponType              = hero.Items.OffHand.GetWeaponType()
		offWeaponDps  = hero.Items.OffHand.GetWeaponDps()
		offWeaponAverageDamage     = hero.Items.OffHand.CalculateAverageWeaponDamage()
		offWeaponAttackSpeedBase   = hero.Items.OffHand.GetWeaponBaseAttackSpeed()
		offWeaponAttackSpeedBonus  = hero.Items.OffHand.GetWeaponAttackSpeedBonus()
		offWeaponCritDamageBonus = hero.Items.OffHand.GetCritDamageBonus()
		attackSpeedBonus           = hero.CalculateTotalAttackSpeedBonus(weaponSetup)
		averageDamageBonus         = hero.CalculateTotalAverageDamageBonus()
		totalCritChance            = hero.CalculateTotalCritChance()
		totalCritDamageBonus       = hero.CalculateTotalCritDamageBonus()
		totalIntelligence          = hero.CalculateTotalIntelligence()
		totalStrength              = hero.CalculateTotalStrength()
		totalDexterity             = hero.CalculateTotalDexterity()
	)

	mainWeaponMainStatBonus := hero.Items.MainHand.GetStrengthBonus() 
	offHandMainStatBonus := hero.Items.OffHand.GetStrengthBonus() 

	switch c := hero.Class; {
	case c == util.UrlValueHeroClassMonk || c == util.UrlValueHeroClassDemonHunter:
		mainWeaponMainStatBonus = hero.Items.MainHand.GetDexterityBonus() 
		offHandMainStatBonus = hero.Items.OffHand.GetDexterityBonus() 
	case c == util.UrlValueHeroClassWizard || c == util.UrlValueHeroClassWitchDoctor:
		mainWeaponMainStatBonus = hero.Items.MainHand.GetIntelligenceBonus() 
		offHandMainStatBonus = hero.Items.OffHand.GetIntelligenceBonus() 
	}

	// Build the actual URL.
	urlValues := url.Values{}

	urlValues.Set(util.UrlKeyBattleTagSystem, dashStyleBattleTag)
	urlValues.Set(util.UrlKeyRealm, realm)
	urlValues.Set(util.UrlKeyHeroId, strconv.FormatFloat(hero.Id, 'f', 0, 64))
	urlValues.Set(util.UrlKeyHeroes, r.FormValue(util.UrlKeyHeroes))

	urlValues.Set(util.UrlKeyHeroName, hero.Name)
	urlValues.Set(util.UrlKeyHeroClass, hero.Class)
	urlValues.Set(util.UrlKeyLevel, strconv.FormatFloat(hero.Level, 'f', 0, 64))
	urlValues.Set(util.UrlKeyStrength, strconv.FormatFloat(totalStrength, 'f', 0, 64))
	urlValues.Set(util.UrlKeyIntelligence, strconv.FormatFloat(totalIntelligence, 'f', 0, 64))
	urlValues.Set(util.UrlKeyDexterity, strconv.FormatFloat(totalDexterity, 'f', 0, 64))

	urlValues.Set(util.UrlKeyWeaponSetup, weaponSetup)
	urlValues.Set(util.UrlKeyMainWeaponType, mainWeaponType)
	urlValues.Set(util.UrlKeyCritChance, strconv.FormatFloat(totalCritChance, 'f', 3, 64))
	urlValues.Set(util.UrlKeyCritDamageBonus, strconv.FormatFloat(totalCritDamageBonus, 'f', 2, 64))
	urlValues.Set(util.UrlKeyAttackSpeedBonus, strconv.FormatFloat(attackSpeedBonus, 'f', 2, 64))
	urlValues.Set(util.UrlKeyAverageDamageBonus, strconv.FormatFloat(averageDamageBonus, 'f', 6, 64))
	urlValues.Set(util.UrlKeyMainWeaponAverageDamage, strconv.FormatFloat(mainWeaponAverageDamage, 'f', 6, 64))
	urlValues.Set(util.UrlKeyMainWeaponAttackSpeedBase, strconv.FormatFloat(mainWeaponAttackSpeedBase, 'f', 2, 64))
	urlValues.Set(util.UrlKeyMainWeaponAttackSpeedBonus, strconv.FormatFloat(mainWeaponAttackSpeedBonus, 'f', 2, 64))
	urlValues.Set(util.UrlKeyOffWeaponAverageDamage, strconv.FormatFloat(offWeaponAverageDamage, 'f', 6, 64))
	urlValues.Set(util.UrlKeyOffWeaponAttackSpeedBase, strconv.FormatFloat(offWeaponAttackSpeedBase, 'f', 2, 64))
	urlValues.Set(util.UrlKeyOffWeaponAttackSpeedBonus, strconv.FormatFloat(offWeaponAttackSpeedBonus, 'f', 2, 64))

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

	// Set up values for our weapon comparator.
	urlValues.Set(util.UrlKeyLwcWeaponSetup, weaponSetup)
	urlValues.Set(util.UrlKeyRwcWeaponSetup, weaponSetup)

	urlValues.Set(util.UrlKeyLwcMainHandWeaponType, mainWeaponType)
	urlValues.Set(util.UrlKeyRwcMainHandWeaponType, mainWeaponType)
	urlValues.Set(util.UrlKeyLwcOffHandWeaponType, offWeaponType)
	urlValues.Set(util.UrlKeyRwcOffHandWeaponType, offWeaponType)

	urlValues.Set(util.UrlKeyLwcMainHandWeaponDps, strconv.FormatFloat(mainWeaponDps - 50.0, 'f', 1, 64))
	urlValues.Set(util.UrlKeyRwcMainHandWeaponDps, strconv.FormatFloat(mainWeaponDps + 50.0, 'f', 1, 64))
	urlValues.Set(util.UrlKeyLwcOffHandWeaponDps, strconv.FormatFloat(offWeaponDps, 'f', 1, 64))
	urlValues.Set(util.UrlKeyRwcOffHandWeaponDps, strconv.FormatFloat(offWeaponDps, 'f', 1, 64))

	urlValues.Set(util.UrlKeyLwcMainHandWeaponAttackSpeedBonus, strconv.FormatFloat(mainWeaponAttackSpeedBonus * 100.0, 'f', 0, 64))
	urlValues.Set(util.UrlKeyRwcMainHandWeaponAttackSpeedBonus, strconv.FormatFloat(mainWeaponAttackSpeedBonus * 100.0, 'f', 0, 64))
	urlValues.Set(util.UrlKeyLwcOffHandWeaponAttackSpeedBonus, strconv.FormatFloat(offWeaponAttackSpeedBonus * 100.0, 'f', 0, 64))
	urlValues.Set(util.UrlKeyRwcOffHandWeaponAttackSpeedBonus, strconv.FormatFloat(offWeaponAttackSpeedBonus * 100.0, 'f', 0, 64))

	urlValues.Set(util.UrlKeyLwcMainHandWeaponCritDamageBonus, strconv.FormatFloat(mainWeaponCritDamageBonus * 100.0, 'f', 0, 64))
	urlValues.Set(util.UrlKeyRwcMainHandWeaponCritDamageBonus, strconv.FormatFloat(mainWeaponCritDamageBonus * 100.0, 'f', 0, 64))
	urlValues.Set(util.UrlKeyLwcOffHandWeaponCritDamageBonus, strconv.FormatFloat(offWeaponCritDamageBonus * 100.0, 'f', 0, 64))
	urlValues.Set(util.UrlKeyRwcOffHandWeaponCritDamageBonus, strconv.FormatFloat(offWeaponCritDamageBonus * 100.0, 'f', 0, 64))

	urlValues.Set(util.UrlKeyLwcMainHandWeaponMainStatBonus, strconv.FormatFloat(mainWeaponMainStatBonus, 'f', 0, 64))
	urlValues.Set(util.UrlKeyRwcMainHandWeaponMainStatBonus, strconv.FormatFloat(mainWeaponMainStatBonus, 'f', 0, 64))
	urlValues.Set(util.UrlKeyLwcOffHandWeaponMainStatBonus, strconv.FormatFloat(offHandMainStatBonus, 'f', 0, 64))
	urlValues.Set(util.UrlKeyRwcOffHandWeaponMainStatBonus, strconv.FormatFloat(offHandMainStatBonus, 'f', 0, 64))
	urlValues.Set(util.UrlKeyLwcOffHandMainStatBonus, strconv.FormatFloat(offHandMainStatBonus, 'f', 0, 64))
	urlValues.Set(util.UrlKeyRwcOffHandMainStatBonus, strconv.FormatFloat(offHandMainStatBonus, 'f', 0, 64))

	/*                                      
	UrlKeyLwcOffHandType                   
	UrlKeyLwcOffHandAverageDamage          
	UrlKeyLwcOffHandAttackSpeedBonus       
	UrlKeyLwcOffHandCritChanceBonus          
    
	UrlKeyRwcOffHandType                   
	UrlKeyRwcOffHandAverageDamage          
	UrlKeyRwcOffHandAttackSpeedBonus       
	UrlKeyRwcOffHandCritChanceBonus           */

	http.Redirect(w, r, "Offensive?"+urlValues.Encode(), 301)
}

func printWeaponComparisonWidget(w http.ResponseWriter, r *http.Request, lookupMap map[string]string) {

	weaponSetup := r.FormValue(lookupMap[util.MapKeyWeaponSetup])

	var (
		mhohSelected = ""
		dwSelected   = ""
		thSelected   = ""
	)

	switch {
	case weaponSetup == util.UrlValueWeaponSetupMainHandOffHand:
		mhohSelected = "selected"
	case weaponSetup == util.UrlValueWeaponSetupDualWield:
		dwSelected = "selected"
	case weaponSetup == util.UrlValueWeaponSetupTwoHander:
		thSelected = "selected"
	}

	fmt.Fprintln(w, `<div class="roundedBorder centerText" style="background-color: #FFE0B2;">`)
	fmt.Fprintln(w, `<table class="fullWidth">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2">`)
	fmt.Fprintf(w, `<select name="%s" onchange="document.getElementById('mainForm').submit();">%s`, lookupMap[util.MapKeyWeaponSetup], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueWeaponSetupMainHandOffHand, mhohSelected, util.WeaponSetupMap[util.UrlValueWeaponSetupMainHandOffHand], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueWeaponSetupDualWield, dwSelected, util.WeaponSetupMap[util.UrlValueWeaponSetupDualWield], "\n")
	if r.FormValue(util.UrlKeyHeroClass) != util.UrlValueHeroClassDemonHunter {
		fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueWeaponSetupTwoHander, thSelected, util.WeaponSetupMap[util.UrlValueWeaponSetupTwoHander], "\n")
	}
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

	if weaponSetup != util.UrlValueWeaponSetupTwoHander {

		fmt.Fprintln(w, `<tr>`)
		fmt.Fprintln(w, `<td style="text-decoration: underline;">Main Hand</td>`)
		fmt.Fprintln(w, `<td style="text-decoration: underline;">Off Hand</td>`)
		fmt.Fprintln(w, `</tr>`)

	}

	fmt.Fprintln(w, `<tr>`)

	if weaponSetup == util.UrlValueWeaponSetupTwoHander {

		fmt.Fprintln(w, `<td colspan="2">`)
		printWeaponWidget(w, r, true, true, lookupMap)
		fmt.Fprintln(w, `</td>`)

	} else {

		fmt.Fprintln(w, `<td class="halfWidth rightBorder">`)
		printWeaponWidget(w, r, false, true, lookupMap)
		fmt.Fprintln(w, `</td>`)

		fmt.Fprintln(w, `<td class="halfWidth">`)

		if weaponSetup == util.UrlValueWeaponSetupDualWield {

			printWeaponWidget(w, r, false, false, lookupMap)

			// Print the Offhand values into hidden input so that we remember them.
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandType], r.FormValue(lookupMap[util.MapKeyOffHandType]), "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandAverageDamage], r.FormValue(lookupMap[util.MapKeyOffHandAverageDamage]), "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandAttackSpeedBonus], r.FormValue(lookupMap[util.MapKeyOffHandAttackSpeedBonus]), "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandCritChanceBonus], r.FormValue(lookupMap[util.MapKeyOffHandCritChanceBonus]), "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandMainStatBonus], r.FormValue(lookupMap[util.MapKeyOffHandMainStatBonus]), "\n")

		} else {

			printOffHandWidget(w, r, lookupMap)

			// Print the weapon Offhand values into hidden input so that we remember them.
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandWeaponType], r.FormValue(lookupMap[util.MapKeyOffHandWeaponType]), "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandWeaponDps], r.FormValue(lookupMap[util.MapKeyOffHandWeaponDps]), "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandWeaponAttackSpeedBonus], r.FormValue(lookupMap[util.MapKeyOffHandWeaponAttackSpeedBonus]), "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandWeaponCritDamageBonus], r.FormValue(lookupMap[util.MapKeyOffHandWeaponCritDamageBonus]), "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, lookupMap[util.MapKeyOffHandWeaponMainStatBonus], r.FormValue(lookupMap[util.MapKeyOffHandWeaponMainStatBonus]), "\n")
			
		}

		fmt.Fprintln(w, `</td>`)

	}

	fmt.Fprintln(w, `<td>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
}

func printWeaponWidget(w http.ResponseWriter, r *http.Request, twoHand bool, mainHand bool, lookupMap map[string]string) {

	var (
		weaponTypeMapKey        = util.MapKeyMainHandWeaponType
		weaponDpsMapKey         = util.MapKeyMainHandWeaponDps
		weaponAttackSpeedMapKey = util.MapKeyMainHandWeaponAttackSpeedBonus
		weaponCritDamageMapKey  = util.MapKeyMainHandWeaponCritDamageBonus
		weaponMainStatMapKey    = util.MapKeyMainHandWeaponMainStatBonus
	)

	if !mainHand {
		weaponTypeMapKey = util.MapKeyOffHandWeaponType
		weaponDpsMapKey = util.MapKeyOffHandWeaponDps
		weaponAttackSpeedMapKey = util.MapKeyOffHandWeaponAttackSpeedBonus
		weaponCritDamageMapKey = util.MapKeyOffHandWeaponCritDamageBonus
		weaponMainStatMapKey = util.MapKeyOffHandWeaponMainStatBonus
	}

	weaponType := r.FormValue(lookupMap[weaponTypeMapKey])
	weaponTypeLookupMap := util.WearableOhWeaponMap

	if twoHand {
		weaponTypeLookupMap = util.WearableThWeaponMap
	}

	validWeaponTypes := weaponTypeLookupMap[r.FormValue(util.UrlKeyHeroClass)]

	fmt.Fprintln(w, `<table class="centerBlock">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2">`)
	fmt.Fprintf(w, `<select name="%s" onchange="showUpdateButton();">%s`, lookupMap[weaponTypeMapKey], "\n")

	for i := 0; i < len(validWeaponTypes); i++ {

		selected := ""

		if weaponType == validWeaponTypes[i] {
			selected = "selected"
		}

		fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, validWeaponTypes[i], selected, util.WeaponTypeMap[validWeaponTypes[i]], "\n")

	}

	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">DPS:</td>`)
	fmt.Fprintf(w, `<td class="tableRight"><input class="mediumInput" name="%s" value="%s" /></td>%s`, lookupMap[weaponDpsMapKey], r.FormValue(lookupMap[weaponDpsMapKey]), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Atk Speed:</td>`)
	fmt.Fprintf(w, `<td class="tableRight"><input class="smallInput" name="%s" value="%s" /></td>%s`, lookupMap[weaponAttackSpeedMapKey], r.FormValue(lookupMap[weaponAttackSpeedMapKey]), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Crit Dmg:</td>`)
	fmt.Fprintf(w, `<td class="tableRight"><input class="smallInput" name="%s" value="%s" /></td>%s`, lookupMap[weaponCritDamageMapKey], r.FormValue(lookupMap[weaponCritDamageMapKey]), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Main Stat:</td>`)
	fmt.Fprintf(w, `<td class="tableRight"><input class="smallInput" name="%s" value="%s" /></td>%s`, lookupMap[weaponMainStatMapKey], r.FormValue(lookupMap[weaponMainStatMapKey]), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
}

func printOffHandWidget(w http.ResponseWriter, r *http.Request, lookupMap map[string]string) {

	var (
		offHandTypeMapKey = util.MapKeyOffHandType
		offHandAverageDamageMapKey = util.MapKeyOffHandAverageDamage
		offHandAttackSpeedMapKey   = util.MapKeyMainHandWeaponAttackSpeedBonus
		offHandCritChanceMapKey    = util.MapKeyOffHandCritChanceBonus
		offHandMainStatMapKey      = util.MapKeyOffHandMainStatBonus
	)

	offHandType := r.FormValue(lookupMap[offHandTypeMapKey])
	validOffHandTypes := util.WearableOhMap[r.FormValue(util.UrlKeyHeroClass)]

	fmt.Fprintln(w, `<table class="centerBlock">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2">`)
	fmt.Fprintf(w, `<select name="%s" onchange="showUpdateButton();">%s`, lookupMap[offHandTypeMapKey], "\n")

	for i := 0; i < len(validOffHandTypes); i++ {

		selected := ""

		if offHandType == validOffHandTypes[i] {
			selected = "selected"
		}

		fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, validOffHandTypes[i], selected, util.OffHandTypeMap[validOffHandTypes[i]], "\n")

	}

	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Avg Damage:</td>`)
	fmt.Fprintf(w, `<td class="tableRight"><input class="smallInput" name="%s" value="%s" /></td>%s`, lookupMap[offHandAverageDamageMapKey], r.FormValue(lookupMap[offHandAverageDamageMapKey]), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Atk Speed:</td>`)
	fmt.Fprintf(w, `<td class="tableRight"><input class="smallInput" name="%s" value="%s" /></td>%s`, lookupMap[offHandAttackSpeedMapKey], r.FormValue(lookupMap[offHandAttackSpeedMapKey]), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Crit Chance:</td>`)
	fmt.Fprintf(w, `<td class="tableRight"><input class="smallInput" name="%s" value="%s" /></td>%s`, lookupMap[offHandCritChanceMapKey], r.FormValue(lookupMap[offHandCritChanceMapKey]), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Main Stat:</td>`)
	fmt.Fprintf(w, `<td class="tableRight"><input class="smallInput" name="%s" value="%s" /></td>%s`, lookupMap[offHandMainStatMapKey], r.FormValue(lookupMap[offHandMainStatMapKey]), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
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
