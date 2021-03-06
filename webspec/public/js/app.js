'use strict';

var app = angular.module('app',  ["ngRoute", "ngAnimate"]);
/*
app.config(['$routeProvider', function($routeProvider) {
  $routeProvider.
    when('/success', {templateUrl: 'ok.html'}).
    otherwise({controller: 'Spec'});
}]);
*/

/*
app.animation('.alert', function() {
  return {
    addClass: function(element, className,done) { 
      console.log("on addClass");
      $(element).fadeOut(3000,done);
    },

    removeClass: function(element, className, done) { 
      console.log("on removeClass");
      $(element).fadeOut(3000,done);
    }
  }
});
*/

app.controller('Spec', function($scope, $http, $location) {
  $scope.Sections = {};
  $scope.states = ['PLANNING', 'AGREED', 'IMPLEMENTING', 'FINISHED', 'OBSOLET' ];

  $scope.filter = {
    "PLANNING": false, 
    "AGREED": false,
    "IMPLEMENTING": false, 
    "FINISHED": false, 
    "OBSOLET": false
  };

  $scope.resetCount = function(state) {
    $scope.count = {
      "PLANNING": 0, 
      "AGREED": 0,
      "IMPLEMENTING": 0, 
      "FINISHED": 0, 
      "OBSOLET": 0
    };
  }

  $scope.resetCount();

  $scope.isFiltered = function(state, resp) {
    //isFiltered
    if ($scope.filteredPerson != "ALL") {
      if (resp != $scope.filteredPerson) {
        return true
      }
    }
    return $scope.filter[state];
  }
  //$scope.planning = false;
  
  /*
  $scope.setFilter = function() {
    // console.log(state);
    console.log($scope.filter);
    console.log(this);
  }
  */
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

  $scope.filteredPerson = "ALL";


$scope.loadTranslations = function(callback) {
    $http.get(window.location.pathname+"/../translations.json").success(function(data) { 
      $scope.translations = data;
      if (callback) {
        callback();
      }
    });
  }


  $scope.loadInfo = function(callback) {
    $http.get(window.location.pathname+'/../spec.json').success(function(data) { 
      //console.log(data);
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
      $scope.personFilter = ["ALL"];
      //$scope.personFilter = [];
      $scope.persons = [];
      $.each($scope.INFO.Persons, function(a,b){
        $scope.persons.push(a);
        $scope.personFilter.push(a);
        //console.log(a);
        //console.log(b);
      });

      $scope.resetcountResponsibles();
      $scope.setSectionNumbers();
      //console.log($scope.countResponsibles);

      //console.log($scope.persons);
      //delete $scope.Spec.Sections;
      //console.log($scope.Spec);
      if (callback) {
        callback();
      }
    });
  }

/*
$scope.store();
    //$scope[section] = $scope[section+"S"][this.$index];
    var text;
    if ($scope.currentSection == "OVERVIEW") {
      $scope.sectionIndex = -1;
      text = $scope.Sections.OVERVIEW.Text;
      $scope.paragraph = $scope.Sections.OVERVIEW;
    } else {
      text = $scope.paragraphs[this.$index].Text;
      $scope.sectionIndex = this.$index;
      $scope.paragraph = $scope.paragraphs[this.$index];
    }
    $scope.enableCodeHighlighting(text);
    $scope.CommentAuthor = "";
    $scope.CommentAuthor = "";
    $scope.CommentText = "";
    $scope.cleanupMessages();
    $scope.textchanged = false;
*/

  var valueOfNotFinishedness = {
    //"PLANNING": 0.95,
    "PLANNING": 1,
    "AGREED": 0.8,
    "IMPLEMENTING": 0.75
  }

  function progressForSection(paragraphs, progress) {
    var notFinished = 0;
    var notAgreed = 0; 
    var total = 0;
    for (var i = paragraphs.length - 1; i >= 0; i--) {
       if ($scope.filteredPerson != "ALL") {
        if (paragraphs[i].Responsible !=  $scope.filteredPerson) {
          continue;
        }
       }
       total++;
       if (progress) {            
          $scope.count[paragraphs[i].State] += 1;
       }
       var st = paragraphs[i].State
       if (st != "FINISHED" && st != "OBSOLET") {
          if (progress) {            
            notFinished = notFinished + valueOfNotFinishedness[st];
          } else {
            notFinished++;
          }
          if (st == "PLANNING") {
            notAgreed++;
          }
       }       
    };
    return {
      "notFinished": notFinished,
      "notAgreed": notAgreed,
      "total": total
    }
  }

  $scope.sectionNumbers = {};

  $scope.setSectionNumbers = function() {
    var allNotA = 0;
    var allNotF = 0;
    var allTotal = 0;
    console.log("setSectionNumbers called");
      $scope.sectionNumbers = {};
      $.each($scope.Sections, function(name, sec){
        if (name != "OVERVIEW") {
          var r = progressForSection(sec, false);
          var notA = Math.floor(r.notAgreed);
          var notF = Math.floor(r.notFinished);
          if (r.total > 0) {
          //  if (notA + notF > 0) {
            allNotA += notA;
            allNotF += notF;
            allTotal += r.total,
            $scope.sectionNumbers[name] = "" + notA + "/" + notF + "/" + r.total;
            //$scope.sectionNumbers[name] = "" + notA + "/" + notF ;
          }
        }
      });
      if (allTotal > 0) {
        $scope.sectionNumbers["OVERVIEW"] = "" + allNotA + "/" + allNotF + "/" + allTotal;
      }
  }

  $scope.setCounts = function() {
    $scope.resetCount();
    $scope.setSectionNumbers();
    $scope.progressFinished = 0;
    $scope.progressNotAgreed = 0;
    //console.log("set counts called for "+ $scope.currentSection);

    if ($scope.currentSection == "OVERVIEW") {
      //$scope.Se
      var total = 0;
      var notFinished = 0;
      var notAgreed = 0;
      $.each($scope.Sections, function(name, sec){
        if (name != "OVERVIEW") {
          var r = progressForSection(sec, true);
          total += r.total;
          notFinished += r.notFinished;
          notAgreed += r.notAgreed;          
        }
      });

      if (total != 0) {
        $scope.progressFinished = Math.floor(((total - notFinished) * 100) / total);
        $scope.progressNotAgreed = Math.floor((notAgreed * 100) / total);
      }
    } else {
      var r = progressForSection($scope.paragraphs, true);
      if (r.total != 0) {
        $scope.progressFinished = Math.floor(((r.total - r.notFinished) * 100) / r.total);
        $scope.progressNotAgreed = Math.floor((r.notAgreed * 100) / r.total);
      }
    }
    //var total = $scope.paragraphs.length;
     
  }

  $scope.filterClass = function(state) {
    var active = $scope.filter[state] ? "filteractive " : "";
    if ($scope.count[state] > 0 && $scope.currentSection != "OVERVIEW" ) {
      return active + "btn-primary";
    }
    /*
    var active = $scope.filter[state] ? "active " : "";
    if ($scope.count[state] > 0) {
      if ($scope.filter[state]) {
        return active + "btn-danger";
      }
      return active + "btn-success";
    } 
    return active + "btn-default";
    */
    return active + "btn-default";
  }

  $scope.resetcountResponsibles = function() {
    $scope.countResponsibles = {"ALL": 0};
    $.each($scope.INFO.Persons, function(a,b){
        $scope.countResponsibles[a] = 0;
    });
  }

  $scope.setcountResponsibles = function(paragraph) {
    for (var i = paragraph.length - 1; i >= 0; i--) {
       $scope.countResponsibles["ALL"] += 1;
       $scope.countResponsibles[paragraph[i].Responsible] += 1;
    };
  }
  
  $scope.setParagraphs = function(section) {
    $scope.store();
    //$scope[section] = $scope[section+"S"][this.$index];
    $scope.currentSection = section;
    if (section == "OVERVIEW") {
      //console.log("setting OVERVIEW");
      $scope.paragraphs = [];
      $scope.resetCount();
      $scope.setPara();
      $scope.resetcountResponsibles();
      $.each($scope.Sections, function(name, sec){
        if (name != "OVERVIEW") {
          $scope.setcountResponsibles(sec);
        }
      });
      //$scope.setcountResponsibles($scope.paragraphs);
      $scope.setCounts();
    } else {
      // $scope.store(true);    
    //$scope[section] = {};
    //$scope.paragraph = {Comments: {}, Text: ""};
      $scope.paragraph = {Comments: {}, Text: ""};
      $scope.enableCodeHighlighting("");
      $scope.sectionIndex = -1;

      //$scope.paragraph = null;
      /*
      if ($scope.codeMirrorSet) {
        //console.log("resetting CodeMirror");
        $scope.enableCodeHighlighting("");
        //$('#Text .CodeMirror')[0].CodeMirror.refresh();
        window.setTimeout(refreshText, 100);
      } 
      */    


      $scope.paragraphs = $scope.Sections[section];
      $scope.resetcountResponsibles();
      $scope.setcountResponsibles($scope.paragraphs);
      /*
      for (var i = $scope.paragraphs.length - 1; i >= 0; i--) {
         $scope.countResponsibles["all"] += 1;
         $scope.countResponsibles[$scope.paragraphs[i].Responsible] += 1;
      };
      */
      $scope.setCounts();
      $scope.CommentAuthor = "";
      $scope.CommentAuthor = "";
      $scope.CommentText = "";
      $scope.sectionform.$setPristine();
      //$scope.textchanged = false;
    //  $scope.cleanupMessages();
    }
  }

  $scope.success = function() {
    $scope.fail_message = "";
    $scope.ok_message = "Saved";
    window.setTimeout($scope.cleanupMessages, 3000);
  }  

  $scope.cleanupMessages = function() {
  //  console.log("do cleanup");
    $scope.fail_message = "";
    $scope.ok_message = "";
  }

  $scope.error = function(data) {
    $scope.fail_message = data;
    $scope.ok_message = "";
    // window.setTimeout($scope.cleanupMessages, 4000);
  }

  $scope.saveInfo = function() {
    var url = window.location.pathname+'/../saveInfo';
    $scope.infoform.$setPristine();
    $http.post(url, $scope.INFO).success(function(data) {
      $scope.loadInfo();
      $scope.success();
    }).error(function(data){
      $scope.error(data);
    });
  };

  /*
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
  */

  $scope.progressFinished = 0;
  $scope.progressNotAgreed = 0;

  $scope.translations = {};

  $scope.translate = function (s) {
    if ($scope.translations[s]) {
      return $scope.translations[s];
    }
    return s;
  }

  $scope.validate = function() {
    $scope.store();
    $scope.paragraph = null;
    $scope.currentSection = "";
    $scope.sectionIndex = -1;
    $scope.paragraphs = [];
    $http.post(window.location.pathname+"/../validate").success(function(data) {
      $scope.validationMessage = data;
    });
  }

  $scope.changeParagraphSection = function(position, targetsection) {
    $scope.store();
    var url = window.location.pathname+'/../changeParagraphSection?section='+$scope.currentSection+"&position="+position+"&targetsection="+targetsection;
    $scope.sectionform.$setPristine();

    $http.post(url).success(function(data) {
      var fn = function() {
        $scope.setParagraphs($scope.currentSection);        
      };
      $scope.loadInfo(fn);
      $scope.success();      
    }).error(function(data){
      $scope.error(data);
    });
  }

  $scope.moveParagraph = function(from, to) {
    $scope.store();
    var url = window.location.pathname+'/../moveParagraph?section='+$scope.currentSection+"&from="+from+"&to="+to;
    $scope.sectionform.$setPristine();

    $http.post(url).success(function(data) {
      var fn = function() {
        $scope.setParagraphs($scope.currentSection);
        /*
        $scope.sectionIndex = to;
        //console.log("setting " + $scope.currentSection + " at "+ $scope.sectionIndex );
        var text = $scope.paragraphs[$scope.sectionIndex].Text;
        $scope.paragraph = $scope.paragraphs[$scope.sectionIndex];
        $scope.enableCodeHighlighting(text);      
        $scope.CommentAuthor = "";
        $scope.CommentAuthor = "";
        $scope.CommentText = "";
        */
      };
      $scope.loadInfo(fn);
      $scope.success();      
    }).error(function(data){
      $scope.error(data);
    });
  }

  $scope.saveSection = function(reload) {
    
    var url = window.location.pathname+'/../saveSection?section='+$scope.currentSection+"&position="+$scope.sectionIndex;
    //var loc = this.$location;
    //console.log(loc);
    $scope.sectionform.$setPristine();
    //console.log($scope.paragraph);
    $http.post(url, $scope.paragraph).success(function(data) {
      /*
      if ($scope.sectionIndex == -1) {
        $scope.loadSection($scope.currentSection);
      }
      */
      
      if (reload) {
        var fn = function() {
          var oldIndex = $scope.sectionIndex
          $scope.setParagraphs($scope.currentSection);
          if ($scope.currentSection != "OVERVIEW" && oldIndex != -1) {
              $scope.sectionIndex = oldIndex;
              //console.log("setting " + $scope.currentSection + " at "+ $scope.sectionIndex );
              var text = $scope.paragraphs[$scope.sectionIndex].Text;
              $scope.paragraph = $scope.paragraphs[$scope.sectionIndex];
              $scope.enableCodeHighlighting(text);
              $scope.CommentAuthor = "";
              $scope.CommentAuthor = "";
              $scope.CommentText = "";
              $scope.sectionform.$setPristine();
              //$('#paragraph-list li#paragraph-'+$scope.sectionIndex).addClass("active");
          }
        };

        $scope.loadInfo(fn);
      } else {
        $scope.loadInfo();
      }

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

  $scope.setCommentForAuthor = function() {
    $scope.CommentText = $scope.paragraph.Comments[$scope.CommentAuthor];
  }

  $scope.setTranslations = function(lang, url) {
      $scope.TranslationsLanguage = lang;
      $scope.TranslationsURL = url;
  }

  $scope.saveTranslations = function() {
    if ($scope.TranslationsLanguage) {
      //$scope.textchanged = true;
      $scope.infoform.$setDirty();
      $scope.INFO.Translations[$scope.TranslationsLanguage] = $scope.TranslationsURL;
    }
  }  

  $scope.removeTranslations = function(lang) {
    $scope.infoform.$setDirty();
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
      $scope.infoform.$setDirty();
      $scope.INFO.Related[$scope.RelatedName] = $scope.RelatedURL;
    }
  }  

  $scope.removeRelated = function(name) {
    $scope.infoform.$setDirty();
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
      $scope.infoform.$setDirty();
      $scope.INFO.SupersededBy[$scope.SupersededByName] = $scope.SupersededByURL;
    }
  }  

  $scope.removeSupersededBy = function(name) {
    $scope.infoform.$setDirty();
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
      $scope.infoform.$setDirty();
      $scope.INFO.Resources[$scope.ResourcesName] = $scope.ResourcesURL;
    }
  }  

  $scope.removeResources = function(name) {
    $scope.infoform.$setDirty();
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
      $scope.infoform.$setDirty();
      $scope.INFO.Persons[$scope.PersonsShort] = $scope.PersonsFull;
    }
  }  

  $scope.removePersons = function(short) {
    $scope.infoform.$setDirty();
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
        $scope.infoform.$setDirty();
        $scope.INFO.RequestedBy[$scope.INFO.RequestedBy.length] = $scope.RequestedByName;
      }
    }
  }  

  $scope.removeRequestedBy = function(name) {
    $scope.infoform.$setDirty();
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
    $scope.sectionform.$setDirty();
    delete $scope.paragraph.Comments[author];
  }

  $scope.saveComment = function() {
    if ($scope.CommentAuthor) {
      $scope.paragraph.Comments[$scope.CommentAuthor] = $scope.CommentText;
      $scope.sectionform.$setDirty();
    }
  }

  $scope.deleteSection = function() {
    var url = window.location.pathname+'/../deleteSection?section='+$scope.currentSection+"&position="+$scope.sectionIndex;
    $scope.sectionform.$setPristine();
    $http.post(url).success(function(data) {
      //$scope.loadSection($scope.currentSection);
      var fn = function() {
        $scope.setParagraphs($scope.currentSection);
      };
      $scope.loadInfo(fn);
      $scope.success();
      //$scope.newContent($scope.currentSection);
      //console.log(data);
    }).error(function(data){
      $scope.error(data);
    });
  };
  
  var myCodeMirror;

  function refreshText() {
    //console.log("refreshing codemirror");
    $('#Text .CodeMirror')[0].CodeMirror.refresh();
    //$scope.codeMirrorSet = true;
  }
 

