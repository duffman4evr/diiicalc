package com.diiicalc.core;

import com.diiicalc.api.Rune;
import com.diiicalc.api.Skill;
import com.diiicalc.core.modifiers.*;

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
      // ----
      // Wizard
      // ----

      // Active

      INJECTOR_MAP.put(ActiveSkills.SLUG_WIZARD_MAGIC_WEAPON, new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.skillDamage.add(new SkillDamageModifier()
            {
               @Override
               public double get()
               {
                  return 0.1;
               }
            });
         }
      });

      INJECTOR_MAP.put(ActiveSkills.SLUG_WIZARD_MAGIC_WEAPON + "|c", new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.skillDamage.add(new SkillDamageModifier()
            {
               @Override
               public double get()
               {
                  return 0.20;
               }
            });
         }
      });

      INJECTOR_MAP.put(ActiveSkills.SLUG_WIZARD_FROST_NOVA + "|e", new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.critChance.add(new CritChanceModifier()
            {
               @Override
               public double get()
               {
                  return 0.10;
               }
            });
         }
      });

      INJECTOR_MAP.put(ActiveSkills.SLUG_WIZARD_FROST_NOVA + "|a", new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.skillDamage.add(new SkillDamageModifier()
            {
               @Override
               public double get()
               {
                  return 0.60;
               }
            });
         }
      });

      INJECTOR_MAP.put(ActiveSkills.SLUG_WIZARD_SLOW_TIME + "|a", new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.skillDamage.add(new SkillDamageModifier()
            {
               @Override
               public double get()
               {
                  return 0.60;
               }
            });
         }
      });

      // Passive

      INJECTOR_MAP.put(PassiveSkills.SLUG_WIZARD_BLUR, new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.incomingDamage.add(new IncomingDamageModifier()
            {
               @Override
               public void addModifier(Map<String, Double> defenseModifiers)
               {
                  defenseModifiers.put("Blur", 0.80);
               }
            });
         }
      });

      INJECTOR_MAP.put(PassiveSkills.SLUG_WIZARD_GLASS_CANNON, new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.skillDamage.add(new SkillDamageModifier()
            {
               @Override
               public double get()
               {
                  return 0.15;
               }
            });

            statModifiers.allResist.add(new AllResistModifier()
            {
               @Override
               public double get(double totalAllResist)
               {
                  return totalAllResist * -0.1;
               }
            });

            statModifiers.armor.add(new ArmorModifier()
            {
               @Override
               public double get(double totalArmor)
               {
                  return totalArmor * -0.1;
               }
            });
         }
      });

      INJECTOR_MAP.put(PassiveSkills.SLUG_WIZARD_COLD_BLOODED, new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.skillDamage.add(new SkillDamageModifier()
            {
               @Override
               public double get()
               {
                  return 0.1;
               }
            });
         }
      });

      INJECTOR_MAP.put(PassiveSkills.SLUG_WIZARD_CONFLAGRATION, new ModifierInjector()
      {
         @Override
         public void inject(StatModifiers statModifiers)
         {
            statModifiers.critChance.add(new CritChanceModifier()
            {
               @Override
               public double get()
               {
                  return 0.06;
               }
            });
         }
      });
   }
}
