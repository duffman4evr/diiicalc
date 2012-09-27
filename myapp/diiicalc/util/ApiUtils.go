package util

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiProfileLookupResponse struct {
	ErrorCode   string    `json:"code"`
	ErrorReason string    `json:"reason"`
	Heroes      []ApiHero `json:"heroes"`
}

type ApiHero struct {
	Name         string    `json:"name"`
	Class        string    `json:"class"`
	Id           float64   `json:"id"`
	Level        float64   `json:"level"`
	ParagonLevel float64   `json:"paragonLevel"`
	Hardcore     bool      `json:"hardcore"`
	Gender       float64   `json:"gender"`
	LastUpdated  float64   `json:"last-updated"`
	Dead         bool      `json:"dead"`
	Stats        ApiStats  `json:"stats"`
	Skills       ApiSkills `json:"skills"`
	Items        ApiItems  `json:"skills"`
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
	CritDamagePercent      ApiMinMax `json:"Crit_Damage_Percent"`
	CritChancePercent      ApiMinMax `json:"Crit_Chance_Percent"`
	CritPercentBonusCapped ApiMinMax `json:"Crit_Percent_Bonus_Capped"`
	AttacksPerSecondBonus  ApiMinMax `json:"Attacks_Per_Second_Percent"`
	Strength               ApiMinMax `json:"Strength_Item"`
	Intelligence           ApiMinMax `json:"Intelligence_Item"`
	Dexterity              ApiMinMax `json:"Dexterity_Item"`
	Vitality               ApiMinMax `json:"Vitality_Item"`
	// TODO remove these if not used
	//Sockets              ApiMinMax `json:"Sockets"`
	//Crossbow             ApiMinMax `json:"Crossbow"`
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

	// This describes the attacks per second bonus on a
	// weapon only. E.G. it is 0.11 if there is an 11% attack
	// speed bonus.
	AttacksPerSecondWeaponBonus ApiMinMax `json:"Attacks_Per_Second_Item_Percent"`

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

func (self *ApiItems) LookUpAll(realm string, r *http.Request) {

	numItems := 13
	doneChannel := make(chan int, numItems)

	// Look up each item in a goroutine. 
	// Send a signal when done.
	go func() {
		LookUpItem(&self.Head, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.Torso, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.Feet, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.Hands, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.Shoulders, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.Legs, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.Bracers, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.MainHand, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.OffHand, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.Waist, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.RightFinger, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.LeftFinger, realm, r)
		doneChannel <- 1
	}()
	go func() {
		LookUpItem(&self.Neck, realm, r)
		doneChannel <- 1
	}()

	for i := 0; i < numItems; i++ {
		<-doneChannel
	}
}

func (self *ApiHero) DeduceWeaponSetup() string {

	if self.Items.MainHand.Type.TwoHanded {
		return UrlValueWeaponSetupTwoHander
	}

	switch oh := self.Items.OffHand.Type.Id; {
	case oh == "Orb":
		fallthrough
	case oh == "Shield":
		fallthrough
	case oh == "Mojo":
		fallthrough
	case oh == "Quiver":
		fallthrough
	case oh == "": // In the case that they are not wearing an offhand at all.
		return UrlValueWeaponSetupMainHandOffHand
	}

	return UrlValueWeaponSetupDualWield
}

func (self *ApiHero) CalculateTotalAverageDamageBonus() float64 {

	bonus := 0.0

	bonus += self.Items.Head.CalculateAverageDamageBonus()
	bonus += self.Items.Torso.CalculateAverageDamageBonus()
	bonus += self.Items.Feet.CalculateAverageDamageBonus()
	bonus += self.Items.Hands.CalculateAverageDamageBonus()
	bonus += self.Items.Shoulders.CalculateAverageDamageBonus()
	bonus += self.Items.Legs.CalculateAverageDamageBonus()
	bonus += self.Items.Bracers.CalculateAverageDamageBonus()
	bonus += self.Items.Waist.CalculateAverageDamageBonus()
	bonus += self.Items.RightFinger.CalculateAverageDamageBonus()
	bonus += self.Items.LeftFinger.CalculateAverageDamageBonus()
	bonus += self.Items.Neck.CalculateAverageDamageBonus()

	return bonus
}

func (self *ApiHero) CalculateTotalAttackSpeedBonus(weaponSetup string) float64 {

	bonus := 0.0

	if weaponSetup == UrlValueWeaponSetupDualWield {
		bonus += 0.15
	}
	fmt.Printf("\n\n1:%v\n", bonus)
	bonus += self.Items.Head.GetAttackSpeedBonus()
	bonus += self.Items.Torso.GetAttackSpeedBonus()
	bonus += self.Items.Feet.GetAttackSpeedBonus()
	bonus += self.Items.Hands.GetAttackSpeedBonus()
	bonus += self.Items.Shoulders.GetAttackSpeedBonus()
	bonus += self.Items.Legs.GetAttackSpeedBonus()
	bonus += self.Items.Bracers.GetAttackSpeedBonus()
	bonus += self.Items.Waist.GetAttackSpeedBonus()
	bonus += self.Items.RightFinger.GetAttackSpeedBonus()
	bonus += self.Items.LeftFinger.GetAttackSpeedBonus()
	bonus += self.Items.Neck.GetAttackSpeedBonus()
	fmt.Printf("\n\n2:%v\n", bonus)

	return bonus
}

func (self *ApiHero) CalculateTotalCritChance() float64 {

	chance := 0.05

	chance += self.Items.Head.GetCritChanceBonus()
	chance += self.Items.Torso.GetCritChanceBonus()
	chance += self.Items.Feet.GetCritChanceBonus()
	chance += self.Items.Hands.GetCritChanceBonus()
	chance += self.Items.Shoulders.GetCritChanceBonus()
	chance += self.Items.Legs.GetCritChanceBonus()
	chance += self.Items.Bracers.GetCritChanceBonus()
	chance += self.Items.MainHand.GetCritChanceBonus()
	chance += self.Items.OffHand.GetCritChanceBonus()
	chance += self.Items.Waist.GetCritChanceBonus()
	chance += self.Items.RightFinger.GetCritChanceBonus()
	chance += self.Items.LeftFinger.GetCritChanceBonus()
	chance += self.Items.Neck.GetCritChanceBonus()

	return chance
}

func (self *ApiHero) CalculateTotalCritDamageBonus() float64 {

	damage := 0.50

	damage += self.Items.Head.GetCritDamageBonus()
	damage += self.Items.Torso.GetCritDamageBonus()
	damage += self.Items.Feet.GetCritDamageBonus()
	damage += self.Items.Hands.GetCritDamageBonus()
	damage += self.Items.Shoulders.GetCritDamageBonus()
	damage += self.Items.Legs.GetCritDamageBonus()
	damage += self.Items.Bracers.GetCritDamageBonus()
	damage += self.Items.MainHand.GetCritDamageBonus()
	damage += self.Items.OffHand.GetCritDamageBonus()
	damage += self.Items.Waist.GetCritDamageBonus()
	damage += self.Items.RightFinger.GetCritDamageBonus()
	damage += self.Items.LeftFinger.GetCritDamageBonus()
	damage += self.Items.Neck.GetCritDamageBonus()

	return damage
}

func (self *ApiHero) CalculateTotalStrength() float64 {

	strength := 0.0

	// Get base value from class and level.
	if self.Class == UrlValueHeroClassBarbarian {
		strength += 10
		strength += (self.Level - 1) * 3.0
		strength += self.ParagonLevel * 3.0
	} else {
		strength += 8
		strength += (self.Level - 1)
		strength += self.ParagonLevel
	}

	// Add in gear bonuses.
	strength += self.Items.Head.GetStrengthBonus()
	strength += self.Items.Torso.GetStrengthBonus()
	strength += self.Items.Feet.GetStrengthBonus()
	strength += self.Items.Hands.GetStrengthBonus()
	strength += self.Items.Shoulders.GetStrengthBonus()
	strength += self.Items.Legs.GetStrengthBonus()
	strength += self.Items.Bracers.GetStrengthBonus()
	strength += self.Items.MainHand.GetStrengthBonus()
	strength += self.Items.OffHand.GetStrengthBonus()
	strength += self.Items.Waist.GetStrengthBonus()
	strength += self.Items.RightFinger.GetStrengthBonus()
	strength += self.Items.LeftFinger.GetStrengthBonus()
	strength += self.Items.Neck.GetStrengthBonus()

	return strength
}

func (self *ApiHero) CalculateTotalDexterity() float64 {

	dexterity := 0.0

	// Get base value from class and level.
	if self.Class == UrlValueHeroClassMonk || self.Class == UrlValueHeroClassDemonHunter {
		dexterity += 10
		dexterity += (self.Level - 1) * 3.0
		dexterity += self.ParagonLevel * 3.0
	} else {
		dexterity += 8
		dexterity += (self.Level - 1)
		dexterity += self.ParagonLevel
	}

	// Add in gear bonuses.
	dexterity += self.Items.Head.GetDexterityBonus()
	dexterity += self.Items.Torso.GetDexterityBonus()
	dexterity += self.Items.Feet.GetDexterityBonus()
	dexterity += self.Items.Hands.GetDexterityBonus()
	dexterity += self.Items.Shoulders.GetDexterityBonus()
	dexterity += self.Items.Legs.GetDexterityBonus()
	dexterity += self.Items.Bracers.GetDexterityBonus()
	dexterity += self.Items.MainHand.GetDexterityBonus()
	dexterity += self.Items.OffHand.GetDexterityBonus()
	dexterity += self.Items.Waist.GetDexterityBonus()
	dexterity += self.Items.RightFinger.GetDexterityBonus()
	dexterity += self.Items.LeftFinger.GetDexterityBonus()
	dexterity += self.Items.Neck.GetDexterityBonus()

	return dexterity
}

func (self *ApiHero) CalculateTotalIntelligence() float64 {

	intelligence := 0.0

	// Get base value from class and level.
	if self.Class == UrlValueHeroClassWizard || self.Class == UrlValueHeroClassWitchDoctor {
		intelligence += 10
		intelligence += (self.Level - 1) * 3.0
		intelligence += self.ParagonLevel * 3.0
	} else {
		intelligence += 8
		intelligence += self.Level - 1
		intelligence += self.ParagonLevel
	}

	// Add in gear bonuses.
	intelligence += self.Items.Head.GetIntelligenceBonus()
	intelligence += self.Items.Torso.GetIntelligenceBonus()
	intelligence += self.Items.Feet.GetIntelligenceBonus()
	intelligence += self.Items.Hands.GetIntelligenceBonus()
	intelligence += self.Items.Shoulders.GetIntelligenceBonus()
	intelligence += self.Items.Legs.GetIntelligenceBonus()
	intelligence += self.Items.Bracers.GetIntelligenceBonus()
	intelligence += self.Items.MainHand.GetIntelligenceBonus()
	intelligence += self.Items.OffHand.GetIntelligenceBonus()
	intelligence += self.Items.Waist.GetIntelligenceBonus()
	intelligence += self.Items.RightFinger.GetIntelligenceBonus()
	intelligence += self.Items.LeftFinger.GetIntelligenceBonus()
	intelligence += self.Items.Neck.GetIntelligenceBonus()

	return intelligence
}

func (self *ApiHero) CalculateTotalVitality() float64 {

	vitality := 0.0

	// Get base value from class and level.
	vitality += 9
	vitality += (self.Level - 1) * 2.0
	vitality += self.ParagonLevel * 2.0

	// Add in gear bonuses.
	vitality += self.Items.Head.GetVitalityBonus()
	vitality += self.Items.Torso.GetVitalityBonus()
	vitality += self.Items.Feet.GetVitalityBonus()
	vitality += self.Items.Hands.GetVitalityBonus()
	vitality += self.Items.Shoulders.GetVitalityBonus()
	vitality += self.Items.Legs.GetVitalityBonus()
	vitality += self.Items.Bracers.GetVitalityBonus()
	vitality += self.Items.MainHand.GetVitalityBonus()
	vitality += self.Items.OffHand.GetVitalityBonus()
	vitality += self.Items.Waist.GetVitalityBonus()
	vitality += self.Items.RightFinger.GetVitalityBonus()
	vitality += self.Items.LeftFinger.GetVitalityBonus()
	vitality += self.Items.Neck.GetVitalityBonus()

	return vitality
}

func (self *ApiItem) GetWeaponDps() float64 {
	return self.Dps.Min
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

func (self *ApiItem) GetWeaponBaseAttackSpeed() float64 {
	return self.Attributes.AttacksPerSecondBase.Min
}

func (self *ApiItem) GetWeaponAttackSpeedBonus() float64 {
	return self.Attributes.AttacksPerSecondWeaponBonus.Min
}

func (self *ApiItem) GetAttackSpeedBonus() float64 {
	return self.Attributes.AttacksPerSecondBonus.Min
}

func (self *ApiItem) GetCritChanceBonus() float64 {
	return self.Attributes.CritPercentBonusCapped.Min
}

func (self *ApiItem) GetCritDamageBonus() float64 {

	damage := self.Attributes.CritDamagePercent.Min

	for i := 0; i < len(self.Gems); i++ {
		damage += self.Gems[i].Attributes.CritDamagePercent.Min
	}

	return damage
}

func (self *ApiItem) GetStrengthBonus() float64 {

	strength := self.Attributes.Strength.Min

	for i := 0; i < len(self.Gems); i++ {
		strength += self.Gems[i].Attributes.Strength.Min
	}

	return strength
}

func (self *ApiItem) GetIntelligenceBonus() float64 {

	intelligence := self.Attributes.Intelligence.Min

	for i := 0; i < len(self.Gems); i++ {
		intelligence += self.Gems[i].Attributes.Intelligence.Min
	}

	return intelligence
}

func (self *ApiItem) GetVitalityBonus() float64 {

	vitality := self.Attributes.Vitality.Min

	for i := 0; i < len(self.Gems); i++ {
		vitality += self.Gems[i].Attributes.Vitality.Min
	}

	return vitality
}

func (self *ApiItem) GetDexterityBonus() float64 {

	dexterity := self.Attributes.Dexterity.Min

	for i := 0; i < len(self.Gems); i++ {
		dexterity += self.Gems[i].Attributes.Dexterity.Min
	}

	return dexterity
}

func (self *ApiItem) GetWeaponType() (weaponType string) {

	switch t := self.Type; {
	case t.Id == "Dagger":
		weaponType = UrlValueWeaponTypeDagger
	case t.Id == "Sword":
		if t.TwoHanded {
			weaponType = UrlValueWeaponTypeThSword
		} else {
			weaponType = UrlValueWeaponTypeOhSword
		}
	case t.Id == "Mace":
		if t.TwoHanded {
			weaponType = UrlValueWeaponTypeThMace
		} else {
			weaponType = UrlValueWeaponTypeOhMace
		}
	case t.Id == "Axe":
		if t.TwoHanded {
			weaponType = UrlValueWeaponTypeThAxe
		} else {
			weaponType = UrlValueWeaponTypeOhAxe
		}
	case t.Id == "Polearm":
		weaponType = UrlValueWeaponTypePolearm
	case t.Id == "Spear":
		weaponType = UrlValueWeaponTypeSpear
	case t.Id == "Mighty Weapon":
		if t.TwoHanded {
			weaponType = UrlValueWeaponTypeThMightyWeapon
		} else {
			weaponType = UrlValueWeaponTypeOhMightyWeapon
		}
	case t.Id == "Bow":
		weaponType = UrlValueWeaponTypeBow
	case t.Id == "Xbow":
		weaponType = UrlValueWeaponTypeCrossbow
	case t.Id == "HandXbow":
		weaponType = UrlValueWeaponTypeHandCrossbow
	}

	return
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
