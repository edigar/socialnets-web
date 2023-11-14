$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.unlike-post', unlikePost);
$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost)

function createPost(event) {
    event.preventDefault();

    $.ajax({
        url: "/post",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function (){
        Swal.fire("Ops...", "Erro ao criar publicação.", "error");
    });
}

function likePost(event) {
    event.preventDefault();

    const postId = event.target.closest('div').getAttribute('data-post-id');
    event.target.disabled = true;

    $.ajax({
        url: `/post/${postId}/like`,
        method: "POST",
    }).done(function () {
        const likesCount = $(event.target).next('span');
        const likes = parseInt(likesCount.text());
        likesCount.text(likes + 1)

        $(event.target).addClass('unlike-post');
        $(event.target).addClass('text-danger');
        $(event.target).removeClass('like-post');
    }).fail(function (){
        Swal.fire("Ops...", "Erro ao curtir publicação.", "error");
    }).always(function () {
        event.target.disabled = false;
    });
}

function unlikePost(event) {
    event.preventDefault();

    const postId = event.target.closest('div').getAttribute('data-post-id');
    event.target.disabled = true;

    $.ajax({
        url: `/post/${postId}/unlike`,
        method: "POST",
    }).done(function () {
        const likesCount = $(event.target).next('span');
        const likes = parseInt(likesCount.text());
        likesCount.text(likes - 1)

        $(event.target).removeClass('unlike-post');
        $(event.target).removeClass('text-danger');
        $(event.target).addClass('like-post');
    }).fail(function (){
        Swal.fire("Ops...", "Erro ao descurtir publicação.", "error");
    }).always(function () {
        event.target.disabled = false;
    });
}

function updatePost() {
    $(this).prop('disabled', true);
    const postId = $(this).data('post-id');

    $.ajax({
        url: `/post/${postId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function () {
        Swal.fire('Sucesso!', 'Publicação criada com sucesso.', 'success').then(function() {
            window.location="/home";
        });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao editar publicação.", "error");
    }).always(function() {
        $('#update-post').prop('disabled', false);
    })
}

function deletePost(event) {
    event.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicção? Essa ação é irreversível.",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning",
    }).then(function(confirmation) {
        if (!confirmation.value) return;
        const element = $(event.target);
        const post = element.closest('div');
        const postId = post.data('post-id');
        const box = post.parent().closest('div');

        element.prop('disabled', true);

        $.ajax({
            url: `/post/${postId}`,
            method: "DELETE",
        }).done(function () {
            box.fadeOut("slow", function () {
                $(this).remove();
            })
        }).fail(function (){
            Swal.fire("Ops...", "Erro ao excluir a publicação.", "error");
        })
    })
}