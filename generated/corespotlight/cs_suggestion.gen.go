// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CSSuggestion */


/* debug [class_header]: Header for CSSuggestion */
// The class instance for the [CSSuggestion] class.
var (
	CSSuggestionClass     _CSSuggestionClass
	CSSuggestionClassOnce sync.Once
)

func getCSSuggestionClass() _CSSuggestionClass {
	CSSuggestionClassOnce.Do(func() {
		CSSuggestionClass = _CSSuggestionClass{objc.GetClass("CSSuggestion")}
	})
	return CSSuggestionClass
}

type _CSSuggestionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSSuggestion */
// An interface definition for the [CSSuggestion] class.
type ICSSuggestion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CSSuggestion */
	// properties:
	LocalizedAttributedSuggestion() foundation.AttributedString
	SuggestionKind() CSSuggestionKind
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSSuggestion */
	// methods:
	Compare(other ICSSuggestion) ComparisonResult /* not a class type */
	CompareByRank(other ICSSuggestion) ComparisonResult /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSSuggestion */
// Alloc allocates a new instance without initialization.
func (cc _CSSuggestionClass) Alloc() CSSuggestion {
	rv := objc.Send[CSSuggestion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSSuggestionClass) New() CSSuggestion {
	rv := objc.Send[CSSuggestion](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSSuggestion) Init() CSSuggestion {
	rv := objc.Send[CSSuggestion](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSSuggestion) Autorelease() CSSuggestion {
	rv := objc.Send[CSSuggestion](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSuggestion creates a new CSSuggestion instance.
func NewCSSuggestion() CSSuggestion {
	return getCSSuggestionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSSuggestion */
// The kind of suggestion to use in a query.
//
// Your app uses objects to populate a contextual menu of suggestions.


// The kind of suggestion to use in a query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion
type CSSuggestion struct {
	objectivec.Object
}

// CSSuggestionFrom constructs a [CSSuggestion] from an unsafe.Pointer.
//
// The kind of suggestion to use in a query.
func CSSuggestionFrom(ptr unsafe.Pointer) CSSuggestion {
	return CSSuggestion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSSuggestion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSSuggestion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSSuggestion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSSuggestion */

// Compares the suggestion with a second specified suggestion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/compare(_:)
func (c_ CSSuggestion) Compare(other ICSSuggestion) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](c_.ID, objc.Sel("compare:"), other)
	return rv
}/* debug [instance_methods/method]: Compare */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/compare(byRank:)
func (c_ CSSuggestion) CompareByRank(other ICSSuggestion) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](c_.ID, objc.Sel("compareByRank:"), other)
	return rv
}/* debug [instance_methods/method]: CompareByRank */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSSuggestion */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/localizedAttributedSuggestion-oq3b
func (c_ CSSuggestion) LocalizedAttributedSuggestion() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("localizedAttributedSuggestion"))
	return rv
}/* debug [instance_properties/getter]: localizedAttributedSuggestion */


// The type of suggestion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/suggestionKind-swift.property
func (c_ CSSuggestion) SuggestionKind() CSSuggestionKind {
	rv := objc.Send[CSSuggestionKind](c_.ID, objc.Sel("suggestionKind"))
	return rv
}/* debug [instance_properties/getter]: suggestionKind */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSSuggestion */



