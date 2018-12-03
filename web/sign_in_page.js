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
    type: "GET",
    url: url + "?user_name="+email,
    data: JSON.stringify({user_name: email, pass_word: password}),
    contentType: "application/json; charset=utf-8", // this
    dataType: "json", // and this
    success: function (data) {
        console.log(data);
        if (data.length >= 1) {
            console.log(data.length);
            window.location.href="http://localhost:8080/dashboard.html"
        }
    },
    error: function (errormessage) {
        console.log(errormessage);
    }
});
});