package dogs

import (
	"fmt"
	"testing"
)

func BenchmarkToDogYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToDogYears(10)
	}
}

func BenchmarkToHumanYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToHumanYears(56)
	}
}

func ExampleToDogYears() {
	fmt.Println(ToDogYears(10))
	// Output:
	// 70
}

func ExampleToHumanYears() {
	fmt.Println(ToHumanYears(56))
	// Output:
	// 8
}

func TestToDogYears(t *testing.T) {
	n := ToDogYears(10)
	if n != 70 {
		t.Error("Got:", n, "want:", 70)
	}
}

func TestToHumanYears(t *testing.T) {
	n := ToHumanYears(56)
	if n != 8 {
		t.Error("Got:", n, "want:", 8)
	}
}
