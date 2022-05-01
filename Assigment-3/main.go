package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type OtherPage struct {
	Title string
	News  string
}

func otherHandler(w http.ResponseWriter, r *http.Request) {
	Data()
	p := OtherPage{Title: "Auto Reload DATA Projects", News: "update every 15 second"}
	t, err := template.ParseFiles("data.html")
	fmt.Println(err)
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>WELCOME</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/data/", otherHandler)
	http.ListenAndServe(":8080", nil)
}

func Data() {
	rand.Seed(time.Now().Unix())
	Water := rand.Intn(15)
	Wind := rand.Intn(15)
	fmt.Println("Status")
	fmt.Println("Water  : ", Water, "Meter")
	fmt.Println("Wind   : ", Wind, "Meter Per Second")

	fmt.Scanln(Water)
	if Water <= 5 {
		fmt.Println("Status Water : Aman!")
	}
	if Water >= 6 && Water <= 8 {
		fmt.Println("Status Water : Siaga!")
	}
	if Water >= 8 {
		fmt.Println("Status Water : Bahaya!")
	}

	fmt.Scanln(Wind)
	if Wind <= 6 {
		fmt.Println("Status Wind : Aman!")
	}
	if Wind >= 7 && Wind <= 15 {
		fmt.Println("Status Wind : Siaga!")
	}
	if Wind > 15 {
		fmt.Println("Status Wind : Bahaya!")
	}
}
