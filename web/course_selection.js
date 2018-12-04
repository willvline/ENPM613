
$('.form-course-selection').submit(function(e){
    e.preventDefault();
   
    // $('input:checkbox').click(function () { 
    //     this.blur(); 
    //     this.focus(); 
    //     }); 
 
    var htmleasy = $("#htmleasy").is(':checked');
    var jseasy = $("#jseasy").is(':checked');
    var csseasy = $("#csseasy").is(':checked');
    var htmlmedium = $("#htmlmedium").is(':checked');
    var jsmedium = $("#jsmedium").is(':checked');
    var cssmedium = $("#cssmedium").is(':checked');
    var htmldifficult = $("#htmldifficult").is(':checked');
    var jsdifficult = $("#jsdifficult").is(':checked');
    var cssdifficult = $("#cssdifficult").is(':checked');
    
    var jsonData = {
        course_records: {}
    };

    if (htmleasy) {
        jsonData.course_records["htmleasy"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }
    if (jseasy) {
        jsonData.course_records["jseasy"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }
    if (csseasy) {
        jsonData.course_records["csseasy"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }
    if (htmlmedium) {
        jsonData.course_records["htmlmedium"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }
    if (jsmedium) {
        jsonData.course_records["jsmedium"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }
    if (cssmedium) {
        jsonData.course_records["cssmedium"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }
    if (htmldifficult) {
        jsonData.course_records["htmldifficult"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }
    if (jsdifficult) {
        jsonData.course_records["jsdifficult"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }
    if (cssdifficult) {
        jsonData.course_records["cssdifficult"] = {
            chapter1: false,
            chapter2: false,
            chapter3: false,
            };
    }

    console.log(jsonData);

    var url = $(this).attr('action');
    $.ajax({
        type: "PATCH",
        url: "http://localhost:8000/student",
    
        data: JSON.stringify(jsonData),
        xhrFields: { withCredentials: true },
        crossDomain: true,
        // data: JSON.stringify({
        //     htmlmedium: htmlmedium,
        //     jsmedium: jsmedium,
        //     cssmedium: cssmedium
        // }),
        // data: JSON.stringify({
        //     htmldifficult: htmldifficult,
        //     jsdifficult: jsdifficult,
        //     cssdifficult: cssdifficult
        // }),
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