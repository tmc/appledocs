// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [CachedImageRep] class.
var (
	CachedImageRepClass     _CachedImageRepClass
	CachedImageRepClassOnce sync.Once
)

func getCachedImageRepClass() _CachedImageRepClass {
	CachedImageRepClassOnce.Do(func() {
		CachedImageRepClass = _CachedImageRepClass{objc.GetClass("NSCachedImageRep")}
	})
	return CachedImageRepClass
}

type _CachedImageRepClass struct {
	class objc.Class
}





// An interface definition for the [CachedImageRep] class.
type ICachedImageRep interface {
	IImageRep
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CachedImageRepClass) Alloc() CachedImageRep {
	rv := objc.Send[CachedImageRep](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CachedImageRepClass) New() CachedImageRep {
	rv := objc.Send[CachedImageRep](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CachedImageRep) Init() CachedImageRep {
	rv := objc.Send[CachedImageRep](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CachedImageRep) Autorelease() CachedImageRep {
	rv := objc.Send[CachedImageRep](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCachedImageRep creates a new CachedImageRep instance.
func NewCachedImageRep() CachedImageRep {
	return getCachedImageRepClass().New()
}





// An object that stores image data in a form that can be readily transferred to the screen.
//
// An object differs from other image representation objects in that it simply stores the already rendered image, whereas other image representation objects generally have knowledge about how to render the image from source data. You typically do not use this class directly. Instead, and its other image representation objects create instances of as needed to cache versions of the rendered image. This caching speeds up screen-based drawing for existing images during subsequent rendering operations. Cached image representations are also used to capture drawing commands for images created programmatically by locking focus on an image.


// An object that stores image data in a form that can be readily transferred to the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCachedImageRep
type CachedImageRep struct {
	ImageRep
}

// CachedImageRepFrom constructs a [CachedImageRep] from an unsafe.Pointer.
//
// An object that stores image data in a form that can be readily transferred to the screen.
func CachedImageRepFrom(ptr unsafe.Pointer) CachedImageRep {
	return CachedImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}






// Returns a cached image representation initialized with the specified image characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCachedImageRep/initWithSize:depth:separate:alpha:
func NewCachedImageRepWithSizeDepthSeparateAlpha(size corefoundation.CGSize, depth WindowDepth, flag bool, alpha bool) CachedImageRep {
	instance := getCachedImageRepClass().Alloc()
	rv := objc.Send[CachedImageRep](instance.ID, objc.Sel("initWithSize:depth:separate:alpha:"), size, depth, flag, alpha)
	rv.Autorelease()
	return rv
}


// Returns a cached image representation initialized for drawing in the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCachedImageRep/initWithWindow:rect:
func NewCachedImageRepWithWindowRect(win IWindow, rect corefoundation.CGRect) CachedImageRep {
	instance := getCachedImageRepClass().Alloc()
	rv := objc.Send[CachedImageRep](instance.ID, objc.Sel("initWithWindow:rect:"), win, rect)
	rv.Autorelease()
	return rv
}



























