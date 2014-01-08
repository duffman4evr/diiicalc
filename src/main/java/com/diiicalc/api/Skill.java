package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown=true)
public class Skill implements Comparable<Skill>
{
   @JsonProperty("slug")
   private String slug;

   @JsonProperty("icon")
   private String icon;

   @JsonProperty("name")
   private String name;

   @JsonProperty("tooltipUrl")
   private String tooltipUrl;

   @JsonProperty("skillCalcId")
   private String skillCalcId;

   public Skill() { }

   public Skill(String slug, String icon)
   {
      this.slug = slug;
      this.icon = icon;
   }

   public Skill(String slug, String icon, String name)
   {
      this.slug = slug;
      this.icon = icon;
      this.name = name;
   }

   public String getSlug()
   {
      return slug;
   }

   public String getIcon()
   {
      return icon;
   }

   public String getName()
   {
      return name;
   }

   public String getTooltipUrl()
   {
      return tooltipUrl;
   }

   public String getSkillCalcId()
   {
      return skillCalcId;
   }

   @Override
   public boolean equals(Object o)
   {
      if (this == o)
      {
         return true;
      }
      if (o == null || getClass() != o.getClass())
      {
         return false;
      }

      Skill skill = (Skill) o;

      if (slug != null ? !slug.equals(skill.slug) : skill.slug != null)
      {
         return false;
      }

      return true;
   }

   @Override
   public int hashCode()
   {
      return slug != null ? slug.hashCode() : 0;
   }

   @Override
   public int compareTo(Skill o)
   {
      return this.slug.compareTo(o.slug);
   }
}
