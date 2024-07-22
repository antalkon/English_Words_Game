function suesses_v(message) {
    const alertBox = document.getElementById('alertBox_su');
    alertBox.classList.remove('hidden');
    document.getElementById('suText').textContent = message;
    
    setTimeout(() => {
      alertBox.classList.add('hidden');
    }, 3000); // 3000 milliseconds = 3 seconds
  }
  