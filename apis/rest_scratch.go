package main

import "net/http"

func server() {
	http.ListenAndServe(":8080")
}
