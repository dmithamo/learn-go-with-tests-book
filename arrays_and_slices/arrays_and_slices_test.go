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
