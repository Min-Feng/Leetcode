package main

import (
	"fmt"
	"time"
)

func reverse(x int) int {
	var newValue, inputValue int64

	inputValue = int64(x)
	for inputValue != 0 {
		newValue = newValue*10 + inputValue%10
		inputValue = inputValue / 10
		fmt.Println(newValue, inputValue)
		time.Sleep(200 * time.Millisecond)
	}

	//如果發生overflow的現象,兩者數字一定不同
	returnValue := int32(newValue)
	if int64(returnValue) != newValue {
		return 0
	}

	return int(returnValue)
}

func main() {
	reverse(12400)
}
