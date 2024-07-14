function pendingEmail(){
    let email = document.getElementById("email")
    const data = {
        email: email.value,
    };
    
    // Вызов функции с примером URL
    postData('/api/v1/email/pending', data);
}