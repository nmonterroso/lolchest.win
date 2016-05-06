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
			text: 'Champion Name',
			reverse: false,
			sortType: profileConst.sortTypeChamp
		},
		chestFirst: {
			text: 'Chest Earned',
			reverse: false,
			sortType: profileConst.sortTypeChest
		},
		chestLast: {
			text: 'Chest Available',
			reverse: true,
			sortType: profileConst.sortTypeChest
		},
		gradeBest: {
			text: 'Best Grade',
			reverse: false,
			sortType: profileConst.sortTypeGrade
		}
	};

	$scope.selectedSortOrder = $scope.sortOptions.chestLast;
}]);
