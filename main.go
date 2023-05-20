package main

import (
	"ascii-art-web/asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) { // с помощью  "w" обращаемся к конкретной страничке, "r" - отслеживает подключение к конкретной странице (POST)
	tmpl, err := template.ParseFiles("ui/templates/index.html") // создаём переменную с шаблонами html
	if err != nil {
		// fmt.Fprintf(w, err.Error())
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		// fmt.Fprintf(w, err.Error())
		return
	}
}

func error_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ERROR, BIATCH")
}

type ViewData struct {
	Output string
}

func get_article(w http.ResponseWriter, r *http.Request) {
	output_title := r.FormValue("output_title")
	if output_title == "" {
		fmt.Fprintf(w, "The input field is empty, please don't be shy, write something :3")
		return
	}
	style := r.FormValue("ascii-style")
	output, err := asciiart.AsciiArt(output_title, style)
	if err != nil {
		// fmt.Fprintf(w, "Something wrong...")
		return
	}

	data := ViewData{
		Output: output,
	}

	tmpl, err := template.ParseFiles("ui/templates/index.html")
	if err != nil {
		// fmt.Fprintf(w, err.Error())
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		// fmt.Fprintf(w, err.Error())
		return
	}
}

func post_article(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "POST" {
		title := r.FormValue("title")
		style := r.FormValue("ascii-style")
		output, err := asciiart.AsciiArt(title, style)
		if err != nil {
			fmt.Fprintf(w, "Something wrong...")
			return
		}

		data := ViewData{
			Output: output,
		}

		tmpl, err := template.ParseFiles("ui/templates/index.html")
		if err != nil {
			// fmt.Fprintf(w, err.Error())
			return
		}
		if title == "" {
			fmt.Fprintf(w, "The input field is empty, please don't be shy, write something :3")
			return
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			// fmt.Fprintf(w, err.Error())
			return
		}
	} else if r.Method == "GET" {
		get_article(w, r)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", index)                                                    // главная страничка
	http.HandleFunc("/error/", error_page)                                         // страничка ошибки на будущее
	http.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui")))) // каждый раз когда идёт обращение нач-ся на "ui" ui убирается по оставшемуся пути ищется соответствующий файл
	http.HandleFunc("/save_article", post_article)
	http.HandleFunc("/get_article/", get_article)
	http.ListenAndServe(":8181", nil)
}

// fmt.Println("Server is listening...http://127.0.0.1:8181/")
// err := http.ListenAndServe(":8181", nil) // устанавливаем порт веб-сервера
// if err != nil {
// 	log.Fatal("ListenAndServe: ", err)
// }
// http.ListenAndServe(":8181", http.FileServer(http.Dir("static")))
// // http.HandleFunc("/", какая-то_функция)
