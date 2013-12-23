package com.diiicalc.core;

import com.diiicalc.api.Gem;
import com.diiicalc.api.Item;
import com.diiicalc.api.SetRank;
import com.diiicalc.api.ValueRange;

import java.util.*;

class StatsFromItems
{
   // Base.
   private double dexterity;
   private double strength;
   private double intelligence;
   private double vitality;

   // Defensive.
   private double lifePercent;
   private double lifeOnHit;
   private double lifeRegen;

   private double blockChance;
   private double blockAmountMin;
   private double blockAmountDelta;

   private double armor;

   private double resistArcane;
   private double resistFire;
   private double resistLightning;
   private double resistPoison;
   private double resistCold;
   private double resistPhysical;

   private double resistAll;

   // Offensive.
   private double critChance;
   private double critDamageBonus;

   private Map<String, ValueRange> damageBonuses = new HashMap<String, ValueRange>();

   private double attackSpeedBonus;

   private Weapon mainWeapon;
   private Weapon offWeapon;

   StatsFromItems(Map<String, Item> itemMap)
   {
      Map<String, List<SetRank>> setDefinition = new HashMap<String, List<SetRank>>();
      Map<String, Integer> setOwnership = new HashMap<String, Integer>();

      List<Map<String, ValueRange>> attributesList = new ArrayList<Map<String, ValueRange>>();

      for (Map.Entry<String, Item> entry : itemMap.entrySet())
      {
         String slotType = entry.getKey();
         Item item = entry.getValue();

         if (Constants.ITEM_SLOT_MAIN_HAND.equals(slotType))
         {
            this.mainWeapon = new Weapon(item);
         }

         if (Constants.ITEM_SLOT_OFF_HAND.equals(slotType))
         {
            if (Utils.isWeapon(item.getType().getId()))
            {
               this.offWeapon = new Weapon(item);
            }
         }

         // Armor is quirky, deal with it here.
         item.getAttributesRaw().remove(Constants.ITEM_ATTR_ARMOR);
         item.getAttributesRaw().remove(Constants.ITEM_ATTR_ARMOR_BONUS);

         item.getAttributesRaw().put(Constants.ITEM_ATTR_ARMOR, item.getArmor());

         // Base attributes.
         attributesList.add(item.getAttributesRaw());

         // Gems.
         if (item.getGems() != null)
         {
            for (Gem gem : item.getGems())
            {
               attributesList.add(gem.getAttributesRaw());
            }
         }

         // Sets.
         if (item.getSet() != null)
         {
            String slug = item.getSet().getSlug();

            if (!setDefinition.containsKey(slug))
            {
               setDefinition.put(slug, item.getSet().getRanks());
            }

            Integer ownershipCount = setOwnership.get(slug);

            if (ownershipCount == null)
            {
               ownershipCount = 0;
            }

            ownershipCount++;

            setOwnership.put(slug, ownershipCount);
         }
      }

      // Do Set post-processing.
      for (Map.Entry<String, Integer> entry : setOwnership.entrySet())
      {
         String slug = entry.getKey();
         int ownership = entry.getValue();

         List<SetRank> setRanks = setDefinition.get(slug);

         for (SetRank setRank : setRanks)
         {
            if (setRank.getRequiredOwnership() > ownership)
            {
               continue;
            }

            attributesList.add(setRank.getAttributesRaw());
         }
      }

      // Process the attribute maps.
      for (Map<String, ValueRange> attributesRaw : attributesList)
      {
         this.addAttributes(attributesRaw);
      }
   }

