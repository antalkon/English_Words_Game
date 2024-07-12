const cookieWindow = document.getElementById('cookie');
const cookieAccept = localStorage.getItem('CookieAccept');
const cookieModal = document.getElementById('cookieModal');

if (cookieAccept === 'true' && cookieModal) {
    cookieWindow.remove();
} else{
    cookieWindow.style.display = 'block'
}