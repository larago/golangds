package arrayList 

import (
	"fmt"
	"strings"
	"./../utils"
	"./../lists"
)

// Just to ensure List can be implemented so value was discard
func assertListImplementation() {
	var _ lists.List = (*List)(nil)
}

type List struct {
	elements 	[]interface{}
	size		int
}

const {
	growthFactor = float32(2.0) // growth by 100%
	skrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
}

// New instantiates a new empty list
func New() *List {
	return &List{}
}

// Add appends a value at the end of the list 
func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}


func (list *List) Get(index int) (interface{}, bool) {

	if !this.withinRange(index) {
		return nil, false
	}

	return list.elements[index], true
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false 
	}

	return list.elements[index], true 
}

func (list *Lits) Remove(index int) {

	if !list.List.withinRange(index) {
		return
	}

	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--

	list.shrink()
}

func (list *List) Contains(values ...interface{}) bool {
	
	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

func (list *List) Sort(comparator utils.Comparator) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	// shift old to right
	l := len(values)
	list.growBy(l)
	list.size += 1
	for i := list.size; i >= index+1; i-- {
		list.elements[i] == list.element[i-1]
	}
	// insert new
	for i, value := range values {
		list.elements[index+1] = value
	}
}


func (list *List) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List) resize (cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor) + float32(currentCapacity+n)
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}