   private void addAttributes(Map<String, ValueRange> attributes)
   {
      // Base stats.
      {
         ValueRange dexterity = attributes.get(Constants.ITEM_ATTR_DEXTERITY);
         ValueRange strength = attributes.get(Constants.ITEM_ATTR_STRENGTH);
         ValueRange intelligence = attributes.get(Constants.ITEM_ATTR_INTELLIGENCE);
         ValueRange vitality = attributes.get(Constants.ITEM_ATTR_VITALITY);

         if (dexterity != null) this.dexterity += dexterity.getMin();
         if (strength != null) this.strength += strength.getMin();
         if (intelligence != null) this.intelligence += intelligence.getMin();
         if (vitality != null) this.vitality += vitality.getMin();
      }

      // Base stats (try to get bugged names).
      {
         ValueRange dexterity = attributes.get(Constants.BUG_ITEM_ATTR_DEXTERITY);
         ValueRange strength = attributes.get(Constants.BUG_ITEM_ATTR_STRENGTH);
         ValueRange intelligence = attributes.get(Constants.BUG_ITEM_ATTR_INTELLIGENCE);
         ValueRange vitality = attributes.get(Constants.BUG_ITEM_ATTR_VITALITY);

         if (dexterity != null) this.dexterity += dexterity.getMin();
         if (strength != null) this.strength += strength.getMin();
         if (intelligence != null) this.intelligence += intelligence.getMin();
         if (vitality != null) this.vitality += vitality.getMin();
      }

      // Direct life-related.
      {
         ValueRange lifePercent = attributes.get(Constants.ITEM_ATTR_LIFE_PERCENT);
         ValueRange lifeOnHit = attributes.get(Constants.ITEM_ATTR_LIFE_ON_HIT);
         ValueRange lifeRegen = attributes.get(Constants.ITEM_ATTR_LIFE_REGEN);

         if (lifePercent != null) this.lifePercent += lifePercent.getMin();
         if (lifeOnHit != null) this.lifeOnHit += lifeOnHit.getMin();
         if (lifeRegen != null) this.lifeRegen += lifeRegen.getMin();
      }

      // Shield.
      {
         ValueRange blockChance = attributes.get(Constants.ITEM_ATTR_BLOCK_CHANCE);
         ValueRange blockChanceBonus = attributes.get(Constants.ITEM_ATTR_BLOCK_CHANCE_BONUS);
         ValueRange blockAmountMin = attributes.get(Constants.ITEM_ATTR_BLOCK_AMOUNT_MIN);
         ValueRange blockAmountDelta = attributes.get(Constants.ITEM_ATTR_BLOCK_AMOUNT_DELTA);

         if (blockChance != null) this.blockChance += blockChance.getMin();
         if (blockChanceBonus != null) this.blockChance += blockChanceBonus.getMin();
         if (blockAmountMin != null) this.blockAmountMin += blockAmountMin.getMin();
         if (blockAmountDelta != null) this.blockAmountDelta += blockAmountDelta.getMin();
      }

      // Armor.
      {
         ValueRange armor = attributes.get(Constants.ITEM_ATTR_ARMOR);
         ValueRange armorBonus = attributes.get(Constants.ITEM_ATTR_ARMOR_BONUS);

         if (armor != null) this.armor += armor.getMin();
         if (armorBonus != null) this.armor += armorBonus.getMin();
      }

      // Resists.
      {
         ValueRange resistAll = attributes.get(Constants.ITEM_ATTR_RESIST_ALL);
         ValueRange resistArcane = attributes.get(Constants.ITEM_ATTR_RESIST_ARCANE);
         ValueRange resistFire = attributes.get(Constants.ITEM_ATTR_RESIST_FIRE);
         ValueRange resistLightning = attributes.get(Constants.ITEM_ATTR_RESIST_LIGHTNING);
         ValueRange resistPoison = attributes.get(Constants.ITEM_ATTR_RESIST_POISON);
         ValueRange resistCold = attributes.get(Constants.ITEM_ATTR_RESIST_COLD);
         ValueRange resistPhysical = attributes.get(Constants.ITEM_ATTR_RESIST_PHYSICAL);

         if (resistAll != null) this.resistAll += resistAll.getMin();
         if (resistArcane != null) this.resistArcane += resistArcane.getMin();
         if (resistFire != null) this.resistFire += resistFire.getMin();
         if (resistLightning != null) this.resistLightning += resistLightning.getMin();
         if (resistPoison != null) this.resistPoison += resistPoison.getMin();
         if (resistCold != null) this.resistCold += resistCold.getMin();
         if (resistPhysical != null) this.resistPhysical += resistPhysical.getMin();
      }

      // Bonus damage.
      {
         List<String> damageTypes = new ArrayList<String>();

         damageTypes.add(Constants.PHYSICAL_DAMAGE_TYPE);
         damageTypes.addAll(Arrays.asList(Constants.ELEMENTAL_DAMAGE_TYPES));

         for (String damageType : damageTypes)
         {
            ValueRange min = attributes.get(Constants.ITEM_ATTR_DAMAGE_MIN_PREFIX + damageType);
            ValueRange delta = attributes.get(Constants.ITEM_ATTR_DAMAGE_DELTA_PREFIX + damageType);

            if (min != null && delta != null)
            {
               this.damageBonuses.put(damageType, ValueRange.fromMinAndDelta(min, delta));
            }
         }
      }

      // Crit chance.
      {
         ValueRange critChance = attributes.get(Constants.ITEM_ATTR_CRIT_CHANCE);

         if (critChance != null) this.critChance += critChance.getMin();
      }

      // Crit damage.
      {
         ValueRange critDamageBonus = attributes.get(Constants.ITEM_ATTR_CRIT_DMG);

         if (critDamageBonus != null) this.critDamageBonus += critDamageBonus.getMin();
      }

      // Attack speed bonus.
      {
         ValueRange attackSpeedBonus = attributes.get(Constants.ITEM_ATTR_ATTACK_SPEED_PERCENT);

         if (attackSpeedBonus != null) this.attackSpeedBonus += attackSpeedBonus.getMin();
      }
   }

