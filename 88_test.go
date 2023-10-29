package ll88

import (
	"testing"
)

func TestMerge1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	want := []int{1, 2, 2, 3, 5, 6}

	got, err := Merge(nums1, m, nums2, n)

	if !equal(got, want) || err != nil {
		t.Fatalf("you fucked it up this time, got %q, wanted %q: Error: %q", got, want, err)
	}
}
func TestMerge2(t *testing.T) {
	nums1 := []int{1}
	m := 1
	nums2 := []int{}
	n := 0

	want := []int{1}

	got, err := Merge(nums1, m, nums2, n)

	if !equal(got, want) || err != nil {
		t.Fatalf("you fucked it up this time, got %q, wanted %q: Error: %q", got, want, err)
	}
}
func TestMerge3(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1

	want := []int{1}

	got, err := Merge(nums1, m, nums2, n)

	if !equal(got, want) || err != nil {
		t.Fatalf("you fucked it up this time, got %q, wanted %q: Error: %q", got, want, err)
	}
}

func TestMerge4(t *testing.T) {
	nums1 := []int{1, 0}
	m := 1
	nums2 := []int{2}
	n := 1

	want := []int{1, 2}

	got, err := Merge(nums1, m, nums2, n)

	if !equal(got, want) || err != nil {
		t.Fatalf("you fucked it up this time, got %q, wanted %q: Error: %q", got, want, err)
	}
}

func TestMerge5(t *testing.T) {
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1

	want := []int{1, 2}

	got, err := Merge(nums1, m, nums2, n)

	if !equal(got, want) || err != nil {
		t.Fatalf("you fucked it up this time, got %q, wanted %q: Error: %q", got, want, err)
	}
}

// nums1 =
// [2,0]
// m =
// 1
// nums2 =
// [1]
// n =
// 1

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
