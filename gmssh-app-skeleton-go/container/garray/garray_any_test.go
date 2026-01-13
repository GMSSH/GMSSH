package garray

import (
	"reflect"
	"testing"
)

func TestStrArray(t *testing.T) {
	strArray := NewDefaultArray[string]()
	strArray.Append("1")
	strArray.Append("2")
	strArray.Append("3")

	tests := []string{"1", "2", "3"}
	if !reflect.DeepEqual(strArray.Array(), tests) {
		t.Errorf("excepted:%#v,  got:%#v", strArray.Array(), tests)
	}
}

func TestIntArray(t *testing.T) {
	strArray := NewDefaultArray[int]()
	strArray.Append(1)
	strArray.Append(2)
	strArray.Append(3)

	tests := []int{1, 2, 3}
	if !reflect.DeepEqual(strArray.Array(), tests) {
		t.Errorf("excepted:%#v,  got:%#v", strArray.Array(), tests)
	}
}
