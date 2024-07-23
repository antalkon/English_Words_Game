function suesses_v(message) {
    const alertBox = document.getElementById('alertBox_su');
    alertBox.classList.remove('hidden');
    document.getElementById('suText').textContent = message;
    
    setTimeout(() => {
      alertBox.classList.add('hidden');
    }, 3000); // 3000 milliseconds = 3 seconds
  }

  function error_v(message) {
    const alertBox = document.getElementById('alertBox_er');
    alertBox.classList.remove('hidden');
    document.getElementById('erText').textContent = message;
    
    setTimeout(() => {
      alertBox.classList.add('hidden');
    }, 3000); // 3000 milliseconds = 3 seconds
  }
  
  function warning_v(message) {
    const alertBox = document.getElementById('alertBox_wa');
    alertBox.classList.remove('hidden');
    document.getElementById('waText').textContent = message;
    
    setTimeout(() => {
      alertBox.classList.add('hidden');
    }, 3000); // 3000 milliseconds = 3 seconds
  }
  