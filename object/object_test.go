package object

import "testing"

func TestStringHashKey(t *testing.T) {
	foo1 := &String{Value: "Hello World"}
	foo2 := &String{Value: "Hello World"}
	bar1 := &String{Value: "Konichiwa"}
	bar2 := &String{Value: "Konichiwa"}

	if foo1.HashKey() != foo2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if bar1.HashKey() != bar2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if foo1.HashKey() == bar1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
