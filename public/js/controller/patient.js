(function () {
	angular.module('patient.controller', [])
		.controller('PatientController', ['$scope', '$http', function ($scope, $http) {
			$scope.patients = {};
			$scope.fec = moment().format('DD/MM/YYYY');
			$scope.fecInit = "";
			$scope.fecNac = "";

			$scope.patient = {
				ID: 0,
				dateInit: "",
				nomApe: "",
				age: 0,
				sex: "",
				dateNac: "",
				address: "",
				ocupation: "",
				telCel: "",
				alergies: "",
				operations: "",
				diabettes: false,
				hipertension: false,
				others: "",
				treatMedics: ""
			};

			$scope.accion = "";
			$scope.sex = ""

			findAllPatient($scope, $http);
			
			var f = new Date();
			console.log(moment("1991-08-04").format());

			$scope.showNuevo = function () {
				console.log("Click!!");
				// $dialog.dialog({}).open('components/patient/form-patient.html');
				// $('#frmPatient').modal('show');
			}

			$scope.loadPatients = function () {}

			$scope.newPatient = function () {
				clearFields($scope);
				$scope.accion = "NUEVO";
			}

			$scope.updatePatient = function (p) {
				console.log(p);
				clearFields($scope);
				$scope.accion = "MODIFICAR";
			}

			$scope.getSex = function () {
				// alert($scope.sex);
			}

			$scope.save = function () {
				// console.log($scope.patient.dateInit);
				// console.log($scope.patient.dateNac);
				console.log($scope.patient);
			}
		}]);

	function doCreatePatient(p, $http) {
		$http({
			method: 'post',
			url: 'api/patient/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {

			},
			function error(response) {

			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}

	function clearFields($scope) {
		$scope.patient.ID = 0;
		$scope.patient.dateInit = new Date();
		$scope.patient.nomApe = "";
		$scope.patient.age = 0;
		$scope.patient.sex = "";
		$scope.patient.dateNac = new Date("04/08/1991");
		$scope.patient.address = "";
		$scope.patient.ocupation = "";
		$scope.patient.telCel = "";
		$scope.patient.alergies = "";
		$scope.patient.operations = "";
		$scope.patient.diabettes = false;
		$scope.patient.hipertension = false;
		$scope.patient.others = "";
		$scope.patient.treatMedics = "";
	}

	function findAllPatient($scope, $http) {
		$http({
			method: 'get',
			url: 'api/patients/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				// console.log(response.data);
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