package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Andromeda")

	if result != "Hello Andromeda" {
		t.Fail()
	}

	fmt.Println("TestHelloWorldEko Done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		t.FailNow()
	}

	fmt.Println("TestHelloWorldKhannedy Done")

}