package com.diiicalc.core;

import com.diiicalc.api.Rune;
import com.diiicalc.api.Skill;
import com.diiicalc.api.Skills;
import com.diiicalc.core.modifiers.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public class StatModifiers
{
   List<ArmorModifier> armor = new ArrayList<ArmorModifier>();
   List<IncomingDamageModifier> incomingDamage = new ArrayList<IncomingDamageModifier>();
   List<CritChanceModifier> critChance = new ArrayList<CritChanceModifier>();
   List<CritDamageModifier> critDamage = new ArrayList<CritDamageModifier>();
   List<SkillDamageModifier> skillDamage = new ArrayList<SkillDamageModifier>();

   public StatModifiers(Skills skills, Map<String, String> activeSkillsOverride, Map<String, String> passiveSkillsOverride)
   {
      if (activeSkillsOverride != null)
      {
         for (Map.Entry<String, String> entry : activeSkillsOverride.entrySet())
         {
            Skill skill = ActiveSkills.lookup(entry.getKey());
            Rune rune = ActiveSkills.lookupRune(skill.getSlug(), entry.getValue());

            this.injectSkill(skill, rune);
         }
      }
      else
      {
         for (Skills.ActiveSkill activeSkill : skills.getActive())
         {
            Skill skill = activeSkill.getSkill();
            Rune rune = activeSkill.getRune();

            this.injectSkill(skill, rune);
         }
      }

      if (passiveSkillsOverride != null)
      {
         for (Map.Entry<String, String> entry : passiveSkillsOverride.entrySet())
         {
            Skill skill = PassiveSkills.lookup(entry.getKey());

            this.injectSkill(skill, null);
         }
      }
      else
      {
         for (Skills.PassiveSkill activeSkill : skills.getPassive())
         {
            Skill skill = activeSkill.getSkill();

            this.injectSkill(skill, null);
         }
      }
   }

   private void injectSkill(Skill skill, Rune rune)
   {
      ModifierInjector modifierInjector = Injectors.lookup(skill, rune);

      if (modifierInjector == null)
      {
         return;
      }

      modifierInjector.inject(this);
   }
}
