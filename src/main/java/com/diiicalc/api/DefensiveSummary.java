package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonProperty;

public class DefensiveSummary
{
   @JsonProperty("effectiveLife")
   private double effectiveLife;

   @JsonProperty("effectiveLifeWithoutDodge")
   private double effectiveLifeWithoutDodge;

   @JsonProperty("armor")
   private double armor;

   @JsonProperty("allResist")
   private double allResist;

   @JsonProperty("dodgeChance")
   private double dodgeChance;

   @JsonProperty("totalIncomingDamageModifier")
   private double totalIncomingDamageModifier;

   @JsonProperty("resistAllDamageModifier")
   private double resistAllDamageModifier;

   @JsonProperty("armorDamageModifier")
   private double armorDamageModifier;

   @JsonProperty("oneArmorEhp")
   private double oneArmorEhp;

   @JsonProperty("oneAllResistEhp")
   private double oneAllResistEhp;

   @JsonProperty("oneVitalityEhp")
   private double oneVitalityEhp;

   @JsonProperty("onePercentLifeEhp")
   private double onePercentLifeEhp;

   public DefensiveSummary
   (
      double effectiveLife,
      double effectiveLifeWithoutDodge,
      double armor,
      double allResist,
      double dodgeChance,
      double totalIncomingDamageModifier,
      double resistAllDamageModifier,
      double armorDamageModifier,
      double oneArmorEhp,
      double oneAllResistEhp,
      double oneVitalityEhp,
      double onePercentLifeEhp
   )
   {
      this.effectiveLife = effectiveLife;
      this.effectiveLifeWithoutDodge = effectiveLifeWithoutDodge;
      this.armor = armor;
      this.allResist = allResist;
      this.dodgeChance = dodgeChance;
      this.totalIncomingDamageModifier = totalIncomingDamageModifier;
      this.resistAllDamageModifier = resistAllDamageModifier;
      this.armorDamageModifier = armorDamageModifier;
      this.oneArmorEhp = oneArmorEhp;
      this.oneAllResistEhp = oneAllResistEhp;
      this.oneVitalityEhp = oneVitalityEhp;
      this.onePercentLifeEhp = onePercentLifeEhp;
   }
}
