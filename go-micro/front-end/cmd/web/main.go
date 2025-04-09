package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	// НОВИЙ МАРШРУТ ЕБАУТ
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		render(w, "about.page.gohtml")
	})

	// МАРШРУТ ФОРМИ
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		render(w, "form.page.gohtml")
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/form", http.StatusSeeOther)
			return
		}
		r.ParseForm()
		pes_amount := r.FormValue("pes_amount")
		fmt.Fprintf(w, "Ого, у вас %s песів, це круто!", pes_amount)
	})

	fmt.Println("Running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func render(w http.ResponseWriter, t string) {
	fmt.Println("Rendering:", t)
	files := []string{
		"cmd/web/templates/base.layout.gohtml",
		"cmd/web/templates/header.partial.gohtml",
		"cmd/web/templates/footer.partial.gohtml",
		fmt.Sprintf("cmd/web/templates/%s", t),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Parse error:", err)
		http.Error(w, "Template error", 500)
		return
	}
	data := map[string]interface{}{
		"CurrentTime": time.Now().Format("02 Jan 2006 15:04:05"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Execute error:", err)
	}
}
