// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSMutableFontCollection */


/* debug [class_header]: Header for NSMutableFontCollection */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableFontCollection */
// An interface definition for the [MutableFontCollection] class.
type IMutableFontCollection interface {
	IFontCollection
	
/* debug [class_interface_properties]: Properties for MutableFontCollection */
	// properties:
	ExclusionDescriptors() []FontDescriptor
	SetExclusionDescriptors(value []FontDescriptor)
	QueryDescriptors() []FontDescriptor
	SetQueryDescriptors(value []FontDescriptor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableFontCollection */
	// methods:
	AddQueryForDescriptors(descriptors []FontDescriptor)
	RemoveQueryForDescriptors(descriptors []FontDescriptor)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableFontCollection */
// Alloc allocates a new instance without initialization.
func (mc _MutableFontCollectionClass) Alloc() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableFontCollection */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableFontCollection */

// Creates a mutable font collection containing the fonts that match the specified font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(descriptors:)
func NewMutableFontCollectionWithDescriptors(queryDescriptors []FontDescriptor) MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(getMutableFontCollectionClass().class), objc.Sel("fontCollectionWithDescriptors:"), queryDescriptors)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableFontCollectionWithDescriptors */


// Creates a mutable font collection containing fonts suitable for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(locale:)
func NewMutableFontCollectionWithLocale(locale foundation.Locale) MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(getMutableFontCollectionClass().class), objc.Sel("fontCollectionWithLocale:"), locale)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableFontCollectionWithLocale */


// Creates a mutable named font collection object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:)
func NewMutableFontCollectionWithName(name FontCollectionName /* typedef */) MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(getMutableFontCollectionClass().class), objc.Sel("fontCollectionWithName:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableFontCollectionWithName */


// Creates a mutable font collection with the specified name and font visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:visibility:)
func NewMutableFontCollectionWithNameVisibility(name FontCollectionName /* typedef */, visibility FontCollectionVisibility) MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(getMutableFontCollectionClass().class), objc.Sel("fontCollectionWithName:visibility:"), name, visibility)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableFontCollectionWithNameVisibility */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableFontCollection */

// Creates a mutable font collection containing the fonts that match the specified font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(descriptors:)
func (mc _MutableFontCollectionClass) FontCollectionWithDescriptors(queryDescriptors []FontDescriptor) IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithDescriptors:"), queryDescriptors)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FontCollectionWithDescriptors) */


// Creates a mutable font collection containing fonts suitable for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(locale:)
func (mc _MutableFontCollectionClass) FontCollectionWithLocale(locale foundation.Locale) IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithLocale:"), locale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FontCollectionWithLocale) */


// Creates a mutable named font collection object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:)
func (mc _MutableFontCollectionClass) FontCollectionWithName(name FontCollectionName /* typedef */) IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FontCollectionWithName) */


// Creates a mutable font collection with the specified name and font visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:visibility:)
func (mc _MutableFontCollectionClass) FontCollectionWithNameVisibility(name FontCollectionName /* typedef */, visibility FontCollectionVisibility) IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithName:visibility:"), name, visibility)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FontCollectionWithNameVisibility) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableFontCollection */

// The mutable font collection that matches all registered fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/withAllAvailableDescriptors
func (mc _MutableFontCollectionClass) FontCollectionWithAllAvailableDescriptors() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.class), objc.Sel("fontCollectionWithAllAvailableDescriptors"))
	return rv
}/* debug [class_properties_class/property]: fontCollectionWithAllAvailableDescriptors */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableFontCollection */

// Edits the query and exclusion arrays by adding the specified font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/addQuery(for:)
func (m_ MutableFontCollection) AddQueryForDescriptors(descriptors []FontDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addQueryForDescriptors:"), descriptors)
}/* debug [instance_methods/method]: AddQueryForDescriptors */


// Edits the query and exclusion arrays by removing the specified font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/removeQuery(for:)
func (m_ MutableFontCollection) RemoveQueryForDescriptors(descriptors []FontDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeQueryForDescriptors:"), descriptors)
}/* debug [instance_methods/method]: RemoveQueryForDescriptors */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableFontCollection */

// The font descriptors to exclude from query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/exclusionDescriptors
func (m_ MutableFontCollection) ExclusionDescriptors() []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](m_.ID, objc.Sel("exclusionDescriptors"))
	return rv
}/* debug [instance_properties/getter]: exclusionDescriptors */


// The font descriptors to exclude from query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/exclusionDescriptors
func (m_ MutableFontCollection) SetExclusionDescriptors(value []FontDescriptor) {
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
}/* debug [instance_properties/setter]: exclusionDescriptors */


// The font descriptors to include in query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/queryDescriptors
func (m_ MutableFontCollection) QueryDescriptors() []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](m_.ID, objc.Sel("queryDescriptors"))
	return rv
}/* debug [instance_properties/getter]: queryDescriptors */


// The font descriptors to include in query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/queryDescriptors
func (m_ MutableFontCollection) SetQueryDescriptors(value []FontDescriptor) {
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
}/* debug [instance_properties/setter]: queryDescriptors */


// The mutable font collection that matches all registered fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/withAllAvailableDescriptors
func (m_ MutableFontCollection) FontCollectionWithAllAvailableDescriptors() IMutableFontCollection {
	rv := objc.Send[MutableFontCollection](m_.ID, objc.Sel("fontCollectionWithAllAvailableDescriptors"))
	return rv
}/* debug [instance_properties/getter]: fontCollectionWithAllAvailableDescriptors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableFontCollection */


