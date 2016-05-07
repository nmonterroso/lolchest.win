'use strict';

/* Searchbar Directive */

var searchBar = angular.module('lolApp.searchBar');

searchBar.directive("searchBar", function() {
	return {
		restrict: 'E',
		scope: {
		},
		templateUrl: '/app/shared/search-bar/search-bar.html'
	};
});