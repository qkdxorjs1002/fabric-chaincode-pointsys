'use strict';

var app = angular.module('application', []);

app.controller('AppCtrl', function ($scope, appFactory) {
    $("#success_add").hide();
    $("#success_query").hide();
    $("#success_delete").hide();
    $("#success_update").hide();
    $scope.initAB = function () {
        appFactory.initAB($scope.memberid, function (data) {
            if (data == "")
                $scope.init_ab = "success";
            $("#success_add").show();
        });
    }
    $scope.queryAB = function () {
        appFactory.queryAB($scope.memberid, function (data) {
            $scope.query_ab = data;
            $("#success_query").show();
        });
    }

    $scope.deleteAB = function () {
        appFactory.deleteAB($scope.memberid, function (data) {
            $scope.delete_ab = data;
            $("#success_delete").show();
        });
    }

    $scope.updateAB = function () {
        appFactory.updateAB($scope.member, function (data) {
            $scope.update_ab = data;
            $("#success_update").show();
        });
    }
});

app.factory('appFactory', function ($http) {

    var factory = {};

    factory.initAB = function (name, callback) {
        $http.get('/member/add?name=' + name).success(function (output) {
            callback(output)
        });
    }
    factory.queryAB = function (name, callback) {
        $http.get('/member/query?name=' + name).success(function (output) {
            callback(output)
        });
    }

    factory.deleteAB = function (name, callback) {
        $http.get('/member/delete?name=' + name).success(function (output) {
            callback(output)
        });
    }

    factory.updateAB = function (data, callback) {
        $http.get('/point/update?name=' + data.name + '&point=' + data.point).success(function (output) {
            callback(output)
        });
    }
    return factory;
});