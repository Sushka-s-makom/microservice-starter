package main

import (
	"log"
	"net/http"
)

func main() {
	// 1) Создаём "маршрутизатор" (роутер) — штуку, которая решает,
	// какой обработчик вызвать для какого URL.
	mux := http.NewServeMux()

	// 2) Регистрируем endpoint /health.
	// Когда прилетит GET /health — выполнится функция healthHandler.
	mux.HandleFunc("/health", healthHandler) // функция healthHandler — обработчик ручки /health.
	mux.HandleFunc("/hello", helloHandler)

	// 3) Задаём адрес, на котором будет слушать сервер.
	// ":8080" означает "на всех интерфейсах, порт 8080".
	addr := ":8080"

	log.Printf("listening on %s", addr)

	// 4) Запуск сервера.
	// ListenAndServe блокирует main() и обслуживает запросы, пока сервис не остановят
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
