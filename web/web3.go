package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, Redseanet!")
}
func main() {
	http.HandleFunc("/", sayHello)
	log.Println("启动成功，可以通过 localhost:8000 访问")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("List 8000")
	}
}
