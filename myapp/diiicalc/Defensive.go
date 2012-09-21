package diiicalc

import (
	"diiicalc/defensive"
	"diiicalc/util"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

func defensivePage(w http.ResponseWriter, r *http.Request) {

	// First, check if the user switched heroes on the drop-down.
	var (
		pageHeroId         = r.FormValue(util.UrlKeyHeroId)
		dropDownHeroId     = r.FormValue(util.UrlKeyHeroIdUser)
		dashStyleBattleTag = r.FormValue(util.UrlKeyBattleTagSystem)
		realm              = r.FormValue(util.UrlKeyRealm)
	)

	if dropDownHeroId != "" && pageHeroId != dropDownHeroId {
		// If they switched heroes, do a redirect.
		redirectToDefensivePage(dropDownHeroId, dashStyleBattleTag, realm, w, r)
		return
	}

	// Parse information out of the request.
	baseStats := defensive.NewBaseStats(r)
	skillChoices := ParseDefensiveSkillChoices(r)

	// Derive actual stats using the skills chosen by user or profile.
	derivedStats := defensive.NewDerivedStats(baseStats)

	for i := 0; i < len(skillChoices); i++ {
		skillChoices[i].ModifyDefensiveDerivedStats(derivedStats)
	}

	// Populate metadata about the stats (Effective Life, etc.).
	metaStats := defensive.NewMetaStats(derivedStats)

	var (
		// Grab the comparison pieces from the request.
		leftCompareValue, _   = strconv.ParseFloat(r.FormValue(util.UrlKeyLeftCompareValue), 64)
		centerCompareValue, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyCenterCompareValue), 64)
		rightCompareValue, _  = strconv.ParseFloat(r.FormValue(util.UrlKeyRightCompareValue), 64)

		// Do the actual comparisons.
		leftCompareChanges   = metaStats.CalculateStatChangeEffect(r.FormValue(util.UrlKeyLeftCompareType), leftCompareValue)
		centerCompareChanges = metaStats.CalculateStatChangeEffect(r.FormValue(util.UrlKeyCenterCompareType), centerCompareValue)
		rightCompareChanges  = metaStats.CalculateStatChangeEffect(r.FormValue(util.UrlKeyRightCompareType), rightCompareValue)

		effectiveLifeGainForOneMoreResist, _, _ = metaStats.ComputeStatChangesForResistChange(1.0)
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
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyDexterity, r.FormValue(util.UrlKeyDexterity), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyVitality, r.FormValue(util.UrlKeyVitality), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyLifePercent, r.FormValue(util.UrlKeyLifePercent), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyLifeOnHit, r.FormValue(util.UrlKeyLifeOnHit), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyLifeRegen, r.FormValue(util.UrlKeyLifeRegen), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyArmor, r.FormValue(util.UrlKeyArmor), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyBlockMin, r.FormValue(util.UrlKeyBlockMin), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyBlockMax, r.FormValue(util.UrlKeyBlockMax), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyBlockChance, r.FormValue(util.UrlKeyBlockChance), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyResistArcane, r.FormValue(util.UrlKeyResistArcane), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyResistFire, r.FormValue(util.UrlKeyResistFire), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyResistLightning, r.FormValue(util.UrlKeyResistLightning), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyResistPoison, r.FormValue(util.UrlKeyResistPoison), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyResistCold, r.FormValue(util.UrlKeyResistCold), "\n")
	fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyResistPhysical, r.FormValue(util.UrlKeyResistPhysical), "\n")

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
	fmt.Fprintln(w, `<div style="font-size: 24px; margin: 10px;"><span>Defensive Stat Summary for </span>`)
	printHeroSelect(w, r.FormValue(util.UrlKeyHeroes), r.FormValue(util.UrlKeyHeroId))
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td style="width: 100px;">`)
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `</tr>`)
	fmt.Fprintln(w, `</table>`)

	// Main summary.
	fmt.Fprintln(w, `<div class="roundedBorder" style="float: left; width: 560px;">`)
	fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">Effective Life: <span style="font-weight: bold;">%s</span></div>%s`, util.GenerateCommaLadenValue(metaStats.EffectiveLife), "\n")
	fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">Mitigation: <span style="font-weight: bold;">%.2f%%</span></div>%s`, metaStats.TotalMitigation*100.0, "\n")

	if metaStats.EffectiveLifeOnHit > 1 {
		fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">Effective Life on Hit: <span style="font-weight: bold;">%s</span></div>%s`, util.GenerateCommaLadenValue(metaStats.EffectiveLifeOnHit), "\n")
	}

	effectiveLifeFromBlock := metaStats.EffectiveLife - metaStats.EffectiveLifeNoShield

	if effectiveLifeFromBlock > 1.0 {
		fmt.Fprintf(w, `<div style="font-size: 18px; margin: 5px;">Effective Life From Block: <span style="font-weight: bold;">%s</span></div>%s`, util.GenerateCommaLadenValue(effectiveLifeFromBlock), "\n")
	}

	fmt.Fprintln(w, `</div>`)

	// Skill selection.
	fmt.Fprintln(w, `<div class="roundedBorder centerText" style="float: right; width: 323px; height:65px; background-color: #B2D1B2; display: table;">`)
	fmt.Fprintln(w, `<div style="display: table-cell; vertical-align: middle;">`)
	fmt.Fprintln(w, `<table class="centerBlock">`)

	for i := 0; i < len(skillChoices); i++ {
		skillChoices[i].PrintDefensiveHtml(w)
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
		fmt.Fprintf(w, `<td class="halfWidth tableRight" style="font-weight: bold;">%.2f%%</td>%s`, metaStats.MitigationSources[value]*100, "\n")
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
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%.0f" onkeyup="showUpdateButton();" />%s`, util.UrlKeyLeftCompareValue, leftCompareValue, "\n")
	printDefensiveComparisonSelect(w, util.UrlKeyLeftCompareType, r.FormValue(util.UrlKeyLeftCompareType))
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%.0f" onkeyup="showUpdateButton();" />%s`, util.UrlKeyCenterCompareValue, centerCompareValue, "\n")
	printDefensiveComparisonSelect(w, util.UrlKeyCenterCompareType, r.FormValue(util.UrlKeyCenterCompareType))
	fmt.Fprintln(w, `</td>`)

	fmt.Fprintln(w, `<td class="centerText thirdWidth">`)
	fmt.Fprintf(w, `<input name="%s" type="text" size="10" value="%.0f" onkeyup="showUpdateButton();" />%s`, util.UrlKeyRightCompareValue, rightCompareValue, "\n")
	printDefensiveComparisonSelect(w, util.UrlKeyRightCompareType, r.FormValue(util.UrlKeyRightCompareType))
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

	hero, heroLookupError := util.LookUpHero(dashStyleBattleTag, heroId, realm, r)

	if heroLookupError != nil {
		return
	}

	// Figure out the base life and life % bonus.
	baseLife := util.ComputeBaseLifeForHero(hero.Stats.Vitality, hero.Level)
	lifePercentBonus := (hero.Stats.Life / baseLife) - 1.0

	// Account for Blizzard issue where they report the wrong resistance.
	if hero.Class == util.UrlValueHeroClassWizard || hero.Class == util.UrlValueHeroClassWitchDoctor {
		hero.Stats.ResistArcane += (hero.Level - 1) * (0.3)
		hero.Stats.ResistFire += (hero.Level - 1) * (0.3)
		hero.Stats.ResistLightning += (hero.Level - 1) * (0.3)
		hero.Stats.ResistPoison += (hero.Level - 1) * (0.3)
		hero.Stats.ResistCold += (hero.Level - 1) * (0.3)
		hero.Stats.ResistPhysical += (hero.Level - 1) * (0.3)
	} else {
		hero.Stats.ResistArcane += (hero.Level - 1) * (0.1)
		hero.Stats.ResistFire += (hero.Level - 1) * (0.1)
		hero.Stats.ResistLightning += (hero.Level - 1) * (0.1)
		hero.Stats.ResistPoison += (hero.Level - 1) * (0.1)
		hero.Stats.ResistCold += (hero.Level - 1) * (0.1)
		hero.Stats.ResistPhysical += (hero.Level - 1) * (0.1)
	}

	// Account for some passive skills giving issues due to Diablo 3 API reporting buffed stats.
	armorBonus := 1.0
	resistBonus := 1.0

	for i := 0; i < len(hero.Skills.Passive); i++ {
		slug := hero.Skills.Passive[i].Skill.Slug

		if slug == new(ToughAsNailsSkillChoice).GetSkillSlug() {
			armorBonus += .25
		}

		if slug == new(GlassCannonSkillChoice).GetSkillSlug() {
			armorBonus += -0.1
			resistBonus += -0.1
		}
	}

	for i := 0; i < len(hero.Skills.Passive); i++ {
		slug := hero.Skills.Passive[i].Skill.Slug

		if slug == new(NervesOfSteelSkillChoice).GetSkillSlug() {
			hero.Stats.Armor -= hero.Stats.Vitality * armorBonus
		}

		if slug == new(SeizeTheInitiativeSkillChoice).GetSkillSlug() {
			hero.Stats.Armor -= hero.Stats.Dexterity * armorBonus
		}
	}

	hero.Stats.Armor /= armorBonus

	hero.Stats.ResistArcane /= resistBonus
	hero.Stats.ResistFire /= resistBonus
	hero.Stats.ResistLightning /= resistBonus
	hero.Stats.ResistPoison /= resistBonus
	hero.Stats.ResistCold /= resistBonus
	hero.Stats.ResistPhysical /= resistBonus

	// Build the actual URL.
	urlValues := url.Values{}

	urlValues.Set(util.UrlKeyBattleTagSystem, dashStyleBattleTag)
	urlValues.Set(util.UrlKeyRealm, realm)
	urlValues.Set(util.UrlKeyHeroId, strconv.FormatFloat(hero.Id, 'f', 0, 64))
	urlValues.Set(util.UrlKeyHeroes, r.FormValue(util.UrlKeyHeroes))

	urlValues.Set(util.UrlKeyHeroName, hero.Name)
	urlValues.Set(util.UrlKeyHeroClass, hero.Class)
	urlValues.Set(util.UrlKeyLevel, strconv.FormatFloat(hero.Level, 'f', 0, 64))
	urlValues.Set(util.UrlKeyDexterity, strconv.FormatFloat(hero.Stats.Dexterity, 'f', 0, 64))
	urlValues.Set(util.UrlKeyVitality, strconv.FormatFloat(hero.Stats.Vitality, 'f', 0, 64))
	urlValues.Set(util.UrlKeyLifePercent, strconv.FormatFloat(lifePercentBonus, 'f', 3, 64))
	urlValues.Set(util.UrlKeyLifeOnHit, strconv.FormatFloat(hero.Stats.LifeOnHit, 'f', 0, 64))
	urlValues.Set(util.UrlKeyLifeRegen, "0")
	urlValues.Set(util.UrlKeyArmor, strconv.FormatFloat(hero.Stats.Armor, 'f', 0, 64))
	urlValues.Set(util.UrlKeyBlockMin, strconv.FormatFloat(hero.Stats.BlockAmountMin, 'f', 0, 64))
	urlValues.Set(util.UrlKeyBlockMax, strconv.FormatFloat(hero.Stats.BlockAmountMax, 'f', 0, 64))
	urlValues.Set(util.UrlKeyBlockChance, strconv.FormatFloat(hero.Stats.BlockChance, 'f', 3, 64))
	urlValues.Set(util.UrlKeyResistArcane, strconv.FormatFloat(hero.Stats.ResistArcane, 'f', 0, 64))
	urlValues.Set(util.UrlKeyResistFire, strconv.FormatFloat(hero.Stats.ResistFire, 'f', 0, 64))
	urlValues.Set(util.UrlKeyResistLightning, strconv.FormatFloat(hero.Stats.ResistLightning, 'f', 0, 64))
	urlValues.Set(util.UrlKeyResistPoison, strconv.FormatFloat(hero.Stats.ResistPoison, 'f', 0, 64))
	urlValues.Set(util.UrlKeyResistCold, strconv.FormatFloat(hero.Stats.ResistCold, 'f', 0, 64))
	urlValues.Set(util.UrlKeyResistPhysical, strconv.FormatFloat(hero.Stats.ResistPhysical, 'f', 0, 64))

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
		leftCompareValue   = "250"
		centerCompareValue = "80"
		rightCompareValue  = "100"

		leftCompareType   = util.UrlValueCompareTypeArmor
		centerCompareType = util.UrlValueCompareTypeResist
		rightCompareType  = util.UrlValueCompareTypeVitality
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
	if r.FormValue(util.UrlKeyRightCompareType) != "" {
		rightCompareType = r.FormValue(util.UrlKeyRightCompareType)
	}

	urlValues.Set(util.UrlKeyLeftCompareType, leftCompareType)
	urlValues.Set(util.UrlKeyCenterCompareType, centerCompareType)
	urlValues.Set(util.UrlKeyRightCompareType, rightCompareType)

	http.Redirect(w, r, "Defensive?"+urlValues.Encode(), 301)
}

