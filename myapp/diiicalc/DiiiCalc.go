package diiicalc

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

const (
	// URL keys.
	urlKeyBattleTagUser   = "btu"
	urlKeyBattleTagSystem = "bts"
	urlKeyHeroIdUser      = "hiu"
	urlKeyFindButton      = "fi"
	urlKeyDefensiveButton = "vd"
	urlKeyRealm           = "re"
	urlKeyHeroes          = "he"

	urlKeyLeftCompareType   = "lct"
	urlKeyCenterCompareType = "cct"
	urlKeyRightCompareType  = "rct"

	urlKeyLeftCompareValue   = "lcv"
	urlKeyCenterCompareValue = "ccv"
	urlKeyRightCompareValue  = "rcv"

	urlKeyHeroId          = "hi"
	urlKeyHeroName        = "hn"
	urlKeyHeroClass       = "hc"
	urlKeyNumberOfWeapons = "nw"
	urlKeyLevel           = "lv"
	urlKeyDexterity       = "de"
	urlKeyVitality        = "vi"
	urlKeyArmor           = "ar"
	urlKeyLifePercent     = "lp"
	urlKeyLifeOnHit       = "lh"
	urlKeyLifeRegen       = "lr"
	urlKeyBlockMin        = "bi"
	urlKeyBlockMax        = "ba"
	urlKeyBlockChance     = "bc"
	urlKeyResistArcane    = "ra"
	urlKeyResistFire      = "rf"
	urlKeyResistLightning = "rl"
	urlKeyResistPoison    = "rp"
	urlKeyResistCold      = "rc"
	urlKeyResistPhysical  = "ry"

	urlKeySkill1 = "s1"
	urlKeySkill2 = "s2"
	urlKeySkill3 = "s3"
	urlKeySkill4 = "s4"
	urlKeySkill5 = "s5"
	urlKeySkill6 = "s6"

	urlKeyRune1 = "r1"
	urlKeyRune2 = "r2"
	urlKeyRune3 = "r3"
	urlKeyRune4 = "r4"
	urlKeyRune5 = "r5"
	urlKeyRune6 = "r6"

	urlKeyPassive1 = "p1"
	urlKeyPassive2 = "p2"
	urlKeyPassive3 = "p3"

	// URL values.
	urlValueCompareTypeVitality    = "vit"
	urlValueCompareTypeResist      = "res"
	urlValueCompareTypeArmor       = "arm"
	urlValueCompareTypePercentLife = "pli"
	urlValueCompareTypeDexterity   = "dex"

	urlValueHeroClassBarbarian   = "barbarian"
	urlValueHeroClassMonk        = "monk"
	urlValueHeroClassWitchDoctor = "witch-doctor"
	urlValueHeroClassWizard      = "wizard"
	urlValueHeroClassDemonHunter = "demon-hunter"

	urlValueRealmUs = "us"
	urlValueRealmEu = "eu"
	urlValueRealmTw = "tw"
	urlValueRealmKr = "kr"
)

var (
	urlKeysActiveSkills  = []string{urlKeySkill1, urlKeySkill2, urlKeySkill3, urlKeySkill4, urlKeySkill5, urlKeySkill6}
	urlKeysActiveRunes   = []string{urlKeyRune1, urlKeyRune2, urlKeyRune3, urlKeyRune4, urlKeyRune5, urlKeyRune6}
	urlKeysPassiveSkills = []string{urlKeyPassive1, urlKeyPassive2, urlKeyPassive3}

	heroClassMap = make(map[string]string)
	compareTypeMap = make(map[string]string)
)

