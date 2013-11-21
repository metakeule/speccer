'use strict';

var app = angular.module('app',  ["ngRoute"]);
/*
app.config(['$routeProvider', function($routeProvider) {
  $routeProvider.
    when('/success', {templateUrl: 'ok.html'}).
    otherwise({controller: 'Spec'});
}]);
*/


app.controller('Spec', function($scope, $http, $location) {
  $scope.Sections = {};
  $scope.states = ['PLANNING', 'APPROVED', 'PARTLY_IMPLEMENTED', 'FULLY_IMPLEMENTED', 'OBSOLET' ];
  
  /*
  $scope.loadSection = function(section) {
    $http.get('/spec.json?section='+section).success(function(data) { 
      if (section == "OVERVIEW") {
        $scope.OVERVIEW = data; 
        // $scope.setContent("OVERVIEW");
        // $scope.setContent("OVERVIEW");
        //<script src="http://strapdownjs.com/v/0.2/strapdown.js"></script>
        //document.write('<script src="http://strapdownjs.com/v/0.2/strapdown.js"></script>');
      } else {
        $scope.Sections[section] = data; 
      }
    });
  } 
  */

  $scope.loadInfo = function() {
    $http.get('/spec.json').success(function(data) { 
      console.log(data);
      $scope.INFO = data.INFO; 
        //$scope.loadSection('SCENARIO');
  //$scope.loadSection('FEATURE');
  //$scope.loadSection('DEFINITION');
  //$scope.loadSection('NONGOAL');
  //$scope.loadSection('UNDECIDED');
  //$scope.loadSection('CONTRADICTION');
  //$scope.loadSection('OVERVIEW');
      $scope.Sections.SCENARIO = data.Sections.SCENARIO; 
      $scope.Sections.FEATURE = data.Sections.FEATURE; 
      $scope.Sections.DEFINITION = data.Sections.DEFINITION; 
      $scope.Sections.NONGOAL = data.Sections.NONGOAL; 
      $scope.Sections.UNDECIDED = data.Sections.UNDECIDED; 
      $scope.Sections.CONTRADICTION = data.Sections.CONTRADICTION; 
      $scope.Sections.OVERVIEW = data.OVERVIEW; 
      //delete $scope.Spec.Sections;
      //console.log($scope.Spec);
    });
  }

  $scope.success = function() {
    $scope.fail_message = "";
    $scope.ok_message = "Success";
    // window.setTimeout($scope.cleanupMessages, 2000);
  }  

  $scope.cleanupMessages = function() {
    // console.log("do cleanup");
    $scope.fail_message = "";
    $scope.ok_message = "";
  }

  $scope.error = function(data) {
    $scope.fail_message = data;
    $scope.ok_message = "";
    // window.setTimeout($scope.cleanupMessages, 4000);
  }

  $scope.saveInfo = function() {
    var url = '/saveInfo';
    $http.post(url, $scope.INFO).success(function(data) {
      $scope.loadInfo();
      $scope.success();
    }).error(function(data){
      $scope.error(data);
    });
  };

  $scope.saveAll = function() {
    var url = '/saveAll';
    var spec = {
      INFO: $scope.INFO,
      OVERVIEW: $scope.Sections.OVERVIEW,
      Sections: {
        SCENARIO: $scope.Sections.SCENARIO,
        FEATURE: $scope.Sections.FEATURE,
        DEFINITION: $scope.Sections.DEFINITION,
        NONGOAL: $scope.Sections.NONGOAL,
        UNDECIDED: $scope.Sections.UNDECIDED,
        CONTRADICTION: $scope.Sections.CONTRADICTION
      }
    }
    $http.post(url, spec).success(function(data) {
      $scope.loadInfo();
      $scope.success();
    }).error(function(data){
      $scope.error(data);
    });
  };

  $scope.saveSection = function() {
    
    var url = '/saveSection?section='+$scope.currentSection+"&position="+$scope.sectionIndex;
    //var loc = this.$location;
    //console.log(loc);
    $http.post(url, $scope.paragraph).success(function(data) {
      /*
      if ($scope.sectionIndex == -1) {
        $scope.loadSection($scope.currentSection);
      }
      */
      $scope.loadInfo();
      $scope.success();
  //    $location.path("/success");
      //console.log(data);
    }).error(function(data){
      $scope.error(data);
    });
  };

  $scope.setComment = function(author, comment) {
    if (author) {
      $scope.CommentAuthor = author;
      $scope.CommentText = comment;
    }
  }

  $scope.setTranslations = function(lang, url) {
      $scope.TranslationsLanguage = lang;
      $scope.TranslationsURL = url;
  }

  $scope.saveTranslations = function() {
    if ($scope.TranslationsLanguage) {
      $scope.INFO.Translations[$scope.TranslationsLanguage] = $scope.TranslationsURL;
    }
  }  

  $scope.removeTranslations = function(lang) {
    if ($scope.TranslationsLanguage == lang) {
      $scope.TranslationsLanguage = "";
      $scope.TranslationsURL = "";
    }
    delete $scope.INFO.Translations[lang];
  }


  $scope.setRelated = function(name, url) {
      $scope.RelatedName = name;
      $scope.RelatedURL = url;
  }

  $scope.saveRelated = function() {
    if ($scope.RelatedName) {
      $scope.INFO.Related[$scope.RelatedName] = $scope.RelatedURL;
    }
  }  

  $scope.removeRelated = function(name) {
    if ($scope.RelatedName == name) {
      $scope.RelatedName = "";
      $scope.RelatedURL = "";
    }
    delete $scope.INFO.Related[name];
  }

  $scope.setSupersededBy = function(name, url) {
      $scope.SupersededByName = name;
      $scope.SupersededByURL = url;
  }

  $scope.saveSupersededBy = function() {
    if ($scope.SupersededByName) {
      $scope.INFO.SupersededBy[$scope.SupersededByName] = $scope.SupersededByURL;
    }
  }  

  $scope.removeSupersededBy = function(name) {
    if ($scope.SupersededByName == name) {
      $scope.SupersededByName = "";
      $scope.SupersededByURL = "";
    }
    delete $scope.INFO.SupersededBy[name];
  }

  $scope.setResources = function(name, url) {
      $scope.ResourcesName = name;
      $scope.ResourcesURL = url;
  }

  $scope.saveResources = function() {
    if ($scope.ResourcesName) {
      $scope.INFO.Resources[$scope.ResourcesName] = $scope.ResourcesURL;
    }
  }  

  $scope.removeResources = function(name) {
    if ($scope.ResourcesName == name) {
      $scope.ResourcesName = "";
      $scope.ResourcesURL = "";
    }
    delete $scope.INFO.Resources[name];
  }

  $scope.setPersons = function(short, full) {
      $scope.PersonsShort = short;
      $scope.PersonsFull = full;
  }

  $scope.savePersons = function() {
    if ($scope.PersonsShort) {
      $scope.INFO.Persons[$scope.PersonsShort] = $scope.PersonsFull;
    }
  }  

  $scope.removePersons = function(short) {
    if ($scope.PersonsShort == short) {
      $scope.PersonsShort = "";
      $scope.PersonsFull = "";
    }
    delete $scope.INFO.Persons[short];
  }

  $scope.addRequestedBy = function() {
    if ($scope.RequestedByName) {
      var found = false;
      for (var i = $scope.INFO.RequestedBy.length - 1; i >= 0; i--) {
        if ($scope.INFO.RequestedBy[i] === $scope.RequestedByName) {
          found = true;
        }
      };
      if (found == false) {
        $scope.INFO.RequestedBy[$scope.INFO.RequestedBy.length] = $scope.RequestedByName;
      }
    }
  }  

  $scope.removeRequestedBy = function(name) {
    if ($scope.RequestedByName == name) {
      $scope.RequestedByName = "";
    }
    for (var i = $scope.INFO.RequestedBy.length - 1; i >= 0; i--) {
        if ($scope.INFO.RequestedBy[i] === name) {
          $scope.INFO.RequestedBy.splice(i,1);
        }
      };
  }

  /*
  mapList1("Persons", "Persons", "short", "full"),
      mapList2("Persons", "Short", "Full"),
  */

  $scope.removeComment = function(author) {
    if ($scope.CommentAuthor == author) {
      $scope.CommentAuthor = "";
      $scope.CommentText = "";   
    }
    delete $scope.paragraph.Comments[author];
  }

  $scope.saveComment = function() {
    if ($scope.CommentAuthor) {
      $scope.paragraph.Comments[$scope.CommentAuthor] = $scope.CommentText;
    }
  }

  $scope.deleteSection = function() {
    var url = '/deleteSection?section='+$scope.currentSection+"&position="+$scope.sectionIndex;
    $http.post(url).success(function(data) {
      $scope.loadSection($scope.currentSection);
      $scope.success();
      $scope.newContent($scope.currentSection);
      //console.log(data);
    }).error(function(data){
      $scope.error(data);
    });
  };
  
  var myCodeMirror;

  function refreshText() {
    $('#Text .CodeMirror')[0].CodeMirror.refresh();
  }
 

  $scope.enableCodeHighlighting = function(text) {
    
    if ($scope.codeMirrorSet === false){
      
      myCodeMirror = CodeMirror($('#Text')[0], {
        // value: text,
        mode:  "markdown",
        dragDrop: false
      });

      myCodeMirror.on("change",   function() {
        $scope.paragraph.Text = myCodeMirror.getValue();
        $('#preview').html(markdown.toHTML(myCodeMirror.getValue()));        
      });
  
      $scope.codeMirrorSet = true;
      //return
      window.setTimeout(refreshText, 100);
    }
    $('#Text .CodeMirror')[0].CodeMirror.setValue(text);
  }

  $scope.setContent = function(section) {
    //$scope[section] = $scope[section+"S"][this.$index];
    $scope.currentSection = section;
    var text;
    if (section == "OVERVIEW") {
      $scope.sectionIndex = -1;
      text = $scope.Sections.OVERVIEW.Text;
      $scope.paragraph = $scope.Sections.OVERVIEW;
    } else {
      text = $scope.Sections[section][this.$index].Text;
      $scope.sectionIndex = this.$index;
      $scope.paragraph = $scope.Sections[section][this.$index];
    }
    $scope.enableCodeHighlighting(text);
    $scope.CommentAuthor = "";
    $scope.CommentAuthor = "";
    $scope.CommentText = "";
    $scope.cleanupMessages();
  }

  $scope.newContent = function(section) {
    //$scope[section] = {};
    $scope.paragraph = {Comments: {}, Text: ""};
    $scope.enableCodeHighlighting("");
    $scope.sectionIndex = -1;
    $scope.currentSection = section;
    $scope.cleanupMessages();
  }

  $scope.codeMirrorSet = false;
  //var myCodeMirror ;
  /*
  var myCodeMirror = CodeMirror($('#Text')[0], {
    value: "",
    mode:  "markdown"
  });
  */
  $scope.loadInfo();
  //$scope.loadSection('SCENARIO');
  //$scope.loadSection('FEATURE');
  //$scope.loadSection('DEFINITION');
  //$scope.loadSection('NONGOAL');
  //$scope.loadSection('UNDECIDED');
  //$scope.loadSection('CONTRADICTION');
  //$scope.loadSection('OVERVIEW');

  //myCodeMirror.setValue("hiho");

});

