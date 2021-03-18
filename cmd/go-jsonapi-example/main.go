/*
Разработать CRUD (REST API) для модели автомобиля, который имеет следующие поля:

1. Уникальный идентификатор (любой тип, общение с БД не является критерием чего-либо, можно сделать и in-memory хранилище на время жизни сервиса)
2. Бренд автомобиля (текст)
3. Модель автомобиля (текст)
4. Цена автомобиля (целое, не может быть меньше 0)
5. Статус автомобиля (В пути, На складе, Продан, Снят с продажи)
6. Пробег (целое)

Формат ответа api - json api (https://jsonapi.org/)
*/

package main

import (
	"fmt"
	"net/http"

	"go-jsonapi-example/internal/model"
	"go-jsonapi-example/internal/resource"
	"go-jsonapi-example/internal/storage"

	"github.com/manyminds/api2go"
)

func main() {
	port := 8080
	baseURL := fmt.Sprintf("http://localhost:%d", port)
	api := api2go.NewAPIWithBaseURL("v1", baseURL)
	carStorage := storage.NewCarStorage()
	api.AddResource(model.Car{}, resource.CarResource{CarStorage: carStorage})

	fmt.Printf("Listening on :%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), api.Handler())
}
