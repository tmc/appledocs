// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INIntentResolutionResult] class.
var (
	INIntentResolutionResultClass     _INIntentResolutionResultClass
	INIntentResolutionResultClassOnce sync.Once
)

func getINIntentResolutionResultClass() _INIntentResolutionResultClass {
	INIntentResolutionResultClassOnce.Do(func() {
		INIntentResolutionResultClass = _INIntentResolutionResultClass{objc.GetClass("INIntentResolutionResult")}
	})
	return INIntentResolutionResultClass
}

type _INIntentResolutionResultClass struct {
	class objc.Class
}

// An interface definition for the [INIntentResolutionResult] class.
type IINIntentResolutionResult interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A resolution result for a parameter of an intent object.
//
// An object describes how your app resolves a parameter of an intent object. This class is abstract and you don’t create instances of it directly. Instead, you use class methods to instantiate the appropriate subclass whose type matches the type of data that you’re trying to resolve. This class defines methods that are common to all resolution result objects. During the resolution phase of a request, each parameter you resolve requires a resolution result object of a specific type. When creating that resolution result object, use the class method that represents your resolution. Subclasses of define methods for returning a successful resolution and may contain other methods for confirming values or disambiguating from among several possible values. Use the methods of this class when a value isn’t required, when a required value is missing, or when the value specified by the user doesn’t correspond to a supported solution. When creating resolution result objects, always strive toward successful resolutions. If needed, take advantage of information you’ve regarding the user’s behavior patterns or interactions with your app. For example, if the user always has the same workout goals, use the previous workout goals as default values rather than asking the user to specify goal information.

// A resolution result for a parameter of an intent object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntentResolutionResult
type INIntentResolutionResult struct {
	objectivec.Object
}

// INIntentResolutionResultFrom constructs a [INIntentResolutionResult] from an unsafe.Pointer.
//
// A resolution result for a parameter of an intent object.
func INIntentResolutionResultFrom(ptr unsafe.Pointer) INIntentResolutionResult {
	return INIntentResolutionResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INIntentResolutionResultClass) Alloc() INIntentResolutionResult {
	rv := objc.Send[INIntentResolutionResult](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INIntentResolutionResultClass) New() INIntentResolutionResult {
	rv := objc.Send[INIntentResolutionResult](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INIntentResolutionResult) Init() INIntentResolutionResult {
	rv := objc.Send[INIntentResolutionResult](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INIntentResolutionResult) Autorelease() INIntentResolutionResult {
	rv := objc.Send[INIntentResolutionResult](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINIntentResolutionResult creates a new INIntentResolutionResult instance.
func NewINIntentResolutionResult() INIntentResolutionResult {
	return getINIntentResolutionResultClass().New()
}