//  $scope.textchanged = false;

  $scope.enableCodeHighlighting = function(text) {
    
    if ($scope.codeMirrorSet === false){
      
      myCodeMirror = CodeMirror($('#Text')[0], {
        // value: text,
        mode:  "markdown",
        dragDrop: false,
        // theme: "monokai"
        //theme: "xq-light"
        //theme: "eclipse"
        theme: "webspec"
      });

      myCodeMirror.on("change",   function() {
        if ($scope.paragraph) {
          $scope.paragraph.Text = myCodeMirror.getValue();
          $scope.sectionform.$setDirty();
        } 
        $('#preview').html(markdown.toHTML(myCodeMirror.getValue()));        
      });
  
      $scope.codeMirrorSet = true;
      window.setTimeout(refreshText, 100);
      //return
    }
    $('#Text .CodeMirror')[0].CodeMirror.setValue(text);
    $scope.sectionform.$setPristine();
  }

/*
  $scope.dirty = function() {
    return ($scope.textchanged || $('#section-form').hasClass("ng-dirty") || $('#info-form').hasClass("ng-dirty"));
  }
*/

  $scope.store = function(reload) {
    // console.log("store called");
    if ($scope.currentSection) {
      // console.log("has current section " + $scope.currentSection);
      if ($scope.sectionform.$dirty) {
        //console.log("storing section");
        if (reload) {
          $scope.saveSection(true);
        } else {
          $scope.saveSection(false);
        }
        
        //$('#section-form').removeClass("ng-dirty");
      }
    } else {
      if ($scope.infoform.$dirty ) {
        //console.log("storing info");
        //$scope.saveAll();
        $scope.saveInfo();
        
        //$('#info-form').removeClass("ng-dirty");
      }
    }
    //$scope.textchanged = false;
  }

  $scope.forcestore = function() {
    if ($scope.currentSection) {
      //console.log("storing section");
      $scope.saveSection(true);
      //$('#section-form').removeClass("ng-dirty");
    } else {
      //console.log("storing info");
      $scope.saveInfo();
      //$('#info-form').removeClass("ng-dirty");
    }
    //$scope.textchanged = false;
  }

  $scope.unsetCurrentSection = function() {
    $scope.store();
    //console.log("unsetting current section");
    $scope.currentSection = null;
    $scope.paragraphs = [];
    $scope.resetCount();
    $scope.resetcountResponsibles();
  }

  $scope.handleDropped = function(item, box, x, y) {
    //console.log(item);  
    //console.log(box);  
    var boxid = $(box).attr('id');
    var from = $(item).attr("id");
    var halfheight = $("#paragraph-list li").height() / 2 ;
    //var halfheight = 0
    if (boxid == 'paragraph-list') {
      var to = 0;
      $("#paragraph-list li").each(
        function() {
          var $this = $(this);
       //   console.log(this);  
         // console.log("top: "+$this.offset().top + " halfheight: " + halfheight + " y: "+y);
          if (($this.offset().top - halfheight)< y) {
            to =  parseInt($this.attr("id"), 10) ;
           // console.log("set to "+$this.attr("id"));
          }
        }
      );
      // remove the first item ("new")
     //to = to + 1;
     //console.log(to);
      if (from != to) {
        $scope.moveParagraph(from, to);
      }
    } else {
      var targetSection = $(box).attr('section');
      if (targetSection != $scope.currentSection) {
        $scope.changeParagraphSection(from, targetSection);
      }
      // console.log(targetSection);
    }
    
  }

  $scope.setPara = function() {
    $scope.store(true);
    //$scope[section] = $scope[section+"S"][this.$index];
    var text;
    if ($scope.currentSection == "OVERVIEW") {
      $scope.sectionIndex = -1;
      text = $scope.Sections.OVERVIEW.Text;
      $scope.paragraph = $scope.Sections.OVERVIEW;
    } else {
      text = $scope.paragraphs[this.$index].Text;
      $scope.sectionIndex = this.$index;
      $scope.paragraph = $scope.paragraphs[this.$index];
    }
    $scope.enableCodeHighlighting(text);
    $scope.CommentAuthor = "";
    $scope.CommentAuthor = "";
    $scope.CommentText = "";
    //$scope.cleanupMessages();
    $scope.sectionform.$setPristine();
    //$scope.textchanged = false;
  }

