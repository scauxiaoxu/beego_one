package models

import "fmt"

func coverPanic() {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string message", message)
	case error:
		fmt.Println("error message", message)
		//default:
		//	fmt.Println("unknown message", message)
	}
}
