package main

import (
	//"errors"
	"fmt"
)

func main() {
	err:=moo();
	fmt.Println(err)

	//baseErr:=errors.unwrap()
	//fmt.Println(baseErr)

}

func cat() error {
	return fmt.Errorf("cat() Error")
}

func moo() error {
	cat();
	return fmt.Errorf("moo() Error")
}