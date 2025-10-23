// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INPlacemarkResolutionResult] class.
var (
	INPlacemarkResolutionResultClass     _INPlacemarkResolutionResultClass
	INPlacemarkResolutionResultClassOnce sync.Once
)

func getINPlacemarkResolutionResultClass() _INPlacemarkResolutionResultClass {
	INPlacemarkResolutionResultClassOnce.Do(func() {
		INPlacemarkResolutionResultClass = _INPlacemarkResolutionResultClass{objc.GetClass("INPlacemarkResolutionResult")}
	})
	return INPlacemarkResolutionResultClass
}

type _INPlacemarkResolutionResultClass struct {
	class objc.Class
}

// An interface definition for the [INPlacemarkResolutionResult] class.
type IINPlacemarkResolutionResult interface {
	IINIntentResolutionResult
	// properties:
	// methods:
}

// A resolution result for placemark information associated with an intent.
//
// An object is what you return when resolving parameters containing an object. Use the creation method that best reflects your ability to successfully resolve the parameter. For additional resolution options, see .


// A resolution result for placemark information associated with an intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlacemarkResolutionResult
type INPlacemarkResolutionResult struct {
	INIntentResolutionResult
}

// INPlacemarkResolutionResultFrom constructs a [INPlacemarkResolutionResult] from an unsafe.Pointer.
//
// A resolution result for placemark information associated with an intent.
func INPlacemarkResolutionResultFrom(ptr unsafe.Pointer) INPlacemarkResolutionResult {
	return INPlacemarkResolutionResult{
		INIntentResolutionResult: INIntentResolutionResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INPlacemarkResolutionResultClass) Alloc() INPlacemarkResolutionResult {
	rv := objc.Send[INPlacemarkResolutionResult](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INPlacemarkResolutionResultClass) New() INPlacemarkResolutionResult {
	rv := objc.Send[INPlacemarkResolutionResult](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INPlacemarkResolutionResult) Init() INPlacemarkResolutionResult {
	rv := objc.Send[INPlacemarkResolutionResult](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INPlacemarkResolutionResult) Autorelease() INPlacemarkResolutionResult {
	rv := objc.Send[INPlacemarkResolutionResult](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINPlacemarkResolutionResult creates a new INPlacemarkResolutionResult instance.
func NewINPlacemarkResolutionResult() INPlacemarkResolutionResult {
	return getINPlacemarkResolutionResultClass().New()
}



// Creates an object whose resolution requires that the user must confirm the value before proceeding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlacemarkResolutionResult/confirmationRequired(with:)
func (ic _INPlacemarkResolutionResultClass) ConfirmationRequiredWithPlacemarkToConfirm(placemarkToConfirm objc.IObject /* cross-framework Placemark */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("confirmationRequiredWithPlacemarkToConfirm:"), placemarkToConfirm)
	return rv
}



