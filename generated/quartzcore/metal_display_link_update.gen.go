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
	metalDisplayLinkUpdateClass     _MetalDisplayLinkUpdateClass
	metalDisplayLinkUpdateClassOnce sync.Once
)

func getMetalDisplayLinkUpdateClass() _MetalDisplayLinkUpdateClass {
	metalDisplayLinkUpdateClassOnce.Do(func() {
		metalDisplayLinkUpdateClass = _MetalDisplayLinkUpdateClass{objc.GetClass("CAMetalDisplayLinkUpdate")}
	})
	return metalDisplayLinkUpdateClass
}

type _MetalDisplayLinkUpdateClass struct {
	class objc.Class
}

// An interface definition for the [MetalDisplayLinkUpdate] class.
type IMetalDisplayLinkUpdate interface {
	objectivec.IObject
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


// A deadline that indicates when your app needs to finish rendering to the drawable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update/targetTimestamp
func (m_ MetalDisplayLinkUpdate) TargetTimestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targetTimestamp"))
	return rv
}



