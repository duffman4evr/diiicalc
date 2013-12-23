package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown=true)
public class BlizzardApiError
{
   @JsonProperty("code")
   private String code;

   @JsonProperty("reason")
   private String reason;

   public String getCode()
   {
      return code;
   }

   public String getReason()
   {
      return reason;
   }
}
