package test

import (
	"fmt"
	"time"
)

// HelloTest public function
func HelloTest() {
	helloTest()
}

// helloTest private function
func helloTest() {
	time.Sleep(time.Second * 3)
	fmt.Println("hello test!")
}
