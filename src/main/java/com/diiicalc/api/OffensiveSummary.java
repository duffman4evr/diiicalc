package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonProperty;

public class OffensiveSummary
{
   @JsonProperty("dps")
   private double dps;

   @JsonProperty("weaponDamage")
   private double weaponDamage;

   @JsonProperty("attacksPerSecond")
   private double attacksPerSecond;

   @JsonProperty("critChance")
   private double critChance;

   @JsonProperty("critDamage")
   private double critDamage;

   @JsonProperty("onePrimaryStatDps")
   private double primaryStatDps;

   @JsonProperty("onePercentCritChanceDps")
   private double critChanceDps;

   @JsonProperty("onePercentCritDamageDps")
   private double critDamageDps;

   @JsonProperty("onePercentAttackSpeedDps")
   private double attackSpeedDps;

   public OffensiveSummary
   (
      double dps,
      double weaponDamage,
      double attacksPerSecond,
      double critChance,
      double critDamage,
      double primaryStatDps,
      double critChanceDps,
      double critDamageDps,
      double attackSpeedDps
   )
   {
      this.dps = dps;
      this.weaponDamage = weaponDamage;
      this.attacksPerSecond = attacksPerSecond;
      this.critChance = critChance;
      this.critDamage = critDamage;
      this.primaryStatDps = primaryStatDps;
      this.critChanceDps = critChanceDps;
      this.critDamageDps = critDamageDps;
      this.attackSpeedDps = attackSpeedDps;
   }
}
