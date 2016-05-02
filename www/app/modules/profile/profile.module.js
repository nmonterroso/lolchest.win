'use strict';

/* Profile Module */

var profileMod = angular.module('app.profile', ['ngResource']);

profileMod.config(['$routeProvider',
	function($routeProvider) {
		$routeProvider
			.when('/profile/:summonerName', {
				templateUrl: 'app/modules/profile/profile.html',
				controller: 'ProfileCtrl'
			})
	}]);