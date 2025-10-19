// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableFontCollection] class.
var (
	mutableFontCollectionClass     _MutableFontCollectionClass
	mutableFontCollectionClassOnce sync.Once
)

func getMutableFontCollectionClass() _MutableFontCollectionClass {
	mutableFontCollectionClassOnce.Do(func() {
		mutableFontCollectionClass = _MutableFontCollectionClass{objc.GetClass("NSMutableFontCollection")}
	})
	return mutableFontCollectionClass
}

type _MutableFontCollectionClass struct {
	class objc.Class
}

// An interface definition for the [MutableFontCollection] class.
type IMutableFontCollection interface {
	IFontCollection
}

// A mutable collection of font descriptors taken together as a single object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection

type MutableFontCollection struct {
	FontCollection
}

// MutableFontCollectionFrom constructs a [MutableFontCollection] from an unsafe.Pointer.
//
// A mutable collection of font descriptors taken together as a single object.
func MutableFontCollectionFrom(ptr unsafe.Pointer) MutableFontCollection {
	return MutableFontCollection{
		FontCollection: FontCollectionFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (mc _MutableFontCollectionClass) Alloc() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableFontCollectionClass) New() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableFontCollection) Init() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableFontCollection) Autorelease() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableFontCollection creates a new MutableFontCollection instance.
func NewMutableFontCollection() MutableFontCollection {
	return getMutableFontCollectionClass().New()
}




