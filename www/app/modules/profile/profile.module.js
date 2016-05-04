'use strict';

/* Profile Module */

var profileMod = angular.module('lolApp.profile', ['ngResource']);

profileMod.constant('profileConst', {
	'sortTypeAlphabetical': 'a-z',
	'sortTypeGrade': 'grade',
	'gradeOrder': ['S+', 'S', 'S-', 'A+', 'A', 'A-', 'B+', 'B', 'B-', 'C+', 'C', 'C-', 'D+', 'D', 'D-', undefined] // TODO: move grade order elsewhere?
});

profileMod.config(['$routeProvider',
	function($routeProvider) {
		$routeProvider
			.when('/profile/:summonerName', {
				templateUrl: 'app/modules/profile/profile.html',
				controller: 'ProfileCtrl'
			})
	}]);