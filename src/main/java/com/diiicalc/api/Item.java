package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.HashMap;
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

   @JsonProperty("tooltipParams")
   private String tooltipParams;

   @JsonProperty("type")
   private Type type;

   @JsonProperty("attributesRaw")
   private Map<String, ValueRange> attributesRaw = new HashMap<String, ValueRange>();

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

   public static class Type
   {
      @JsonProperty("twoHanded")
      private boolean twoHanded;

      @JsonProperty("id")
      private String id;

      public boolean isTwoHanded()
      {
         return twoHanded;
      }

      public String getId()
      {
         return id;
      }
   }
}
