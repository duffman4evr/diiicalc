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
   private final double effectiveLifeWithoutDodge;

   public DefensiveStats
   (
      double armor,
      double resistAll,
      double dodgeChance,
      double averageBlockAmount,
      Map<String, Double> incomingDamageModifiers,
      double incomingDamageModifier,
      double effectiveLife,
      double effectiveLifeWithoutDodge
   )
   {
      this.armor = armor;
      this.resistAll = resistAll;
      this.dodgeChance = dodgeChance;
      this.averageBlockAmount = averageBlockAmount;
      this.incomingDamageModifiers = incomingDamageModifiers;
      this.totalIncomingDamageModifier = incomingDamageModifier;
      this.effectiveLife = effectiveLife;
      this.effectiveLifeWithoutDodge = effectiveLifeWithoutDodge;
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

   public double getEffectiveLifeWithoutDodge()
   {
      return effectiveLifeWithoutDodge;
   }
}
