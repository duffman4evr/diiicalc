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

   @JsonProperty("tooltipParams")
   private String tooltipParams;

   @JsonProperty("skillCalcId")
   private String skillCalcId;

   @JsonProperty("order")
   private long order;

   public String getSlug()
   {
      return slug;
   }

   public String getType()
   {
      return type;
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
