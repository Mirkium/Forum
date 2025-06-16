const container = document.getElementById('form-container');
const toRegisterBtn = document.getElementById('to-register');
const toLoginBtn = document.getElementById('to-login');
const registerForm = document.getElementById('registerForm');

toRegisterBtn.addEventListener('click', () => {
    container.classList.add('form-active');
});

toLoginBtn.addEventListener('click', () => {
    container.classList.remove('form-active');
    registerForm.reset();
})

