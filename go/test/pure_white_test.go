/* test package for go's test
RUN:
locate in test pacakge, then run `go test .`
*/
package test

import (
	"fmt"
	"testing"
)

func TestPureFunc1(t *testing.T) {
	fmt.Println("test func 1111 starting...")
	helloTest()
	fmt.Println("test func 1111 end.")
}
func TestPureFunc2(t *testing.T) {
	fmt.Println("test func 2222 starting...")
	helloTest()
	fmt.Println("test func 2222 end.")
}
