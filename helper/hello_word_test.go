package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Andromeda")

	if result != "Hello Andromeda" {
		panic("Result is not Hello Andromeda")
	}
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		panic("Result is not Hello Khannedy")
	}
}