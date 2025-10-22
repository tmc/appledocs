// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [leftEyeClosed] class.
var (
	LeftEyeClosedClass     _leftEyeClosedClass
	LeftEyeClosedClassOnce sync.Once
)

func getleftEyeClosedClass() _leftEyeClosedClass {
	LeftEyeClosedClassOnce.Do(func() {
		LeftEyeClosedClass = _leftEyeClosedClass{objc.GetClass("leftEyeClosed")}
	})
	return LeftEyeClosedClass
}

type _leftEyeClosedClass struct {
	class objc.Class
}

// An interface definition for the [leftEyeClosed] class.
type IleftEyeClosed interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyeClosed-c.ivar

type leftEyeClosed struct {
	objectivec.Object
}

// leftEyeClosedFrom constructs a [leftEyeClosed] from an unsafe.Pointer.
func leftEyeClosedFrom(ptr unsafe.Pointer) leftEyeClosed {
	return leftEyeClosed{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _leftEyeClosedClass) Alloc() leftEyeClosed {
	rv := objc.Send[leftEyeClosed](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _leftEyeClosedClass) New() leftEyeClosed {
	rv := objc.Send[leftEyeClosed](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ leftEyeClosed) Init() leftEyeClosed {
	rv := objc.Send[leftEyeClosed](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ leftEyeClosed) Autorelease() leftEyeClosed {
	rv := objc.Send[leftEyeClosed](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewleftEyeClosed creates a new leftEyeClosed instance.
func NewleftEyeClosed() leftEyeClosed {
	return getleftEyeClosedClass().New()
}




