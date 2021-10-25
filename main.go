package main

import (
	"calc"
	"fmt"
)

func main() {
	newLog := NewSimpleLog()
	calc.Calculate(newLog)
	fmt.Println(newLog.All())
}
