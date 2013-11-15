package main

import (
	"fmt"
	"github.com/msiebuhr/go-tag-test" // Whatever is on master
	//"github.com/msiebuhr/go-tag-test@master" // Ditto
	//"github.com/msiebuhr/go-tag-test@v0.0.0" // Tag v0.0.0 / First release
	//"github.com/msiebuhr/go-tag-test@v0.0.1" // Tag v0.0.1 / Second release
)

func main() {
	fmt.Println(TagTest.Tag())
}
