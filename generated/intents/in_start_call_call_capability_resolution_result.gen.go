// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INStartCallCallCapabilityResolutionResult] class.
var (
	INStartCallCallCapabilityResolutionResultClass     _INStartCallCallCapabilityResolutionResultClass
	INStartCallCallCapabilityResolutionResultClassOnce sync.Once
)

func getINStartCallCallCapabilityResolutionResultClass() _INStartCallCallCapabilityResolutionResultClass {
	INStartCallCallCapabilityResolutionResultClassOnce.Do(func() {
		INStartCallCallCapabilityResolutionResultClass = _INStartCallCallCapabilityResolutionResultClass{objc.GetClass("INStartCallCallCapabilityResolutionResult")}
	})
	return INStartCallCallCapabilityResolutionResultClass
}

type _INStartCallCallCapabilityResolutionResultClass struct {
	class objc.Class
}

// An interface definition for the [INStartCallCallCapabilityResolutionResult] class.
type IINStartCallCallCapabilityResolutionResult interface {
	objectivec.IObject
}

// A resolution result for the call capability for the call.
//
// You return an object when resolving parameters containing an value. Use the creation method that best reflects your ability to resolve the parameter successfully. The resolved value can be different than the original . This flexibility allows app extensions to apply business logic constraints. For additional resolution operators, see .


// A resolution result for the call capability for the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallCallCapabilityResolutionResult
type INStartCallCallCapabilityResolutionResult struct {
	objectivec.Object
}

// INStartCallCallCapabilityResolutionResultFrom constructs a [INStartCallCallCapabilityResolutionResult] from an unsafe.Pointer.
//
// A resolution result for the call capability for the call.
func INStartCallCallCapabilityResolutionResultFrom(ptr unsafe.Pointer) INStartCallCallCapabilityResolutionResult {
	return INStartCallCallCapabilityResolutionResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartCallCallCapabilityResolutionResultClass) Alloc() INStartCallCallCapabilityResolutionResult {
	rv := objc.Send[INStartCallCallCapabilityResolutionResult](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartCallCallCapabilityResolutionResultClass) New() INStartCallCallCapabilityResolutionResult {
	rv := objc.Send[INStartCallCallCapabilityResolutionResult](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartCallCallCapabilityResolutionResult) Init() INStartCallCallCapabilityResolutionResult {
	rv := objc.Send[INStartCallCallCapabilityResolutionResult](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartCallCallCapabilityResolutionResult) Autorelease() INStartCallCallCapabilityResolutionResult {
	rv := objc.Send[INStartCallCallCapabilityResolutionResult](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartCallCallCapabilityResolutionResult creates a new INStartCallCallCapabilityResolutionResult instance.
func NewINStartCallCallCapabilityResolutionResult() INStartCallCallCapabilityResolutionResult {
	return getINStartCallCallCapabilityResolutionResultClass().New()
}




