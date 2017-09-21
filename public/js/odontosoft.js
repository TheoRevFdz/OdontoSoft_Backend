(function () {
	var login = angular.module('login', [
		'login.controller'
	]);

	var patient = angular.module('patient', [
		'patient.controller',
		'patient.directives'
	]);

	var odonto = angular.module('odonto', [
		'odonto.directives'
	]);

	var control = angular.module('control', [
		'control.controller',
		'control.directives'
	]);
})();