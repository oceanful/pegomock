// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyMapOfIntToInt() map[int]int {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(map[int]int))(nil)).Elem()))
	var nullValue map[int]int
	return nullValue
}

func EqMapOfIntToInt(value map[int]int) map[int]int {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue map[int]int
	return nullValue
}

func NotEqMapOfIntToInt(value map[int]int) map[int]int {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue map[int]int
	return nullValue
}

func MapOfIntToIntThat(matcher pegomock.ArgumentMatcher) map[int]int {
	pegomock.RegisterMatcher(matcher)
	var nullValue map[int]int
	return nullValue
}
