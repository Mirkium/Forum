function toggleComments(threadId) {
    const section = document.getElementById('comments-' + threadId);
    if (section.style.display === 'none') {
        section.style.display = 'block';
    } else {
        section.style.display = 'none';
    }
}

function showCommentForm(threadId) {
    const form = document.getElementById('comment-form-' + threadId);
    form.style.display = 'block';
}