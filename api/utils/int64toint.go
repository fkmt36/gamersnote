package utils

import "math"

// Int64ToInt int64をintに変換します
func Int64ToInt(i int64) int {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0
	}
	return int(i)
}
