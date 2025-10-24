// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	ExclusionDescriptors() []FontDescriptor
	MatchingDescriptors() []FontDescriptor
	QueryDescriptors() []FontDescriptor
	// methods:
	MatchingDescriptorsForFamily(family objc.IObject /* cross-framework: NSString */) []FontDescriptor
	MatchingDescriptorsForFamilyOptions(family objc.IObject /* cross-framework: NSString */, options foundation.IDictionary) []FontDescriptor
	MatchingDescriptorsWithOptions(options foundation.IDictionary) []FontDescriptor
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



// Returns a font collection matching the given descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(descriptors:)
func NewFontCollectionWithDescriptors(queryDescriptors []FontDescriptor) FontCollection {
	rv := objc.Send[FontCollection](objc.ID(getFontCollectionClass().class), objc.Sel("fontCollectionWithDescriptors:"), queryDescriptors)
	return rv
}


// Returns a collection of fonts matching the given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(locale:)
func NewFontCollectionWithLocale(locale foundation.Locale) FontCollection {
	rv := objc.Send[FontCollection](objc.ID(getFontCollectionClass().class), objc.Sel("fontCollectionWithLocale:"), locale)
	return rv
}


// Creates a named font collection object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(name:)
func NewFontCollectionWithName(name objc.IObject /* cross-framework: FontCollectionName */) FontCollection {
	rv := objc.Send[FontCollection](objc.ID(getFontCollectionClass().class), objc.Sel("fontCollectionWithName:"), name)
	return rv
}


// Creates a font collection with the specified name and font visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(name:visibility:)
func NewFontCollectionWithNameVisibility(name objc.IObject /* cross-framework: FontCollectionName */, visibility FontCollectionVisibility) FontCollection {
	rv := objc.Send[FontCollection](objc.ID(getFontCollectionClass().class), objc.Sel("fontCollectionWithName:visibility:"), name, visibility)
	return rv
}



// Remove from view the named font collection with the specified visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/hide(withName:visibility:)
func (fc _FontCollectionClass) HideFontCollectionWithNameVisibilityError(name objc.IObject /* cross-framework: FontCollectionName */, visibility FontCollectionVisibility, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("hideFontCollectionWithName:visibility:error:"), name, visibility, error_)
	return rv
}


// Returns a font collection matching the given descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(descriptors:)
func (fc _FontCollectionClass) FontCollectionWithDescriptors(queryDescriptors []FontDescriptor) IFontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.class), objc.Sel("fontCollectionWithDescriptors:"), queryDescriptors)
	return rv
}


// Returns a collection of fonts matching the given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(locale:)
func (fc _FontCollectionClass) FontCollectionWithLocale(locale foundation.Locale) IFontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.class), objc.Sel("fontCollectionWithLocale:"), locale)
	return rv
}


// Creates a named font collection object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(name:)
func (fc _FontCollectionClass) FontCollectionWithName(name objc.IObject /* cross-framework: FontCollectionName */) IFontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.class), objc.Sel("fontCollectionWithName:"), name)
	return rv
}


// Creates a font collection with the specified name and font visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(name:visibility:)
func (fc _FontCollectionClass) FontCollectionWithNameVisibility(name objc.IObject /* cross-framework: FontCollectionName */, visibility FontCollectionVisibility) IFontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.class), objc.Sel("fontCollectionWithName:visibility:"), name, visibility)
	return rv
}


// Renames the font collection with the specified name and visibility to the second name specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/rename(fromName:visibility:toName:)
func (fc _FontCollectionClass) RenameFontCollectionWithNameVisibilityToNameError(oldName objc.IObject /* cross-framework: FontCollectionName */, visibility FontCollectionVisibility, newName objc.IObject /* cross-framework: FontCollectionName */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("renameFontCollectionWithName:visibility:toName:error:"), oldName, visibility, newName, outError)
	return rv
}


// Make the given font collection visible by giving it a name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/show(_:withName:visibility:)
func (fc _FontCollectionClass) ShowFontCollectionWithNameVisibilityError(collection IFontCollection, name objc.IObject /* cross-framework: FontCollectionName */, visibility FontCollectionVisibility, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("showFontCollection:withName:visibility:error:"), collection, name, visibility, error_)
	return rv
}


// Returns all named collections visible to this process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/allFontCollectionNames
func (fc _FontCollectionClass) AllFontCollectionNames() []string {
	rv := objc.Send[[]string](objc.ID(fc.class), objc.Sel("allFontCollectionNames"))
	return rv
}

// The font collection that matches all registered fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/withAllAvailableDescriptors
func (fc _FontCollectionClass) FontCollectionWithAllAvailableDescriptors() FontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.class), objc.Sel("fontCollectionWithAllAvailableDescriptors"))
	return rv
}

// Returns an array of font descriptors matching the logical descriptors for the given font family.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/matchingDescriptors(forFamily:)
func (f_ FontCollection) MatchingDescriptorsForFamily(family objc.IObject /* cross-framework: NSString */) []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("matchingDescriptorsForFamily:"), family)
	return rv
}


// Returns an array of font descriptors matching the logical descriptors for the given font family and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/matchingDescriptors(forFamily:options:)
func (f_ FontCollection) MatchingDescriptorsForFamilyOptions(family objc.IObject /* cross-framework: NSString */, options foundation.IDictionary) []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("matchingDescriptorsForFamily:options:"), family, options)
	return rv
}


// Returns an array of font descriptors matching the logical descriptors with the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/matchingDescriptors(options:)
func (f_ FontCollection) MatchingDescriptorsWithOptions(options foundation.IDictionary) []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("matchingDescriptorsWithOptions:"), options)
	return rv
}


// Returns all named collections visible to this process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/allFontCollectionNames
func (f_ FontCollection) AllFontCollectionNames() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("allFontCollectionNames"))
	return rv
}


// A list of query font descriptors whose matching results are excluded from the list of matching descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/exclusionDescriptors
func (f_ FontCollection) ExclusionDescriptors() []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("exclusionDescriptors"))
	return rv
}


// An array of font descriptors matching the logical descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/matchingDescriptors
func (f_ FontCollection) MatchingDescriptors() []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("matchingDescriptors"))
	return rv
}


// An array of font descriptors whose matching results produce the collection’s matching descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/queryDescriptors
func (f_ FontCollection) QueryDescriptors() []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("queryDescriptors"))
	return rv
}


// The font collection that matches all registered fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/withAllAvailableDescriptors
func (f_ FontCollection) FontCollectionWithAllAvailableDescriptors() IFontCollection {
	rv := objc.Send[FontCollection](f_.ID, objc.Sel("fontCollectionWithAllAvailableDescriptors"))
	return rv
}


