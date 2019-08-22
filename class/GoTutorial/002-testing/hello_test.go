package main

import "testing"

func TestFoo(t *testing.T) {
    want := "Hello, world."
    if got := foo(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}