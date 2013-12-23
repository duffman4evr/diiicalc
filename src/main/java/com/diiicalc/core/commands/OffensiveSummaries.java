package com.diiicalc.core.commands;

import com.diiicalc.api.DefensiveSummary;
import com.diiicalc.api.Hero;
import com.diiicalc.api.Item;
import com.diiicalc.api.OffensiveSummary;
import com.diiicalc.core.*;
import com.yammer.metrics.annotation.Timed;

import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Map;

@Path("/offensive-summaries")
@Produces(MediaType.APPLICATION_JSON)
public class OffensiveSummaries
{
   @GET
   @Timed
   @Path("/{heroId}")
   public OffensiveSummary getSingle(@PathParam("heroId") long heroId, @QueryParam("battleTag") String battleTag) throws Exception
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

      OffensiveStats normalOffensiveStats = Utils.computeOffensiveStats(statTotals);

      statTotals.addPrimaryStat(1);

      OffensiveStats primaryStatValue = Utils.computeOffensiveStats(statTotals);

      statTotals.addPrimaryStat(-1);
      statTotals.addCritChance(1);

      OffensiveStats critChanceValue = Utils.computeOffensiveStats(statTotals);

      statTotals.addCritChance(-1);
      statTotals.addCritDamage(1);

      OffensiveStats critDamageValue = Utils.computeOffensiveStats(statTotals);

      statTotals.addCritDamage(-1);
      statTotals.addAttackSpeed(1);

      OffensiveStats attackSpeedValue = Utils.computeOffensiveStats(statTotals);

      return new OffensiveSummary
      (
         normalOffensiveStats.getDps(),
         normalOffensiveStats.getAttacksPerSecond(),
         normalOffensiveStats.getCritChance(),
         normalOffensiveStats.getCritDamage(),
         primaryStatValue.getDps() - normalOffensiveStats.getDps(),
         critChanceValue.getDps() - normalOffensiveStats.getDps(),
         critDamageValue.getDps() - normalOffensiveStats.getDps(),
         attackSpeedValue.getDps() - normalOffensiveStats.getDps()
      );
   }
}
