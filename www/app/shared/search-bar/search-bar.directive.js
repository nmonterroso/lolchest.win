'use strict';

/* Searchbar Directive */

var searchBar = angular.module('lolSearchBar');

searchBar.directive("searchBar", ['$window', '$location', function($window, $location) {
	return {
		restrict: 'E',
		scope: {
			formClasses: "@",
			allowInlineBypass: "="
		},
		templateUrl: '/app/shared/search-bar/search-bar.html',
		link: function(scope) {
			var breakpoint = 984;
			scope.navigateTo = function(route) {
				$location.path(route)
			};

			var w = angular.element($window);
			scope.inlineRegionSelector = w.width() > breakpoint;

			scope.$watch(function() {
				return w.width();
			}, function (val, oldVal) {
				if (oldVal != val) {
					scope.inlineRegionSelector = val > breakpoint;
				}
			});

			w.bind('resize', function() {
				scope.$apply();
			});
		}
	};
}]);
