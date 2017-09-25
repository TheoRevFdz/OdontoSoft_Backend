(function () {
	angular.module('control.controller', [])
		.controller('ControlController', ['$scope', '$http', function ($scope, $http) {
			$scope.controls = {};
			$scope.works = {};
			$scope.accion = "";
			$scope.work = "";

			findAllControls($scope, $http);
			getAllWorks($scope, $http);

			$scope.mostrar = function (w) {
				alert(w.ID + " - " + w.name);
			}

			$scope.newControl = function () {
				$scope.accion = "NUEVO";
			};

			$scope.modifyControl = function () {
				$scope.accion = "MODIFICAR";
			};
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
				alert(response.message);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}

	function getAllWorks($scope, $http) {
		$http({
			method: 'GET',
			url: 'api/works/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				console.log(response.data);
				$scope.works = response.data;
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