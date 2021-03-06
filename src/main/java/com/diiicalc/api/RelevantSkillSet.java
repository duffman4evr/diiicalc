package com.diiicalc.api;

import com.diiicalc.core.ActiveSkills;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.ArrayList;
import java.util.List;

public class RelevantSkillSet
{
   @JsonProperty("active")
   private List<RelevantActiveSkill> active = new ArrayList<RelevantActiveSkill>();

   @JsonProperty("passive")
   private List<Skill> passive = new ArrayList<Skill>();

   public List<RelevantActiveSkill> getActive()
   {
      return active;
   }

   public List<Skill> getPassive()
   {
      return passive;
   }

   public static class RelevantActiveSkill
   {
      @JsonProperty("skill")
      Skill skill;

      @JsonProperty("runes")
      List<Rune> runes = new ArrayList<Rune>();

      @JsonProperty("allowRunelessUsage")
      boolean allowRunelessUsage = true;

      public RelevantActiveSkill(Skill skill)
      {
         this.skill = skill;
      }

      public List<Rune> getRunes()
      {
         return runes;
      }

      public Skill getSkill()
      {
         return skill;
      }

      public void disallowRunelessUsage()
      {
         this.allowRunelessUsage = false;
      }

      public void addRune(String runeType)
      {
         this.runes.add(ActiveSkills.lookupRune(this.skill.getSlug(), runeType));
      }
   }
}
