package com.diiicalc.core;

import com.diiicalc.api.Rune;
import com.diiicalc.api.Skill;
import com.diiicalc.core.modifiers.ArmorModifier;
import com.diiicalc.core.modifiers.SkillDamageModifier;

import java.util.HashMap;
import java.util.Map;

public class Injectors
{
   public static ModifierInjector lookup(Skill skill, Rune rune)
   {
      String lookupKey = skill.getSlug();

      if (rune != null)
      {
         lookupKey = lookupKey + "|" + rune.getType();
      }

      return INJECTOR_MAP.get(lookupKey);
   }

   private static final Map<String, ModifierInjector> INJECTOR_MAP = new HashMap<String, ModifierInjector>();

   static
   {
      INJECTOR_MAP.put("big-bad-voodoo|a", new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers modifiers)
         {
            modifiers.armor.add(new ArmorModifier()
            {
               @Override
               public double calculateArmor(double armorFromStrength, double armorFromItems)
               {
                  return 0;
               }
            });
         }
      });

      INJECTOR_MAP.put(PassiveSkills.SLUG_WIZARD_GLASS_CANNON, new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers modifiers)
         {
            modifiers.skillDamage.add(new SkillDamageModifier()
            {
               @Override
               public double getSkillDamageBonus()
               {
                  return 0.15;
               }
            });
         }
      });
   }
}
