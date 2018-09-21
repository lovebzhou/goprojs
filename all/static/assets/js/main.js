$(document).ready(function(){
    console.log("main.js on document ready");
    foo();
  });

  function foo() {
    $("#console").html("<p>hello world2!</p>");
    $("#log").text("hello")
  }