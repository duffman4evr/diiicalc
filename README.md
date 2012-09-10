RELEASE NOTES

------------VERSION 1.0:------------

===UPDATES===

- Initial release. Ability to query the D3 US API for data and calculate EHP.

------------VERSION 1.1:------------

===UPDATES===

- You can now look up BattleTags with both '-' and '#' format. Example: Mytag#1234 and Mytag-1234 both work now.
- Class names are now presented next to character names during character selection.
- You can now select EU / TW / KR region characters!
- Quick switch between characters in the EHP screen has been added. A drop-down at the top of the page allows you to switch.
- The calculator now leverages a new Life on Hit field exposed by the Diablo 3 API. Life on Hit will now be completely accurate, and the calculator should now be making less API calls overall.
- Enchantress follower armor bonus added.
- Shield blocking is now part of the EHP calculation. I'm not totally confident on my math though, so looking for testers on this.
- EHP with dodge is now shown for Monks and Demon Hunters. In these cases, dodge is simply shown as another mitigation source.
- Look and feel has been updated a tiny bit. A slighly grey background should make the calculator easier on the eyes ;)
- Back button functionality has been improved. No more 'confirm resubmit' issues.

------------VERSION 1.2:------------

===UPDATES===

- System-wide issues with calculating Armor/Resistance bonuses due to skills have been fixed. HUGE thanks to ANBUCyrus for finding this out.
- The calculator now reports accurate numbers for characters that logged off with defensive passives. Note: logging off with an active Enchantress with the armor buff will still cause problems.
- Dodge is now factored in to EHP for all hero class types.
- The bottom comparison section has been updated to allow you to select which stat to compare. Life % bonuses and Dexterity have been added.
- 'Leap - Iron Impact' was previously giving a bonus of only 200% armor, rather than 300%. This has been fixed. Bug found by: ANBUCyrus.
- 'Fists of Thunder' support added.
- 'The Guardian's Path' support added.
- 'Mantra of Evasion' now affects Dodge %. Issue found by: TyrialFrost.
- Issue where 'Seize the Initiative' skill wasn't being detected was fixed. Bug found by: pmilkman.
- Life % is now factored in to EHP gains due to vitality gains.
- The Diablo 3 API mis-reports resistances for all characters by a small amount. A workaround for this has been implemented.

------------VERSION 1.3:------------

===UPDATES===

- A 'back to search' link has been added for convenience.
- Certain common group buffs like War Cry and Mantra of Evasion have been added to all profiles.
- Effective Life on Hit is shown in the main information bubble.
- Fixed some goofy centering that was present in character search pages.
- Added a way for people to contact me.

===KNOWN ISSUES===

- The Diablo 3 API reports buffed (rather than unbuffed) values to this calculator.  Due to this, the calculator will have the best accuracy if you log your character off without the Enchantress armor buff. Once Blizzard fixes the D3 API, this will no longer be necessary.
- Life regen is not yet taken into account. Getting a value for this requires many requests to the Diablo 3 profile API, and applications are limited in the amount of requests they can make.
- Life steal (percentage-based) is not taken into account due the offensive stat calculation requirements for such a stat.

