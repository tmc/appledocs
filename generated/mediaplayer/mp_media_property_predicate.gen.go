// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaPropertyPredicate] class.
var (
	MediaPropertyPredicateClass     _MediaPropertyPredicateClass
	MediaPropertyPredicateClassOnce sync.Once
)

func getMediaPropertyPredicateClass() _MediaPropertyPredicateClass {
	MediaPropertyPredicateClassOnce.Do(func() {
		MediaPropertyPredicateClass = _MediaPropertyPredicateClass{objc.GetClass("MPMediaPropertyPredicate")}
	})
	return MediaPropertyPredicateClass
}

type _MediaPropertyPredicateClass struct {
	class objc.Class
}

// An interface definition for the [MediaPropertyPredicate] class.
type IMediaPropertyPredicate interface {
	IMediaPredicate
	// properties:
	// methods:
}

// A set of predicates for defining a filter in a media query.
//
// Use one or more objects to define the filter in a media query to retrieve a subset of media items from the Music library. A predicate in this context is a statement of a logical condition that you want to test each media item against. The query retrieves the items that satisfy that condition. You define Music library queries, and retrieve query results, using the class. and describe the media items and media item collections that you can retrieve with a query.


// A set of predicates for defining a filter in a media query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPropertyPredicate
type MediaPropertyPredicate struct {
	MediaPredicate
}

// MediaPropertyPredicateFrom constructs a [MediaPropertyPredicate] from an unsafe.Pointer.
//
// A set of predicates for defining a filter in a media query.
func MediaPropertyPredicateFrom(ptr unsafe.Pointer) MediaPropertyPredicate {
	return MediaPropertyPredicate{
		MediaPredicate: MediaPredicateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaPropertyPredicateClass) Alloc() MediaPropertyPredicate {
	rv := objc.Send[MediaPropertyPredicate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaPropertyPredicateClass) New() MediaPropertyPredicate {
	rv := objc.Send[MediaPropertyPredicate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPropertyPredicate) Init() MediaPropertyPredicate {
	rv := objc.Send[MediaPropertyPredicate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPropertyPredicate) Autorelease() MediaPropertyPredicate {
	rv := objc.Send[MediaPropertyPredicate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPropertyPredicate creates a new MediaPropertyPredicate instance.
func NewMediaPropertyPredicate() MediaPropertyPredicate {
	return getMediaPropertyPredicateClass().New()
}



// Creates a media property predicate with the default comparison type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPropertyPredicate/init(value:forProperty:)
func NewMediaPropertyPredicateWithValueForProperty(value objectivec.IObject, property objc.IObject /* cross-framework: NSString */) MediaPropertyPredicate {
	rv := objc.Send[MediaPropertyPredicate](objc.ID(getMediaPropertyPredicateClass().class), objc.Sel("predicateWithValue:forProperty:"), value, property)
	return rv
}


// Creates a media property predicate with a specified comparison type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPropertyPredicate/init(value:forProperty:comparisonType:)
func NewMediaPropertyPredicateWithValueForPropertyComparisonType(value objectivec.IObject, property objc.IObject /* cross-framework: NSString */, comparisonType MediaPredicateComparison) MediaPropertyPredicate {
	rv := objc.Send[MediaPropertyPredicate](objc.ID(getMediaPropertyPredicateClass().class), objc.Sel("predicateWithValue:forProperty:comparisonType:"), value, property, comparisonType)
	return rv
}



// Creates a media property predicate with the default comparison type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPropertyPredicate/init(value:forProperty:)
func (mc _MediaPropertyPredicateClass) PredicateWithValueForProperty(value objectivec.IObject, property objc.IObject /* cross-framework: NSString */) IMediaPropertyPredicate {
	rv := objc.Send[MediaPropertyPredicate](objc.ID(mc.class), objc.Sel("predicateWithValue:forProperty:"), value, property)
	return rv
}


// Creates a media property predicate with a specified comparison type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPropertyPredicate/init(value:forProperty:comparisonType:)
func (mc _MediaPropertyPredicateClass) PredicateWithValueForPropertyComparisonType(value objectivec.IObject, property objc.IObject /* cross-framework: NSString */, comparisonType MediaPredicateComparison) IMediaPropertyPredicate {
	rv := objc.Send[MediaPropertyPredicate](objc.ID(mc.class), objc.Sel("predicateWithValue:forProperty:comparisonType:"), value, property, comparisonType)
	return rv
}


