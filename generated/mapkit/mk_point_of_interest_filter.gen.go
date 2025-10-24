// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPointOfInterestFilter */


/* debug [class_header]: Header for MKPointOfInterestFilter */
// The class instance for the [MKPointOfInterestFilter] class.
var (
	MKPointOfInterestFilterClass     _MKPointOfInterestFilterClass
	MKPointOfInterestFilterClassOnce sync.Once
)

func getMKPointOfInterestFilterClass() _MKPointOfInterestFilterClass {
	MKPointOfInterestFilterClassOnce.Do(func() {
		MKPointOfInterestFilterClass = _MKPointOfInterestFilterClass{objc.GetClass("MKPointOfInterestFilter")}
	})
	return MKPointOfInterestFilterClass
}

type _MKPointOfInterestFilterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPointOfInterestFilter */
// An interface definition for the [MKPointOfInterestFilter] class.
type IMKPointOfInterestFilter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKPointOfInterestFilter */
	// properties:
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPointOfInterestFilter */
	// methods:
	ExcludesCategory(category MKPointOfInterestCategory /* typedef */) bool
	IncludesCategory(category MKPointOfInterestCategory /* typedef */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPointOfInterestFilter */
// Alloc allocates a new instance without initialization.
func (mc _MKPointOfInterestFilterClass) Alloc() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPointOfInterestFilterClass) New() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPointOfInterestFilter) Init() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPointOfInterestFilter) Autorelease() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPointOfInterestFilter creates a new MKPointOfInterestFilter instance.
func NewMKPointOfInterestFilter() MKPointOfInterestFilter {
	return getMKPointOfInterestFilterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPointOfInterestFilter */
// A filter that includes or excludes point of interest categories from a map view, local search, or local search completer.
//
// You can apply a point of interest filter in a map view ( ), a local search request ( ), a search completer ( ), and in snapshot options ( ).


// A filter that includes or excludes point of interest categories from a map view, local search, or local search completer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter
type MKPointOfInterestFilter struct {
	objectivec.Object
}

// MKPointOfInterestFilterFrom constructs a [MKPointOfInterestFilter] from an unsafe.Pointer.
//
// A filter that includes or excludes point of interest categories from a map view, local search, or local search completer.
func MKPointOfInterestFilterFrom(ptr unsafe.Pointer) MKPointOfInterestFilter {
	return MKPointOfInterestFilter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPointOfInterestFilter */

// Initialize the point of interest filter with a list of categories to exclude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter/init(excluding:)
func NewMKPointOfInterestFilterExcludingCategories(categories []string) MKPointOfInterestFilter {
	instance := getMKPointOfInterestFilterClass().Alloc()
	rv := objc.Send[MKPointOfInterestFilter](instance.ID, objc.Sel("initExcludingCategories:"), categories)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPointOfInterestFilterExcludingCategories */


// Initialize the point of interest filter with a list of categories to include.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter/init(including:)
func NewMKPointOfInterestFilterIncludingCategories(categories []string) MKPointOfInterestFilter {
	instance := getMKPointOfInterestFilterClass().Alloc()
	rv := objc.Send[MKPointOfInterestFilter](instance.ID, objc.Sel("initIncludingCategories:"), categories)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPointOfInterestFilterIncludingCategories */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPointOfInterestFilter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPointOfInterestFilter */

// A filter that excludes all point of interest categories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter/excludingAll
func (mc _MKPointOfInterestFilterClass) FilterExcludingAllCategories() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](objc.ID(mc.class), objc.Sel("filterExcludingAllCategories"))
	return rv
}/* debug [class_properties_class/property]: filterExcludingAllCategories */

// A filter that includes all point of interest categories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter/includingAll
func (mc _MKPointOfInterestFilterClass) FilterIncludingAllCategories() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](objc.ID(mc.class), objc.Sel("filterIncludingAllCategories"))
	return rv
}/* debug [class_properties_class/property]: filterIncludingAllCategories */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPointOfInterestFilter */

// Returns a Boolean value indicating whether the filter excludes the point of interest category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter/excludes(_:)
func (m_ MKPointOfInterestFilter) ExcludesCategory(category MKPointOfInterestCategory /* typedef */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("excludesCategory:"), category)
	return rv
}/* debug [instance_methods/method]: ExcludesCategory */


// Returns a Boolean value indicating whether the filter includes the point of interest category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter/includes(_:)
func (m_ MKPointOfInterestFilter) IncludesCategory(category MKPointOfInterestCategory /* typedef */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("includesCategory:"), category)
	return rv
}/* debug [instance_methods/method]: IncludesCategory */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPointOfInterestFilter */

// A filter that excludes all point of interest categories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter/excludingAll
func (m_ MKPointOfInterestFilter) FilterExcludingAllCategories() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("filterExcludingAllCategories"))
	return rv
}/* debug [instance_properties/getter]: filterExcludingAllCategories */


// A filter that includes all point of interest categories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointOfInterestFilter/includingAll
func (m_ MKPointOfInterestFilter) FilterIncludingAllCategories() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("filterIncludingAllCategories"))
	return rv
}/* debug [instance_properties/getter]: filterIncludingAllCategories */


// A filter that lists point of interest categories to include or exclude in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/pointofinterestfilter
func (m_ MKPointOfInterestFilter) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// A filter that lists point of interest categories to include or exclude in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/pointofinterestfilter
func (m_ MKPointOfInterestFilter) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPointOfInterestFilter */


