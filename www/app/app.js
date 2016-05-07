'use strict';

/* App Module */

var lolApp = angular.module('lolApp', ['ngRoute', 'lolFilters', 'lolApp.profile', 'lolApp.home']);

lolApp.config(['$locationProvider', '$routeProvider',
	function($locationProvider, $routeProvider) {
		$locationProvider.html5Mode(true);
		$routeProvider
			.when('/404', {
				templateUrl: 'app/shared/404.html'
			})
			.otherwise({
				redirectTo: '/404'
			});
	}]);