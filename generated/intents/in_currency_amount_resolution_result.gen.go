// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INCurrencyAmountResolutionResult] class.
var (
	INCurrencyAmountResolutionResultClass     _INCurrencyAmountResolutionResultClass
	INCurrencyAmountResolutionResultClassOnce sync.Once
)

func getINCurrencyAmountResolutionResultClass() _INCurrencyAmountResolutionResultClass {
	INCurrencyAmountResolutionResultClassOnce.Do(func() {
		INCurrencyAmountResolutionResultClass = _INCurrencyAmountResolutionResultClass{objc.GetClass("INCurrencyAmountResolutionResult")}
	})
	return INCurrencyAmountResolutionResultClass
}

type _INCurrencyAmountResolutionResultClass struct {
	class objc.Class
}

// An interface definition for the [INCurrencyAmountResolutionResult] class.
type IINCurrencyAmountResolutionResult interface {
	IINIntentResolutionResult
}

// A resolution result for a currency amount associated with an intent.
//
// An object is what you return when resolving parameters containing an object. Use the creation method that best reflects your ability to successfully resolve the parameter. For additional resolution options, see .

// A resolution result for a currency amount associated with an intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCurrencyAmountResolutionResult
type INCurrencyAmountResolutionResult struct {
	INIntentResolutionResult
}

// INCurrencyAmountResolutionResultFrom constructs a [INCurrencyAmountResolutionResult] from an unsafe.Pointer.
//
// A resolution result for a currency amount associated with an intent.
func INCurrencyAmountResolutionResultFrom(ptr unsafe.Pointer) INCurrencyAmountResolutionResult {
	return INCurrencyAmountResolutionResult{
		INIntentResolutionResult: INIntentResolutionResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INCurrencyAmountResolutionResultClass) Alloc() INCurrencyAmountResolutionResult {
	rv := objc.Send[INCurrencyAmountResolutionResult](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCurrencyAmountResolutionResultClass) New() INCurrencyAmountResolutionResult {
	rv := objc.Send[INCurrencyAmountResolutionResult](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCurrencyAmountResolutionResult) Init() INCurrencyAmountResolutionResult {
	rv := objc.Send[INCurrencyAmountResolutionResult](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCurrencyAmountResolutionResult) Autorelease() INCurrencyAmountResolutionResult {
	rv := objc.Send[INCurrencyAmountResolutionResult](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCurrencyAmountResolutionResult creates a new INCurrencyAmountResolutionResult instance.
func NewINCurrencyAmountResolutionResult() INCurrencyAmountResolutionResult {
	return getINCurrencyAmountResolutionResultClass().New()
}
