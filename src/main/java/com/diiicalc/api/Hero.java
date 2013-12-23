package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown=true)
public class Hero
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

   @JsonProperty("items")
   private Map<String, ItemSummary> items;

   @JsonProperty("skills")
   private Skills skills;

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

   public Map<String, ItemSummary> getItems()
   {
      return items;
   }

   public Skills getSkills()
   {
      return skills;
   }

   @JsonIgnoreProperties(ignoreUnknown=true)
   public static class ItemSummary
   {
      @JsonProperty("tooltipParams")
      private String tooltipParams;

      public String getTooltipParams()
      {
         return tooltipParams;
      }
   }
}
