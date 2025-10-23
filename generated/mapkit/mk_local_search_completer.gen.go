// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKLocalSearchCompleter] class.
type IMKLocalSearchCompleter interface {
	objectivec.IObject
	// properties:
	AddressFilter() unsafe.Pointer
	SetAddressFilter(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	FilterType() unsafe.Pointer
	SetFilterType(value unsafe.Pointer)
	IsSearching() bool /* primitive/slice/pointer. */
	SetIsSearching(value bool /* primitive/slice/pointer. */)
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	QueryFragment() string /* primitive/slice/pointer. */
	SetQueryFragment(value string /* primitive/slice/pointer. */)
	Region() unsafe.Pointer
	SetRegion(value unsafe.Pointer)
	RegionPriority() unsafe.Pointer
	SetRegionPriority(value unsafe.Pointer)
	ResultTypes() unsafe.Pointer
	SetResultTypes(value unsafe.Pointer)
	Results() unsafe.Pointer
	SetResults(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MKLocalSearchCompleterClass) Alloc() MKLocalSearchCompleter {
	rv := objc.Send[MKLocalSearchCompleter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A filter that lists which address options to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/addressfilter
func (m_ MKLocalSearchCompleter) AddressFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addressFilter"))
	return rv
}


// A filter that lists which address options to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/addressfilter
func (m_ MKLocalSearchCompleter) SetAddressFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddressFilter:"), value)
}


// The object that receives the completion results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/delegate
func (m_ MKLocalSearchCompleter) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}


// The object that receives the completion results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/delegate
func (m_ MKLocalSearchCompleter) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}


// The filter options for the search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/filtertype-swift.property
func (m_ MKLocalSearchCompleter) FilterType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("filterType"))
	return rv
}


// The filter options for the search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/filtertype-swift.property
func (m_ MKLocalSearchCompleter) SetFilterType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFilterType:"), value)
}


// A Boolean value that indicates whether a search operation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/issearching
func (m_ MKLocalSearchCompleter) IsSearching() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSearching"))
	return rv
}


// A Boolean value that indicates whether a search operation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/issearching
func (m_ MKLocalSearchCompleter) SetIsSearching(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSearching:"), value)
}


// A filter that lists point of interest categories to include or exclude in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/pointofinterestfilter
func (m_ MKLocalSearchCompleter) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// A filter that lists point of interest categories to include or exclude in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/pointofinterestfilter
func (m_ MKLocalSearchCompleter) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}


// The search string that you want completions for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/queryfragment
func (m_ MKLocalSearchCompleter) QueryFragment() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("queryFragment"))
	return rv
}


// The search string that you want completions for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/queryfragment
func (m_ MKLocalSearchCompleter) SetQueryFragment(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueryFragment:"), objc.String(value))
}


// The region that defines the geographic scope of the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/region
func (m_ MKLocalSearchCompleter) Region() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("region"))
	return rv
}


// The region that defines the geographic scope of the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/region
func (m_ MKLocalSearchCompleter) SetRegion(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}


// A value that indicates the importance of the configured region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/regionpriority
func (m_ MKLocalSearchCompleter) RegionPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("regionPriority"))
	return rv
}


// A value that indicates the importance of the configured region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/regionpriority
func (m_ MKLocalSearchCompleter) SetRegionPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegionPriority:"), value)
}


// The types of search completions to include.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/resulttypes
func (m_ MKLocalSearchCompleter) ResultTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("resultTypes"))
	return rv
}


// The types of search completions to include.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/resulttypes
func (m_ MKLocalSearchCompleter) SetResultTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultTypes:"), value)
}


// The most recently received search completions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/results
func (m_ MKLocalSearchCompleter) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("results"))
	return rv
}


// The most recently received search completions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/results
func (m_ MKLocalSearchCompleter) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResults:"), value)
}



