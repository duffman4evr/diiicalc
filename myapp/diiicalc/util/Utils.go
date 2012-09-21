package util

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Some package private helper functions.
func createSystemBattleTag(userBattleTag string) (systemBattleTag string, err error) {

	systemBattleTag = strings.Replace(userBattleTag, "#", "-", -1)

	match, _ := regexp.MatchString(`[^0-9][^\-]*-[0-9]+`, systemBattleTag)

	if !match {
		err = errors.New("Given BattleTag is not of the form 'name-number' or 'name#number'.")
	}

	return
}

// Some publicly exposed helper functions.
func ComputeBaseLifeForHero(vitality float64, level float64) (life float64) {
	return 36 + (4 * level) + DeriveLifeFromVitality(vitality, level)
}

func DeriveLifeFromVitality(vitality float64, level float64) (life float64) {
	var lifePerVit float64

	if level < 35 {
		lifePerVit = 10.0
	} else {
		lifePerVit = level - 25.0
	}

	return vitality * lifePerVit
}

func AddDodge(dodge float64, mitigationSources *map[string]float64) {

	oldDodge := (*mitigationSources)[MitigationSourceDodge]
	addedDodge := (1 - oldDodge) * dodge

	(*mitigationSources)[MitigationSourceDodge] = oldDodge + addedDodge

}

func ComputeDodgeChanceFromDexterity(dex float64) (dodgeChance float64) {

	//     Dex     | Chance per Dex, with 1 == 100% dodge chance
	// ------------+----------------
	//    1 - 100  |      .001
	//  101 - 500  |      .00025
	//  501 - 1000 |      .00020
	// 1001 - 8000 |      .00010

	dodgeChance = 0

	if dex <= 100 {
		return dex * .001
	} else {
		dodgeChance += 0.1
		dex -= 100
	}

	if dex <= 400 {
		return dodgeChance + (dex * .00025)
	} else {
		dodgeChance += 0.1
		dex -= 400
	}

	if dex <= 500 {
		return dodgeChance + (dex * .00020)
	} else {
		dodgeChance += 0.1
		dex -= 500
	}

	dodgeChance += dex * .00010

	if dodgeChance > 0.999 {
		dodgeChance = 0.999
	}

	return dodgeChance
}

func ComputeArmorFromDr(dr float64, lvl float64) (armor float64) {
	return 50.0 * lvl * dr / (1.0 - dr)
}

func GenerateCommaLadenValue(f float64) (value string) {

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

func GetSignForValue(f float64) (sign string) {

	if f >= 0 {
		sign = "+"
	}

	return
}

func GetColorForValue(f float64) (color string) {

	if f < 0 {
		color = "#A31919"
	} else {
		color = "#008A2E"
	}

	return color
}

func FindMin(f ...float64) (min float64) {

	length := len(f)

	min = f[0]

	for i := 1; i < length; i++ {
		newVal := f[i]
		if newVal < min {
			min = newVal
		}
	}

	return min
}

func FindMax(f ...float64) (max float64) {

	length := len(f)

	max = f[0]

	for i := 1; i < length; i++ {
		newVal := f[i]
		if newVal > max {
			max = newVal
		}
	}

	return max
}
