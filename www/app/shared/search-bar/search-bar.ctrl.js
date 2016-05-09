'use strict';

/* Profile Controller */

var searchBar = angular.module('lolSearchBar', []);

var REGIONS_LIST = [{
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

searchBar.constant('searchBarConst', {
	'regions': REGIONS_LIST,
	'getRegionObjFromLowerTag': function(tag) {
		for (var i = 0, l = REGIONS_LIST.length; i < l; i++) {
			var regionObj = REGIONS_LIST[i];
			if (regionObj.tag.toLowerCase() === tag) {
				return regionObj;
			}
		}
		return null;
	}
});

searchBar.controller('SearchBarCtrl', ['$routeParams', '$rootScope', '$scope', 'searchBarConst', function($routeParams, $rootScope, $scope, searchBarConst) {
	// looks for the region specified in the route params and tries to set selected region to that
	if ($routeParams && $routeParams.region) {
		$scope.selectedRegion = searchBarConst.getRegionObjFromLowerTag($routeParams.region);
	}

	$scope.regions = searchBarConst.regions;

	if (!$scope.selectedRegion) {
		$scope.selectedRegion = searchBarConst.regions[0];
	}

	$scope.setRegion = function(index) {
		$scope.selectedRegion = searchBarConst.regions[index];
	}
}]);
