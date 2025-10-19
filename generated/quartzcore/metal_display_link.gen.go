// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetalDisplayLink] class.
var metalDisplayLinkClass = _MetalDisplayLinkClass{objc.GetClass("CAMetalDisplayLink")}

type _MetalDisplayLinkClass struct {
	class objc.Class
}

// An interface definition for the [MetalDisplayLink] class.
type IMetalDisplayLink interface {
	objectivec.IObject
}

// A class your Metal app uses to register for callbacks to synchronize its animations for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink

type MetalDisplayLink struct {
	objectivec.Object
}

// MetalDisplayLinkFrom constructs a [MetalDisplayLink] from an unsafe.Pointer.
//
// A class your Metal app uses to register for callbacks to synchronize its animations for a display.
func MetalDisplayLinkFrom(ptr unsafe.Pointer) MetalDisplayLink {
	return MetalDisplayLink{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MetalDisplayLinkClass) Alloc() MetalDisplayLink {
	rv := objc.Send[MetalDisplayLink](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MetalDisplayLinkClass) New() MetalDisplayLink {
	rv := objc.Send[MetalDisplayLink](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetalDisplayLink) Init() MetalDisplayLink {
	rv := objc.Send[MetalDisplayLink](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetalDisplayLink) Autorelease() MetalDisplayLink {
	rv := objc.Send[MetalDisplayLink](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetalDisplayLink creates a new MetalDisplayLink instance.
func NewMetalDisplayLink() MetalDisplayLink {
	return metalDisplayLinkClass.New()
}




