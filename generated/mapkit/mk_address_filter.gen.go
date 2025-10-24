// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKAddressFilter */


/* debug [class_header]: Header for MKAddressFilter */
// The class instance for the [MKAddressFilter] class.
var (
	MKAddressFilterClass     _MKAddressFilterClass
	MKAddressFilterClassOnce sync.Once
)

func getMKAddressFilterClass() _MKAddressFilterClass {
	MKAddressFilterClassOnce.Do(func() {
		MKAddressFilterClass = _MKAddressFilterClass{objc.GetClass("MKAddressFilter")}
	})
	return MKAddressFilterClass
}

type _MKAddressFilterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKAddressFilter */
// An interface definition for the [MKAddressFilter] class.
type IMKAddressFilter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKAddressFilter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKAddressFilter */
	// methods:
	ExcludesOptions(options MKAddressFilterOption) bool
	IncludesOptions(options MKAddressFilterOption) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKAddressFilter */
// Alloc allocates a new instance without initialization.
func (mc _MKAddressFilterClass) Alloc() MKAddressFilter {
	rv := objc.Send[MKAddressFilter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKAddressFilterClass) New() MKAddressFilter {
	rv := objc.Send[MKAddressFilter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKAddressFilter) Init() MKAddressFilter {
	rv := objc.Send[MKAddressFilter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKAddressFilter) Autorelease() MKAddressFilter {
	rv := objc.Send[MKAddressFilter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKAddressFilter creates a new MKAddressFilter instance.
func NewMKAddressFilter() MKAddressFilter {
	return getMKAddressFilterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKAddressFilter */
// An object that filters which address options to include or exclude in search results.
//
// Use this object to filter search results by criteria, such as country, region, and municipality. See for more information.


// An object that filters which address options to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter
type MKAddressFilter struct {
	objectivec.Object
}

// MKAddressFilterFrom constructs a [MKAddressFilter] from an unsafe.Pointer.
//
// An object that filters which address options to include or exclude in search results.
func MKAddressFilterFrom(ptr unsafe.Pointer) MKAddressFilter {
	return MKAddressFilter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKAddressFilter */

// Creates an address filter with options for excluding results in a search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/init(excluding:)
func NewMKAddressFilterExcludingOptions(options MKAddressFilterOption) MKAddressFilter {
	instance := getMKAddressFilterClass().Alloc()
	rv := objc.Send[MKAddressFilter](instance.ID, objc.Sel("initExcludingOptions:"), options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKAddressFilterExcludingOptions */


// Creates an address filter with options for including results in a search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/init(including:)
func NewMKAddressFilterIncludingOptions(options MKAddressFilterOption) MKAddressFilter {
	instance := getMKAddressFilterClass().Alloc()
	rv := objc.Send[MKAddressFilter](instance.ID, objc.Sel("initIncludingOptions:"), options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKAddressFilterIncludingOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKAddressFilter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKAddressFilter */

// A list of categories to exclude from a search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/excludingAll
func (mc _MKAddressFilterClass) FilterExcludingAll() MKAddressFilter {
	rv := objc.Send[MKAddressFilter](objc.ID(mc.class), objc.Sel("filterExcludingAll"))
	return rv
}/* debug [class_properties_class/property]: filterExcludingAll */

// A list of categories to include in a search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/includingAll
func (mc _MKAddressFilterClass) FilterIncludingAll() MKAddressFilter {
	rv := objc.Send[MKAddressFilter](objc.ID(mc.class), objc.Sel("filterIncludingAll"))
	return rv
}/* debug [class_properties_class/property]: filterIncludingAll */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKAddressFilter */

// Indicates whether options are excluded from filtering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/excludes(_:)
func (m_ MKAddressFilter) ExcludesOptions(options MKAddressFilterOption) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("excludesOptions:"), options)
	return rv
}/* debug [instance_methods/method]: ExcludesOptions */


// Indicates whether options are included for filtering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/includes(_:)
func (m_ MKAddressFilter) IncludesOptions(options MKAddressFilterOption) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("includesOptions:"), options)
	return rv
}/* debug [instance_methods/method]: IncludesOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKAddressFilter */

// A list of categories to exclude from a search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/excludingAll
func (m_ MKAddressFilter) FilterExcludingAll() IMKAddressFilter {
	rv := objc.Send[MKAddressFilter](m_.ID, objc.Sel("filterExcludingAll"))
	return rv
}/* debug [instance_properties/getter]: filterExcludingAll */


// A list of categories to include in a search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/includingAll
func (m_ MKAddressFilter) FilterIncludingAll() IMKAddressFilter {
	rv := objc.Send[MKAddressFilter](m_.ID, objc.Sel("filterIncludingAll"))
	return rv
}/* debug [instance_properties/getter]: filterIncludingAll */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKAddressFilter */


