
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
        url: "http:localhost:8000/signup",
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