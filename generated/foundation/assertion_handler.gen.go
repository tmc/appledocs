// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssertionHandler] class.
var assertionHandlerClass = _AssertionHandlerClass{objc.GetClass("NSAssertionHandler")}

type _AssertionHandlerClass struct {
	class objc.Class
}

// An object that logs an assertion to the console. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler

type AssertionHandler struct {
	objectivec.Object
}

// AssertionHandlerFrom constructs a [AssertionHandler] from an unsafe.Pointer.
//
// An object that logs an assertion to the console.
func AssertionHandlerFrom(ptr unsafe.Pointer) AssertionHandler {
	return AssertionHandler{objectivec.Object{objc.ID(ptr)}}
}



