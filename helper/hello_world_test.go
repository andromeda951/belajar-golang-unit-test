package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

// tidak perlu menambah-nambah function test, gunakan table test
func TestHelloWorldKurniawan(t *testing.T) {
	result := HelloWorld("Kurniawan")

	if result != "Hello Kurniawan" {
		t.Fatal("Result must be 'Hello Kurniawan'")
	}

	fmt.Println("TestHelloWorldKurniawan Done")

}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux OS")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run() // eksekusi semua unit test

	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T)  {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})

	t.Run("Kurniawa", func(t *testing.T)  {
		result := HelloWorld("Kurniawa")
		require.Equal(t, "Hello Kurniawa", result, "Result must be 'Hello Kurniawa'")
	})

	// tidak perlu menambah-nambah subtest, gunakan table test
	t.Run("Khannedy", func(t *testing.T)  {
		result := HelloWorld("Khannedy")
		require.Equal(t, "Hello Khannedy", result, "Result must be 'Hello Khannedy'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Eko)",
			request: "Eko",
			expected: "Hello Eko",
		},
		{
			name: "HelloWorld(Kurniawan)",
			request: "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name: "HelloWorld(Khannedy)",
			request: "Khannedy",
			expected: "Hello Khannedy",
		},
		{
			name: "HelloWorld(Budi)",
			request: "Budi",
			expected: "Hello Budi",
		},
		{
			name: "HelloWorld(Joko)",
			request: "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name,func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}







