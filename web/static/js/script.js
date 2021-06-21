$(document).ready(function () {
    $('button.clearBtn').click(function (event) {
        event.preventDefault();
        $(this).closest('form').find("input[type=text]").val("");
        $(this).closest('form').find("input").removeClass("is-invalid");
    });
});