   public double getDexterity()
   {
      return dexterity;
   }

   public void setDexterity(double dexterity)
   {
      this.dexterity = dexterity;
   }

   public double getStrength()
   {
      return strength;
   }

   public void setStrength(double strength)
   {
      this.strength = strength;
   }

   public double getIntelligence()
   {
      return intelligence;
   }

   public void setIntelligence(double intelligence)
   {
      this.intelligence = intelligence;
   }

   public double getVitality()
   {
      return vitality;
   }

   public void setVitality(double vitality)
   {
      this.vitality = vitality;
   }

   public double getLifePercent()
   {
      return lifePercent;
   }

   public void setLifePercent(double lifePercent)
   {
      this.lifePercent = lifePercent;
   }

   public double getLifeOnHit()
   {
      return lifeOnHit;
   }

   public void setLifeOnHit(double lifeOnHit)
   {
      this.lifeOnHit = lifeOnHit;
   }

   public double getLifeRegen()
   {
      return lifeRegen;
   }

   public void setLifeRegen(double lifeRegen)
   {
      this.lifeRegen = lifeRegen;
   }

   public double getBlockChance()
   {
      return blockChance;
   }

   public void setBlockChance(double blockChance)
   {
      this.blockChance = blockChance;
   }

   public double getBlockAmountMin()
   {
      return blockAmountMin;
   }

   public void setBlockAmountMin(double blockAmountMin)
   {
      this.blockAmountMin = blockAmountMin;
   }

   public double getBlockAmountDelta()
   {
      return blockAmountDelta;
   }

   public void setBlockAmountDelta(double blockAmountDelta)
   {
      this.blockAmountDelta = blockAmountDelta;
   }

   public double getArmor()
   {
      return armor;
   }

   public void setArmor(double armor)
   {
      this.armor = armor;
   }

   public double getResistArcane()
   {
      return resistArcane;
   }

   public void setResistArcane(double resistArcane)
   {
      this.resistArcane = resistArcane;
   }

   public double getResistFire()
   {
      return resistFire;
   }

   public void setResistFire(double resistFire)
   {
      this.resistFire = resistFire;
   }

   public double getResistLightning()
   {
      return resistLightning;
   }

   public void setResistLightning(double resistLightning)
   {
      this.resistLightning = resistLightning;
   }

   public double getResistPoison()
   {
      return resistPoison;
   }

   public void setResistPoison(double resistPoison)
   {
      this.resistPoison = resistPoison;
   }

   public double getResistCold()
   {
      return resistCold;
   }

   public void setResistCold(double resistCold)
   {
      this.resistCold = resistCold;
   }

   public double getResistPhysical()
   {
      return resistPhysical;
   }

   public void setResistPhysical(double resistPhysical)
   {
      this.resistPhysical = resistPhysical;
   }

   public double getResistAll()
   {
      return resistAll;
   }

   public void setResistAll(double resistAll)
   {
      this.resistAll = resistAll;
   }

   public double getCritChance()
   {
      return critChance;
   }

   public void setCritChance(double critChance)
   {
      this.critChance = critChance;
   }

   public double getCritDamageBonus()
   {
      return critDamageBonus;
   }

   public void setCritDamageBonus(double critDamageBonus)
   {
      this.critDamageBonus = critDamageBonus;
   }

   public Map<String, ValueRange> getDamageBonuses()
   {
      return damageBonuses;
   }

   public double getAttackSpeedBonus()
   {
      return attackSpeedBonus;
   }

   public void setAttackSpeedBonus(double attackSpeedBonus)
   {
      this.attackSpeedBonus = attackSpeedBonus;
   }

   public Weapon getMainWeapon()
   {
      return mainWeapon;
   }

   public void setMainWeapon(Weapon mainWeapon)
   {
      this.mainWeapon = mainWeapon;
   }

   public Weapon getOffWeapon()
   {
      return offWeapon;
   }

   public void setOffWeapon(Weapon offWeapon)
   {
      this.offWeapon = offWeapon;
   }

