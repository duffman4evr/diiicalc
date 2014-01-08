package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown=true)
public class Rune
{
   @JsonProperty("slug")
   private String slug;

   @JsonProperty("type")
   private String type;

   @JsonProperty("name")
   private String name;

   @JsonProperty("description")
   private String description;

   @JsonProperty("tooltipParams")
   private String tooltipParams;

   @JsonProperty("skillCalcId")
   private String skillCalcId;

   @JsonProperty("order")
   private long order;

   public Rune() { }

   public Rune(String slug, String name)
   {
      this.slug = slug;
      this.name = name;
   }

   public Rune(String slug, String name, String description)
   {
      this.slug = slug;
      this.name = name;
      this.description = description;
   }

   public String getSlug()
   {
      return slug;
   }

   public String getType()
   {
      return type;
   }

   public String getName()
   {
      return name;
   }

   public String getDescription()
   {
      return description;
   }

   public String getTooltipParams()
   {
      return tooltipParams;
   }

   public String getSkillCalcId()
   {
      return skillCalcId;
   }

   public long getOrder()
   {
      return order;
   }
}
