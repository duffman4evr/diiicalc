package com.diiicalc.core;

import com.diiicalc.api.*;
import com.diiicalc.core.modifiers.ArmorModifier;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.utils.URIBuilder;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class Utils
{
   public static final ObjectMapper JSON_MAPPER = new ObjectMapper();

   public static <T> T doGet(BattlenetRealm realm, String path, Class<T> typeClass) throws Exception
   {
      return Utils.doGetWithArgs(realm, path, null, typeClass);
   }

   public static <T> T doGetWithArgs(BattlenetRealm realm, String path, Map<String, String> queryArgs, Class<T> typeClass) throws Exception
   {
      CloseableHttpClient httpClient = HttpClients.createDefault();

      URIBuilder builder;

      if (realm == BattlenetRealm.US)
      {
         builder = new URIBuilder("http://us.battle.net");
      }
      else
      {
         builder = new URIBuilder("http://eu.battle.net");
      }

      builder.setPath(path);

      if (queryArgs != null && queryArgs.size() > 0)
      {
         for (Map.Entry<String, String> entry : queryArgs.entrySet())
         {
            builder.setParameter(entry.getKey(), entry.getValue());
         }
      }

      HttpGet httpget = new HttpGet(builder.build());

      CloseableHttpResponse response = httpClient.execute(httpget);

      try
      {
         HttpEntity entity = response.getEntity();

         if (entity != null)
         {
            String entityString = EntityUtils.toString(entity);

            BlizzardApiError error = JSON_MAPPER.readValue(entityString, BlizzardApiError.class);

            if (error.getCode() != null)
            {
               throw new Exception("Not sure what to do about exceptions yet.");
            }

            return JSON_MAPPER.readValue(entityString, typeClass);
         }
      }
      finally
      {
         response.close();
      }

      return null;
   }

   public static Map<String, Item> pullDownItemMap(Hero hero) throws Exception
   {
      Map<String, Item> itemMap = new HashMap<String, Item>();

      for (Map.Entry<String, Hero.ItemSummary> entry : hero.getItems().entrySet())
      {
         String itemPath = Constants.DATA_API_URL_PREFIX + "/" + entry.getValue().getTooltipParams();

         itemMap.put(entry.getKey(), Utils.doGet(BattlenetRealm.US, itemPath, Item.class));
      }

      return itemMap;
   }

   public static DefensiveStats computeDefensiveStats(StatTotals statTotals, long monsterLevel)
   {
      double armor = statTotals.getArmor();
      double resistAll = statTotals.getAllResist();
      double dodgeChance = statTotals.getDodgeChance();
      double averageBlockAmount = statTotals.getAverageBlockAmount();

      Map<String, Double> incomingDamageModifiers = statTotals.getIncomingDamageModifiers(monsterLevel);

      double totalIncomingDamageModifier = 1.0;

      for (Map.Entry<String, Double> damageModifier : incomingDamageModifiers.entrySet())
      {
         totalIncomingDamageModifier *= damageModifier.getValue();
      }

      double effectiveLifeMultiplier = (1 / totalIncomingDamageModifier) * (1 + (averageBlockAmount / statTotals.getLife()));
      double effectiveLife = statTotals.getLife() * effectiveLifeMultiplier;

      DefensiveStats defensiveStats = new DefensiveStats
      (
         armor,
         resistAll,
         dodgeChance,
         averageBlockAmount,
         incomingDamageModifiers,
         totalIncomingDamageModifier,
         effectiveLife
      );

      return defensiveStats;
   }

   public static OffensiveStats computeOffensiveStats(StatTotals statTotals)
   {
      double critChance = statTotals.getCritChance();
      double critDamage = 1 + statTotals.getCritDamageBonus();

      // Looks like the equation is this:
      //       Weapon Term: [AverageWeaponDamageForBothMainAndOffHand + DamageBonusFromNonWeaponItems] *
      // Attack Speed Term: [AverageWeaponAttackSpeedForBothMainAndOffHand * (1 + AttackSpeedBonusesFromNonWeaponItemsIncludingDualWieldingBonus)] *
      //    Main Stat Term: [1 + (MainStat * .01)] *
      //         Crit Term: [1 + (CritChance * CritDamageBonus)] *
      //        Skill Term: [1 + SkillBonus]
      double weaponTerm = statTotals.getAverageWeaponDamage() + statTotals.getAverageBonusDamage();
      double attackSpeedTerm = statTotals.getWeaponAttackSpeed()  * (1 + statTotals.getAttackSpeedBonus());
      double mainStatTerm = 1 + (statTotals.getMainStat() * 0.01);
      double critTerm = 1 + (critChance * (critDamage - 1));
      double skillTerm = 1 + statTotals.getSkillDamageBonus();

      double weaponDamage = weaponTerm * mainStatTerm * critTerm * skillTerm;
      double dps = weaponDamage * attackSpeedTerm;

      return new OffensiveStats(dps, weaponDamage, attackSpeedTerm, critChance, critDamage);
   }

   // Some publicly exposed helper functions.
   public static long computeBaseVitality(long level, long paragonLevel)
   {
      long levelWithParagon = level + paragonLevel;

      return 7 + (2 * levelWithParagon);
   }

   public static long computeBaseDexterity(long level, long paragonLevel, String heroType)
   {
      if (Constants.HERO_TYPE_DEMON_HUNTER.equals(heroType) || Constants.HERO_TYPE_MONK.equals(heroType))
      {
         return 7 + (3 * (level + paragonLevel));
      }
      else
      {
         return 7 + (level + paragonLevel);
      }
   }

   public static long computeBaseIntelligence(long level, long paragonLevel, String heroType)
   {
      if (Constants.HERO_TYPE_WITCH_DOCTOR.equals(heroType) || Constants.HERO_TYPE_WIZARD.equals(heroType))
      {
         return 7 + (3 * (level + paragonLevel));
      }
      else
      {
         return 7 + (level + paragonLevel);
      }
   }

   public static long computeBaseStrength(long level, long paragonLevel, String heroType)
   {
      if (Constants.HERO_TYPE_BARBARIAN.equals(heroType))
      {
         return 7 + (3 * (level + paragonLevel));
      }
      else
      {
         return 7 + (level + paragonLevel);
      }
   }

   public static long computeLifePerVitality(long level)
   {
      long lifePerVit = 10;

      if (level > 35)
      {
         lifePerVit = level - 25;
      }

      return lifePerVit;
   }

   public static double computeDodgeChanceFromDexterity(double dex)
   {
      double dodgeChance = 0;

      if (dex <= 100)
      {
         return dex * .001;
      }
      else
      {
         dodgeChance += 0.1;
         dex -= 100;
      }

      if (dex <= 400)
      {
         return dodgeChance + (dex * .00025);
      }
      else
      {
         dodgeChance += 0.1;
         dex -= 400;
      }

      if (dex <= 500)
      {
         return dodgeChance + (dex * .00020);
      }
      else
      {
         dodgeChance += 0.1;
         dex -= 500;
      }

      dodgeChance += dex * .00010;

      if (dodgeChance > 0.999)
      {
         dodgeChance = 0.999;
      }

      return dodgeChance;
   }

   public static double findMin(double ... args)
   {
      double min = args[0];

      for (double d : args)
      {
         if (d < min)
         {
            min = d;
         }
      }

      return min;
   }

   public static ModifierInjector lookupInjector(Skill skill, Rune rune)
   {
      String lookupKey = skill.getSlug() + "|" + rune.getSlug();

      return INJECTOR_MAP.get(lookupKey);
   }

   public static boolean isWeapon(String type)
   {
      return WEAPON_TYPES.contains(type);
   }

   public static Double getAttackSpeedForWeaponType(Item.Type type)
   {
      return WEAPON_TYPE_TO_ATTACK_SPEED_MAP.get(type);
   }

   // -----
   // Private stuff.
   // -----
   private static final Map<String, ModifierInjector> INJECTOR_MAP = new HashMap<String, ModifierInjector>();
   private static Set<String> WEAPON_TYPES = new HashSet<String>();
   private static final Map<Item.Type, Double> WEAPON_TYPE_TO_ATTACK_SPEED_MAP = new HashMap<Item.Type, Double>();

   static
   {
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_DAGGER);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_SWORD);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_MACE);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_AXE);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_POLEARM);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_SPEAR);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_MIGHTY_WEAPON);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_BOW);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_CROSSBOW);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_HAND_CROSSBOW);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_WAND);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_DIABO);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_FIST_WEAPON);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_CEREMONIAL_KNIFE);
      WEAPON_TYPES.add(Constants.WEAPON_TYPE_STAFF);

      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_DAGGER, false), 1.5);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_SWORD, false), 1.4);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_SWORD, true), 1.1);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_MACE, false), 1.2);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_MACE, true), 0.9);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_AXE, false), 1.3);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_AXE, true), 1.1);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_POLEARM, true), 0.95);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_SPEAR, false), 1.5);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_MIGHTY_WEAPON, false), 1.3);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_MIGHTY_WEAPON, true), 1.0);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_BOW, true), 1.4);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_CROSSBOW, true), 1.1);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_HAND_CROSSBOW, false), 1.6);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_WAND, false), 1.4);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_DIABO, true), 1.1);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_FIST_WEAPON, false), 1.4);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_CEREMONIAL_KNIFE, false), 1.4);
      WEAPON_TYPE_TO_ATTACK_SPEED_MAP.put(new Item.Type(Constants.WEAPON_TYPE_STAFF, true), 1.0);

      INJECTOR_MAP.put("big-bad-voodoo|big-bad-voodoo-a", new ModifierInjector()
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
   }
}