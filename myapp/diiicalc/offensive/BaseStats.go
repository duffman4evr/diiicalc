package offensive

import (
	"diiicalc/util"
	"net/http"
	"strconv"
)

type BaseStats struct {
	HeroClass                  string
	Level                      float64
	Strength                   float64
	Dexterity                  float64
	Intelligence               float64
	CritChance                 float64
	CritDamageBonus            float64
	AverageDamageBonus         float64
	AttackSpeedBonus           float64
	WeaponSetup                string
	MainWeaponType             string
	MainWeaponAverageDamage    float64
	MainWeaponAttackSpeedBase  float64
	MainWeaponAttackSpeedBonus float64
	OffWeaponAverageDamage     float64
	OffWeaponAttackSpeedBase   float64
	OffWeaponAttackSpeedBonus  float64
}

func NewBaseStats(r *http.Request) *BaseStats {

	self := new(BaseStats)

	// Parse in all the source values from the request.
	self.HeroClass = r.FormValue(util.UrlKeyHeroClass)
	self.Level, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyLevel), 64)
	self.Strength, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyStrength), 64)
	self.Dexterity, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyDexterity), 64)
	self.Intelligence, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyIntelligence), 64)
	self.CritChance, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyCritChance), 64)
	self.CritDamageBonus, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyCritDamageBonus), 64)
	self.AttackSpeedBonus, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyAttackSpeedBonus), 64)
	self.AverageDamageBonus, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyAverageDamageBonus), 64)
	self.WeaponSetup = r.FormValue(util.UrlKeyWeaponSetup)
	self.MainWeaponType = r.FormValue(util.UrlKeyMainWeaponType)
	self.MainWeaponAverageDamage, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyMainWeaponAverageDamage), 64)
	self.MainWeaponAttackSpeedBase, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyMainWeaponAttackSpeedBase), 64)
	self.MainWeaponAttackSpeedBonus, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyMainWeaponAttackSpeedBonus), 64)
	self.OffWeaponAverageDamage, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyOffWeaponAverageDamage), 64)
	self.OffWeaponAttackSpeedBase, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyOffWeaponAttackSpeedBase), 64)
	self.OffWeaponAttackSpeedBonus, _ = strconv.ParseFloat(r.FormValue(util.UrlKeyOffWeaponAttackSpeedBonus), 64)

	return self
}
