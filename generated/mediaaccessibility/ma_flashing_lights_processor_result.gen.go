// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MAFlashingLightsProcessorResult] class.
var (
	MAFlashingLightsProcessorResultClass     _MAFlashingLightsProcessorResultClass
	MAFlashingLightsProcessorResultClassOnce sync.Once
)

func getMAFlashingLightsProcessorResultClass() _MAFlashingLightsProcessorResultClass {
	MAFlashingLightsProcessorResultClassOnce.Do(func() {
		MAFlashingLightsProcessorResultClass = _MAFlashingLightsProcessorResultClass{objc.GetClass("MAFlashingLightsProcessorResult")}
	})
	return MAFlashingLightsProcessorResultClass
}

type _MAFlashingLightsProcessorResultClass struct {
	class objc.Class
}

// An interface definition for the [MAFlashingLightsProcessorResult] class.
type IMAFlashingLightsProcessorResult interface {
	objectivec.IObject
}

// An object that reports the result of the flashing lights processor.
//
// An object is the result of calling . This object indicates whether the method successfully processed the input surface, the intensity of flashing lights in the input surface, and the amount of mitigation in the output surface.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessorResult
type MAFlashingLightsProcessorResult struct {
	objectivec.Object
}

// MAFlashingLightsProcessorResultFrom constructs a [MAFlashingLightsProcessorResult] from an unsafe.Pointer.
//
// An object that reports the result of the flashing lights processor.
func MAFlashingLightsProcessorResultFrom(ptr unsafe.Pointer) MAFlashingLightsProcessorResult {
	return MAFlashingLightsProcessorResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MAFlashingLightsProcessorResultClass) Alloc() MAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MAFlashingLightsProcessorResultClass) New() MAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MAFlashingLightsProcessorResult) Init() MAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MAFlashingLightsProcessorResult) Autorelease() MAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMAFlashingLightsProcessorResult creates a new MAFlashingLightsProcessorResult instance.
func NewMAFlashingLightsProcessorResult() MAFlashingLightsProcessorResult {
	return getMAFlashingLightsProcessorResultClass().New()
}




