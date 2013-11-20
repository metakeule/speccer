'use strict';

var app = angular.module('app', []);

app.controller('Spec', function($scope, $http) {
  $http.get('/spec.json?section=SCENARIO').success(function(data) {
    $scope.SCENARIOS = data
  });
  $http.get('/spec.json?section=FEATURE').success(function(data) {
    $scope.FEATURES = data
  });
  $http.get('/spec.json?section=DEFINITION').success(function(data) {
    $scope.DEFINITIONS = data
  });
  $http.get('/spec.json?section=NONGOAL').success(function(data) {
    $scope.NONGOALS = data
  });
  $http.get('/spec.json?section=UNDECIDED').success(function(data) {
    $scope.UNDECIDEDS = data
  });

  $scope.updateSection = function(section) {
    //var x = $scope[section];
    //x.Position = this.$index;    
    //x.Section = section;
    //console.log(x);
    $http.post('/updateSection?section='+section+"&position="+$scope.sectionIndex, $scope[section]).success(function(data) {
      console.log(data);
    });
  };
  $scope.createSection = function() {
    console.log(this);
    /*
    if (this.text) {
      this.list.push(this.text);
      this.text = '';
    }
    */
  };

  $scope.setContent = function(section) {
    $scope[section] = $scope[section+"S"][this.$index];
    $scope.sectionIndex = this.$index;
    console.log(section);
    console.log(this.$index);
    console.log(this);
  }
});

