// Code generated by "stringer -type Enumerated"; DO NOT EDIT

package main

import "fmt"

const (
	_Enumerated_name_0 = "A"
	_Enumerated_name_1 = "B"
	_Enumerated_name_2 = "C"
	_Enumerated_name_3 = "D"
)

var (
	_Enumerated_index_0 = [...]uint8{0, 1}
	_Enumerated_index_1 = [...]uint8{0, 1}
	_Enumerated_index_2 = [...]uint8{0, 1}
	_Enumerated_index_3 = [...]uint8{0, 1}
)

func (i Enumerated) String() string {
	switch {
	case i == 2:
		return _Enumerated_name_0
	case i == 4:
		return _Enumerated_name_1
	case i == 8:
		return _Enumerated_name_2
	case i == 16:
		return _Enumerated_name_3
	default:
		return fmt.Sprintf("Enumerated(%d)", i)
	}
}