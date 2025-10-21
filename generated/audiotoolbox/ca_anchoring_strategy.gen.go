// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AnchoringStrategy] class.
var (
	AnchoringStrategyClass     _AnchoringStrategyClass
	AnchoringStrategyClassOnce sync.Once
)

func getAnchoringStrategyClass() _AnchoringStrategyClass {
	AnchoringStrategyClassOnce.Do(func() {
		AnchoringStrategyClass = _AnchoringStrategyClass{objc.GetClass("CAAnchoringStrategy")}
	})
	return AnchoringStrategyClass
}

type _AnchoringStrategyClass struct {
	class objc.Class
}

// An interface definition for the [AnchoringStrategy] class.
type IAnchoringStrategy interface {
	objectivec.IObject
}

// The center of a head-tracked spatial experience.
//
// The Objective-C version of the Swift type.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAAnchoringStrategy
type AnchoringStrategy struct {
	objectivec.Object
}

// AnchoringStrategyFrom constructs a [AnchoringStrategy] from an unsafe.Pointer.
//
// The center of a head-tracked spatial experience.
func AnchoringStrategyFrom(ptr unsafe.Pointer) AnchoringStrategy {
	return AnchoringStrategy{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AnchoringStrategyClass) Alloc() AnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AnchoringStrategyClass) New() AnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AnchoringStrategy) Init() AnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AnchoringStrategy) Autorelease() AnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAnchoringStrategy creates a new AnchoringStrategy instance.
func NewAnchoringStrategy() AnchoringStrategy {
	return getAnchoringStrategyClass().New()
}




