package main

import "fmt"

func main() {
	authed := false
	aa := &authed
	fmt.Println(authed, *aa)
	authed = true
	fmt.Println(authed, *aa)
}
