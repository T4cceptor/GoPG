package main

import (
	"fmt"
	"log"
	"os"

	"container/list"
	"reflect"

	"example.com/greetings"
)

var logger = log.New(os.Stdout, "greetings: ", 1)

func printList() {
	var x list.List
	x.PushBack(1)
	x.PushBack("test")
	x.PushBack(map[string]int{"a": 123})

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		fmt.Println(reflect.TypeOf(e.Value))
	}
}

func checkRandomEntryFunc() {
	tests := []string{
		"Hi, Welcome!",
		"Great to see you!",
		"Hail! Well met!",
	}
	fmt.Println(greetings.RandomEntry(tests))

	tests2 := []int{1, 2, 3, 4}
	fmt.Println(greetings.RandomEntry(tests2))

}

func extendingSlice() {
	var test []int = []int{1, 2, 3}
	for i, v := range test {
		fmt.Printf("%v: %v\n", i, v)
	}
	test = append(test, 123)
	fmt.Println(test[3])
}

// Function Constructors - chaining of functions as return types
// Note: this enables us to write decorators and closures!
func makeFunction(name string) func() func() func() {
	fmt.Println("00000")
	// res is a closure!
	// we can add arbitrary parameters to makeFunction
	// and "store" them in the returned function
	var res = func() func() func() {
		fmt.Println("test1")
		return makeFunction2(name)
	}
	return res
}

func makeFunction2(name string) func() func() {
	fmt.Println("11111")
	return func() func() {
		return makeFunction3(name)
	}
}

func makeFunction3(name string) func() {
	fmt.Println("33333")
	return func() {
		fmt.Printf(name)
	}
}

func functionReturnTest() {
	f := makeFunction("hellooo")
	f()()()
}

func main() {
	// printList()
	// checkRandomEntryFunc()
	// extendingSlice()
	// functionReturnTest()
	main2()

	// type Node struct {
	// 	Next  *Node
	// 	Value interface{}
	// }
	// var first *Node = Node()

	// visited := make(map[*Node]bool)
	// for n := first; n != nil; n = n.Next {
	// 	if visited[n] {
	// 		fmt.Println("cycle detected")
	// 		break
	// 	}
	// 	visited[n] = true
	// 	fmt.Println(n.Value)
	// }

	// type Person struct {
	// 	Name  string
	// 	Likes []string
	// }
	// var people []*Person

	// likes := make(map[string][]*Person)
	// for _, p := range people {
	// 	for _, l := range p.Likes {
	// 		likes[l] = append(likes[l], p)
	// 	}
	// }
}

func main2() {

	// var test string
	var err error

	// test, err = greetings.PrintMessage("test")

	messages, err := greetings.GetMessages([]string{"test1", "test2"})
	if err != nil {
		logger.Fatal(err)
	}

	for messageIndex, message := range messages {
		fmt.Printf("%v: %v\n", messageIndex, message)
	}
	// fmt.Println(test)
}
