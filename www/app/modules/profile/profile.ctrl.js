'use strict';

/* Profile Controller */

var profileMod = angular.module('app.profile');

profileMod.controller('ProfileCtrl', ['$scope','$routeParams', 'Summoner', function($scope, $routeParams, Summoner) {

	var summonerData = Summoner.get({
		summonerName: $routeParams.summonerName
	}, function() {
		$scope.masteries = summonerData.champMastery || [];
	});

	$scope.orderProp = 'champName'
}]);
