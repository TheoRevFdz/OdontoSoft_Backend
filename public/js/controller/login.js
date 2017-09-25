(function () {
	angular.module('login.controller', [])
		.controller('LoginController', ['$scope', '$http', function ($scope, $http) {
			$scope.login = {
				user: '',
				pass: ''
			};

			$scope.ingresar = function () {
				login($scope, $http);
			};
		}]);

	function login($scope, $http) {
		// console.log("Hola desde AngularJS!");
		$http({
			method: 'post',
			url: 'api/login',
			data: {
				username: $scope.login.user,
				password: $scope.login.pass
			},
			headers: 'Content-Type: application/json'
		}).then(
			function success(response) {
				alert("Acceso concedido!");
			},
			function error(response) {
				alert(response.data.message);
			}
		).catch(
			function fail(response) {
				console.log("Excep: " + response);
			}
		);
	}
})();
