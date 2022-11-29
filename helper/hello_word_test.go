package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Andromeda")

	if result != "Hello Andromeda" {
		t.Error("Result must be 'Hello Eko'")
	}

	fmt.Println("TestHelloWorldEko Done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		t.Fatal("Result must be 'Hello Khannedy'")
	}

	fmt.Println("TestHelloWorldKhannedy Done")

}