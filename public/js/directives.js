(function () {
	angular.module('odonto.directives', [])
		.directive('menuBar', function () {
			return {
				restrict: 'E',
				templateUrl: 'components/menu-bar2.html'
			};
		});

	angular.module('patient.directives', [])
		.directive('menuBar', function () {
			return {
				restrict: 'E',
				templateUrl: 'components/menu-bar2.html'
			};
		})
		.directive('formPatient', function () {
			return {
				restrict: 'E',
				templateUrl: 'components/patient/form-patient.html'
			};
		});
})();