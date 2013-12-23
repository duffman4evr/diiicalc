package com.diiicalc.core.modifiers;

import java.util.Map;

public interface IncomingDamageModifier
{
   public void addModifier(Map<String, Double> defenseModifiers);
}
