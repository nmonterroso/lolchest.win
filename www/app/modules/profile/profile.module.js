'use strict';

/* Profile Module */

var profileMod = angular.module('lolApp.profile', ['ngResource']);

profileMod.constant('profileConst', {
	'sortTypeChamp': 'champ',
	'sortTypeGrade': 'grade',
	'sortTypeChest': 'chest',
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

profileMod.directive("profileHeader", function() {
	return {
		restrict: 'E',
		scope: {
			name: '=',
			icon: '=',
			earnedChests: '=',
			totalChests: '='
		},
		templateUrl: '/app/modules/profile/header/profile-header.html'
	};
});
