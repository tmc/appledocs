// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CustomImageRep] class.
var (
	CustomImageRepClass     _CustomImageRepClass
	CustomImageRepClassOnce sync.Once
)

func getCustomImageRepClass() _CustomImageRepClass {
	CustomImageRepClassOnce.Do(func() {
		CustomImageRepClass = _CustomImageRepClass{objc.GetClass("NSCustomImageRep")}
	})
	return CustomImageRepClass
}

type _CustomImageRepClass struct {
	class objc.Class
}

// An interface definition for the [CustomImageRep] class.
type ICustomImageRep interface {
	IImageRep
}

// An object that uses a delegate object to render an image from a custom format.
//
// When called upon to produce an image, an sends a message to its delegate to do the actual drawing. You can use this class to support custom image formats without going to the trouble of subclassing directly.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep
type CustomImageRep struct {
	ImageRep
}

// CustomImageRepFrom constructs a [CustomImageRep] from an unsafe.Pointer.
//
// An object that uses a delegate object to render an image from a custom format.
func CustomImageRepFrom(ptr unsafe.Pointer) CustomImageRep {
	return CustomImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomImageRepClass) Alloc() CustomImageRep {
	rv := objc.Send[CustomImageRep](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomImageRepClass) New() CustomImageRep {
	rv := objc.Send[CustomImageRep](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomImageRep) Init() CustomImageRep {
	rv := objc.Send[CustomImageRep](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomImageRep) Autorelease() CustomImageRep {
	rv := objc.Send[CustomImageRep](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomImageRep creates a new CustomImageRep instance.
func NewCustomImageRep() CustomImageRep {
	return getCustomImageRepClass().New()
}

// Returns a representation of an image initialized with the specified delegate information.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(draw:delegate:)
func NewCustomImageRepWithDrawSelectorDelegate(selector objc.SEL, delegate objc.ID) CustomImageRep {
	instance := getCustomImageRepClass().Alloc()
	rv := objc.Send[CustomImageRep](instance.ID, objc.Sel("initWithDrawSelector:delegate:"), selector, delegate)
	rv.Autorelease()
	return rv
}

// The delegate object that renders the image for the image representation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/delegate
func (c_ CustomImageRep) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}
