// A "Hello World" program that prints a greeting with the current time.
package main

import (
	"fmt"
	"time"
	//"strings"
)

//import "time"

func getSomething() { return }

func getInteger() int { return 123 }

// greeting returns a pleasant, semi-useful greeting.
func greeting() string {
	return "Hello world, the time is: " + time.Now().String()
}

func main2() {
	fmt.Println("Starting Go Program")
	fmt.Println(greeting())
	fmt.Println("Ending Go Program")
	// fmt.Println "Ending Go Program"
}
