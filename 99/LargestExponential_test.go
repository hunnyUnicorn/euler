package LargestExponential

import "testing"

func TestFindLargestNumber(t *testing.T) {

	gotBase, gotExponent := findLargestNumber(3, []int{4, 3, 2}, []int{7, 7, 11})
	wantedBase := 4
	wantedExponent := 7

	if gotBase != wantedBase || gotExponent != wantedExponent {
		t.Errorf("Error in find largest number : got %d %d, wanted %d %d", gotBase, gotExponent, wantedBase, wantedExponent)
	}
}

var values ByValue

func TestLength(t *testing.T) {
	values = ByValue{ExponentValues{0, 0, 0}, ExponentValues{1, 1, 1}}
	gotLength := values.Len()
	wantedLength := 2

	if gotLength != wantedLength {
		t.Errorf("Error in get length : got %d, wanted %d", gotLength, wantedLength)
	}
}

func TestLess(t *testing.T) {
	values = ByValue{ExponentValues{0, 0, 0}, ExponentValues{1, 1, 1}}
	gotBool := values.Less(0, 1)
	wantedBool := false

	if gotBool != wantedBool {
		t.Errorf("Error in compare : got %v, wanted %v", gotBool, wantedBool)
	}
}

func TestSwap(t *testing.T) {
	values = ByValue{ExponentValues{0, 0, 0}, ExponentValues{1, 1, 1}}
	values.Swap(0, 1)

	if values[0].base != 1 || values[0].exponent != 1 || values[1].base != 0 || values[1].exponent != 0 {
		t.Errorf("Error in swap")
	}
}
