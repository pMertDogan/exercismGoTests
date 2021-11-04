package cards

func isIndexInRange(slice []int, indexToCheck int) bool {
	if len(slice)-1 >= indexToCheck && indexToCheck >= 0 {

		return true
	} else {
		return false
	}
}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if isIndexInRange(slice, index) {
		return slice[index], true
	} else {
		return 0, false
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if isIndexInRange(slice, index) {
		slice[index] = value
		return slice
	} else {
		return append(slice, value)
	}
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length >= 1 {

		a := make([]int, length)
		for i := range a {
			a[i] = value
		}
		return a
	}
	var b []int
	return b

}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {

	if index < len(slice) && index >= 0 {
	return append(slice[:index], slice[index+1:]...)
	}else {
		// 	return slice
		// }
	// if index < len(slice) && index >= 0 {
	// 	var b []int
	// 	b = append(b, slice[:index]...)

	// 	b = append(b, slice[index+1:]...)
	// 	return b
	// } else {
	// 	return slice
	// }
	return slice
}

}
