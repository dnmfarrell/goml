package ml

import (
	"testing"
)

func TestConsL(t *testing.T) {
	xs := RangeInf[int]{0, 1}.Eval()
	if xs.Empty() {
		t.Error("ConsL.Empty is true")
	}
	if !xs.Lazy() {
		t.Error("ConsL.Lazy is false")
	}
	if xs.Head() != 0 {
		t.Errorf("Head() returns %d", xs.Head())
	}
	if xs.Tail().Head() != 1 {
		t.Errorf("Tail().Head() returns %d", xs.Tail().Head())
	}
}
