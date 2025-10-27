// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Touch] class.
var (
	TouchClass     _TouchClass
	TouchClassOnce sync.Once
)

func getTouchClass() _TouchClass {
	TouchClassOnce.Do(func() {
		TouchClass = _TouchClass{objc.GetClass("NSTouch")}
	})
	return TouchClass
}

type _TouchClass struct {
	class objc.Class
}





// An interface definition for the [Touch] class.
type ITouch interface {
	objectivec.IObject
	

	// properties:
	Identity() unsafe.Pointer
	Resting() bool
	NormalizedPosition() corefoundation.CGPoint
	Phase() TouchPhase
	Type() TouchType
	DeviceSize() corefoundation.CGSize
	SetDeviceSize(value corefoundation.CGSize)
	IsResting() bool
	SetIsResting(value bool)


	

	// methods:
	PreviousLocationInView(view IView) corefoundation.CGPoint


}





// Alloc allocates a new instance without initialization.
func (tc _TouchClass) Alloc() Touch {
	rv := objc.Send[Touch](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TouchClass) New() Touch {
	rv := objc.Send[Touch](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Touch) Init() Touch {
	rv := objc.Send[Touch](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Touch) Autorelease() Touch {
	rv := objc.Send[Touch](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTouch creates a new Touch instance.
func NewTouch() Touch {
	return getTouchClass().New()
}





// A snapshot of a particular touch at an instant in time.
//
// A touch event is not persistent throughout the touch. A touch creates new instances as it progresses. Use the identity property to follow a specific touch across its lifetime. Touches do not have a corresponding screen location. The first touch of a touch collection latches to the view underlying the cursor using the same hit detection as mouse events. Additional touches on the same device latch to the same view. Latches remain on views until the user ends a touch or an event cancels it.


// A snapshot of a particular touch at an instant in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch
type Touch struct {
	objectivec.Object
}

// TouchFrom constructs a [Touch] from an unsafe.Pointer.
//
// A snapshot of a particular touch at an instant in time.
func TouchFrom(ptr unsafe.Pointer) Touch {
	return Touch{objectivec.Object{objc.ID(ptr)}}
}




















// Indicates the previous location of the touch in the view’s coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/previousLocation(in:)
func (t_ Touch) PreviousLocationInView(view IView) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t_.ID, objc.Sel("previousLocationInView:"), view)
	return rv
}







// The changes to a particular touch during its lifetime.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/identity
func (t_ Touch) Identity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("identity"))
	return rv
}


// The indicator for a resting touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/isResting
func (t_ Touch) Resting() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resting"))
	return rv
}


// The normalized position of the touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/normalizedPosition
func (t_ Touch) NormalizedPosition() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t_.ID, objc.Sel("normalizedPosition"))
	return rv
}


// The current phase of the touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/phase-swift.property
func (t_ Touch) Phase() TouchPhase {
	rv := objc.Send[TouchPhase](t_.ID, objc.Sel("phase"))
	return rv
}


// A type of touch from a Touch Bar interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/type
func (t_ Touch) Type() TouchType {
	rv := objc.Send[TouchType](t_.ID, objc.Sel("type"))
	return rv
}


// The range of the touch device in points, such as 72 ppi.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/devicesize
func (t_ Touch) DeviceSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t_.ID, objc.Sel("deviceSize"))
	return rv
}


// The range of the touch device in points, such as 72 ppi.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/devicesize
func (t_ Touch) SetDeviceSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDeviceSize:"), value)
}


// The indicator for a resting touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/isresting
func (t_ Touch) IsResting() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isResting"))
	return rv
}


// The indicator for a resting touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/isresting
func (t_ Touch) SetIsResting(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsResting:"), value)
}








