package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

func inputvalue() {
	var name string
	fmt.Printf("ป้อนชื่อ :")
	fmt.Scanf("%s", &name)

	fmt.Println("Hello =", name)
}
