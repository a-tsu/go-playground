package main

import (
    "fmt"
    "net/http"
    "log"
    "time"
)

var str string = "default"

func setStr(str1 string, str2 string) {
    str = str1 + " " + str2
}

var flag = make(chan bool)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("access came\n")
    flag <- true
    fmt.Fprintf(w, str)
}

func call1sec(ch chan bool) {
    t := time.NewTicker(time.Second)
    for {
        select {
            case <-t.C:
                fmt.Printf("1sec passed.\n")
            case <-ch:
                fmt.Printf("ch has come\n")
        }
    }
    defer t.Stop()
}

func main() {
    setStr("hello", "world")
    go call1sec(flag)
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8089", nil))
}
