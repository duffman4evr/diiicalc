package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.ArrayList;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown=true)
public class CareerProfile
{
   @JsonProperty("heroes")
   private List<HeroSummary> heroes = new ArrayList<HeroSummary>();

   public List<HeroSummary> getHeroes()
   {
      return heroes;
   }

   @JsonIgnoreProperties(ignoreUnknown=true)
   public static class HeroSummary
   {
      @JsonProperty("id")
      private long id;

      @JsonProperty("name")
      private String name;

      @JsonProperty("level")
      private int level;

      @JsonProperty("paragonLevel")
      private int paragonLevel;

      @JsonProperty("class")
      private String type;

      @JsonProperty("hardcore")
      private boolean hardcore;

      @JsonProperty("dead")
      private boolean dead;

      @JsonProperty("gender")
      private int gender;

      public long getId()
      {
         return id;
      }

      public String getName()
      {
         return name;
      }

      public int getLevel()
      {
         return level;
      }

      public int getParagonLevel()
      {
         return paragonLevel;
      }

      public String getType()
      {
         return type;
      }

      public boolean isHardcore()
      {
         return hardcore;
      }

      public boolean isDead()
      {
         return dead;
      }

      public int getGender()
      {
         return gender;
      }
   }
}
