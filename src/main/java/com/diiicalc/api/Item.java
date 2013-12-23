package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown=true)
public class Item
{
   @JsonProperty("id")
   private String id;

   @JsonProperty("name")
   private String name;

   @JsonProperty("icon")
   private String icon;

   @JsonProperty("armor")
   private ValueRange armor;

   @JsonProperty("tooltipParams")
   private String tooltipParams;

   @JsonProperty("type")
   private Type type;

   @JsonProperty("attributesRaw")
   private Map<String, ValueRange> attributesRaw = new HashMap<String, ValueRange>();

   @JsonProperty("gems")
   private List<Gem> gems = new ArrayList<Gem>();

   @JsonProperty("set")
   private Set set;

   public String getId()
   {
      return id;
   }

   public String getName()
   {
      return name;
   }

   public String getIcon()
   {
      return icon;
   }

   public ValueRange getArmor()
   {
      return armor;
   }

   public String getTooltipParams()
   {
      return tooltipParams;
   }

   public Type getType()
   {
      return type;
   }

   public Map<String, ValueRange> getAttributesRaw()
   {
      return attributesRaw;
   }

   public List<Gem> getGems()
   {
      return gems;
   }

   public Set getSet()
   {
      return set;
   }

   public static class Type
   {
      @JsonProperty("twoHanded")
      private boolean twoHanded;

      @JsonProperty("id")
      private String id;

      public Type() { }

      public Type(String id, boolean twoHanded)
      {
         this.id = id;
         this.twoHanded = twoHanded;
      }

      public boolean isTwoHanded()
      {
         return twoHanded;
      }

      public String getId()
      {
         return id;
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

         Type type = (Type) o;

         if (twoHanded != type.twoHanded)
         {
            return false;
         }
         if (id != null ? !id.equals(type.id) : type.id != null)
         {
            return false;
         }

         return true;
      }

      @Override
      public int hashCode()
      {
         int result = (twoHanded ? 1 : 0);
         result = 31 * result + (id != null ? id.hashCode() : 0);
         return result;
      }
   }
}
