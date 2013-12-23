DiiiCalcApp = new Backbone.Marionette.Application();

// Regions.
DiiiCalcApp.addRegions
({
   contentRegion: "#content"
});

// Models.
DiiiCalcApp.CareerProfileModel = Backbone.Model.extend( { urlRoot: '/service/career-profiles' } );
DiiiCalcApp.HeroModel = Backbone.Model.extend( { urlRoot: '/service/heroes' } );

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
      'paste  .battle-tag-input': 'clearErrors'
   },
   onShow: function()
   {
      if (DiiiCalcApp.battleTag)
      {
         $('.battle-tag-input').val(DiiiCalcApp.battleTag);
      }
   },
   findHeroes: function()
   {
      var battleTag = $('.battle-tag-input').val().replace('#', '-');
      var careerProfileModel = new DiiiCalcApp.CareerProfileModel({ id: battleTag });

      careerProfileModel.fetch
      ({
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
      alert("hello" + heroId);
   }
});

DiiiCalcApp.HeroLayout = Backbone.Marionette.Layout.extend
({
   template: '#heroTemplate',
   regions:
   {
      offensiveStatsRegion: '#offensive-stats',
      defensiveStatsRegion: '#defensive-stats'
   },
   onShow: function()
   {
      if (this.model.get('customer'))
      {
         var that = this;
         var connectors = new ProxyClients();
         connectors.fetch
         ({
            data: $.param({ customerId: this.model.get('customer').get('id') }),
            success: function(response)
            {
               var tempModel = new Backbone.Model
               ({
                  customerId: that.model.get('customer').get('id'),
                  proxyClients: response.models,
                  username: username,
                  password: password
               });

               that.proxyClientsRegion.show(new ProxyClientsView({ model: tempModel }));
            }
         });

         var connectorsWithData = new ProxyClients();
         connectorsWithData.fetch
         ({
            data: $.param({ customerId: this.model.get('customer').get('id'), requireData: 'true' }),
            success: function(response)
            {
               that.renderReports(that, response.models);
            }
         });
      }
   },
});

// Controller.
DiiiCalcApp.Controller = Marionette.Controller.extend
({
   showHeroForId: function(heroId)
   {
      var heroModel = new DiiiCalcApp.HeroModel({ id: heroId });
      heroModel.fetch
      ({
         data: $.param({ battleTag: DiiiCalcApp.battleTag }),
         success: function(model)
         {
            var heroView = new DiiiCalcApp.HeroView({ model: model });
            DiiiCalcApp.contentRegion.show(heroView);
         }
      });
   },
   showHeroesForBattleTag: function(battleTag)
   {
      var careerProfileModel = new DiiiCalcApp.CareerProfileModel({ id: battleTag });
      careerProfileModel.fetch
      ({
         success: function(model)
         {
            DiiiCalcApp.controller.showHeroesForModel(model);
         }
      });
   },
   showHeroesForModel: function(model)
   {
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
            "heroes/:id": "showHeroesForBattleTag"
         }
      });
   }
);

DiiiCalcApp.start();

Backbone.history.start();

////////

