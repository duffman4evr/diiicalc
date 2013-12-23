package com.diiicalc.core;

import java.util.Map;

public class DefensiveStats
{
   private final double armor;
   private final double resistAll;
   private final double dodgeChance;
   private final double averageBlockAmount;
   private final Map<String, Double> incomingDamageModifiers;
   private final double totalIncomingDamageModifier;
   private final double effectiveLife;

   public DefensiveStats
   (
      double armor,
      double resistAll,
      double dodgeChance,
      double averageBlockAmount,
      Map<String, Double> incomingDamageModifiers,
      double incomingDamageModifier,
      double effectiveLife
   )
   {
      this.armor = armor;
      this.resistAll = resistAll;
      this.dodgeChance = dodgeChance;
      this.averageBlockAmount = averageBlockAmount;
      this.incomingDamageModifiers = incomingDamageModifiers;
      this.totalIncomingDamageModifier = incomingDamageModifier;
      this.effectiveLife = effectiveLife;
   }

   public double getArmor()
   {
      return armor;
   }

   public double getResistAll()
   {
      return resistAll;
   }

   public double getDodgeChance()
   {
      return dodgeChance;
   }

   public double getAverageBlockAmount()
   {
      return averageBlockAmount;
   }

   public Map<String, Double> getIncomingDamageModifiers()
   {
      return incomingDamageModifiers;
   }

   public double getTotalIncomingDamageModifier()
   {
      return totalIncomingDamageModifier;
   }

   public double getEffectiveLife()
   {
      return effectiveLife;
   }
}
