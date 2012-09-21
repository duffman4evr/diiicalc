package util

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ApiProfileLookupResponse struct {
	ErrorCode   string    `json:"code"`
	ErrorReason string    `json:"reason"`
	Heroes      []ApiHero `json:"heroes"`
}

type ApiHero struct {
	Name        string    `json:"name"`
	Class       string    `json:"class"`
	Id          float64   `json:"id"`
	Level       float64   `json:"level"`
	Hardcore    bool      `json:"hardcore"`
	Gender      float64   `json:"gender"`
	LastUpdated float64   `json:"last-updated"`
	Dead        bool      `json:"dead"`
	Stats       ApiStats  `json:"stats"`
	Skills      ApiSkills `json:"skills"`
	Items       ApiItems  `json:"skills"`
}

type ApiStats struct {
	Life              float64 `json:"life"`
	Dps               float64 `json:"damage"`
	AttackSpeed       float64 `json:"attackSpeed"`
	Armor             float64 `json:"armor"`
	Strength          float64 `json:"strength"`
	Dexterity         float64 `json:"dexterity"`
	Vitality          float64 `json:"vitality"`
	Intelligence      float64 `json:"intelligence"`
	ResistPhysical    float64 `json:"physicalResist"`
	ResistFire        float64 `json:"fireResist"`
	ResistCold        float64 `json:"coldResist"`
	ResistLightning   float64 `json:"lightningResist"`
	ResistPoison      float64 `json:"poisonResist"`
	ResistArcane      float64 `json:"arcaneResist"`
	CritDamage        float64 `json:"critDamage"`
	DamageIncrease    float64 `json:"damageIncrease"`
	CritChance        float64 `json:"critChance"`
	DamageReduction   float64 `json:"damageReduction"`
	BlockChance       float64 `json:"blockChance"`
	ThornsDamage      float64 `json:"thorns"`
	LifeSteal         float64 `json:"lifeSteal"`
	LifePerKill       float64 `json:"lifePerKill"`
	GoldFind          float64 `json:"goldFind"`
	MagicFind         float64 `json:"magicFind"`
	BlockAmountMin    float64 `json:"blockAmountMin"`
	BlockAmountMax    float64 `json:"blockAmountMax"`
	LifeOnHit         float64 `json:"lifeOnHit"`
	PrimaryResource   float64 `json:"primaryResource"`
	SecondaryResource float64 `json:"secondaryResource"`
}

type ApiSkills struct {
	Active  []ApiActiveSkill  `json:"active"`
	Passive []ApiPassiveSkill `json:"passive"`
}

type ApiActiveSkill struct {
	Skill ApiSkill `json:"skill"`
	Rune  ApiRune  `json:"rune"`
}

type ApiPassiveSkill struct {
	Skill ApiSkill `json:"skill"`
}

type ApiSkill struct {
	Slug              string  `json:"slug"`
	Name              string  `json:"name"`
	Icon              string  `json:"icon"`
	Level             float64 `json:"level"`
	CategorySlug      string  `json:"categorySlug"`
	TooltipUrl        string  `json:"tooltipUrl"`
	Description       string  `json:"description"`
	SimpleDescription string  `json:"simpleDescription"`
	SkillCalcId       string  `json:"skillCalcId"`
}

type ApiRune struct {
	Slug              string  `json:"slug"`
	Type              string  `json:"type"`
	Name              string  `json:"name"`
	Level             float64 `json:"level"`
	Description       string  `json:"description"`
	SimpleDescription string  `json:"simpleDescription"`
	TooltipParams     string  `json:"tooltipParams"`
	SkillCalcId       string  `json:"skillCalcId"`
	Order             float64 `json:"order"`
}

type ApiItems struct {
	Head        ApiItem `json:"head"`
	Torso       ApiItem `json:"torso"`
	Feet        ApiItem `json:"feet"`
	Hands       ApiItem `json:"hands"`
	Shoulders   ApiItem `json:"shoulders"`
	Legs        ApiItem `json:"legs"`
	Bracers     ApiItem `json:"bracers"`
	MainHand    ApiItem `json:"mainHand"`
	OffHand     ApiItem `json:"offHand"`
	Waist       ApiItem `json:"waist"`
	RightFinger ApiItem `json:"rightFinger"`
	LeftFinger  ApiItem `json:"leftFinger"`
	Neck        ApiItem `json:"neck"`
}

