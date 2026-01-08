package main

import (
	"encoding/json"
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

// healthHandler — обработчик ручки /health.
// Его задача: быстро и стабильно ответить "я жив".
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// В реальном сервисе иногда проверяют метод (GET),
	// чтобы случайные POST/PUT не проходили.
	if r.Method != http.MethodGet {
		// 405 Method Not Allowed — корректный ответ, если метод не тот.
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Укажем, что отдаём JSON.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Явно задаём статус 200 OK (хотя по умолчанию он и так 200, если ничего не менять).
	w.WriteHeader(http.StatusOK)

	// Отдаём JSON: {"status":"ok"}.
	// json.NewEncoder удобен: кодирует и пишет в w.
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
