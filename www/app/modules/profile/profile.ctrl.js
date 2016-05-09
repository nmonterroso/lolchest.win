'use strict';

/* Profile Controller */

var profileMod = angular.module('lolApp.profile');

profileMod.controller(
	'ProfileCtrl', 
	[
		'$location', '$scope','$routeParams', 'Summoner', 'profileConst', 'searchBarConst',
		function($location, $scope, $routeParams, Summoner, profileConst, searchBarConst) {

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
			}, function(resp) {
				$scope.loading = false;

				var code = resp.data.code || 500;
				var params = null;
				var regionObj = null;
				var regionName = null;

				switch (code) {
					case 404:
						regionObj = searchBarConst.getRegionObjFromLowerTag($routeParams.region);
						if (regionObj) {
							regionName = regionObj.name;
						}
						params = {
							type: 'summoner_not_found',
							name: $routeParams.summonerName,
							region: $routeParams.region,
							regionName: regionName
						};
						break;
				}

				$location.path("error");
				if (params != null) {
					$location.search(params);
				}
				$location.replace();
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
			};

			var profileHeight = null;
			$scope.ensureHeight = function() {
				if (profileHeight == null) {
					var profile = document.getElementById('profile');
					profileHeight = profile.offsetHeight;
					profile.style.height = profileHeight + 'px';
				}
			};
		}]);