func printDefensiveComparisonSelect(w http.ResponseWriter, urlKey string, selectedCompareType string) {

	var (
		armorSelected       = ""
		resistSelected      = ""
		vitalitySelected    = ""
		percentLifeSelected = ""
		dexteritySelected   = ""
	)

	switch {
	case selectedCompareType == util.UrlValueCompareTypeArmor:
		armorSelected = "selected"
	case selectedCompareType == util.UrlValueCompareTypeResist:
		resistSelected = "selected"
	case selectedCompareType == util.UrlValueCompareTypeVitality:
		vitalitySelected = "selected"
	case selectedCompareType == util.UrlValueCompareTypePercentLife:
		percentLifeSelected = "selected"
	case selectedCompareType == util.UrlValueCompareTypeDexterity:
		dexteritySelected = "selected"
	}

	fmt.Fprintf(w, `<select name="%s" onchange="showUpdateButton();">%s`, urlKey, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypeArmor, armorSelected, util.CompareTypeMap[util.UrlValueCompareTypeArmor], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypeResist, resistSelected, util.CompareTypeMap[util.UrlValueCompareTypeResist], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypeVitality, vitalitySelected, util.CompareTypeMap[util.UrlValueCompareTypeVitality], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypePercentLife, percentLifeSelected, util.CompareTypeMap[util.UrlValueCompareTypePercentLife], "\n")
	fmt.Fprintf(w, `<option value="%s" %s>%s</option>%s`, util.UrlValueCompareTypeDexterity, dexteritySelected, util.CompareTypeMap[util.UrlValueCompareTypeDexterity], "\n")
	fmt.Fprintln(w, `</select>`)

}
