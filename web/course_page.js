//alert('course page js connected');
$(function() {
    $("#sidebar").load("sidebar.html");
    $("#screen").load("introduction.html");
});

$('#introduction').on('click', function() {
    $("li a").removeClass("active");
    $(this).addClass("active");
    $("#screen").load("introduction.html");
});

$('#lecture').on('click', function() {
    $("li a").removeClass("active");
    $(this).addClass("active");
    $("#screen").load("lecture_page.html");
});

$('#assignment').on('click', function() {
    $("li a").removeClass("active");
    $(this).addClass("active");
    $("#screen").load("assignment.html");
});

$('#grades').on('click', function() {
    $("li a").removeClass("active");
    $(this).addClass("active");
    $("#screen").load("grades.html");
});