'use strict';

/* Filters */

var lolChestFilters = angular.module('lolChestFilters', []);

lolChestFilters.filter('checkmark', function() {

	return function(input) {
		return input ? '\u2713' : '\u2718';
	};
});