package golang

import (
	"testing"
	"reflect"
)

func TestFizzBuzzFirst *testing.T) {
	original := map[int]int{
		1: 1,
		2: 2,
	}
	expected := map[int]int{
		1: 1,
		2: 2,
	}

	result := fizzbuzz(original)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}
}

// func TestFizzBuzz *testing.T) {
// 	original := map[int]int{
// 		1: 1,
// 		2: 2,
// 		3: 3,
// 		4: 4,
// 		5: 5,
// 		6: 6,
// 		7: 7,
// 		8: 8,
// 		9: 9,
// 		10: 10,
// 		15: 15,
// 		30: 30,
// 	}
// 	expected := map[int]int{
// 		1: 1,
// 		2: 2,
// 		3: 3,
// 		4: 4,
// 		5: 5,
// 		6: 6,
// 		7: 7,
// 		8: 8,
// 		9: 9,
// 		10: 10,
// 		15: 15,
// 		30: 30,
// 	}

// 	result := fizzbuzz(original)

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("%v is expected but got %v", expected, result)
// 	}
// }
