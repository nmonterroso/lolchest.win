'use strict';

/* Profile Controller */

var lolChestApp = angular.module('lolChestApp');

lolChestApp.controller('ProfileController', ['$scope','$routeParams', 'Summoner', function($scope, $routeParams, Summoner) {

	var summonerData = Summoner.get({
		summonerName: $routeParams.summonerName
	}, function() {
		$scope.masteries = summonerData.champMastery || [];
	});

	$scope.orderProp = 'champName'
}]);
