package defensive

import (
	"net/http"
	"strconv"
)

const (
	// URL keys.
	urlKeyHeroClass       = "hc"
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
)

type BaseStats struct {
	HeroClass       string
	Level           float64
	Dexterity       float64
	Vitality        float64
	Armor           float64
	LifePercent     float64
	LifeOnHit       float64
	LifeRegen       float64
	ResistArcane    float64
	ResistFire      float64
	ResistLightning float64
	ResistPoison    float64
	ResistCold      float64
	ResistPhysical  float64
	BlockAmountMin  float64
	BlockAmountMax  float64
	BlockChance     float64
}

func NewBaseStats(r *http.Request) *BaseStats {

	self := new(BaseStats)

	// Parse in all the source values from the request.
	self.HeroClass = r.FormValue(urlKeyHeroClass)
	self.Level, _ = strconv.ParseFloat(r.FormValue(urlKeyLevel), 64)
	self.Dexterity, _ = strconv.ParseFloat(r.FormValue(urlKeyDexterity), 64)
	self.Vitality, _ = strconv.ParseFloat(r.FormValue(urlKeyVitality), 64)
	self.Armor, _ = strconv.ParseFloat(r.FormValue(urlKeyArmor), 64)
	self.LifePercent, _ = strconv.ParseFloat(r.FormValue(urlKeyLifePercent), 64)
	self.LifeOnHit, _ = strconv.ParseFloat(r.FormValue(urlKeyLifeOnHit), 64)
	self.LifeRegen, _ = strconv.ParseFloat(r.FormValue(urlKeyLifeRegen), 64)
	self.BlockAmountMin, _ = strconv.ParseFloat(r.FormValue(urlKeyBlockMin), 64)
	self.BlockAmountMax, _ = strconv.ParseFloat(r.FormValue(urlKeyBlockMax), 64)
	self.BlockChance, _ = strconv.ParseFloat(r.FormValue(urlKeyBlockChance), 64)
	self.ResistArcane, _ = strconv.ParseFloat(r.FormValue(urlKeyResistArcane), 64)
	self.ResistFire, _ = strconv.ParseFloat(r.FormValue(urlKeyResistFire), 64)
	self.ResistLightning, _ = strconv.ParseFloat(r.FormValue(urlKeyResistLightning), 64)
	self.ResistPoison, _ = strconv.ParseFloat(r.FormValue(urlKeyResistPoison), 64)
	self.ResistCold, _ = strconv.ParseFloat(r.FormValue(urlKeyResistCold), 64)
	self.ResistPhysical, _ = strconv.ParseFloat(r.FormValue(urlKeyResistPhysical), 64)

	return self
}
