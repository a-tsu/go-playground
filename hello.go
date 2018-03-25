package main

import (
	"fmt"
    "net/http"
    "log"
)

var str string = "default"

func setStr(str1 string, str2 string) {
    str = str1 + " " + str2
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, str)
}

func main() {
    setStr("hello", "world")
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8089", nil))
}
