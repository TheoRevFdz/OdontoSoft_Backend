(function () {
	angular.module('tratment.controller', [])
		.controller('TreatmentController', ['$scope', '$http', function ($scope, $http) {
			$scope.treatments = {};
		}]);
})();