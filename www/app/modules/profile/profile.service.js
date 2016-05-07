'use strict';

/* Profile Service */

var profileMod = angular.module('lolApp.profile');

profileMod.factory('Summoner', ['$resource', function($resource) {

	return $resource('http://localhost:8080/na/:summonerName');
}]);