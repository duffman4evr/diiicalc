package com.diiicalc.core.commands;

import com.diiicalc.api.DefensiveSummary;
import com.diiicalc.api.Hero;
import com.diiicalc.api.Item;
import com.diiicalc.core.*;
import com.fasterxml.jackson.core.type.TypeReference;
import com.yammer.metrics.annotation.Timed;

import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Map;

@Path("/defensive-summaries")
@Produces(MediaType.APPLICATION_JSON)
public class DefensiveSummaries
{
   @GET
   @Timed
   @Path("/{heroId}")
   public DefensiveSummary getSingle
   (
      @PathParam("heroId") long heroId,
      @QueryParam("battleTag") String battleTag,
      @QueryParam("monsterLevel") Long monsterLevel,
      @QueryParam("realm") BattlenetRealm realm,
      @QueryParam("activeSkills") String activeSkillsJson,
      @QueryParam("passiveSkills") String passiveSkillsJson
   )
      throws Exception
   {
      if (battleTag == null)
      {
         throw new WebApplicationException(Response.status(Response.Status.BAD_REQUEST)
            .type(MediaType.TEXT_PLAIN_TYPE)
            .entity("No BattleTag given.")
            .build());
      }

      if (monsterLevel == null)
      {
         monsterLevel = 63L;
      }

      Map<String, String> activeSkillOverrides = null;
      Map<String, String> passiveSkillOverrides = null;

      if (activeSkillsJson != null)
      {
         activeSkillOverrides = Utils.JSON_MAPPER.readValue(activeSkillsJson, new TypeReference<Map<String,String>>() { });
      }

      if (passiveSkillsJson != null)
      {
         passiveSkillOverrides = Utils.JSON_MAPPER.readValue(passiveSkillsJson, new TypeReference<Map<String,String>>() { });
      }

      String heroPath = Constants.PROFILE_API_URL_PREFIX + "/" + battleTag + "/hero/" + heroId;

      Hero hero = Utils.doGet(realm, heroPath, Hero.class);

      Map<String, Item> itemMap = Utils.pullDownItemMap(hero, realm);

      StatTotals statTotals = new StatTotals(itemMap, hero, activeSkillOverrides, passiveSkillOverrides);

      DefensiveStats normalDefensiveStats = Utils.computeDefensiveStats(statTotals, monsterLevel);

      statTotals.addResistAll(1);

      DefensiveStats resistAllValue = Utils.computeDefensiveStats(statTotals, monsterLevel);

      statTotals.addResistAll(-1);
      statTotals.addVitality(1);

      DefensiveStats vitalityValue = Utils.computeDefensiveStats(statTotals, monsterLevel);

      statTotals.addVitality(-1);
      statTotals.addArmor(1);

      DefensiveStats armorValue = Utils.computeDefensiveStats(statTotals, monsterLevel);

      statTotals.addArmor(-1);
      statTotals.addPercentLife(0.01);

      DefensiveStats percentLifeValue = Utils.computeDefensiveStats(statTotals, monsterLevel);

      return new DefensiveSummary
      (
         normalDefensiveStats.getEffectiveLife(),
         normalDefensiveStats.getArmor(),
         normalDefensiveStats.getResistAll(),
         normalDefensiveStats.getDodgeChance(),
         normalDefensiveStats.getTotalIncomingDamageModifier(),
         normalDefensiveStats.getIncomingDamageModifiers().get(Constants.DAMAGE_MODIFIER_RESISTS),
         normalDefensiveStats.getIncomingDamageModifiers().get(Constants.DAMAGE_MODIFIER_ARMOR),
         armorValue.getEffectiveLife() - normalDefensiveStats.getEffectiveLife(),
         resistAllValue.getEffectiveLife() - normalDefensiveStats.getEffectiveLife(),
         vitalityValue.getEffectiveLife() - normalDefensiveStats.getEffectiveLife(),
         percentLifeValue.getEffectiveLife() - normalDefensiveStats.getEffectiveLife()
      );
   }
}
