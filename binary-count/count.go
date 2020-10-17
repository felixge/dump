package count

func Count(n int) int {
	var (
		curCount = -1
		maxCount = 0
	)
	for n > 0 {
		one := n&1 > 0
		if one {
			if curCount > maxCount {
				maxCount = curCount
			}
			curCount = 0
		} else if curCount != -1 {
			curCount++
		}
		n = n >> 1
	}
	return maxCount
}
