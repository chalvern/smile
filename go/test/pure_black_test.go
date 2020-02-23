/* test package for go's test
RUN:
locate in test pacakge, then run `go test .`
*/
package test_test

import (
	"fmt"
	"testing"

	"github.com/chalvern/smile/go/test"
)

func TestMain(m *testing.M) {
	fmt.Println("main starting")
	m.Run()
	fmt.Println("main end.")
}

func TestPureFunc1(t *testing.T) {
	fmt.Println("test_test func 1111 starting...")
	test.HelloTest()
	fmt.Println("test_test func 1111 end.")
}
func TestPureFunc2(t *testing.T) {
	fmt.Println("test_test func 2222 starting...")
	test.HelloTest()
	fmt.Println("test_test func 2222 end.")
}
