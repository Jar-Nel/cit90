package main

import "testing"

func BenchmarkEcho1(b *testing.B) {
	files:= []string {"cmd","Word","Word","Test"}
	for i:=0; i<b.N; i++{
		echo1(files)
	}
}

func BenchmarkEcho2(b *testing.B) {
	files:= []string {"cmd","Word","Word","Test"}
	for i:=0; i<b.N; i++{
		echo2(files)
	}
}

func BenchmarkEcho3(b *testing.B) {
	files:= []string {"cmd","Word","Word","Test"}
	for i:=0; i<b.N; i++{
		echo3(files)
	}
}