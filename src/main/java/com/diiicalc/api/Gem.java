package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.HashMap;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown=true)
public class Gem
{
   @JsonProperty("item")
   private Item item;

   @JsonProperty("attributesRaw")
   private Map<String, ValueRange> attributesRaw = new HashMap<String, ValueRange>();

   public Item getItem()
   {
      return item;
   }

   public Map<String, ValueRange> getAttributesRaw()
   {
      return attributesRaw;
   }
}
