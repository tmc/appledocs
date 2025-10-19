// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [rightEyeClosed] class.
var (
	rightEyeClosedClass     _rightEyeClosedClass
	rightEyeClosedClassOnce sync.Once
)

func getrightEyeClosedClass() _rightEyeClosedClass {
	rightEyeClosedClassOnce.Do(func() {
		rightEyeClosedClass = _rightEyeClosedClass{objc.GetClass("rightEyeClosed")}
	})
	return rightEyeClosedClass
}

type _rightEyeClosedClass struct {
	class objc.Class
}

// An interface definition for the [rightEyeClosed] class.
type IrightEyeClosed interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyeClosed-c.ivar
type rightEyeClosed struct {
	objectivec.Object
}

// rightEyeClosedFrom constructs a [rightEyeClosed] from an unsafe.Pointer.
func rightEyeClosedFrom(ptr unsafe.Pointer) rightEyeClosed {
	return rightEyeClosed{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _rightEyeClosedClass) Alloc() rightEyeClosed {
	rv := objc.Send[rightEyeClosed](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _rightEyeClosedClass) New() rightEyeClosed {
	rv := objc.Send[rightEyeClosed](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ rightEyeClosed) Init() rightEyeClosed {
	rv := objc.Send[rightEyeClosed](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ rightEyeClosed) Autorelease() rightEyeClosed {
	rv := objc.Send[rightEyeClosed](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrightEyeClosed creates a new rightEyeClosed instance.
func NewrightEyeClosed() rightEyeClosed {
	return getrightEyeClosedClass().New()
}




