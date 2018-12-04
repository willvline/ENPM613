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

// $('.form-signup').submit(function(e){
//   e.preventDefault();
//   var firstName = $('#firstName').val();
//   var lastName  = $('#lastName').val();
//   var userName  = $('#userName').val();
//   var email     = $('#email').val();
//   var password = $('#password').val();
//   var url = $(this).attr('action');
 
//   $.ajax({
//     type: "POST",
//     url: url,
//     data: JSON.stringify({
//         user_name:  firstName+" "+lastName, 
//         email:      email,
//         pass_word:  password
//     }),
//     contentType: "application/json; charset=utf-8", // this
//     dataType: "json", // and this
//     success: function (data) {
//         console.log(data);
//         //window.location.href="http://localhost:8080/sign_in_page.html"
//     },
//     error: function (errormessage) {
//         console.log(errormessage);
//     }
// });
// });
$('.form-signup').submit(function(e){
    e.preventDefault();
    var firstName = $('#firstName').val();
    var lastName  = $('#lastName').val();
    var userName  = $('#userName').val();
    var email     = $('#email').val();
    var password = $('#password').val();
    var passwordAgain = $('#passwordAgain').val();
    if(password !== passwordAgain){
        alert("passwords does not match");
        return false;
        //window.location.href="/sign_up_page.html";
    }
    var url = $(this).attr('action');
    $.ajax({
        type: "POST",
        url: url + "/signup",
        data: JSON.stringify({
            user_name:    email, 
            email:        email,
            pass_word:    password
        }),
        contentType: "application/json; charset=utf-8", // this
        dataType: "json", // and this
        success: function (data) {
            console.log(data);
            if(data.length >= 1){
            window.location.href="http://localhost:8080/sign_in_page.html";
            }
        },
        error: function (errormessage) {
            //alert(errormessage.JSON());
            alert(errormessage.responseJSON.error);
            
            console.log(errormessage);
    }});

  });