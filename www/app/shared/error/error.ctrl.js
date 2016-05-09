'use strict';

angular
    .module('lolApp.error', ['ngResource'])
    .controller('ErrorCtrl', ['$location', '$scope', function($location, $scope) {
        $scope.message = "Uh Oh! An error occurred :(";

        switch ($location.search().type) {
            case "summoner_not_found":
                $scope.message = $location.search().name+" was not found";
                break;
            case "unknown_path":
                $scope.message = "page not found";
        }
    }]);
