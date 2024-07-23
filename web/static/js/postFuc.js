async function sendPostRequest(data, url) {
    try {
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      });
  
      if (!response.ok) {
        const errorData = await response.json();
        error_v(`${response.status} - ${errorData.message}`)
        throw new Error(`Error: ${response.status} - ${errorData.message}`);
      }
  
      const result = await response.json();
      console.log('Success:', result.message);
      suesses_v(result.message)
    } catch (error) {
      console.error('Request failed:', error.message);
      error_v(error.message)

    }
  }
  
