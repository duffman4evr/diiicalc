package com.diiicalc.core.resources;

import com.diiicalc.api.Hero;
import com.diiicalc.core.BattlenetRealm;
import com.diiicalc.core.Constants;
import com.diiicalc.core.Utils;
import com.yammer.metrics.annotation.Timed;

import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/heroes")
@Produces(MediaType.APPLICATION_JSON)
public class Heroes
{
   @GET
   @Timed
   @Path("/{heroId}")
   public Hero getSingle(@PathParam("heroId") long heroId, @QueryParam("battleTag") String battleTag) throws Exception
   {
      if (battleTag == null)
      {
         throw new WebApplicationException(Response.status(Response.Status.BAD_REQUEST)
            .type(MediaType.TEXT_PLAIN_TYPE)
            .entity("No BattleTag given.")
            .build());
      }

      String path = Constants.PROFILE_API_URL_PREFIX + "/" + battleTag + "/" + heroId;

      return Utils.doGet(BattlenetRealm.US, path, Hero.class);
   }
}
