// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AutomaticAnchoringStrategy] class.
var (
	AutomaticAnchoringStrategyClass     _AutomaticAnchoringStrategyClass
	AutomaticAnchoringStrategyClassOnce sync.Once
)

func getAutomaticAnchoringStrategyClass() _AutomaticAnchoringStrategyClass {
	AutomaticAnchoringStrategyClassOnce.Do(func() {
		AutomaticAnchoringStrategyClass = _AutomaticAnchoringStrategyClass{objc.GetClass("CAAutomaticAnchoringStrategy")}
	})
	return AutomaticAnchoringStrategyClass
}

type _AutomaticAnchoringStrategyClass struct {
	class objc.Class
}

// An interface definition for the [AutomaticAnchoringStrategy] class.
type IAutomaticAnchoringStrategy interface {
	IAnchoringStrategy
}

// A system-defined anchoring strategy.


// A system-defined anchoring strategy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAAutomaticAnchoringStrategy
type AutomaticAnchoringStrategy struct {
	AnchoringStrategy
}

// AutomaticAnchoringStrategyFrom constructs a [AutomaticAnchoringStrategy] from an unsafe.Pointer.
//
// A system-defined anchoring strategy.
func AutomaticAnchoringStrategyFrom(ptr unsafe.Pointer) AutomaticAnchoringStrategy {
	return AutomaticAnchoringStrategy{
		AnchoringStrategy: AnchoringStrategyFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AutomaticAnchoringStrategyClass) Alloc() AutomaticAnchoringStrategy {
	rv := objc.Send[AutomaticAnchoringStrategy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AutomaticAnchoringStrategyClass) New() AutomaticAnchoringStrategy {
	rv := objc.Send[AutomaticAnchoringStrategy](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AutomaticAnchoringStrategy) Init() AutomaticAnchoringStrategy {
	rv := objc.Send[AutomaticAnchoringStrategy](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AutomaticAnchoringStrategy) Autorelease() AutomaticAnchoringStrategy {
	rv := objc.Send[AutomaticAnchoringStrategy](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAutomaticAnchoringStrategy creates a new AutomaticAnchoringStrategy instance.
func NewAutomaticAnchoringStrategy() AutomaticAnchoringStrategy {
	return getAutomaticAnchoringStrategyClass().New()
}




