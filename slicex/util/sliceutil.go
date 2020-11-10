package util

//对比两个切片
func CompareSlices(x []int, y []int) (int, bool) {
	if len(x) != len(y) {
		return -1, false
	}
	for post := range x {
		if x[post] != y[post] {
			return post, false
		}
	}
	return -1, true
}
