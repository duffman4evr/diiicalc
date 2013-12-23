package com.diiicalc.core;

public class OffensiveStats
{
   private double dps;
   private double weaponDamage;
   private double attacksPerSecond;
   private double critChance;
   private double critDamage;

   public OffensiveStats
   (
      double dps,
      double weaponDamage,
      double attacksPerSecond,
      double critChance,
      double critDamage
   )
   {
      this.dps = dps;
      this.weaponDamage = weaponDamage;
      this.attacksPerSecond = attacksPerSecond;
      this.critChance = critChance;
      this.critDamage = critDamage;
   }

   public double getDps()
   {
      return dps;
   }

   public double getWeaponDamage()
   {
      return weaponDamage;
   }

   public double getAttacksPerSecond()
   {
      return attacksPerSecond;
   }

   public double getCritChance()
   {
      return critChance;
   }

   public double getCritDamage()
   {
      return critDamage;
   }
}
