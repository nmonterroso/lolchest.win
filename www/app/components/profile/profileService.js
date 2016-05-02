'use strict';

/* Profile Service */

var profileService = angular.module('ProfileService', ['ngResource']);

profileService.factory('Summoner', ['$resource', function($resource) {

	return $resource('http://localhost:8080/summoner/:summonerName');
}]);