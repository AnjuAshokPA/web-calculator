package routers

import (
	"net/http"

	"github.com/AnjuAshokPA/web-calculator/handlers"
)

// Routers function is for hosting the webserver
func Routers() error {
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("./views"))))
	http.HandleFunc("/", handlers.Create)
	http.HandleFunc("/calculator", handlers.Calculator)
	err := http.ListenAndServe(":8081", nil)
	return err
}
