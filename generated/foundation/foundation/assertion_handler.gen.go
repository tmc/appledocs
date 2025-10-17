// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AssertionHandler] class.
var AssertionHandlerClass objc.Class

func init() {
	AssertionHandlerClass = objc.GetClass("NSAssertionHandler")
}

type AssertionHandler struct {
	objc.ID
}

func AssertionHandlerFrom(ptr unsafe.Pointer) AssertionHandler {
	return AssertionHandler{
		ID: objc.ID(ptr),
	}
}




