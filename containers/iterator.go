package containers

type InteratorWithIndex interface {
	Next() bool
	Value() interface{}
	Index() int
	Begin()
	First() bool
}

type IteratorWithKey interface {
	Next() bool
	Value() interface{}
	Key() interface{}
	Begin()
	First() bool
}
