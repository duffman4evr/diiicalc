package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.*;

public class SkillChoiceSet
{
   @JsonProperty("relevantSkillSet")
   RelevantSkillSet relevantSkillSet;

   @JsonProperty("passiveChoices")
   Map<String, String> passiveChoices = new HashMap<String, String>();

   @JsonProperty("activeChoices")
   Map<String, String> activeChoices = new HashMap<String, String>();

   public SkillChoiceSet(RelevantSkillSet relevantSkillSet)
   {
      this.relevantSkillSet = relevantSkillSet;
   }

   public void addSkillChoices(Skills skills)
   {
      // Passive.
      for (Skill relevantPassive : this.relevantSkillSet.getPassive())
      {
         for (Skills.PassiveSkill passive : skills.getPassive())
         {
            if (passive.getSkill().getSlug().equals(relevantPassive.getSlug()))
            {
               passiveChoices.put(relevantPassive.getSlug(), null);
            }
         }
      }

      // Active.
      for (Skills.ActiveSkill activeSkillChoice : skills.getActive())
      {
         for (RelevantSkillSet.RelevantActiveSkill relevantActiveSkill : this.relevantSkillSet.getActive())
         {
            if (!relevantActiveSkill.getSkill().equals(activeSkillChoice.getSkill()))
            {
               continue;
            }

            // If it is a relevant active skill, check if a relevant rune is chosen.
            List<Rune> relevantRunes = relevantActiveSkill.getRunes();

            String chosenRuneSlug = null;

            for (Rune relevantRune : relevantRunes)
            {
               if (activeSkillChoice.getRune().getSlug().equals(relevantRune.getSlug()))
               {
                  chosenRuneSlug = relevantRune.getSlug();
                  break;
               }
            }

            this.activeChoices.put(activeSkillChoice.getSkill().getSlug(), chosenRuneSlug);
         }
      }
   }
}
