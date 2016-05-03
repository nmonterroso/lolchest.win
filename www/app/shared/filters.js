'use strict';

/* Filters */

var lolFilters = angular.module('lolFilters', []);

lolFilters.filter('checkmark', function() {

	return function(input) {
		return input ? '\u2713' : '\u2718';
	};
});