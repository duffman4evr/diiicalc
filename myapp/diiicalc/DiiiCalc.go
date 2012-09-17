package diiicalc

import (
	"bytes"
	"diiicalc/util"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Special init funtion. Sets up URL handling.
func init() {
	http.HandleFunc("/CharacterFind", characterFindPage)
	http.HandleFunc("/Defensive", defensivePage)
	http.HandleFunc("/Offensive", offensivePage)
}

func characterFindPage(w http.ResponseWriter, r *http.Request) {

	var (
		// Check where we are by looking up whether a button was hit.
		findButtonUsed      = r.FormValue(util.UrlKeyFindButton) != ""
		defensiveButtonUsed = r.FormValue(util.UrlKeyDefensiveButton) != ""
		offensiveButtonUsed = r.FormValue(util.UrlKeyOffensiveButton) != ""

		// BattleTag lookup variables.
		battleTag          = r.FormValue(util.UrlKeyBattleTagUser)
		dashStyleBattleTag = r.FormValue(util.UrlKeyBattleTagSystem)
		realm              = r.FormValue(util.UrlKeyRealm)
		heroId             = r.FormValue(util.UrlKeyHeroIdUser)

		heroes               []util.ApiHero
		battleTagLookupError error
	)

	if findButtonUsed {

		// Take the given BattleTag and look it up in the Battle.net API.
		// Show the user this page again, but with hero select options instead.
		heroes, dashStyleBattleTag, battleTagLookupError = util.LookUpBattleTag(battleTag, realm, r)

	} else if defensiveButtonUsed {

		// Try to look up the given BattleTag and hero ID with the Battle.net API.
		// Here, we trust the values that come in and just redirect to the hero page.
		redirectToDefensivePage(heroId, dashStyleBattleTag, realm, w, r)
		return

	} else if offensiveButtonUsed {

		// Try to look up the given BattleTag and hero ID with the Battle.net API.
		// Here, we trust the values that come in and just redirect to the hero page.
		redirectToOffensivePage(heroId, dashStyleBattleTag, realm, w, r)
		return

	}

	printHtmlIntro(w)

	fmt.Fprintln(w, `<div style="font-size: 24px; margin-bottom: 30px;">Diablo III Stat Calculator</div>`)

	fmt.Fprintln(w, `<form action="CharacterFind" method="GET">`)

	fmt.Fprintln(w, `<table style="margin-left: auto; margin-right: auto;">`)

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
			fmt.Fprintln(w, `<td class="tableLeft">Hero:</td>`)
			fmt.Fprintln(w, `<td class="tableRight">`)
			fmt.Fprintf(w, `<select autofocus="autofocus" name="%s" id="%s">`, util.UrlKeyHeroIdUser, util.UrlKeyHeroIdUser)

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

			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyBattleTagSystem, dashStyleBattleTag, "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyRealm, realm, "\n")
			fmt.Fprintf(w, `<input type="hidden" name="%s" value="%s" />%s`, util.UrlKeyHeroes, buffer.String(), "\n")

			fmt.Fprintf(w, `<input style="margin-top: 30px; font-size: 24px;" name="%s" type="submit" value="View Defensive Stats" />%s`, util.UrlKeyDefensiveButton, "\n")
			fmt.Fprintf(w, `<input style="margin-top: 30px; font-size: 24px;" name="%s" type="submit" value="View Offensive Stats" />%s`, util.UrlKeyOffensiveButton, "\n")

		}
	}

	fmt.Fprintln(w, `</form>`)

	printHtmlOutro(w)
}

func printBattleTagInput(w http.ResponseWriter, battleTag string, realm string) {

	// TODO don't like this. Is there a more idiomatic way?
	usSelected := ""
	euSelected := ""
	twSelected := ""
	krSelected := ""

	switch {
	case realm == util.UrlValueRealmUs:
		usSelected = "selected"
	case realm == util.UrlValueRealmEu:
		euSelected = "selected"
	case realm == util.UrlValueRealmTw:
		twSelected = "selected"
	case realm == util.UrlValueRealmKr:
		krSelected = "selected"
	}

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">BattleTag:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<input autofocus="autofocus" name="%s" id="%s" placeholder="mytag#1234" type="text" size="25" value="%s" />%s`, util.UrlKeyBattleTagUser, util.UrlKeyBattleTagUser, battleTag, "\n")
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintln(w, `<td class="tableLeft">Region:</td>`)
	fmt.Fprintln(w, `<td class="tableRight">`)
	fmt.Fprintf(w, `<select name="%s">%s`, util.UrlKeyRealm, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>US</option>%s`, util.UrlValueRealmUs, usSelected, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>EU</option>%s`, util.UrlValueRealmEu, euSelected, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>TW</option>%s`, util.UrlValueRealmTw, twSelected, "\n")
	fmt.Fprintf(w, `<option value="%s" %s>KR</option>%s`, util.UrlValueRealmKr, krSelected, "\n")
	fmt.Fprintln(w, `</select>`)
	fmt.Fprintln(w, `</td>`)
	fmt.Fprintln(w, `</tr>`)

	fmt.Fprintln(w, `<tr>`)
	fmt.Fprintf(w, `<td colspan="2"><input type="submit" style="margin-top: 15px; font-size: 24px;" name="%s" value="Find" /></td>%s`, util.UrlKeyFindButton, "\n")
	fmt.Fprintln(w, `</tr>`)

}

func printHeroSelect(w http.ResponseWriter, heroes string, selectedHeroId string) {

	fmt.Fprintf(w, `<select onchange="document.getElementById('defensiveForm').submit();" style="font-size: 24px; width:auto; height:auto; vertical-align: middle;" name="%s" id="%s">`, util.UrlKeyHeroIdUser, util.UrlKeyHeroIdUser)

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

	// TODO use a proper style sheet and JS file instead of this crap.

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

 	.footer
 	{
		position: fixed;
		width: 100%;
		bottom: 0px;
		left: 0px;
		z-index: 1;
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
	fmt.Fprintln(w, `<div class="footer" style="margin-bottom: 5px; font-size: 12px;"><span>Concerns? Email me: <a href="mailto:duffman4evr@gmail.com">duffman4evr@gmail.com</a></span></div>`)
	fmt.Fprintln(w, `</div>`)
	fmt.Fprintln(w, `</body>`)
	fmt.Fprintln(w, `</html>`)
}