// Special init funtion.
func init() {

	heroClassMap[urlValueHeroClassBarbarian] = "Barbarian"
	heroClassMap[urlValueHeroClassMonk] = "Monk"
	heroClassMap[urlValueHeroClassWitchDoctor] = "Witch Doctor"
	heroClassMap[urlValueHeroClassWizard] = "Wizard"
	heroClassMap[urlValueHeroClassDemonHunter] = "Demon Hunter"

	compareTypeMap[urlValueCompareTypeVitality] = "Vitality"
	compareTypeMap[urlValueCompareTypeResist] = "Resist"
	compareTypeMap[urlValueCompareTypeArmor] = "Armor"
	compareTypeMap[urlValueCompareTypePercentLife] = `% Life`
	compareTypeMap[urlValueCompareTypeDexterity] = "Dexterity"

	http.HandleFunc("/CharacterFind", characterFind)
	http.HandleFunc("/Defensive", defensive)

}

func characterFind(w http.ResponseWriter, r *http.Request) {

	// Turn off the back button cache.
	w.Header().Set("Cache-Control", "no-cache, max-age=0, must-revalidate, no-store")

	var (
		// Check where we are by looking up whether a button was hit.
		findButtonUsed      = r.FormValue(urlKeyFindButton) != ""
		defensiveButtonUsed = r.FormValue(urlKeyDefensiveButton) != ""

		// BattleTag lookup variables.
		battleTag          = r.FormValue(urlKeyBattleTagUser)
		dashStyleBattleTag = r.FormValue(urlKeyBattleTagSystem)
		realm              = r.FormValue(urlKeyRealm)
		heroId             = r.FormValue(urlKeyHeroIdUser)

		heroes               []ApiHero
		battleTagLookupError error
	)

	if findButtonUsed {

		// Take the given BattleTag and look it up in the Battle.net API.
		// Show the user this page again, but with hero select options instead.
		heroes, dashStyleBattleTag, battleTagLookupError = lookUpBattleTag(battleTag, realm, r)

	} else if defensiveButtonUsed {

		// Try to look up the given BattleTag and hero ID with the Battle.net API.
		// Here, we trust the values that come in and just redirect to the hero page.
		redirectToDefensivePage(heroId, dashStyleBattleTag, realm, w, r)
		return

	} else {

		// No buttons were used. The user may have come in with a URL that tries to use
		// HTTP GET. We don't really want any of that, so just show a clean page in this case.

	}

	printHtmlIntro(w)

	fmt.Fprintln(w, `<div style="font-size: 24px; margin-bottom: 30px;">Diablo III Defensive Stat Calculator</div>`)

	fmt.Fprintln(w, `<form action="CharacterFind" method="GET">`)

	fmt.Fprintln(w, `<table class="fullWidth">`)

	if !findButtonUsed {

		printBattleTagInput(w, battleTag, realm)

	} else {

		if battleTagLookupError != nil {

			printBattleTagInput(w, battleTag, realm)

			fmt.Fprintln(w, `<tr>`)
			fmt.Fprintf(w, `<td colspan="2" class="centerText" style="color: #990000;">%s</td>%s`, battleTagLookupError.Error(), "\n")
			fmt.Fprintln(w, `</tr>`)
			fmt.Fprintln(w, `</table>`)

		} else {

			var buffer bytes.Buffer

			fmt.Fprintln(w, `<tr>`)
			fmt.Fprintln(w, `<td class="halfWidth tableLeft">Hero:</td>`)
			fmt.Fprintln(w, `<td class="halfWidth tableRight">`)
			fmt.Fprintf(w, `<select autofocus="autofocus" name="%s" id="%s">`, urlKeyHeroIdUser, urlKeyHeroIdUser)

			for i := 0; i < len(heroes); i++ {

				heroSelectionString := heroes[i].GenerateSelectionString()

				if i != 0 {
					buffer.WriteString("|")
				}

				buffer.WriteString(strconv.FormatFloat(heroes[i].Id, 'f', 0, 64))
				buffer.WriteString(",")
				buffer.WriteString(heroSelectionString)

				fmt.Fprintf(w, `<option value="%.0f">%s</option>%s`, heroes[i].Id, heroSelectionString, "\n")
			}

			fmt.Fprintln(w, `</select>`)
			fmt.Fprintln(w, `</td>`)
			fmt.Fprintln(w, `</tr>`)

			fmt.Fprintln(w, `</table>`)

			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyBattleTagSystem, dashStyleBattleTag, "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyRealm, realm, "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyHeroes, buffer.String(), "\n")

			fmt.Fprintf(w, `<input style="margin-top: 30px; font-size: 24px;" name="%s" type="submit" value="View Defensive Stats" />%s`, urlKeyDefensiveButton, "\n")

		}
	}

	fmt.Fprintln(w, `</form>`)

	printHtmlOutro(w)
}

