// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyMapOfStringToInterface() map[string]interface{} {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(map[string]interface{}))(nil)).Elem()))
	var nullValue map[string]interface{}
	return nullValue
}

func EqMapOfStringToInterface(value map[string]interface{}) map[string]interface{} {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue map[string]interface{}
	return nullValue
}

func NotEqMapOfStringToInterface(value map[string]interface{}) map[string]interface{} {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue map[string]interface{}
	return nullValue
}

func MapOfStringToInterfaceThat(matcher pegomock.ArgumentMatcher) map[string]interface{} {
	pegomock.RegisterMatcher(matcher)
	var nullValue map[string]interface{}
	return nullValue
}
