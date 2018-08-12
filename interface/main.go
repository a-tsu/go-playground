package main

import (
    "fmt"
)

type api interface {
    getLine() string
}

type Data struct {
    line string
}

func (d *Data) getLine() string {
    key := "don't call this at the test " + d.line
    return key
}

func main() {
    d := &Data{
        line:  "a-tsu",
    }
    var i api = d
    fmt.Println(i.getLine())
}
