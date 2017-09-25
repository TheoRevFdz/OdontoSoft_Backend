(function () {
	angular.module('treatment.controller', [])
		.controller('TreatmentController', ['$scope', '$http', function ($scope, $http) {
			$scope.treatments = {};

			findAllTreatments($scope, $http);
		}]);

	function findAllTreatments($scope, $http) {
		$http({
			method: 'GET',
			url: 'api/treatments/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				console.log(response.data);
				$scope.treatments = response.data;
			},
			function error(response) {
				alert(response.message);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}
})();