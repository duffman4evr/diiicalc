package com.diiicalc.core;

import com.diiicalc.api.Skill;

import java.util.HashMap;
import java.util.Map;

public class PassiveSkills
{
   public static final String SLUG_WIZARD_POWER_HUNGRY = "power-hungry";
   public static final String SLUG_WIZARD_BLUR = "blur";
   public static final String SLUG_WIZARD_EVOCATION = "evocation";
   public static final String SLUG_WIZARD_GLASS_CANNON = "glass-cannon";
   public static final String SLUG_WIZARD_PRODIGY = "prodigy";
   public static final String SLUG_WIZARD_ASTRAL_PRESENCE = "astral-presence";
   public static final String SLUG_WIZARD_ILLUSIONIST = "illusionist";
   public static final String SLUG_WIZARD_COLD_BLOODED = "cold-blooded";
   public static final String SLUG_WIZARD_CONFLAGRATION = "conflagration";
   public static final String SLUG_WIZARD_PARALYSIS = "paralysis";
   public static final String SLUG_WIZARD_GALVANIZING_WARD = "galvanizing-ward";
   public static final String SLUG_WIZARD_TEMPORAL_FLUX = "temporal-flux";
   public static final String SLUG_WIZARD_CRITICAL_MASS = "critical-mass";
   public static final String SLUG_WIZARD_ARCANE_DYNAMO = "arcane-dynamo";
   public static final String SLUG_WIZARD_UNSTABLE_ANOMALY = "unstable-anomaly";

   public static final String SLUG_BARBARIAN_POUND_OF_FLESH = "pound-of-flesh";
   public static final String SLUG_BARBARIAN_RUTHLESS = "ruthless";
   public static final String SLUG_BARBARIAN_NERVES_OF_STEEL = "nerves-of-steel";
   public static final String SLUG_BARBARIAN_WEAPONS_MASTER = "weapons-master";
   public static final String SLUG_BARBARIAN_INSPIRING_PRESENCE = "inspiring-presence";
   public static final String SLUG_BARBARIAN_BERSERKER_RAGE = "berserker-rage";
   public static final String SLUG_BARBARIAN_BLOODTHIRST = "bloodthirst";
   public static final String SLUG_BARBARIAN_ANIMOSITY = "animosity";
   public static final String SLUG_BARBARIAN_SUPERSTITION = "superstition";
   public static final String SLUG_BARBARIAN_TOUGH_AS_NAILS = "tough-as-nails";
   public static final String SLUG_BARBARIAN_NO_ESCAPE = "no-escape";
   public static final String SLUG_BARBARIAN_RELENTLESS = "relentless";
   public static final String SLUG_BARBARIAN_BRAWLER = "brawler";
   public static final String SLUG_BARBARIAN_JUGGERNAUT = "juggernaut";
   public static final String SLUG_BARBARIAN_UNFORGIVING = "unforgiving";
   public static final String SLUG_BARBARIAN_BOON_OF_BULKATHOS = "boon-of-bulkathos";

   public static final String[] ALL_WIZARD_SLUGS =
   {
      SLUG_WIZARD_POWER_HUNGRY,
      SLUG_WIZARD_BLUR,
      SLUG_WIZARD_EVOCATION,
      SLUG_WIZARD_GLASS_CANNON,
      SLUG_WIZARD_PRODIGY,
      SLUG_WIZARD_ASTRAL_PRESENCE,
      SLUG_WIZARD_ILLUSIONIST,
      SLUG_WIZARD_COLD_BLOODED,
      SLUG_WIZARD_CONFLAGRATION,
      SLUG_WIZARD_PARALYSIS,
      SLUG_WIZARD_GALVANIZING_WARD,
      SLUG_WIZARD_TEMPORAL_FLUX,
      SLUG_WIZARD_CRITICAL_MASS,
      SLUG_WIZARD_ARCANE_DYNAMO,
      SLUG_WIZARD_UNSTABLE_ANOMALY,
   };

   public static final String[] ALL_BARBARIAN_SLUGS =
   {
      SLUG_BARBARIAN_POUND_OF_FLESH,
      SLUG_BARBARIAN_RUTHLESS,
      SLUG_BARBARIAN_NERVES_OF_STEEL,
      SLUG_BARBARIAN_WEAPONS_MASTER,
      SLUG_BARBARIAN_INSPIRING_PRESENCE,
      SLUG_BARBARIAN_BERSERKER_RAGE,
      SLUG_BARBARIAN_BLOODTHIRST,
      SLUG_BARBARIAN_ANIMOSITY,
      SLUG_BARBARIAN_SUPERSTITION,
      SLUG_BARBARIAN_TOUGH_AS_NAILS,
      SLUG_BARBARIAN_NO_ESCAPE,
      SLUG_BARBARIAN_RELENTLESS,
      SLUG_BARBARIAN_BRAWLER,
      SLUG_BARBARIAN_JUGGERNAUT,
      SLUG_BARBARIAN_UNFORGIVING,
      SLUG_BARBARIAN_BOON_OF_BULKATHOS,
   };

