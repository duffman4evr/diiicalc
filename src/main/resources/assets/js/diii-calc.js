// Helpers.
Number.prototype.numberFormat = function (decimals, dec_point, thousands_sep)
{
   dec_point = typeof dec_point !== 'undefined' ? dec_point : '.';
   thousands_sep = typeof thousands_sep !== 'undefined' ? thousands_sep : ',';

   var parts = this.toFixed(decimals).toString().split('.');
   parts[0] = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, thousands_sep);

   return parts.join(dec_point);
}

jQuery.fn.exists = function(){ return this.length > 0; }

// ----
// DIII Calc
// ----

DiiiCalcApp = new Backbone.Marionette.Application();

// Constants.
DiiiCalcApp.CHECK_MARK_CLASS = "glyphicon glyphicon-ok";

// Regions.
DiiiCalcApp.addRegions
({
   contentRegion: "#content"
});

// Models.
DiiiCalcApp.CareerProfileModel = Backbone.Model.extend( { urlRoot: '/service/career-profiles' } );
DiiiCalcApp.DefensiveSummaryModel = Backbone.Model.extend( { urlRoot: '/service/defensive-summaries' } );
DiiiCalcApp.OffensiveSummaryModel = Backbone.Model.extend( { urlRoot: '/service/offensive-summaries' } );
DiiiCalcApp.SkillChoiceSetModel = Backbone.Model.extend( { urlRoot: '/service/skill-choice-sets' } );

// Views / Layouts.
DiiiCalcApp.BattleTagLookupView = Backbone.Marionette.ItemView.extend
({
   template: '#battleTagSearchTemplate',
   events:
   {
      'submit .battle-tag-form' : 'findHeroes',
      'keyup  .battle-tag-input': 'clearErrors',
      'change .battle-tag-input': 'clearErrors',
      'input  .battle-tag-input': 'clearErrors',
      'paste  .battle-tag-input': 'clearErrors',
      'click  .select-realm'    : 'clickRealm'
   },
   onShow: function()
   {
      DiiiCalcApp.realm = "US";
      this.updateRealmUi();

      if (DiiiCalcApp.battleTag)
      {
         $('.battle-tag-input').val(DiiiCalcApp.battleTag);
      }
   },
   clickRealm: function(ev)
   {
      var clickTarget = $(ev.originalEvent.target);
      var dropdown = clickTarget.parent().parent().parent();
      DiiiCalcApp.realm = clickTarget.attr('id');
      dropdown.removeClass('open');
      this.updateRealmUi();
      return false;
   },
   updateRealmUi: function()
   {
      $('.select-realm').each
      (
         function(index, rawRalm)
         {
            var realm = $(rawRalm);
            var span = realm.find("span");

            if (realm.attr('id') === DiiiCalcApp.realm)
            {
               $('.dropdown-toggle').html(DiiiCalcApp.realm + '  <span class="caret"></span>');
               span.attr("class", DiiiCalcApp.CHECK_MARK_CLASS);
            }
            else
            {
               span.attr("class", "");
            }
         }
      )
   },
   findHeroes: function()
   {
      var battleTag = $('.battle-tag-input').val().replace('#', '-');
      var careerProfileModel = new DiiiCalcApp.CareerProfileModel({ id: battleTag });

      careerProfileModel.fetch
      ({
         data: $.param({ realm: DiiiCalcApp.realm }),
         success: function(model)
         {
            DiiiCalcApp.battleTag = battleTag;
            DiiiCalcApp.controller.showHeroesForModel(model);
            DiiiCalcApp.router.navigate('heroes/' + battleTag)
         },
         error: function()
         {
            DiiiCalcApp.battleTag = null;
            $('.help-block').html('BattleTag not found, try again.')
            $('.battle-tag-input').select();
         }
      });

      return false;
   },
   clearErrors: function()
   {
      $('.help-block').html('')
   }
});

DiiiCalcApp.HeroesView = Backbone.Marionette.ItemView.extend
({
   template: '#heroesTemplate',
   events:
   {
      'click .heroes-table' : 'clickHero'
   },
   clickHero: function(ev)
   {
      var clickTarget = $(ev.originalEvent.target);
      var heroId = clickTarget.parent().attr('id');
      DiiiCalcApp.controller.showBreakdownForHero(heroId);
      DiiiCalcApp.router.navigate('breakdowns/' + heroId);
   }
});

