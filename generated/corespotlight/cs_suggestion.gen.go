// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CSSuggestion] class.
type ICSSuggestion interface {
	objectivec.IObject
	Compare(other unsafe.Pointer) unsafe.Pointer
	CompareByRank(other unsafe.Pointer) unsafe.Pointer
}

// The kind of suggestion to use in a query.
//
// Your app uses objects to populate a contextual menu of suggestions.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CSSuggestionClass) Alloc() CSSuggestion {
	rv := objc.Send[CSSuggestion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Compares the suggestion with a second specified suggestion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/compare(_:)
func (c_ CSSuggestion) Compare(other unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compare:"), other)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/compare(byRank:)
func (c_ CSSuggestion) CompareByRank(other unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compareByRank:"), other)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/localizedAttributedSuggestion-oq3b
func (c_ CSSuggestion) LocalizedAttributedSuggestion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("localizedAttributedSuggestion"))
	return rv
}

// The type of suggestion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/suggestionKind-swift.property
func (c_ CSSuggestion) SuggestionKind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("suggestionKind"))
	return rv
}



