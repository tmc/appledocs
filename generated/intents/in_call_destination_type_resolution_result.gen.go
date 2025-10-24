// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INCallDestinationTypeResolutionResult] class.
var (
	INCallDestinationTypeResolutionResultClass     _INCallDestinationTypeResolutionResultClass
	INCallDestinationTypeResolutionResultClassOnce sync.Once
)

func getINCallDestinationTypeResolutionResultClass() _INCallDestinationTypeResolutionResultClass {
	INCallDestinationTypeResolutionResultClassOnce.Do(func() {
		INCallDestinationTypeResolutionResultClass = _INCallDestinationTypeResolutionResultClass{objc.GetClass("INCallDestinationTypeResolutionResult")}
	})
	return INCallDestinationTypeResolutionResultClass
}

type _INCallDestinationTypeResolutionResultClass struct {
	class objc.Class
}

// An interface definition for the [INCallDestinationTypeResolutionResult] class.
type IINCallDestinationTypeResolutionResult interface {
	IINIntentResolutionResult
}

// A resolution result for the destination type of a call.
//
// You return an object when resolving parameters containing an value. Use the creation method that best reflects your ability to resolve the parameter successfully. For additional resolution operators, see .

// A resolution result for the destination type of a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallDestinationTypeResolutionResult
type INCallDestinationTypeResolutionResult struct {
	INIntentResolutionResult
}

// INCallDestinationTypeResolutionResultFrom constructs a [INCallDestinationTypeResolutionResult] from an unsafe.Pointer.
//
// A resolution result for the destination type of a call.
func INCallDestinationTypeResolutionResultFrom(ptr unsafe.Pointer) INCallDestinationTypeResolutionResult {
	return INCallDestinationTypeResolutionResult{
		INIntentResolutionResult: INIntentResolutionResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INCallDestinationTypeResolutionResultClass) Alloc() INCallDestinationTypeResolutionResult {
	rv := objc.Send[INCallDestinationTypeResolutionResult](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCallDestinationTypeResolutionResultClass) New() INCallDestinationTypeResolutionResult {
	rv := objc.Send[INCallDestinationTypeResolutionResult](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCallDestinationTypeResolutionResult) Init() INCallDestinationTypeResolutionResult {
	rv := objc.Send[INCallDestinationTypeResolutionResult](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCallDestinationTypeResolutionResult) Autorelease() INCallDestinationTypeResolutionResult {
	rv := objc.Send[INCallDestinationTypeResolutionResult](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCallDestinationTypeResolutionResult creates a new INCallDestinationTypeResolutionResult instance.
func NewINCallDestinationTypeResolutionResult() INCallDestinationTypeResolutionResult {
	return getINCallDestinationTypeResolutionResultClass().New()
}
