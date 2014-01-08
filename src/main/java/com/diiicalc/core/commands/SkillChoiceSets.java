package com.diiicalc.core.commands;

import com.diiicalc.api.*;
import com.diiicalc.core.*;
import com.yammer.metrics.annotation.Timed;

import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Map;

@Path("/skill-choice-sets")
@Produces(MediaType.APPLICATION_JSON)
public class SkillChoiceSets
{
   @GET
   @Timed
   @Path("/{heroId}")
   public SkillChoiceSet getSingle(@PathParam("heroId") long heroId, @QueryParam("battleTag") String battleTag) throws Exception
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

      RelevantSkillSet relevantSkillSet = Utils.getRelevantSkillSetForHeroType(hero.getType());

      SkillChoiceSet skillChoiceSet = new SkillChoiceSet(relevantSkillSet);

      skillChoiceSet.addSkillChoices(hero.getSkills());

      return skillChoiceSet;
   }
}
