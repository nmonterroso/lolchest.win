'use strict';

/* Profile Controller */

var searchBar = angular.module('lolApp.searchBar', []);

searchBar.controller('SearchBarCtrl', ['$scope', function($scope) {

	$scope.regions = [{
		tag: 'NA',
		name: 'North America'
	}, {
		tag: 'EUW',
		name: 'Europe West'
	}, {
		tag: 'EUNE',
		name: 'Europe Nordic & East'
	}, {
		tag: 'BR',
		name: 'Brazil'
	}, {
		tag: 'TR',
		name: 'Turkey'
	}, {
		tag: 'RU',
		name: 'Russia'
	}, {
		tag: 'KR',
		name: 'Korea'
	}, {
		tag: 'JP',
		name: 'Japan'
	}, {
		tag: 'LAN',
		name: 'Latin America North'
	}, {
		tag: 'LAS',
		name: 'Latin America South'
	}, {
		tag: 'OCE',
		name: 'Oceania'
	}];

	$scope.selectedRegion = $scope.regions[0];

	$scope.setRegion = function(index) {
		$scope.selectedRegion = $scope.regions[index];
	}
}]);
