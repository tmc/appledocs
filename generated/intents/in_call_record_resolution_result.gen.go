// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INCallRecordResolutionResult] class.
var (
	INCallRecordResolutionResultClass     _INCallRecordResolutionResultClass
	INCallRecordResolutionResultClassOnce sync.Once
)

func getINCallRecordResolutionResultClass() _INCallRecordResolutionResultClass {
	INCallRecordResolutionResultClassOnce.Do(func() {
		INCallRecordResolutionResultClass = _INCallRecordResolutionResultClass{objc.GetClass("INCallRecordResolutionResult")}
	})
	return INCallRecordResolutionResultClass
}

type _INCallRecordResolutionResultClass struct {
	class objc.Class
}

// An interface definition for the [INCallRecordResolutionResult] class.
type IINCallRecordResolutionResult interface {
	IINIntentResolutionResult
}

// A resolution result for the details of a call.
//
// You return an object when resolving parameters containing an value. The can be different than the original . This flexibility allows app extensions to apply business logic constraints. Use to continue with a value. For additional resolution operators, see .


// A resolution result for the details of a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallRecordResolutionResult
type INCallRecordResolutionResult struct {
	INIntentResolutionResult
}

// INCallRecordResolutionResultFrom constructs a [INCallRecordResolutionResult] from an unsafe.Pointer.
//
// A resolution result for the details of a call.
func INCallRecordResolutionResultFrom(ptr unsafe.Pointer) INCallRecordResolutionResult {
	return INCallRecordResolutionResult{
		INIntentResolutionResult: INIntentResolutionResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INCallRecordResolutionResultClass) Alloc() INCallRecordResolutionResult {
	rv := objc.Send[INCallRecordResolutionResult](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCallRecordResolutionResultClass) New() INCallRecordResolutionResult {
	rv := objc.Send[INCallRecordResolutionResult](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCallRecordResolutionResult) Init() INCallRecordResolutionResult {
	rv := objc.Send[INCallRecordResolutionResult](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCallRecordResolutionResult) Autorelease() INCallRecordResolutionResult {
	rv := objc.Send[INCallRecordResolutionResult](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCallRecordResolutionResult creates a new INCallRecordResolutionResult instance.
func NewINCallRecordResolutionResult() INCallRecordResolutionResult {
	return getINCallRecordResolutionResultClass().New()
}




