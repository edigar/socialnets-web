$('#login').on('submit', login)

function login(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function () {
        window.location = "/home"
    }).fail(function (jqXHR, textStatus, errorThrown) {
        //console.log("XHR: ", jqXHR, "textStatus: ", textStatus, "Error: ", errorThrown);
        Swal.fire("Ops...", "Usuário ou senha incorretos", "error");
    });
}