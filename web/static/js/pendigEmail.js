function sendPendingEmail(){
    const email = document.getElementById("email")
    const data = {
        email: email.value,
      };
      
      const url = '/api/v1/email/pending';
      
      sendPostRequest(data, url);
      
}
  