'use strict';

/* App Module */

var lolChestApp = angular.module('lolChestApp', ['ngRoute', 'lolChestFilters', 'ProfileService']);

lolChestApp.config(['$routeProvider',
	function($routeProvider) {
		$routeProvider
			.when('/', {
				templateUrl: 'app/components/home/homeView.html'
			})
			.when('/profile/:summonerName', {
				templateUrl: 'app/components/profile/profileView.html',
				controller: 'ProfileController'
			})
			.when('/404', {
				templateUrl: 'app/shared/404.html'
			})
			.otherwise({
				redirectTo: '/404'
			});
	}]);