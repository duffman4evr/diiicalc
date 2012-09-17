package util

const (
	// URL keys.
	UrlKeyBattleTagUser   = "btu"
	UrlKeyBattleTagSystem = "bts"
	UrlKeyHeroIdUser      = "hiu"
	UrlKeyFindButton      = "fi"
	UrlKeyDefensiveButton = "vd"
	UrlKeyOffensiveButton = "vo"
	UrlKeyRealm           = "re"
	UrlKeyHeroes          = "he"

	UrlKeyLeftCompareType   = "lct"
	UrlKeyCenterCompareType = "cct"
	UrlKeyRightCompareType  = "rct"

	UrlKeyLeftCompareValue   = "lcv"
	UrlKeyCenterCompareValue = "ccv"
	UrlKeyRightCompareValue  = "rcv"

	UrlKeyHeroId       = "hi"
	UrlKeyHeroName     = "hn"
	UrlKeyHeroClass    = "hc"
	UrlKeyLevel        = "lv"
	UrlKeyStrength     = "st"
	UrlKeyIntelligence = "it"
	UrlKeyDexterity    = "de"
	UrlKeyVitality     = "vi"

	// Offensive
	UrlKeyCritChance                 = "cc"
	UrlKeyCritDamage                 = "cd"
	UrlKeyAttackSpeedBonus           = "sb"
	UrlKeyAverageDamageBonus         = "ab"
	UrlKeyTwoHandedWeapon            = "th"
	UrlKeyMainWeaponAverageDamage    = "da"
	UrlKeyMainWeaponAttackSpeedBase  = "wa"
	UrlKeyMainWeaponAttackSpeedBonus = "wb"
	UrlKeyMainWeaponType             = "wt"

	// Defensive
	UrlKeyArmor           = "ar"
	UrlKeyLifePercent     = "lp"
	UrlKeyLifeOnHit       = "lh"
	UrlKeyLifeRegen       = "lr"
	UrlKeyBlockMin        = "bi"
	UrlKeyBlockMax        = "ba"
	UrlKeyBlockChance     = "bc"
	UrlKeyResistArcane    = "ra"
	UrlKeyResistFire      = "rf"
	UrlKeyResistLightning = "rl"
	UrlKeyResistPoison    = "rp"
	UrlKeyResistCold      = "rc"
	UrlKeyResistPhysical  = "ry"

	UrlKeySkill1 = "s1"
	UrlKeySkill2 = "s2"
	UrlKeySkill3 = "s3"
	UrlKeySkill4 = "s4"
	UrlKeySkill5 = "s5"
	UrlKeySkill6 = "s6"

	UrlKeyRune1 = "r1"
	UrlKeyRune2 = "r2"
	UrlKeyRune3 = "r3"
	UrlKeyRune4 = "r4"
	UrlKeyRune5 = "r5"
	UrlKeyRune6 = "r6"

	UrlKeyPassive1 = "p1"
	UrlKeyPassive2 = "p2"
	UrlKeyPassive3 = "p3"

	// URL values.
	UrlValueCompareTypeVitality       = "vit"
	UrlValueCompareTypeResist         = "res"
	UrlValueCompareTypeArmor          = "arm"
	UrlValueCompareTypePercentLife    = "pli"
	UrlValueCompareTypeDexterity      = "dex"
	UrlValueCompareTypeStrength       = "str"
	UrlValueCompareTypeIntelligence   = "int"
	UrlValueCompareTypeAttackSpeed    = "apd"
	UrlValueCompareTypeCritChance     = "crc"
	UrlValueCompareTypeCritDamage     = "crd"
	UrlValueCompareTypeAverageWeapDmg = "awd"
	UrlValueCompareTypeMinMaxWeapDmg  = "mmd"

	UrlValueHeroClassBarbarian   = "barbarian"
	UrlValueHeroClassMonk        = "monk"
	UrlValueHeroClassWitchDoctor = "witch-doctor"
	UrlValueHeroClassWizard      = "wizard"
	UrlValueHeroClassDemonHunter = "demon-hunter"

	UrlValueRealmUs = "us"
	UrlValueRealmEu = "eu"
	UrlValueRealmTw = "tw"
	UrlValueRealmKr = "kr"

	UrlValueWeaponTypeDagger       = "dagger"
	UrlValueWeaponTypeSword        = "sword"
	UrlValueWeaponTypeMace         = "mace"
	UrlValueWeaponTypeAxe          = "axe"
	UrlValueWeaponTypePolearm      = "polearm"
	UrlValueWeaponTypeSpear        = "spear"
	UrlValueWeaponTypeMightyWeapon = "mightyWeapon"

	// Other
	MitigationSourceArmor       = "Armor"
	MitigationSourceResistances = "Resistance"
	MitigationSourceMeleeClass  = "Monk/Barb Bonus"
	MitigationSourceDodge       = "Dodge"
)

var (
	UrlKeysActiveSkills  = []string{UrlKeySkill1, UrlKeySkill2, UrlKeySkill3, UrlKeySkill4, UrlKeySkill5, UrlKeySkill6}
	UrlKeysActiveRunes   = []string{UrlKeyRune1, UrlKeyRune2, UrlKeyRune3, UrlKeyRune4, UrlKeyRune5, UrlKeyRune6}
	UrlKeysPassiveSkills = []string{UrlKeyPassive1, UrlKeyPassive2, UrlKeyPassive3}

	HeroClassMap = map[string]string{
		UrlValueHeroClassBarbarian:   "Barbarian",
		UrlValueHeroClassMonk:        "Monk",
		UrlValueHeroClassWitchDoctor: "Witch Doctor",
		UrlValueHeroClassWizard:      "Wizard",
		UrlValueHeroClassDemonHunter: "Demon Hunter",
	}

	CompareTypeMap = map[string]string{
		UrlValueCompareTypeVitality:     "Vitality",
		UrlValueCompareTypeResist:       "Resist",
		UrlValueCompareTypeArmor:        "Armor",
		UrlValueCompareTypePercentLife:  `% Life`,
		UrlValueCompareTypeDexterity:    "Dexterity",
		UrlValueCompareTypeStrength:     "Strength",
		UrlValueCompareTypeIntelligence: "Intelligence",
		UrlValueCompareTypeAttackSpeed:  "Attack Speed",
		UrlValueCompareTypeCritChance:   "Crit Chance",
		UrlValueCompareTypeCritDamage:   "Crit Damage",
	}
)
