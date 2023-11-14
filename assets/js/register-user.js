$('#register-user').on('submit', createUser)

function createUser(event) {
    event.preventDefault();

    let name = $('#name').val();
    let email = $('#email').val();
    let nick = $('#nick').val();
    let password = $('#password').val();
    let passwordConfirmation = $('#password-confirmation').val();

    if (password !== passwordConfirmation) {
        Swal.fire("Ops...", "As senhas não coincidem.", "error");
        return
    }

    $.ajax({
        url: "/usuario",
        method: "POST",
        data: {
            name: name,
            email: email,
            nick: nick,
            password: password,
        }
    }).done(function () {
        Swal.fire("Sucesso!", "Usuário criado com sucesso.", "success").then(function () {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: email,
                    password: password,
                }
            }).done(function () {
                window.location = "/home";
            }).fail(function () {
                Swal.fire("Ops...", "Erro ao autenticar o usuário.", "error");
            });
        });
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao cadastrar usuário.", "error");
    });
}