package main

import (
	"fmt"
	"sync"
)

var (
	//counter int32          // counter is a variable incremented by all goroutines.
	wg sync.WaitGroup // wg is used to wait for the program to finish.
	//mutex   sync.Mutex     // mutex is used to define a critical section of code.
)

func main() {
	wg.Add(2) // Add a count of two, one for each goroutine.
	//defer wg.Wait()
	var a, b int
	a = 20
	b = 2
	go addtion(a, b)
	go multiply(a, b)

	wg.Wait() // Wait for the goroutines to finish.

}

func multiply(a int, b int) {
	//defer wg.Done() // Schedule the call to Done to tell main we are done.
	//mutex.Lock()
	//defer mutex.Unlock()
	//for i := 0; i < 3; i++ {
	//mutex.Lock()
	//{
	//fmt.Println(lang)
	//counter++
	//}
	//mutex.Unlock()
	//}
	//wg.Done()
	result := a * b
	fmt.Printf("multiplication of %d and %d is %d \n", a, b, result)
	wg.Done()
}
func addtion(a int, b int) {

	result := a + b
	fmt.Printf("addition of %d and %d is %d \n", a, b, result)
	wg.Done()
}