/*
$.fn.serializeObject = function ()
{
   var o = {};
   var a = this.serializeArray();
   $.each(a, function ()
   {
      if (o[this.name] !== undefined)
      {
         if (!o[this.name].push)
         {
            o[this.name] = [o[this.name]];
         }
         o[this.name].push(this.value || '');
      } else
      {
         o[this.name] = this.value || '';
      }
   });
   return o;
};

jQuery.fn.exists = function() { return this.length > 0; }

var username = null;
var password = null;
var userRole = null;
var firstName = null;

$.ajaxSetup
({
   'beforeSend': function(xhr)
   {
      if (username && password)
      {
         xhr.setRequestHeader('Authorization', "Basic " + btoa(username + ":" + password));
      }
   }
});

function isUserLoggedIn()
{
   return (username != null);
}

var ProxyConfig = Backbone.Model.extend( { urlRoot: '/service/admin/connector-configs' } );
var Customers = Backbone.Collection.extend( { url: '/service/admin/customers' } );
var Customer = Backbone.Model.extend( { urlRoot: '/service/admin/customers' } );
var ProxyClients = Backbone.Collection.extend( { url: '/service/admin/connectors' } );
var ProxyClient = Backbone.Model.extend( { urlRoot: '/service/admin/connectors' } );
var DeviceUsageReport = Backbone.Model.extend( { urlRoot: '/service/admin/telemetry/reports/device-usages' } );
var UsageReport = Backbone.Model.extend( { urlRoot: '/service/admin/telemetry/reports/usage' } );
var BundleSnapshot = Backbone.Model.extend( { urlRoot: '/service/admin/telemetry/snapshots' } );

var HeaderView = Backbone.Marionette.ItemView.extend
({
   template: '#headerViewTemplate',
   events:
   {
      'click .navbar-btn': 'signOut'
   },
   signOut: function()
   {
      username = null;
      password = null;
      userRole = null;
      firstName = null;

      diConnectRouter.navigate('login', { trigger: true } );
   }
});

var LoginView = Backbone.Marionette.ItemView.extend
({
   isLoginView: true,
   template: '#loginViewTemplate',
   events:
   {
      'submit .form-signin': 'signIn',
      'click .update-pass-btn': 'updatePass',
      'keyup #newPass2': 'validatePasswords',
      'change #newPass2': 'validatePasswords',
      'input #newPass2': 'validatePasswords',
      'paste #newPass2': 'validatePasswords',
      'keyup #bad-creds-modal': 'failModalKeyPress',
      'keyup #pass-reset-modal': 'updatePassModalKeyPress'
   },
   signIn: function()
   {
      var signInDetails = $(".form-signin").serializeObject();

      username = signInDetails.username;
      password = signInDetails.password;
      userRole = null;
      firstName = null;

      $.ajax
      ({
         'url'      : '/service/admin/login',
         'type'     : 'POST',
         'success': function(response)
         {
            userRole = response['role-id'];
            firstName = response['first-name'];

            diConnectRouter.navigate('customers', { trigger: true } );
         },
         'error': function(response)
         {
            var passReset = response.getResponseHeader("x-password-reset");

            if (passReset)
            {
               $('#pass-reset-modal').modal('show').on
               (
                  'shown.bs.modal',
                  function()
                  {
                     $('#newPass1').focus();
                  }
               );
            }
            else
            {
               username = undefined;
               password = undefined;

               $('#bad-creds-modal').modal('show');
            }
         }
      });

      return false;
   },
   validatePasswords: function()
   {
      if (!this.events.hasOwnProperty('keyup #newPass1'))
      {
         this.events['keyup #newPass1'] = 'validatePasswords';
         this.events['change #newPass1'] = 'validatePasswords';
         this.events['input #newPass1'] = 'validatePasswords';
         this.events['paste #newPass1'] = 'validatePasswords';

         this.delegateEvents();
      }

      var pass1 = $('#newPass1');
      var pass2 = $('#newPass2');

      var valid = true;

      if (pass1.val() !== pass2.val())
      {
         $('.pass-form-group').addClass("has-error");
         $('.help-block-pass').html("Passwords do not match.");
         valid = false;
      }
      else
      {
         $('.pass-form-group').removeClass("has-error");
         $('.help-block-pass').html("&nbsp;");
      }

      if (valid)
      {
         $('.update-pass-btn').removeAttr("disabled");
      }
      else
      {
         $('.update-pass-btn').attr("disabled", "disabled");
      }
   },
   updatePass: function()
   {
      var newPass = $('#newPass1').val();

      $.ajax
      ({
         'url'      : '/service/admin/login/reset-password',
         'type'     : 'PUT',
         'headers' :
         {
            'x-new-password': newPass
         },
         'success': function(response)
         {
            password = newPass;
            userRole = response['role-id'];
            firstName = response['first-name'];

            $('#pass-reset-modal').modal('hide').on
            (
               'hidden.bs.modal',
               function()
               {
                  diConnectRouter.navigate('customers', { trigger: true } );
               }
            );
         },
         'error': function()
         {
            alert("Unrecoverable error.")
         }
      });
   },
   failModalKeyPress: function(e)
   {
      if (e.which === 13)
      {
         $('#bad-creds-modal').modal('hide').on
         (
            'hidden.bs.modal',
            function()
            {
               $('#pw-input').select();
            }
         );
      }
   },
   updatePassModalKeyPress: function(e)
   {
      if (e.which === 13)
      {
         this.updatePass();
      }
   }
});

var CustomersView = Backbone.Marionette.ItemView.extend({ template: '#customersViewTemplate' });
var ProxyClientsView = Backbone.Marionette.ItemView.extend({ template: '#proxyClientsViewTemplate' });

CustomerLayout = Backbone.Marionette.Layout.extend
({
   template: "#customerLayoutTemplate",
   regions:
   {
      reportRegion: "#customer-report",
      proxyClientsRegion: "#proxy-client-list"
   },
   onShow: function()
   {
      if (this.model.get('customer'))
      {
         var that = this;
         var connectors = new ProxyClients();
         connectors.fetch
         ({
            data: $.param({ customerId: this.model.get('customer').get('id') }),
            success: function(response)
            {
               var tempModel = new Backbone.Model
               ({
                  customerId: that.model.get('customer').get('id'),
                  proxyClients: response.models,
                  username: username,
                  password: password
               });

               that.proxyClientsRegion.show(new ProxyClientsView({ model: tempModel }));
            }
         });

         var connectorsWithData = new ProxyClients();
         connectorsWithData.fetch
         ({
            data: $.param({ customerId: this.model.get('customer').get('id'), requireData: 'true' }),
            success: function(response)
            {
               that.renderReports(that, response.models);
            }
         });
      }
   },
   renderReports: function(that, models)
   {
      if (models.length === 0)
      {
         return;
      }

      var proxyClientId = models[0].get('proxy-client-id');
      var bundleSnapshot = new BundleSnapshot({ id: proxyClientId });

      bundleSnapshot.fetch
      ({
         success: function(model)
         {
            var tempModel = new Backbone.Model({ bundleSnapshot: model, connectorModels: models });
            var customerReportView = new CustomerReportView({ model: tempModel });
            that.reportRegion.show(customerReportView);
         }
      });
   },
   events:
   {
      'submit .edit-customer-form': 'save',
      'click .delete-customer': 'showDeleteModal',
      'click .delete-modal': 'erase',
      'keyup #name': 'nameChange',
      'change #name': 'nameChange',
      'input #name': 'nameChange',
      'paste #name': 'nameChange'
   },
   save: function(ev)
   {
      ev.preventDefault();
      var jsonObject = $(".edit-customer-form").serializeObject();
      var customer = new Customer();
      customer.save
      (
         jsonObject,
         {
            success: function()
            {
               if ($("#id").exists())
               {
                  $('.help-block').html("Saved!");
               }
               else
               {
                  diConnectRouter.navigate('', { trigger: true } );
               }
            },
            error: function(model, xhr)
            {
               $('.name-form-group').addClass("has-error");
               $('.help-block').html(xhr.responseText);
               $('.update-customer').attr("disabled", "disabled");
            }
         }
      );
      return false;
   },
   showDeleteModal: function()
   {
      $('#delete-modal').modal('show');
   },
   erase: function()
   {
      var that = this;
      $('#delete-modal').modal('hide').on
      (
         'hidden.bs.modal',
         function()
         {
            that.model.get('customer').destroy
            ({
               success: function()
               {
                  diConnectRouter.navigate('', { trigger: true } );
               }
            });
         }
      );
   },
   nameChange: function(ev)
   {
      if (ev.which === 13)
      {
         return;
      }

      var name = $(ev.currentTarget);

      var valid = true;

      if (name.val().trim() == "")
      {
         $('.name-form-group').addClass("has-error");
         $('.help-block').html("Name must have content.");
         valid = false;
      }
      else
      {
         $('.name-form-group').removeClass("has-error");
         $('.help-block').html("&nbsp;")
      }

      if (valid)
      {
         $('.update-customer').removeAttr("disabled");
      }
      else
      {
         $('.update-customer').attr("disabled", "disabled");
      }
   }
});

var CustomerReportView = Backbone.Marionette.ItemView.extend
({
   template: '#customerReportViewTemplate',
   onShow: function()
   {
      setTimeout(this.populateLineGraph, 50, this, true);
      setTimeout(this.populatePieChart, 50, this, true);
   },
   events:
   {
      'click .usage-timespan-day': 'lineStartDay',
      'click .usage-timespan-week': 'lineStartWeek',
      'click .device-timespan-day': 'pieStartDay',
      'click .device-timespan-week': 'pieStartWeek',
      'change .connector-report-select': 'changeReport'
   },
   changeReport: function()
   {
      var that = this;
      var bundleSnapshot = new BundleSnapshot({ id: $('.connector-report-select :selected').val() });

      bundleSnapshot.fetch
      ({
         success: function(model)
         {
            that.model = new Backbone.Model({ bundleSnapshot: model, connectorModels: that.model.get('connectorModels') });
            that.render();
            that.populateLineGraph(that, false);
            that.populatePieChart(that, false)
         }
      });
   },
   lineStartDay: function()
   {
      $('.usage-timespan-week').removeClass("active");
      $('.usage-timespan-day').addClass("active");
      this.populateLineGraph(this, false);
   },
   lineStartWeek: function()
   {
      $('.usage-timespan-week').addClass("active");
      $('.usage-timespan-day').removeClass("active");
      this.populateLineGraph(this, false);
   },
   pieStartDay: function()
   {
      $('.device-timespan-week').removeClass("active");
      $('.device-timespan-day').addClass("active");
      this.populatePieChart(this, false);
   },
   pieStartWeek: function()
   {
      $('.device-timespan-week').addClass("active");
      $('.device-timespan-day').removeClass("active");
      this.populatePieChart(this, false);
   },
   populatePieChart: function(that, firstDraw)
   {
      var startTime = 0;
      if ($('.device-timespan-week').hasClass("active"))
      {
         startTime = Date.now() - (7 * 24 * 60 * 60 * 1000);
      }
      else
      {
         startTime = Date.now() - (24 * 60 * 60 * 1000);
      }
      that.deviceUsageReport = new DeviceUsageReport({ id: that.model.get('bundleSnapshot').id });
      that.deviceUsageReport.fetch
      ({
         data: $.param({ start: startTime }),
         success: function(data)
         {
            var pieData = [];

            for (var i = 0; i < data.keys().length; i++)
            {
               var key = data.keys()[i];
               if (key == "id")
               {
                  continue;
               }
               pieData.push({ key: data.get(key)['device-type'], y: data.get(key)['usage-count'] })
            }

            $('#deviceTypesChart svg').empty();

            nv.addGraph(function()
            {
               var width = 250;
               var height = 250;

               var chart = nv.models.pieChart()
                  .x(function(d) { return d.key })
                  .y(function(d) { return d.y })
                  .color(d3.scale.category10().range())
                  .width(width)
                  .height(height);

               d3.select("#deviceTypesChart svg")
                  .datum(pieData)
                  .transition().duration(1200)
                  .attr('width', width)
                  .attr('height', height)
                  .call(chart);

               return chart;
            });
         }
      });
   },
   populateLineGraph: function(that, firstDraw)
   {
      var startTime = 0;
      if ($('.usage-timespan-week').hasClass("active"))
      {
         startTime = Date.now() - (7 * 24 * 60 * 60 * 1000);
      }
      else
      {
         startTime = Date.now() - (24 * 60 * 60 * 1000);
      }
      that.usageReport = new UsageReport({ id: that.model.get('bundleSnapshot').id });
      that.usageReport.fetch
      ({
         data: $.param({ start: startTime }),
         success: function(data)
         {
            var lineData = [];

            for (var i = 0; i < data.keys().length; i++)
            {
               var key = data.keys()[i];
               if (key == "id")
               {
                  continue;
               }
               lineData.push({ x: data.get(key)['timestamp'], y: data.get(key)['usage-count'] })
            }

            var graphData = [{ values: lineData, key: 'Usage', color: '#ff7f0e' }];

            // We must clear the dom here or else the re-renders will fail.
            $('#usageChart svg').empty();

            nv.addGraph(function()
            {
               var chart = nv.models.lineChart();

               chart.xAxis
                  .axisLabel('')
                  .tickFormat(function (d) { return ''; });

               chart.yAxis
                  .axisLabel('Usage Count')
                  .tickFormat(d3.format(',r'));

               d3.select('#usageChart svg')
                  .datum(graphData)
                  .transition().duration(500)
                  .call(chart);

               return chart;
            });
         }
      });
   }
});

var ProxyClientLayout = Backbone.Marionette.Layout.extend
({
   template: "#proxyClientLayoutTemplate",
   regions:
   {
      configRegion: "#proxy-config"
   },
   onShow: function()
   {
      var that = this;

      var proxyClientModel = this.model.get('proxyClient');

      if (proxyClientModel)
      {
         var proxyClientId = this.model.get('proxyClient').get('proxy-client-id');
         this.proxyConfig = new ProxyConfig({ id: proxyClientId });

         this.proxyConfig.fetch
         ({
            success: function(model)
            {
               that.configView = new ProxyConfigView({ model: model });
               that.configRegion.show(that.configView);
            }
         });
      }
   },
   events:
   {
      'submit .edit-proxy-client-form'   : 'save',
      'click .delete-proxy-client'       : 'showDeleteModal',
      'click .delete-modal'              : 'erase',
      'change #default-writeable-dir'    : 'validateForm',
      'change #default-heartbeat-period' : 'validateForm',
      'keyup #type'                      : 'validateForm',
      'change #type'                     : 'validateForm',
      'input #type'                      : 'validateForm',
      'paste #type'                      : 'validateForm',
      'keyup #heartbeat-period'          : 'validateForm',
      'change #heartbeat-period'         : 'validateForm',
      'input #heartbeat-period'          : 'validateForm',
      'paste #heartbeat-period'          : 'validateForm',
      'keyup #writeable-dir'             : 'validateForm',
      'change #writeable-dir'            : 'validateForm',
      'input #writeable-dir'             : 'validateForm',
      'paste #writeable-dir'             : 'validateForm'
   },
   save: function(ev)
   {
      if (this.configView)
      {
         this.configView.save();
      }

      var jsonObject = $(ev.currentTarget).serializeObject();

      if (jsonObject['regen-api-key'] == 'on')
      {
         jsonObject['api-key'] = null;
      }
      if (jsonObject['regen-secret-key'] == 'on')
      {
         jsonObject['secret-key'] = null;
      }

      var proxyClient = new ProxyClient();

      proxyClient.save
      (
         jsonObject,
         {
            success: function()
            {
               diConnectRouter.navigate('customers/edit/' + jsonObject['customer-id'], { trigger: true } );
            }
         }
      );

      return false;
   },
   showDeleteModal: function()
   {
      $('#delete-modal').modal('show');
   },
   erase: function()
   {
      var that = this;
      $('#delete-modal').modal('hide').on
      (
         'hidden.bs.modal',
         function()
         {
            that.model.get('proxyClient').destroy
            ({
               success: function()
               {
                  diConnectRouter.navigate('customers/edit/' + $("#customer-id").val(), { trigger: true } );
               }
            });
         }
      );
   },
   validateForm: function()
   {
      var type = $('#type');
      var heartbeat = $('#heartbeat-period');

      var writeableDir = $('#writeable-dir');
      var writeableDirCheckbox = $('#default-writeable-dir');

      var valid = true;

      if (type.val().trim() == "")
      {
         $('.type-form-group').addClass("has-error");
         $('.help-block-type').html("Type must have content.");
         valid = false;
      }
      else
      {
         $('.type-form-group').removeClass("has-error");
         $('.help-block-type').html("&nbsp;")
      }

      if (heartbeat.exists())
      {
         var heartbeatCheckbox = $('#default-heartbeat-period');

         if (heartbeat.exists() && isNaN(heartbeat.val()) && !heartbeatCheckbox.is(":checked"))
         {
            $('.heartbeat-form-group').addClass("has-error");
            $('.help-block-heartbeat').html("Heartbeat must be a number.");
            valid = false;
         }
         else
         {
            $('.heartbeat-form-group').removeClass("has-error");
            $('.help-block-heartbeat').html("&nbsp;")
         }

         if (heartbeatCheckbox.is(":checked"))
         {
            heartbeat.attr("disabled", "disabled");
         }
         else
         {
            heartbeat.removeAttr("disabled");
         }
      }

      if (writeableDir.exists())
      {
         if (writeableDir.val().trim() == "" && !writeableDirCheckbox.is(":checked"))
         {
            $('.writeable-dir-form-group').addClass("has-error");
            $('.help-block-writeable-dir').html("Writeable dir must have content.");
            valid = false;
         }
         else
         {
            $('.writeable-dir-form-group').removeClass("has-error");
            $('.help-block-writeable-dir').html("&nbsp;")
         }

         if (writeableDirCheckbox.is(":checked"))
         {
            writeableDir.attr("disabled", "disabled");
         }
         else
         {
            writeableDir.removeAttr("disabled");
         }
      }

      if (valid)
      {
         $('.update-proxy-client').removeAttr("disabled");
      }
      else
      {
         $('.update-proxy-client').attr("disabled", "disabled");
      }

      $('.download-config').attr("disabled", "disabled");
   }
});

var ProxyConfigView = Backbone.Marionette.ItemView.extend
({
   template : function(serializedModel)
   {
      return _.template($('#proxyConfigViewTemplate').html(), serializedModel, {variable: 'proxyConfig'});
   },
   save: function()
   {
      var configDetails = { };
      if (!$('#default-writeable-dir').is(':checked'))
      {
         configDetails['writeable-dir'] = $('#writeable-dir').val();
      }
      if (!$('#default-heartbeat-period').is(':checked'))
      {
         configDetails['heartbeat-period'] = $('#heartbeat-period').val();
      }
      this.model.save(configDetails);
   }
});

function showHeaderView()
{
   var tempModel = new Backbone.Model({ firstName: firstName });
   loginRegion.show(new HeaderView({ model: tempModel }));
}

var mainRegion = new Backbone.Marionette.Region({ el: "#content" });
var loginRegion = new Backbone.Marionette.Region({ el: "#header" });

var DiConnectController = Marionette.Controller.extend
({
   login: function()
   {
      showHeaderView();

      mainRegion.show(new LoginView());
   },
   customers: function()
   {
      if (!isUserLoggedIn()) { this.reRouteToLogin(); return; }

      showHeaderView();

      var customers = new Customers();

      customers.fetch
      ({
         success: function(response)
         {
            var tempModel = new Backbone.Model({ customers: response.models });
            mainRegion.show(new CustomersView({ model: tempModel}));
         }
      });
   },
   customerNew: function()
   {
      if (!isUserLoggedIn()) { this.reRouteToLogin(); return; }

      showHeaderView();
      var tempModel = new Backbone.Model({ customer: null, userRole: userRole });
      mainRegion.show(new CustomerLayout({ model: tempModel }));
   },
   customerEdit: function(term)
   {
      if (!isUserLoggedIn()) { this.reRouteToLogin(); return; }

      showHeaderView();
      var customer = new Customer({id: term});
      customer.fetch
      ({
         success: function(model)
         {
            var tempModel = new Backbone.Model({ customer: model, userRole: userRole });
            mainRegion.show(new CustomerLayout({ model: tempModel }));
         }
      })
   },
   customerReport: function(term)
   {
      if (!isUserLoggedIn()) { this.reRouteToLogin(); return; }

      showHeaderView();
      mainRegion.show(new CustomerReportView({id: term}));
   },
   proxyClientNew: function(term)
   {
      if (!isUserLoggedIn()) { this.reRouteToLogin(); return; }

      showHeaderView();

      var tempModel = new Backbone.Model
      ({
         proxyClient: null,
         customer: null,
         customerId: term
      });

      mainRegion.show(new ProxyClientLayout({ model: tempModel }));
   },
   proxyClientEdit: function(term)
   {
      if (!isUserLoggedIn()) { this.reRouteToLogin(); return; }

      showHeaderView();

      var proxyClient = new ProxyClient({ id: term });

      proxyClient.fetch
      ({
         success: function(proxyClientModel)
         {
            var customer = new Customer({ id: proxyClientModel.get('customer-id') });

            customer.fetch
            ({
               success: function(customerModel)
               {
                  var tempModel = new Backbone.Model
                  ({
                     customer: customerModel,
                     proxyClient: proxyClientModel,
                     username: username,
                     password: password,
                     userRole: userRole
                  });

                  mainRegion.show(new ProxyClientLayout({model: tempModel}));
               }
            });
         }
      })
   },

   reRouteToLogin: function()
   {
      diConnectRouter.navigate('login', { trigger: true} );
   }
});

var DiConnectRouter = Backbone.Marionette.AppRouter.extend
({
   appRoutes:
   {
      ''                              : 'customers',
      'login'                         : 'login',
      'customers'                     : 'customers',
      'customers/new'                 : 'customerNew',
      'customers/edit/:id'            : 'customerEdit',
      'customers/report/:id'          : 'customerReport',
      'proxy-clients/new/:customerId' : 'proxyClientNew',
      'proxy-clients/edit/:id'        : 'proxyClientEdit'
   }
});

var diConnectRouter = new DiConnectRouter({ controller: new DiConnectController() })

Backbone.history.start();
   */