'use strict';

/* App Module */

var lolApp = angular.module('lolApp', ['ngRoute', 'lolFilters', 'lolApp.profile', 'lolApp.home', 'lolApp.searchBar']);

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

lolApp.controller('AppController', ['$rootScope', function($rootScope) {
	$rootScope.$on('$routeChangeSuccess', function (event, currentRoute) {
		console.log("ROUTE CHANGE", event, currentRoute);
		switch(currentRoute.templateUrl) {
			case 'app/modules/home/home.html':
				$rootScope.bodyClass = 'home-view';
				break;
			case 'app/modules/profile/profile.html':
				$rootScope.bodyClass = 'profile-view';
				break;
		}
	});
}]);
