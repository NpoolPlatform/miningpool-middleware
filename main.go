package main

import "fmt"

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	err = returnAErr()
	aha, err := returnErr()
	fmt.Println(aha)
}

func returnAErr() error {
	return fmt.Errorf("sdfasdf111111111111111a")
}

func returnErr() (string, error) {
	return "eadgf", fmt.Errorf("sdfasdfa")
}
