//Метод выполняется при полной загрузке DOM, когда мы первый раз заходим на сайт
document.addEventListener("DOMContentLoaded", function() {
	//Создаем AJAX-запрос GET на страницу /index.json
	var xhr = new XMLHttpRequest();
	//Метод выполняется при изменении состояния запроса
	xhr.onreadystatechange = function() {
		//Если положительный ответ
		if(xhr.readyState == 4 && xhr.status == 200) {
			//Получаем ответ от сервера и парсим его в JSON-объект
			var jsonResponse = JSON.parse(xhr.responseText) ;
			//Ищем элемент с id=welcome и вставляем  в тело секцию welcome из результата
			document.getElementById("welcome").innerHTML = jsonResponse.welcome;
			document.getElementById("description").innerHTML = jsonResponse.description;

			//Выводим сообщение с текстом результата от сервера
			alert(JSON.stringify(jsonResponse));
		}
	};
	xhr.open("GET", "/index.json", true);
	//Отправляем AJAX-запрос
	xhr.send();
});