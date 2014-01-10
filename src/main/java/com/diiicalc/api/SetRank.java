package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.HashMap;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown=true)
public class SetRank
{
   @JsonProperty("required")
   private int requiredOwnership;

   @JsonProperty("attributesRaw")
   private Map<String, ValueRange> attributesRaw = new HashMap<String, ValueRange>();

   public int getRequiredOwnership()
   {
      return requiredOwnership;
   }

   public Map<String, ValueRange> getAttributesRaw()
   {
      return attributesRaw;
   }
}
