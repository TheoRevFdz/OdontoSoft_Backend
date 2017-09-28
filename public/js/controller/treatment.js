(function () {
	angular.module('treatment.controller', [])
		.controller('TreatmentController', ['$scope', '$http', function ($scope, $http) {
			$scope.accion = "";
			$scope.accionTD = "";
			$scope.treatments = {};
			$scope.patients = {};
			$scope.works = {};
			$scope.treatment = {
				ID: 0,
				patientId: "",
				select: false
			};
			$scope.treatmentDetail = {
				ID: 0,
				workId: "",
				quantity: 0,
				treatmentId: 0,
				precio: 0
			};

			$scope.btnState = true;

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

			$scope.addTD = function () {
				$scope.accionTD = "AGREGAR";
			};

			$scope.modifyTD = function () {
				$scope.accionTD = "MODIFICAR";
			};

			$scope.selectWork = function () {
				console.log($scope.treatmentDetail);
			}

			$scope.selectTreatmentRow = function (id) {
				console.log("ID: " + id);
				$scope.btnState = false;
				for (let i = 0; i < $scope.treatments.length; i++) {
					if ($scope.treatments[i].ID == id) {
						$scope.treatments[i].state = {
							background: "#51a0c7",
							color: "white"
						};
					} else {
						$scope.treatments[i].state = {
							background: "#eceeef",
							color: "black"
						};
					}
				}
				$scope.treatmentDetail.treatmentId = id;
			};

			$scope.trStyle = {
				color: "#F5F5F5",
				backgroundColor: '#51a0c7'
			}

			$scope.executeTD = function () {
				switch ($scope.accionTD) {
					case "AGREGAR":
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
				patientId: $scope.treatment.patientId.ID
			}
		}).then(
			function success(response) {
				alert(response.data.message);
				$scope.patients = {};
				findAllTreatments($scope, $http);
				findLastTreatment($scope, $http);
				console.log($scope.treatment);
				$scope.treatmentDetail.treatmentId = $scope.treatment.ID;
				// clearFielsTreatments($scope);
			},
			function error(response) {
				console.log($scope.treatment);
				alert(response.data.message);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}

	function findLastTreatment($scope, $http) {
		$http({
			method: 'GET',
			url: 'api/last-treatment/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				// console.log(response.data);
				$scope.treatment = response.data;
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

	function clearFielsTreatments($scope) {
		$scope.treatment.ID = 0;
		$scope.treatment.patientId = 0;
	}

})();