DiiiCalcApp.HeroLayout = Backbone.Marionette.Layout.extend
({
   template : function(model)
   {
      return _.template($("#heroTemplate").html(), model, { variable: 'hero' });
   },
   regions:
   {
      offensiveStatsRegion: '#offensive-stats',
      defensiveStatsRegion: '#defensive-stats',
      skillsRegion: '#skills'
   },
   onShow: function()
   {
      var that = this;

      // Show spinners while we load.
      var offensiveLoadingModel = new Backbone.Model({ heading: "Offensive", panelType: "warning" });
      var defensiveLoadingModel = new Backbone.Model({ heading: "Defensive", panelType: "info"});

      this.offensiveStatsRegion.show(new DiiiCalcApp.ProgressPanelView({ model: offensiveLoadingModel }));
      this.defensiveStatsRegion.show(new DiiiCalcApp.ProgressPanelView({ model: defensiveLoadingModel }));

      // Load data and show it in views.
      var offensiveModel = new DiiiCalcApp.OffensiveSummaryModel({ id: this.model.id });
      var defensiveModel = new DiiiCalcApp.DefensiveSummaryModel({ id: this.model.id });
      var skillChoiceSetModel = new DiiiCalcApp.SkillChoiceSetModel({ id: this.model.id });

      offensiveModel.fetch
      ({
         data: $.param({ battleTag: DiiiCalcApp.battleTag, realm: DiiiCalcApp.realm }),
         success: function(model)
         {
            DiiiCalcApp.onePrimaryStatDps = model.get("onePrimaryStatDps");
            DiiiCalcApp.onePercentCritChanceDps = model.get("onePercentCritChanceDps");
            DiiiCalcApp.onePercentCritDamageDps = model.get("onePercentCritDamageDps");
            DiiiCalcApp.onePercentAttackSpeedDps = model.get("onePercentAttackSpeedDps");

            var view = new DiiiCalcApp.OffensiveSummaryView({ model: model });

            that.offensiveStatsRegion.show(view);
         }
      });

      defensiveModel.fetch
      ({
         data: $.param({ battleTag: DiiiCalcApp.battleTag, realm: DiiiCalcApp.realm }),
         success: function(model)
         {
            DiiiCalcApp.oneArmorEhp = model.get("oneArmorEhp");
            DiiiCalcApp.oneAllResistEhp = model.get("oneAllResistEhp");
            DiiiCalcApp.oneVitalityEhp = model.get("oneVitalityEhp");
            DiiiCalcApp.onePercentLifeEhp = model.get("onePercentLifeEhp");

            var view = new DiiiCalcApp.DefensiveSummaryView({ model: model });

            that.defensiveStatsRegion.show(view);
         }
      });

      skillChoiceSetModel.fetch
      ({
         data: $.param({ battleTag: DiiiCalcApp.battleTag, realm: DiiiCalcApp.realm }),
         success: function(model)
         {
            DiiiCalcApp.passiveChoices = model.get('passiveChoices');
            DiiiCalcApp.activeChoices = model.get('activeChoices');

            var view = new DiiiCalcApp.SkillsView({ model: model });

            that.skillsRegion.show(view);
         }
      });
   }
});

DiiiCalcApp.ProgressPanelView = Backbone.Marionette.ItemView.extend
({
   template: '#progressPanelTemplate',
   onShow: function()
   {
      var opts = {
         lines: 13, // The number of lines to draw
         length: 8, // The length of each line
         width: 6, // The line thickness
         radius: 21, // The radius of the inner circle
         corners: 1, // Corner roundness (0..1)
         rotate: 0, // The rotation offset
         direction: 1, // 1: clockwise, -1: counterclockwise
         color: '#000', // #rgb or #rrggbb or array of colors
         speed: 1.3, // Rounds per second
         trail: 68, // Afterglow percentage
         shadow: false, // Whether to render a shadow
         hwaccel: true, // Whether to use hardware acceleration
         className: 'spinner', // The CSS class to assign to the spinner
         zIndex: 2e9, // The z-index (defaults to 2000000000)
         top: 'auto', // Top position relative to parent in px
         left: 'auto' // Left position relative to parent in px
      };

      var targets = document.getElementsByClassName("spinner-target");

      for (var i = 0; i < targets.length; i++)
      {
         opts.color = targets[i].style.color;
         new Spinner(opts).spin(targets[i]);
      }
   }
});

