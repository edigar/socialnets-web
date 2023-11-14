$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser);

function unfollow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuario/${userId}/parar-de-seguir`,
        method: "POST",
    }).done(function () {
        window.location = `/usuario/${userId}`;
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário.", "error");
        $('#unfollow').prop('disabled', false);
    })
}

function follow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuario/${userId}/seguir`,
        method: "POST",
    }).done(function () {
        window.location = `/usuario/${userId}`;
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao seguir o usuário.", "error");
        $('#follow').prop('disabled', false);
    })
}

function edit(event) {
    event.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success").then(function() {
            window.location = "/perfil";
        })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error")
    });
}

function updatePassword(event) {
    event.preventDefault();

    const currentPassword = $('#password').val();
    const newPassword = $('#new-password').val();
    const passwordConfirmation = $('#password-confirmation').val();

    if(newPassword !== passwordConfirmation) {
        Swal.fire("Ops...", "As senhas não coincidem!", "warning");
        return;
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            current: currentPassword,
            new: newPassword
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Senha alterada com sucesso.", "success").then(function () {
            window.location = "/perfil";
        })
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao atualizar a senha!", "error")
    });
}

function deleteUser() {
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja apagar sua conta? essa é uma ação irreversível.",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning",
    }).then(function(confirmation) {
        if (confirmation.value) {
            $.ajax({
                url: "/deletar-usuario",
                method: "DELETE",
            }).done(function(){
                Swal.fire("Sucesso!", "Seu usuário foi exluido com sucesso.", "success").then(function () {
                    window.location = "/logout";
                })
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao exlcuir o seu usuário", "error")
            })
        }
    });
}