/*
  $scope.newPara = function() {
    $scope.store(true);    
    //$scope[section] = {};
    $scope.paragraph = {Comments: {}, Text: ""};
    $scope.enableCodeHighlighting("");
    $scope.sectionIndex = -1;
    $scope.CommentAuthor = "";
    $scope.CommentAuthor = "";
    $scope.CommentText = "";
    //$scope.cleanupMessages();
    $scope.sectionform.$setPristine();
  }
*/
/*
  $scope.setContent = function(section) {
    $scope.store();
    //$scope[section] = $scope[section+"S"][this.$index];
    $scope.currentSection = section;
    var text;
    if (section == "OVERVIEW") {
      $scope.sectionIndex = -1;
      text = $scope.Sections.OVERVIEW.Text;
      $scope.paragraph = $scope.Sections.OVERVIEW;
      $scope.setParagraphs("OVERVIEW");
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
    $scope.textchanged = false;
  }
*/

/*
  $scope.newContent = function(section) {
    $scope.store();    
    //$scope[section] = {};
    $scope.paragraph = {Comments: {}, Text: ""};
    $scope.enableCodeHighlighting("");
    $scope.sectionIndex = -1;
    $scope.currentSection = section;
    $scope.cleanupMessages();
  }
*/
  $scope.codeMirrorSet = false;
  //var myCodeMirror ;
  /*
  var myCodeMirror = CodeMirror($('#Text')[0], {
    value: "",
    mode:  "markdown"
  });
  */

  $scope.loadTranslations(function(){
    $scope.loadInfo();
  });
  //$scope.loadSection('SCENARIO');
  //$scope.loadSection('FEATURE');
  //$scope.loadSection('DEFINITION');
  //$scope.loadSection('NONGOAL');
  //$scope.loadSection('UNDECIDED');
  //$scope.loadSection('CONTRADICTION');
  //$scope.loadSection('OVERVIEW');

  //myCodeMirror.setValue("hiho");
  //$scope.textchanged = false;



