package diiicalc

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
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
	Damage            float64 `json:"damage"`
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
	Id            string               `json:"id"`
	Name          string               `json:"name"`
	Icon          string               `json:"icon"`
	DisplayColor  string               `json:"displayColor"`
	TooltipParams string               `json:"tooltipParams"`
	RawAttributes ApiItemAttributesRaw `json:"attributesRaw"`
}

type ApiItemAttributesRaw struct {
	LifeOnHit ApiMinMaxNumber `json:"Hitpoints_On_Hit"`
}

type ApiMinMaxNumber struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// Utilities on types.
func (self *ApiHero) GenerateSelectionString() string {

	var buffer bytes.Buffer

	buffer.WriteString(self.Name)
	buffer.WriteString(" (")
	buffer.WriteString(heroClassMap[self.Class])
	buffer.WriteString(")")

	return buffer.String()
}

// Static utilities.
func lookUpBattleTag(battleTag string, realm string, r *http.Request) (heroes []ApiHero, dashStyleBattleTag string, battleTagLookupError error) {

	// If the BattleTag is malformed, return an error.
	dashStyleBattleTag = strings.Replace(battleTag, "#", "-", -1)

	match, regexError := regexp.MatchString(`[^0-9][^\-]*-[0-9]+`, dashStyleBattleTag)

	if regexError != nil {
		battleTagLookupError = errors.New("Invalid regular expression.")
		return
	}

	if !match {
		battleTagLookupError = errors.New("Given BattleTag is not of the form 'name-number' or 'name#number'.")
		return
	}

	var buffer bytes.Buffer

	buffer.WriteString("http://")
	buffer.WriteString(realm)
	buffer.WriteString(".battle.net/api/d3/profile/")
	buffer.WriteString(dashStyleBattleTag)
	buffer.WriteString("/")

	// We can't use the regular old http.Get method from within
	// the app engine, so use what google provides.
	appEngineContext := appengine.NewContext(r)
	httpClient := urlfetch.Client(appEngineContext)
	resp, err := httpClient.Get(buffer.String())

	if err != nil {
		battleTagLookupError = errors.New("Error contacting the Diablo 3 API.")
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		battleTagLookupError = errors.New("Malformed response from the Diablo 3 API.")
		return
	}

	var lookupResponse ApiProfileLookupResponse

	json.Unmarshal(bodyBytes, &lookupResponse)

	if len(lookupResponse.Heroes) == 0 {
		battleTagLookupError = errors.New("Could not find profile for the given BattleTag.")
		return
	}

	heroes = lookupResponse.Heroes

	return
}

func lookUpHero(battleTag string, heroId string, realm string, r *http.Request) (hero ApiHero, heroLookupError error) {

	// We can trust the battleTag to be good here.

	var buffer bytes.Buffer

	buffer.WriteString("http://")
	buffer.WriteString(realm)
	buffer.WriteString(".battle.net/api/d3/profile/")
	buffer.WriteString(battleTag)
	buffer.WriteString("/hero/")
	buffer.WriteString(heroId)

	lookupUrl := buffer.String()

	fmt.Printf("lookUpHero: Looking Up %s%s", lookupUrl, "\n")

	// We can't use the regular old http.Get method from within
	// the app engine, so use what google provides.
	appEngineContext := appengine.NewContext(r)
	httpClient := urlfetch.Client(appEngineContext)
	resp, err := httpClient.Get(lookupUrl)

	if err != nil {
		heroLookupError = errors.New("Error contacting the Diablo 3 API.")
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		heroLookupError = errors.New("Malformed response from the Diablo 3 API.")
		return
	}

	err = json.Unmarshal(bodyBytes, &hero)

	return
}
