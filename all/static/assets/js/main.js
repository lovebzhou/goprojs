$(document).ready(function(){
    console.log("main.js on document ready");
    foo();
  });

  function foo() {
    $("#console").html("<p>hello world2!</p>");
    $("#log").text("hello");
  }

  function dummy() {
    var cat = {name: "Tom",
      walk: function () {
        console.log(this.name + ' is walking');
      }
    };

    var dog = Object.create(cat);

    console.log(dog.name);
    dog.walk();
  }