'use strict';

/* App Module */

var lolApp = angular.module('lolApp', ['ngRoute', 'lolFilters', 'lolSearchBar', 'lolApp.profile', 'lolApp.home', 'lolApp.error']);

lolApp.config(['$locationProvider', '$routeProvider',
	function($locationProvider, $routeProvider) {
		$locationProvider.html5Mode(true);

		$routeProvider
			.when('/error', {
				templateUrl: 'app/shared/error/error.html',
				controller: 'ErrorCtrl'
			})
			.otherwise({
				redirectTo: '/'
			});
	}]);

lolApp.controller('AppController', ['$rootScope', function($rootScope) {
	// kinda ghetto but we realized this too late :(
	var setResponsive = function(enabled) {
		var value = enabled ? 'width=device-width, initial-scale=1': '';
		document.getElementById('meta-responsive').setAttribute('content', value);
	};

	$rootScope.$on('$routeChangeSuccess', function (event, currentRoute) {
		setResponsive(false);

		switch(currentRoute.templateUrl) {
			case 'app/modules/home/home.html':
				$rootScope.bodyClass = 'home-view';
				break;
			case 'app/modules/profile/profile.html':
				$rootScope.bodyClass = 'profile-view';
				setResponsive(true);
				break;
			case 'app/shared/error/error.html':
				$rootScope.bodyClass = 'error-view';
				break;
		}
	});
}]);
