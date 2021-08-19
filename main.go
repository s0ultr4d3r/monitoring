package main

import (
	"fmt"
)

func main() {
	a := parse()
	fmt.Println("catch")
	b := prblmSlc(a)
	c := uniqPrblms(b)
	for _, el := range c {
		fmt.Println(el)
	}

}
