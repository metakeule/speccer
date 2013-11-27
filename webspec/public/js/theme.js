$(document).ready(function(){
  
  $.ajax({
    dataType: "json",
    url: window.location.pathname+"/../translations.json",
    async: false,
    success: function(data) {
      // console.log(data);

      function tr(s) {
        if (data[s]) {
          return data[s];
        }
        return s;
      }


      var ul = $('<ul class="nav" />');
      var ulli = $('<li class="dropdown"><a href="#" class="dropdown-toggle" data-toggle="dropdown">' + tr('GOTO') + '<b class="caret"></b></a></li>');
      var drop = $('<ul class="dropdown-menu" />');
      $('<li><a href="#">' + tr('TOP') + '</a></li>').appendTo(drop);
      var i = 0;
      $('h2').each(
        function() {
          var link = $('<a href="#header-'+i+'" />');
          var li = $('<li />');
          link.text($(this).text());
          link.appendTo(li);
          li.appendTo(drop);
          $(this).attr('id', "header-"+i);
          i++;
        }
      );

      drop.appendTo(ulli);
      ulli.appendTo(ul);
      ul.appendTo($('.navbar-inner .container'));

      $('<form class="navbar-form pull-left"><a class="btn btn-default" id="markdown-btn" href="#">'+tr('MARKDOWN')+'</a>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; ' + tr('USAGE') + ' &nbsp;<select id="usageselecter">' + $('#selectusage').html() + '</select>&nbsp; &nbsp; ' + tr('THEME') + ' &nbsp;<select id="themeselecter">' + $('#selecttheme').html() + '</select></form>').appendTo($('.navbar-inner .container'));
      $('body.hide-navbar .navbar').hide();
      function navigate() {
        window.location.href = window.location.pathname+"?theme="+$('#themeselecter').val() + "&usage=" +$('#usageselecter').val();
      }

      $('#themeselecter').change(function() {navigate();});
      $('#usageselecter').change(function() {navigate();});
      $('#markdown-btn').click(function(){
        // console.log("hi");
        // console.log(window.location.pathname+"/../spec.md?usage="+$('#usageselecter').val());
        window.location.href = window.location.pathname+"/../spec.md?usage="+$('#usageselecter').val();
      });
      $('#selecttheme').hide();
      $('#selectusage').hide();
    }
  });
});