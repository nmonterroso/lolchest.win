'use strict';

/* Profile Module */

var profileMod = angular.module('lolApp.profile', ['ngResource']);

profileMod.constant('profileConst', {
	'sortTypeChamp': 'champ',
	'sortTypeGrade': 'grade',
	'sortTypeChest': 'chest',
	'gradeOrder': ['S+', 'S', 'S-', 'A+', 'A', 'A-', 'B+', 'B', 'B-', 'C+', 'C', 'C-', 'D+', 'D', 'D-', undefined]
});

profileMod.config(['$routeProvider',
	function($routeProvider) {
		$routeProvider
			.when('/:region/:summonerName', {
				templateUrl: 'app/modules/profile/profile.html',
				controller: 'ProfileCtrl'
			})
	}]);
