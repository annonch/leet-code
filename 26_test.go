package ll26

import (
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {
	nums := []int{1, 1, 2}

	wantVal := 2
	wantNums := []int{1, 2, 0}
	gotVal, gotNums, err := RemoveDuplicates(nums)

	if !equal(gotNums, wantNums) || wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %v %v, wanted %v %v: Error: %q", gotVal, gotNums, wantVal, wantNums, err)
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	wantVal := 5
	wantNums := []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0}
	gotVal, gotNums, err := RemoveDuplicates(nums)

	if !equal(gotNums, wantNums) || wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %v %v, wanted %v %v: Error: %q", gotVal, gotNums, wantVal, wantNums, err)
	}
}

func TestRemoveDuplicates3(t *testing.T) {
	nums := []int{0}

	wantVal := 1
	wantNums := []int{0}
	gotVal, gotNums, err := RemoveDuplicates(nums)

	if !equal(gotNums, wantNums) || wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %v %v, wanted %v %v: Error: %q", gotVal, gotNums, wantVal, wantNums, err)
	}
}

func TestRemoveDuplicates4(t *testing.T) {
	nums := []int{}

	wantVal := 0
	wantNums := []int{}
	gotVal, gotNums, err := RemoveDuplicates(nums)

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
