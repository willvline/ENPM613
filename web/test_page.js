

$('.form-signin').submit(function(e){
  var email = $('#inputEmail').val();
  var password = $('#inputPassword').val();
  var course = $('#htmlesay').val();



  $.ajax({
    type: "GET",
    url: "http://localhost:8000/student",
    xhrFields: { withCredentials: true },
    crossDomain: true,
    data: JSON.stringify({
        user_name:      email,
        pass_word:      password
    }),
    contentType: "application/json; charset=utf-8", // this
    dataType: "json", // and this
    success: function (data, xhr) {
        //返回一个 json object list
        //[{student_id: "5c05dd666d39b5139927a6ff", user_name: "terp@umd.edu", pass_word: "jiushiwo", email: "terp@umd.edu", grades: {…}, …]
        console.log(data[0].student_id);
        //window.location.href="http://localhost:8080/dashboard.html"
        $("#username").attr("value", data[0].user_name);
        $("#email").attr("value", data[0].email);
    },
    error: function (errormessage) {
        console.log(errormessage);
        alert("Username and password don't match!");
    }

});
});