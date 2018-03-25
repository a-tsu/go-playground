package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestSetStrSuccess(t *testing.T) {
    setStr("hoge", "huga")
    expected := "hoge huga"

    if str != expected {
        t.Errorf("got: %v\nwant: %v", str, expected)
    }
}


func TestHttpSuccess(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(handler))
    defer ts.Close()

    res, err := http.Get( ts.URL )
    if err != nil {
        t.Error("unexpected")
        return
    }

    if res.StatusCode != 200 {
        t.Error("Status code error")
        return
    }

}
