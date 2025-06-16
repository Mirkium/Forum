function seeComments(el) {
    const comments = el.nextElementSibling;
    if (comments && comments.classList.contains('comments')) {
        comments.style.display = (comments.style.display === "none" ||
                                   comments.style.display === '') 
                                   ? "block" : "none";
    }
}