'use strict';

/* Shared Filters */

var profileMod = angular.module('lolApp.profile');

profileMod
/**
 * Sort a list of champion masteries
 *
 * @param {Array} champMasteries - The list of champion masteries from profile
 * @param {String} sortOpts - Object containing sort options
 */
	.filter('sortMasteries', ['profileConst', function(profileConst) {

		// helper function for sorting strings alphabetically A-Z, or Z-A if reversed
		function sortAlphabetical(strA, strB, reverse) {
			var value = 0;
			if (strA > strB) {
				value = 1;
			} else {
				value = -1;
			}
			return reverse ? value * -1 : value;
		}

		// helper function for sorting masteries by grade S+ to D-, or reverse of.
		// if grades are equal, sort alphabetically
		function sortGrade(masteryA, masteryB, reverse) {
			var value = 0;
			var gradeA = profileConst.gradeOrder.indexOf(masteryA.highestGrade);
			var gradeB = profileConst.gradeOrder.indexOf(masteryB.highestGrade);
			if (gradeA === gradeB) {
				value = sortAlphabetical(masteryA.champName, masteryB.champName, false);
			} else {
				value = gradeA > gradeB ? 1 : -1;
				value = reverse ? value * -1 : value;
			}
			return value;
		}

		// helper function for sorting masteries by chest awarded to chest available, or reverse of.
		// if chest availabilities are equal, sort by grade
		function sortChest(masteryA, masteryB, reverse) {
			var value;
			if (masteryA.chestIsAvailable === masteryB.chestIsAvailable) {
				value = sortGrade(masteryA, masteryB, false);
			} else if (masteryA.chestIsAvailable) {
				value = reverse ? -1 : 1;
			} else {
				value = reverse ? 1 : -1;
			}
			return value;
		}

		// filter function
		return function(champMasteries, sortOpts) {
			if (!champMasteries) {
				return [];
			}

			var reverse = sortOpts.reverse;
			return champMasteries.sort(function(a, b) {
				switch(sortOpts.sortType) {
					case profileConst.sortTypeChamp:
						return sortAlphabetical(a.champName, b.champName, reverse);
					case profileConst.sortTypeChest:
						return sortChest(a, b, reverse);
					case profileConst.sortTypeGrade:
						return sortGrade(a, b, reverse);
				}
			});
		}
	}]);