DiiiCalcApp.DefensiveSummaryView = Backbone.Marionette.ItemView.extend
({
   template: '#defensiveSummaryTemplate',
   events:
   {
      'submit .battle-tag-form' : 'findHeroes',
      'keyup  .armor-ehp-input': 'armorChanged',
      'change .armor-ehp-input': 'armorChanged',
      'input  .armor-ehp-input': 'armorChanged',
      'paste  .armor-ehp-input': 'armorChanged',
      'keyup  .resist-ehp-input': 'resistChanged',
      'change .resist-ehp-input': 'resistChanged',
      'input  .resist-ehp-input': 'resistChanged',
      'paste  .resist-ehp-input': 'resistChanged',
      'keyup  .percent-life-ehp-input': 'percentLifeChanged',
      'change .percent-life-ehp-input': 'percentLifeChanged',
      'input  .percent-life-ehp-input': 'percentLifeChanged',
      'paste  .percent-life-ehp-input': 'percentLifeChanged',
      'keyup  .vitality-ehp-input': 'vitalityChanged',
      'change .vitality-ehp-input': 'vitalityChanged',
      'input  .vitality-ehp-input': 'vitalityChanged',
      'paste  .vitality-ehp-input': 'vitalityChanged'
   },
   onShow: function()
   {
      $(".armor-ehp-input").val("200")
      $(".resist-ehp-input").val("80")
      $(".vitality-ehp-input").val("100");
      $(".percent-life-ehp-input").val("12")

      this.armorChanged();
      this.resistChanged();
      this.vitalityChanged();
      this.percentLifeChanged();
   },
   armorChanged: function()
   {
      var input = $(".armor-ehp-input").val();
      var ehpValue = DiiiCalcApp.oneArmorEhp * input;

      $(".armor-ehp-output").html(ehpValue.numberFormat(0) + " EHP");
   },
   resistChanged: function()
   {
      var input = $(".resist-ehp-input").val();
      var ehpValue = DiiiCalcApp.oneAllResistEhp * input;

      $(".resist-ehp-output").html(ehpValue.numberFormat(0) + " EHP");
   },
   vitalityChanged: function()
   {
      var input = $(".vitality-ehp-input").val();
      var ehpValue = DiiiCalcApp.oneVitalityEhp * input;

      $(".vitality-ehp-output").html(ehpValue.numberFormat(0) + " EHP");
   },
   percentLifeChanged: function()
   {
      var input = $(".percent-life-ehp-input").val();
      var ehpValue = DiiiCalcApp.onePercentLifeEhp * input;

      $(".percent-life-ehp-output").html(ehpValue.numberFormat(0) + " EHP");
   }
});

DiiiCalcApp.OffensiveSummaryView = Backbone.Marionette.ItemView.extend
({
   template: '#offensiveSummaryTemplate',
   events:
   {
      'submit .battle-tag-form' : 'findHeroes',
      'keyup  .primary-stat-dps-input': 'primaryStatChanged',
      'change .primary-stat-dps-input': 'primaryStatChanged',
      'input  .primary-stat-dps-input': 'primaryStatChanged',
      'paste  .primary-stat-dps-input': 'primaryStatChanged',
      'keyup  .crit-chance-dps-input': 'critChanceChanged',
      'change .crit-chance-dps-input': 'critChanceChanged',
      'input  .crit-chance-dps-input': 'critChanceChanged',
      'paste  .crit-chance-dps-input': 'critChanceChanged',
      'keyup  .crit-damage-dps-input': 'critDamageChanged',
      'change .crit-damage-dps-input': 'critDamageChanged',
      'input  .crit-damage-dps-input': 'critDamageChanged',
      'paste  .crit-damage-dps-input': 'critDamageChanged',
      'keyup  .attack-speed-dps-input': 'attackSpeedChanged',
      'change .attack-speed-dps-input': 'attackSpeedChanged',
      'input  .attack-speed-dps-input': 'attackSpeedChanged',
      'paste  .attack-speed-dps-input': 'attackSpeedChanged'
   },
   onShow: function()
   {
      $(".primary-stat-dps-input").val("100")
      $(".crit-chance-dps-input").val("5")
      $(".crit-damage-dps-input").val("50")
      $(".attack-speed-dps-input").val("9");

      this.primaryStatChanged();
      this.critChanceChanged();
      this.critDamageChanged();
      this.attackSpeedChanged();
   },
   primaryStatChanged: function()
   {
      var input = $(".primary-stat-dps-input").val();
      var dpsValue = DiiiCalcApp.onePrimaryStatDps * input;

      $(".primary-stat-dps-output").html(dpsValue.numberFormat(0) + " DPS");
   },
   critChanceChanged: function()
   {
      var input = $(".crit-chance-dps-input").val();
      var dpsValue = DiiiCalcApp.onePercentCritChanceDps * input;

      $(".crit-chance-dps-output").html(dpsValue.numberFormat(0) + " DPS");
   },
   critDamageChanged: function()
   {
      var input = $(".crit-damage-dps-input").val();
      var dpsValue = DiiiCalcApp.onePercentCritDamageDps * input;

      $(".crit-damage-dps-output").html(dpsValue.numberFormat(0) + " DPS");
   },
   attackSpeedChanged: function()
   {
      var input = $(".attack-speed-dps-input").val();
      var dpsValue = DiiiCalcApp.onePercentAttackSpeedDps * input;

      $(".attack-speed-dps-output").html(dpsValue.numberFormat(0) + " DPS");
   }
});

