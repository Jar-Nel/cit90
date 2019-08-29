package main

import "testing"

func TestFoo(t *testing.T) {
	want:=42
	if got:=foo(); got != want{
		t.Errorf("foo() error  want: %v, got %v",want,got)
	}
}