package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldGhozi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ghozi")
	}
}

func BenchmarkHelloWorldHumama(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Humama")
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("This unit test not running mac OS")
	}
	result := HelloWorld("Ghozi")
	require.Equal(t, "Hello Ghozi", result, "Result must be 'Hello Ghozi'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ghozi")

	require.Equal(t, "Hello Ghozi", result, "Result must be 'Hello Ghozi'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ghozi")

	assert.Equal(t, "Hello Ghozi", result, "Result must be 'Hello Ghozi'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldGhozi(t *testing.T) {
	result := HelloWorld("Ghozi")

	if result != "Hello Ghozi" {
		t.Error("Result must be 'Hello Ghozi'")
	}

	fmt.Println("TestHelloWorldGhozi Done")
}

func TestHelloWorldHumama(t *testing.T) {
	result := HelloWorld("Humama")

	if result != "Hello Humama" {
		t.Fatal("Result must be 'Hello Humama'")
	}

	fmt.Println("TestHelloWorldHumama Done")
}
