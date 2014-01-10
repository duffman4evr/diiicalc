package com.diiicalc.core.resources;

import com.diiicalc.api.CareerProfile;
import com.diiicalc.core.BattlenetRealm;
import com.diiicalc.core.Constants;
import com.diiicalc.core.Utils;
import com.yammer.metrics.annotation.Timed;

import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;

@Path("/career-profiles")
@Produces(MediaType.APPLICATION_JSON)
public class CareerProfiles
{
   @GET
   @Timed
   @Path("/{battleTag}")
   public CareerProfile getSingle(@PathParam("battleTag") String battleTag, @QueryParam("realm") BattlenetRealm realm) throws Exception
   {
      String path = Constants.PROFILE_API_URL_PREFIX + "/" + battleTag + "/";

      return Utils.doGet(realm, path, CareerProfile.class);
   }
}