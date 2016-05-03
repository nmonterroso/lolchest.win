'use strict';

/* Shared Filters */

var lolFilters = angular.module('lolFilters', []);

lolFilters
/**
 * Filter a list for elements with start with the given string.
 *
 * @param {Array} items - The list of items to filter
 * @param {String} prefix - The prefix to search for
 * @param {String} [itemProperty] - An optional property to use for the search
 *                                  if filtering a list of objects.
 */
	.filter('startsWith', function() {
		return function(items, prefix, itemProperty) {
			return items.filter(function(item) {
				if (prefix === undefined) {
					return true; // this is when the search prefix is blank
				}
				var findIn = itemProperty ? item[itemProperty] : item;
				return findIn.toString().toLowerCase().indexOf(prefix.toLowerCase()) === 0;
			});
		};
	});