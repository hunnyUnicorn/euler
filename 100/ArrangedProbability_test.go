package ArrangedProbability

import (
	"math/big"
	"testing"
)

func Test1(t *testing.T) {

	gotCountofBlueDisk, gotTotalCountofDisk := solve(1, 2, 2)
	wantedCountofBlueDisk := big.NewInt(3)
	wantedTotalCountofDisk := big.NewInt(4)

	if gotCountofBlueDisk.Cmp(wantedCountofBlueDisk) != 0 || gotTotalCountofDisk.Cmp(wantedTotalCountofDisk) != 0 {
		t.Errorf("Error in solve : got %d %d, wanted %d %d", gotCountofBlueDisk, gotTotalCountofDisk, wantedCountofBlueDisk, wantedTotalCountofDisk)
	}
}

func Test2(t *testing.T) {

	gotCountofBlueDisk, gotTotalCountofDisk := solve(1, 2, 5)
	wantedCountofBlueDisk := big.NewInt(15)
	wantedTotalCountofDisk := big.NewInt(21)

	if gotCountofBlueDisk.Cmp(wantedCountofBlueDisk) != 0 || gotTotalCountofDisk.Cmp(wantedTotalCountofDisk) != 0 {
		t.Errorf("Error in solve : got %d %d, wanted %d %d", gotCountofBlueDisk, gotTotalCountofDisk, wantedCountofBlueDisk, wantedTotalCountofDisk)
	}
}

func Test3(t *testing.T) {

	gotCountofBlueDisk, gotTotalCountofDisk := solve(1, 2, 30)
	wantedCountofBlueDisk := big.NewInt(85)
	wantedTotalCountofDisk := big.NewInt(120)

	if gotCountofBlueDisk.Cmp(wantedCountofBlueDisk) != 0 || gotTotalCountofDisk.Cmp(wantedTotalCountofDisk) != 0 {
		t.Errorf("Error in solve : got %d %d, wanted %d %d", gotCountofBlueDisk, gotTotalCountofDisk, wantedCountofBlueDisk, wantedTotalCountofDisk)
	}
}

func TestAdd(t *testing.T) {

	gotValue := add(big.NewInt(1), big.NewInt(4))
	wantedValue := big.NewInt(5)

	if gotValue.Cmp(wantedValue) != 0 {
		t.Errorf("Error in add : got %d, wanted %d", gotValue, wantedValue)
	}
}

func TestSub(t *testing.T) {

	gotValue := sub(big.NewInt(5), big.NewInt(1))
	wantedValue := big.NewInt(4)

	if gotValue.Cmp(wantedValue) != 0 {
		t.Errorf("Error in sub : got %d, wanted %d", gotValue, wantedValue)
	}
}

func TestMul(t *testing.T) {

	gotValue := mul(big.NewInt(5), big.NewInt(2))
	wantedValue := big.NewInt(10)

	if gotValue.Cmp(wantedValue) != 0 {
		t.Errorf("Error in Mul : got %d, wanted %d", gotValue, wantedValue)
	}
}

func TestDiv(t *testing.T) {

	gotValue := div(big.NewInt(10), big.NewInt(2))
	wantedValue := big.NewInt(5)

	if gotValue.Cmp(wantedValue) != 0 {
		t.Errorf("Error in div : got %d, wanted %d", gotValue, wantedValue)
	}
}

func TestMod(t *testing.T) {

	gotValue := mod(big.NewInt(5), big.NewInt(2))
	wantedValue := big.NewInt(1)

	if gotValue.Cmp(wantedValue) != 0 {
		t.Errorf("Error in mod : got %d, wanted %d", gotValue, wantedValue)
	}
}
