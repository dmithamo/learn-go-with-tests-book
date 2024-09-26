package arraysandslices

import (
	testinghelpers "learninggowithtests/helpers"
	"testing"
)

func TestSumMembers(t *testing.T) {
	t.Run("slice with any number of members", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}
		want := 15
		got := SumMembers(given)

		testinghelpers.AssetCorrectMsg(t, got, want, given)
	})
}

func TestMultiSliceSums(t *testing.T) {
	t.Run("several slices", func(t *testing.T) {
		given := [][]int{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 6}}
		want := []int{15, 20}
		got := MultiSliceSum(given)
		testinghelpers.AssetCorrectMsg(t, got, want, given)
	})

	// Test is unimportant - does not affect coverage, hence its
	//   concerns are already addressed in the above test
	// t.Run("single slice", func(t *testing.T) {
	// 	given := [][]int{[]int{1, 2, 3, 4, 5}}
	// 	want := []int{15}
	// 	got := MultiSliceSum(given)
	// 	testinghelpers.AssetCorrectMsg(t, got, want, given)
	// })
}

func TestMultiSliceSumTails(t *testing.T) {
	given := [][]int{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 6}, []int{1, 2}, []int{0, 9}}
	want := []int{14, 18, 2, 9}
	t.Run("handles several non-empty slices", func(t *testing.T) {
		got := MultiSliceSumTails(given)
		testinghelpers.AssetCorrectMsg(t, got, want, given)
	})

	t.Run("handles empty slices", func(t *testing.T) {
		given = append(given, []int{})
		want = append(want, 0)
		got := MultiSliceSumTails(given)
		testinghelpers.AssetCorrectMsg(t, got, want, given)
	})
}
