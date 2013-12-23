package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown=true)
public class Skill
{
   @JsonProperty("slug")
   private String slug;

   @JsonProperty("icon")
   private String icon;

   @JsonProperty("tooltipUrl")
   private String tooltipParams;

   @JsonProperty("skillCalcId")
   private String skillCalcId;

   public String getSlug()
   {
      return slug;
   }

   public String getIcon()
   {
      return icon;
   }

   public String getTooltipParams()
   {
      return tooltipParams;
   }

   public String getSkillCalcId()
   {
      return skillCalcId;
   }
}
