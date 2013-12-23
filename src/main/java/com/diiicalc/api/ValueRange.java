package com.diiicalc.api;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ValueRange
{
   @JsonProperty("min")
   private double min;

   @JsonProperty("max")
   private double max;

   public ValueRange() { }

   public ValueRange(double min, double max)
   {
      this.min = min;
      this.max = max;
   }

   public double getMin()
   {
      return min;
   }

   public double getMax()
   {
      return max;
   }

   public ValueRange scale(double factor)
   {
      return new ValueRange(this.min * factor, this.max * factor);
   }

   public ValueRange add(ValueRange other)
   {
      return new ValueRange(this.min + other.getMin(), this.max + other.getMax());
   }

   public static ValueRange fromMinAndDelta(ValueRange min, ValueRange delta)
   {
      return new ValueRange(min.getMin(), min.getMin() + delta.getMin());
   }

   public static ValueRange fromMinAndMax(ValueRange min, ValueRange max)
   {
      return new ValueRange(min.getMin(), max.getMin());
   }

   public double getAverage()
   {
      return (this.min + this.max) * 0.5;
   }

   public String toString()
   {
      return this.min + " - " + this.max;
   }
}
