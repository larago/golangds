package utils

// Comparator will make type assertion (see IntComparator for example)
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//		negative ,  if a < b
//		zero 	 ,	if a == b
//		positive ,	if a > b

type Comparator func(a, b interface{}) int

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
	switch {
	case diff < 0:
		return -1
	case diff > 0:
		return 1
	// they are the same string
	default:
		return 0
	}
}

// IntComparator provides a basic comparasion on int
func IntComparator(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int8Comparator provides a basic comparasion on Int8
func Int8Comparator(a, b interface{}) int {
	aAsserted := a.(int8)
	bAsserted := b.(int8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int16Comparator provides a basic comparasion on Int16
func Int16Comparator(a, b interface{}) int {
	aAsserted := a.(int16)
	bAsserted := b.(int16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int32Comparator provides a basic comparasion on int32
func Int32Comparator(a, b interface{}) int {
	aAsserted := a.(int32)
	bAsserted := b.(int32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int64Comparator provides a basic comparasion on int64
func Int64Comparator(a, b interface{}) int {
	aAsserted := a.(int64)
	bAsserted := b.(int64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
