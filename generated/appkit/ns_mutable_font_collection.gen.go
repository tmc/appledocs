// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableFontCollection] class.
var (
	MutableFontCollectionClass     _MutableFontCollectionClass
	MutableFontCollectionClassOnce sync.Once
)

func getMutableFontCollectionClass() _MutableFontCollectionClass {
	MutableFontCollectionClassOnce.Do(func() {
		MutableFontCollectionClass = _MutableFontCollectionClass{objc.GetClass("NSMutableFontCollection")}
	})
	return MutableFontCollectionClass
}

type _MutableFontCollectionClass struct {
	class objc.Class
}

// An interface definition for the [MutableFontCollection] class.
type IMutableFontCollection interface {
	IFontCollection
}

// A mutable collection of font descriptors taken together as a single object.
//
// You can use this class to modify the search queries for the font descriptors used by the parent class.
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


// The font descriptors to exclude from query results.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection/exclusiondescriptors
func (m_ MutableFontCollection) ExclusionDescriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("exclusionDescriptors"))
	return rv
}


// SetExclusionDescriptors sets the value of the exclusionDescriptors property.
// The font descriptors to exclude from query results.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection/exclusiondescriptors
func (m_ MutableFontCollection) SetExclusionDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExclusionDescriptors:"), value)
}

// The font descriptors to include in query results.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection/querydescriptors
func (m_ MutableFontCollection) QueryDescriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("queryDescriptors"))
	return rv
}


// SetQueryDescriptors sets the value of the queryDescriptors property.
// The font descriptors to include in query results.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutablefontcollection/querydescriptors
func (m_ MutableFontCollection) SetQueryDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueryDescriptors:"), value)
}



