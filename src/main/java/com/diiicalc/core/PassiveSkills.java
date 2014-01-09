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

   public static final String SLUG_WITCH_DOCTOR_JUNGLE_FORTITUDE = "jungle-fortitude";
   public static final String SLUG_WITCH_DOCTOR_CIRCLE_OF_LIFE = "circle-of-life";
   public static final String SLUG_WITCH_DOCTOR_SPIRITUAL_ATTUNEMENT = "spiritual-attunement";
   public static final String SLUG_WITCH_DOCTOR_GRUESOME_FEAST = "gruesome-feast";
   public static final String SLUG_WITCH_DOCTOR_BLOOD_RITUAL = "blood-ritual";
   public static final String SLUG_WITCH_DOCTOR_BAD_MEDICINE = "bad-medicine";
   public static final String SLUG_WITCH_DOCTOR_ZOMBIE_HANDLER = "zombie-handler";
   public static final String SLUG_WITCH_DOCTOR_PIERCE_THE_VEIL = "pierce-the-veil";
   public static final String SLUG_WITCH_DOCTOR_SPIRIT_VESSEL = "spirit-vessel";
   public static final String SLUG_WITCH_DOCTOR_FETISH_SYCOPHANTS = "fetish-sycophants";
   public static final String SLUG_WITCH_DOCTOR_RUSH_OF_ESSENCE = "rush-of-essence";
   public static final String SLUG_WITCH_DOCTOR_VISION_QUEST = "vision-quest";
   public static final String SLUG_WITCH_DOCTOR_FIERCE_LOYALTY = "fierce-loyalty";
   public static final String SLUG_WITCH_DOCTOR_GRAVE_INJUSTICE = "grave-injustice";
   public static final String SLUG_WITCH_DOCTOR_TRIBAL_RITES = "tribal-rites";

   public static final String SLUG_DEMON_HUNTER_THRILL_OF_THE_HUNT = "thrill-of-the-hunt";
   public static final String SLUG_DEMON_HUNTER_TACTICAL_ADVANTAGE = "tactical-advantage";
   public static final String SLUG_DEMON_HUNTER_VENGEANCE = "vengeance";
   public static final String SLUG_DEMON_HUNTER_STEADY_AIM = "steady-aim";
   public static final String SLUG_DEMON_HUNTER_CULL_THE_WEAK = "cull-the-weak";
   public static final String SLUG_DEMON_HUNTER_NIGHT_STALKER = "night-stalker";
   public static final String SLUG_DEMON_HUNTER_BROODING = "brooding";
   public static final String SLUG_DEMON_HUNTER_HOT_PURSUIT = "hot-pursuit";
   public static final String SLUG_DEMON_HUNTER_ARCHERY = "archery";
   public static final String SLUG_DEMON_HUNTER_NUMBING_TRAPS = "numbing-traps";
   public static final String SLUG_DEMON_HUNTER_PERFECTIONIST = "perfectionist";
   public static final String SLUG_DEMON_HUNTER_CUSTOM_ENGINEERING = "custom-engineering";
   public static final String SLUG_DEMON_HUNTER_GRENADIER = "grenadier";
   public static final String SLUG_DEMON_HUNTER_SHARPSHOOTER = "sharpshooter";
   public static final String SLUG_DEMON_HUNTER_BALLISTICS = "ballistics";

   public static final String SLUG_MONK_RESOLVE = "resolve";
   public static final String SLUG_MONK_FLEET_FOOTED = "fleet-footed";
   public static final String SLUG_MONK_EXALTED_SOUL = "exalted-soul";
   public static final String SLUG_MONK_TRANSCENDENCE = "transcendence";
   public static final String SLUG_MONK_CHANT_OF_RESONANCE = "chant-of-resonance";
   public static final String SLUG_MONK_SEIZE_THE_INITIATIVE = "seize-the-initiative";
   public static final String SLUG_MONK_THE_GUARDIANS_PATH = "the-guardians-path";
   public static final String SLUG_MONK_SIXTH_SENSE = "sixth-sense";
   public static final String SLUG_MONK_PACIFISM = "pacifism";
   public static final String SLUG_MONK_BEACON_OF_YTAR = "beacon-of-ytar";
   public static final String SLUG_MONK_GUIDING_LIGHT = "guiding-light";
   public static final String SLUG_MONK_ONE_WITH_EVERYTHING = "one-with-everything";
   public static final String SLUG_MONK_COMBINATION_STRIKE = "combination-strike";
   public static final String SLUG_MONK_NEAR_DEATH_EXPERIENCE = "near-death-experience";

   private static final String[] ALL_WIZARD_SLUGS =
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

   private static final String[] ALL_BARBARIAN_SLUGS =
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

   private static final String[] ALL_WITCH_DOCTOR_SLUGS =
   {
      SLUG_WITCH_DOCTOR_JUNGLE_FORTITUDE,
      SLUG_WITCH_DOCTOR_CIRCLE_OF_LIFE,
      SLUG_WITCH_DOCTOR_SPIRITUAL_ATTUNEMENT,
      SLUG_WITCH_DOCTOR_GRUESOME_FEAST,
      SLUG_WITCH_DOCTOR_BLOOD_RITUAL,
      SLUG_WITCH_DOCTOR_BAD_MEDICINE,
      SLUG_WITCH_DOCTOR_ZOMBIE_HANDLER,
      SLUG_WITCH_DOCTOR_PIERCE_THE_VEIL,
      SLUG_WITCH_DOCTOR_SPIRIT_VESSEL,
      SLUG_WITCH_DOCTOR_FETISH_SYCOPHANTS,
      SLUG_WITCH_DOCTOR_RUSH_OF_ESSENCE,
      SLUG_WITCH_DOCTOR_VISION_QUEST,
      SLUG_WITCH_DOCTOR_FIERCE_LOYALTY,
      SLUG_WITCH_DOCTOR_GRAVE_INJUSTICE,
      SLUG_WITCH_DOCTOR_TRIBAL_RITES,
   };

   private static final String[] ALL_DEMON_HUNTER_SLUGS =
   {
      SLUG_DEMON_HUNTER_THRILL_OF_THE_HUNT,
      SLUG_DEMON_HUNTER_TACTICAL_ADVANTAGE,
      SLUG_DEMON_HUNTER_VENGEANCE,
      SLUG_DEMON_HUNTER_STEADY_AIM,
      SLUG_DEMON_HUNTER_CULL_THE_WEAK,
      SLUG_DEMON_HUNTER_NIGHT_STALKER,
      SLUG_DEMON_HUNTER_BROODING,
      SLUG_DEMON_HUNTER_HOT_PURSUIT,
      SLUG_DEMON_HUNTER_ARCHERY,
      SLUG_DEMON_HUNTER_NUMBING_TRAPS,
      SLUG_DEMON_HUNTER_PERFECTIONIST,
      SLUG_DEMON_HUNTER_CUSTOM_ENGINEERING,
      SLUG_DEMON_HUNTER_GRENADIER,
      SLUG_DEMON_HUNTER_SHARPSHOOTER,
      SLUG_DEMON_HUNTER_BALLISTICS,
   };

   private static final String[] ALL_MONK_SLUGS =
   {
      SLUG_MONK_RESOLVE,
      SLUG_MONK_FLEET_FOOTED,
      SLUG_MONK_EXALTED_SOUL,
      SLUG_MONK_TRANSCENDENCE,
      SLUG_MONK_CHANT_OF_RESONANCE,
      SLUG_MONK_SEIZE_THE_INITIATIVE,
      SLUG_MONK_THE_GUARDIANS_PATH,
      SLUG_MONK_SIXTH_SENSE,
      SLUG_MONK_PACIFISM,
      SLUG_MONK_BEACON_OF_YTAR,
      SLUG_MONK_GUIDING_LIGHT,
      SLUG_MONK_ONE_WITH_EVERYTHING,
      SLUG_MONK_COMBINATION_STRIKE,
      SLUG_MONK_NEAR_DEATH_EXPERIENCE,
   };

   public static Skill lookup(String slug)
   {
      return SKILL_MAP.get(slug);
   }

   private static Map<String, Skill> SKILL_MAP = new HashMap<String, Skill>();

   public static final Map<String, String[]> HERO_TYPE_TO_SLUGS_MAP = new HashMap<String, String[]>();

   static
   {
      // Populate Hero -> Slug[] map.
      HERO_TYPE_TO_SLUGS_MAP.put(Constants.HERO_TYPE_BARBARIAN, ALL_BARBARIAN_SLUGS);
      HERO_TYPE_TO_SLUGS_MAP.put(Constants.HERO_TYPE_MONK, ALL_MONK_SLUGS);
      HERO_TYPE_TO_SLUGS_MAP.put(Constants.HERO_TYPE_WIZARD, ALL_WIZARD_SLUGS);
      HERO_TYPE_TO_SLUGS_MAP.put(Constants.HERO_TYPE_WITCH_DOCTOR, ALL_WITCH_DOCTOR_SLUGS);
      HERO_TYPE_TO_SLUGS_MAP.put(Constants.HERO_TYPE_DEMON_HUNTER, ALL_DEMON_HUNTER_SLUGS);

      // This code is GENERATED. See SkillDownloader for details.
      SKILL_MAP.put("resolve", new Skill("resolve", "monk_passive_resolve", "Resolve"));
      SKILL_MAP.put("fleet-footed", new Skill("fleet-footed", "monk_passive_fleetfooted", "Fleet Footed"));
      SKILL_MAP.put("exalted-soul", new Skill("exalted-soul", "monk_passive_exaltedsoul", "Exalted Soul"));
      SKILL_MAP.put("transcendence", new Skill("transcendence", "monk_passive_transcendence", "Transcendence"));
      SKILL_MAP.put("chant-of-resonance", new Skill("chant-of-resonance", "monk_passive_chantofresonance", "Chant of Resonance"));
      SKILL_MAP.put("seize-the-initiative", new Skill("seize-the-initiative", "monk_passive_seizetheinitiative", "Seize the Initiative"));
      SKILL_MAP.put("the-guardians-path", new Skill("the-guardians-path", "monk_passive_theguardianspath", "The Guardian's Path"));
      SKILL_MAP.put("sixth-sense", new Skill("sixth-sense", "monk_passive_sixthsense", "Sixth Sense"));
      SKILL_MAP.put("pacifism", new Skill("pacifism", "monk_passive_pacifism", "Pacifism"));
      SKILL_MAP.put("beacon-of-ytar", new Skill("beacon-of-ytar", "monk_passive_beaconofytar", "Beacon of Ytar"));
      SKILL_MAP.put("guiding-light", new Skill("guiding-light", "monk_passive_guidinglight", "Guiding Light"));
      SKILL_MAP.put("one-with-everything", new Skill("one-with-everything", "monk_passive_onewitheverything", "One With Everything"));
      SKILL_MAP.put("combination-strike", new Skill("combination-strike", "monk_passive_combinationstrike", "Combination Strike"));
      SKILL_MAP.put("near-death-experience", new Skill("near-death-experience", "monk_passive_neardeathexperience", "Near Death Experience"));

      SKILL_MAP.put("thrill-of-the-hunt", new Skill("thrill-of-the-hunt", "demonhunter_passive_thrillofthehunt", "Thrill of the Hunt"));
      SKILL_MAP.put("tactical-advantage", new Skill("tactical-advantage", "demonhunter_passive_tacticaladvantage", "Tactical Advantage"));
      SKILL_MAP.put("vengeance", new Skill("vengeance", "demonhunter_passive_vengeance", "Vengeance"));
      SKILL_MAP.put("steady-aim", new Skill("steady-aim", "demonhunter_passive_steadyaim", "Steady Aim"));
      SKILL_MAP.put("cull-the-weak", new Skill("cull-the-weak", "demonhunter_passive_culltheweak", "Cull the Weak"));
      SKILL_MAP.put("night-stalker", new Skill("night-stalker", "demonhunter_passive_nightstalker", "Night Stalker"));
      SKILL_MAP.put("brooding", new Skill("brooding", "demonhunter_passive_brooding", "Brooding"));
      SKILL_MAP.put("hot-pursuit", new Skill("hot-pursuit", "demonhunter_passive_hotpursuit", "Hot Pursuit"));
      SKILL_MAP.put("archery", new Skill("archery", "demonhunter_passive_archery", "Archery"));
      SKILL_MAP.put("numbing-traps", new Skill("numbing-traps", "demonhunter_passive_numbingtraps", "Numbing Traps"));
      SKILL_MAP.put("perfectionist", new Skill("perfectionist", "demonhunter_passive_perfectionist", "Perfectionist"));
      SKILL_MAP.put("custom-engineering", new Skill("custom-engineering", "demonhunter_passive_customengineering", "Custom Engineering"));
      SKILL_MAP.put("grenadier", new Skill("grenadier", "demonhunter_passive_grenadier", "Grenadier"));
      SKILL_MAP.put("sharpshooter", new Skill("sharpshooter", "demonhunter_passive_sharpshooter", "Sharpshooter"));
      SKILL_MAP.put("ballistics", new Skill("ballistics", "demonhunter_passive_ballistics", "Ballistics"));

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
      SKILL_MAP.put("jungle-fortitude", new Skill("jungle-fortitude", "witchdoctor_passive_junglefortitude", "Jungle Fortitude"));
      SKILL_MAP.put("circle-of-life", new Skill("circle-of-life", "witchdoctor_passive_circleoflife", "Circle of Life"));
      SKILL_MAP.put("spiritual-attunement", new Skill("spiritual-attunement", "witchdoctor_passive_spiritualattunement", "Spiritual Attunement"));
      SKILL_MAP.put("gruesome-feast", new Skill("gruesome-feast", "witchdoctor_passive_gruesomefeast", "Gruesome Feast"));
      SKILL_MAP.put("blood-ritual", new Skill("blood-ritual", "witchdoctor_passive_bloodritual", "Blood Ritual"));
      SKILL_MAP.put("bad-medicine", new Skill("bad-medicine", "witchdoctor_passive_badmedicine", "Bad Medicine"));
      SKILL_MAP.put("zombie-handler", new Skill("zombie-handler", "witchdoctor_passive_zombiehandler", "Zombie Handler"));
      SKILL_MAP.put("pierce-the-veil", new Skill("pierce-the-veil", "witchdoctor_passive_piercetheveil", "Pierce the Veil"));
      SKILL_MAP.put("spirit-vessel", new Skill("spirit-vessel", "witchdoctor_passive_spiritvessel", "Spirit Vessel"));
      SKILL_MAP.put("fetish-sycophants", new Skill("fetish-sycophants", "witchdoctor_passive_fetishsycophants", "Fetish Sycophants"));
      SKILL_MAP.put("rush-of-essence", new Skill("rush-of-essence", "witchdoctor_passive_rushofessence", "Rush of Essence"));
      SKILL_MAP.put("vision-quest", new Skill("vision-quest", "witchdoctor_passive_visionquest", "Vision Quest"));
      SKILL_MAP.put("fierce-loyalty", new Skill("fierce-loyalty", "witchdoctor_passive_fierceloyalty", "Fierce Loyalty"));
      SKILL_MAP.put("grave-injustice", new Skill("grave-injustice", "witchdoctor_passive_graveinjustice", "Grave Injustice"));
      SKILL_MAP.put("tribal-rites", new Skill("tribal-rites", "witchdoctor_passive_tribalrites", "Tribal Rites"));

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
