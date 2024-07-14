// Функция для обработки ошибок с текстом сообщения
function handleError(errorMsg) {
    console.error("Error:", errorMsg);
    // Здесь можно добавить дополнительные действия по обработке ошибок
}

// Функция для обработки успешного ответа с текстом сообщения
function handleSuccess(successMsg) {
    console.log("Success:", successMsg);
    // Здесь можно добавить дополнительные действия по успешному ответу
}

// Основная функция для выполнения POST запроса
async function postData(url = '', data = {}) {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data) // Преобразование объекта data в JSON строку
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const jsonResponse = await response.json();

        if (jsonResponse.error) {
            handleError(jsonResponse.error);
        } else if (jsonResponse.success) {
            handleSuccess(jsonResponse.success);
        } else {
            handleError('Unknown response structure');
        }
    } catch (error) {
        handleError(`Fetch error: ${error.message}`);
    }
}

// Пример данных для отправки

