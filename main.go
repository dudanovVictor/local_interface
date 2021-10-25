package main

import (
	"fmt"
	"local_interface/calc"
)

func main() {
	newLog := NewSimpleLog()
	calc.Calculate(newLog)
	fmt.Println(newLog.All())
}
