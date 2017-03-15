package utils

import "time"

// Comparator will make type assertion (see IntComparator for example)
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//		negative ,  if a < b
//		zero 	 ,	if a == b
//		positive ,	if a > b

type Comparator func(a, b interfave{}) int 

func StringComparator(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	// check minmum letters if they are same
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	// if above check passes, check their length
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	// return val according to diff
	switch diff {
		case diff < 0:
		return -1
		case diff > 0:
		return 1
		// they are the same string
		default:
		return 0
	}
}