package main

import (
    "testing"
    "net/http"
)

func TestSetStrSuccess(t *testing.T) {
    setStr("hoge", "huga")
    expected := "hoge huga"

    if str != expected {
        t.Errorf("got: %v\nwant: %v", str, expected)
    }
}
