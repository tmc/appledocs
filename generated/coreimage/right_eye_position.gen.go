// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [rightEyePosition] class.
var (
	RightEyePositionClass     _rightEyePositionClass
	RightEyePositionClassOnce sync.Once
)

func getrightEyePositionClass() _rightEyePositionClass {
	RightEyePositionClassOnce.Do(func() {
		RightEyePositionClass = _rightEyePositionClass{objc.GetClass("rightEyePosition")}
	})
	return RightEyePositionClass
}

type _rightEyePositionClass struct {
	class objc.Class
}





// An interface definition for the [rightEyePosition] class.
type IrightEyePosition interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _rightEyePositionClass) Alloc() rightEyePosition {
	rv := objc.Send[rightEyePosition](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _rightEyePositionClass) New() rightEyePosition {
	rv := objc.Send[rightEyePosition](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ rightEyePosition) Init() rightEyePosition {
	rv := objc.Send[rightEyePosition](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ rightEyePosition) Autorelease() rightEyePosition {
	rv := objc.Send[rightEyePosition](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrightEyePosition creates a new rightEyePosition instance.
func NewrightEyePosition() rightEyePosition {
	return getrightEyePositionClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyePosition-c.ivar
type rightEyePosition struct {
	objectivec.Object
}

// rightEyePositionFrom constructs a [rightEyePosition] from an unsafe.Pointer.
func rightEyePositionFrom(ptr unsafe.Pointer) rightEyePosition {
	return rightEyePosition{objectivec.Object{objc.ID(ptr)}}
}































