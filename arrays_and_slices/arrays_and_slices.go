package arraysandslices

func SumMembers(members []int) int {
	sum := 0
	for _, member := range members {
		sum += member
	}
	return sum
}
