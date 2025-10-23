// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INStartCallContactResolutionResult] class.
var (
	INStartCallContactResolutionResultClass     _INStartCallContactResolutionResultClass
	INStartCallContactResolutionResultClassOnce sync.Once
)

func getINStartCallContactResolutionResultClass() _INStartCallContactResolutionResultClass {
	INStartCallContactResolutionResultClassOnce.Do(func() {
		INStartCallContactResolutionResultClass = _INStartCallContactResolutionResultClass{objc.GetClass("INStartCallContactResolutionResult")}
	})
	return INStartCallContactResolutionResultClass
}

type _INStartCallContactResolutionResultClass struct {
	class objc.Class
}

// An interface definition for the [INStartCallContactResolutionResult] class.
type IINStartCallContactResolutionResult interface {
	objectivec.IObject
}

// A resolution result for the contact for the call.
//
// You return an object when resolving parameters containing an value. Use the creation method that best reflects your ability to resolve the parameter successfully. The resolved value can be different than the original . This flexibility allows app extensions to apply business logic constraints. For additional resolution operators, see .


// A resolution result for the contact for the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallContactResolutionResult
type INStartCallContactResolutionResult struct {
	objectivec.Object
}

// INStartCallContactResolutionResultFrom constructs a [INStartCallContactResolutionResult] from an unsafe.Pointer.
//
// A resolution result for the contact for the call.
func INStartCallContactResolutionResultFrom(ptr unsafe.Pointer) INStartCallContactResolutionResult {
	return INStartCallContactResolutionResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartCallContactResolutionResultClass) Alloc() INStartCallContactResolutionResult {
	rv := objc.Send[INStartCallContactResolutionResult](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartCallContactResolutionResultClass) New() INStartCallContactResolutionResult {
	rv := objc.Send[INStartCallContactResolutionResult](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartCallContactResolutionResult) Init() INStartCallContactResolutionResult {
	rv := objc.Send[INStartCallContactResolutionResult](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartCallContactResolutionResult) Autorelease() INStartCallContactResolutionResult {
	rv := objc.Send[INStartCallContactResolutionResult](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartCallContactResolutionResult creates a new INStartCallContactResolutionResult instance.
func NewINStartCallContactResolutionResult() INStartCallContactResolutionResult {
	return getINStartCallContactResolutionResultClass().New()
}




