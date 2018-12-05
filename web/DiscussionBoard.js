$('#submit').on('click', function(e) {
    var user_name = $("#user_name").val();
    var content = $("#content").val();
    var time = new Date($.now());
    var post_date = "" + time.getMonth()+'/'+time.getDate()+'/'+time.getFullYear() + ' ' + time.getHours() + ':' + time.getMinutes() + ':' + time.getSeconds();
    $.ajax({
        type: "POST",
        url: "http://localhost:8000/discussionboard",
        data: JSON.stringify({
            poster_name:    user_name, 
            content:        content,
            post_date:      post_date
        }),
        contentType: "application/json; charset=utf-8", // this
        dataType: "json", // and this
        success: function (data) {
            
        },
        error: function (errormessage) {
            //alert(errormessage.JSON());
            alert(errormessage.responseJSON.error);
            
            console.log(errormessage);
    }}).done(function(data){
        $("#box").prepend('<hr>' + '<strong>'+data[0].poster_name+'</strong>' + '<br>'+data[0].post_date + '<p><em>'+data[0].content+'</em></p>')
    });
});
