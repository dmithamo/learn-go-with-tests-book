package arraysandslices

// SumMembers returns the sum of the members of a slice
func SumMembers(members []int) int {
    if len(members) == 0 {
        return 0
    }

    sum := 0
    for _, member := range members {
        sum += member
    }
    return sum
}

// MultiSliceSum adds members of each slice and returns slice with sums
func MultiSliceSum(slices [][]int) []int {
    sums := make([]int, 0, len(slices))
    for _, slice := range slices {
        sums = append(sums, SumMembers(slice))
    }

    return sums
}

func MultiSliceSumTails(slices [][]int) []int {
    sums := make([]int, 0, len(slices))
    for _, slice := range slices {
        head := 0
        if len(slice) > 0 {
            head = slice[0]
        }
        sums = append(sums, SumMembers(slice)-head)
    }

    return sums
}
