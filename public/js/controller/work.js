(function () {
	angular.module('work.controller', [])
		.controller('WorkController', ['$scope', '$http', function ($scope, $http) {
			$scope.works = {};
			$scope.work = {
				ID: 0,
				name: "",
				price: 0
			};
			$scope.accion = "";

			findAllWorks($scope, $http);

			$scope.newWork = function () {
				$scope.accion = "NUEVA";
				clearFields($scope);
			};

			$scope.modifyWork = function (w) {
				$scope.accion = "MODIFICAR";
				asignedFields(w, $scope);
			};

			$scope.execute = function () {
				switch ($scope.accion) {
					case "NUEVA":
						console.log($scope.work);
						createWork($scope, $http);
						break;

					case "MODIFICAR":
						updateWork($scope, $http);
						console.log($scope.accion);
						break;
				}
			};

		}]);

	function findAllWorks($scope, $http) {
		$http({
			method: 'get',
			url: 'api/works/',
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				$scope.works = response.data;
			},
			function erroe(response) {
				alert(response.message);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}

	function createWork($scope, $http) {
		$http({
			method: 'post',
			url: 'api/crud/work/',
			headers: 'Content-Type: application/json',
			data: {
				ID: 0,
				name: $scope.work.name,
				price: parseFloat($scope.work.price)
			}
		}).then(
			function success(response) {
				alert(response.data.message);
				$scope.works = {};
				findAllWorks($scope, $http);
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

	function updateWork($scope, $http) {
		$http({
			method: 'put',
			url: 'api/crud/work/',
			headers: 'Content-Type: application/json',
			data: {
				ID: $scope.work.ID,
				name: $scope.work.name,
				price: parseFloat($scope.work.price)
			}
		}).then(
			function success(response) {
				alert(response.data.message);
				$scope.works = {};
				findAllWorks($scope, $http);
			},
			function erroe(response) {
				alert(response.message);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}

	function clearFields($scope) {
		$scope.work.ID = 0;
		$scope.work.name = "";
		$scope.work.price = 0;
	}

	function asignedFields(w, $scope) {
		$scope.work.ID = w.ID;
		$scope.work.name = w.name;
		$scope.work.price = w.price;
	}
})();