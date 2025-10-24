// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetalDisplayLinkUpdate] class.
var (
	MetalDisplayLinkUpdateClass     _MetalDisplayLinkUpdateClass
	MetalDisplayLinkUpdateClassOnce sync.Once
)

func getMetalDisplayLinkUpdateClass() _MetalDisplayLinkUpdateClass {
	MetalDisplayLinkUpdateClassOnce.Do(func() {
		MetalDisplayLinkUpdateClass = _MetalDisplayLinkUpdateClass{objc.GetClass("CAMetalDisplayLinkUpdate")}
	})
	return MetalDisplayLinkUpdateClass
}

type _MetalDisplayLinkUpdateClass struct {
	class objc.Class
}

// An interface definition for the [MetalDisplayLinkUpdate] class.
type IMetalDisplayLinkUpdate interface {
	objectivec.IObject
	// properties:
	TargetTimestamp() float64
	Drawable() MetalDrawable /* not a class type */
	SetDrawable(value MetalDrawable /* not a class type */)
	TargetPresentationTimestamp() float64
	SetTargetPresentationTimestamp(value float64)
	// methods:
}

// Stores information about a single update from a Metal display link instance.


// Stores information about a single update from a Metal display link instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update
type MetalDisplayLinkUpdate struct {
	objectivec.Object
}

// MetalDisplayLinkUpdateFrom constructs a [MetalDisplayLinkUpdate] from an unsafe.Pointer.
//
// Stores information about a single update from a Metal display link instance.
func MetalDisplayLinkUpdateFrom(ptr unsafe.Pointer) MetalDisplayLinkUpdate {
	return MetalDisplayLinkUpdate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetalDisplayLinkUpdateClass) Alloc() MetalDisplayLinkUpdate {
	rv := objc.Send[MetalDisplayLinkUpdate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetalDisplayLinkUpdateClass) New() MetalDisplayLinkUpdate {
	rv := objc.Send[MetalDisplayLinkUpdate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetalDisplayLinkUpdate) Init() MetalDisplayLinkUpdate {
	rv := objc.Send[MetalDisplayLinkUpdate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetalDisplayLinkUpdate) Autorelease() MetalDisplayLinkUpdate {
	rv := objc.Send[MetalDisplayLinkUpdate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetalDisplayLinkUpdate creates a new MetalDisplayLinkUpdate instance.
func NewMetalDisplayLinkUpdate() MetalDisplayLinkUpdate {
	return getMetalDisplayLinkUpdateClass().New()
}



// A deadline that indicates when your app needs to finish rendering to the drawable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update/targetTimestamp
func (m_ MetalDisplayLinkUpdate) TargetTimestamp() float64 {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("targetTimestamp"))
	return rv
}


// The Metal drawable your app uses to render the next frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametaldisplaylink/update/drawable
func (m_ MetalDisplayLinkUpdate) Drawable() MetalDrawable /* not a class type */ {
	rv := objc.Send[MetalDrawable](m_.ID, objc.Sel("drawable"))
	return rv
}


// The Metal drawable your app uses to render the next frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametaldisplaylink/update/drawable
func (m_ MetalDisplayLinkUpdate) SetDrawable(value MetalDrawable /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawable:"), value)
}


// The time the system estimates until the display of the next frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametaldisplaylink/update/targetpresentationtimestamp
func (m_ MetalDisplayLinkUpdate) TargetPresentationTimestamp() float64 {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("targetPresentationTimestamp"))
	return rv
}


// The time the system estimates until the display of the next frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametaldisplaylink/update/targetpresentationtimestamp
func (m_ MetalDisplayLinkUpdate) SetTargetPresentationTimestamp(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetPresentationTimestamp:"), value)
}



