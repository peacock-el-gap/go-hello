package main

import (
	"fmt"

	"github.com/peacock-el-gap/go/stringutils"
)

func main() {
	fmt.Println("hello, world")
	fmt.Println(stringutils.Reverse("ABCDEFGHIJK"))
	fmt.Println(stringutils.Reverse("123456"))
	fmt.Println(stringutils.Reverse(stringutils.Reverse("ABC1234")))
}
