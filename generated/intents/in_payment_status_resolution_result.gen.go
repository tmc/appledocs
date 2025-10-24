// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INPaymentStatusResolutionResult] class.
var (
	INPaymentStatusResolutionResultClass     _INPaymentStatusResolutionResultClass
	INPaymentStatusResolutionResultClassOnce sync.Once
)

func getINPaymentStatusResolutionResultClass() _INPaymentStatusResolutionResultClass {
	INPaymentStatusResolutionResultClassOnce.Do(func() {
		INPaymentStatusResolutionResultClass = _INPaymentStatusResolutionResultClass{objc.GetClass("INPaymentStatusResolutionResult")}
	})
	return INPaymentStatusResolutionResultClass
}

type _INPaymentStatusResolutionResultClass struct {
	class objc.Class
}

// An interface definition for the [INPaymentStatusResolutionResult] class.
type IINPaymentStatusResolutionResult interface {
	IINIntentResolutionResult
	// properties:
	// methods:
}

// A resolution result for the payment status of a bill during a search.
//
// An object is what you return when resolving parameters containing an value. Use the creation method that best reflects your ability to resolve the parameter successfully. For additional resolution options, see .

// A resolution result for the payment status of a bill during a search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPaymentStatusResolutionResult
type INPaymentStatusResolutionResult struct {
	INIntentResolutionResult
}

// INPaymentStatusResolutionResultFrom constructs a [INPaymentStatusResolutionResult] from an unsafe.Pointer.
//
// A resolution result for the payment status of a bill during a search.
func INPaymentStatusResolutionResultFrom(ptr unsafe.Pointer) INPaymentStatusResolutionResult {
	return INPaymentStatusResolutionResult{
		INIntentResolutionResult: INIntentResolutionResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INPaymentStatusResolutionResultClass) Alloc() INPaymentStatusResolutionResult {
	rv := objc.Send[INPaymentStatusResolutionResult](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INPaymentStatusResolutionResultClass) New() INPaymentStatusResolutionResult {
	rv := objc.Send[INPaymentStatusResolutionResult](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INPaymentStatusResolutionResult) Init() INPaymentStatusResolutionResult {
	rv := objc.Send[INPaymentStatusResolutionResult](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INPaymentStatusResolutionResult) Autorelease() INPaymentStatusResolutionResult {
	rv := objc.Send[INPaymentStatusResolutionResult](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINPaymentStatusResolutionResult creates a new INPaymentStatusResolutionResult instance.
func NewINPaymentStatusResolutionResult() INPaymentStatusResolutionResult {
	return getINPaymentStatusResolutionResultClass().New()
}
