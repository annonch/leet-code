package ll150

import (
	"testing"
)

func TestMajorityElement1(t *testing.T) {
	nums := []int{3, 2, 3}

	wantVal := 3
	gotVal, err := MajorityElement(nums)

	if wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %v, wanted %v: Error: %q", gotVal, wantVal, err)
	}
}

func TestMajorityElement2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}

	wantVal := 2
	gotVal, err := MajorityElement(nums)

	if wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %v, wanted %v: Error: %q", gotVal, wantVal, err)
	}
}

func TestMajorityElement3(t *testing.T) {
	nums := []int{3}

	wantVal := 3
	gotVal, err := MajorityElement(nums)

	if wantVal != gotVal || err != nil {
		t.Fatalf("you fucked it up this time, got %v, wanted %v: Error: %q", gotVal, wantVal, err)
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
