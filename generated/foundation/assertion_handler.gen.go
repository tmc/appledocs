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

// An interface definition for the [AssertionHandler] class.
type IAssertionHandler interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (ac _AssertionHandlerClass) Alloc() AssertionHandler {
	rv := objc.Send[AssertionHandler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AssertionHandlerClass) New() AssertionHandler {
	rv := objc.Send[AssertionHandler](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssertionHandler) Init() AssertionHandler {
	rv := objc.Send[AssertionHandler](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssertionHandler) Autorelease() AssertionHandler {
	rv := objc.Send[AssertionHandler](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssertionHandler creates a new AssertionHandler instance.
func NewAssertionHandler() AssertionHandler {
	return assertionHandlerClass.New()
}




