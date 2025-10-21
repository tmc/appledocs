// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INDailyRoutineRelevanceProvider] class.
var (
	INDailyRoutineRelevanceProviderClass     _INDailyRoutineRelevanceProviderClass
	INDailyRoutineRelevanceProviderClassOnce sync.Once
)

func getINDailyRoutineRelevanceProviderClass() _INDailyRoutineRelevanceProviderClass {
	INDailyRoutineRelevanceProviderClassOnce.Do(func() {
		INDailyRoutineRelevanceProviderClass = _INDailyRoutineRelevanceProviderClass{objc.GetClass("INDailyRoutineRelevanceProvider")}
	})
	return INDailyRoutineRelevanceProviderClass
}

type _INDailyRoutineRelevanceProviderClass struct {
	class objc.Class
}

// An interface definition for the [INDailyRoutineRelevanceProvider] class.
type IINDailyRoutineRelevanceProvider interface {
	objectivec.IObject
}

// The provider class that specifies a relevant daily routine.
//
// Ask the user for permission to use their location before providing shortcuts to that include a daily routine relevance provider. If the user gives your app permission to access their location Always, shortcuts your app provides can influence widget stacks.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INDailyRoutineRelevanceProvider
type INDailyRoutineRelevanceProvider struct {
	objectivec.Object
}

// INDailyRoutineRelevanceProviderFrom constructs a [INDailyRoutineRelevanceProvider] from an unsafe.Pointer.
//
// The provider class that specifies a relevant daily routine.
func INDailyRoutineRelevanceProviderFrom(ptr unsafe.Pointer) INDailyRoutineRelevanceProvider {
	return INDailyRoutineRelevanceProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INDailyRoutineRelevanceProviderClass) Alloc() INDailyRoutineRelevanceProvider {
	rv := objc.Send[INDailyRoutineRelevanceProvider](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INDailyRoutineRelevanceProviderClass) New() INDailyRoutineRelevanceProvider {
	rv := objc.Send[INDailyRoutineRelevanceProvider](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INDailyRoutineRelevanceProvider) Init() INDailyRoutineRelevanceProvider {
	rv := objc.Send[INDailyRoutineRelevanceProvider](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INDailyRoutineRelevanceProvider) Autorelease() INDailyRoutineRelevanceProvider {
	rv := objc.Send[INDailyRoutineRelevanceProvider](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINDailyRoutineRelevanceProvider creates a new INDailyRoutineRelevanceProvider instance.
func NewINDailyRoutineRelevanceProvider() INDailyRoutineRelevanceProvider {
	return getINDailyRoutineRelevanceProviderClass().New()
}




// Creates a daily routine relevance provider with the specified situation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INDailyRoutineRelevanceProvider/init(situation:)
func NewINDailyRoutineRelevanceProviderWithSituation(situation unsafe.Pointer) INDailyRoutineRelevanceProvider {
	instance := getINDailyRoutineRelevanceProviderClass().Alloc()
	rv := objc.Send[INDailyRoutineRelevanceProvider](instance.ID, objc.Sel("initWithSituation:"), situation)
	rv.Autorelease()
	return rv
}


// The relevant daily routine situation for the provider.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INDailyRoutineRelevanceProvider/situation-swift.property
func (i_ INDailyRoutineRelevanceProvider) Situation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("situation"))
	return rv
}


