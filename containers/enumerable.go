package containers

type EnumerableWithIndex interface{
	// Each call the given function once for each element, passing that element's index and value
	Each(func(index int, value interface{}))

	Any(func(index int, value interface{}) bool) bool

	All(func(index int, value interface{}) bool) bool
 
	Find(func(index int, value interface{}) bool) (int, interface{})
}

type EnumerableWithKey interface {
	Each(func(key interface{}, value interface{}))

	Any(func(key interface{}, value interface{}) bool) bool

	All(func(key interface{}, value interface{}) bool) bool 

	Find(func(key interface{}, value interface{}) bool) (interface{}, interface{})
}