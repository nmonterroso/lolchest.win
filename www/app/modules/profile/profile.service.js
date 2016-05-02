'use strict';

/* Profile Service */

var profileMod = angular.module('app.profile');

profileMod.factory('Summoner', ['$resource', function($resource) {

	return $resource('http://localhost:8080/summoner/:summonerName');
}]);