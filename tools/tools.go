package tools

// type DistanceRow = interfaces.DistanceRow
// type DistanceMatrix = interfaces.DistanceMatrix
// type DistanceTable = interfaces.DistanceTable

func Min64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	} else if a == b {
		return a
	} else {
		return b
	}
}

func Max64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	} else if a == b {
		return a
	} else {
		return b
	}
}
