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
		});
})();