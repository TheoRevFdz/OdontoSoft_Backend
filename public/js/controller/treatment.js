(function () {
	angular.module('treatment.controller', [])
		.controller('TreatmentController', ['$scope', '$http', function ($scope, $http) {
			$scope.accion = "";
			$scope.treatments = {};
			$scope.patients = {};
			$scope.treatment = {
				patientId: ""
			};

			findAllTreatments($scope, $http);
			getAllPatients($scope, $http);

			$scope.newTreatment = function () {
				$scope.accion = "NUEVO";
			};

			$scope.modifyTreatment = function () {
				$scope.accion = "MODIFICAR";
			};
		}]);

	function getAllPatients($scope, $http) {
		$http({
			method: 'GET',
			url: 'api/patients/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				console.log(response.data);
				$scope.patients = response.data;
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