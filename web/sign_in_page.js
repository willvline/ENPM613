//alert('js connected');
// $('.form-signin').submit(function(e){
//       e.preventDefault();
//       var email = $('#inputEmail').val();
//       var password = $('#inputPassword').val();
//       var url = $(this).attr('action');
//       $.post(url, {email: email, password: password}).
//       done(function(data){
//         console.log(data);
//       });
//     });

$('.form-signin').submit(function(e){
  e.preventDefault();
  var email = $('#inputEmail').val();
  var password = $('#inputPassword').val();
  var url = $(this).attr('action');
  $.ajax({
    type: "POST",
    url: url+"/auth",
    data: JSON.stringify({
        user_name:      email,
        email:          email,
        pass_word:      password
    }),
    contentType: "application/json; charset=utf-8", // this
    dataType: "json", // and this
    success: function (data) {
        console.log(data);
        if (data) {
             window.location.href="http://localhost:8080/dashboard.html"
        }else {
            alert("Username and password don't match!");
        }
    },
    error: function (errormessage) {
        console.log(errormessage);
    }
});
});