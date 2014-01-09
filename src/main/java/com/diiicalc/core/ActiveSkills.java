package com.diiicalc.core;

import com.diiicalc.api.Rune;
import com.diiicalc.api.Skill;

import java.util.HashMap;
import java.util.Map;

public class ActiveSkills
{
   public static final String SLUG_WIZARD_MAGIC_MISSLE = "magic-missile";
   public static final String SLUG_WIZARD_RAY_OF_FROST = "ray-of-frost";
   public static final String SLUG_WIZARD_SHOCK_PULSE = "shock-pulse";
   public static final String SLUG_WIZARD_FROST_NOVA = "frost-nova";
   public static final String SLUG_WIZARD_ARCANE_ORB = "arcane-orb";
   public static final String SLUG_WIZARD_DIAMOND_SKIN = "diamond-skin";
   public static final String SLUG_WIZARD_WAVE_OF_FORCE = "wave-of-force";
   public static final String SLUG_WIZARD_SPECTRAL_BLADE = "spectral-blade";
   public static final String SLUG_WIZARD_ARCANE_TORRENT = "arcane-torrent";
   public static final String SLUG_WIZARD_ENERGY_TWISTER = "energy-twister";
   public static final String SLUG_WIZARD_ICE_ARMOR = "ice-armor";
   public static final String SLUG_WIZARD_ELECTROCUTE = "electrocute";
   public static final String SLUG_WIZARD_SLOW_TIME = "slow-time";
   public static final String SLUG_WIZARD_STORM_ARMOR = "storm-armor";
   public static final String SLUG_WIZARD_EXPLOSIVE_BLAST = "explosive-blast";
   public static final String SLUG_WIZARD_MAGIC_WEAPON = "magic-weapon";
   public static final String SLUG_WIZARD_HYDRA = "hydra";
   public static final String SLUG_WIZARD_DISINTEGRATE = "disintegrate";
   public static final String SLUG_WIZARD_FAMILIAR = "familiar";
   public static final String SLUG_WIZARD_TELEPORT = "teleport";
   public static final String SLUG_WIZARD_MIRROR_IMAGE = "mirror-image";
   public static final String SLUG_WIZARD_METEOR = "meteor";
   public static final String SLUG_WIZARD_BLIZZARD = "blizzard";
   public static final String SLUG_WIZARD_ENERGY_ARMOR = "energy-armor";
   public static final String SLUG_WIZARD_ARCHON = "archon";

   public static final String SLUG_BARBARIAN_BASH = "bash";
   public static final String SLUG_BARBARIAN_HAMMER_OF_THE_ANCIENTS = "hammer-of-the-ancients";
   public static final String SLUG_BARBARIAN_CLEAVE = "cleave";
   public static final String SLUG_BARBARIAN_GROUND_STOMP = "ground-stomp";
   public static final String SLUG_BARBARIAN_REND = "rend";
   public static final String SLUG_BARBARIAN_LEAP = "leap";
   public static final String SLUG_BARBARIAN_ANCIENT_SPEAR = "ancient-spear";
   public static final String SLUG_BARBARIAN_FRENZY = "frenzy";
   public static final String SLUG_BARBARIAN_SEISMIC_SLAM = "seismic-slam";
   public static final String SLUG_BARBARIAN_REVENGE = "revenge";
   public static final String SLUG_BARBARIAN_WEAPON_THROW = "weapon-throw";
   public static final String SLUG_BARBARIAN_SPRINT = "sprint";
   public static final String SLUG_BARBARIAN_THREATENING_SHOUT = "threatening-shout";
   public static final String SLUG_BARBARIAN_EARTHQUAKE = "earthquake";
   public static final String SLUG_BARBARIAN_WHIRLWIND = "whirlwind";
   public static final String SLUG_BARBARIAN_FURIOUS_CHARGE = "furious-charge";
   public static final String SLUG_BARBARIAN_IGNORE_PAIN = "ignore-pain";
   public static final String SLUG_BARBARIAN_BATTLE_RAGE = "battle-rage";
   public static final String SLUG_BARBARIAN_CALL_OF_THE_ANCIENTS = "call-of-the-ancients";
   public static final String SLUG_BARBARIAN_OVERPOWER = "overpower";
   public static final String SLUG_BARBARIAN_WAR_CRY = "war-cry";
   public static final String SLUG_BARBARIAN_WRATH_OF_THE_BERSERKER = "wrath-of-the-berserker";

   public static final String SLUG_WITCH_DOCTOR_POISON_DART = "poison-dart";
   public static final String SLUG_WITCH_DOCTOR_CORPSE_SPIDERS = "corpse-spiders";
   public static final String SLUG_WITCH_DOCTOR_PLAGUE_OF_TOADS = "plague-of-toads";
   public static final String SLUG_WITCH_DOCTOR_FIREBOMB = "firebomb";
   public static final String SLUG_WITCH_DOCTOR_GRASP_OF_THE_DEAD = "grasp-of-the-dead";
   public static final String SLUG_WITCH_DOCTOR_FIREBATS = "firebats";
   public static final String SLUG_WITCH_DOCTOR_HAUNT = "haunt";
   public static final String SLUG_WITCH_DOCTOR_LOCUST_SWARM = "locust-swarm";
   public static final String SLUG_WITCH_DOCTOR_SUMMON_ZOMBIE_DOGS = "summon-zombie-dogs";
   public static final String SLUG_WITCH_DOCTOR_HORRIFY = "horrify";
   public static final String SLUG_WITCH_DOCTOR_SPIRIT_WALK = "spirit-walk";
   public static final String SLUG_WITCH_DOCTOR_HEX = "hex";
   public static final String SLUG_WITCH_DOCTOR_SOUL_HARVEST = "soul-harvest";
   public static final String SLUG_WITCH_DOCTOR_SACRIFICE = "sacrifice";
   public static final String SLUG_WITCH_DOCTOR_MASS_CONFUSION = "mass-confusion";
   public static final String SLUG_WITCH_DOCTOR_ZOMBIE_CHARGER = "zombie-charger";
   public static final String SLUG_WITCH_DOCTOR_SPIRIT_BARRAGE = "spirit-barrage";
   public static final String SLUG_WITCH_DOCTOR_ACID_CLOUD = "acid-cloud";
   public static final String SLUG_WITCH_DOCTOR_WALL_OF_ZOMBIES = "wall-of-zombies";
   public static final String SLUG_WITCH_DOCTOR_GARGANTUAN = "gargantuan";
   public static final String SLUG_WITCH_DOCTOR_BIG_BAD_VOODOO = "big-bad-voodoo";
   public static final String SLUG_WITCH_DOCTOR_FETISH_ARMY = "fetish-army";

   public static final String SLUG_DEMON_HUNTER_HUNGERING_ARROW = "hungering-arrow";
   public static final String SLUG_DEMON_HUNTER_ENTANGLING_SHOT = "entangling-shot";
   public static final String SLUG_DEMON_HUNTER_BOLA_SHOT = "bola-shot";
   public static final String SLUG_DEMON_HUNTER_GRENADES = "grenades";
   public static final String SLUG_DEMON_HUNTER_IMPALE = "impale";
   public static final String SLUG_DEMON_HUNTER_RAPID_FIRE = "rapid-fire";
   public static final String SLUG_DEMON_HUNTER_CHAKRAM = "chakram";
   public static final String SLUG_DEMON_HUNTER_ELEMENTAL_ARROW = "elemental-arrow";
   public static final String SLUG_DEMON_HUNTER_CALTROPS = "caltrops";
   public static final String SLUG_DEMON_HUNTER_SMOKE_SCREEN = "smoke-screen";
   public static final String SLUG_DEMON_HUNTER_SHADOW_POWER = "shadow-power";
   public static final String SLUG_DEMON_HUNTER_VAULT = "vault";
   public static final String SLUG_DEMON_HUNTER_PREPARATION = "preparation";
   public static final String SLUG_DEMON_HUNTER_COMPANION = "companion";
   public static final String SLUG_DEMON_HUNTER_MARKED_FOR_DEATH = "marked-for-death";
   public static final String SLUG_DEMON_HUNTER_EVASIVE_FIRE = "evasive-fire";
   public static final String SLUG_DEMON_HUNTER_FAN_OF_KNIVES = "fan-of-knives";
   public static final String SLUG_DEMON_HUNTER_SPIKE_TRAP = "spike-trap";
   public static final String SLUG_DEMON_HUNTER_SENTRY = "sentry";
   public static final String SLUG_DEMON_HUNTER_STRAFE = "strafe";
   public static final String SLUG_DEMON_HUNTER_MULTISHOT = "multishot";
   public static final String SLUG_DEMON_HUNTER_CLUSTER_ARROW = "cluster-arrow";
   public static final String SLUG_DEMON_HUNTER_RAIN_OF_VENGEANCE = "rain-of-vengeance";

   public static final String SLUG_MONK_FISTS_OF_THUNDER = "fists-of-thunder";
   public static final String SLUG_MONK_DEADLY_REACH = "deadly-reach";
   public static final String SLUG_MONK_CRIPPLING_WAVE = "crippling-wave";
   public static final String SLUG_MONK_WAY_OF_THE_HUNDRED_FISTS = "way-of-the-hundred-fists";
   public static final String SLUG_MONK_LASHING_TAIL_KICK = "lashing-tail-kick";
   public static final String SLUG_MONK_TEMPEST_RUSH = "tempest-rush";
   public static final String SLUG_MONK_WAVE_OF_LIGHT = "wave-of-light";
   public static final String SLUG_MONK_BLINDING_FLASH = "blinding-flash";
   public static final String SLUG_MONK_BREATH_OF_HEAVEN = "breath-of-heaven";
   public static final String SLUG_MONK_SERENITY = "serenity";
   public static final String SLUG_MONK_INNER_SANCTUARY = "inner-sanctuary";
   public static final String SLUG_MONK_DASHING_STRIKE = "dashing-strike";
   public static final String SLUG_MONK_EXPLODING_PALM = "exploding-palm";
   public static final String SLUG_MONK_SWEEPING_WIND = "sweeping-wind";
   public static final String SLUG_MONK_CYCLONE_STRIKE = "cyclone-strike";
   public static final String SLUG_MONK_SEVEN_SIDED_STRIKE = "sevensided-strike";
   public static final String SLUG_MONK_MYSTIC_ALLY = "mystic-ally";
   public static final String SLUG_MONK_MANTRA_OF_EVASION = "mantra-of-evasion";
   public static final String SLUG_MONK_MANTRA_OF_RETRIBUTION = "mantra-of-retribution";
   public static final String SLUG_MONK_MANTRA_OF_HEALING = "mantra-of-healing";
   public static final String SLUG_MONK_MANTRA_OF_CONVICTION = "mantra-of-conviction";

   private static final String[] ALL_WIZARD_SLUGS =
   {
      SLUG_WIZARD_MAGIC_MISSLE,
      SLUG_WIZARD_RAY_OF_FROST,
      SLUG_WIZARD_SHOCK_PULSE,
      SLUG_WIZARD_FROST_NOVA,
      SLUG_WIZARD_ARCANE_ORB,
      SLUG_WIZARD_DIAMOND_SKIN,
      SLUG_WIZARD_WAVE_OF_FORCE,
      SLUG_WIZARD_SPECTRAL_BLADE,
      SLUG_WIZARD_ARCANE_TORRENT,
      SLUG_WIZARD_ENERGY_TWISTER,
      SLUG_WIZARD_ICE_ARMOR,
      SLUG_WIZARD_ELECTROCUTE,
      SLUG_WIZARD_SLOW_TIME,
      SLUG_WIZARD_STORM_ARMOR,
      SLUG_WIZARD_EXPLOSIVE_BLAST,
      SLUG_WIZARD_MAGIC_WEAPON,
      SLUG_WIZARD_HYDRA,
      SLUG_WIZARD_DISINTEGRATE,
      SLUG_WIZARD_FAMILIAR,
      SLUG_WIZARD_TELEPORT,
      SLUG_WIZARD_MIRROR_IMAGE,
      SLUG_WIZARD_METEOR,
      SLUG_WIZARD_BLIZZARD,
      SLUG_WIZARD_ENERGY_ARMOR,
      SLUG_WIZARD_ARCHON,
   };

   private static final String[] ALL_BARBARIAN_SLUGS =
   {
      SLUG_BARBARIAN_BASH,
      SLUG_BARBARIAN_HAMMER_OF_THE_ANCIENTS,
      SLUG_BARBARIAN_CLEAVE,
      SLUG_BARBARIAN_GROUND_STOMP,
      SLUG_BARBARIAN_REND,
      SLUG_BARBARIAN_LEAP,
      SLUG_BARBARIAN_ANCIENT_SPEAR,
      SLUG_BARBARIAN_FRENZY,
      SLUG_BARBARIAN_SEISMIC_SLAM,
      SLUG_BARBARIAN_REVENGE,
      SLUG_BARBARIAN_WEAPON_THROW,
      SLUG_BARBARIAN_SPRINT,
      SLUG_BARBARIAN_THREATENING_SHOUT,
      SLUG_BARBARIAN_EARTHQUAKE,
      SLUG_BARBARIAN_WHIRLWIND,
      SLUG_BARBARIAN_FURIOUS_CHARGE,
      SLUG_BARBARIAN_IGNORE_PAIN,
      SLUG_BARBARIAN_BATTLE_RAGE,
      SLUG_BARBARIAN_CALL_OF_THE_ANCIENTS,
      SLUG_BARBARIAN_OVERPOWER,
      SLUG_BARBARIAN_WAR_CRY,
      SLUG_BARBARIAN_WRATH_OF_THE_BERSERKER,
   };

   private static final String[] ALL_WITCH_DOCTOR_SLUGS =
   {
      SLUG_WITCH_DOCTOR_POISON_DART,
      SLUG_WITCH_DOCTOR_CORPSE_SPIDERS,
      SLUG_WITCH_DOCTOR_PLAGUE_OF_TOADS,
      SLUG_WITCH_DOCTOR_FIREBOMB,
      SLUG_WITCH_DOCTOR_GRASP_OF_THE_DEAD,
      SLUG_WITCH_DOCTOR_FIREBATS,
      SLUG_WITCH_DOCTOR_HAUNT,
      SLUG_WITCH_DOCTOR_LOCUST_SWARM,
      SLUG_WITCH_DOCTOR_SUMMON_ZOMBIE_DOGS,
      SLUG_WITCH_DOCTOR_HORRIFY,
      SLUG_WITCH_DOCTOR_SPIRIT_WALK,
      SLUG_WITCH_DOCTOR_HEX,
      SLUG_WITCH_DOCTOR_SOUL_HARVEST,
      SLUG_WITCH_DOCTOR_SACRIFICE,
      SLUG_WITCH_DOCTOR_MASS_CONFUSION,
      SLUG_WITCH_DOCTOR_ZOMBIE_CHARGER,
      SLUG_WITCH_DOCTOR_SPIRIT_BARRAGE,
      SLUG_WITCH_DOCTOR_ACID_CLOUD,
      SLUG_WITCH_DOCTOR_WALL_OF_ZOMBIES,
      SLUG_WITCH_DOCTOR_GARGANTUAN,
      SLUG_WITCH_DOCTOR_BIG_BAD_VOODOO,
      SLUG_WITCH_DOCTOR_FETISH_ARMY,
   };

   private static final String[] ALL_DEMON_HUNTER_SLUGS =
   {
      SLUG_DEMON_HUNTER_HUNGERING_ARROW,
      SLUG_DEMON_HUNTER_ENTANGLING_SHOT,
      SLUG_DEMON_HUNTER_BOLA_SHOT,
      SLUG_DEMON_HUNTER_GRENADES,
      SLUG_DEMON_HUNTER_IMPALE,
      SLUG_DEMON_HUNTER_RAPID_FIRE,
      SLUG_DEMON_HUNTER_CHAKRAM,
      SLUG_DEMON_HUNTER_ELEMENTAL_ARROW,
      SLUG_DEMON_HUNTER_CALTROPS,
      SLUG_DEMON_HUNTER_SMOKE_SCREEN,
      SLUG_DEMON_HUNTER_SHADOW_POWER,
      SLUG_DEMON_HUNTER_VAULT,
      SLUG_DEMON_HUNTER_PREPARATION,
      SLUG_DEMON_HUNTER_COMPANION,
      SLUG_DEMON_HUNTER_MARKED_FOR_DEATH,
      SLUG_DEMON_HUNTER_EVASIVE_FIRE,
      SLUG_DEMON_HUNTER_FAN_OF_KNIVES,
      SLUG_DEMON_HUNTER_SPIKE_TRAP,
      SLUG_DEMON_HUNTER_SENTRY,
      SLUG_DEMON_HUNTER_STRAFE,
      SLUG_DEMON_HUNTER_MULTISHOT,
      SLUG_DEMON_HUNTER_CLUSTER_ARROW,
      SLUG_DEMON_HUNTER_RAIN_OF_VENGEANCE,
   };

   private static final String[] ALL_MONK_SLUGS =
   {
      SLUG_MONK_FISTS_OF_THUNDER,
      SLUG_MONK_DEADLY_REACH,
      SLUG_MONK_CRIPPLING_WAVE,
      SLUG_MONK_WAY_OF_THE_HUNDRED_FISTS,
      SLUG_MONK_LASHING_TAIL_KICK,
      SLUG_MONK_TEMPEST_RUSH,
      SLUG_MONK_WAVE_OF_LIGHT,
      SLUG_MONK_BLINDING_FLASH,
      SLUG_MONK_BREATH_OF_HEAVEN,
      SLUG_MONK_SERENITY,
      SLUG_MONK_INNER_SANCTUARY,
      SLUG_MONK_DASHING_STRIKE,
      SLUG_MONK_EXPLODING_PALM,
      SLUG_MONK_SWEEPING_WIND,
      SLUG_MONK_CYCLONE_STRIKE,
      SLUG_MONK_SEVEN_SIDED_STRIKE,
      SLUG_MONK_MYSTIC_ALLY,
      SLUG_MONK_MANTRA_OF_EVASION,
      SLUG_MONK_MANTRA_OF_RETRIBUTION,
      SLUG_MONK_MANTRA_OF_HEALING,
      SLUG_MONK_MANTRA_OF_CONVICTION,
   };

