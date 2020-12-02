package closest

import "math"

func findClosestToZero(integers []int) int {
	if len(integers) == 0 {
		return 0
	}
	var closest = integers[0]
	for _, n := range integers {
		abs_n := int(math.Abs(float64(n)))
		abs_closest := int(math.Abs(float64(closest)))
		if abs_n < abs_closest {
			closest = n
		}
	}
	return closest;
}