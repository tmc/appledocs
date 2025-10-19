// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [errorCorrectionLevel] class.
var errorCorrectionLevelClass = _errorCorrectionLevelClass{objc.GetClass("errorCorrectionLevel")}

type _errorCorrectionLevelClass struct {
	class objc.Class
}

// An interface definition for the [errorCorrectionLevel] class.
type IerrorCorrectionLevel interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectionLevel-c.ivar

type errorCorrectionLevel struct {
	objectivec.Object
}

// errorCorrectionLevelFrom constructs a [errorCorrectionLevel] from an unsafe.Pointer.
func errorCorrectionLevelFrom(ptr unsafe.Pointer) errorCorrectionLevel {
	return errorCorrectionLevel{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ec _errorCorrectionLevelClass) Alloc() errorCorrectionLevel {
	rv := objc.Send[errorCorrectionLevel](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ec _errorCorrectionLevelClass) New() errorCorrectionLevel {
	rv := objc.Send[errorCorrectionLevel](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ errorCorrectionLevel) Init() errorCorrectionLevel {
	rv := objc.Send[errorCorrectionLevel](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ errorCorrectionLevel) Autorelease() errorCorrectionLevel {
	rv := objc.Send[errorCorrectionLevel](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewerrorCorrectionLevel creates a new errorCorrectionLevel instance.
func NewerrorCorrectionLevel() errorCorrectionLevel {
	return errorCorrectionLevelClass.New()
}




