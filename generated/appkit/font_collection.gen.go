// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontCollection] class.
var (
	fontCollectionClass     _FontCollectionClass
	fontCollectionClassOnce sync.Once
)

func getFontCollectionClass() _FontCollectionClass {
	fontCollectionClassOnce.Do(func() {
		fontCollectionClass = _FontCollectionClass{objc.GetClass("NSFontCollection")}
	})
	return fontCollectionClass
}

type _FontCollectionClass struct {
	class objc.Class
}

// An interface definition for the [FontCollection] class.
type IFontCollection interface {
	objectivec.IObject
}

// A font collection, which is a group of font descriptors taken together as a single object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection
type FontCollection struct {
	objectivec.Object
}

// FontCollectionFrom constructs a [FontCollection] from an unsafe.Pointer.
//
// A font collection, which is a group of font descriptors taken together as a single object.
func FontCollectionFrom(ptr unsafe.Pointer) FontCollection {
	return FontCollection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FontCollectionClass) Alloc() FontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FontCollectionClass) New() FontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontCollection) Init() FontCollection {
	rv := objc.Send[FontCollection](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontCollection) Autorelease() FontCollection {
	rv := objc.Send[FontCollection](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontCollection creates a new FontCollection instance.
func NewFontCollection() FontCollection {
	return getFontCollectionClass().New()
}




