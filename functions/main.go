package main

import "fmt"

// entry point
// func <reciever> <func_name>() <return> {}
func main() {
	// func expression
	// with it, it is possible to create
	// a function inside a function
	greetings := func() {
		fmt.Println("This is func expression")
		fmt.Println("Here are all the ways to write functions:")
	}
	greetings()

	// test of the passing a func to a variable
	assignedFunc := returnFunc
	fmt.Println(assignedFunc(2, 5))

	fmt.Println(returnFunc(1, 1))
	fmt.Println(namedReturnFunc(2, 1))
	fmt.Println(multipleReturnsFunc(6, 3))
	multipleReturnes, _ := multipleReturnsFunc(10, 5)
	// there should be 2 returned values
	// if only one is needed, just put "_" instead of a var name
	fmt.Println(multipleReturnes)

	// calling the callback function which takes a function as a parameter
	// n int is a created variable
	// not much use in GoLang
	callbackFunc([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(n int) {
		fmt.Print(n)
	})

	farewell := funcRetunsFunc()
	fmt.Println(farewell())

}

func returnFunc(numberA, numberB int) string {
	sum := numberA + numberB
	return fmt.Sprint(numberA, "+", numberB, "=", sum)
}

func namedReturnFunc(numberA, numberB int) (substrc string) {
	s := numberA - numberB
	substrc = fmt.Sprint(numberA, "-", numberB, "=", s)
	return // since the only one return has been assigned
	// the return doesn't need to be assigned
}

func multipleReturnsFunc(numberA, numberB int) (string, string) {
	div1 := numberA / numberB
	div2 := numberB / numberA
	return fmt.Sprint(numberA, "/", numberB, "=", div1),
		fmt.Sprint(numberB, "/", numberA, "=", div2)
}

// variadic = optional = "..."
// for params the dots str in the front of type
// for args the dots are in the back of variable

func funcRetunsFunc() func() string {
	// function which returns a <func() string> type
	fmt.Println("This function returns a funcition")
	return func() string {
		return fmt.Sprint("The function is inside of a function")
	}
}

func clouser() {
	outsideVar := 1

	{
		insideVar := 10
		fmt.Println(outsideVar)
		fmt.Println(insideVar)
	}
	fmt.Println(outsideVar)
	// the inside var is not accessable
}

// callback func takes a function as a parameter
func callbackFunc(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
