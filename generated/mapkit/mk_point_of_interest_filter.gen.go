// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKPointOfInterestFilter] class.
type IMKPointOfInterestFilter interface {
	objectivec.IObject
	// properties:
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MKPointOfInterestFilterClass) Alloc() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A filter that lists point of interest categories to include or exclude in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/pointofinterestfilter
func (m_ MKPointOfInterestFilter) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// A filter that lists point of interest categories to include or exclude in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearchcompleter/pointofinterestfilter
func (m_ MKPointOfInterestFilter) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}



