'use strict';

/* Profile Controller */

var profileMod = angular.module('lolApp.profile');

profileMod.controller('ProfileCtrl', ['$scope','$routeParams', 'Summoner', 'profileConst', function($scope, $routeParams, Summoner, profileConst) {

	var summonerData = Summoner.get({
		summonerName: $routeParams.summonerName
	}, function() {
		$scope.masteries = summonerData.champMastery || [];
	});

	$scope.sortOptions = {
		alpha: {
			text: 'Champion',
			reverse: false,
			sortType: profileConst.sortTypeAlphabetical,
			property: 'champName'
		},
		revAlpha: {
			text: 'Reverse Champion',
			reverse: true,
			sortType: profileConst.sortTypeAlphabetical,
			property: 'champName'
		},
		chestFirst: {
			text: 'Chests Awarded',
			reverse: false,
			sortType: profileConst.sortTypeAlphabetical,
			property: 'chestIsAvailable'
		},
		chestLast: {
			text: 'Chests Available',
			reverse: true,
			sortType: profileConst.sortTypeAlphabetical,
			property: 'chestIsAvailable'
		},
		gradeBest: {
			text: 'Best Grades',
			reverse: false,
			sortType: profileConst.sortTypeGrade,
			property: 'highestGrade'
		},
		gradeWorst: {
			text: 'Worst Grades',
			reverse: true,
			sortType: profileConst.sortTypeGrade,
			property: 'highestGrade'
		}
	};

	// TODO: sort within sort, eg. sort by grades after sorting by chests
	$scope.sortFunc = function(item) {
		var property = item[$scope.selectedSortOrder.property];
		switch($scope.selectedSortOrder.sortType) {
			case profileConst.sortTypeAlphabetical:
				return property;
			case profileConst.sortTypeGrade:
				return profileConst.gradeOrder.indexOf(property);
		}
	};

	$scope.selectedSortOrder = $scope.sortOptions.gradeBest;
}]);
