'use strict';

angular
  .module('lolApp.error', ['ngResource'])
  .controller('ErrorCtrl', [
    '$location', '$scope',
    function($location, $scope) {
      $scope.message = "Uh Oh! An error occurred :(";
      var params = $location.search();

      switch (params.type) {
        case "summoner_not_found":
          $scope.message = 'Summoner "' + params.name + '" was not found';
          if (params.regionName) {
            $scope.message += ' in ' + params.regionName;
          }
          break;
      }
    }]);
