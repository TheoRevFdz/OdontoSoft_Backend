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
			console.log(moment("1991-08-04").format() + "Z");
			console.log(new Date(moment("2017-09-12").format()));
			console.log(moment(new Date()).format());

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
				doCreatePatient($scope, $http);
				// console.log($scope.patient.dateInit);
				// console.log($scope.patient.dateNac);
				console.log($scope.patient);
			}
		}]);

	function doCreatePatient($scope, $http) {
		$http({
			method: 'post',
			url: 'api/crud/patients/',
			headers: 'Content-Type: application/json',
			data: {
				ID: 0,
				dateInit: moment($scope.patient.dateInit).format() + "Z",
				nomApe: $scope.patient.nomApe,
				age: $scope.patient.age,
				sex: $scope.patient.sex,
				dateNac: moment($scope.patient.dateNac).format() + "Z",
				address: $scope.patient.address,
				ocupation: $scope.patient.ocupation,
				telCel: $scope.patient.telCel,
				alergies: $scope.patient.alergies,
				operations: $scope.patient.operations,
				diabettes: $scope.patient.diabettes,
				hipertension: $scope.patient.hipertension,
				others: $scope.patient.others,
				treatMedics: $scope.patient.treatMedics
			}
		}).then(
			function success(response) {
				alert(response.data);
				console.log(response);
			},
			function error(response) {
				alert(response);
				console.log(response);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}

	function clearFields($scope) {
		$scope.patient.ID = 0;
		$scope.patient.dateInit = "";
		$scope.patient.nomApe = "";
		$scope.patient.age = 0;
		$scope.patient.sex = "";
		$scope.patient.dateNac = "";
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