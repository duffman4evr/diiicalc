package com.diiicalc.core;

import com.diiicalc.api.Rune;
import com.diiicalc.api.Skill;
import com.diiicalc.api.Skills;
import com.diiicalc.core.modifiers.*;

import java.util.ArrayList;
import java.util.List;

public class StatModifiers
{
   List<ArmorModifier> armor = new ArrayList<ArmorModifier>();
   List<IncomingDamageModifier> incomingDamage = new ArrayList<IncomingDamageModifier>();
   List<CritChanceModifier> critChance = new ArrayList<CritChanceModifier>();
   List<CritDamageModifier> critDamage = new ArrayList<CritDamageModifier>();
   List<SkillDamageModifier> skillDamage = new ArrayList<SkillDamageModifier>();

   public StatModifiers(Skills skills)
   {
      for (Skills.ActiveSkill activeSkill : skills.getActive())
      {
         Skill skill = activeSkill.getSkill();
         Rune rune = activeSkill.getRune();

         ModifierInjector modifierInjector = Utils.lookupInjector(skill, rune);

         if (modifierInjector == null)
         {
            continue;
         }

         modifierInjector.inject(this);
      }
   }
}