   public static Skill lookup(String slug)
   {
      return SKILL_MAP.get(slug);
   }

   private static Map<String, Skill> SKILL_MAP = new HashMap<String, Skill>();

   static
   {
      // This code is GENERATED. See SkillDownloader for details.
      SKILL_MAP.put("power-hungry", new Skill("power-hungry", "wizard_passive_powerhungry", "Power Hungry"));
      SKILL_MAP.put("blur", new Skill("blur", "wizard_passive_blur", "Blur"));
      SKILL_MAP.put("evocation", new Skill("evocation", "wizard_passive_evocation", "Evocation"));
      SKILL_MAP.put("glass-cannon", new Skill("glass-cannon", "wizard_passive_glasscannon", "Glass Cannon"));
      SKILL_MAP.put("prodigy", new Skill("prodigy", "wizard_passive_prodigy", "Prodigy"));
      SKILL_MAP.put("astral-presence", new Skill("astral-presence", "wizard_passive_astralpresence", "Astral Presence"));
      SKILL_MAP.put("illusionist", new Skill("illusionist", "wizard_passive_illusionist", "Illusionist"));
      SKILL_MAP.put("cold-blooded", new Skill("cold-blooded", "wizard_passive_coldblooded", "Cold Blooded"));
      SKILL_MAP.put("conflagration", new Skill("conflagration", "wizard_passive_conflagration", "Conflagration"));
      SKILL_MAP.put("paralysis", new Skill("paralysis", "wizard_passive_paralysis", "Paralysis"));
      SKILL_MAP.put("galvanizing-ward", new Skill("galvanizing-ward", "wizard_passive_galvanizingward", "Galvanizing Ward"));
      SKILL_MAP.put("temporal-flux", new Skill("temporal-flux", "wizard_passive_temporalflux", "Temporal Flux"));
      SKILL_MAP.put("critical-mass", new Skill("critical-mass", "wizard_passive_criticalmass", "Critical Mass"));
      SKILL_MAP.put("arcane-dynamo", new Skill("arcane-dynamo", "wizard_passive_arcanedynamo", "Arcane Dynamo"));
      SKILL_MAP.put("unstable-anomaly", new Skill("unstable-anomaly", "wizard_passive_unstableanomaly", "Unstable Anomaly"));
      SKILL_MAP.put("pound-of-flesh", new Skill("pound-of-flesh", "barbarian_passive_poundofflesh", "Pound of Flesh"));
      SKILL_MAP.put("ruthless", new Skill("ruthless", "barbarian_passive_ruthless", "Ruthless"));
      SKILL_MAP.put("nerves-of-steel", new Skill("nerves-of-steel", "barbarian_passive_nervesofsteel", "Nerves Of Steel"));
      SKILL_MAP.put("weapons-master", new Skill("weapons-master", "barbarian_passive_weaponsmaster", "Weapons Master"));
      SKILL_MAP.put("inspiring-presence", new Skill("inspiring-presence", "barbarian_passive_inspiringpresence", "Inspiring Presence"));
      SKILL_MAP.put("berserker-rage", new Skill("berserker-rage", "barbarian_passive_berserkerrage", "Berserker Rage"));
      SKILL_MAP.put("bloodthirst", new Skill("bloodthirst", "barbarian_passive_bloodthirst", "Bloodthirst"));
      SKILL_MAP.put("animosity", new Skill("animosity", "barbarian_passive_animosity", "Animosity"));
      SKILL_MAP.put("superstition", new Skill("superstition", "barbarian_passive_superstition", "Superstition"));
      SKILL_MAP.put("tough-as-nails", new Skill("tough-as-nails", "barbarian_passive_toughasnails", "Tough as Nails"));
      SKILL_MAP.put("no-escape", new Skill("no-escape", "barbarian_passive_noescape", "No Escape"));
      SKILL_MAP.put("relentless", new Skill("relentless", "barbarian_passive_relentless", "Relentless"));
      SKILL_MAP.put("brawler", new Skill("brawler", "barbarian_passive_brawler", "Brawler"));
      SKILL_MAP.put("juggernaut", new Skill("juggernaut", "barbarian_passive_juggernaut", "Juggernaut"));
      SKILL_MAP.put("unforgiving", new Skill("unforgiving", "barbarian_passive_unforgiving", "Unforgiving"));
      SKILL_MAP.put("boon-of-bulkathos", new Skill("boon-of-bulkathos", "barbarian_passive_boonofbulkathos", "Boon of Bul-Kathos"));
   }
}
