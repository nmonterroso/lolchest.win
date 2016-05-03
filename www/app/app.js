'use strict';

/* App Module */

var lolApp = angular.module('lolApp', ['ngRoute', 'lolFilters', 'lolApp.profile', 'lolApp.home']);

lolApp.config(['$routeProvider',
	function($routeProvider) {
		$routeProvider
			.when('/404', {
				templateUrl: 'app/shared/404.html'
			})
			.otherwise({
				redirectTo: '/404'
			});
	}]);