type ApiItem struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	Icon             string            `json:"icon"`
	DisplayColor     string            `json:"displayColor"`
	TooltipParams    string            `json:"tooltipParams"`
	RequiredLevel    float64           `json:"requiredLevel"`
	ItemLevel        float64           `json:"itemLevel"`
	BonusAffixes     float64           `json:"bonusAffixes"`
	TypeName         string            `json:"typeName"`
	Type             ApiItemType       `json:"type"`
	Dps              ApiMinMax         `json:"dps"`
	AttacksPerSecond ApiMinMax         `json:"attacksPerSecond"`
	Attributes       ApiItemAttributes `json:"attributesRaw"`
	Gems             []ApiGem          `json:"gems"`

	// These two are the base damage of the item plus
	// any bonuses that are to the PHYSICAL min or max value. 
	// E.G:
	//   + 22 Minumum Damage
	// That value is then multiplied by any + DMG% 
	// modifiers on the item.
	// E.G:
	//   + 50% Damage
	// The result of that calculation shows up in these 
	// two values in the API from Blizzard. If we want
	// the true base damage, we can find them in the
	// 'attributesRaw' json key.
	MinDamage ApiMinMax `json:"minDamage"`
	MaxDamage ApiMinMax `json:"maxDamage"`
}

type ApiMinMax struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type ApiGem struct {
	Item                    ApiItem           `json:"item"`
	Attributes              ApiItemAttributes `json:"attributesRaw"`
	HumanReadableAttributes []string          `json:"attributes"`
}

type ApiItemType struct {
	Id        string `json:"id"`
	TwoHanded bool   `json:"twoHanded"`
}

type ApiItemAttributes struct {
	CritDamagePercent    ApiMinMax `json:"Crit_Damage_Percent"`
	Sockets              ApiMinMax `json:"Sockets"`
	Crossbow             ApiMinMax `json:"Crossbow"`
	AttacksPerSecondBase ApiMinMax `json:"Attacks_Per_Second_Item"`

	// These values describe damage affixes that are usually present on
	// non weapons. Particularly: rings, amulets, offhands.
	DamageMinPhysical  ApiMinMax `json:"Damage_Min#Physical"`
	DamageMinArcane    ApiMinMax `json:"Damage_Min#Arcane"`
	DamageMinCold      ApiMinMax `json:"Damage_Min#Cold"`
	DamageMinFire      ApiMinMax `json:"Damage_Min#Fire"`
	DamageMinHoly      ApiMinMax `json:"Damage_Min#Holy"`
	DamageMinLightning ApiMinMax `json:"Damage_Min#Lightning"`
	DamageMinPoison    ApiMinMax `json:"Damage_Min#Poison"`

	DamageDeltaPhysical  ApiMinMax `json:"Damage_Delta#Physical"`
	DamageDeltaArcane    ApiMinMax `json:"Damage_Delta#Arcane"`
	DamageDeltaCold      ApiMinMax `json:"Damage_Delta#Cold"`
	DamageDeltaFire      ApiMinMax `json:"Damage_Delta#Fire"`
	DamageDeltaHoly      ApiMinMax `json:"Damage_Delta#Holy"`
	DamageDeltaLightning ApiMinMax `json:"Damage_Delta#Lightning"`
	DamageDeltaPoison    ApiMinMax `json:"Damage_Delta#Poison"`

	// -----
	// Following are properties for weapons specifically.
	// -----

	// This describes the attacks per second bonus on an
	// item. E.G. it is 0.11 if there is an 11% attack
	// speed bonus.
	AttacksPerSecondBonus ApiMinMax `json:"Attacks_Per_Second_Item_Percent"`

	// These values describe the minimums and deltas of 
	// both physical and elemental damage on this weapon.
	// The 'Damage_Weapon_Min#Physical' is the *true* 
	// base minimum damage of this weapon. That value is
	// then added to any bonuses found of the same type
	// and then further multiplied by any percentage
	// bonuses found for that same type of damage.
	WeaponDamageMinPhysical  ApiMinMax `json:"Damage_Weapon_Min#Physical"`
	WeaponDamageMinArcane    ApiMinMax `json:"Damage_Weapon_Min#Arcane"`
	WeaponDamageMinCold      ApiMinMax `json:"Damage_Weapon_Min#Cold"`
	WeaponDamageMinFire      ApiMinMax `json:"Damage_Weapon_Min#Fire"`
	WeaponDamageMinHoly      ApiMinMax `json:"Damage_Weapon_Min#Holy"`
	WeaponDamageMinLightning ApiMinMax `json:"Damage_Weapon_Min#Lightning"`
	WeaponDamageMinPoison    ApiMinMax `json:"Damage_Weapon_Min#Poison"`

	WeaponDamageDeltaPhysical  ApiMinMax `json:"Damage_Weapon_Delta#Physical"`
	WeaponDamageDeltaArcane    ApiMinMax `json:"Damage_Weapon_Delta#Arcane"`
	WeaponDamageDeltaCold      ApiMinMax `json:"Damage_Weapon_Delta#Cold"`
	WeaponDamageDeltaFire      ApiMinMax `json:"Damage_Weapon_Delta#Fire"`
	WeaponDamageDeltaHoly      ApiMinMax `json:"Damage_Weapon_Delta#Holy"`
	WeaponDamageDeltaLightning ApiMinMax `json:"Damage_Weapon_Delta#Lightning"`
	WeaponDamageDeltaPoison    ApiMinMax `json:"Damage_Weapon_Delta#Poison"`

	WeaponDamageMinBonusPhysical ApiMinMax `json:"Damage_Weapon_Bonus_Min#Physical`
	WeaponDamageMaxBonusPhysical ApiMinMax `json:"Damage_Weapon_Bonus_Max#Physical`

	WeaponDamagePercentBonusPhysical ApiMinMax `json:"Damage_Weapon_Percent_Bonus#Physical`
}

