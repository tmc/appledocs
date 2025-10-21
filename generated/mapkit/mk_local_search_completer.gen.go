// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// A utility object for generating a list of completion strings based on a partial search string that you provide.
//
// You use an object to retrieve auto-complete suggestions for your own map-based search controls. As the user types text, you feed the current text string into the search completer object, which delivers possible string completions that match locations or points of interest. You create and configure objects yourself. You must always assign a delegate object to the search completer so that you can receive the search results that it generates. Specify a search region to restrict results to a designated area. The following code shows a simple example of a view controller that stores the object in a property. The view controller itself acts as the delegate for the completer and the view controller uses the region associated with an object that’s part of the view controller’s interface. Completer objects are long-lived objects, so you can store strong references to them and reuse them later in your code. Listing 1. Creating and configuring a search completer Update the value of the completer’s property to begin a search query. You can update this property in real time as the user types new characters into a text field because the completer object waits a short amount of time for the query string to stabilize. When modifications to the query string stop, the completer initiates a new search and returns the results to your delegate as an array of objects.
//
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


// The filter options for the search results.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/filterType-swift.property
func (m_ MKLocalSearchCompleter) FilterType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("filterType"))
	return rv
}


// SetFilterType sets the value of the filterType property.
// The filter options for the search results.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/filterType-swift.property
func (m_ MKLocalSearchCompleter) SetFilterType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFilterType:"), value)
}

// The most recently received search completions.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/results
func (m_ MKLocalSearchCompleter) Results() []MKLocalSearchCompletion {
	rv := objc.Send[[]MKLocalSearchCompletion](m_.ID, objc.Sel("results"))
	return rv
}