   public static Skill lookup(String slug)
   {
      return SKILL_MAP.get(slug);
   }

   public static Rune lookupRune(String skillSlug, String runeLetter)
   {
      return RUNE_MAP.get(skillSlug + "-" + runeLetter);
   }

   private static Map<String, Skill> SKILL_MAP = new HashMap<String, Skill>();
   private static Map<String, Rune> RUNE_MAP = new HashMap<String, Rune>();

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
      SKILL_MAP.put("fists-of-thunder", new Skill("fists-of-thunder", "monk_fistsofthunder", "Fists of Thunder"));
      RUNE_MAP.put("fists-of-thunder-a", new Rune("fists-of-thunder-a", "Thunderclap", "Teleport to the target and release an electric shockwave with every punch that hits all enemies within 6 yards of your primary target for 35% weapon damage as Lightning."));
      RUNE_MAP.put("fists-of-thunder-b", new Rune("fists-of-thunder-b", "Bounding Light", "Every third punch releases chain lightning instead of knocking enemies back.  Each lightning strike inflicts 73% weapon damage as Lightning."));
      RUNE_MAP.put("fists-of-thunder-c", new Rune("fists-of-thunder-c", "Static Charge", "Your primary target is charged with static electricity for 5 seconds and takes 37% weapon damage as Lightning when you attack other enemies with Fists of Thunder."));
      RUNE_MAP.put("fists-of-thunder-d", new Rune("fists-of-thunder-d", "Quickening", "Critical Hits generate an additional 15 Spirit for each target hit."));
      RUNE_MAP.put("fists-of-thunder-e", new Rune("fists-of-thunder-e", "Lightning Flash", "Increases your chance to Dodge by 16% for 2 seconds."));
      SKILL_MAP.put("deadly-reach", new Skill("deadly-reach", "monk_deadlyreach", "Deadly Reach"));
      RUNE_MAP.put("deadly-reach-a", new Rune("deadly-reach-a", "Foresight", "The third strike increases the damage of all attacks by 18% for 30 seconds."));
      RUNE_MAP.put("deadly-reach-b", new Rune("deadly-reach-b", "Piercing Trident", "Increases the area of effect of the second and third strikes."));
      RUNE_MAP.put("deadly-reach-c", new Rune("deadly-reach-c", "Scattered Blows", "The third strike is replaced with an attack that will hit up to 6 nearby enemies within 15 yards for 170% weapon damage as Lightning."));
      RUNE_MAP.put("deadly-reach-d", new Rune("deadly-reach-d", "Strike from Beyond", "Critical Hits generate an additional 10 Spirit for each target hit."));
      RUNE_MAP.put("deadly-reach-e", new Rune("deadly-reach-e", "Keen Eye", "The third strike increases your Armor by 50% for 4 seconds."));
      SKILL_MAP.put("crippling-wave", new Skill("crippling-wave", "monk_cripplingwave", "Crippling Wave"));
      RUNE_MAP.put("crippling-wave-a", new Rune("crippling-wave-a", "Mangle", "Increase damage to 143% weapon damage."));
      RUNE_MAP.put("crippling-wave-b", new Rune("crippling-wave-b", "Tsunami", "The range of Crippling Wave's third strike is increased to 17 yards and the effect of the movement speed reduction is increased to 60%."));
      RUNE_MAP.put("crippling-wave-c", new Rune("crippling-wave-c", "Concussion", "Enemies hit by Crippling Wave inflict 20% less damage for 3 seconds."));
      RUNE_MAP.put("crippling-wave-d", new Rune("crippling-wave-d", "Rising Tide", "Critical Hits generate an additional 5 Spirit for each target hit."));
      RUNE_MAP.put("crippling-wave-e", new Rune("crippling-wave-e", "Breaking Wave", "Enemies hit by Crippling Wave take 10% additional damage from all attacks for 3 seconds."));
      SKILL_MAP.put("way-of-the-hundred-fists", new Skill("way-of-the-hundred-fists", "monk_wayofthehundredfists", "Way of the Hundred Fists"));
      RUNE_MAP.put("way-of-the-hundred-fists-a", new Rune("way-of-the-hundred-fists-a", "Fists of Fury", "Affected targets will take an additional 100% weapon damage as Holy over 5 seconds. Also adds a short dash to the first strike."));
      RUNE_MAP.put("way-of-the-hundred-fists-b", new Rune("way-of-the-hundred-fists-b", "Hands of Lightning", "Increases the number of hits in the second strike from 7 to 10."));
      RUNE_MAP.put("way-of-the-hundred-fists-c", new Rune("way-of-the-hundred-fists-c", "Blazing Fists", "Critical Hits increase your attack speed and movement speed by 5% for 5 seconds. This effect can stack up to 3 times."));
      RUNE_MAP.put("way-of-the-hundred-fists-d", new Rune("way-of-the-hundred-fists-d", "Spirited Salvo", "Every activation of the skill has a 15% chance to generate 15 additional Spirit."));
      RUNE_MAP.put("way-of-the-hundred-fists-e", new Rune("way-of-the-hundred-fists-e", "Windforce Flurry", "The third strike generates a wave of wind that deals 250% weapon damage as Physical to enemies directly ahead of you."));
      SKILL_MAP.put("lashing-tail-kick", new Skill("lashing-tail-kick", "monk_lashingtailkick", "Lashing Tail Kick"));
      RUNE_MAP.put("lashing-tail-kick-a", new Rune("lashing-tail-kick-a", "Vulture Claw Kick", "Release a torrent of fire that burns nearby enemies for 517% weapon damage as Fire and causes Knockback."));
      RUNE_MAP.put("lashing-tail-kick-b", new Rune("lashing-tail-kick-b", "Spinning Flame Kick", "Hurl a column of fire that burns through enemies, causing 588% weapon damage as Fire to each enemy it strikes."));
      RUNE_MAP.put("lashing-tail-kick-c", new Rune("lashing-tail-kick-c", "Hand of Ytar", "Attack enemies at long range, slowing the movement speed of affected targets by 80% for 2 seconds."));
      RUNE_MAP.put("lashing-tail-kick-d", new Rune("lashing-tail-kick-d", "Sweeping Armada", "Increases Knockback distance by 150% and slows the movement speed of struck enemies by 60% for 2 seconds."));
      RUNE_MAP.put("lashing-tail-kick-e", new Rune("lashing-tail-kick-e", "Scorpion Sting", "Enemies have a 50% chance to be stunned for 1.5 seconds instead of being knocked back."));
      SKILL_MAP.put("tempest-rush", new Skill("tempest-rush", "monk_tempestrush", "Tempest Rush"));
      RUNE_MAP.put("tempest-rush-a", new Rune("tempest-rush-a", "Bluster", "Enemies knocked back have their damage reduced by 20% for the duration of the effect."));
      RUNE_MAP.put("tempest-rush-b", new Rune("tempest-rush-b", "Tailwind", "Increases the movement speed of Tempest Rush by 25%."));
      RUNE_MAP.put("tempest-rush-c", new Rune("tempest-rush-c", "Slipstream", "Reduces damage taken while running by 25%."));
      RUNE_MAP.put("tempest-rush-d", new Rune("tempest-rush-d", "Northern Breeze", "Reduces the channeling cost of Tempest Rush to 8 Spirit."));
      RUNE_MAP.put("tempest-rush-e", new Rune("tempest-rush-e", "Flurry", "Increases the potency of the hobbling effect, slowing enemy movement by 80%."));
      SKILL_MAP.put("wave-of-light", new Skill("wave-of-light", "monk_waveoflight", "Wave of Light"));
      RUNE_MAP.put("wave-of-light-a", new Rune("wave-of-light-a", "Wall of Light", "Increases damage of the initial strike to 1202% weapon damage as Holy."));
      RUNE_MAP.put("wave-of-light-b", new Rune("wave-of-light-b", "Explosive Light", "Release bursts of energy that deal 914% weapon damage as Holy to nearby enemies."));
      RUNE_MAP.put("wave-of-light-c", new Rune("wave-of-light-c", "Pillar of the Ancients", "Summon an ancient pillar that deals 595% weapon damage followed by an additional 595% weapon damage after 2 seconds."));
      RUNE_MAP.put("wave-of-light-d", new Rune("wave-of-light-d", "Empowered Wave", "Reduces the cost of Wave of Light to 40 Spirit."));
      RUNE_MAP.put("wave-of-light-e", new Rune("wave-of-light-e", "Blinding Light", "Critical Hits Stun enemies for 3 seconds."));
      SKILL_MAP.put("blinding-flash", new Skill("blinding-flash", "monk_blindingflash", "Blinding Flash"));
      RUNE_MAP.put("blinding-flash-a", new Rune("blinding-flash-a", "Faith in the Light", "For 3 seconds after using Blinding Flash, all of your attacks are empowered to deal 30% additional weapon damage as Holy."));
      RUNE_MAP.put("blinding-flash-b", new Rune("blinding-flash-b", "Blinding Echo", "6 seconds after using Blinding Flash, a second flash of light will blind enemies within 20 yards for 0.5 seconds."));
      RUNE_MAP.put("blinding-flash-c", new Rune("blinding-flash-c", "Blinded and Confused", "Blinded enemies have a 25% chance to attack each other."));
      RUNE_MAP.put("blinding-flash-d", new Rune("blinding-flash-d", "Self Reflection", "Increases the duration enemies are blinded to 4 seconds."));
      RUNE_MAP.put("blinding-flash-e", new Rune("blinding-flash-e", "Searing Light", "Increases the chance elite enemies will miss attacks to 60%."));
      SKILL_MAP.put("breath-of-heaven", new Skill("breath-of-heaven", "monk_breathofheaven", "Breath of Heaven"));
      RUNE_MAP.put("breath-of-heaven-a", new Rune("breath-of-heaven-a", "Circle of Scorn", "Breath of Heaven also sears enemies for 80% weapon damage as Holy."));
      RUNE_MAP.put("breath-of-heaven-b", new Rune("breath-of-heaven-b", "Circle of Life", "Increases the healing power of Breath of Heaven to 8063 - 9675 Life."));
      RUNE_MAP.put("breath-of-heaven-c", new Rune("breath-of-heaven-c", "Blazing Wrath", "Breath of Heaven increases the damage of your attacks by 15% for 45 seconds."));
      RUNE_MAP.put("breath-of-heaven-d", new Rune("breath-of-heaven-d", "Infused with Light", "Gain 8 additional Spirit from Spirit generating attacks for 5 seconds after using Breath of Heaven."));
      RUNE_MAP.put("breath-of-heaven-e", new Rune("breath-of-heaven-e", "Penitent Flame", "Enemies exposed to Breath of Heaven run away in Fear for 1.5 seconds."));
      SKILL_MAP.put("serenity", new Skill("serenity", "monk_serenity", "Serenity"));
      RUNE_MAP.put("serenity-a", new Rune("serenity-a", "Peaceful Repose", "When activated, Serenity heals you for 6202 - 7752 Life."));
      RUNE_MAP.put("serenity-b", new Rune("serenity-b", "Instant Karma", "While Serenity is active, 50% of all projectiles and melee attacks are reflected back at the attacker."));
      RUNE_MAP.put("serenity-c", new Rune("serenity-c", "Ascension", "Increases the duration of Serenity to 4 seconds."));
      RUNE_MAP.put("serenity-d", new Rune("serenity-d", "Tranquility", "Extends the protective shield to allies within 45 yards for 1 seconds, and makes them immune to control impairing effects like Slow and Frozen."));
      RUNE_MAP.put("serenity-e", new Rune("serenity-e", "Reap What Is Sown", "When Serenity ends, the shield explodes, dealing 30% of the damage absorbed by Serenity as Holy damage to enemies within 20 yards. The damage to each enemy cannot exceed 100% of your maximum Life."));
      SKILL_MAP.put("inner-sanctuary", new Skill("inner-sanctuary", "monk_innersanctuary", "Inner Sanctuary"));
      RUNE_MAP.put("inner-sanctuary-a", new Rune("inner-sanctuary-a", "Forbidden Palace", "You and your allies standing in the area of effect of Inner Sanctuary deal 15% additional damage."));
      RUNE_MAP.put("inner-sanctuary-b", new Rune("inner-sanctuary-b", "Consecration", "Increases the duration of Inner Sanctuary to 7 seconds."));
      RUNE_MAP.put("inner-sanctuary-c", new Rune("inner-sanctuary-c", "Circle of Protection", "You and your allies standing in the area of effect of Inner Sanctuary take 35% less damage."));
      RUNE_MAP.put("inner-sanctuary-d", new Rune("inner-sanctuary-d", "Safe Haven", "You and your allies standing in the area of effect of Inner Sanctuary regenerate 1550 Life per second."));
      RUNE_MAP.put("inner-sanctuary-e", new Rune("inner-sanctuary-e", "Sanctified Ground", "When Inner Sanctuary expires, it becomes sanctified ground for 6 seconds, slowing the movement of all enemies that move through it by 80%."));
      SKILL_MAP.put("dashing-strike", new Skill("dashing-strike", "monk_dashingstrike", "Dashing Strike"));
      RUNE_MAP.put("dashing-strike-a", new Rune("dashing-strike-a", "Soaring Skull", "Launch yourself through the air and slow all enemies along your path by 60% for 2 seconds."));
      RUNE_MAP.put("dashing-strike-b", new Rune("dashing-strike-b", "Way of the Falling Star", "After striking an enemy, your movement speed is increased 25% for 3 seconds."));
      RUNE_MAP.put("dashing-strike-c", new Rune("dashing-strike-c", "Blinding Speed", "Receive a 20% increased chance to Dodge for 3 seconds."));
      RUNE_MAP.put("dashing-strike-d", new Rune("dashing-strike-d", "Quicksilver", "Reduces the cost of Dashing Strike to 10 Spirit."));
      RUNE_MAP.put("dashing-strike-e", new Rune("dashing-strike-e", "Flying Side Kick", "Perform a flying kick that has a 60% chance to Stun your target for 1.5 seconds."));
      SKILL_MAP.put("exploding-palm", new Skill("exploding-palm", "monk_explodingpalm", "Exploding Palm"));
      RUNE_MAP.put("exploding-palm-a", new Rune("exploding-palm-a", "Impending Doom", "Increases the duration of the Bleed effect to deal 745% weapon damage as Physical over 15 seconds."));
      RUNE_MAP.put("exploding-palm-b", new Rune("exploding-palm-b", "Creeping Demise", "Also reduces your target's movement speed by 80% for 3 seconds."));
      RUNE_MAP.put("exploding-palm-c", new Rune("exploding-palm-c", "The Flesh is Weak", "Also causes the target to take 12% additional damage for 9 seconds."));
      RUNE_MAP.put("exploding-palm-d", new Rune("exploding-palm-d", "Strong Spirit", "If the target explodes after bleeding, gain 5 Spirit for each enemy caught in the blast."));
      RUNE_MAP.put("exploding-palm-e", new Rune("exploding-palm-e", "Essence Burn", "Instead of bleeding, the target will burn for 745% weapon damage as Fire over 9 seconds. If the target dies while burning, it explodes causing all nearby enemies to burn for 60% weapon damage as Fire over 3 seconds. This effect can happen multiple times."));
      SKILL_MAP.put("sweeping-wind", new Skill("sweeping-wind", "monk_sweepingwind", "Sweeping Wind"));
      RUNE_MAP.put("sweeping-wind-a", new Rune("sweeping-wind-a", "Blade Storm", "Intensify the vortex, increasing the damage per stack to 26% weapon damage. This increases the damage with 3 stacks to 78% weapon damage."));
      RUNE_MAP.put("sweeping-wind-b", new Rune("sweeping-wind-b", "Fire Storm", "Increases the radius of the vortex to 14 yards and changes the damage dealt to Fire."));
      RUNE_MAP.put("sweeping-wind-c", new Rune("sweeping-wind-c", "Cyclone", "While your vortex is at the maximum stack count, Critical Hits have a chance to spawn a lightning tornado that periodically electrocutes nearby enemies for 26% weapon damage as Lightning. Each spawned lightning tornado lasts 3 seconds."));
      RUNE_MAP.put("sweeping-wind-d", new Rune("sweeping-wind-d", "Inner Storm", "As long as your vortex is at the maximum stack count, you gain 3 Spirit per second."));
      RUNE_MAP.put("sweeping-wind-e", new Rune("sweeping-wind-e", "Master of Wind", "Increases the duration of the vortex to 20 seconds."));
      SKILL_MAP.put("cyclone-strike", new Skill("cyclone-strike", "monk_cyclonestrike", "Cyclone Strike"));
      RUNE_MAP.put("cyclone-strike-a", new Rune("cyclone-strike-a", "Sunburst", "Changes the blast into an explosion of fire that has a 35% chance to Fear enemies for 1.5 seconds."));
      RUNE_MAP.put("cyclone-strike-b", new Rune("cyclone-strike-b", "Implosion", "Increases the distance enemies will be pulled towards you to 34 yards."));
      RUNE_MAP.put("cyclone-strike-c", new Rune("cyclone-strike-c", "Soothing Breeze", "Cyclone Strike heals you and all allies within 24 yards for 1240 Life."));
      RUNE_MAP.put("cyclone-strike-d", new Rune("cyclone-strike-d", "Eye of the Storm", "Reduces the Spirit cost of Cyclone Strike to 30 Spirit."));
      RUNE_MAP.put("cyclone-strike-e", new Rune("cyclone-strike-e", "Wall of Wind", "After using Cyclone Strike, gain a 20% chance to dodge attacks for 3 seconds."));
      SKILL_MAP.put("sevensided-strike", new Skill("sevensided-strike", "monk_sevensidedstrike", "Seven-Sided Strike"));
      RUNE_MAP.put("sevensided-strike-a", new Rune("sevensided-strike-a", "Sudden Assault", "Teleport to the target, increasing damage done to 2310% weapon damage over 7 strikes."));
      RUNE_MAP.put("sevensided-strike-b", new Rune("sevensided-strike-b", "Several-Sided Strike", "Increases the number of strikes to 9."));
      RUNE_MAP.put("sevensided-strike-c", new Rune("sevensided-strike-c", "Pandemonium", "Enemies hit by Seven-Sided Strike have a 25% chance to be stunned for 7 seconds by each hit."));
      RUNE_MAP.put("sevensided-strike-d", new Rune("sevensided-strike-d", "Sustained Attack", "Reduces the cooldown of Seven-Sided Strike to 23 seconds."));
      RUNE_MAP.put("sevensided-strike-e", new Rune("sevensided-strike-e", "Fulminating Onslaught", "Each strike explodes, dealing 254% weapon damage as Holy in a 7 yard radius around the target."));
      SKILL_MAP.put("mystic-ally", new Skill("mystic-ally", "monk_mystically", "Mystic Ally"));
      RUNE_MAP.put("mystic-ally-a", new Rune("mystic-ally-a", "Fire Ally", "Imbue the ally with the essence of fire. The ally gains the ability to unleash a flaming kick for 80% weapon damage as Fire plus an additional 40% of your weapon damage per second as Fire for 2 seconds to all enemies in a straight line."));
      RUNE_MAP.put("mystic-ally-b", new Rune("mystic-ally-b", "Water Ally", "Imbue the ally with the essence of water. The ally gains the ability to perform a wave attack that deals 120% of your weapon damage as Physical and slows the movement of affected targets by 60% for 2 seconds."));
      RUNE_MAP.put("mystic-ally-c", new Rune("mystic-ally-c", "Earth Ally", "Imbue the ally with the essence of earth. Maximum Life for you and the ally is increased by 10%. The ally also gains the ability to create a wave of earth, dealing 60% of your weapon damage as Physical to a single enemy and forcing that enemy to attack the ally for 3 seconds."));
      RUNE_MAP.put("mystic-ally-d", new Rune("mystic-ally-d", "Air Ally", "Imbue the ally with the essence of air. Every attack made by the ally has a 2% chance to generate 100 Spirit for you. In addition, the ally is surrounded in a torrent of wind that deals 10% of your weapon damage per second as Physical to all nearby enemies."));
      RUNE_MAP.put("mystic-ally-e", new Rune("mystic-ally-e", "Eternal Ally", "Imbue the ally with the essence of life. When the ally dies, it has a 100% chance to be reborn after 3 seconds. In addition, the physical damage of the ally's basic attack is increased to 60% of your weapon damage per swing."));
      SKILL_MAP.put("mantra-of-evasion", new Skill("mantra-of-evasion", "monk_mantraofevasion", "Mantra of Evasion"));
      RUNE_MAP.put("mantra-of-evasion-a", new Rune("mantra-of-evasion-a", "Backlash", "Successfully dodging an attack has a chance to create a burst of flame dealing 35% weapon damage as Fire to all nearby enemies."));
      RUNE_MAP.put("mantra-of-evasion-b", new Rune("mantra-of-evasion-b", "Perseverance", "Mantra of Evasion also reduces the duration of all control impairing effects like Slow or Frozen by 20%."));
      RUNE_MAP.put("mantra-of-evasion-c", new Rune("mantra-of-evasion-c", "Hard Target", "Mantra of Evasion also increases Armor by 20%."));
      RUNE_MAP.put("mantra-of-evasion-d", new Rune("mantra-of-evasion-d", "Wind through the Reeds", "Mantra of Evasion also increases movement speed by 8%."));
      RUNE_MAP.put("mantra-of-evasion-e", new Rune("mantra-of-evasion-e", "Divine Protection", "When you or an ally under the effect of Mantra of Evasion is reduced below 25% Life, a shield of protection forms around that target, reducing damage taken by 80% for 3 seconds. Each target can be protected at most once every 90 seconds by this effect."));
      SKILL_MAP.put("mantra-of-retribution", new Skill("mantra-of-retribution", "monk_mantraofretribution", "Mantra of Retribution"));
      RUNE_MAP.put("mantra-of-retribution-a", new Rune("mantra-of-retribution-a", "Retaliation", "Increases the amount of damage reflected by the Mantra to 60%. The Mantra will now reflect ranged damage as well as melee damage."));
      RUNE_MAP.put("mantra-of-retribution-b", new Rune("mantra-of-retribution-b", "Transgression", "Increases attack speed for you and your allies by 8%."));
      RUNE_MAP.put("mantra-of-retribution-c", new Rune("mantra-of-retribution-c", "Indignation", "When taking damage from the Mantra of Retribution, enemies have a 10% chance to be stunned for 2 seconds."));
      RUNE_MAP.put("mantra-of-retribution-d", new Rune("mantra-of-retribution-d", "Against All Odds", "When reflecting damage done to you, Mantra of Retribution has a chance to restore 3 Spirit."));
      RUNE_MAP.put("mantra-of-retribution-e", new Rune("mantra-of-retribution-e", "Collateral Damage", "An attacker that is damaged by Mantra of Retribution has a 30% chance to suffer a feedback blast, dealing 45% weapon damage as Holy to itself and nearby enemies."));
      SKILL_MAP.put("mantra-of-healing", new Skill("mantra-of-healing", "monk_mantraofhealing", "Mantra of Healing"));
      RUNE_MAP.put("mantra-of-healing-a", new Rune("mantra-of-healing-a", "Sustenance", "Increases the Life regeneration granted by Mantra of Healing to 1240 Life per second."));
      RUNE_MAP.put("mantra-of-healing-b", new Rune("mantra-of-healing-b", "Boon of Inspiration", "Mantra of Healing also heals 186 Life when hitting an enemy."));
      RUNE_MAP.put("mantra-of-healing-c", new Rune("mantra-of-healing-c", "Heavenly Body", "Mantra of Healing also increases Vitality by 10%."));
      RUNE_MAP.put("mantra-of-healing-d", new Rune("mantra-of-healing-d", "Circular Breathing", "Mantra of Healing also regenerates 3 Spirit per second."));
      RUNE_MAP.put("mantra-of-healing-e", new Rune("mantra-of-healing-e", "Time of Need", "Mantra of Healing also increases resistances to all damage types by 20%."));
      SKILL_MAP.put("mantra-of-conviction", new Skill("mantra-of-conviction", "monk_mantraofconviction", "Mantra of Conviction"));
      RUNE_MAP.put("mantra-of-conviction-a", new Rune("mantra-of-conviction-a", "Overawe", "Increases the strength of Mantra of Conviction so that enemies take 24% additional damage and 48% for the first 3 seconds."));
      RUNE_MAP.put("mantra-of-conviction-b", new Rune("mantra-of-conviction-b", "Submission", "Enemies affected by Mantra of Conviction take 12% weapon damage per second as Holy."));
      RUNE_MAP.put("mantra-of-conviction-c", new Rune("mantra-of-conviction-c", "Dishearten", "Slows the movement of enemies within 20 yards by 30%."));
      RUNE_MAP.put("mantra-of-conviction-d", new Rune("mantra-of-conviction-d", "Reclamation", "You and your allies have a 30% chance to be healed for 279 - 341 Life when using melee attacks on an enemy under the effects of Mantra of Conviction."));
      RUNE_MAP.put("mantra-of-conviction-e", new Rune("mantra-of-conviction-e", "Intimidation", "Enemies affected by Mantra of Conviction deal 10% less damage."));
      SKILL_MAP.put("hungering-arrow", new Skill("hungering-arrow", "demonhunter_hungeringarrow", "Hungering Arrow"));
      RUNE_MAP.put("hungering-arrow-a", new Rune("hungering-arrow-a", "Cinder Arrow", "Light the arrow on fire, dealing 35% additional weapon damage as Fire over 3 seconds."));
      RUNE_MAP.put("hungering-arrow-b", new Rune("hungering-arrow-b", "Shatter Shot", "If the arrow successfully pierces the first target, the arrow splits into 3 arrows."));
      RUNE_MAP.put("hungering-arrow-c", new Rune("hungering-arrow-c", "Devouring Arrow", "Each consecutive pierce increases the damage of the arrow by 70%."));
      RUNE_MAP.put("hungering-arrow-d", new Rune("hungering-arrow-d", "Puncturing Arrow", "Increase the chance for the arrow to pierce to 50%."));
      RUNE_MAP.put("hungering-arrow-e", new Rune("hungering-arrow-e", "Spray of Teeth", "Successful Critical Hits cause a burst of bone to explode from the target, dealing 50% weapon damage to enemies in that area."));
      SKILL_MAP.put("entangling-shot", new Skill("entangling-shot", "demonhunter_entanglingshot", "Entangling Shot"));
      RUNE_MAP.put("entangling-shot-a", new Rune("entangling-shot-a", "Heavy Burden", "Increase the movement slow duration to 4 seconds."));
      RUNE_MAP.put("entangling-shot-b", new Rune("entangling-shot-b", "Chain Gang", "Hit up to 4 targets."));
      RUNE_MAP.put("entangling-shot-c", new Rune("entangling-shot-c", "Shock Collar", "Strike targets with electrified chains that do an additional 70% weapon damage as Lightning over 2 seconds."));
      RUNE_MAP.put("entangling-shot-d", new Rune("entangling-shot-d", "Justice is Served", "Increase the Hatred generated to 6 per shot."));
      RUNE_MAP.put("entangling-shot-e", new Rune("entangling-shot-e", "Bounty Hunter", "Gain 6% of the damage dealt as Life."));
      SKILL_MAP.put("bola-shot", new Skill("bola-shot", "demonhunter_bolashot", "Bola Shot"));
      RUNE_MAP.put("bola-shot-a", new Rune("bola-shot-a", "Volatile Explosives", "Increase the explosion radius to 20 yards."));
      RUNE_MAP.put("bola-shot-b", new Rune("bola-shot-b", "Acid Strike", "Shoot 3 bolas that each deal 160% weapon damage as Poison. The bolas no longer explode for area damage to nearby targets."));
      RUNE_MAP.put("bola-shot-c", new Rune("bola-shot-c", "Thunder Ball", "When the bola explodes, it deals 160% weapon damage as Lightning and has a 35% chance to Stun the primary target for 1.5 seconds."));
      RUNE_MAP.put("bola-shot-d", new Rune("bola-shot-d", "Bitter Pill", "When the bola explodes, you have a 15% chance to gain 2 Discipline."));
      RUNE_MAP.put("bola-shot-e", new Rune("bola-shot-e", "Imminent Doom", "Augment the bola to deal 216% weapon damage as Arcane to the target and 149% weapon damage as Arcane to all other targets within 14 yards, but increases the explosion delay to 2 seconds."));
      SKILL_MAP.put("grenades", new Skill("grenades", "demonhunter_grenades", "Grenades"));
      RUNE_MAP.put("grenades-a", new Rune("grenades-a", "Gas Grenades", "Throw gas grenades that explode for 95% weapon damage as Poison and leave a cloud that deals an additional 25% weapon damage per second as Poison for 3 seconds to enemies who stand in the area."));
      RUNE_MAP.put("grenades-b", new Rune("grenades-b", "Cluster Grenades", "Throw cluster grenades that deal 112% weapon damage as Fire over an 8 yard radius."));
      RUNE_MAP.put("grenades-c", new Rune("grenades-c", "Fire Bomb", "Throw a single grenade that deals 124% weapon damage as Fire."));
      RUNE_MAP.put("grenades-d", new Rune("grenades-d", "Tinkerer", "Increases Hatred generation to 6 Hatred."));
      RUNE_MAP.put("grenades-e", new Rune("grenades-e", "Stun Grenades", "Hurl grenades that have a 25% chance to Stun enemies for 1.5 seconds."));
      SKILL_MAP.put("impale", new Skill("impale", "demonhunter_impale", "Impale"));
      RUNE_MAP.put("impale-a", new Rune("impale-a", "Overpenetration", "The knife will pierce through all enemies in a straight line."));
      RUNE_MAP.put("impale-b", new Rune("impale-b", "Impact", "Impale causes Knockback and has a 65% chance to Stun enemies for 1.5 seconds."));
      RUNE_MAP.put("impale-c", new Rune("impale-c", "Chemical Burn", "Your target will also Bleed for 220% weapon damage as Physical over 2 seconds."));
      RUNE_MAP.put("impale-d", new Rune("impale-d", "Awareness", "After the initial throw, release multiple blades centered on you, dealing 75% weapon damage to all enemies within 10 yards."));
      RUNE_MAP.put("impale-e", new Rune("impale-e", "Grievous Wounds", "Critical Hits cause 100% additional damage."));
      SKILL_MAP.put("rapid-fire", new Skill("rapid-fire", "demonhunter_rapidfire", "Rapid Fire"));
      RUNE_MAP.put("rapid-fire-a", new Rune("rapid-fire-a", "Bombardment", "Rapidly fire grenades that explode for 414% weapon damage as Fire to all enemies within a 8 yard radius."));
      RUNE_MAP.put("rapid-fire-b", new Rune("rapid-fire-b", "High Velocity", "Fire poison arrows that have a 40% chance to pierce through enemies."));
      RUNE_MAP.put("rapid-fire-c", new Rune("rapid-fire-c", "Fire Support", "While channeling Rapid Fire, launch 2 homing rockets every second. Each rocket deals 145% weapon damage as Physical to nearby targets."));
      RUNE_MAP.put("rapid-fire-d", new Rune("rapid-fire-d", "Withering Fire", "Reduces the initial Hatred cost to 10, and ignites your arrows, causing them to deal Fire damage."));
      RUNE_MAP.put("rapid-fire-e", new Rune("rapid-fire-e", "Web Shot", "Slows the movement of affected targets by 80% for 1 seconds."));
      SKILL_MAP.put("chakram", new Skill("chakram", "demonhunter_chakram", "Chakram"));
      RUNE_MAP.put("chakram-a", new Rune("chakram-a", "Twin Chakrams", "A second Chakram mirrors the first.  Each Chakram deals 114% weapon damage as Physical."));
      RUNE_MAP.put("chakram-b", new Rune("chakram-b", "Boomerang", "The Chakram path turns into a loop, dealing 230% weapon damage as Lightning to enemies along the path."));
      RUNE_MAP.put("chakram-c", new Rune("chakram-c", "Serpentine", "The Chakram follows a slow curve, dealing 230% weapon damage as Poison to enemies along the path."));
      RUNE_MAP.put("chakram-d", new Rune("chakram-d", "Razor Disk", "The Chakram spirals out from the targeted location dealing 187% weapon damage as Arcane to enemies along the path."));
      RUNE_MAP.put("chakram-e", new Rune("chakram-e", "Shuriken Cloud", "Surround yourself with spinning Chakrams for 120 seconds, dealing 34% weapon damage per second as Physical to nearby enemies."));
      SKILL_MAP.put("elemental-arrow", new Skill("elemental-arrow", "demonhunter_elementalarrow", "Elemental Arrow"));
      RUNE_MAP.put("elemental-arrow-a", new Rune("elemental-arrow-a", "Frost Arrow", "Fire a frost arrow that splits into multiple arrows after hitting its target, dealing 170% weapon damage as Cold. Affected enemies have their movement speed slowed by 60% for 1 seconds."));
      RUNE_MAP.put("elemental-arrow-b", new Rune("elemental-arrow-b", "Ball Lightning", "Fire a slow-moving arrow that electrocutes enemies along its path for 155% weapon damage as Lightning."));
      RUNE_MAP.put("elemental-arrow-c", new Rune("elemental-arrow-c", "Screaming Skull", "Grants a 40% chance to shoot a skull that will Fear affected enemies for 1.5 seconds."));
      RUNE_MAP.put("elemental-arrow-d", new Rune("elemental-arrow-d", "Nether Tentacles", "Shadow tentacles deal 155% weapon damage to enemies along its path and return 3% of damage dealt as Life for you."));
      RUNE_MAP.put("elemental-arrow-e", new Rune("elemental-arrow-e", "Lightning Bolts", "Fire electrified bolts that Stun enemies for 1.5 seconds on a Critical Hit."));
      SKILL_MAP.put("caltrops", new Skill("caltrops", "demonhunter_caltrops", "Caltrops"));
      RUNE_MAP.put("caltrops-a", new Rune("caltrops-a", "Jagged Spikes", "Enemies in the area also take 270% weapon damage as Physical over 6 seconds."));
      RUNE_MAP.put("caltrops-b", new Rune("caltrops-b", "Hooked Spines", "Increase the slowing amount to 80%."));
      RUNE_MAP.put("caltrops-c", new Rune("caltrops-c", "Torturous Ground", "When the trap is sprung, all enemies in the area are immobilized for 2 seconds."));
      RUNE_MAP.put("caltrops-d", new Rune("caltrops-d", "Carved Stakes", "Reduces the cost of Caltrops to 4 Discipline."));
      RUNE_MAP.put("caltrops-e", new Rune("caltrops-e", "Bait the Trap", "Become empowered while standing in the area of effect, gaining an additional 10% Critical Hit Chance with all attacks."));
      SKILL_MAP.put("smoke-screen", new Skill("smoke-screen", "demonhunter_smokescreen", "Smoke Screen"));
      RUNE_MAP.put("smoke-screen-a", new Rune("smoke-screen-a", "Choking Gas", "Leave behind a cloud of gas that deals 700% weapon damage as Physical to enemies in the area over 5 seconds."));
      RUNE_MAP.put("smoke-screen-b", new Rune("smoke-screen-b", "Lingering Fog", "Increase the duration of the effect to 1.5 seconds."));
      RUNE_MAP.put("smoke-screen-c", new Rune("smoke-screen-c", "Breathe Deep", "While invisible you gain 12 Hatred per second."));
      RUNE_MAP.put("smoke-screen-d", new Rune("smoke-screen-d", "Special Recipe", "Reduce the cost to 12 Discipline."));
      RUNE_MAP.put("smoke-screen-e", new Rune("smoke-screen-e", "Displacement", "Gain 35% movement speed when activated."));
      SKILL_MAP.put("shadow-power", new Skill("shadow-power", "demonhunter_shadowpower", "Shadow Power"));
      RUNE_MAP.put("shadow-power-a", new Rune("shadow-power-a", "Night Bane", "Gain an additional 3 Hatred per second while Shadow Power is active."));
      RUNE_MAP.put("shadow-power-b", new Rune("shadow-power-b", "Shadow Glide", "Gain 30% bonus to movement speed while Shadow Power is active."));
      RUNE_MAP.put("shadow-power-c", new Rune("shadow-power-c", "Gloom", "Reduce incoming damage by 35% while Shadow Power is active."));
      RUNE_MAP.put("shadow-power-d", new Rune("shadow-power-d", "Well of Darkness", "Decreases the Discipline cost to 12."));
      RUNE_MAP.put("shadow-power-e", new Rune("shadow-power-e", "Blood Moon", "Increases damage done as Life to 25%."));
      SKILL_MAP.put("vault", new Skill("vault", "demonhunter_vault", "Vault"));
      RUNE_MAP.put("vault-a", new Rune("vault-a", "Trail of Cinders", "Leave a trail of fire in your wake that inflicts 300% weapon damage as Fire over 3 seconds."));
      RUNE_MAP.put("vault-b", new Rune("vault-b", "Acrobatics", "Removes the Discipline cost but adds a 10 second cooldown."));
      RUNE_MAP.put("vault-c", new Rune("vault-c", "Action Shot", "As you travel, shoot arrows for 75% weapon damage at nearby targets."));
      RUNE_MAP.put("vault-d", new Rune("vault-d", "Tumble", "After using Vault, your next Vault within 6 seconds has its Discipline cost reduced by 50%."));
      RUNE_MAP.put("vault-e", new Rune("vault-e", "Rattling Roll", "All enemies within 8 yards of your destination are knocked back and stunned for 1.5 seconds."));
      SKILL_MAP.put("preparation", new Skill("preparation", "demonhunter_preparation", "Preparation"));
      RUNE_MAP.put("preparation-a", new Rune("preparation-a", "Punishment", "Restore all Hatred for 25 Discipline. Preparation has no cooldown."));
      RUNE_MAP.put("preparation-b", new Rune("preparation-b", "Invigoration", "Increase maximum Discipline by 10 for 5 seconds when using Preparation."));
      RUNE_MAP.put("preparation-c", new Rune("preparation-c", "Focused Mind", "Gain 45 Discipline over 15 seconds instead of restoring it immediately."));
      RUNE_MAP.put("preparation-d", new Rune("preparation-d", "Battle Scars", "Gain 60% Life after using Preparation."));
      RUNE_MAP.put("preparation-e", new Rune("preparation-e", "Backup Plan", "There is a 30% chance that Preparation's cooldown will not be triggered."));
      SKILL_MAP.put("companion", new Skill("companion", "demonhunter_companion", "Companion"));
      RUNE_MAP.put("companion-a", new Rune("companion-a", "Spider Companion", "Summon a spider instead of a raven. The spider's attacks also Slow the movement of enemies by 60% for 2 seconds."));
      RUNE_MAP.put("companion-b", new Rune("companion-b", "Boar Companion", "Summon a boar instead of a raven. The boar increases your Life regeneration by 310 per second. In addition the boar increases your resistances to all damage types by 15%."));
      RUNE_MAP.put("companion-c", new Rune("companion-c", "Wolf Companion", "Summon a wolf instead of a raven. The wolf attacks for 94% of your weapon damage as Physical."));
      RUNE_MAP.put("companion-d", new Rune("companion-d", "Bat Companion", "Summon a bat instead of a raven. The bat grants you 3 Hatred per second."));
      RUNE_MAP.put("companion-e", new Rune("companion-e", "Ferret Companion", "Summon ferrets instead of a raven. The ferrets collect gold for you and increase gold found on monsters by 10%."));
      SKILL_MAP.put("marked-for-death", new Skill("marked-for-death", "demonhunter_markedfordeath", "Marked for Death"));
      RUNE_MAP.put("marked-for-death-a", new Rune("marked-for-death-a", "Grim Reaper", "An additional 12% of damage done to the target is also divided among all enemies within 20 yards."));
      RUNE_MAP.put("marked-for-death-b", new Rune("marked-for-death-b", "Contagion", "When the target is killed, the ability spreads to 2 other nearby targets. This effect can chain repeatedly."));
      RUNE_MAP.put("marked-for-death-c", new Rune("marked-for-death-c", "Valley of Death", "Mark an area on the ground 12 yards wide for 15 seconds.  Enemies in the area take 12% additional damage."));
      RUNE_MAP.put("marked-for-death-d", new Rune("marked-for-death-d", "Mortal Enemy", "Attacks you make against the marked target generate 3 Hatred."));
      RUNE_MAP.put("marked-for-death-e", new Rune("marked-for-death-e", "Death Toll", "Heal attackers for 1% of the damage done to the marked target."));
      SKILL_MAP.put("evasive-fire", new Skill("evasive-fire", "demonhunter_evasivefire", "Evasive Fire"));
      RUNE_MAP.put("evasive-fire-a", new Rune("evasive-fire-a", "Hardened", "Instead of backflipping, your Armor is increased by 25% for 3 seconds. This does not cost Discipline."));
      RUNE_MAP.put("evasive-fire-b", new Rune("evasive-fire-b", "Covering Fire", "Shoot a spread of bolts that hit up to 3 targets for 130% weapon damage each."));
      RUNE_MAP.put("evasive-fire-c", new Rune("evasive-fire-c", "Parting Gift", "Whenever a backflip is triggered, leave a poison bomb behind that explodes for 55% weapon damage as Poison in a 12 yard radius after 0.6 seconds. Turns Evasive Fire into Poison damage."));
      RUNE_MAP.put("evasive-fire-d", new Rune("evasive-fire-d", "Surge", "Reduces the cost of the backflip to 2 Discipline. Turns Evasive Fire into Lightning damage."));
      RUNE_MAP.put("evasive-fire-e", new Rune("evasive-fire-e", "Displace", "Increase the distance of the backflip to 30 yards."));
      SKILL_MAP.put("fan-of-knives", new Skill("fan-of-knives", "demonhunter_fanofknives", "Fan of Knives"));
      RUNE_MAP.put("fan-of-knives-a", new Rune("fan-of-knives-a", "Hail of Knives", "Increase the radius to damage all enemies within 20 yards."));
      RUNE_MAP.put("fan-of-knives-b", new Rune("fan-of-knives-b", "Assassin's Knives", "Throw long-range knives that deal 70% weapon damage to 5 additional targets."));
      RUNE_MAP.put("fan-of-knives-c", new Rune("fan-of-knives-c", "Fan of Daggers", "Imbue your knives with a 65% chance to Stun enemies for 2 seconds."));
      RUNE_MAP.put("fan-of-knives-d", new Rune("fan-of-knives-d", "Crippling Razors", "Increase the amount enemies are slowed to 80% for 2 seconds."));
      RUNE_MAP.put("fan-of-knives-e", new Rune("fan-of-knives-e", "Retaliate", "Surround yourself with whirling blades that deal 464% weapon damage to all enemies if you are struck in the next 10 seconds."));
      SKILL_MAP.put("spike-trap", new Skill("spike-trap", "demonhunter_spiketrap", "Spike Trap"));
      RUNE_MAP.put("spike-trap-a", new Rune("spike-trap-a", "Long Fuse", "Increases the arming time to 2 seconds but increases damage to 371% weapon damage."));
      RUNE_MAP.put("spike-trap-b", new Rune("spike-trap-b", "Echoing Blast", "Your Spike Traps can explode up to 3 times, dealing 275% weapon damage as Poison to all enemies within 8 yards each time."));
      RUNE_MAP.put("spike-trap-c", new Rune("spike-trap-c", "Sticky Trap", "Plant a bomb on an enemy rather than on the ground. If the target dies within 30 seconds, the bomb explodes dealing 404% weapon damage to all enemies within 16 yards."));
      RUNE_MAP.put("spike-trap-d", new Rune("spike-trap-d", "Scatter", "Simultaneously place all 3 traps."));
      RUNE_MAP.put("spike-trap-e", new Rune("spike-trap-e", "Lightning Rod", "When the trap is triggered it releases a pulse of lightning that will bounce to up to 3 enemies for 275% weapon damage as Lightning."));
      SKILL_MAP.put("sentry", new Skill("sentry", "demonhunter_sentry", "Sentry"));
      RUNE_MAP.put("sentry-a", new Rune("sentry-a", "Chain of Torment", "Create a tether between you and the Sentry that does 125% weapon damage every second to every enemy it touches."));
      RUNE_MAP.put("sentry-b", new Rune("sentry-b", "Vigilant Watcher", "Reduces the cooldown of Sentry to 6 seconds."));
      RUNE_MAP.put("sentry-c", new Rune("sentry-c", "Spitfire Turret", "The turret will also fire homing rockets aimed at random nearby targets for 30% weapon damage as Fire."));
      RUNE_MAP.put("sentry-d", new Rune("sentry-d", "Aid Station", "Heals nearby allies for 2.0% of their maximum Life per second."));
      RUNE_MAP.put("sentry-e", new Rune("sentry-e", "Guardian Turret", "The turret also creates a shield that reduces damage taken by allies by 15%."));
      SKILL_MAP.put("strafe", new Skill("strafe", "demonhunter_strafe", "Strafe"));
      RUNE_MAP.put("strafe-a", new Rune("strafe-a", "Demolition", "Throw out bouncy grenades that explode for 187% weapon damage to targets within 9 yards."));
      RUNE_MAP.put("strafe-b", new Rune("strafe-b", "Emberstrafe", "Leave a trail of fire in your wake that inflicts 65% weapon damage as Fire over 2 seconds."));
      RUNE_MAP.put("strafe-c", new Rune("strafe-c", "Rocket Storm", "In addition to regular firing, fire off homing rockets for 60% weapon damage as Fire."));
      RUNE_MAP.put("strafe-d", new Rune("strafe-d", "Drifting Shadow", "Movement speed increased to 100% of normal running speed while strafing."));
      RUNE_MAP.put("strafe-e", new Rune("strafe-e", "Stinging Steel", "Throw out knives rather than arrows that do an extra 100% damage on successful Critical Hits."));
      SKILL_MAP.put("multishot", new Skill("multishot", "demonhunter_multishot", "Multishot"));
      RUNE_MAP.put("multishot-a", new Rune("multishot-a", "Full Broadside", "Increase the damage of Multishot to 215% weapon damage."));
      RUNE_MAP.put("multishot-b", new Rune("multishot-b", "Burst Fire", "Every time you fire, generate a shock pulse that damages nearby enemies for 65% weapon damage as Arcane."));
      RUNE_MAP.put("multishot-c", new Rune("multishot-c", "Arsenal", "Every use also fires 3 rockets at nearby enemies that deal 60% weapon damage as Fire each."));
      RUNE_MAP.put("multishot-d", new Rune("multishot-d", "Fire at Will", "Cost reduced to 15 Hatred. Deals 165% weapon damage as Lightning."));
      RUNE_MAP.put("multishot-e", new Rune("multishot-e", "Suppression Fire", "Every enemy hit grants 1 Discipline."));
      SKILL_MAP.put("cluster-arrow", new Skill("cluster-arrow", "demonhunter_clusterarrow", "Cluster Arrow"));
      RUNE_MAP.put("cluster-arrow-a", new Rune("cluster-arrow-a", "Loaded for Bear", "Increases the damage of the explosion at the impact location to 304% weapon damage as Fire."));
      RUNE_MAP.put("cluster-arrow-b", new Rune("cluster-arrow-b", "Shooting Stars", "Instead of releasing grenades, shoots up to 3 rockets at nearby enemies dealing 175% weapon damage as Physical each."));
      RUNE_MAP.put("cluster-arrow-c", new Rune("cluster-arrow-c", "Cluster Bombs", "Launch the cluster through the air, dropping bombs in a straight line that each explode for 230% weapon damage as Fire."));
      RUNE_MAP.put("cluster-arrow-d", new Rune("cluster-arrow-d", "Maelstrom", "Instead of releasing grenades, the cluster releases shadow energy that deals 165% weapon damage as Physical to nearby enemies. You will gain 4% of the damage done as Life."));
      RUNE_MAP.put("cluster-arrow-e", new Rune("cluster-arrow-e", "Dazzling Arrow", "Enemies hit by grenades have a 55% chance to be stunned for 2 seconds and changes the damage to Physical."));
      SKILL_MAP.put("rain-of-vengeance", new Skill("rain-of-vengeance", "demonhunter_rainofvengeance", "Rain of Vengeance"));
      RUNE_MAP.put("rain-of-vengeance-a", new Rune("rain-of-vengeance-a", "Beastly Bombs", "Summon 20 Shadow Beasts to drop bombs on enemies, dealing 245% weapon damage each."));
      RUNE_MAP.put("rain-of-vengeance-b", new Rune("rain-of-vengeance-b", "Dark Cloud", "Launch a massive volley of guided arrows that rain down on enemies for 792% weapon damage over 12 seconds."));
      RUNE_MAP.put("rain-of-vengeance-c", new Rune("rain-of-vengeance-c", "Anathema", "Summon a Shadow Beast that drops grenades from the sky dealing 3300% weapon damage over 15 seconds."));
      RUNE_MAP.put("rain-of-vengeance-d", new Rune("rain-of-vengeance-d", "Flying Strike", "A group of 8 Shadow Beasts plummet from the sky at a targeted location dealing 100% weapon damage each and stunning enemies for 2 seconds."));
      RUNE_MAP.put("rain-of-vengeance-e", new Rune("rain-of-vengeance-e", "Stampede", "Summon a wave of 10 Shadow Beasts to tear across the ground, knocking back enemies and dealing 120% weapon damage each."));
      SKILL_MAP.put("magic-missile", new Skill("magic-missile", "wizard_magicmissile", "Magic Missile"));
      RUNE_MAP.put("magic-missile-a", new Rune("magic-missile-a", "Charged Blast", "Increases the damage of Magic Missile to 163% weapon damage as Arcane."));
      RUNE_MAP.put("magic-missile-b", new Rune("magic-missile-b", "Split", "Fire 3 missiles that each deal 56% weapon damage as Arcane."));
      RUNE_MAP.put("magic-missile-c", new Rune("magic-missile-c", "Penetrating Blast", "Missiles have a 100% chance to pierce through their target and hit additional enemies."));
      RUNE_MAP.put("magic-missile-d", new Rune("magic-missile-d", "Attunement", "Whenever Magic Missile hits a target you gain 4 Arcane Power."));
      RUNE_MAP.put("magic-missile-e", new Rune("magic-missile-e", "Seeker", "Missiles track the nearest enemy and their damage is increased to 138% weapon damage as Arcane."));
      SKILL_MAP.put("ray-of-frost", new Skill("ray-of-frost", "wizard_rayoffrost", "Ray of Frost"));
      RUNE_MAP.put("ray-of-frost-a", new Rune("ray-of-frost-a", "Snow Blast", "Using continuously on a single target increases damage over 1.5 seconds to inflict a maximum of 364% weapon damage as Cold."));
      RUNE_MAP.put("ray-of-frost-b", new Rune("ray-of-frost-b", "Sleet Storm", "Create a swirling storm around you, dealing 364% weapon damage as Cold to all enemies caught within it."));
      RUNE_MAP.put("ray-of-frost-c", new Rune("ray-of-frost-c", "Numb", "Increase the amount the target's movement is slowed to 80% for 4 seconds."));
      RUNE_MAP.put("ray-of-frost-d", new Rune("ray-of-frost-d", "Cold Blood", "Reduce casting cost to 10 Arcane Power."));
      RUNE_MAP.put("ray-of-frost-e", new Rune("ray-of-frost-e", "Black Ice", "Enemies killed with Ray of Frost leave behind a patch of ice that deals 504% weapon damage as Cold to enemies moving through it over 3 seconds."));
      SKILL_MAP.put("shock-pulse", new Skill("shock-pulse", "wizard_shockpulse", "Shock Pulse"));
      RUNE_MAP.put("shock-pulse-a", new Rune("shock-pulse-a", "Fire Bolts", "Cast bolts of fire that each deal 195% weapon damage as Fire."));
      RUNE_MAP.put("shock-pulse-b", new Rune("shock-pulse-b", "Living Lightning", "Conjure a being of lightning that drifts forward, electrocuting nearby enemies for 53% weapon damage as Lightning."));
      RUNE_MAP.put("shock-pulse-c", new Rune("shock-pulse-c", "Piercing Orb", "Merge the bolts in a single giant orb that oscillates forward dealing 105% weapon damage as Lightning to everything it hits."));
      RUNE_MAP.put("shock-pulse-d", new Rune("shock-pulse-d", "Lightning Affinity", "Every target hit by a pulse restores 3 Arcane Power."));
      RUNE_MAP.put("shock-pulse-e", new Rune("shock-pulse-e", "Explosive Bolts", "Slain enemies explode, dealing 105% weapon damage as Lightning to every enemy within 10 yards."));
      SKILL_MAP.put("frost-nova", new Skill("frost-nova", "wizard_frostnova", "Frost Nova"));
      RUNE_MAP.put("frost-nova-a", new Rune("frost-nova-a", "Bone Chill", "Enemies take 15% more damage while frozen or chilled by Frost Nova."));
      RUNE_MAP.put("frost-nova-b", new Rune("frost-nova-b", "Shatter", "A frozen enemy that is killed has a 50% chance of releasing another Frost Nova."));
      RUNE_MAP.put("frost-nova-c", new Rune("frost-nova-c", "Frozen Mist", "Frost Nova no longer freezes enemies, but instead leaves behind a mist of frost that deals 160% weapon damage as Cold over 8 seconds."));
      RUNE_MAP.put("frost-nova-d", new Rune("frost-nova-d", "Cold Snap", "Reduce cooldown of Frost Nova to 9 seconds."));
      RUNE_MAP.put("frost-nova-e", new Rune("frost-nova-e", "Deep Freeze", "If Frost Nova hits at least 5 targets, you gain a 15% bonus to Critical Hit Chance for 12 seconds."));
      SKILL_MAP.put("arcane-orb", new Skill("arcane-orb", "wizard_arcaneorb", "Arcane Orb"));
      RUNE_MAP.put("arcane-orb-a", new Rune("arcane-orb-a", "Obliteration", "Increase the damage of the explosion to deal 260% weapon damage as Arcane."));
      RUNE_MAP.put("arcane-orb-b", new Rune("arcane-orb-b", "Arcane Nova", "Modify the orb to deal 200% weapon damage as Arcane to all enemies within 20 yards."));
      RUNE_MAP.put("arcane-orb-c", new Rune("arcane-orb-c", "Arcane Orbit", "Create 4 Arcane Orbs that orbit you, exploding for 80% weapon damage as Arcane when enemies get close."));
      RUNE_MAP.put("arcane-orb-d", new Rune("arcane-orb-d", "Tap the Source", "Reduce casting cost to 20 Arcane Power."));
      RUNE_MAP.put("arcane-orb-e", new Rune("arcane-orb-e", "Celestial Orb", "The orb will pierce through targets, damaging any enemy it passes through."));
      SKILL_MAP.put("diamond-skin", new Skill("diamond-skin", "wizard_diamondskin", "Diamond Skin"));
      RUNE_MAP.put("diamond-skin-a", new Rune("diamond-skin-a", "Mirror Skin", "Reflects 100% of damage absorbed back at the attacker."));
      RUNE_MAP.put("diamond-skin-b", new Rune("diamond-skin-b", "Enduring Skin", "Increases the duration of Diamond Skin to 9 seconds."));
      RUNE_MAP.put("diamond-skin-c", new Rune("diamond-skin-c", "Crystal Shell", "Increases the maximum amount of damage absorbed to 21707 damage."));
      RUNE_MAP.put("diamond-skin-d", new Rune("diamond-skin-d", "Prism", "Reduces Arcane Power cost of all spells by 7 while Diamond Skin is active."));
      RUNE_MAP.put("diamond-skin-e", new Rune("diamond-skin-e", "Diamond Shards", "When Diamond Skin wears off, diamond shards explode in all directions dealing 210% weapon damage as Physical to nearby enemies."));
      SKILL_MAP.put("wave-of-force", new Skill("wave-of-force", "wizard_waveofforce", "Wave of Force"));
      RUNE_MAP.put("wave-of-force-a", new Rune("wave-of-force-a", "Forceful Wave", "Increases damage to 260% weapon damage as Physical, but reduces Knockback."));
      RUNE_MAP.put("wave-of-force-b", new Rune("wave-of-force-b", "Exploding Wave", "Enemies hit have a 40% chance to cause a smaller Wave of Force that deals 100% weapon damage as Physical and knocks back enemies caught in its wake."));
      RUNE_MAP.put("wave-of-force-c", new Rune("wave-of-force-c", "Teleporting Wave", "Enemies caught in the wave are teleported to a random location."));
      RUNE_MAP.put("wave-of-force-d", new Rune("wave-of-force-d", "Force Affinity", "Reduce casting cost to 15 Arcane Power and the cooldown is reduced to 9 seconds."));
      RUNE_MAP.put("wave-of-force-e", new Rune("wave-of-force-e", "Impactful Wave", "Increases the distance enemies are knocked back and Stuns all affected enemies for 2 seconds."));
      SKILL_MAP.put("spectral-blade", new Skill("spectral-blade", "wizard_spectralblade", "Spectral Blade"));
      RUNE_MAP.put("spectral-blade-a", new Rune("spectral-blade-a", "Deep Cuts", "Enemies hit by the blade will Bleed for an additional 45% weapon damage over 3 seconds."));
      RUNE_MAP.put("spectral-blade-b", new Rune("spectral-blade-b", "Thrown Blade", "Extends the reach of Spectral Blade to 20 yards."));
      RUNE_MAP.put("spectral-blade-c", new Rune("spectral-blade-c", "Impactful Blades", "Hits Slow the movement of enemies by 80% for 1 second."));
      RUNE_MAP.put("spectral-blade-d", new Rune("spectral-blade-d", "Siphoning Blade", "Every enemy hit grants 3 Arcane Power."));
      RUNE_MAP.put("spectral-blade-e", new Rune("spectral-blade-e", "Healing Blades", "Whenever the blades do damage, you are healed for 5% of the damage done."));
      SKILL_MAP.put("arcane-torrent", new Skill("arcane-torrent", "wizard_arcanetorrent", "Arcane Torrent"));
      RUNE_MAP.put("arcane-torrent-a", new Rune("arcane-torrent-a", "Disruption", "Targets hit by Arcane Torrent become disrupted for 6 seconds, causing them to take 15% additional damage from any attacks that deal Arcane damage."));
      RUNE_MAP.put("arcane-torrent-b", new Rune("arcane-torrent-b", "Cascade", "Enemies killed by Arcane Torrent have a 100% chance to fire a new missile at a nearby enemy dealing 285% weapon damage as Arcane."));
      RUNE_MAP.put("arcane-torrent-c", new Rune("arcane-torrent-c", "Arcane Mines", "Instead of firing projectiles, lay Arcane mines that arm after 2 seconds. These mines explode when an enemy approaches, dealing 340% weapon damage as Arcane. Enemies caught in the explosion have their movement and attack speeds reduced by 30% for 3 seconds."));
      RUNE_MAP.put("arcane-torrent-d", new Rune("arcane-torrent-d", "Power Stone", "Every missile hit has a 2% chance to leave behind a Power Stone that grants Arcane Power when picked up."));
      RUNE_MAP.put("arcane-torrent-e", new Rune("arcane-torrent-e", "Death Blossom", "Unleash a torrent of power beyond your control. You no longer direct where the projectiles go, but their damage is increased to 670% weapon damage as Arcane."));
      SKILL_MAP.put("energy-twister", new Skill("energy-twister", "wizard_energytwister", "Energy Twister"));
      RUNE_MAP.put("energy-twister-a", new Rune("energy-twister-a", "Gale Force", "Increases the damage of Energy Twister to 468% weapon damage as Arcane."));
      RUNE_MAP.put("energy-twister-b", new Rune("energy-twister-b", "Raging Storm", "When two Energy Twisters collide, they merge into a tornado with increased area of effect that causes 360% weapon damage as Arcane over 6 seconds."));
      RUNE_MAP.put("energy-twister-c", new Rune("energy-twister-c", "Storm Chaser", "Casting Energy Twister grants you a Wind Charge. You can store up to 3 Wind Charges at a time. Casting a Signature spell releases all Wind Charges as a giant Energy Twister that deals 75% weapon damage as Arcane per Wind Charge.The following skills are Signature spells: Magic Missile Shock Pulse Spectral Blade Electrocute"));
      RUNE_MAP.put("energy-twister-d", new Rune("energy-twister-d", "Mistral Breeze", "Reduces casting cost of Energy Twister to 20 Arcane Power."));
      RUNE_MAP.put("energy-twister-e", new Rune("energy-twister-e", "Wicked Wind", "Twisters no longer travel but spin in place, dealing 252% weapon damage as Arcane over 6 seconds to everything caught in them."));
      SKILL_MAP.put("ice-armor", new Skill("ice-armor", "wizard_icearmor", "Ice Armor"));
      RUNE_MAP.put("ice-armor-a", new Rune("ice-armor-a", "Jagged Ice", "Melee attackers also take 130% weapon damage as Cold."));
      RUNE_MAP.put("ice-armor-b", new Rune("ice-armor-b", "Chilling Aura", "Lower the temperature of the air around you. Nearby enemies are chilled, slowing their movement speed by 30%."));
      RUNE_MAP.put("ice-armor-c", new Rune("ice-armor-c", "Frozen Storm", "A whirling storm of ice builds around you that deals 30% weapon damage as Cold over 3 seconds after casting Ice Armor."));
      RUNE_MAP.put("ice-armor-d", new Rune("ice-armor-d", "Crystallize", "Whenever you are struck by a melee attack, your Armor is increased by 20% for 30 seconds. This effect can stack up to 3 times."));
      RUNE_MAP.put("ice-armor-e", new Rune("ice-armor-e", "Ice Reflect", "Melee attacks have a 25% chance to create a Frost Nova centered on the attacker, dealing 75% weapon damage as Cold."));
      SKILL_MAP.put("electrocute", new Skill("electrocute", "wizard_electrocute", "Electrocute"));
      RUNE_MAP.put("electrocute-a", new Rune("electrocute-a", "Lightning Blast", "Create streaks of lightning that pierce through targets, hitting all enemies for 86% weapon damage as Lightning."));
      RUNE_MAP.put("electrocute-b", new Rune("electrocute-b", "Chain Lightning", "Increases the maximum number of enemies that can be electrocuted to 6."));
      RUNE_MAP.put("electrocute-c", new Rune("electrocute-c", "Arc Lightning", "Blast a cone of lightning that causes 115% weapon damage as Lightning to all affected targets."));
      RUNE_MAP.put("electrocute-d", new Rune("electrocute-d", "Surge of Power", "Gain 1 Arcane Power for every enemy hit by Electrocute."));
      RUNE_MAP.put("electrocute-e", new Rune("electrocute-e", "Forked Lightning", "Critical Hits release 4 charged bolts in random directions, dealing 55% weapon damage as Lightning to any targets hit."));
      SKILL_MAP.put("slow-time", new Skill("slow-time", "wizard_slowtime", "Slow Time"));
      RUNE_MAP.put("slow-time-a", new Rune("slow-time-a", "Time Warp", "Enemies caught in the bubble of warped time take 20% more damage."));
      RUNE_MAP.put("slow-time-b", new Rune("slow-time-b", "Miasma", "Slow Time effects cling to enemies for 3 seconds after they have left the bubble."));
      RUNE_MAP.put("slow-time-c", new Rune("slow-time-c", "Time Shell", "Increases the potency of the movement speed reduction to 80%."));
      RUNE_MAP.put("slow-time-d", new Rune("slow-time-d", "Perpetuity", "Reduces the cooldown of Slow Time to 12 seconds."));
      RUNE_MAP.put("slow-time-e", new Rune("slow-time-e", "Stretch Time", "Time is sped up for any allies standing in the area, increasing their attack speed by 10%."));
      SKILL_MAP.put("storm-armor", new Skill("storm-armor", "wizard_stormarmor", "Storm Armor"));
      RUNE_MAP.put("storm-armor-a", new Rune("storm-armor-a", "Thunder Storm", "Increase the damage of the shock to 130% weapon damage as Lightning."));
      RUNE_MAP.put("storm-armor-b", new Rune("storm-armor-b", "Scramble", "Increases your movement speed by 25% for 3 seconds whenever you are hit by melee or ranged attacks."));
      RUNE_MAP.put("storm-armor-c", new Rune("storm-armor-c", "Reactive Armor", "Ranged and melee attackers are shocked for 70% weapon damage as Lightning."));
      RUNE_MAP.put("storm-armor-d", new Rune("storm-armor-d", "Power of the Storm", "Reduce the Arcane Power cost of all skills by 3 while Storm Armor is active."));
      RUNE_MAP.put("storm-armor-e", new Rune("storm-armor-e", "Shocking Aspect", "Critical Hits have a chance to electrocute a nearby enemy for 35% weapon damage as Lightning."));
      SKILL_MAP.put("explosive-blast", new Skill("explosive-blast", "wizard_explosiveblast", "Explosive Blast"));
      RUNE_MAP.put("explosive-blast-a", new Rune("explosive-blast-a", "Short Fuse", "Immediately release the energy of Explosive Blast for 225% weapon damage as Physical."));
      RUNE_MAP.put("explosive-blast-b", new Rune("explosive-blast-b", "Obliterate", "Increases the explosion radius to 18 yards for 225% weapon damage as Physical."));
      RUNE_MAP.put("explosive-blast-c", new Rune("explosive-blast-c", "Time Bomb", "Explosive Blast detonates from the point it was originally cast after 2.5 seconds for 315% weapon damage as Physical."));
      RUNE_MAP.put("explosive-blast-d", new Rune("explosive-blast-d", "Unleashed", "Reduces the casting cost of Explosive Blast to 10 Arcane Power."));
      RUNE_MAP.put("explosive-blast-e", new Rune("explosive-blast-e", "Chain Reaction", "A chain of 3 consecutive explosions cascade off you, each causing 97% weapon damage as Physical."));
      SKILL_MAP.put("magic-weapon", new Skill("magic-weapon", "wizard_magicweapon", "Magic Weapon"));
      RUNE_MAP.put("magic-weapon-a", new Rune("magic-weapon-a", "Venom", "Attacks poison enemies, dealing 15% weapon damage as Poison over 3 seconds."));
      RUNE_MAP.put("magic-weapon-b", new Rune("magic-weapon-b", "Electrify", "Attacks have a chance to cause lightning to arc to 3 nearby enemies, dealing 10% weapon damage as Lightning."));
      RUNE_MAP.put("magic-weapon-c", new Rune("magic-weapon-c", "Force Weapon", "Increases the damage bonus of Magic Weapon to 15% damage, and gives up to a 2% chance to Knockback any enemies hit."));
      RUNE_MAP.put("magic-weapon-d", new Rune("magic-weapon-d", "Conduit", "Attacks have a chance to restore 1 Arcane Power."));
      RUNE_MAP.put("magic-weapon-e", new Rune("magic-weapon-e", "Blood Magic", "Attacks recover 1.5% of damage caused as Life."));
      SKILL_MAP.put("hydra", new Skill("hydra", "wizard_hydra", "Hydra"));
      RUNE_MAP.put("hydra-a", new Rune("hydra-a", "Frost Hydra", "Summon a Frost Hydra that breathes a short range cone of frost, causing 36% weapon damage as Cold to all enemies in the cone."));
      RUNE_MAP.put("hydra-b", new Rune("hydra-b", "Lightning Hydra", "Summon a Lightning Hydra that electrocutes enemies for 64% weapon damage as Lightning."));
      RUNE_MAP.put("hydra-c", new Rune("hydra-c", "Venom Hydra", "Summon a poison breathing Hydra that leaves a pool of acid that causes 18% weapon damage per second as Poison to enemies who remain in the pool."));
      RUNE_MAP.put("hydra-d", new Rune("hydra-d", "Mammoth Hydra", "Summon a Mammoth Hydra that breathes a river of flame at nearby enemies, dealing 67% weapon damage per second as Fire to enemies caught on the burning ground."));
      RUNE_MAP.put("hydra-e", new Rune("hydra-e", "Arcane Hydra", "Summon an Arcane Hydra that spits Arcane Orbs, which explode on impact, causing 60% weapon damage as Arcane to enemies near the explosion."));
      SKILL_MAP.put("disintegrate", new Skill("disintegrate", "wizard_disintegrate", "Disintegrate"));
      RUNE_MAP.put("disintegrate-a", new Rune("disintegrate-a", "Intensify", "Damage increases slowly over time to inflict a maximum of 286% weapon damage as Arcane."));
      RUNE_MAP.put("disintegrate-b", new Rune("disintegrate-b", "Convergence", "Increase the width of the beam allowing it to hit more enemies."));
      RUNE_MAP.put("disintegrate-c", new Rune("disintegrate-c", "Entropy", "The beam fractures into a short-ranged cone that deals 253% weapon damage as Arcane."));
      RUNE_MAP.put("disintegrate-d", new Rune("disintegrate-d", "Chaos Nexus", "When casting the beam you become charged with energy that spits out at nearby enemies doing 44% weapon damage as Arcane."));
      RUNE_MAP.put("disintegrate-e", new Rune("disintegrate-e", "Volatility", "Enemies killed by the beam have a 35% chance to explode causing 395% weapon damage as Arcane to all enemies within 8 yards."));
      SKILL_MAP.put("familiar", new Skill("familiar", "wizard_familiar", "Familiar"));
      RUNE_MAP.put("familiar-a", new Rune("familiar-a", "Sparkflint", "Summon a fiery Familiar that increases the damage of all attacks by 12% while Familiar is active."));
      RUNE_MAP.put("familiar-b", new Rune("familiar-b", "Cannoneer", "The Familiar's projectiles explode on impact, dealing 20% weapon damage as Arcane to all enemies within 6 yards."));
      RUNE_MAP.put("familiar-c", new Rune("familiar-c", "Vigoron", "While the Familiar is active, you regenerate 620 Life per second."));
      RUNE_MAP.put("familiar-d", new Rune("familiar-d", "Arcanot", "While the Familiar is active, you regenerate 2 Arcane Power per second."));
      RUNE_MAP.put("familiar-e", new Rune("familiar-e", "Ancient Guardian", "Summon a protective Familiar. When you are below 50% Life the Familiar will fully absorb damage from 1 attack every 6 seconds."));
      SKILL_MAP.put("teleport", new Skill("teleport", "wizard_teleport", "Teleport"));
      RUNE_MAP.put("teleport-a", new Rune("teleport-a", "Calamity", "Casts a low power Wave of Force upon arrival, dealing 265% weapon damage as Physical to all nearby enemies."));
      RUNE_MAP.put("teleport-b", new Rune("teleport-b", "Fracture", "Summon 2 decoys for 8 seconds after teleporting."));
      RUNE_MAP.put("teleport-c", new Rune("teleport-c", "Safe Passage", "For 4 seconds after you Teleport, you will take 30% less damage."));
      RUNE_MAP.put("teleport-d", new Rune("teleport-d", "Reversal", "Casting Teleport again within 8 seconds will instantly return you to your original location."));
      RUNE_MAP.put("teleport-e", new Rune("teleport-e", "Wormhole", "After casting Teleport, there is a 1 second delay before the cooldown begins, allowing you to Teleport again."));
      SKILL_MAP.put("mirror-image", new Skill("mirror-image", "wizard_mirrorimage", "Mirror Image"));
      RUNE_MAP.put("mirror-image-a", new Rune("mirror-image-a", "Mirror Mimics", "Spells cast by your Mirror Images will do 10% of the damage of your own spells."));
      RUNE_MAP.put("mirror-image-b", new Rune("mirror-image-b", "Duplicates", "Summon 5 Mirror Images that have 25% of your Life each."));
      RUNE_MAP.put("mirror-image-c", new Rune("mirror-image-c", "Simulacrum", "Increase the Life of your Mirror Images to 100% of your own."));
      RUNE_MAP.put("mirror-image-d", new Rune("mirror-image-d", "Extension of Will", "The duration of your Mirror Images is increased to 10 seconds and their Life is increased to 29% of your Life."));
      RUNE_MAP.put("mirror-image-e", new Rune("mirror-image-e", "Mocking Demise", "When a Mirror Image is destroyed, it explodes, doing 45% weapon damage as Physical and has a 50% chance to Stun for 2 seconds."));
      SKILL_MAP.put("meteor", new Skill("meteor", "wizard_meteor", "Meteor"));
      RUNE_MAP.put("meteor-a", new Rune("meteor-a", "Molten Impact", "Increases the damage of the Meteor impact to 390% weapon damage as Fire and the molten fire to 90% weapon damage as Fire over 3 seconds."));
      RUNE_MAP.put("meteor-b", new Rune("meteor-b", "Meteor Shower", "Unleash a volley of 7 smaller Meteors that each strike for 104% weapon damage as Fire."));
      RUNE_MAP.put("meteor-c", new Rune("meteor-c", "Comet", "Transforms the Meteor to ice that deals 312% weapon damage as Cold. The impact site is covered in a freezing mist that deals 72% weapon damage as Cold and Slows enemy movement by 60% over 3 seconds."));
      RUNE_MAP.put("meteor-d", new Rune("meteor-d", "Star Pact", "Reduces the casting cost of Meteor to 35 Arcane Power and the damage type to Arcane."));
      RUNE_MAP.put("meteor-e", new Rune("meteor-e", "Liquefy", "If the initial impact of the Meteor causes a Critical Hit, the molten fire duration is increased to 8 seconds."));
      SKILL_MAP.put("blizzard", new Skill("blizzard", "wizard_blizzard", "Blizzard"));
      RUNE_MAP.put("blizzard-a", new Rune("blizzard-a", "Unrelenting Storm", "Increases the duration of Blizzard to deal 680% weapon damage as Cold over 8 seconds."));
      RUNE_MAP.put("blizzard-b", new Rune("blizzard-b", "Stark Winter", "Increases the size of Blizzard to cover 22 yards, dealing 510% weapon damage as Cold over 6 seconds."));
      RUNE_MAP.put("blizzard-c", new Rune("blizzard-c", "Grasping Chill", "After the Blizzard ends, the ground is covered in a low lying mist for 3 seconds that Slows the movement speed of enemies by 60%."));
      RUNE_MAP.put("blizzard-d", new Rune("blizzard-d", "Snowbound", "Reduces the casting cost of Blizzard to 20 Arcane Power."));
      RUNE_MAP.put("blizzard-e", new Rune("blizzard-e", "Frozen Solid", "Enemies caught in the Blizzard have a 20% chance to be Frozen for 3 seconds."));
      SKILL_MAP.put("energy-armor", new Skill("energy-armor", "wizard_energyarmor", "Energy Armor"));
      RUNE_MAP.put("energy-armor-a", new Rune("energy-armor-a", "Prismatic Armor", "Increases all of your resistances by 25% while Energy Armor is active."));
      RUNE_MAP.put("energy-armor-b", new Rune("energy-armor-b", "Energy Tap", "Rather than decreasing your maximum Arcane Power, Energy Armor increases it by 20 while it is active."));
      RUNE_MAP.put("energy-armor-c", new Rune("energy-armor-c", "Force Armor", "While Energy Armor is active, incoming attacks that would deal more than 35% of your maximum Life are reduced to deal 35% of your maximum Life instead."));
      RUNE_MAP.put("energy-armor-d", new Rune("energy-armor-d", "Absorption", "You have a chance to gain 4 Arcane Power whenever you are hit by a ranged or melee attack."));
      RUNE_MAP.put("energy-armor-e", new Rune("energy-armor-e", "Pinpoint Barrier", "Increases your chance to critically hit by 5% while Energy Armor is active."));
      SKILL_MAP.put("archon", new Skill("archon", "wizard_archon", "Archon"));
      RUNE_MAP.put("archon-a", new Rune("archon-a", "Improved Archon", "Increases the damage of all Archon abilities by 25%."));
      RUNE_MAP.put("archon-b", new Rune("archon-b", "Slow Time", "Archon form can cast Slow Time that lasts for 15 seconds."));
      RUNE_MAP.put("archon-c", new Rune("archon-c", "Teleport", "Archon form can now cast Teleport with a cooldown of 10 seconds."));
      RUNE_MAP.put("archon-d", new Rune("archon-d", "Pure Power", "Decreases the cooldown of Archon to 100 seconds."));
      RUNE_MAP.put("archon-e", new Rune("archon-e", "Arcane Destruction", "An explosion erupts around you when you transform, causing 1600% weapon damage as Arcane to all enemies within 15 yards."));
      SKILL_MAP.put("poison-dart", new Skill("poison-dart", "witchdoctor_poisondart", "Poison Dart"));
      RUNE_MAP.put("poison-dart-a", new Rune("poison-dart-a", "Flaming Dart", "Ignite the dart so that it deals 180% weapon damage as Fire at once."));
      RUNE_MAP.put("poison-dart-b", new Rune("poison-dart-b", "Splinters", "Shoot 3 Poison Darts that deal 60% weapon damage as Poison each."));
      RUNE_MAP.put("poison-dart-c", new Rune("poison-dart-c", "Numbing Dart", "Toxins in the Poison Dart reduce the target's movement speed by 60% for 2 seconds."));
      RUNE_MAP.put("poison-dart-d", new Rune("poison-dart-d", "Spined Dart", "Gain 29 Mana every time a Poison Dart hits an enemy."));
      RUNE_MAP.put("poison-dart-e", new Rune("poison-dart-e", "Snake to the Face", "Transform your Poison Dart into a snake that has a 30% chance to Stun the enemy for 1.5 seconds."));
      SKILL_MAP.put("corpse-spiders", new Skill("corpse-spiders", "witchdoctor_corpsespider", "Corpse Spiders"));
      RUNE_MAP.put("corpse-spiders-a", new Rune("corpse-spiders-a", "Blazing Spiders", "Summon fire spiders that deal a total of 156% weapon damage as Fire."));
      RUNE_MAP.put("corpse-spiders-b", new Rune("corpse-spiders-b", "Spider Queen", "Summon a spider queen that births spiderlings, dealing 630% weapon damage as Poison over 15 seconds.You may only have one spider queen summoned at a time."));
      RUNE_MAP.put("corpse-spiders-c", new Rune("corpse-spiders-c", "Leaping Spiders", "Summon jumping spiders that leap up to 25 yards to reach their target and attack for a total of 144% weapon damage as Physical."));
      RUNE_MAP.put("corpse-spiders-d", new Rune("corpse-spiders-d", "Widowmakers", "Summon widowmaker spiders that return 4 Mana to you per hit."));
      RUNE_MAP.put("corpse-spiders-e", new Rune("corpse-spiders-e", "Medusa Spiders", "Summon paralyzing spiders that have a 25% chance to Slow enemies' movement by 60% with every attack."));
      SKILL_MAP.put("plague-of-toads", new Skill("plague-of-toads", "witchdoctor_plagueoftoads", "Plague of Toads"));
      RUNE_MAP.put("plague-of-toads-a", new Rune("plague-of-toads-a", "Explosive Toads", "Mutate to fire bullfrogs that explode for 169% weapon damage as Fire."));
      RUNE_MAP.put("plague-of-toads-b", new Rune("plague-of-toads-b", "Rain of Toads", "Cause toads to rain from the sky that deal 130% weapon damage as Poison to enemies in the area over 2 seconds."));
      RUNE_MAP.put("plague-of-toads-c", new Rune("plague-of-toads-c", "Toad of Hugeness", "Summon a giant toad that swallows enemies whole for up to 5 seconds, digesting for 20% of your weapon damage per second as Physical. Adds a 5 second cooldown to Plague of Toads."));
      RUNE_MAP.put("plague-of-toads-d", new Rune("plague-of-toads-d", "Toad Affinity", "Removes the Mana cost of Plague of Toads."));
      RUNE_MAP.put("plague-of-toads-e", new Rune("plague-of-toads-e", "Addling Toads", "Mutate to yellow frogs that deal 130% weapon damage as Poison and have a 15% chance to Confuse affected enemies for 4 seconds."));
      SKILL_MAP.put("firebomb", new Skill("firebomb", "witchdoctor_firebomb", "Firebomb"));
      RUNE_MAP.put("firebomb-a", new Rune("firebomb-a", "Ghost Bomb", "In addition to the base explosion, the skull creates a larger blast that deals an additional 30% weapon damage as Fire to all enemies within 28 yards."));
      RUNE_MAP.put("firebomb-b", new Rune("firebomb-b", "Roll the Bones", "Allows the skull to bounce up to 2 times."));
      RUNE_MAP.put("firebomb-c", new Rune("firebomb-c", "Fire Pit", "The explosion creates a pool of fire that deals 36% weapon damage as Fire over 3 seconds."));
      RUNE_MAP.put("firebomb-d", new Rune("firebomb-d", "Pyrogeist", "Create a column of flame that spews fire at the closest enemy for 640% weapon damage as Fire over 6 seconds. You may only have one Pyrogeist active at a time."));
      RUNE_MAP.put("firebomb-e", new Rune("firebomb-e", "Flash Fire", "Rather than exploding for area damage, each Firebomb can bounce to up to 6 additional targets. Damage is reduced by 15% per bounce."));
      SKILL_MAP.put("grasp-of-the-dead", new Skill("grasp-of-the-dead", "witchdoctor_graspofthedead", "Grasp of the Dead"));
      RUNE_MAP.put("grasp-of-the-dead-a", new Rune("grasp-of-the-dead-a", "Groping Eels", "Increases the damage done to 416% weapon damage as Physical."));
      RUNE_MAP.put("grasp-of-the-dead-b", new Rune("grasp-of-the-dead-b", "Rain of Corpses", "Corpses fall from the sky, dealing 80% weapon damage as Physical over 8 seconds to nearby enemies."));
      RUNE_MAP.put("grasp-of-the-dead-c", new Rune("grasp-of-the-dead-c", "Unbreakable Grasp", "Increases the Slow amount to 80%."));
      RUNE_MAP.put("grasp-of-the-dead-d", new Rune("grasp-of-the-dead-d", "Desperate Grasp", "Reduces the cooldown of Grasp of the Dead to 6 seconds."));
      RUNE_MAP.put("grasp-of-the-dead-e", new Rune("grasp-of-the-dead-e", "Death Is Life", "Enemies who die while in the area of Grasp of the Dead have a 10% chance to produce a health globe or summon a Zombie Dog."));
      SKILL_MAP.put("firebats", new Skill("firebats", "witchdoctor_firebats", "Firebats"));
      RUNE_MAP.put("firebats-a", new Rune("firebats-a", "Dire Bats", "Summon fewer but larger bats that travel up to 40 yards and hit for 220% weapon damage as Fire."));
      RUNE_MAP.put("firebats-b", new Rune("firebats-b", "Hungry Bats", "Rapidly summon bats that seek out nearby enemies for 350% weapon damage as Fire."));
      RUNE_MAP.put("firebats-c", new Rune("firebats-c", "Plague Bats", "Diseased bats fly towards the enemy and infect them. Damage is slow at first, but can increase over time to a maximum of 578% weapon damage as Poison."));
      RUNE_MAP.put("firebats-d", new Rune("firebats-d", "Vampire Bats", "Gain 2.5% of damage done by the bats as Life."));
      RUNE_MAP.put("firebats-e", new Rune("firebats-e", "Cloud of Bats", "Call forth a swirl of bats that damage nearby enemies for 501% weapon damage as Fire. The damage of the bats increases by 20% every second, up to a maximum of 100%."));
      SKILL_MAP.put("haunt", new Skill("haunt", "witchdoctor_haunt", "Haunt"));
      RUNE_MAP.put("haunt-a", new Rune("haunt-a", "Consuming Spirit", "The spirit returns 155 Life per second."));
      RUNE_MAP.put("haunt-b", new Rune("haunt-b", "Lingering Spirit", "If there are no targets left, the spirit will linger for up to 10 seconds looking for new enemies."));
      RUNE_MAP.put("haunt-c", new Rune("haunt-c", "Grasping Spirit", "Slow the movement of haunted targets by 30%."));
      RUNE_MAP.put("haunt-d", new Rune("haunt-d", "Draining Spirit", "The spirit returns 20.4 Mana per second."));
      RUNE_MAP.put("haunt-e", new Rune("haunt-e", "Resentful Spirit", "Summon a vengeful spirit that does 383% weapon damage as Arcane over 2 seconds."));
      SKILL_MAP.put("locust-swarm", new Skill("locust-swarm", "witchdoctor_locust_swarm", "Locust Swarm"));
      RUNE_MAP.put("locust-swarm-a", new Rune("locust-swarm-a", "Searing Locusts", "Engulf the target with burning locusts that deal 468% weapon damage as Fire over 8 seconds."));
      RUNE_MAP.put("locust-swarm-b", new Rune("locust-swarm-b", "Pestilence", "Locust Swarm has a 100% chance to jump to two additional targets instead of one."));
      RUNE_MAP.put("locust-swarm-c", new Rune("locust-swarm-c", "Cloud of Insects", "Increases the duration of the swarm to 10 seconds."));
      RUNE_MAP.put("locust-swarm-d", new Rune("locust-swarm-d", "Devouring Swarm", "Gain 37 Mana for every enemy affected by the swarm."));
      RUNE_MAP.put("locust-swarm-e", new Rune("locust-swarm-e", "Diseased Swarm", "Enemies killed by Locust Swarm leave behind a cloud of locusts that deal 75% weapon damage as Poison over 3 seconds to enemies who stand in the area."));
      SKILL_MAP.put("summon-zombie-dogs", new Skill("summon-zombie-dogs", "witchdoctor_summonzombiedog", "Summon Zombie Dogs"));
      RUNE_MAP.put("summon-zombie-dogs-a", new Rune("summon-zombie-dogs-a", "Burning Dogs", "Your Zombie Dogs burst into flames, burning nearby enemies for 2% of your weapon damage as Fire every second."));
      RUNE_MAP.put("summon-zombie-dogs-b", new Rune("summon-zombie-dogs-b", "Life Link", "Your Zombie Dogs absorb 10% of all damage done to you."));
      RUNE_MAP.put("summon-zombie-dogs-c", new Rune("summon-zombie-dogs-c", "Rabid Dogs", "Your Zombie Dogs gain an infectious bite that deals 9% of your weapon damage as Poison over 3 seconds."));
      RUNE_MAP.put("summon-zombie-dogs-d", new Rune("summon-zombie-dogs-d", "Final Gift", "Your Zombie Dogs have a 15% chance to leave behind a health globe when they die."));
      RUNE_MAP.put("summon-zombie-dogs-e", new Rune("summon-zombie-dogs-e", "Leeching Beasts", "Your Zombie Dogs heal 50% of the damage they deal as Life divided evenly between themselves and you."));
      SKILL_MAP.put("horrify", new Skill("horrify", "witchdoctor_horrify", "Horrify"));
      RUNE_MAP.put("horrify-a", new Rune("horrify-a", "Frightening Aspect", "Gain 100% additional Armor for 8 seconds after casting Horrify."));
      RUNE_MAP.put("horrify-b", new Rune("horrify-b", "Face of Death", "Increases the radius of Horrify to 24 yards."));
      RUNE_MAP.put("horrify-c", new Rune("horrify-c", "Phobia", "Increases the duration horrified enemies run in Fear to 6 seconds."));
      RUNE_MAP.put("horrify-d", new Rune("horrify-d", "Ruthless Terror", "Gain 27 Mana for every horrified enemy."));
      RUNE_MAP.put("horrify-e", new Rune("horrify-e", "Stalker", "Increases movement speed by 20% for 4 seconds after casting Horrify."));
      SKILL_MAP.put("spirit-walk", new Skill("spirit-walk", "witchdoctor_spiritwalk", "Spirit Walk"));
      RUNE_MAP.put("spirit-walk-a", new Rune("spirit-walk-a", "Severance", "Damage enemies you walk through in spirit form for 225% weapon damage as Physical every second for 2 seconds."));
      RUNE_MAP.put("spirit-walk-b", new Rune("spirit-walk-b", "Jaunt", "Increases the duration of Spirit Walk to 3 seconds."));
      RUNE_MAP.put("spirit-walk-c", new Rune("spirit-walk-c", "Umbral Shock", "When Spirit Walk ends, your physical body erupts for 310% weapon damage as Fire to all enemies within 10 yards."));
      RUNE_MAP.put("spirit-walk-d", new Rune("spirit-walk-d", "Honored Guest", "Gain 15% of your maximum Mana every second while Spirit Walk is active."));
      RUNE_MAP.put("spirit-walk-e", new Rune("spirit-walk-e", "Healing Journey", "Gain 7% of your maximum Life every second while Spirit Walk is active."));
      SKILL_MAP.put("hex", new Skill("hex", "witchdoctor_hex", "Hex"));
      RUNE_MAP.put("hex-a", new Rune("hex-a", "Painful Transformation", "Hex causes the target to Bleed for 12% weapon damage as Physical."));
      RUNE_MAP.put("hex-b", new Rune("hex-b", "Angry Chicken", "Transform into an angry chicken for up to 5 seconds that can explode for 215% weapon damage as Physical to all enemies within 12 yards."));
      RUNE_MAP.put("hex-c", new Rune("hex-c", "Unstable Form", "Hexed targets explode when killed, dealing 135% weapon damage as Poison to all enemies within 8 yards."));
      RUNE_MAP.put("hex-d", new Rune("hex-d", "Hedge Magic", "The Fetish Shaman will periodically heal allies for 1861 Life."));
      RUNE_MAP.put("hex-e", new Rune("hex-e", "Jinx", "Hexed targets take 20% additional damage."));
      SKILL_MAP.put("soul-harvest", new Skill("soul-harvest", "witchdoctor_soulharvest", "Soul Harvest"));
      RUNE_MAP.put("soul-harvest-a", new Rune("soul-harvest-a", "Siphon", "Gain 2171 Life for every enemy harvested."));
      RUNE_MAP.put("soul-harvest-b", new Rune("soul-harvest-b", "Soul to Waste", "Increase the duration of Soul Harvest's effect to 60 seconds."));
      RUNE_MAP.put("soul-harvest-c", new Rune("soul-harvest-c", "Languish", "Reduces the movement speed of harvested enemies by 80% for 3 seconds."));
      RUNE_MAP.put("soul-harvest-d", new Rune("soul-harvest-d", "Swallow Your Soul", "Gain 39 Mana for every enemy harvested."));
      RUNE_MAP.put("soul-harvest-e", new Rune("soul-harvest-e", "Vengeful Spirit", "Harvested enemies also take 230% weapon damage as Physical."));
      SKILL_MAP.put("sacrifice", new Skill("sacrifice", "witchdoctor_sacrifice", "Sacrifice"));
      RUNE_MAP.put("sacrifice-a", new Rune("sacrifice-a", "Provoke the Pack", "Each sacrificed Zombie Dog increases your damage by 5% for 30 seconds."));
      RUNE_MAP.put("sacrifice-b", new Rune("sacrifice-b", "For the Master", "Gain 6202 Life for each Zombie Dog you sacrifice."));
      RUNE_MAP.put("sacrifice-c", new Rune("sacrifice-c", "Black Blood", "Ichor erupts from the corpses of the Zombie Dogs and Slows enemies by 60% for 8 seconds."));
      RUNE_MAP.put("sacrifice-d", new Rune("sacrifice-d", "Pride", "Regain 294 Mana for each Zombie Dog you sacrifice."));
      RUNE_MAP.put("sacrifice-e", new Rune("sacrifice-e", "Next of Kin", "Each Zombie Dog you sacrifice has a 35% chance to resurrect as a new Zombie Dog."));
      SKILL_MAP.put("mass-confusion", new Skill("mass-confusion", "witchdoctor_massconfusion", "Mass Confusion"));
      RUNE_MAP.put("mass-confusion-a", new Rune("mass-confusion-a", "Paranoia", "All enemies in the area of Mass Confusion take 20% additional damage for 12 seconds."));
      RUNE_MAP.put("mass-confusion-b", new Rune("mass-confusion-b", "Mass Hysteria", "Up to 6 enemies who aren't Confused are Stunned for 3 seconds."));
      RUNE_MAP.put("mass-confusion-c", new Rune("mass-confusion-c", "Mass Hallucination", "Amid the confusion, a giant spirit rampages through enemies, dealing 22% weapon damage per second as Physical to enemies it passes through."));
      RUNE_MAP.put("mass-confusion-d", new Rune("mass-confusion-d", "Unstable Realm", "Reduces the cooldown of Mass Confusion to 45 seconds."));
      RUNE_MAP.put("mass-confusion-e", new Rune("mass-confusion-e", "Devolution", "Enemies killed while Confused have a 50% chance of spawning a Zombie Dog."));
      SKILL_MAP.put("zombie-charger", new Skill("zombie-charger", "witchdoctor_zombiecharger", "Zombie Charger"));
      RUNE_MAP.put("zombie-charger-a", new Rune("zombie-charger-a", "Zombie Bears", "Summon zombie bears that stampede towards your target. Each bear deals 236% weapon damage as Poison to enemies in the area."));
      RUNE_MAP.put("zombie-charger-b", new Rune("zombie-charger-b", "Wave of Zombies", "Summon 3 Zombie Chargers that each deal 115% weapon damage as Poison."));
      RUNE_MAP.put("zombie-charger-c", new Rune("zombie-charger-c", "Leperous Zombie", "The Zombie Charger leaves behind a cloud of noxious vapors that deals 240% weapon damage as Poison over 3 seconds to enemies caught in it."));
      RUNE_MAP.put("zombie-charger-d", new Rune("zombie-charger-d", "Undeath", "If the Zombie Charger kills any enemies, it will reanimate and charge nearby enemies for 205% weapon damage as Poison. This effect can repeat up to 2 times."));
      RUNE_MAP.put("zombie-charger-e", new Rune("zombie-charger-e", "Explosive Beast", "Summon an explosive Zombie Dog that streaks toward your target before exploding, dealing 236% weapon damage as Fire to all enemies within 9 yards."));
      SKILL_MAP.put("spirit-barrage", new Skill("spirit-barrage", "witchdoctor_spiritbarrage", "Spirit Barrage"));
      RUNE_MAP.put("spirit-barrage-a", new Rune("spirit-barrage-a", "Phlebotomize", "Regain 3% of damage dealt with Spirit Barrage as Life."));
      RUNE_MAP.put("spirit-barrage-b", new Rune("spirit-barrage-b", "Well of Souls", "An additional 3 spirits seek out other targets and deal 65% weapon damage as Physical."));
      RUNE_MAP.put("spirit-barrage-c", new Rune("spirit-barrage-c", "Phantasm", "Summon a spectre that deals 225% weapon damage as Physical over 5 seconds to all enemies within 10 yards."));
      RUNE_MAP.put("spirit-barrage-d", new Rune("spirit-barrage-d", "The Spirit Is Willing", "Gain 44 Mana every time Spirit Barrage hits."));
      RUNE_MAP.put("spirit-barrage-e", new Rune("spirit-barrage-e", "Manitou", "Summon a spectre that hovers over you, unleashing spirit bolts at nearby enemies for 1667% weapon damage as Physical over 20 seconds."));
      SKILL_MAP.put("acid-cloud", new Skill("acid-cloud", "witchdoctor_acidcloud", "Acid Cloud"));
      RUNE_MAP.put("acid-cloud-a", new Rune("acid-cloud-a", "Corpse Bomb", "Raise a corpse from the ground that explodes for 230% weapon damage as Poison to enemies in the area."));
      RUNE_MAP.put("acid-cloud-b", new Rune("acid-cloud-b", "Acid Rain", "Increases the initial area of effect of Acid Cloud to 24 yards."));
      RUNE_MAP.put("acid-cloud-c", new Rune("acid-cloud-c", "Lob Blob Bomb", "The acid on the ground forms into a slime that irradiates nearby enemies for 250% weapon damage as Poison over 5 seconds."));
      RUNE_MAP.put("acid-cloud-d", new Rune("acid-cloud-d", "Slow Burn", "Increases the duration of the acid pools left behind to deal 300% weapon damage as Poison over 6 seconds."));
      RUNE_MAP.put("acid-cloud-e", new Rune("acid-cloud-e", "Kiss of Death", "Spit a cloud of acid that inflicts 127% weapon damage as Poison, followed by 165% weapon damage as Poison over 3 seconds."));
      SKILL_MAP.put("wall-of-zombies", new Skill("wall-of-zombies", "witchdoctor_wallofzombies", "Wall of Zombies"));
      RUNE_MAP.put("wall-of-zombies-a", new Rune("wall-of-zombies-a", "Creepers", "Up to 3 zombies will emerge from the ground and attack nearby enemies for 25% of your weapon damage as Physical per attack."));
      RUNE_MAP.put("wall-of-zombies-b", new Rune("wall-of-zombies-b", "Barricade", "Increases the width of the Wall of Zombies."));
      RUNE_MAP.put("wall-of-zombies-c", new Rune("wall-of-zombies-c", "Dead Rush", "Zombies crawl out of the ground and run in all directions, dealing 445% weapon damage as Physical to nearby enemies."));
      RUNE_MAP.put("wall-of-zombies-d", new Rune("wall-of-zombies-d", "Unrelenting Grip", "Your Wall of Zombies will Slow the movement of enemies by 60% for 5 seconds."));
      RUNE_MAP.put("wall-of-zombies-e", new Rune("wall-of-zombies-e", "Pile On", "Summon a tower of zombies that falls over, dealing 550% weapon damage as Physical to any enemies it hits and knocks them back. Reduces the cooldown to 10 seconds."));
      SKILL_MAP.put("gargantuan", new Skill("gargantuan", "witchdoctor_gargantuan", "Gargantuan"));
      RUNE_MAP.put("gargantuan-a", new Rune("gargantuan-a", "Restless Giant", "When the Gargantuan encounters an elite enemy or is near 5 enemies, it enrages for 15 seconds gaining:  20% movement speed  35% attack speed  200% Physical damageThis effect cannot occur more than once every 120 seconds. Elite enemies include champions, rares, bosses, and other players."));
      RUNE_MAP.put("gargantuan-b", new Rune("gargantuan-b", "Humongoid", "The Gargantuan gains the Cleave ability, allowing its attacks to hit multiple targets for 130% of your weapon damage as Physical."));
      RUNE_MAP.put("gargantuan-c", new Rune("gargantuan-c", "Big Stinker", "The Gargantuan is surrounded by a poison cloud that deals 15% weapon damage as Poison per second to nearby enemies."));
      RUNE_MAP.put("gargantuan-d", new Rune("gargantuan-d", "Wrathful Protector", "Summon a more powerful Gargantuan that only lasts for 15 seconds. The Gargantuan's fists burn with fire, dealing 110% of your weapon damage as Fire and knocking enemies back."));
      RUNE_MAP.put("gargantuan-e", new Rune("gargantuan-e", "Bruiser", "The Gargantuan gains the ability to periodically slam enemies, dealing 200% of your weapon damage as Physical and stunning them for 3 seconds."));
      SKILL_MAP.put("big-bad-voodoo", new Skill("big-bad-voodoo", "witchdoctor_bigbadvoodoo", "Big Bad Voodoo"));
      RUNE_MAP.put("big-bad-voodoo-a", new Rune("big-bad-voodoo-a", "Slam Dance", "The Fetish increases the damage of all nearby allies by 30%."));
      RUNE_MAP.put("big-bad-voodoo-b", new Rune("big-bad-voodoo-b", "Jungle Drums", "Increases the duration of the ritual to 30 seconds."));
      RUNE_MAP.put("big-bad-voodoo-c", new Rune("big-bad-voodoo-c", "Ghost Trance", "The ritual heals all nearby allies for 5% of their maximum Life per second."));
      RUNE_MAP.put("big-bad-voodoo-d", new Rune("big-bad-voodoo-d", "Rain Dance", "The ritual restores 123 Mana per second while standing in the ritual area."));
      RUNE_MAP.put("big-bad-voodoo-e", new Rune("big-bad-voodoo-e", "Boogie Man", "Enemies who die in the ritual area have a 50% chance to resurrect as a Zombie Dog."));
      SKILL_MAP.put("fetish-army", new Skill("fetish-army", "witchdoctor_fetisharmy", "Fetish Army"));
      RUNE_MAP.put("fetish-army-a", new Rune("fetish-army-a", "Fetish Ambush", "Each Fetish deals 250% weapon damage as Physical to any nearby enemy as it is summoned."));
      RUNE_MAP.put("fetish-army-b", new Rune("fetish-army-b", "Legion of Daggers", "Increases number of dagger-wielding Fetishes summoned by 3."));
      RUNE_MAP.put("fetish-army-c", new Rune("fetish-army-c", "Tiki Torchers", "Summon an additional 2 Fetish casters who breathe fire in a cone in front of them that deals 15% of your weapon damage as Fire."));
      RUNE_MAP.put("fetish-army-d", new Rune("fetish-army-d", "Devoted Following", "Decreases the cooldown of Fetish Army to 90 seconds."));
      RUNE_MAP.put("fetish-army-e", new Rune("fetish-army-e", "Head Hunters", "Summon an additional 2 Hunter Fetishes that shoot blowdarts at enemies, dealing 20% of your weapon damage as Poison."));
      SKILL_MAP.put("bash", new Skill("bash", "barbarian_bash", "Bash"));
      RUNE_MAP.put("bash-a", new Rune("bash-a", "Onslaught", "Add 2 reverberations that cause 25% weapon damage per strike. Removes the chance for Knockback."));
      RUNE_MAP.put("bash-b", new Rune("bash-b", "Punish", "Increases the damage of your skills by 8% for 5 seconds after using Bash. This effect stacks up to 3 times."));
      RUNE_MAP.put("bash-c", new Rune("bash-c", "Clobber", "Instead of Knockback, each hit has a 35% chance to Stun the target for 1.5 seconds."));
      RUNE_MAP.put("bash-d", new Rune("bash-d", "Instigation", "Generate 4 additional Fury per attack."));
      RUNE_MAP.put("bash-e", new Rune("bash-e", "Pulverize", "Cause a shockwave that inflicts 38% weapon damage to enemies in a 26 yard line behind the targeted enemy."));
      SKILL_MAP.put("hammer-of-the-ancients", new Skill("hammer-of-the-ancients", "barbarian_hammeroftheancients", "Hammer of the Ancients"));
      RUNE_MAP.put("hammer-of-the-ancients-a", new Rune("hammer-of-the-ancients-a", "Smash", "Strike a smaller area for 406% weapon damage."));
      RUNE_MAP.put("hammer-of-the-ancients-b", new Rune("hammer-of-the-ancients-b", "Rolling Thunder", "Create a shockwave that deals 275% weapon damage to all enemies within 22 yards in front of you."));
      RUNE_MAP.put("hammer-of-the-ancients-c", new Rune("hammer-of-the-ancients-c", "The Devil's Anvil", "Create a tremor at the point of impact for 2 seconds that slows the movement speed of enemies by 80%."));
      RUNE_MAP.put("hammer-of-the-ancients-d", new Rune("hammer-of-the-ancients-d", "Birthright", "Critical Hits have a 10% chance to cause enemies to drop treasure or health globes."));
      RUNE_MAP.put("hammer-of-the-ancients-e", new Rune("hammer-of-the-ancients-e", "Thunderstrike", "Whenever you kill an enemy with Hammer of the Ancients every other enemy within 10 yards is stunned for 3 seconds."));
      SKILL_MAP.put("cleave", new Skill("cleave", "barbarian_cleave", "Cleave"));
      RUNE_MAP.put("cleave-a", new Rune("cleave-a", "Broad Sweep", "Increase damage to 175% weapon damage."));
      RUNE_MAP.put("cleave-b", new Rune("cleave-b", "Gathering Storm", "Enemies cleaved have their movement speed reduced by 80% for 1 seconds."));
      RUNE_MAP.put("cleave-c", new Rune("cleave-c", "Scattering Blast", "On Critical Hits, knock enemies back 9 yards and inflict 60% weapon damage to enemies where they land."));
      RUNE_MAP.put("cleave-d", new Rune("cleave-d", "Reaping Swing", "Generate 3 additional Fury per enemy hit."));
      RUNE_MAP.put("cleave-e", new Rune("cleave-e", "Rupture", "Enemies slain by Cleave explode, causing 85% weapon damage to all other enemies within 8 yards."));
      SKILL_MAP.put("ground-stomp", new Skill("ground-stomp", "barbarian_groundstomp", "Ground Stomp"));
      RUNE_MAP.put("ground-stomp-a", new Rune("ground-stomp-a", "Trembling Stomp", "Enemies in the area also take 76% weapon damage."));
      RUNE_MAP.put("ground-stomp-b", new Rune("ground-stomp-b", "Wrenching Smash", "Increase the area of effect to 24 yards. Enemies are pulled closer before the strike lands."));
      RUNE_MAP.put("ground-stomp-c", new Rune("ground-stomp-c", "Avalanche", "Enemies are knocked back 9 yards and inflict 55% weapon damage to enemies in the landing area."));
      RUNE_MAP.put("ground-stomp-d", new Rune("ground-stomp-d", "Foot of the Mountain", "Increase Fury gained to 30."));
      RUNE_MAP.put("ground-stomp-e", new Rune("ground-stomp-e", "Deafening Crash", "Enemies in the area have their movement speed slowed by 60% for 3 seconds after they recover from being stunned."));
      SKILL_MAP.put("rend", new Skill("rend", "barbarian_rend", "Rend"));
      RUNE_MAP.put("rend-a", new Rune("rend-a", "Lacerate", "Increase damage to 903% weapon damage as Physical over 5 seconds."));
      RUNE_MAP.put("rend-b", new Rune("rend-b", "Ravage", "Increase the range of Rend to hit all enemies within 17 yards."));
      RUNE_MAP.put("rend-c", new Rune("rend-c", "Mutilate", "Affected enemies also have their movement speed reduced by 60%."));
      RUNE_MAP.put("rend-d", new Rune("rend-d", "Blood Lust", "Gain 9% of the damage done by Rend as Life."));
      RUNE_MAP.put("rend-e", new Rune("rend-e", "Bloodbath", "Enemies killed while bleeding cause all enemies within 10 yards to begin bleeding for 100% weapon damage as Physical over 5 seconds."));
      SKILL_MAP.put("leap", new Skill("leap", "barbarian_leap", "Leap"));
      RUNE_MAP.put("leap-a", new Rune("leap-a", "Call of Arreat", "Shockwaves burst forth from the ground increasing the radius of effect to 16 yards and pulling affected enemies towards you."));
      RUNE_MAP.put("leap-b", new Rune("leap-b", "Toppling Impact", "Send enemies hurtling away from where you land."));
      RUNE_MAP.put("leap-c", new Rune("leap-c", "Launch", "Jump into the air with such great force that enemies within 8 yards of the origin of the jump are also slowed by 60% for 3 seconds."));
      RUNE_MAP.put("leap-d", new Rune("leap-d", "Iron Impact", "Gain 100% additional Armor per enemy hit for 3 seconds after landing."));
      RUNE_MAP.put("leap-e", new Rune("leap-e", "Death from Above", "Land with such force that enemies have a 100% chance to be stunned for 3 seconds."));
      SKILL_MAP.put("ancient-spear", new Skill("ancient-spear", "barbarian_ancientspear", "Ancient Spear"));
      RUNE_MAP.put("ancient-spear-a", new Rune("ancient-spear-a", "Harpoon", "Pierce through multiple enemies in a straight line and drag them all back."));
      RUNE_MAP.put("ancient-spear-b", new Rune("ancient-spear-b", "Grappling Hooks", "Throw 3 spears. Each spear will pull back the enemy that it hits."));
      RUNE_MAP.put("ancient-spear-c", new Rune("ancient-spear-c", "Dread Spear", "Gain Life equal to 60% of the damage inflicted."));
      RUNE_MAP.put("ancient-spear-d", new Rune("ancient-spear-d", "Skirmish", "Increases Fury gained to 30."));
      RUNE_MAP.put("ancient-spear-e", new Rune("ancient-spear-e", "Rage Flip", "Enemies hit with Ancient Spear are pulled in the opposite direction and damage is increased to 213% weapon damage."));
      SKILL_MAP.put("frenzy", new Skill("frenzy", "barbarian_frenzy", "Frenzy"));
      RUNE_MAP.put("frenzy-a", new Rune("frenzy-a", "Maniac", "Each Frenzy effect also increases your damage by 4%."));
      RUNE_MAP.put("frenzy-b", new Rune("frenzy-b", "Sidearm", "Each strike has a 25% chance to throw a piercing axe at a nearby enemy that deals 110% weapon damage to all enemies in its path."));
      RUNE_MAP.put("frenzy-c", new Rune("frenzy-c", "Vanguard", "While under the effects of Frenzy, you gain 15% increased movement speed."));
      RUNE_MAP.put("frenzy-d", new Rune("frenzy-d", "Smite", "Add a 20% chance to call down a bolt of lightning from above, stunning your target for 1.5 seconds."));
      RUNE_MAP.put("frenzy-e", new Rune("frenzy-e", "Triumph", "Killing an enemy with Frenzy heals you for 8% of your maximum Life over 6 seconds."));
      SKILL_MAP.put("seismic-slam", new Skill("seismic-slam", "barbarian_seismicslam", "Seismic Slam"));
      RUNE_MAP.put("seismic-slam-a", new Rune("seismic-slam-a", "Shattered Ground", "Increase damage to 288% weapon damage and increases Knockback distance by 100%."));
      RUNE_MAP.put("seismic-slam-b", new Rune("seismic-slam-b", "Rumble", "The ground continues to shudder after the initial strike, damaging enemies in the area for 60% weapon damage as Physical over 2 seconds."));
      RUNE_MAP.put("seismic-slam-c", new Rune("seismic-slam-c", "Stagger", "Add a 70% chance of stunning enemies for 1.5 seconds."));
      RUNE_MAP.put("seismic-slam-d", new Rune("seismic-slam-d", "Strength from Earth", "Reduce Fury cost to 15 Fury."));
      RUNE_MAP.put("seismic-slam-e", new Rune("seismic-slam-e", "Cracking Rift", "Focus the seismic shockwaves along a narrow path to inflict 340% weapon damage to targets along a 42 yard path."));
      SKILL_MAP.put("revenge", new Skill("revenge", "barbarian_revenge", "Revenge"));
      RUNE_MAP.put("revenge-a", new Rune("revenge-a", "Retribution", "Increase damage to 286% weapon damage."));
      RUNE_MAP.put("revenge-b", new Rune("revenge-b", "Provocation", "Increases the chance Revenge will become active to 30% each time you are hit by an attack."));
      RUNE_MAP.put("revenge-c", new Rune("revenge-c", "Grudge", "Knocks enemies back 24 yards whenever Revenge is used."));
      RUNE_MAP.put("revenge-d", new Rune("revenge-d", "Vengeance Is Mine", "Gain 5 Fury and heal for 8.0% of your maximum Life for each enemy hit."));
      RUNE_MAP.put("revenge-e", new Rune("revenge-e", "Best Served Cold", "After using Revenge, your Critical Hit Chance is increased by 10% for 12 seconds."));
      SKILL_MAP.put("weapon-throw", new Skill("weapon-throw", "barbarian_weaponthrow", "Weapon Throw"));
      RUNE_MAP.put("weapon-throw-a", new Rune("weapon-throw-a", "Mighty Throw", "Increase thrown weapon damage to 169% weapon damage."));
      RUNE_MAP.put("weapon-throw-b", new Rune("weapon-throw-b", "Ricochet", "Cause the weapon to ricochet and hit up to 3 targets within 20 yards of each other."));
      RUNE_MAP.put("weapon-throw-c", new Rune("weapon-throw-c", "Throwing Hammer", "Hurl a hammer with a 50% chance to Stun the target for 1.5 seconds."));
      RUNE_MAP.put("weapon-throw-d", new Rune("weapon-throw-d", "Dread Bomb", "Expend all remaining Fury to throw a corpse which inflicts an additional 3% weapon damage for each point of Fury expended to all enemies within 12 yards of the target."));
      RUNE_MAP.put("weapon-throw-e", new Rune("weapon-throw-e", "Stupefy", "Aim for the head, gaining a 20% chance of causing your target to be Confused and attack other enemies for 6 seconds."));
      SKILL_MAP.put("sprint", new Skill("sprint", "barbarian_sprint", "Sprint"));
      RUNE_MAP.put("sprint-a", new Rune("sprint-a", "Marathon", "Increases the movement speed bonus to 50% for 5 seconds."));
      RUNE_MAP.put("sprint-b", new Rune("sprint-b", "Rush", "Increases Dodge Chance by 12% while sprinting."));
      RUNE_MAP.put("sprint-c", new Rune("sprint-c", "Run Like the Wind", "Tornadoes rage in your wake, each one inflicting 60% weapon damage as Physical for 3 seconds to nearby enemies."));
      RUNE_MAP.put("sprint-d", new Rune("sprint-d", "Forced March", "Increase the movement speed of allies within 50 yards by 20% for 3 seconds."));
      RUNE_MAP.put("sprint-e", new Rune("sprint-e", "Gangway", "Slams through enemies, knocking them back and inflicting 25% weapon damage."));
      SKILL_MAP.put("threatening-shout", new Skill("threatening-shout", "barbarian_threateningshout", "Threatening Shout"));
      RUNE_MAP.put("threatening-shout-a", new Rune("threatening-shout-a", "Demoralize", "Affected enemies are also taunted to attack you for 3 seconds."));
      RUNE_MAP.put("threatening-shout-b", new Rune("threatening-shout-b", "Intimidate", "Affected enemies also have their movement speed reduced by 30%."));
      RUNE_MAP.put("threatening-shout-c", new Rune("threatening-shout-c", "Grim Harvest", "Enemies are badly shaken and have a 15% chance to drop additional treasure or health globes."));
      RUNE_MAP.put("threatening-shout-d", new Rune("threatening-shout-d", "Falter", "Affected enemies also have their attack speed reduced by 15% for 5 seconds."));
      RUNE_MAP.put("threatening-shout-e", new Rune("threatening-shout-e", "Terrify", "Enemies are severely demoralized. Each enemy has a 35% chance to flee for 2.5 seconds."));
      SKILL_MAP.put("earthquake", new Skill("earthquake", "barbarian_earthquake", "Earthquake"));
      RUNE_MAP.put("earthquake-a", new Rune("earthquake-a", "Aftershocks", "Secondary tremors knock enemies back and inflict 65% weapon damage as Fire."));
      RUNE_MAP.put("earthquake-b", new Rune("earthquake-b", "Giant's Stride", "Secondary tremors follow your movement and inflict 65% weapon damage as Fire."));
      RUNE_MAP.put("earthquake-c", new Rune("earthquake-c", "Chilling Earth", "Creates an icy patch, causing Earthquake's damage to turn Cold and Slow the movement of enemies by 80%."));
      RUNE_MAP.put("earthquake-d", new Rune("earthquake-d", "The Mountain's Call", "Removes the Fury cost and reduces the cooldown to 105 seconds."));
      RUNE_MAP.put("earthquake-e", new Rune("earthquake-e", "Path of Fire", "Project secondary tremors up to 12 yards ahead of you that inflict 65% weapon damage as Fire."));
      SKILL_MAP.put("whirlwind", new Skill("whirlwind", "barbarian_whirlwind", "Whirlwind"));
      RUNE_MAP.put("whirlwind-a", new Rune("whirlwind-a", "Volcanic Eruption", "Turns Whirlwind into a torrent of magma that inflicts 189% weapon damage as Fire."));
      RUNE_MAP.put("whirlwind-b", new Rune("whirlwind-b", "Dust Devils", "Generate harsh tornadoes that inflict 40% weapon damage to enemies in their path."));
      RUNE_MAP.put("whirlwind-c", new Rune("whirlwind-c", "Hurricane", "Allows you to move at your movement speed while using Whirlwind."));
      RUNE_MAP.put("whirlwind-d", new Rune("whirlwind-d", "Wind Shear", "Gain 1 Fury for every enemy struck."));
      RUNE_MAP.put("whirlwind-e", new Rune("whirlwind-e", "Blood Funnel", "Critical Hits restore 2.0% of your maximum Life."));
      SKILL_MAP.put("furious-charge", new Skill("furious-charge", "barbarian_furiouscharge", "Furious Charge"));
      RUNE_MAP.put("furious-charge-a", new Rune("furious-charge-a", "Battering Ram", "Increase damage at the destination to 283% weapon damage."));
      RUNE_MAP.put("furious-charge-b", new Rune("furious-charge-b", "Dreadnought", "Regain 8% of your maximum Life for each target hit by Furious Charge."));
      RUNE_MAP.put("furious-charge-c", new Rune("furious-charge-c", "Bull Rush", "Any targets who are critically hit by Furious Charge will be stunned for 2.5 seconds."));
      RUNE_MAP.put("furious-charge-d", new Rune("furious-charge-d", "Stamina", "Generate 8 additional Fury for each target hit while charging."));
      RUNE_MAP.put("furious-charge-e", new Rune("furious-charge-e", "Merciless Assault", "Cooldown is reduced by 2 seconds for every target hit. This effect can reduce the cooldown by up to 10 seconds."));
      SKILL_MAP.put("ignore-pain", new Skill("ignore-pain", "barbarian_ignorepain", "Ignore Pain"));
      RUNE_MAP.put("ignore-pain-a", new Rune("ignore-pain-a", "Contempt for Weakness", "Reflects 50% of ignored damage back at the enemy."));
      RUNE_MAP.put("ignore-pain-b", new Rune("ignore-pain-b", "Iron Hide", "Increases duration to 7 seconds."));
      RUNE_MAP.put("ignore-pain-c", new Rune("ignore-pain-c", "Mob Rule", "Extend the effect to nearby allies, reducing damage taken by 65% for 5 seconds."));
      RUNE_MAP.put("ignore-pain-d", new Rune("ignore-pain-d", "Bravado", "When activated, Knockback all enemies within 12 yards and deal 50% weapon damage to them."));
      RUNE_MAP.put("ignore-pain-e", new Rune("ignore-pain-e", "Ignorance is Bliss", "While Ignore Pain is active, gain 20% of all damage dealt as Life."));
      SKILL_MAP.put("battle-rage", new Skill("battle-rage", "barbarian_battlerage", "Battle Rage"));
      RUNE_MAP.put("battle-rage-a", new Rune("battle-rage-a", "Marauder's Rage", "Increase damage bonus to 30%."));
      RUNE_MAP.put("battle-rage-b", new Rune("battle-rage-b", "Ferocity", "While under the effects of Battle Rage, Critical Hits have a chance to increase the duration of Battle Rage by 2 seconds."));
      RUNE_MAP.put("battle-rage-c", new Rune("battle-rage-c", "Swords to Ploughshares", "While under the effects of Battle Rage, Critical Hits have up to a 5% chance to cause enemies to drop additional health globes."));
      RUNE_MAP.put("battle-rage-d", new Rune("battle-rage-d", "Into the Fray", "While under the effects of Battle Rage, Critical Hits have a chance to generate 15 additional Fury."));
      RUNE_MAP.put("battle-rage-e", new Rune("battle-rage-e", "Bloodshed", "While under the effects of Battle Rage, Critical Hits have a chance to cause an explosion of blood dealing 20% of the damage done to all other nearby enemies."));
      SKILL_MAP.put("call-of-the-ancients", new Skill("call-of-the-ancients", "barbarian_calloftheancients", "Call of the Ancients"));
      RUNE_MAP.put("call-of-the-ancients-a", new Rune("call-of-the-ancients-a", "Korlic's Might", "Korlic gains the skill Furious Charge which deals 200% of your weapon damage to all enemies in a line."));
      RUNE_MAP.put("call-of-the-ancients-b", new Rune("call-of-the-ancients-b", "The Council Rises", "The Ancients inflict 66% weapon damage with each attack and have 100% additional Armor."));
      RUNE_MAP.put("call-of-the-ancients-c", new Rune("call-of-the-ancients-c", "Madawc's Madness", "Madawc gains the skill Seismic Slam which deals 180% of your weapon damage to enemies in an arc."));
      RUNE_MAP.put("call-of-the-ancients-d", new Rune("call-of-the-ancients-d", "Duty to the Clan", "Increase duration to 20 seconds."));
      RUNE_MAP.put("call-of-the-ancients-e", new Rune("call-of-the-ancients-e", "Talic's Anger", "Talic gains the skill Leap which deals 250% of your weapon damage to enemies in the area of the leap."));
      SKILL_MAP.put("overpower", new Skill("overpower", "barbarian_overpower", "Overpower"));
      RUNE_MAP.put("overpower-a", new Rune("overpower-a", "Killing Spree", "Your Critical Hit Chance is increased by 10% for 6 seconds."));
      RUNE_MAP.put("overpower-b", new Rune("overpower-b", "Storm of Steel", "Throw up to 3 axes at nearby enemies which inflict 50% weapon damage each."));
      RUNE_MAP.put("overpower-c", new Rune("overpower-c", "Revel", "Heal 8% of your maximum Life for every enemy hit."));
      RUNE_MAP.put("overpower-d", new Rune("overpower-d", "Momentum", "Generate 12 Fury for each enemy hit by Overpower."));
      RUNE_MAP.put("overpower-e", new Rune("overpower-e", "Crushing Advance", "Redirect 35% of incoming melee and ranged damage for 4 seconds after Overpower is activated."));
      SKILL_MAP.put("war-cry", new Skill("war-cry", "barbarian_warcry", "War Cry"));
      RUNE_MAP.put("war-cry-a", new Rune("war-cry-a", "Hardened Wrath", "Increases the Armor bonus to 40%."));
      RUNE_MAP.put("war-cry-b", new Rune("war-cry-b", "Veteran's Warning", "War Cry also grants a 15% bonus to Dodge Chance."));
      RUNE_MAP.put("war-cry-c", new Rune("war-cry-c", "Impunity", "All of your resistances are increased by 20% while affected by War Cry."));
      RUNE_MAP.put("war-cry-d", new Rune("war-cry-d", "Charge!", "Increases Fury gained to 40."));
      RUNE_MAP.put("war-cry-e", new Rune("war-cry-e", "Invigorate", "Increases maximum Life by 10% and regenerates 620 Life per second while affected by War Cry."));
      SKILL_MAP.put("wrath-of-the-berserker", new Skill("wrath-of-the-berserker", "barbarian_wrathoftheberserker", "Wrath of the Berserker"));
      RUNE_MAP.put("wrath-of-the-berserker-a", new Rune("wrath-of-the-berserker-a", "Insanity", "While active your damage is also increased by 100%."));
      RUNE_MAP.put("wrath-of-the-berserker-b", new Rune("wrath-of-the-berserker-b", "Arreat's Wail", "Activating Wrath of the Berserker knocks back all enemies within 12 yards and deals 430% weapon damage to them."));
      RUNE_MAP.put("wrath-of-the-berserker-c", new Rune("wrath-of-the-berserker-c", "Striding Giant", "Increases bonus to Dodge Chance to 60%."));
      RUNE_MAP.put("wrath-of-the-berserker-d", new Rune("wrath-of-the-berserker-d", "Thrive on Chaos", "Every 25 Fury gained while Wrath of the Berserker is active adds 1 second to the duration of the effect."));
      RUNE_MAP.put("wrath-of-the-berserker-e", new Rune("wrath-of-the-berserker-e", "Slaughter", "While Wrath of the Berserker is active, Critical Hits have a chance to cause an eruption of blood dealing 155% weapon damage to enemies within 15 yards."));
   }
}
