package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.ArrayList;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown=true)
public class Set
{
   @JsonProperty("slug")
   private String slug;

   @JsonProperty("ranks")
   private List<SetRank> ranks = new ArrayList<SetRank>();

   public String getSlug()
   {
      return slug;
   }

   public List<SetRank> getRanks()
   {
      return ranks;
   }
}
