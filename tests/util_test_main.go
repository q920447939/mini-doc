package main

import (
	"wahaha/utils/cryptil"
	"fmt"
)

func main() {
	id := cryptil.UniqueId()
	fmt.Println(len(id))
}
