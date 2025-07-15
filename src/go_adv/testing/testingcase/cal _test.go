package testingcase

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(10)
	if result != 55 {
		t.Fatalf("实际值%d,期待值%d", result, 55)
	}
	t.Logf("add(10)执行正确")
}