// Utilities on types.
func (self *ApiHero) GenerateSelectionString() string {

	var buffer bytes.Buffer

	buffer.WriteString(self.Name)
	buffer.WriteString(" (")
	buffer.WriteString(HeroClassMap[self.Class])
	buffer.WriteString(")")

	return buffer.String()
}

func (self *ApiItem) CalculateAverageDamageBonus() float64 {

	var (
		min  = 0.0
		max  = 0.0
		attr = &self.Attributes
	)

	min += attr.DamageMinPhysical.Min
	max += attr.DamageMinPhysical.Min + attr.DamageDeltaPhysical.Min

	min += attr.DamageMinFire.Min
	max += attr.DamageMinFire.Min + attr.DamageDeltaFire.Min

	min += attr.DamageMinArcane.Min
	max += attr.DamageMinArcane.Min + attr.DamageDeltaArcane.Min

	min += attr.DamageMinCold.Min
	max += attr.DamageMinCold.Min + attr.DamageDeltaCold.Min

	min += attr.DamageMinHoly.Min
	max += attr.DamageMinHoly.Min + attr.DamageDeltaHoly.Min

	min += attr.DamageMinLightning.Min
	max += attr.DamageMinLightning.Min + attr.DamageDeltaLightning.Min

	min += attr.DamageMinPoison.Min
	max += attr.DamageMinPoison.Min + attr.DamageDeltaPoison.Min

	return (min + max) / 2
}

func (self *ApiItem) CalculateAverageWeaponDamage() float64 {

	// Start with the correct physical min and max.
	min := self.MinDamage.Min
	max := self.MaxDamage.Max

	attr := &self.Attributes

	// Add elemental damage.
	min += attr.WeaponDamageMinFire.Min
	max += attr.WeaponDamageMinFire.Min + attr.WeaponDamageDeltaFire.Min

	min += attr.WeaponDamageMinArcane.Min
	max += attr.WeaponDamageMinArcane.Min + attr.WeaponDamageDeltaArcane.Min

	min += attr.WeaponDamageMinCold.Min
	max += attr.WeaponDamageMinCold.Min + attr.WeaponDamageDeltaCold.Min

	min += attr.WeaponDamageMinHoly.Min
	max += attr.WeaponDamageMinHoly.Min + attr.WeaponDamageDeltaHoly.Min

	min += attr.WeaponDamageMinLightning.Min
	max += attr.WeaponDamageMinLightning.Min + attr.WeaponDamageDeltaLightning.Min

	min += attr.WeaponDamageMinPoison.Min
	max += attr.WeaponDamageMinPoison.Min + attr.WeaponDamageDeltaPoison.Min

	return (min + max) / 2
}

