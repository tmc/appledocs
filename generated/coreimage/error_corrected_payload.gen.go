// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [errorCorrectedPayload] class.
var (
	ErrorCorrectedPayloadClass     _errorCorrectedPayloadClass
	ErrorCorrectedPayloadClassOnce sync.Once
)

func geterrorCorrectedPayloadClass() _errorCorrectedPayloadClass {
	ErrorCorrectedPayloadClassOnce.Do(func() {
		ErrorCorrectedPayloadClass = _errorCorrectedPayloadClass{objc.GetClass("errorCorrectedPayload")}
	})
	return ErrorCorrectedPayloadClass
}

type _errorCorrectedPayloadClass struct {
	class objc.Class
}

// An interface definition for the [errorCorrectedPayload] class.
type IerrorCorrectedPayload interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/errorCorrectedPayload-c.ivar
type errorCorrectedPayload struct {
	objectivec.Object
}

// errorCorrectedPayloadFrom constructs a [errorCorrectedPayload] from an unsafe.Pointer.
func errorCorrectedPayloadFrom(ptr unsafe.Pointer) errorCorrectedPayload {
	return errorCorrectedPayload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _errorCorrectedPayloadClass) Alloc() errorCorrectedPayload {
	rv := objc.Send[errorCorrectedPayload](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _errorCorrectedPayloadClass) New() errorCorrectedPayload {
	rv := objc.Send[errorCorrectedPayload](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ errorCorrectedPayload) Init() errorCorrectedPayload {
	rv := objc.Send[errorCorrectedPayload](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ errorCorrectedPayload) Autorelease() errorCorrectedPayload {
	rv := objc.Send[errorCorrectedPayload](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewerrorCorrectedPayload creates a new errorCorrectedPayload instance.
func NewerrorCorrectedPayload() errorCorrectedPayload {
	return geterrorCorrectedPayloadClass().New()
}




