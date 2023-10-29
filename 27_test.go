package ll27

import (
	"testing"
)

func TestRemoveElement1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3

	wantVal := 2
	wantNums := []int{2, 2, 0, 0}
	gotVal, gotNums, err := RemoveElement(nums, val)

	if !equal(gotNums, wantNums) || wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %q %q, wanted %q %q: Error: %q", gotVal, gotNums, wantVal, wantNums, err)
	}
}

func TestRemoveElement2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	wantVal := 5
	wantNums := []int{0, 1, 4, 0, 3, 0, 0, 0}
	gotVal, gotNums, err := RemoveElement(nums, val)

	if !equal(gotNums, wantNums) || wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %v %v, wanted %v %v: Error: %q", gotVal, gotNums, wantVal, wantNums, err)
	}
}

// nums.length == 0

func TestRemoveElement3(t *testing.T) {
	nums := []int{}
	val := 2

	wantVal := 0
	wantNums := []int{}
	gotVal, gotNums, err := RemoveElement(nums, val)

	if !equal(gotNums, wantNums) || wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %v %v, wanted %v %v: Error: %q", gotVal, gotNums, wantVal, wantNums, err)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
