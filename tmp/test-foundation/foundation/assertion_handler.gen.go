// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var AssertionHandlerClass _AssertionHandlerClass

func init() {
	AssertionHandlerClass = _AssertionHandlerClass{objc.GetClass("NSAssertionHandler")}
}

type _AssertionHandlerClass struct {
	class objc.Class
}

type AssertionHandler struct {
	objc.ID
}

func AssertionHandlerFrom(ptr unsafe.Pointer) AssertionHandler {
	return AssertionHandler{
		ID: objc.ID(ptr),
	}
}




