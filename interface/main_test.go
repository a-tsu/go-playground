package main

import (
    "testing"
)

type FakeData struct {
    line string
}

func (f *FakeData) getLine() string {
    key := "it's ok to call this at the test " + f.line
    return key
}

func TestPrintSuccess(t *testing.T) {
    f := &FakeData{
        line:  "a-tsu",
    }
    var i api = f
    res := i.getLine()

    expected := "it's ok to call this at the test a-tsu"

    if res != expected {
        t.Errorf("got: %v\nwant: %v", res, expected)
    }
}
