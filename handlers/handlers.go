package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/AnjuAshokPA/web-calculator/operations"
)

// Create handler loads a calculator webpage with title as Hello World
func Create(w http.ResponseWriter, r *http.Request) {
	htmlFileData, err := template.ParseFiles("./views/page.html")
	tmpl := template.Must(htmlFileData, err)
	tmpl.Execute(w, tmpl)
}

// Calculator handler performs arithematic operations on the number given by the user
func Calculator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error while parsing form:%v", err)
		return
	}
	expr := r.FormValue("expr")
	log.Println(expr)
	finalValue, err := operations.Operations(expr)
	if err != nil {
		log.Println("Error while dividing:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("Operation is fatal: %v", err))
		return
	}
	log.Println("The final value is:", finalValue)
	w.WriteHeader(http.StatusOK)
	//log.Fatal(err1)
	json.NewEncoder(w).Encode(fmt.Sprintf("%d", finalValue))
}
