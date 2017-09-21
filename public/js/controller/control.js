(function () {
	angular.module('control.controller', [])
		.controller('ControlController', ['$scope', '$http', function ($scope, $http) {
			$scope.controls = {};

			findAllControls($scope, $http);

			$(".button-collapse").sideNav();
		}]);

	function findAllControls($scope, $http) {
		$http({
			method: 'get',
			url: 'api/controls/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				console.log(response.data);
				$scope.controls = response.data;
			},
			function error(response) {
				alert(response);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}
})();