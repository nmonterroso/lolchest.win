'use strict';

/* Profile Controller */

var profileMod = angular.module('lolApp.profile');

profileMod.controller('ProfileCtrl', ['$scope','$routeParams', 'Summoner', 'profileConst', function($scope, $routeParams, Summoner, profileConst) {

	var summonerData = Summoner.get({
		summonerName: $routeParams.summonerName
	}, function() {
		$scope.masteries = summonerData.champMastery || [];
		$scope.summoner = {
			name: summonerData.name,
			icon: summonerData.profileIconUrl,
			earnedChests: 0,
			totalChests: 0
		};

		$scope.masteries.forEach(function(m) {
			++$scope.summoner.totalChests;
			if (!m.chestIsAvailable) {
				++$scope.summoner.earnedChests
			}
		});

	});

	$scope.sortOptions = {
		alpha: {
			text: 'Champion',
			reverse: false,
			sortType: profileConst.sortTypeChamp
		},
		chestFirst: {
			text: 'Chests Awarded',
			reverse: false,
			sortType: profileConst.sortTypeChest
		},
		chestLast: {
			text: 'Chests Available',
			reverse: true,
			sortType: profileConst.sortTypeChest
		},
		gradeBest: {
			text: 'Best Grades',
			reverse: false,
			sortType: profileConst.sortTypeGrade
		},
		gradeWorst: {
			text: 'Worst Grades',
			reverse: true,
			sortType: profileConst.sortTypeGrade
		}
	};

	$scope.selectedSortOrder = $scope.sortOptions.chestLast;
}]);
