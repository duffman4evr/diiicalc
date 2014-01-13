package com.diiicalc.core;

import com.diiicalc.api.Hero;
import com.diiicalc.api.Item;
import com.diiicalc.api.ValueRange;
import com.diiicalc.core.modifiers.*;

import java.util.HashMap;
import java.util.Map;

public class StatTotals
{
   private String heroClass;
   private long heroLevel;
   private long heroParagonLevel;
   private StatsFromItems statsFromItems;
   private StatModifiers statModifiers;

   public StatTotals(Map<String, Item> itemMap, Hero hero, Map<String, String> activeSkillsOverride, Map<String, String> passiveSkillsOverride)
   {
      this.heroClass = hero.getType();
      this.heroLevel = hero.getLevel();
      this.heroParagonLevel = hero.getParagonLevel();

      this.statsFromItems = new StatsFromItems(itemMap);
      this.statModifiers = new StatModifiers(hero.getSkills(), activeSkillsOverride, passiveSkillsOverride);
   }

   // -----
   // Getters.
   // -----

   public String getHeroClass()
   {
      return heroClass;
   }

   public long getHeroLevel()
   {
      return heroLevel;
   }

   public long getHeroParagonLevel()
   {
      return heroParagonLevel;
   }

   public double getDexterity()
   {
      double base = Utils.computeBaseDexterity(this.getHeroLevel(), this.getHeroParagonLevel(), this.getHeroClass());
      double fromItems = this.statsFromItems.getDexterity();
      return base + fromItems;
   }

   public double getStrength()
   {
      double base = Utils.computeBaseStrength(this.getHeroLevel(), this.getHeroParagonLevel(), this.getHeroClass());
      double fromItems = this.statsFromItems.getStrength();
      return base + fromItems;
   }

   public double getIntelligence()
   {
      double base = Utils.computeBaseIntelligence(this.getHeroLevel(), this.getHeroParagonLevel(), this.getHeroClass());
      double fromItems = this.statsFromItems.getIntelligence();
      return base + fromItems;
   }

   public double getVitality()
   {
      double base = Utils.computeBaseVitality(this.getHeroLevel(), this.getHeroParagonLevel());
      double fromItems = this.statsFromItems.getVitality();
      return base + fromItems;
   }

   public double getMainStat()
   {
      if (Constants.HERO_TYPE_BARBARIAN.equals(this.getHeroClass()))
      {
         return this.getStrength();
      }
      else if (Constants.HERO_TYPE_WIZARD.equals(this.getHeroClass()) || Constants.HERO_TYPE_WITCH_DOCTOR.equals(this.getHeroClass()))
      {
         return this.getIntelligence();
      }
      else
      {
         return this.getDexterity();
      }
   }

   public double getBlockChance()
   {
      return this.statsFromItems.getBlockChance();
   }

   public double getBlockAmountMin()
   {
      return this.statsFromItems.getBlockAmountMin();
   }

   public double getBlockAmountDelta()
   {
      return this.statsFromItems.getBlockAmountDelta();
   }

   public double getAllResist()
   {
      double fromIntelligence = 0.1 * this.getIntelligence();
      double fromAllResistOnItems = this.statsFromItems.getResistAll();
      double fromIndividualResistOnItems = Utils.findMin
      (
         this.statsFromItems.getResistArcane(),
         this.statsFromItems.getResistFire(),
         this.statsFromItems.getResistLightning(),
         this.statsFromItems.getResistPoison(),
         this.statsFromItems.getResistCold(),
         this.statsFromItems.getResistPhysical()
      );

      return fromIntelligence + fromAllResistOnItems + fromIndividualResistOnItems;
   }

   public double getArmor()
   {
      double fromStrength = this.getStrength();
      double fromItems = this.statsFromItems.getArmor();
      double fromModifiers = 0;

      for (ArmorModifier armorModifier : this.statModifiers.armor)
      {
         fromModifiers += armorModifier.calculateArmor(fromStrength, fromItems);
      }

      return fromStrength + fromItems + fromModifiers;
   }

