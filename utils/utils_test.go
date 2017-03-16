package utils

import (
	"strings"
	"testing"
)

func TestTpStringInts(t *testing.T) {

	var value interface{}

	value = int8(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = int16(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = int32(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = int64(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = rune(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestToStringUints(t *testing.T) {

	var value interface{}

	value = int8(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = int16(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = int32(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = int64(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = rune(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestToStringUints(t *testing.T) {
	var value interface{}

	value = uint8(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = uint16(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = uint32(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = uint64(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = byte(1)
	if actualValue, expectedValue := ToString(value), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestToStringFloats(t *testing.T) {
	var value interface{}

	value = float32(1.123456)
	if actualValue, expectedValue := ToString(value), "1.123456"; !strings.HasPrefix(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	value = float64(1.123456)
	if actualValue, expectedValue := ToString(value), "1.123456"; !strings.HasPrefix(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestToStringOther(t *testing.T) {
	var value interface{}

	value = "abc"
	if actualValue, expectedValue := ToString(value), "abc"; actualValue != expectedValue {
		t.Errorf("Got ", args ...interface{})
	}

	value = true
	if actualValue, expectedValue := ToString(value), "abc"; actualValue != expectedValue {
		t.Errorf("Got ", args ...interface{})
	}

	type T struct {
		id	int
		name	string
	}

	if actualValue, expectedValue := ToString(T{1, "abc"}), "{id:1 name:abc}"; actualValue != expectedValue {
		t.Errorf("Got ", args ...interface{})
	}
}
