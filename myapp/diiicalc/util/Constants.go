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
	UrlKeyWeaponSetup                = "ws"
	UrlKeyMainWeaponType             = "wt"
	UrlKeyCritChance                 = "cc"
	UrlKeyCritDamageBonus            = "cd"
	UrlKeyAttackSpeedBonus           = "sb"
	UrlKeyAverageDamageBonus         = "ab"
	UrlKeyMainWeaponAverageDamage    = "da"
	UrlKeyMainWeaponAttackSpeedBase  = "wa"
	UrlKeyMainWeaponAttackSpeedBonus = "wb"
	UrlKeyOffWeaponAverageDamage     = "od"
	UrlKeyOffWeaponAttackSpeedBase   = "oa"
	UrlKeyOffWeaponAttackSpeedBonus  = "ob"

	UrlKeyLwcWeaponSetup                    = "lwcws"
	UrlKeyLwcMainHandWeaponType             = "lwcmhty"
	UrlKeyLwcMainHandWeaponDps              = "lwcwmhdp"
	UrlKeyLwcMainHandWeaponAttackSpeedBonus = "lwcwmhas"
	UrlKeyLwcMainHandWeaponCritDamageBonus  = "lwcwmhcd"
	UrlKeyLwcMainHandWeaponMainStatBonus    = "lwcwmhms"
	UrlKeyLwcOffHandWeaponType              = "lwcohwty"
	UrlKeyLwcOffHandWeaponDps               = "lwcwohwdp"
	UrlKeyLwcOffHandWeaponAttackSpeedBonus  = "lwcwohwas"
	UrlKeyLwcOffHandWeaponCritDamageBonus   = "lwcwohwcd"
	UrlKeyLwcOffHandWeaponMainStatBonus     = "lwcwohwms"
	UrlKeyLwcOffHandType                    = "lwcohoty"
	UrlKeyLwcOffHandAverageDamage           = "lwcwohoad"
	UrlKeyLwcOffHandAttackSpeedBonus        = "lwcwohoas"
	UrlKeyLwcOffHandCritChanceBonus         = "lwcwohocc"
	UrlKeyLwcOffHandMainStatBonus           = "lwcwohoms"

	UrlKeyRwcWeaponSetup                    = "rwcws"
	UrlKeyRwcMainHandWeaponType             = "rwcmhty"
	UrlKeyRwcMainHandWeaponDps              = "rwcwmhdp"
	UrlKeyRwcMainHandWeaponAttackSpeedBonus = "rwcwmhas"
	UrlKeyRwcMainHandWeaponCritDamageBonus  = "rwcwmhcd"
	UrlKeyRwcMainHandWeaponMainStatBonus    = "rwcwmhms"
	UrlKeyRwcOffHandWeaponType              = "rwcohwty"
	UrlKeyRwcOffHandWeaponDps               = "rwcwohwdp"
	UrlKeyRwcOffHandWeaponAttackSpeedBonus  = "rwcwohwas"
	UrlKeyRwcOffHandWeaponCritDamageBonus   = "rwcwohwcd"
	UrlKeyRwcOffHandWeaponMainStatBonus     = "rwcwohwms"
	UrlKeyRwcOffHandType                    = "rwcohoty"
	UrlKeyRwcOffHandAverageDamage           = "rwcwohoad"
	UrlKeyRwcOffHandAttackSpeedBonus        = "rwcwohoas"
	UrlKeyRwcOffHandCritChanceBonus         = "rwcwohocc"
	UrlKeyRwcOffHandMainStatBonus           = "rwcwohoms"

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
	UrlValueCompareTypeMainStat       = "mas"
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

	UrlValueWeaponTypeOhAxe          = "0"
	UrlValueWeaponTypeThAxe          = "1"
	UrlValueWeaponTypeOhMace         = "2"
	UrlValueWeaponTypeThMace         = "3"
	UrlValueWeaponTypeOhMightyWeapon = "4"
	UrlValueWeaponTypeThMightyWeapon = "5"
	UrlValueWeaponTypeOhSword        = "6"
	UrlValueWeaponTypeThSword        = "7"

	UrlValueWeaponTypeBow             = "8"
	UrlValueWeaponTypeCeremonialKnife = "9"
	UrlValueWeaponTypeCrossbow        = "10"
	UrlValueWeaponTypeDagger          = "11"
	UrlValueWeaponTypeDaibo           = "12"
	UrlValueWeaponTypeFistWeapon      = "13"
	UrlValueWeaponTypeHandCrossbow    = "14"
	UrlValueWeaponTypePolearm         = "15"
	UrlValueWeaponTypeSpear           = "16"
	UrlValueWeaponTypeStaff           = "17"
	UrlValueWeaponTypeWand            = "18"

	UrlValueOffHandTypeSource = "source"
	UrlValueOffHandTypeMojo   = "mojo"
	UrlValueOffHandTypeQuiver = "quiver"
	UrlValueOffHandTypeShield = "shield"

	UrlValueWeaponSetupTwoHander       = "th"
	UrlValueWeaponSetupMainHandOffHand = "mo"
	UrlValueWeaponSetupDualWield       = "dw"

	// Other
	MitigationSourceArmor       = "Armor"
	MitigationSourceResistances = "Resistance"
	MitigationSourceMeleeClass  = "Monk/Barb Bonus"
	MitigationSourceDodge       = "Dodge"

	MapKeyWeaponSetup                    = "0"
	MapKeyMainHandWeaponType             = "1"
	MapKeyMainHandWeaponDps              = "2"
	MapKeyMainHandWeaponAttackSpeedBonus = "3"
	MapKeyMainHandWeaponCritDamageBonus  = "4"
	MapKeyMainHandWeaponMainStatBonus    = "5"
	MapKeyOffHandWeaponType              = "6"
	MapKeyOffHandWeaponDps               = "7"
	MapKeyOffHandWeaponAttackSpeedBonus  = "8"
	MapKeyOffHandWeaponCritDamageBonus   = "9"
	MapKeyOffHandWeaponMainStatBonus     = "10"
	MapKeyOffHandType                    = "11"
	MapKeyOffHandAverageDamage           = "12"
	MapKeyOffHandAttackSpeedBonus        = "13"
	MapKeyOffHandCritChanceBonus         = "14"
	MapKeyOffHandMainStatBonus           = "15"
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
		UrlValueCompareTypeVitality:    "Vitality",
		UrlValueCompareTypeResist:      "Resist",
		UrlValueCompareTypeArmor:       "Armor",
		UrlValueCompareTypePercentLife: `% Life`,
		UrlValueCompareTypeDexterity:   "Dexterity",
		UrlValueCompareTypeAttackSpeed: "Attack Speed",
		UrlValueCompareTypeCritChance:  "Crit Chance",
		UrlValueCompareTypeCritDamage:  "Crit Damage",
	}

	WeaponSetupMap = map[string]string{
		UrlValueWeaponSetupTwoHander:       "Two Hander",
		UrlValueWeaponSetupMainHandOffHand: "Main Hand / Off Hand",
		UrlValueWeaponSetupDualWield:       "Dual Wield",
	}

	WeaponTypeMap = map[string]string{
		UrlValueWeaponTypeOhAxe:           "Axe (1.3)",
		UrlValueWeaponTypeThAxe:           "Axe (1.0)",
		UrlValueWeaponTypeOhMace:          "Mace (1.2)",
		UrlValueWeaponTypeThMace:          "Mace (0.9)",
		UrlValueWeaponTypeOhMightyWeapon:  "Mighty Weapon (1.3)",
		UrlValueWeaponTypeThMightyWeapon:  "Mighty Weapon (1.0)",
		UrlValueWeaponTypeOhSword:         "Sword (1.4)",
		UrlValueWeaponTypeThSword:         "Sword (1.1)",
		UrlValueWeaponTypeBow:             "Bow (1.4)",
		UrlValueWeaponTypeCeremonialKnife: "Ceremonial Knife (1.4)",
		UrlValueWeaponTypeCrossbow:        "Crossbow (1.1)",
		UrlValueWeaponTypeDagger:          "Dagger (1.5)",
		UrlValueWeaponTypeDaibo:           "Daibo (1.1)",
		UrlValueWeaponTypeFistWeapon:      "Fist Weapon (1.4)",
		UrlValueWeaponTypeHandCrossbow:    "Hand Crossbow (1.6)",
		UrlValueWeaponTypePolearm:         "Polearm (0.95)",
		UrlValueWeaponTypeSpear:           "Spear (1.2)",
		UrlValueWeaponTypeStaff:           "Staff (1.0)",
		UrlValueWeaponTypeWand:            "Wand (1.4)",
	}

	OffHandTypeMap = map[string]string{
		UrlValueOffHandTypeSource:           "Source",
		UrlValueOffHandTypeMojo:           "Mojo",
		UrlValueOffHandTypeQuiver:          "Quiver",
		UrlValueOffHandTypeShield:          "Shield",
	}

	WearableOhWeaponMap = map[string][]string{
		UrlValueHeroClassBarbarian: []string{
			UrlValueWeaponTypeOhAxe,
			UrlValueWeaponTypeDagger,
			UrlValueWeaponTypeOhMace,
			UrlValueWeaponTypeOhMightyWeapon,
			UrlValueWeaponTypeSpear,
			UrlValueWeaponTypeOhSword,
		},
		UrlValueHeroClassMonk: []string{
			UrlValueWeaponTypeOhAxe,
			UrlValueWeaponTypeDagger,
			UrlValueWeaponTypeFistWeapon,
			UrlValueWeaponTypeOhMace,
			UrlValueWeaponTypeSpear,
			UrlValueWeaponTypeOhSword,
		},
		UrlValueHeroClassWitchDoctor: []string{
			UrlValueWeaponTypeOhAxe,
			UrlValueWeaponTypeCeremonialKnife,
			UrlValueWeaponTypeDagger,
			UrlValueWeaponTypeOhMace,
			UrlValueWeaponTypeSpear,
			UrlValueWeaponTypeOhSword,
		},
		UrlValueHeroClassWizard: []string{
			UrlValueWeaponTypeOhAxe,
			UrlValueWeaponTypeDagger,
			UrlValueWeaponTypeOhMace,
			UrlValueWeaponTypeSpear,
			UrlValueWeaponTypeOhSword,
			UrlValueWeaponTypeWand,
		},
		UrlValueHeroClassDemonHunter: []string{
			UrlValueWeaponTypeBow,
			UrlValueWeaponTypeCrossbow,
			UrlValueWeaponTypeHandCrossbow,
		},
	}

	WearableThWeaponMap = map[string][]string{
		UrlValueHeroClassBarbarian: []string{
			UrlValueWeaponTypeThAxe,
			UrlValueWeaponTypeThMace,
			UrlValueWeaponTypeThMightyWeapon,
			UrlValueWeaponTypePolearm,
			UrlValueWeaponTypeThSword,
		},
		UrlValueHeroClassMonk: []string{
			UrlValueWeaponTypeThAxe,
			UrlValueWeaponTypeDaibo,
			UrlValueWeaponTypeThMace,
			UrlValueWeaponTypePolearm,
			UrlValueWeaponTypeStaff,
			UrlValueWeaponTypeThSword,
		},
		UrlValueHeroClassWitchDoctor: []string{
			UrlValueWeaponTypeThAxe,
			UrlValueWeaponTypeBow,
			UrlValueWeaponTypeCrossbow,
			UrlValueWeaponTypeThMace,
			UrlValueWeaponTypePolearm,
			UrlValueWeaponTypeStaff,
			UrlValueWeaponTypeThSword,
		},
		UrlValueHeroClassWizard: []string{
			UrlValueWeaponTypeThAxe,
			UrlValueWeaponTypeBow,
			UrlValueWeaponTypeCrossbow,
			UrlValueWeaponTypeThMace,
			UrlValueWeaponTypeStaff,
			UrlValueWeaponTypeThSword,
		},
		UrlValueHeroClassDemonHunter: []string{},
	}

	WearableOhMap = map[string][]string{
		UrlValueHeroClassBarbarian: []string{UrlValueOffHandTypeShield},
		UrlValueHeroClassMonk: []string{UrlValueOffHandTypeShield},
		UrlValueHeroClassWitchDoctor: []string{UrlValueOffHandTypeMojo, UrlValueOffHandTypeShield},
		UrlValueHeroClassWizard: []string{UrlValueOffHandTypeSource, UrlValueOffHandTypeShield},
		UrlValueHeroClassDemonHunter: []string{UrlValueOffHandTypeQuiver},
	}

	LeftCompareMap = map[string]string{
		MapKeyWeaponSetup:                    UrlKeyLwcWeaponSetup,
		MapKeyMainHandWeaponType:             UrlKeyLwcMainHandWeaponType,
		MapKeyMainHandWeaponDps:              UrlKeyLwcMainHandWeaponDps,
		MapKeyMainHandWeaponAttackSpeedBonus: UrlKeyLwcMainHandWeaponAttackSpeedBonus,
		MapKeyMainHandWeaponCritDamageBonus:  UrlKeyLwcMainHandWeaponCritDamageBonus,
		MapKeyMainHandWeaponMainStatBonus:    UrlKeyLwcMainHandWeaponMainStatBonus,
		MapKeyOffHandWeaponType:              UrlKeyLwcOffHandWeaponType,
		MapKeyOffHandWeaponDps:               UrlKeyLwcOffHandWeaponDps,
		MapKeyOffHandWeaponAttackSpeedBonus:  UrlKeyLwcOffHandWeaponAttackSpeedBonus,
		MapKeyOffHandWeaponCritDamageBonus:   UrlKeyLwcOffHandWeaponCritDamageBonus,
		MapKeyOffHandWeaponMainStatBonus:     UrlKeyLwcOffHandWeaponMainStatBonus,
		MapKeyOffHandType:                    UrlKeyLwcOffHandType,
		MapKeyOffHandAverageDamage:           UrlKeyLwcOffHandAverageDamage,
		MapKeyOffHandAttackSpeedBonus:        UrlKeyLwcOffHandAttackSpeedBonus,
		MapKeyOffHandCritChanceBonus:         UrlKeyLwcOffHandCritChanceBonus,
		MapKeyOffHandMainStatBonus:           UrlKeyLwcOffHandMainStatBonus,
	}

	RightCompareMap = map[string]string{
		MapKeyWeaponSetup:                    UrlKeyRwcWeaponSetup,
		MapKeyMainHandWeaponType:             UrlKeyRwcMainHandWeaponType,
		MapKeyMainHandWeaponDps:              UrlKeyRwcMainHandWeaponDps,
		MapKeyMainHandWeaponAttackSpeedBonus: UrlKeyRwcMainHandWeaponAttackSpeedBonus,
		MapKeyMainHandWeaponCritDamageBonus:  UrlKeyRwcMainHandWeaponCritDamageBonus,
		MapKeyMainHandWeaponMainStatBonus:    UrlKeyRwcMainHandWeaponMainStatBonus,
		MapKeyOffHandWeaponType:              UrlKeyRwcOffHandWeaponType,
		MapKeyOffHandWeaponDps:               UrlKeyRwcOffHandWeaponDps,
		MapKeyOffHandWeaponAttackSpeedBonus:  UrlKeyRwcOffHandWeaponAttackSpeedBonus,
		MapKeyOffHandWeaponCritDamageBonus:   UrlKeyRwcOffHandWeaponCritDamageBonus,
		MapKeyOffHandWeaponMainStatBonus:     UrlKeyRwcOffHandWeaponMainStatBonus,
		MapKeyOffHandType:                    UrlKeyRwcOffHandType,
		MapKeyOffHandAverageDamage:           UrlKeyRwcOffHandAverageDamage,
		MapKeyOffHandAttackSpeedBonus:        UrlKeyRwcOffHandAttackSpeedBonus,
		MapKeyOffHandCritChanceBonus:         UrlKeyRwcOffHandCritChanceBonus,
		MapKeyOffHandMainStatBonus:           UrlKeyRwcOffHandMainStatBonus,
	}
)
