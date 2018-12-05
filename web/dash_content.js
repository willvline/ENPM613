$('#addCourse').on('click', function () {

    $.ajax({
        type: "GET",
        url: "http://localhost:8000/student",
        xhrFields: { withCredentials: true },
        crossDomain: true,
        success: function (data, xhr) {
            //返回一个 json object list
            //[{student_id: "5c05dd666d39b5139927a6ff", user_name: "terp@umd.edu", pass_word: "jiushiwo", email: "terp@umd.edu", grades: {…}, …]
            console.log(data[0]);
            //window.location.href="http://localhost:8080/dashboard.html"
            var courses = data[0].course_records;
            for (var key in courses) {
                switch (key) {
                    case "jseasy":
                        $("#jseasy_img").css("display", "inline");
                        break;
                    case "jsmedium":
                        $("#jsmedium_img").css("display", "inline");
                        break;
                    case "jsdifficult":
                        $("#jsdifficult_img").css("display", "inline");
                        break;
                    case "htmleasy":
                        $("#htmleasy_img").css("display", "inline");
                        break;
                    case "htmlmedium":
                        $("#htmlmedium_img").css("display", "inline");
                        break;
                    case "htmldifficult":
                        $("#htmldifficult_img").css("display", "inline");
                        break;
                    case "csseasy":
                        $("#csseasy_img").css("display", "inline");
                        break;
                    case "cssmedium":
                        $("#cssmedium_img").css("display", "inline");
                        break;
                    case ("cssdifficult"):
                        $("#cssdifficult_img").css("display", "inline");
                        break;
                }
            }
        },
        error: function (errormessage) {
            console.log(errormessage);
            alert("Username and password don't match!");
        }

    });

    $('#screen').load("lecture_directory_page.html");

});
$(function () {
    $.ajax({
        type: "GET",
        url: "http://localhost:8000/student",
        xhrFields: { withCredentials: true },
        crossDomain: true,
        success: function (data, xhr) {
            //resturn a json object list
            //[{student_id: "5c05dd666d39b5139927a6ff", user_name: "terp@umd.edu", pass_word: "jiushiwo", email: "terp@umd.edu", grades: {…}, …]
            console.log(data[0]);
            //window.location.href="http://localhost:8080/dashboard.html"
            var courses = data[0].course_records;
            for (var key in courses) {
                switch (key) {
                    case "jseasy":
                        $("#js").css("display", "inline");
                        break;
                    case "jsmedium":
                        $("#js").css("display", "inline");
                        break;
                    case "jsdifficult":
                        $("#js").css("display", "inline");
                        break;
                    case "htmleasy":
                        $("#html").css("display", "inline");
                        break;
                    case "htmlmedium":
                        $("#html").css("display", "inline");
                        break;
                    case "htmldifficult":
                        $("#html").css("display", "inline");
                        break;
                    case "csseasy":
                        $("#css").css("display", "inline");
                        break;
                    case "cssmedium":
                        $("#css").css("display", "inline");
                        break;
                    case ("cssdifficult"):
                        $("#css").css("display", "inline");
                        break;
                }
            }
        },
        error: function (errormessage) {
            console.log(errormessage);
            alert("Username and password don't match!");
        }

    });
});