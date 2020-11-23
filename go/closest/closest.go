package closest

func findClosestToZero(integers []int) int {
	var closest = integers[0]
	for _, n := range integers {
		if n < closest {
			closest = n
		}
	}
	return closest;
}