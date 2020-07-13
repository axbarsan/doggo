package object

import (
	"strconv"
	"testing"
)

func TestStringMapKey(t *testing.T) {
	testCases := []struct {
		val1  Mappable
		val2  Mappable
		diff1 Mappable
		diff2 Mappable
	}{
		{
			val1:  &String{Value: "Hello World"},
			val2:  &String{Value: "Hello World"},
			diff1: &String{Value: "My name is johnny"},
			diff2: &String{Value: "My name is johnny"},
		},
		{
			val1:  &Integer{Value: 1},
			val2:  &Integer{Value: 1},
			diff1: &Integer{Value: 10},
			diff2: &Integer{Value: 10},
		},
		{
			val1:  &Boolean{Value: true},
			val2:  &Boolean{Value: true},
			diff1: &Boolean{Value: false},
			diff2: &Boolean{Value: false},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if tc.val1.MapKey() != tc.val2.MapKey() {
				t.Errorf("Case %d: Variables with same content have different hash keys", i)
			}

			if tc.diff1.MapKey() != tc.diff2.MapKey() {
				t.Errorf("Case %d: Variables with same content have different hash keys", i)
			}

			if tc.val1.MapKey() == tc.diff1.MapKey() {
				t.Errorf("Case %d: Variables with different content have same hash keys", i)
			}
		})
	}
}
