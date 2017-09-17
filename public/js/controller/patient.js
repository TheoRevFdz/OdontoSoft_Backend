(function () {
	angular.module('patient.controller', [])
		.controller('PatientController', ['$scope', '$http', function ($scope, $http) {
			$scope.patients = {

			};

			$scope.loadPatients = function () {
				findAllPatient($scope, $http);
			}
		}]);

	function findAllPatient($scope, $http) {
		alert("Patients");
	}
})();