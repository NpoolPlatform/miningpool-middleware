package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	oo, ee := json.Marshal(5)
	fmt.Println(string(oo), ee)
}
