(function () {
	angular.module('odonto.directives', [])
		.directive('menuBar', function () {
			return {
				restrict: 'E',
				templateUrl: 'components/menu-bar.html'
			};
		});

	angular.module('patient.directives', [])
		.directive('menuBar', function () {
			return {
				restrict: 'E',
				templateUrl: 'components/menu-bar.html'
			};
		})
		.directive('formPatient', function () {
			return {
				restrict: 'E',
				templateUrl: 'components/patient/form-patient.html'
			};
		});

	angular.module('control.directives', [])
		.directive('menuBar', function () {
			return {
				restrict: 'E',
				templateUrl: 'components/menu-bar2.html'
			};
		})
		.directive('formControl', function () {
			return {
				restrict: 'E',
				templateUrl: 'components/control/form-control.html'
			}
		});
})();