package diiicalc

import (
	"bytes"
	"strconv"
	"strings"
)

// Some package private helper functions.
func addDodge(dodge float64, mitigationSources *map[string]float64) {

	oldDodge := (*mitigationSources)[MitigationSourceDodge]
	addedDodge := (1 - oldDodge) * dodge

	(*mitigationSources)[MitigationSourceDodge] = oldDodge + addedDodge

}

func getDodgeChanceFromDexterity(dex float64) (dodgeChance float64) {

	//     Dex     | Chance per Dex
	// ------------+----------------
	//    1 - 100  |      .001
	//  101 - 500  |      .00025
	//  501 - 1000 |      .00020
	// 1001 - 8000 |      .00010

	dodgeChance = 0

	if dex <= 100 {
		return dex * .001
	} else {
		dodgeChance += 100 * .001
		dex -= 100
	}

	if dex <= 500 {
		return dodgeChance + (dex * .00025)
	} else {
		dodgeChance += 500 * .00025
		dex -= 500
	}

	if dex <= 500 {
		return dodgeChance + (dex * .00020)
	} else {
		dodgeChance += 500 * .00020
		dex -= 500
	}

	return dodgeChance + (dex * .00010)
}

func getArmorFromDr(dr float64, lvl float64) (armor float64) {
	return 50.0 * lvl * dr / (1.0 - dr)
}

func getCommaLadenValue(f float64) (value string) {

	var buffer bytes.Buffer

	if f < 0 {
		f = -f
		buffer.WriteString("-")
	}

	stringFloat := strconv.FormatFloat(f, 'f', 0, 64)
	chars := strings.Split(stringFloat, "")

	charCount := len(chars)

	for i := 0; i < charCount; i++ {

		buffer.WriteString(chars[i])

		reverseIndex := charCount - i

		if (reverseIndex - 1) == 0 {
			continue
		}

		if (reverseIndex-1)%3 == 0 {
			buffer.WriteString(",")
		}

	}

	value = buffer.String()

	return
}

func getSignForValue(f float64) (sign string) {

	if f >= 0 {
		sign = "+"
	}

	return
}

func getColorForValue(f float64) (color string) {

	if f < 0 {
		color = "#A31919"
	} else {
		color = "#008A2E"
	}

	return
}

func findMin(f ...float64) (min float64) {

	length := len(f)

	min = f[0]

	for i := 1; i < length; i++ {
		newVal := f[i]
		if newVal < min {
			min = newVal
		}
	}

	return
}

func findMax(f ...float64) (max float64) {

	length := len(f)

	max = f[0]

	for i := 1; i < length; i++ {
		newVal := f[i]
		if newVal > max {
			max = newVal
		}
	}

	return
}
