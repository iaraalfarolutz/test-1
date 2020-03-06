package main

import (
	"fmt"
	"github.com/iaraalfarolutz/test-1/falabella_test1"
)

func main() {
	myTLV := []byte("A0511AB398765UJ1N230200")
	myMap, err := tlvToMap.ToMap(myTLV)
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}
