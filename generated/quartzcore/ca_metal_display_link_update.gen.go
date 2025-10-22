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
	Drawable() objc.ID
	TargetPresentationTimestamp() TimeInterval
	TargetTimestamp() TimeInterval
}

// Stores information about a single update from a Metal display link instance.
//
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


// The Metal drawable your app uses to render the next frame.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update/drawable
func (m_ MetalDisplayLinkUpdate) Drawable() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("drawable"))
	return rv
}

// The time the system estimates until the display of the next frame.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update/targetPresentationTimestamp
func (m_ MetalDisplayLinkUpdate) TargetPresentationTimestamp() TimeInterval {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("targetPresentationTimestamp"))
	return rv
}

// A deadline that indicates when your app needs to finish rendering to the drawable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update/targetTimestamp
func (m_ MetalDisplayLinkUpdate) TargetTimestamp() TimeInterval {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("targetTimestamp"))
	return rv
}