//  $("#ok-message").on();
//  $("#fail-message").on();
});




/*
  the following DnD is stolen from  

  http://blog.parkji.co.uk/2013/08/11/native-drag-and-drop-in-angularjs.html

  code:
  http://codepen.io/parkji/pen/JtDro
*/

app.directive('draggable', function() {
  return function(scope, element) {
    // this gives us the native JS object
    var el = element[0];
    
    el.draggable = true;
    
    el.addEventListener(
      'dragstart',
      function(e) {
        e.dataTransfer.effectAllowed = 'move';
     //   console.log("setting id "+this.id);
        e.dataTransfer.setData('Text', this.id);
        this.classList.add('drag');
        return false;
      },
      false
    );
    
    el.addEventListener(
      'dragend',
      function(e) {
        this.classList.remove('drag');
        return false;
      },
      false
    );
  }
});

/*
el.addEventListener(

'drop',

function(e) {

// Stops some browsers from redirecting.

if (e.stopPropagation) e.stopPropagation();

this.classList.remove('over');

scope.$apply(attrs['drop']);
*/

app.directive('droppable', function() {
  return {
    /*
    scope: {
      drop: '&',
      bin: '='
    },
    */
    link: function(scope, element) {
      // again we need the native object
      var el = element[0];
      
      el.addEventListener(
        'dragover',
        function(e) {
          e.dataTransfer.dropEffect = 'move';
          // allows us to drop
          if (e.preventDefault) e.preventDefault();
          this.classList.add('over');
          return false;
        },
        false
      );
      
      el.addEventListener(
        'dragenter',
        function(e) {
          this.classList.add('over');
          return false;
        },
        false
      );
      
      el.addEventListener(
        'dragleave',
        function(e) {
          this.classList.remove('over');
          return false;
        },
        false
      );
      
      el.addEventListener(
        'drop',
        function(e) {
          // Stops some browsers from redirecting.
          if (e.stopPropagation) e.stopPropagation();
          
          this.classList.remove('over');
          
          //var binId = this.id;
          //console.log(binId);
          var item = document.getElementById(e.dataTransfer.getData('Text'));
          //this.appendChild(item);
          // call the passed drop function
          //console.log($(item)[0]);
          var fn = $(item).attr('drop');
          //console.log(e);
          //var id = e.dataTransfer.getData('Text');
          //scope.$apply(item.attrs['drop']);
          scope.$apply(function(x) {
           // console.log(x);
            var userFnName = fn;
            var userFn = x[userFnName];
            userFn(item, el, e.x, e.y);
          });


          //scope.$apply(attrs['drop']);
          /*
          scope.$apply(function(scope) {
            var fn = scope.drop();
            if ('undefined' !== typeof fn) {            
              fn(item.id, binId);
            }
          });
          */
          
          return false;
        },
        false
      );
    }
  }
});
