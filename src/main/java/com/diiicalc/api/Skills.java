package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.ArrayList;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown=true)
public class Skills
{
   @JsonProperty("active")
   private List<ActiveSkill> active = new ArrayList<ActiveSkill>();

   @JsonProperty("passive")
   private List<PassiveSkill> passive = new ArrayList<PassiveSkill>();

   public List<ActiveSkill> getActive()
   {
      return active;
   }

   public List<PassiveSkill> getPassive()
   {
      return passive;
   }

   public static class ActiveSkill
   {
      @JsonProperty("skill")
      private Skill skill;

      @JsonProperty("rune")
      private Rune rune;

      public Skill getSkill()
      {
         return skill;
      }

      public Rune getRune()
      {
         return rune;
      }
   }

   public static class PassiveSkill
   {
      @JsonProperty("skill")
      private Skill active;

      public Skill getActive()
      {
         return active;
      }
   }
}
