(function () {
	angular.module('patient.controller', [])
		.controller('PatientController', ['$scope', '$http', function ($scope, $http) {
			$scope.patients = {};

			findAllPatient($scope, $http);

			$scope.showNuevo = function () {
				console.log("Click!!");
				// $dialog.dialog({}).open('components/patient/form-patient.html');
				$('#frmPatient').modal('show');
			}

			$scope.loadPatients = function () {}
		}]);

	function findAllPatient($scope, $http) {
		// console.log("Patients");

		$http({
			method: 'get',
			url: 'api/patients/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				console.log(response.data);
				$scope.patients = response.data;
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