   public double getLife()
   {
      double lifePerVitality = Utils.computeLifePerVitality(this.getHeroLevel());

      double base = 36;
      double fromLevels = 4 * this.getHeroLevel();
      double fromVitality = lifePerVitality * this.getVitality();

      double beforePercent = base + fromLevels + fromVitality;

      double afterPercent = beforePercent * (1 + this.statsFromItems.getLifePercent());

      return afterPercent;
   }

   public double getDodgeChance()
   {
      double fromDexterity = Utils.computeDodgeChanceFromDexterity(this.getDexterity());

      return fromDexterity;
   }

   public double getAverageBlockAmount()
   {
      return this.getBlockChance() * (this.getBlockAmountMin() + (0.5 * this.getBlockAmountDelta()));
   }

   public Map<String, Double> getIncomingDamageModifiers(long monsterLevel)
   {
      Map<String, Double> incomingDamageModifiers = new HashMap<String, Double>();

      double armor = 1.0 - (this.getArmor() / ((50.0 * monsterLevel) + this.getArmor()));
      double allResist = 1.0 - (this.getAllResist() / ((5.0 * monsterLevel) + this.getAllResist()));

      incomingDamageModifiers.put(Constants.DAMAGE_MODIFIER_ARMOR, armor);
      incomingDamageModifiers.put(Constants.DAMAGE_MODIFIER_RESISTS, allResist);

      if (Constants.HERO_TYPE_BARBARIAN.equals(this.getHeroClass()) || Constants.HERO_TYPE_MONK.equals(this.getHeroClass()))
      {
         incomingDamageModifiers.put(Constants.DAMAGE_MODIFIER_MELEE, 0.30);
      }

      for (IncomingDamageModifier defenseModifier : this.statModifiers.incomingDamage)
      {
         defenseModifier.addModifier(incomingDamageModifiers);
      }

      return incomingDamageModifiers;
   }

   public double getAverageBonusDamage()
   {
      double fromItems = 0;

      for (ValueRange damageBonus :this.statsFromItems.getDamageBonuses().values())
      {
         fromItems += damageBonus.getAverage();
      }

      return fromItems;
   }

   public double getCritChance()
   {
      double base = 0.05;
      double fromItems = this.statsFromItems.getCritChance();
      double fromModifiers = 0;

      for (CritChanceModifier critChanceModifier : this.statModifiers.critChance)
      {
         fromModifiers += critChanceModifier.addCritChance(fromItems);
      }

      return base + fromItems + fromModifiers;
   }

   public double getCritDamageBonus()
   {
      double base = 0.50;
      double fromItems = this.statsFromItems.getCritDamageBonus();
      double fromModifiers = 0;

      for (CritDamageModifier critDamageModifier : this.statModifiers.critDamage)
      {
         fromModifiers += critDamageModifier.addCritDamage(fromItems);
      }

      return base + fromItems + fromModifiers;
   }

   public double getAttackSpeedBonus()
   {
      double fromItems = this.statsFromItems.getAttackSpeedBonus();
      double fromDualWield = (this.getWeaponSetup() == WeaponSetup.DUAL_WIELD) ? 0.15 : 0.0;

      return fromItems + fromDualWield;
   }

   public double getAverageWeaponDamage()
   {
      double averageMainWeaponDamage = this.getAverageDamageForWeapon(this.statsFromItems.getMainWeapon());

      WeaponSetup weaponSetup = this.getWeaponSetup();

      if (weaponSetup != WeaponSetup.DUAL_WIELD)
      {
         return averageMainWeaponDamage;
      }

      double averageOffWeaponDamage = this.getAverageDamageForWeapon(this.statsFromItems.getOffWeapon());

      return new ValueRange(averageMainWeaponDamage, averageOffWeaponDamage).getAverage();
   }

