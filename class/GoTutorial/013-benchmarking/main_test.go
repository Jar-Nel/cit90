package main

import "testing"

func BenchmarkComplit(b *testing.B){
	for i:=0;i<200;i++{
		compLit()
	}
}

func BenchmarkMakeway(b *testing.B){
	for i:=0;i<200;i++{
		makeway()
	}
}