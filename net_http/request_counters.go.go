package main

import (
    "fmt"
    "net/http"
)

var count = 0
var pizzaCount = 0

func main(){

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        count++
        fmt.Fprint(w, count)
    })

    http.HandleFunc("/pizza", func(w http.ResponseWriter, r *http.Request){
        pizzaCount++
        fmt.Fprint(w, pizzaCount)
    })
    http.ListenAndServe(":8080", nil)
}