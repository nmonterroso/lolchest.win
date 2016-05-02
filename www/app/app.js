'use strict';

/* App Module */

var lolChestApp = angular.module('app', ['ngRoute', 'lolChestFilters', 'app.profile', 'app.home']);

lolChestApp.config(['$routeProvider',
	function($routeProvider) {
		$routeProvider
			.when('/404', {
				templateUrl: 'app/shared/404.html'
			})
			.otherwise({
				redirectTo: '/404'
			});
	}]);