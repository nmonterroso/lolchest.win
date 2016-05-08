'use strict';

/* Searchbar Directive */

var searchBar = angular.module('lolSearchBar');

searchBar.directive("searchBar", ['$location', function($location) {
	return {
		restrict: 'E',
		scope: {
			formClasses: "@"
		},
		templateUrl: '/app/shared/search-bar/search-bar.html',
		link: function(scope) {
			scope.navigateTo = function(route) {
				$location.path(route)
			};
		}
	};
}]);