package imgconv

import (
	"fmt"
	"testing"
)

func Testtest(t *testing.T) {
	fmt.Println(test("hello"))
	if test("hello") != "10" {
		t.Fatal("error")
	} else {
		t.Fatal("error")
	}
}

func TestSum(t *testing.T) {
	fmt.Println("hello" == "10")
	if sum(1, 2) != 3 {
		t.Fatal("fail sum func test")
	}
}

// func TestNewImages() {
// 	return
// }
