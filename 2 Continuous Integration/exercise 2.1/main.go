package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "Hello, it is %d:%d", t.Hour(), getMinute(t.Minute(), t.Second()))

	Fibonacci(1000)
}

func getMinute(minute int, second int) int {
	return minute + second/30
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
