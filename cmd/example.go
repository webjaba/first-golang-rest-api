// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(
// 		w,
// 		"Welcome to the Home Page! Request type: %s Protocol: %s",
// 		r.Method,
// 		r.Proto,
// 	)
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "About Us")
// }

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	http.HandleFunc("/about", aboutHandler)

// 	fmt.Println("Server is running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }
