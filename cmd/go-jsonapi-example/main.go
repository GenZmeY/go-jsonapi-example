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

	"context"
	"os"
	"os/signal"
	"syscall"

	"go-jsonapi-example/internal/model"
	"go-jsonapi-example/internal/resource"
	"go-jsonapi-example/internal/storage"

	"github.com/manyminds/api2go"

	"github.com/juju/gnuflag"
)

func main() {
	host := gnuflag.String("host", "localhost", "host")
	port := gnuflag.Int("port", 8080, "port")
	gnuflag.Parse(true)

	addr := fmt.Sprintf("%s:%d", *host, *port)
	baseURL := fmt.Sprintf("http://%s", addr)

	api := api2go.NewAPIWithBaseURL("v1", baseURL)
	api.AddResource(model.Car{}, resource.CarResource{CarStorage: storage.NewCarStorage()})

	server := &http.Server{Addr: addr, Handler: api.Handler()}
	closeHandler(server)

	fmt.Printf("Listening on %s\n", addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Println(err)
		os.Exit(0)
	}
}

func closeHandler(server *http.Server) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}()
}
