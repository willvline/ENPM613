$('#update').on('click', function() {
    var email= $("#email").val();
    var user_name= $("#username").val();
    var pass_word= $("#pass_word").val();
    var firstName= $("#firstName").val();
    var lastName= $("#lastName").val();

    $.ajax({
        type: "PATCH",
        url: "http://localhost:8000/student",
        xhrFields: { withCredentials: true },
        crossDomain: true,
        data: JSON.stringify({
            user_name:      user_name,
            email:          email,
            first_name:     firstName,
            last_name:      lastName,
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
            $("#firstName").attr("value", data[0].first_name);
            $("#lastName").attr("value", data[0].last_name);
        },
        error: function (errormessage) {
            console.log(errormessage);
            alert("Username and password don't match!");
        }
    
    });
});
