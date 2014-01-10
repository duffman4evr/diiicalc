package com.diiicalc.core;

public class Constants
{
   public static final String PROFILE_API_URL_PREFIX = "/api/d3/profile";
   public static final String DATA_API_URL_PREFIX = "/api/d3/data";

   public static final String HERO_TYPE_BARBARIAN    = "barbarian";
   public static final String HERO_TYPE_MONK         = "monk";
   public static final String HERO_TYPE_WITCH_DOCTOR = "witch-doctor";
   public static final String HERO_TYPE_WIZARD       = "wizard";
   public static final String HERO_TYPE_DEMON_HUNTER = "demon-hunter";

   public static final String WEAPON_TYPE_DAGGER = "Dagger";
   public static final String WEAPON_TYPE_SWORD = "Sword";
   public static final String WEAPON_TYPE_SWORD_1H = "Sword1H";
   public static final String WEAPON_TYPE_SWORD_2H = "Sword2H";
   public static final String WEAPON_TYPE_MACE = "Mace";
   public static final String WEAPON_TYPE_MACE_1H = "Mace1H";
   public static final String WEAPON_TYPE_MACE_2H = "Mace2H";
   public static final String WEAPON_TYPE_AXE = "Axe";
   public static final String WEAPON_TYPE_AXE_1H = "Axe1H";
   public static final String WEAPON_TYPE_AXE_2H = "Axe2H";
   public static final String WEAPON_TYPE_POLEARM = "Polearm";
   public static final String WEAPON_TYPE_SPEAR = "Spear";
   public static final String WEAPON_TYPE_MIGHTY_WEAPON = "Mighty Weapon";
   public static final String WEAPON_TYPE_MIGHTY_WEAPON_1H = "MightyWeapon1H";
   public static final String WEAPON_TYPE_MIGHTY_WEAPON_2H = "MightyWeapon2H";
   public static final String WEAPON_TYPE_BOW = "Bow";
   public static final String WEAPON_TYPE_XBOW = "Xbow";
   public static final String WEAPON_TYPE_HAND_XBOW = "HandXbow";
   public static final String WEAPON_TYPE_CROSSBOW = "Crossbow";
   public static final String WEAPON_TYPE_HAND_CROSSBOW = "HandCrossbow";
   public static final String WEAPON_TYPE_WAND = "Wand";
   public static final String WEAPON_TYPE_DIABO = "Diabo";
   public static final String WEAPON_TYPE_FIST_WEAPON = "Fist Weapon";
   public static final String WEAPON_TYPE_FISTWEAPON = "FistWeapon";
   public static final String WEAPON_TYPE_CEREMONIAL_KNIFE = "Ceremonial Knife";
   public static final String WEAPON_TYPE_CEREMONIAL_DAGGER = "CeremonialDagger";
   public static final String WEAPON_TYPE_STAFF = "Staff";

   public static final String DAMAGE_MODIFIER_ARMOR = "Armor";
   public static final String DAMAGE_MODIFIER_RESISTS = "All Resist";
   public static final String DAMAGE_MODIFIER_MELEE = "Melee";

   // -----
   // Item slot names.
   // -----

   public static final String ITEM_SLOT_MAIN_HAND = "mainHand";
   public static final String ITEM_SLOT_OFF_HAND = "offHand";

   // -----
   // Item Attributes
   // -----

   // The four basic stats.
   public static final String ITEM_ATTR_STRENGTH = "Strength_Item";
   public static final String ITEM_ATTR_INTELLIGENCE = "Intelligence_Item";
   public static final String ITEM_ATTR_DEXTERITY = "Dexterity_Item";
   public static final String ITEM_ATTR_VITALITY = "Vitality_Item";

   public static final String BUG_ITEM_ATTR_STRENGTH = "Strength";
   public static final String BUG_ITEM_ATTR_INTELLIGENCE = "Intelligence";
   public static final String BUG_ITEM_ATTR_DEXTERITY = "Dexterity";
   public static final String BUG_ITEM_ATTR_VITALITY = "Vitality";

   // Life related.
   public static final String ITEM_ATTR_LIFE_REGEN = "Hitpoints_Regen_Per_Second";
   public static final String ITEM_ATTR_LIFE_PERCENT = "Hitpoints_Max_Percent_Bonus_Item";
   public static final String ITEM_ATTR_LIFE_ON_HIT = "Hitpoints_On_Hit";

   // Armor related.
   public static final String ITEM_ATTR_ARMOR = "Armor_Item";
   public static final String ITEM_ATTR_ARMOR_BONUS = "Armor_Bonus_Item";

