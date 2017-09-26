(function () {
	angular.module('treatment.controller', [])
		.controller('TreatmentController', ['$scope', '$http', function ($scope, $http) {
			$scope.accion = "";
			$scope.treatments = {};
			$scope.patients = {};
			$scope.works = {};
			$scope.treatment = {
				ID: 0,
				patientId: ""
			};
			$scope.treatmentDetail = {
				ID: 0,
				workId: ""
			};

			findAllTreatments($scope, $http);
			getAllPatients($scope, $http);
			getAllWorks($scope, $http);

			$scope.newTreatment = function () {
				$scope.accion = "NUEVO";
			};

			$scope.modifyTreatment = function () {
				$scope.accion = "MODIFICAR";
			};

			$scope.execute = function () {
				switch ($scope.accion) {
					case "NUEVO":
						createTreatment($scope, $http);
						break;
					case "MODIFICAR":
						updateTreatment($scope, $http);
						break;
				}
			}

			$scope.executeTD = function () {
				switch ($scope.accion) {
					case "NUEVO":
						createTD($scope, $http);
						break;
					case "MODIFICAR":
						updateTD($scope, $http);
						break;
				}
			};

		}]);

	function createTreatment($scope, $http) {
		$http({
			method: 'POST',
			url: 'api/crud/treatment/',
			headers: 'Content-Type: application/json',
			data: {
				ID: 0,
				patientId: $scope.treatment.patientId
			}
		}).then(
			function success(response) {
				alert(response.data.message);
				$scope.patients = {};
				findAllTreatments($scope, $http);
			},
			function erroe(response) {
				alert(response.data.message);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}

	function updateTreatment($scope, $http) {
		$http({
			method: 'PUT',
			url: 'api/crud/treatment/',
			headers: 'Content-Type: application/json',
			data: {
				ID: $scope.treatment.ID,
				patientId: $scope.treatment.patientId
			}
		}).then(
			function success(response) {
				alert(response.data.message);
				$scope.patients = {};
				findAllTreatments($scope, $http);
			},
			function erroe(response) {
				alert(response.data.message);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}

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