DiiiCalcApp.SkillsView = Backbone.Marionette.ItemView.extend
({
   template: '#skillsTemplate',
   events:
   {
      'click .skill-select' : 'skillSelect'
   },
   onShow: function()
   {
      this.updateCheckMarks()
   },
   skillSelect: function(ev)
   {
      var clickTarget = $(ev.originalEvent.target);
      var disable = clickTarget.hasClass("disable-select")
      var dropdown = clickTarget.parent().parent().parent();
      var slug = dropdown.attr("id");
      var isActiveSkill = dropdown.hasClass("active-skill");

      if (isActiveSkill)
      {
         if (disable)
         {
            delete DiiiCalcApp.activeChoices[slug];
         }
         else
         {
            DiiiCalcApp.activeChoices[slug] = clickTarget.attr("id");
         }
      }
      else
      {
         if (disable)
         {
            delete DiiiCalcApp.passiveChoices[slug];
         }
         else
         {
            DiiiCalcApp.passiveChoices[slug] = null;
         }
      }

      dropdown.removeClass('open');
      this.updateCheckMarks();
      return false;
   },
   updateCheckMarks: function()
   {
      $(".active-skill, .passive-skill").each
      (
         function()
         {
            var dropdown = $(this);
            var slug = dropdown.attr("id");
            var activeSelected = DiiiCalcApp.activeChoices.hasOwnProperty(slug);
            var passiveSelected = DiiiCalcApp.passiveChoices.hasOwnProperty(slug);
            var selected = activeSelected || passiveSelected;
            var rune = null;

            if (selected)
            {
               if (activeSelected)
               {
                  rune = DiiiCalcApp.activeChoices[slug];
               }
               if (passiveSelected)
               {
                  rune = null;
               }
            }

            // Clear all existing checkmarks, make it opaque.
            dropdown.find("li").find("span").attr("class", "");

            if (!selected)
            {
               dropdown.find("img").css("opacity", "0.4");
               dropdown.find(".disable-select").find("span").attr("class", DiiiCalcApp.CHECK_MARK_CLASS);
            }
            else
            {
               dropdown.find("img").css("opacity", "1.0");

               if (rune)
               {
                  dropdown.find("#" + rune).find("span").attr("class", DiiiCalcApp.CHECK_MARK_CLASS);
               }
               else
               {
                  var runelessEnable = dropdown.find(".enable-select");

                  if (runelessEnable.exists())
                  {
                     dropdown.find(".enable-select").find("span").attr("class", DiiiCalcApp.CHECK_MARK_CLASS);
                  }
                  else
                  {
                     dropdown.find("img").css("opacity", "0.4");
                     dropdown.find(".disable-select").find("span").attr("class", DiiiCalcApp.CHECK_MARK_CLASS);
                  }
               }
            }
         }
      );
   }
});

// Controller.
DiiiCalcApp.Controller = Marionette.Controller.extend
({
   showBreakdownForHero: function(heroId)
   {
      var heroModel = new Backbone.Model(DiiiCalcApp.heroMap[heroId]);

      DiiiCalcApp.contentRegion.show(new DiiiCalcApp.HeroLayout({ model: heroModel }));
   },
   showHeroesForBattleTag: function(battleTag)
   {
      var careerProfileModel = new DiiiCalcApp.CareerProfileModel({ id: battleTag });
      careerProfileModel.fetch
      ({
         data: $.param({ realm: DiiiCalcApp.realm }),
         success: function(model)
         {
            DiiiCalcApp.controller.showHeroesForModel(model);
         }
      });
   },
   showHeroesForModel: function(model)
   {
      DiiiCalcApp.heroMap = {};

      model.get("heroes").forEach
      (
         function(hero)
         {
            DiiiCalcApp.heroMap[hero.id] = hero;
         }
      );

      var heroesView = new DiiiCalcApp.HeroesView({ model: model });

      DiiiCalcApp.contentRegion.show(heroesView);
   },
   showBattleTagPrompt: function()
   {
      DiiiCalcApp.contentRegion.show(new DiiiCalcApp.BattleTagLookupView())
   }
});

// Initializer.
DiiiCalcApp.addInitializer
(
   function()
   {
      DiiiCalcApp.controller = new DiiiCalcApp.Controller();
      DiiiCalcApp.router = new Marionette.AppRouter
      ({
         controller: DiiiCalcApp.controller,
         appRoutes:
         {
            "": "showBattleTagPrompt",
            "heroes/:id": "showHeroesForBattleTag",
            "breakdowns/:id": "showBreakdownForHero"
         }
      });
   }
);

DiiiCalcApp.start();

Backbone.history.start();