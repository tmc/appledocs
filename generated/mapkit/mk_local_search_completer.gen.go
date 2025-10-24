// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLocalSearchCompleter */


/* debug [class_header]: Header for MKLocalSearchCompleter */
// The class instance for the [MKLocalSearchCompleter] class.
var (
	MKLocalSearchCompleterClass     _MKLocalSearchCompleterClass
	MKLocalSearchCompleterClassOnce sync.Once
)

func getMKLocalSearchCompleterClass() _MKLocalSearchCompleterClass {
	MKLocalSearchCompleterClassOnce.Do(func() {
		MKLocalSearchCompleterClass = _MKLocalSearchCompleterClass{objc.GetClass("MKLocalSearchCompleter")}
	})
	return MKLocalSearchCompleterClass
}

type _MKLocalSearchCompleterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLocalSearchCompleter */
// An interface definition for the [MKLocalSearchCompleter] class.
type IMKLocalSearchCompleter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLocalSearchCompleter */
	// properties:
	AddressFilter() IMKAddressFilter
	SetAddressFilter(value IMKAddressFilter)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	FilterType() MKSearchCompletionFilterType
	SetFilterType(value MKSearchCompletionFilterType)
	Searching() bool
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	QueryFragment() objc.IObject /* cross-framework: NSString */
	SetQueryFragment(value objc.IObject /* cross-framework: NSString */)
	Region() objc.IObject /* cross-framework: MKCoordinateRegion */
	SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */)
	RegionPriority() MKLocalSearchRegionPriority
	SetRegionPriority(value MKLocalSearchRegionPriority)
	Results() []MKLocalSearchCompletion
	ResultTypes() MKLocalSearchCompleterResultType
	SetResultTypes(value MKLocalSearchCompleterResultType)
	IsSearching() bool
	SetIsSearching(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLocalSearchCompleter */
	// methods:
	Cancel()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLocalSearchCompleter */
// Alloc allocates a new instance without initialization.
func (mc _MKLocalSearchCompleterClass) Alloc() MKLocalSearchCompleter {
	rv := objc.Send[MKLocalSearchCompleter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLocalSearchCompleterClass) New() MKLocalSearchCompleter {
	rv := objc.Send[MKLocalSearchCompleter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLocalSearchCompleter) Init() MKLocalSearchCompleter {
	rv := objc.Send[MKLocalSearchCompleter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLocalSearchCompleter) Autorelease() MKLocalSearchCompleter {
	rv := objc.Send[MKLocalSearchCompleter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLocalSearchCompleter creates a new MKLocalSearchCompleter instance.
func NewMKLocalSearchCompleter() MKLocalSearchCompleter {
	return getMKLocalSearchCompleterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLocalSearchCompleter */
// A utility object for generating a list of completion strings based on a partial search string that you provide.
//
// You use an object to retrieve auto-complete suggestions for your own map-based search controls. As the user types text, you feed the current text string into the search completer object, which delivers possible string completions that match locations or points of interest. You create and configure objects yourself. You must always assign a delegate object to the search completer so that you can receive the search results that it generates. Specify a search region to restrict results to a designated area. The following code shows a simple example of a view controller that stores the object in a property. The view controller itself acts as the delegate for the completer and the view controller uses the region associated with an object that’s part of the view controller’s interface. Completer objects are long-lived objects, so you can store strong references to them and reuse them later in your code. Listing 1. Creating and configuring a search completer Update the value of the completer’s property to begin a search query. You can update this property in real time as the user types new characters into a text field because the completer object waits a short amount of time for the query string to stabilize. When modifications to the query string stop, the completer initiates a new search and returns the results to your delegate as an array of objects.


// A utility object for generating a list of completion strings based on a partial search string that you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter
type MKLocalSearchCompleter struct {
	objectivec.Object
}

// MKLocalSearchCompleterFrom constructs a [MKLocalSearchCompleter] from an unsafe.Pointer.
//
// A utility object for generating a list of completion strings based on a partial search string that you provide.
func MKLocalSearchCompleterFrom(ptr unsafe.Pointer) MKLocalSearchCompleter {
	return MKLocalSearchCompleter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLocalSearchCompleter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLocalSearchCompleter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLocalSearchCompleter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLocalSearchCompleter */

// Cancels an in-progress search operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/cancel()
func (m_ MKLocalSearchCompleter) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLocalSearchCompleter */

// A filter that lists which address options to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/addressFilter
func (m_ MKLocalSearchCompleter) AddressFilter() IMKAddressFilter {
	rv := objc.Send[MKAddressFilter](m_.ID, objc.Sel("addressFilter"))
	return rv
}/* debug [instance_properties/getter]: addressFilter */


// A filter that lists which address options to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/addressFilter
func (m_ MKLocalSearchCompleter) SetAddressFilter(value IMKAddressFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddressFilter:"), value)
}/* debug [instance_properties/setter]: addressFilter */


// The object that receives the completion results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/delegate
func (m_ MKLocalSearchCompleter) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The object that receives the completion results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/delegate
func (m_ MKLocalSearchCompleter) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The filter options for the search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/filterType-swift.property
func (m_ MKLocalSearchCompleter) FilterType() MKSearchCompletionFilterType {
	rv := objc.Send[MKSearchCompletionFilterType](m_.ID, objc.Sel("filterType"))
	return rv
}/* debug [instance_properties/getter]: filterType */


// The filter options for the search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/filterType-swift.property
func (m_ MKLocalSearchCompleter) SetFilterType(value MKSearchCompletionFilterType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFilterType:"), value)
}/* debug [instance_properties/setter]: filterType */


// A Boolean value that indicates whether a search operation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/isSearching
func (m_ MKLocalSearchCompleter) Searching() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("searching"))
	return rv
}/* debug [instance_properties/getter]: searching */


// A filter that lists point of interest categories to include or exclude in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/pointOfInterestFilter
func (m_ MKLocalSearchCompleter) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// A filter that lists point of interest categories to include or exclude in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/pointOfInterestFilter
func (m_ MKLocalSearchCompleter) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// The search string that you want completions for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/queryFragment
func (m_ MKLocalSearchCompleter) QueryFragment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("queryFragment"))
	return rv
}/* debug [instance_properties/getter]: queryFragment */


// The search string that you want completions for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/queryFragment
func (m_ MKLocalSearchCompleter) SetQueryFragment(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueryFragment:"), value)
}/* debug [instance_properties/setter]: queryFragment */


// The region that defines the geographic scope of the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/region
func (m_ MKLocalSearchCompleter) Region() objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */


// The region that defines the geographic scope of the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/region
func (m_ MKLocalSearchCompleter) SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}/* debug [instance_properties/setter]: region */


// A value that indicates the importance of the configured region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/regionPriority
func (m_ MKLocalSearchCompleter) RegionPriority() MKLocalSearchRegionPriority {
	rv := objc.Send[MKLocalSearchRegionPriority](m_.ID, objc.Sel("regionPriority"))
	return rv
}/* debug [instance_properties/getter]: regionPriority */


// A value that indicates the importance of the configured region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/regionPriority
func (m_ MKLocalSearchCompleter) SetRegionPriority(value MKLocalSearchRegionPriority) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegionPriority:"), value)
}/* debug [instance_properties/setter]: regionPriority */


// The most recently received search completions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/results
func (m_ MKLocalSearchCompleter) Results() []MKLocalSearchCompletion {
	rv := objc.Send[[]MKLocalSearchCompletion](m_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The types of search completions to include.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/resultTypes
func (m_ MKLocalSearchCompleter) ResultTypes() MKLocalSearchCompleterResultType {
	rv := objc.Send[MKLocalSearchCompleterResultType](m_.ID, objc.Sel("resultTypes"))
	return rv
}/* debug [instance_properties/getter]: resultTypes */


// The types of search completions to include.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/resultTypes
func (m_ MKLocalSearchCompleter) SetResultTypes(value MKLocalSearchCompleterResultType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultTypes:"), value)
}/* debug [instance_properties/setter]: resultTypes */


// A Boolean value that indicates whether a search operation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/issearching
func (m_ MKLocalSearchCompleter) IsSearching() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSearching"))
	return rv
}/* debug [instance_properties/getter]: isSearching */


// A Boolean value that indicates whether a search operation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/issearching
func (m_ MKLocalSearchCompleter) SetIsSearching(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSearching:"), value)
}/* debug [instance_properties/setter]: isSearching */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLocalSearchCompleter */



