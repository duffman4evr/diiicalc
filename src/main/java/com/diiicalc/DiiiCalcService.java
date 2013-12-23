package com.diiicalc;

import com.diiicalc.core.commands.DefensiveSummaries;
import com.diiicalc.core.resources.CareerProfiles;
import com.yammer.dropwizard.Service;
import com.yammer.dropwizard.assets.AssetsBundle;
import com.yammer.dropwizard.config.Bootstrap;
import com.yammer.dropwizard.config.Environment;

public class DiiiCalcService extends Service<DiiiCalcConfiguration>
{
   @Override
   public void initialize(Bootstrap<DiiiCalcConfiguration> bootstrap)
   {
      bootstrap.setName("dii-calc");
      bootstrap.addBundle(new AssetsBundle("/assets/", "/"));
   }

   @Override
   public void run(DiiiCalcConfiguration config, Environment environment) throws ClassNotFoundException
   {
      // Set up resources (top 2 are used by the Connectors).
      environment.addResource(new CareerProfiles());
      environment.addResource(new DefensiveSummaries());
   }

   public static void main(String[] args) throws Exception
   {
      new DiiiCalcService().run(args);
   }
}
