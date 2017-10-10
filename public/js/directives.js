(function() {
  angular.module("odonto.directives", []).directive("menuBar", function() {
    return menuBar();
  });

  angular
    .module("patient.directives", [])
    .directive("menuBar", function() {
      return menuBar();
    })
    .directive("formPatient", function() {
      return {
        restrict: "E",
        templateUrl: "components/patient/form-patient.html"
      };
    });

  angular
    .module("control.directives", [])
    .directive("menuBar", function() {
      return menuBar();
    })
    .directive("formControl", function() {
      return {
        restrict: "E",
        templateUrl: "components/control/form-control.html"
      };
    })
    .directive("selectBrowser", function($timeout, $parse) {
      return selectBrowser($timeout, $parse);
    });

  angular
    .module("work.directives", [])
    .directive("menuBar", function() {
      return menuBar();
    })
    .directive("formWork", function() {
      return {
        restrict: "E",
        templateUrl: "components/work/form-work.html"
      };
    });

  angular
    .module("treatment.directives", [])
    .directive("menuBar", function() {
      return menuBar();
    })
    .directive("formTreatment", function() {
      return {
        restrict: "E",
        templateUrl: "components/treatment/form-treatment.html"
      };
    })
    .directive("selectBrowser", function($timeout, $parse) {
      return selectBrowser($timeout, $parse);
    })
    .directive("formCuracion", function() {
      return {
        restrict: "E",
        templateUrl: "components/treatment/form-curacion.html"
      };
    })
    .directive("formComfirmDeleteCuracion", function() {
      return {
        restrict: "E",
        templateUrl: "components/treatment/form-comfirm_delete_curacion.html"
      };
    });

  function menuBar() {
    return {
      restrict: "E",
      templateUrl: "components/menu-bar.html"
    };
  }

  function selectBrowser($timeout, $parse) {
    return {
      restrict: "AC",
      require: "ngModel",
      link: function(scope, element, attrs) {
        $timeout(function() {
          element.select2();
          element.select2Initialized = true;
        });

        var refreshSelect = function() {
          if (!element.select2Initialized) return;
          $timeout(function() {
            element.trigger("change");
          });
        };

        var recreateSelect = function() {
          if (!element.select2Initialized) return;
          $timeout(function() {
            element.select2("destroy");
            element.select2();
          });
        };

        scope.$watch(attrs.ngModel, refreshSelect);

        if (attrs.ngOptions) {
          var list = attrs.ngOptions.match(/ in ([^ ]*)/)[1];
          // watch for option list change
          scope.$watch(list, recreateSelect);
        }

        if (attrs.ngDisabled) {
          scope.$watch(attrs.ngDisabled, refreshSelect);
        }
      }
    };
  }
})();
