package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonProperty;

public class OffensiveSummary
{
   @JsonProperty("dps")
   private double dps;

   @JsonProperty("attacks-per-second")
   private double attacksPerSecond;

   @JsonProperty("crit-chance")
   private double critChance;

   @JsonProperty("crit-damage")
   private double critDamage;

   @JsonProperty("one-primary-stat-dps")
   private double primaryStatDps;

   @JsonProperty("one-percent-crit-chance-dps")
   private double critChanceDps;

   @JsonProperty("one-percent-crit-damage-dps")
   private double critDamageDps;

   @JsonProperty("one-percent-attack-speed-dps")
   private double attackSpeedDps;

   public OffensiveSummary
   (
      double dps,
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
      this.attacksPerSecond = attacksPerSecond;
      this.critChance = critChance;
      this.critDamage = critDamage;
      this.primaryStatDps = primaryStatDps;
      this.critChanceDps = critChanceDps;
      this.critDamageDps = critDamageDps;
      this.attackSpeedDps = attackSpeedDps;
   }
}