   // Direct damage reduction related.
   public static final String ITEM_ATTR_RANGED_DR_PERCENT = "Damage_Percent_Reduction_From_Ranged";
   public static final String ITEM_ATTR_MELEE_DR_PERCENT = "Damage_Percent_Reduction_From_Melee";
   public static final String ITEM_ATTR_ELITE_DR_PERCENT = "Damage_Percent_Reduction_From_Elites";

   // Block related.
   public static final String ITEM_ATTR_BLOCK_CHANCE = "Block_Chance_Item";
   public static final String ITEM_ATTR_BLOCK_CHANCE_BONUS = "Block_Chance_Bonus_Item";
   public static final String ITEM_ATTR_BLOCK_AMOUNT_MIN = "Block_Amount_Item_Min";
   public static final String ITEM_ATTR_BLOCK_AMOUNT_DELTA = "Block_Amount_Item_Delta";

   // Resistances related.
   public static final String ITEM_ATTR_RESIST_ALL       = "Resistance_All";
   public static final String ITEM_ATTR_RESIST_PHYSICAL  = "Resistance#Physical";
   public static final String ITEM_ATTR_RESIST_ARCANE    = "Resistance#Arcane";
   public static final String ITEM_ATTR_RESIST_COLD      = "Resistance#Cold";
   public static final String ITEM_ATTR_RESIST_FIRE      = "Resistance#Fire";
   public static final String ITEM_ATTR_RESIST_HOLY      = "Resistance#Holy";
   public static final String ITEM_ATTR_RESIST_LIGHTNING = "Resistance#Lightning";
   public static final String ITEM_ATTR_RESIST_POISON    = "Resistance#Poison";

   // Crit related.
   public static final String ITEM_ATTR_CRIT_DMG = "Crit_Damage_Percent";
   public static final String ITEM_ATTR_CRIT_CHANCE = "Crit_Percent_Bonus_Capped";

   // Attack speed related.
   public static final String ITEM_ATTR_ATTACK_SPEED_PERCENT = "Attacks_Per_Second_Percent";

   // These values describe bonus raw damage affixes that are usually present on
   // non weapons. Particularly: rings, amulets, offhands.
   public static final String ITEM_ATTR_DAMAGE_MIN_PREFIX = "Damage_Min#";
   public static final String ITEM_ATTR_DAMAGE_DELTA_PREFIX = "Damage_Delta#";

   // ...
   // Following are properties for weapons specifically.
   // ...

   public static final String WEAPON_ATTR_ATTACKS_PER_SECOND_BASE = "Attacks_Per_Second_Item";

   // This describes the attacks per second bonus on a
   // weapon only. E.G. it is 0.11 if there is an 11% attack
   // speed bonus.
   public static final String WEAPON_ATTR_ATTACK_SPEED_PERCENT_BONUS = "Attacks_Per_Second_Item_Percent";

   // These values describe the minimums and deltas of
   // both physical and elemental damage on this weapon.
   // The 'Damage_Weapon_Min#Physical' is the *true*
   // base minimum damage of this weapon. That value is
   // then added to any bonuses found of the same type
   // and then further multiplied by any percentage
   // bonuses found for that same type of damage.
   public static final String WEAPON_ATTR_PHYSICAL_DAMAGE_BASE_MIN = "Damage_Weapon_Min#Physical";
   public static final String WEAPON_ATTR_PHYSICAL_DAMAGE_BASE_DELTA = "Damage_Weapon_Delta#Physical";

   public static final String WEAPON_ATTR_DAMAGE_MIN_PREFIX = "Damage_Weapon_Min#";
   public static final String WEAPON_ATTR_DAMAGE_DELTA_PREFIX = "Damage_Weapon_Delta#";

   public static final String PHYSICAL_DAMAGE_TYPE = "Physical";

   public static final String[] ELEMENTAL_DAMAGE_TYPES =
   {
      "Arcane",
      "Cold",
      "Fire",
      "Holy",
      "Lightning",
      "Poison",
   };

   public static final String WEAPON_ATTR_PHYSICAL_DAMAGE_BONUS_MIN = "Damage_Weapon_Bonus_Min#Physical";
   public static final String WEAPON_ATTR_PHYSICAL_DAMAGE_BONUS_MAX = "Damage_Weapon_Bonus_Max#Physical";

   public static final String WEAPON_ATTR_DAMAGE_PERCENT_BONUS = "Damage_Weapon_Percent_Bonus#Physical";

   public static final String WEAPON_ATTR_ATTACK_SPEED_RAW_BONUS = "Attacks_Per_Second_Item_Bonus";
}