   public double getWeaponAttackSpeed()
   {
      double mainWeaponAttackSpeed = this.getAttackSpeedForWeapon(this.statsFromItems.getMainWeapon());

      WeaponSetup weaponSetup = this.getWeaponSetup();

      if (weaponSetup != WeaponSetup.DUAL_WIELD)
      {
         return mainWeaponAttackSpeed;
      }

      double offWeaponAttackSpeed = this.getAverageDamageForWeapon(this.statsFromItems.getOffWeapon());

      return new ValueRange(mainWeaponAttackSpeed, offWeaponAttackSpeed).getAverage();
   }

   public double getSkillDamageBonus()
   {
      double skillDamageBonus = 0;

      for (SkillDamageModifier skillDamageModifier : this.statModifiers.skillDamage)
      {
         skillDamageBonus += skillDamageModifier.getSkillDamageBonus();
      }

      return skillDamageBonus;
   }

   // ----
   // Modifiers.
   // ----

   public void addResistAll(double amountToAdd)
   {
      this.statsFromItems.setResistAll(this.statsFromItems.getResistAll() + amountToAdd);
   }

   public void addArmor(double amountToAdd)
   {
      this.statsFromItems.setArmor(this.statsFromItems.getArmor() + amountToAdd);
   }

   public void addVitality(double amountToAdd)
   {
      this.statsFromItems.setVitality(this.statsFromItems.getVitality() + amountToAdd);
   }

   public void addPercentLife(double amountToAdd)
   {
      this.statsFromItems.setLifePercent(this.statsFromItems.getLifePercent() + amountToAdd);
   }

   public void addPrimaryStat(double amountToAdd)
   {
      if (Constants.HERO_TYPE_BARBARIAN.equals(this.getHeroClass()))
      {
         this.statsFromItems.setStrength(this.statsFromItems.getStrength() + amountToAdd);
      }
      else if (Constants.HERO_TYPE_WIZARD.equals(this.getHeroClass()) || Constants.HERO_TYPE_WITCH_DOCTOR.equals(this.getHeroClass()))
      {
         this.statsFromItems.setIntelligence(this.statsFromItems.getIntelligence() + amountToAdd);
      }
      else
      {
         this.statsFromItems.setDexterity(this.statsFromItems.getDexterity() + amountToAdd);
      }
   }

   public void addCritChance(double amountToAdd)
   {
      this.statsFromItems.setCritChance(this.statsFromItems.getCritChance() + amountToAdd);
   }

   public void addCritDamage(double amountToAdd)
   {
      this.statsFromItems.setCritDamageBonus(this.statsFromItems.getCritDamageBonus() + amountToAdd);
   }

   public void addAttackSpeed(double amountToAdd)
   {
      this.statsFromItems.setAttackSpeedBonus(this.statsFromItems.getAttackSpeedBonus() + amountToAdd);
   }

   // -----
   // Private.
   // -----

   private double getAverageDamageForWeapon(StatsFromItems.Weapon weapon)
   {
      ValueRange damageRange = weapon.getBaseDamageRange();
      damageRange = damageRange.add(weapon.getPhysicalDamageBonusRange());
      damageRange = damageRange.scale(1 + weapon.getDamagePercentBonus());

      for (ValueRange elementalDamageRange : weapon.getElementalDamageBonusRanges().values())
      {
         damageRange = damageRange.add(elementalDamageRange);
      }

      return damageRange.getAverage();
   }

   private double getAttackSpeedForWeapon(StatsFromItems.Weapon weapon)
   {
      double attackSpeed = Utils.getAttackSpeedForWeaponType(weapon.getType());

      attackSpeed *= 1 + weapon.getAttackSpeedPercentBonus();
      attackSpeed += weapon.getAttackSpeedRawBonus();

      return attackSpeed;
   }

   private WeaponSetup getWeaponSetup()
   {
      if (this.statsFromItems.getMainWeapon().getType().isTwoHanded())
      {
         return WeaponSetup.TWO_HANDER;
      }

      if (this.statsFromItems.getOffWeapon() != null)
      {
         return WeaponSetup.DUAL_WIELD;
      }

      return WeaponSetup.ONE_HANDER;
   }
}