//самое простое и базовое веб приложение. Простой Веб сервер. HTTP сервер
//после запуска в браузере набрать: http://localhost:8081/
//не забываем останавливать (кнопка стоп справа от плей)

package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world! go_simple_web_app\n"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
