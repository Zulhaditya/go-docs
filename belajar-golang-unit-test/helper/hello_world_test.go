package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
menjalankan unit test:
- go test
- go test -v
- go test -v
- go test -v -run=TestHelloWorldInayah => untuk testing fungsi tertentu
- go test -v ./... => running unit test di semua package (root folder)
*/

// go test -v -run=TestTableHelloWorld
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Inayah",
			request:  "Inayah",
			expected: "Hello Inayah",
		},
		{
			name:     "Wulandari",
			request:  "Wulandari",
			expected: "Hello Wulandari",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

// subtest
// cara running:
// go test -v -run=TestSubTest
// go test -v -run=TestSubTest/Syapiq
func TestSubTest(t *testing.T) {
	t.Run("Syapiq", func(t *testing.T) {
		result := HelloWorld("Syapiq")
		require.Equal(t, "Hello Syapiq", result, "Result must be 'Hello Syapiq'")
	})

	t.Run("Iqbal", func(t *testing.T) {
		result := HelloWorld("Iqbal")
		require.Equal(t, "Hello Iqbal", result, "Result must be 'Hello Iqbal'")
	})

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ackxle")
	if result != "Hello Ackxle" {
		// unit test gagal
		t.Error("Result must be 'Hello Ackxle'")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldInayah(t *testing.T) {
	result := HelloWorld("Inayah")
	if result != "Hello Inayah" {
		// unit test gagal
		t.Fatal("Result must be 'Hello Inayah'")
	}

	fmt.Println("TestHelloWorldInayah done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ackxle")
	assert.Equal(t, "Hello Ackxle", result, "Result must be 'Hello Ackxle'")
	fmt.Println("TestHelloWorldAssert with assert done!")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ackxle")
	require.Equal(t, "Hello Ackxle", result, "Result must be 'Hello Ackxle'")
	fmt.Println("TestHelloWorldAssert with require done!")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("cannot run on mac os")
	}

	result := HelloWorld("Ackxle")
	require.Equal(t, "Hello Ackxle", result, "Result must be 'Hello Ackxle'")
}

// benchmark
/*
menjalankan benchmark:
- go test -v -bench=. => menjalankan unit test dan benchmark
- go test -v -run=NotMatchUnitTest -bench=. => menjalankan benchmark saja
- go test -v -run=NotMatchUnitTest -bench=HelloWorldInayah => menjalankan benchmark tertentu
*/
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ackxle")
	}
}

func BenchmarkHelloWorldInayah(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Inayah Fitri Wulandari")
	}
}

// sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Syapiq", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Syapiq Alfarazi")
		}
	})
	b.Run("Iqbal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Iqbal Ramadhan")
		}
	})
}

// table benchmark
// cara menjalankan:
// go test -v -run=TestTidakAda -bench=BenchmarkTable
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Inayah",
			request: "Inayah",
		},
		{
			name:    "Fitri",
			request: "Fitri",
		},
		{
			name:    "Wulandari",
			request: "Wulandari",
		},
		{
			name:    "Muhammad Zulhaditya Hapiz",
			request: "Muhammad Zulhaditya Hapiz",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
