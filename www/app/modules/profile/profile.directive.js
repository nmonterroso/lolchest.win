'use strict';

/* Profile Module */

var profileMod = angular.module('lolApp.profile');

profileMod.directive("profileHeader", function() {
	return {
		restrict: 'E',
		scope: {
			name: '=',
			icon: '=',
			earnedChests: '=',
			totalChests: '='
		},
		templateUrl: '/app/modules/profile/partials/profile-header.html'
	};
});

profileMod.directive("profileNavBar", function() {
	return {
		restrict: 'E',
		scope: false,
		templateUrl: '/app/modules/profile/partials/profile-nav-bar.html'
	};
});