func defensive(w http.ResponseWriter, r *http.Request) {

	// Turn off the back button cache.
	w.Header().Set("Cache-Control", "no-cache, max-age=0, must-revalidate, no-store")

	// First, check if the user switched heroes ont he drop-down.
	var (
		pageHeroId         = r.FormValue(urlKeyHeroId)
		dropDownHeroId     = r.FormValue(urlKeyHeroIdUser)
		dashStyleBattleTag = r.FormValue(urlKeyBattleTagSystem)
		realm              = r.FormValue(urlKeyRealm)
	)

	if dropDownHeroId != "" && pageHeroId != dropDownHeroId {

		redirectToDefensivePage(dropDownHeroId, dashStyleBattleTag, realm, w, r)
		return

	}

	// If they didn't switch heroes, continue as normal.
	var (
		// Parse information out of the request.
		baseStats    = NewBaseStats(r)
		skillChoices = ParseSkillChoices(r)

		// Derive actual stats using the skills chosen by user or profile.
		derivedStats = NewDerivedStats(baseStats, skillChoices)

		// Populate metadata about the stats (Effective Life, etc.).
		metaStats = NewMetaStats(derivedStats)

		// Grab the comparison pieces from the request.
		leftCompareValue, _   = strconv.ParseFloat(r.FormValue(urlKeyLeftCompareValue), 64)
		centerCompareValue, _ = strconv.ParseFloat(r.FormValue(urlKeyCenterCompareValue), 64)
		rightCompareValue, _  = strconv.ParseFloat(r.FormValue(urlKeyRightCompareValue), 64)

		// Do the actual comparisons.
		leftCompareChanges   = metaStats.CalculateStatChangeEffect(r.FormValue(urlKeyLeftCompareType), leftCompareValue)
		centerCompareChanges = metaStats.CalculateStatChangeEffect(r.FormValue(urlKeyCenterCompareType), centerCompareValue)
		rightCompareChanges  = metaStats.CalculateStatChangeEffect(r.FormValue(urlKeyRightCompareType), rightCompareValue)

		effectiveLifeGainForOneMoreResist, _, _ = metaStats.ComputeStatChangesForResistChange(1.0)
	)

	// Print HTML Stuff.
	printHtmlIntro(w)

	fmt.Fprintln(w, `<form id="defensiveForm" method="GET" autocomplete="off">`)

	// Stuff all of your URL params into the form as hidden elements.
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyBattleTagSystem, r.FormValue(urlKeyBattleTagSystem), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyHeroId, r.FormValue(urlKeyHeroId), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyHeroes, r.FormValue(urlKeyHeroes), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyRealm, r.FormValue(urlKeyRealm), "\n")

	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyHeroName, r.FormValue(urlKeyHeroName), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyHeroClass, r.FormValue(urlKeyHeroClass), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyNumberOfWeapons, r.FormValue(urlKeyNumberOfWeapons), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyLevel, r.FormValue(urlKeyLevel), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyDexterity, r.FormValue(urlKeyDexterity), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyVitality, r.FormValue(urlKeyVitality), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyLifePercent, r.FormValue(urlKeyLifePercent), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyLifeOnHit, r.FormValue(urlKeyLifeOnHit), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyLifeRegen, r.FormValue(urlKeyLifeRegen), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyArmor, r.FormValue(urlKeyArmor), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyBlockMin, r.FormValue(urlKeyBlockMin), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyBlockMax, r.FormValue(urlKeyBlockMax), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyBlockChance, r.FormValue(urlKeyBlockChance), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyResistArcane, r.FormValue(urlKeyResistArcane), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyResistFire, r.FormValue(urlKeyResistFire), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyResistLightning, r.FormValue(urlKeyResistLightning), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyResistPoison, r.FormValue(urlKeyResistPoison), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyResistCold, r.FormValue(urlKeyResistCold), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyResistPhysical, r.FormValue(urlKeyResistPhysical), "\n")

	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeySkill1, r.FormValue(urlKeySkill1), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeySkill2, r.FormValue(urlKeySkill2), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeySkill3, r.FormValue(urlKeySkill3), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeySkill4, r.FormValue(urlKeySkill4), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeySkill5, r.FormValue(urlKeySkill5), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeySkill6, r.FormValue(urlKeySkill6), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyPassive1, r.FormValue(urlKeyPassive1), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyPassive2, r.FormValue(urlKeyPassive2), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyPassive3, r.FormValue(urlKeyPassive3), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyRune1, r.FormValue(urlKeyRune1), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyRune2, r.FormValue(urlKeyRune2), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyRune3, r.FormValue(urlKeyRune3), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyRune4, r.FormValue(urlKeyRune4), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyRune5, r.FormValue(urlKeyRune5), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, urlKeyRune6, r.FormValue(urlKeyRune6), "\n")

	// Header.
	fmt.Fprintln(w, `<div style="font-size: 24px; margin: 10px;"><span>Defensive Stat Summary for </span>`)
	printHeroSelect(w, r.FormValue(urlKeyHeroes), r.FormValue(urlKeyHeroId))
	fmt.Fprintln(w, `</div>`)

	// Main summary.
	fmt.Fprintln(w, `<div class="roundedBorder" style="float: left; width: 560px;">`)
	fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">Effective Life: <span style="font-weight: bold;">%s</span></div>%s`, getCommaLadenValue(metaStats.EffectiveLife), "\n")
	fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">Mitigation: <span style="font-weight: bold;">%.2f `, metaStats.TotalMitigation*100.0)
	fmt.Fprintln(w, `%</span></div>`)

	effectiveLifeFromBlock := metaStats.EffectiveLife - metaStats.EffectiveLifeNoShield

	if effectiveLifeFromBlock > 1.0 {
		fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">Effective Life From Block: <span style="font-weight: bold;">%s</span></div>%s`, getCommaLadenValue(effectiveLifeFromBlock), "\n")
	}

	fmt.Fprintln(w, `</div>`)

	// Skill selection.
	fmt.Fprintln(w, `<div class="roundedBorder centerText" style="float: right; width: 323px; height:65px; background-color: #B2D1B2; display: table;">`)
	fmt.Fprintln(w, `<div style="display: table-cell; vertical-align: middle;">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	for i := 0; i < len(skillChoices); i++ {
		skillChoices[i].PrintHtml(w)
	}

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</div>`)

	// Mitigation Sources and Stat Equivalencies.
	fmt.Fprintln(w, `<div class="roundedBorder" style="width: 442px; clear: both; float: left">`)
	fmt.Fprintln(w, `<table class="fullWidth">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2" class="centerText" style="text-decoration: underline; font-size: 20px;">Mitigation Sources</td>`)
	fmt.Fprintln(w, `</tr>`)

	alphabeticalMitigationSources := []string{}

	for key := range metaStats.MitigationSources {
		alphabeticalMitigationSources = append(alphabeticalMitigationSources, key)
	}

	sort.Strings(alphabeticalMitigationSources)

	for _, value := range alphabeticalMitigationSources {
		fmt.Fprintln(w, `<tr>`)
		fmt.Fprintf(w, `<td class="halfWidth tableLeft">%s:</td>`, value)
		fmt.Fprintf(w, `<td class="halfWidth tableRight" style="font-weight: bold;">%.2f `, metaStats.MitigationSources[value]*100)
		fmt.Fprintln(w, `%</td>`)
		fmt.Fprintln(w, `</tr>`)
	}

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)

	fmt.Fprintln(w, `<div class="roundedBorder" style="width: 442px; float: right;">`)
	fmt.Fprintln(w, `<table class="fullWidth">`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td colspan="2" class="centerText" style="text-decoration: underline; font-size: 20px;">Stat Equivalencies</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">1 Vitality =</td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight">%.0f Effective Life</td>%s`, metaStats.ComputeEffectiveLifeChangeForVitChange(1.0), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">1 All Resist =</td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight">%.0f Armor</td>%s`, metaStats.ComputeArmorEquivalentForResistChange(1.0), "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">1 All Resist =</td>`)
	fmt.Fprintf(w, `<td class="halfWidth tableRight">%.0f Effective Life</td>%s`, effectiveLifeGainForOneMoreResist, "\n")
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)

	// Print stat comparison utility.
	fmt.Fprintln(w, `<div id="statCompare" style="clear: both;">`)
	fmt.Fprintln(w, `<div class="centerText" style="margin-top: 10px; margin-bottom: 10px; font-size: 24px;">Compare Stats:</div>`)

	fmt.Fprintln(w, `<table class="fullWidth" style="margin-top:10px">`)

	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%.0f" onkeyup="showUpdateButton();" />%s`, urlKeyLeftCompareValue, leftCompareValue, "\n")
	printComparisonSelect(w, urlKeyLeftCompareType, r.FormValue(urlKeyLeftCompareType))
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%.0f" onkeyup="showUpdateButton();" />%s`, urlKeyCenterCompareValue, centerCompareValue, "\n")
	printComparisonSelect(w, urlKeyCenterCompareType, r.FormValue(urlKeyCenterCompareType))
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%.0f" onkeyup="showUpdateButton();" />%s`, urlKeyRightCompareValue, rightCompareValue, "\n")
	printComparisonSelect(w, urlKeyRightCompareType, r.FormValue(urlKeyRightCompareType))
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)

	fmt.Fprintln(w, `<td class="thirdWidth rightBorder">`)
	fmt.Fprintln(w, `<div class="fullWidth centerText">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	for i := 0; i < len(leftCompareChanges); i++ {

		fmt.Fprintln(w, `<tr>`)
		fmt.Fprintf(w, `<td class="tableLeft" style="color: %s;">%s</td>%s`, leftCompareChanges[i].Color, leftCompareChanges[i].Value, "\n")
		fmt.Fprintf(w, `<td class="tableRight">%s</td>%s`, leftCompareChanges[i].Name, "\n")
		fmt.Fprintln(w, `</tr>`)

	}

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="thirdWidth rightBorder">`)
	fmt.Fprintln(w, `<div class="fullWidth centerText">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	for i := 0; i < len(centerCompareChanges); i++ {

		fmt.Fprintln(w, `<tr>`)
		fmt.Fprintf(w, `<td class="tableLeft" style="color: %s;">%s</td>%s`, centerCompareChanges[i].Color, centerCompareChanges[i].Value, "\n")
		fmt.Fprintf(w, `<td class="tableRight">%s</td>%s`, centerCompareChanges[i].Name, "\n")
		fmt.Fprintln(w, `</tr>`)

	}

	fmt.Fprintln(w, `</table>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="thirdWidth">`)
	fmt.Fprintln(w, `<div class="fullWidth centerText">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	for i := 0; i < len(rightCompareChanges); i++ {

		fmt.Fprintln(w, `<tr>`)
		fmt.Fprintf(w, `<td class="tableLeft" style="color: %s;">%s</td>%s`, rightCompareChanges[i].Color, rightCompareChanges[i].Value, "\n")
		fmt.Fprintf(w, `<td class="tableRight">%s</td>%s`, rightCompareChanges[i].Name, "\n")
		fmt.Fprintln(w, `</tr>`)

	}

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

func redirectToDefensivePage(heroId string, dashStyleBattleTag string, realm string, w http.ResponseWriter, r *http.Request) {

	hero, heroLookupError := lookUpHero(dashStyleBattleTag, heroId, realm, r)

	if heroLookupError != nil {
		return
	}

	var numberOfWeapons int64 = 1

	if hero.Items.MainHand.Id != hero.Items.OffHand.Id {
		numberOfWeapons = 2
	}

	baseLife := getLifeFromVitality(hero.Stats.Vitality, hero.Level)
	lifePercentBonus := (hero.Stats.Life / baseLife) - 1

	urlValues := url.Values{}

	urlValues.Set(urlKeyBattleTagSystem, dashStyleBattleTag)
	urlValues.Set(urlKeyRealm, realm)
	urlValues.Set(urlKeyHeroId, strconv.FormatFloat(hero.Id, 'f', 0, 64))
	urlValues.Set(urlKeyHeroes, r.FormValue(urlKeyHeroes))

	urlValues.Set(urlKeyHeroName, hero.Name)
	urlValues.Set(urlKeyHeroClass, hero.Class)
	urlValues.Set(urlKeyNumberOfWeapons, strconv.FormatInt(numberOfWeapons, 10))
	urlValues.Set(urlKeyLevel, strconv.FormatFloat(hero.Level, 'f', 0, 64))
	urlValues.Set(urlKeyDexterity, strconv.FormatFloat(hero.Stats.Dexterity, 'f', 0, 64))
	urlValues.Set(urlKeyVitality, strconv.FormatFloat(hero.Stats.Vitality, 'f', 0, 64))
	urlValues.Set(urlKeyLifePercent, strconv.FormatFloat(lifePercentBonus, 'f', 0, 64))
	urlValues.Set(urlKeyLifeOnHit, strconv.FormatFloat(hero.Stats.LifeOnHit, 'f', 0, 64))
	urlValues.Set(urlKeyLifeRegen, "0")
	urlValues.Set(urlKeyArmor, strconv.FormatFloat(hero.Stats.Armor, 'f', 0, 64))
	urlValues.Set(urlKeyBlockMin, strconv.FormatFloat(hero.Stats.BlockAmountMin, 'f', 0, 64))
	urlValues.Set(urlKeyBlockMax, strconv.FormatFloat(hero.Stats.BlockAmountMax, 'f', 0, 64))
	urlValues.Set(urlKeyBlockChance, strconv.FormatFloat(hero.Stats.BlockChance, 'f', 3, 64))
	urlValues.Set(urlKeyResistArcane, strconv.FormatFloat(hero.Stats.ResistArcane, 'f', 0, 64))
	urlValues.Set(urlKeyResistFire, strconv.FormatFloat(hero.Stats.ResistFire, 'f', 0, 64))
	urlValues.Set(urlKeyResistLightning, strconv.FormatFloat(hero.Stats.ResistLightning, 'f', 0, 64))
	urlValues.Set(urlKeyResistPoison, strconv.FormatFloat(hero.Stats.ResistPoison, 'f', 0, 64))
	urlValues.Set(urlKeyResistCold, strconv.FormatFloat(hero.Stats.ResistCold, 'f', 0, 64))
	urlValues.Set(urlKeyResistPhysical, strconv.FormatFloat(hero.Stats.ResistPhysical, 'f', 0, 64))

	urlValues.Set(urlKeySkill1, hero.Skills.Active[0].Skill.Slug)
	urlValues.Set(urlKeySkill2, hero.Skills.Active[1].Skill.Slug)
	urlValues.Set(urlKeySkill3, hero.Skills.Active[2].Skill.Slug)
	urlValues.Set(urlKeySkill4, hero.Skills.Active[3].Skill.Slug)
	urlValues.Set(urlKeySkill5, hero.Skills.Active[4].Skill.Slug)
	urlValues.Set(urlKeySkill6, hero.Skills.Active[5].Skill.Slug)

	urlValues.Set(urlKeyRune1, hero.Skills.Active[0].Rune.Slug)
	urlValues.Set(urlKeyRune2, hero.Skills.Active[1].Rune.Slug)
	urlValues.Set(urlKeyRune3, hero.Skills.Active[2].Rune.Slug)
	urlValues.Set(urlKeyRune4, hero.Skills.Active[3].Rune.Slug)
	urlValues.Set(urlKeyRune5, hero.Skills.Active[4].Rune.Slug)
	urlValues.Set(urlKeyRune6, hero.Skills.Active[5].Rune.Slug)

	urlValues.Set(urlKeyPassive1, hero.Skills.Passive[0].Skill.Slug)
	urlValues.Set(urlKeyPassive2, hero.Skills.Passive[1].Skill.Slug)
	urlValues.Set(urlKeyPassive3, hero.Skills.Passive[2].Skill.Slug)

	var (
		leftCompareValue   = "250"
		centerCompareValue = "80"
		rightCompareValue  = "100"

		leftCompareType    = urlValueCompareTypeArmor
		centerCompareType = urlValueCompareTypeResist
		rightCompareType  = urlValueCompareTypeVitality
	)

	if r.FormValue(urlKeyLeftCompareValue) != "" {
		leftCompareValue = r.FormValue(urlKeyLeftCompareValue)
	}
	if r.FormValue(urlKeyCenterCompareValue) != "" {
		centerCompareValue = r.FormValue(urlKeyCenterCompareValue)
	}
	if r.FormValue(urlKeyRightCompareValue) != "" {
		rightCompareValue = r.FormValue(urlKeyRightCompareValue)
	}

	urlValues.Set(urlKeyLeftCompareValue, leftCompareValue)
	urlValues.Set(urlKeyCenterCompareValue, centerCompareValue)
	urlValues.Set(urlKeyRightCompareValue, rightCompareValue)

	if r.FormValue(urlKeyLeftCompareType) != "" {
		leftCompareType = r.FormValue(urlKeyLeftCompareType)
	}
	if r.FormValue(urlKeyCenterCompareType) != "" {
		centerCompareType = r.FormValue(urlKeyCenterCompareType)
	}
	if r.FormValue(urlKeyRightCompareType) != "" {
		rightCompareType = r.FormValue(urlKeyRightCompareType)
	}

	urlValues.Set(urlKeyLeftCompareType, leftCompareType)
	urlValues.Set(urlKeyCenterCompareType, centerCompareType)
	urlValues.Set(urlKeyRightCompareType, rightCompareType)

	http.Redirect(w, r, "Defensive?"+urlValues.Encode(), 301)
}

func printComparisonSelect(w http.ResponseWriter, urlKey string, selectedCompareType string) {

	var (
		armorSelected       = ""
		resistSelected      = ""
		vitalitySelected    = ""
		percentLifeSelected = ""
		dexteritySelected   = ""
	)

	switch {
	case selectedCompareType == urlValueCompareTypeArmor:
		armorSelected = "selected"
	case selectedCompareType == urlValueCompareTypeResist:
		resistSelected = "selected"
	case selectedCompareType == urlValueCompareTypeVitality:
		vitalitySelected = "selected"
	case selectedCompareType == urlValueCompareTypePercentLife:
		percentLifeSelected = "selected"
	case selectedCompareType == urlValueCompareTypeDexterity:
		dexteritySelected = "selected"
	}

	fmt.Fprintf(w, `<select name="%s" onchange="showUpdateButton();">%s`, urlKey, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, urlValueCompareTypeArmor, armorSelected, compareTypeMap[urlValueCompareTypeArmor], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, urlValueCompareTypeResist, resistSelected, compareTypeMap[urlValueCompareTypeResist], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, urlValueCompareTypeVitality, vitalitySelected, compareTypeMap[urlValueCompareTypeVitality], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, urlValueCompareTypePercentLife, percentLifeSelected, compareTypeMap[urlValueCompareTypePercentLife], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, urlValueCompareTypeDexterity, dexteritySelected, compareTypeMap[urlValueCompareTypeDexterity], "\n")
	fmt.Fprintln(w, `</select>`)

}

func printBattleTagInput(w http.ResponseWriter, battleTag string, realm string) {

	usSelected := ""
	euSelected := ""
	twSelected := ""
	krSelected := ""

	switch {
	case realm == urlValueRealmUs:
		usSelected = "selected"
	case realm == urlValueRealmEu:
		euSelected = "selected"
	case realm == urlValueRealmTw:
		twSelected = "selected"
	case realm == urlValueRealmKr:
		krSelected = "selected"
	}

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">BattleTag:</td>`)
	fmt.Fprintln(w, `<td class="halfWidth tableRight">`)
	fmt.Fprintf(w, `<input autofocus="autofocus" name="%s" id="%s" placeholder="mytag#1234" type="text" size="25" value="%s" />%s`, urlKeyBattleTagUser, urlKeyBattleTagUser, battleTag, "\n")
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="halfWidth tableLeft">Region:</td>`)
	fmt.Fprintln(w, `<td class="halfWidth tableRight">`)
	fmt.Fprintf(w, `<select name="%s">%s`, urlKeyRealm, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>US</option>%s`, urlValueRealmUs, usSelected, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>EU</option>%s`, urlValueRealmEu, euSelected, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>TW</option>%s`, urlValueRealmTw, twSelected, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>KR</option>%s`, urlValueRealmKr, krSelected, "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintf(w, `<td colspan="2"><input type="submit" style="margin-top: 15px;" name="%s" value="Find" /></td>%s`, urlKeyFindButton, "\n")
	fmt.Fprintln(w, `</tr>`)

}

func printHeroSelect(w http.ResponseWriter, heroes string, selectedHeroId string) {

	fmt.Fprintf(w, `<select onchange="document.getElementById('defensiveForm').submit();" style="font-size: 24px; width:auto; height:auto; vertical-align: middle;" name="%s" id="%s">`, urlKeyHeroIdUser, urlKeyHeroIdUser)

	heroTokens := strings.Split(heroes, "|")

	for i := 0; i < len(heroTokens); i++ {

		keyValuePair := strings.Split(heroTokens[i], ",")

		selected := ""

		if keyValuePair[0] == selectedHeroId {
			selected = "selected"
		}

		fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, keyValuePair[0], selected, keyValuePair[1], "\n")
	}

	fmt.Fprintln(w, `</select>`)

}

func printHtmlIntro(w http.ResponseWriter) {

	fmt.Fprintln(w, `<!DOCTYPE HTML>`)
	fmt.Fprintln(w, `<html>`)

	fmt.Fprintln(w, `<head>`)
	fmt.Fprintln(w, `<style type="text/css">
	.centerBlock
	{
		margin: 0 auto;
	}

	.centerText
	{
		text-align: center;
	}

	.fullWidth
	{
		width: 100%;
	}

	.halfWidth
	{
		width: 50%;
	}

	.thirdWidth
	{
		width: 33.33%;
	}

	.twoThirdsWidth
	{
		width: 66.66%;
	}

	.tableLeft
	{
		text-align: right;
	}

	.tableRight
	{
		text-align: left;
	}

	table {
		border-spacing: 0px;
	}

	.rightBorder
	{
		border-right-style: solid;
		border-width: 2px;
		border-color: #555555;
	}

	.roundedBorder
	{
		border-style: solid;
		border-width: 1px;
		border-radius: 7px;
		margin: 3px;
		border-color: #555555;
		background-color: #DDEDFF;
		box-shadow:rgba(0,0,0,0.5) 0px 0px 24px;
	}
</style>`)

	fmt.Fprintln(w, `<script type="text/javascript">
	  
      function showUpdateButton()
      {
          document.getElementById("updateButton").setAttribute("style", "display: inline;")
      }

	</script>`)
	fmt.Fprintln(w, `</head>`)

	fmt.Fprintln(w, `<body style=" color: #2D2D2D; background-color: #CACACA; font-family: 'MS Sans Serif', Geneva, sans-serif;">`)

	fmt.Fprintln(w, `<div class="centerBlock centerText" style="width: 900px;">`)
}

func printHtmlOutro(w http.ResponseWriter) {
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</body>`)
	fmt.Fprintln(w, `</html>`)
}
