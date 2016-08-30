package controllers

import (
	"github.com/astaxie/beego"
)

//Структура контроллера
type MainController struct {
	beego.Controller
}

//Структура ответа клиенту
type SimpleResponse struct {
	Welcome string `json:"welcome"`
	Description string `json:"description"`
}

func (c *MainController) Get() {
	//Если поступил простой get-запрос (Когда мы заходим на сайт)
	if c.Ctx.Input.Param(":ext") == "" {
		//Указываем шаблон 
		c.TplName = "index.html"
		//Рендерим шаблон и отправляем клиенту
		c.Render()
	//Если мы запросили страницу с расширением json (в нашем случае /index.json)	
	} else if c.Ctx.Input.Param(":ext") == "json" {
		//Создаем объект структуры SimpleResponse, заполняем поля объекта данными
		simpleResponse := SimpleResponse{Welcome: "Welcome my Friend!", Description: "This is very simple site on beego."}
		//Передаем объект в секцию JSON контроллера
		c.Data["json"] = simpleResponse
		//Возвращаем JSON-ответ клиенту
		c.ServeJSON()
	}
}
