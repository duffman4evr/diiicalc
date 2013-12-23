package com.diiicalc.core.commands;

import com.diiicalc.api.DefensiveSummary;
import com.diiicalc.api.Hero;
import com.diiicalc.api.Item;
import com.diiicalc.core.*;
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
   public DefensiveSummary getSingle(@PathParam("heroId") long heroId, @QueryParam("battleTag") String battleTag) throws Exception
   {
      if (battleTag == null)
      {
         throw new WebApplicationException(Response.status(Response.Status.BAD_REQUEST)
            .type(MediaType.TEXT_PLAIN_TYPE)
            .entity("No BattleTag given.")
            .build());
      }

      String heroPath = Constants.PROFILE_API_URL_PREFIX + "/" + battleTag + "/hero/" + heroId;

      Hero hero = Utils.doGet(BattlenetRealm.US, heroPath, Hero.class);

      Map<String, Item> itemMap = Utils.pullDownItemMap(hero);

      StatTotals statTotals = new StatTotals(itemMap, hero);

      DefensiveStats normalDefensiveStats = Utils.computeDefensiveStats(statTotals);

      statTotals.addResistAll(1);

      DefensiveStats resistAllValue = Utils.computeDefensiveStats(statTotals);

      statTotals.addResistAll(-1);
      statTotals.addVitality(1);

      DefensiveStats vitalityValue = Utils.computeDefensiveStats(statTotals);

      statTotals.addVitality(-1);
      statTotals.addArmor(1);

      DefensiveStats armorValue = Utils.computeDefensiveStats(statTotals);

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
         vitalityValue.getEffectiveLife() - normalDefensiveStats.getEffectiveLife()
      );
   }
}
