package main

import "fmt"

func main() {
	ok, err := say("Hello World")
	if err != nil {
		panic(err.Error())
	}
	switch ok {
	case true:
		fmt.Println("Everything Ok")
	default:
		fmt.Println("Got an error")
	}
}

func say(what string) (bool, error) {
	if what == "" {
		return false, fmt.Errorf("Empty String")
	}
	fmt.Println(what)
	return true, nil
}
