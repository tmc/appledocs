// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ExclusionDescriptors() []FontDescriptor
	SetExclusionDescriptors(value []FontDescriptor)
	QueryDescriptors() []FontDescriptor
	SetQueryDescriptors(value []FontDescriptor)
	// methods:
	AddQueryForDescriptors(descriptors []FontDescriptor)
	RemoveQueryForDescriptors(descriptors []FontDescriptor)
}

// A mutable collection of font descriptors taken together as a single object.
//
// You can use this class to modify the search queries for the font descriptors used by the parent class.


// A mutable collection of font descriptors taken together as a single object.
//
// [Full Topic]
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



// Creates a mutable font collection containing the fonts that match the specified font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(descriptors:)
func NewMutableFontCollectionWithDescriptors(queryDescriptors []FontDescriptor) MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(getMutableFontCollectionClass().class), objc.Sel("fontCollectionWithDescriptors:"), queryDescriptors)
	return rv
}


// Creates a mutable font collection containing fonts suitable for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(locale:)
func NewMutableFontCollectionWithLocale(locale foundation.Locale) MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(getMutableFontCollectionClass().class), objc.Sel("fontCollectionWithLocale:"), locale)
	return rv
}


// Creates a mutable named font collection object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:)
func NewMutableFontCollectionWithName(name objc.IObject /* cross-framework: FontCollectionName */) MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(getMutableFontCollectionClass().class), objc.Sel("fontCollectionWithName:"), name)
	return rv
}


// Creates a mutable font collection with the specified name and font visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:visibility:)
func NewMutableFontCollectionWithNameVisibility(name objc.IObject /* cross-framework: FontCollectionName */, visibility FontCollectionVisibility) MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(getMutableFontCollectionClass().class), objc.Sel("fontCollectionWithName:visibility:"), name, visibility)
	return rv
}



// Creates a mutable font collection containing the fonts that match the specified font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(descriptors:)
func (mc _MutableFontCollectionClass) FontCollectionWithDescriptors(queryDescriptors []FontDescriptor) IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithDescriptors:"), queryDescriptors)
	return rv
}


// Creates a mutable font collection containing fonts suitable for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(locale:)
func (mc _MutableFontCollectionClass) FontCollectionWithLocale(locale foundation.Locale) IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithLocale:"), locale)
	return rv
}


// Creates a mutable named font collection object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:)
func (mc _MutableFontCollectionClass) FontCollectionWithName(name objc.IObject /* cross-framework: FontCollectionName */) IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithName:"), name)
	return rv
}


// Creates a mutable font collection with the specified name and font visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:visibility:)
func (mc _MutableFontCollectionClass) FontCollectionWithNameVisibility(name objc.IObject /* cross-framework: FontCollectionName */, visibility FontCollectionVisibility) IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithName:visibility:"), name, visibility)
	return rv
}


// The mutable font collection that matches all registered fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/withAllAvailableDescriptors
func (mc _MutableFontCollectionClass) FontCollectionWithAllAvailableDescriptors() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithAllAvailableDescriptors"))
	return rv
}

// Edits the query and exclusion arrays by adding the specified font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/addQuery(for:)
func (m_ MutableFontCollection) AddQueryForDescriptors(descriptors []FontDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addQueryForDescriptors:"), descriptors)
}


// Edits the query and exclusion arrays by removing the specified font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/removeQuery(for:)
func (m_ MutableFontCollection) RemoveQueryForDescriptors(descriptors []FontDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeQueryForDescriptors:"), descriptors)
}


// The font descriptors to exclude from query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/exclusionDescriptors
func (m_ MutableFontCollection) ExclusionDescriptors() []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](m_.ID, objc.Sel("exclusionDescriptors"))
	return rv
}


// The font descriptors to exclude from query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/exclusionDescriptors
func (m_ MutableFontCollection) SetExclusionDescriptors(value []FontDescriptor) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setExclusionDescriptors:"), nsArray)
}


// The font descriptors to include in query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/queryDescriptors
func (m_ MutableFontCollection) QueryDescriptors() []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](m_.ID, objc.Sel("queryDescriptors"))
	return rv
}


// The font descriptors to include in query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/queryDescriptors
func (m_ MutableFontCollection) SetQueryDescriptors(value []FontDescriptor) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueryDescriptors:"), nsArray)
}


// The mutable font collection that matches all registered fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/withAllAvailableDescriptors
func (m_ MutableFontCollection) FontCollectionWithAllAvailableDescriptors() IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](m_.ID, objc.Sel("fontCollectionWithAllAvailableDescriptors"))
	return rv
}


