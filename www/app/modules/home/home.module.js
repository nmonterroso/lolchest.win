'use strict';

/* Home Module */

var homeMod = angular.module('lolApp.home', []);

homeMod.config(['$routeProvider',
	function($routeProvider) {
		$routeProvider
			.when('/', {
				templateUrl: 'app/modules/home/home.html'
			})
	}]);