   public static class Weapon
   {
      // To determine average damage.
      private ValueRange baseDamageRange;
      private ValueRange physicalDamageBonusRange = new ValueRange();
      private double damagePercentBonus;
      private Map<String, ValueRange> elementalDamageBonusRanges = new HashMap<String, ValueRange>();

      // To determine attacks per second.
      private Item.Type type;
      private double attackSpeedPercentBonus;
      private double attackSpeedRawBonus;

      public Weapon() { }

      private Weapon(Item item)
      {
         Map<String, ValueRange> attributes = item.getAttributesRaw();

         // Base damage.
         {
            ValueRange min = attributes.get(Constants.WEAPON_ATTR_PHYSICAL_DAMAGE_BASE_MIN);
            ValueRange delta = attributes.get(Constants.WEAPON_ATTR_PHYSICAL_DAMAGE_BASE_DELTA);

            if (min != null && delta != null)
            {
               this.baseDamageRange = ValueRange.fromMinAndDelta(min, delta);
            }
         }

         // Bonus physical damage.
         {
            ValueRange min = attributes.get(Constants.WEAPON_ATTR_PHYSICAL_DAMAGE_BONUS_MIN);
            ValueRange max = attributes.get(Constants.WEAPON_ATTR_PHYSICAL_DAMAGE_BONUS_MAX);

            if (min != null && max != null)
            {
               this.physicalDamageBonusRange = ValueRange.fromMinAndMax(min, max);
            }
         }

         // Physical % bonus.
         {
            ValueRange damagePercentBonus = attributes.get(Constants.WEAPON_ATTR_DAMAGE_PERCENT_BONUS);

            if (damagePercentBonus != null) this.damagePercentBonus = damagePercentBonus.getMin();
         }

         // Elemental damage (we account for more than one type even though the game doesn't allow it).
         for (String damageType : Constants.ELEMENTAL_DAMAGE_TYPES)
         {
            ValueRange min = attributes.get(Constants.WEAPON_ATTR_DAMAGE_MIN_PREFIX + damageType);
            ValueRange delta = attributes.get(Constants.WEAPON_ATTR_DAMAGE_DELTA_PREFIX + damageType);

            if (min != null && delta != null)
            {
               elementalDamageBonusRanges.put(damageType, ValueRange.fromMinAndDelta(min, delta));
            }
         }

         // Type (determines base attack speed).
         this.type = item.getType();

         // % bonus to attack speed.
         {
            ValueRange attackSpeedPercentBonus = attributes.get(Constants.WEAPON_ATTR_ATTACK_SPEED_PERCENT_BONUS);

            if (attackSpeedPercentBonus != null) this.attackSpeedPercentBonus = attackSpeedPercentBonus.getMin();
         }

         // Raw bonus to attack speed.
         {
            ValueRange attackSpeedRawBonus = attributes.get(Constants.WEAPON_ATTR_ATTACK_SPEED_RAW_BONUS);

            if (attackSpeedRawBonus != null) this.attackSpeedRawBonus = attackSpeedRawBonus.getMin();
         }
      }

      public ValueRange getBaseDamageRange()
      {
         return baseDamageRange;
      }

      public void setBaseDamageRange(ValueRange baseDamageRange)
      {
         this.baseDamageRange = baseDamageRange;
      }

      public ValueRange getPhysicalDamageBonusRange()
      {
         return physicalDamageBonusRange;
      }

      public void setPhysicalDamageBonusRange(ValueRange physicalDamageBonusRange)
      {
         this.physicalDamageBonusRange = physicalDamageBonusRange;
      }

      public double getDamagePercentBonus()
      {
         return damagePercentBonus;
      }

      public void setDamagePercentBonus(double damagePercentBonus)
      {
         this.damagePercentBonus = damagePercentBonus;
      }

      public Map<String, ValueRange> getElementalDamageBonusRanges()
      {
         return elementalDamageBonusRanges;
      }

      public Item.Type getType()
      {
         return type;
      }

      public void setType(Item.Type type)
      {
         this.type = type;
      }

      public double getAttackSpeedPercentBonus()
      {
         return attackSpeedPercentBonus;
      }

      public void setAttackSpeedPercentBonus(double attackSpeedPercentBonus)
      {
         this.attackSpeedPercentBonus = attackSpeedPercentBonus;
      }

      public double getAttackSpeedRawBonus()
      {
         return attackSpeedRawBonus;
      }

      public void setAttackSpeedRawBonus(double attackSpeedRawBonus)
      {
         this.attackSpeedRawBonus = attackSpeedRawBonus;
      }
   }
}
