package main

import (
	"encoding/json"
	"net/http"
)

// healthHandler — обработчик ручки /health.
// Его задача: быстро и стабильно ответить "я жив".
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем метод (GET),
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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 1) Разрешаем только GET — так аккуратнее.
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 2) Читаем query параметр name.
	// Пример: /hello?name=alex
	name := r.URL.Query().Get("name")

	// 3) Если name не передали — используем значение по умолчанию.
	if name == "" {
		name = "world"
	}

	// 4) Формируем JSON-ответ.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "hello, " + name,
	})

}
