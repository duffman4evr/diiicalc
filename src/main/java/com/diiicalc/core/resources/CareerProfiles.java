package com.diiicalc.core.resources;

import com.diiicalc.api.CareerProfile;
import com.diiicalc.core.BattlenetRealm;
import com.diiicalc.core.Constants;
import com.diiicalc.core.Utils;
import com.yammer.metrics.annotation.Timed;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

@Path("/career-profiles")
@Produces(MediaType.APPLICATION_JSON)
public class CareerProfiles
{
   @GET
   @Timed
   @Path("/{battleTag}")
   public CareerProfile getSingle(@PathParam("battleTag") String battleTag) throws Exception
   {
      String path = Constants.PROFILE_API_URL_PREFIX + "/" + battleTag + "/";

      return Utils.doGet(BattlenetRealm.US, path, CareerProfile.class);
   }
}

/*

@Path("/admin/customers")
@Produces(MediaType.APPLICATION_JSON)
public class CustomerResource
{
   private CustomerDao dao;

   public CustomerResource(CustomerDao dao)
   {
      this.dao = dao;
   }

   @GET
   @Timed
   public List<CustomerVo> getAll(@Auth @BasicAuth AdminClient adminClient)
   {
      Utils.requireRole(adminClient, Constants.USER_ROLE_FE);

      return this.dao.getAll();
   }

   @GET
   @Timed
   @Path("/{customerId}")
   public CustomerVo getSingle(@Auth @BasicAuth AdminClient adminClient, @PathParam("customerId") long customerId)
   {
      Utils.requireRole(adminClient, Constants.USER_ROLE_FE);

      return this.dao.findByCustomerId(customerId);
   }

   @POST
   @Timed
   public void create(@Auth @BasicAuth AdminClient adminClient, CustomerVo customerVo)
   {
      Utils.requireRole(adminClient, Constants.USER_ROLE_FE);

      if ("".equals(customerVo.getName().trim()))
      {
         throw new WebApplicationException(Response.status(Response.Status.BAD_REQUEST)
            .type(MediaType.TEXT_PLAIN_TYPE)
            .entity("Must have a non-empty name.")
            .build());
      }

      CustomerVo duplicate = this.dao.findByName(customerVo.getName());

      if (duplicate != null)
      {
         throw new WebApplicationException(Response.status(Response.Status.CONFLICT)
            .entity("Name is a duplicate" + (duplicate.isDisabled() ? ", which is already in use by a disabled customer." : "."))
            .build());
      }

      this.dao.create(customerVo.getName());
   }

   @PUT
   @Timed
   @Path("/{customerId}")
   public void update(@Auth @BasicAuth AdminClient adminClient, @PathParam("customerId") long customerId, CustomerVo customerVo)
   {
      Utils.requireRole(adminClient, Constants.USER_ROLE_FE);

      if ("".equals(customerVo.getName().trim()))
      {
         throw new WebApplicationException(Response.status(Response.Status.BAD_REQUEST)
            .type(MediaType.TEXT_PLAIN_TYPE)
            .entity("Must have a non-empty name.")
            .build());
      }

      CustomerVo duplicate = this.dao.findByName(customerVo.getName());

      if (duplicate != null && duplicate.getCustomerId() != customerId)
      {
         throw new WebApplicationException(Response.status(Response.Status.CONFLICT)
            .entity("Name is a duplicate" + (duplicate.isDisabled() ? ", which is already in use by a disabled customer." : "."))
            .build());
      }

      this.dao.update(customerId, customerVo.getName());
   }

   @DELETE
   @Timed
   @Path("/{customerId}")
   public void delete(@Auth @BasicAuth AdminClient adminClient, @PathParam("customerId") long customerId)
   {
      Utils.requireRole(adminClient, Constants.USER_ROLE_ADMIN);

      this.dao.disable(customerId);
   }
}

*/