func (self *ApiItem) GetWeaponType() (weaponType string) {

	switch t := self.Type.Id; {
	case t == "Dagger":
		weaponType = UrlValueWeaponTypeDagger
	case t == "Sword":
		weaponType = UrlValueWeaponTypeSword
	case t == "Mace":
		weaponType = UrlValueWeaponTypeMace
	case t == "Axe":
		weaponType = UrlValueWeaponTypeAxe
	case t == "Polearm":
		weaponType = UrlValueWeaponTypePolearm
	case t == "Spear":
		weaponType = UrlValueWeaponTypeSpear
	case t == "Mighty Weapon":
		weaponType = UrlValueWeaponTypeMightyWeapon
	case t == "Bow":
		weaponType = UrlValueWeaponTypeBow
	case t == "Xbow":
		weaponType = UrlValueWeaponTypeCrossbow
	case t == "HandXbow":
		weaponType = UrlValueWeaponTypeHandCrossbow
	}

	return
}

func (self *ApiItem) IsTwoHandedWeapon() bool {
	return self.Type.TwoHanded
}

// Static utilities.
func LookUpBattleTag(battleTag string, realm string, r *http.Request) (heroes []ApiHero, systemBattleTag string, err error) {

	// If the BattleTag is malformed, return an error.
	systemBattleTag, err = createSystemBattleTag(battleTag)

	if err != nil {
		return
	}

	var buffer bytes.Buffer

	buffer.WriteString("http://")
	buffer.WriteString(realm)
	buffer.WriteString(".battle.net/api/d3/profile/")
	buffer.WriteString(systemBattleTag)
	buffer.WriteString("/")

	appEngineContext := appengine.NewContext(r)
	httpClient := urlfetch.Client(appEngineContext)
	resp, err := httpClient.Get(buffer.String())

	if err != nil {
		err = errors.New("Error contacting the Diablo 3 API.")
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		err = errors.New("Malformed response from the Diablo 3 API.")
		return
	}

	var lookupResponse ApiProfileLookupResponse

	json.Unmarshal(bodyBytes, &lookupResponse)

	if len(lookupResponse.Heroes) == 0 {
		err = errors.New("Could not find profile for the given BattleTag.")
		return
	}

	heroes = lookupResponse.Heroes

	return
}

func LookUpHero(battleTag string, heroId string, realm string, r *http.Request) (hero ApiHero, err error) {

	// If the BattleTag is malformed, return an error.
	systemBattleTag, err := createSystemBattleTag(battleTag)

	if err != nil {
		return
	}

	var buffer bytes.Buffer

	buffer.WriteString("http://")
	buffer.WriteString(realm)
	buffer.WriteString(".battle.net/api/d3/profile/")
	buffer.WriteString(systemBattleTag)
	buffer.WriteString("/hero/")
	buffer.WriteString(heroId)

	lookupUrl := buffer.String()

	appEngineContext := appengine.NewContext(r)
	httpClient := urlfetch.Client(appEngineContext)
	resp, err := httpClient.Get(lookupUrl)

	if err != nil {
		err = errors.New("Error contacting the Diablo 3 API.")
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		err = errors.New("Malformed response from the Diablo 3 API.")
		return
	}

	err = json.Unmarshal(bodyBytes, &hero)

	return
}

func LookUpItem(item *ApiItem, realm string, r *http.Request) (err error) {

	var buffer bytes.Buffer

	buffer.WriteString("http://")
	buffer.WriteString(realm)
	buffer.WriteString(".battle.net/api/d3/data/")
	buffer.WriteString(item.TooltipParams)

	var (
		lookupUrl        = buffer.String()
		appEngineContext = appengine.NewContext(r)
		httpClient       = urlfetch.Client(appEngineContext)
	)

	resp, err := httpClient.Get(lookupUrl)

	if err != nil {
		err = errors.New("Error contacting the Diablo 3 API.")
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		err = errors.New("Malformed response from the Diablo 3 API.")
		return
	}

	err = json.Unmarshal(bodyBytes, item)

	return
}
