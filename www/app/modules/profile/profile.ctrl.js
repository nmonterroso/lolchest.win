'use strict';

/* Profile Controller */

var profileMod = angular.module('lolApp.profile');

profileMod.controller('ProfileCtrl', ['$scope','$routeParams', 'Summoner', 'profileConst', function($scope, $routeParams, Summoner, profileConst) {

	$scope.loading = true;

	var summonerData = Summoner.get({
		region: $routeParams.region,
		summonerName: $routeParams.summonerName
	}, function() {
		$scope.loading = false;
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

	$scope.sortOptions = [
		{
			text: 'Champion Name',
			reverse: false,
			sortType: profileConst.sortTypeChamp
		},
		{
			text: 'Chest Earned',
			reverse: false,
			sortType: profileConst.sortTypeChest
		},
		{
			text: 'Chest Available',
			reverse: true,
			sortType: profileConst.sortTypeChest
		},
		{
			text: 'Best Grade',
			reverse: false,
			sortType: profileConst.sortTypeGrade
		}
	];

	$scope.selectedSortOrder = $scope.sortOptions[1];

	$scope.setSortOrder = function(index) {
		$scope.selectedSortOrder = $scope.sortOptions[index];
	}
}]);
