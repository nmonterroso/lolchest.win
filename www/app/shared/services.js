'use strict';

/* Filters */

var lolChestServices = angular.module('lolChestServices', ['ngResource']);


lolChestServices.factory('Summoner', ['$resource',
	function($resource){
		return $resource('http://localhost:8080/summoner/:summonerName');
	}]);