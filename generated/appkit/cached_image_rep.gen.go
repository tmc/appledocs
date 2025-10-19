// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CachedImageRep] class.
var (
	cachedImageRepClass     _CachedImageRepClass
	cachedImageRepClassOnce sync.Once
)

func getCachedImageRepClass() _CachedImageRepClass {
	cachedImageRepClassOnce.Do(func() {
		cachedImageRepClass = _CachedImageRepClass{objc.GetClass("NSCachedImageRep")}
	})
	return cachedImageRepClass
}

type _CachedImageRepClass struct {
	class objc.Class
}

// An interface definition for the [CachedImageRep] class.
type ICachedImageRep interface {
	IImageRep
}

// An object that stores image data in a form that can be readily transferred to the screen. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (cc _CachedImageRepClass) Alloc() CachedImageRep {
	rv := objc.Send[CachedImageRep](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




