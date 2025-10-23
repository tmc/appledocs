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
	FontCollectionClass     _FontCollectionClass
	FontCollectionClassOnce sync.Once
)

func getFontCollectionClass() _FontCollectionClass {
	FontCollectionClassOnce.Do(func() {
		FontCollectionClass = _FontCollectionClass{objc.GetClass("NSFontCollection")}
	})
	return FontCollectionClass
}

type _FontCollectionClass struct {
	class objc.Class
}

// An interface definition for the [FontCollection] class.
type IFontCollection interface {
	objectivec.IObject
	// properties:
	ExclusionDescriptors() IFontDescriptor
	SetExclusionDescriptors(value IFontDescriptor)
	MatchingDescriptors() IFontDescriptor
	SetMatchingDescriptors(value IFontDescriptor)
	QueryDescriptors() IFontDescriptor
	SetQueryDescriptors(value IFontDescriptor)
	// methods:
}

// A font collection, which is a group of font descriptors taken together as a single object.
//
// You can publicize the font collection as a named collection and it is presented through the System user interface such as the font panel and Font Book. The queries can be modified using the subclass.


// A font collection, which is a group of font descriptors taken together as a single object.
//
// [Full Topic]
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



// A list of query font descriptors whose matching results are excluded from the list of matching descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/exclusiondescriptors
func (f_ FontCollection) ExclusionDescriptors() IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("exclusionDescriptors"))
	return rv
}


// A list of query font descriptors whose matching results are excluded from the list of matching descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/exclusiondescriptors
func (f_ FontCollection) SetExclusionDescriptors(value IFontDescriptor) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setExclusionDescriptors:"), value)
}


// An array of font descriptors matching the logical descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/matchingdescriptors
func (f_ FontCollection) MatchingDescriptors() IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("matchingDescriptors"))
	return rv
}


// An array of font descriptors matching the logical descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/matchingdescriptors
func (f_ FontCollection) SetMatchingDescriptors(value IFontDescriptor) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMatchingDescriptors:"), value)
}


// An array of font descriptors whose matching results produce the collection’s matching descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/querydescriptors
func (f_ FontCollection) QueryDescriptors() IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("queryDescriptors"))
	return rv
}


// An array of font descriptors whose matching results produce the collection’s matching descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontcollection/querydescriptors
func (f_ FontCollection) SetQueryDescriptors(value IFontDescriptor) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setQueryDescriptors:"), value)
}



