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

   public static final String[] ALL_WIZARD_SLUGS =
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

   public static final String[] ALL_BARBARIAN_SLUGS =
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

   static
   {
      // This code is GENERATED. See SkillDownloader for details.
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
