$('#help').on('click', function() {
    $("#screen").load("Help.html");
});

$('#Dashboard').on('click', function() {
    $("#screen").load("dashboard.html");
});

$('#Myaccount').on('click', function() {
    $("#screen").load("accountpage.html");
    $.ajax({
        type: "GET",
        url: "http://localhost:8000/student",
        xhrFields: { withCredentials: true },
        crossDomain: true,
        // data: JSON.stringify({
        //     user_name:      email,
        //     pass_word:      password
        // }),
        // contentType: "application/json; charset=utf-8", // this
        // dataType: "json", // and this
        success: function (data, xhr) {
            //返回一个 json object list
            //[{student_id: "5c05dd666d39b5139927a6ff", user_name: "terp@umd.edu", pass_word: "jiushiwo", email: "terp@umd.edu", grades: {…}, …]
            console.log(data[0]);
            //window.location.href="http://localhost:8080/dashboard.html"
            $("#username").attr("value", data[0].user_name);
            $("#email").attr("value", data[0].email);
            $("#firstName").attr("value", data[0].first_name);
            $("#lastName").attr("value", data[0].last_name);
            
        },
        error: function (errormessage) {
            console.log(errormessage);
            alert("Username and password don't match!");
        }
    
    });
});

$('#DiscussionBoard').on('click', function() {
    $("#screen").load("DiscussionBoard.html");
    $.ajax({
        type: "GET",
        url: "http://localhost:8000/discussionboard",
        // data: JSON.stringify({
        //     user_name:    email, 
        //     email:        email,
        //     pass_word:    password
        // }),
        // contentType: "application/json; charset=utf-8", // this
        dataType: "json", // and this
        success: function (data) {
            //window.location.href="http://localhost:8080/sign_in_page.html";
        },
        error: function (errormessage) {
            //alert(errormessage.JSON());
            console.log(errormessage);
            alert(errormessage.responseJSON.error);
            
            console.log(errormessage);
    }}).done(function(data){
        console.log(data);
        var i;
        var posts = "";
        for (i = 0; i < data.length; i++) { 
            posts = posts + '<strong>'+data[i].poster_name+'</strong>' + '<br><strong>'+data[i].post_date+'</strong>' + '<p><em>'+data[i].content+'</em></p >';
            // $("#box").append('<strong>'+data[i].poster_name+'</strong>' + '<br><strong>'+data[i].post_date+'</strong>' + '<p><em>'+data[i].content+'</em></p >')
        }
        $("#box").html(posts);  
    });
});