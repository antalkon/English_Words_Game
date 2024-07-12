// Получение значения CookieAccept из локального хранилища

// Функция для сохранения значения CookieAccept = true в локальном хранилище и удаления модального окна
function setCookieAccept() {
    localStorage.setItem('CookieAccept', 'true');
    if (cookieModal) {
        cookieModal.remove();
    }
}

// Проверка значения CookieAccept и удаление модального окна, если значение 'true'

