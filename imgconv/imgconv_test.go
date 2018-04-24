package imgconv

import (
	"log"
	"os"
	"testing"
)

type MockImagefile struct{}

func TestNewImages(t *testing.T) {
	filename, imglist, err := NewImages("../images")

	if filename == nil {
		t.Fatal("fail error func test")
	}
	if imglist == nil {
		t.Fatal("fail error func test")
	}
	if err != nil {
		t.Fatal("fail error func test")
	}
}

func TestImgconv(t *testing.T) {
	t.Helper()
	dir, _ := os.Getwd()
	println(dir)

	filepath, image, err := NewImages("../images")
	if err != nil {
		log.Fatal(err)
	}

	outType := []struct {
		outtype string
	}{
		{outtype: "png"},
		{outtype: "jpg"},
	}

	for _, c := range outType {

		err = Imgconv(c.outtype, filepath, image)
		if err != nil {
			t.Fatal("fail error func test")
		}
		// _, err := os.Stat(filepath[0] + c.outtype)
		// if err != nil {
		// 	t.Fatal("fail error func test")
		// }
	}
}

/*
package main

import (
	"fmt"
)

func main() {
	s := &StudentImpl{}
	// s := &StudentMock{}
	name, age := Show(s)
	fmt.Println("name: ", name)
	fmt.Println("age: ", age)
}

func Show(s Student) (string, int) {
	name := s.Name()
	age := s.Age()
	return name, age
}

type Student interface {
	Name() string
	Age() int
}

// db access
type StudentImpl struct{}

func (s *StudentImpl) Name() string {
	name := "Taro"
	//fmt.Println(name)
	return name
}

func (s *StudentImpl) Age() int {
	age := 15
	//fmt.Println(age)
	return age
}
*/

/*
package main

import (
	"testing"
)

// mock
type StudentMock struct{}

func (s *StudentMock) Name() string {
	name := "Mock Taro"
	return name
}

func (s *StudentMock) Age() int {
	age := 100
	return age
}

func Testshow(t *testing.T) {
	s := &StudentMock{}
	name, age := Show(s)
	if name != "Mock Taro" {
		t.Fatalf("failed test %#v", name)
	}
	if age != 100 {
		t.Fatal("failed test")
	}
}
*/
