package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}
