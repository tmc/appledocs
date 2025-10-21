// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ReshapeGradient] class.
var (
	ReshapeGradientClass     _ReshapeGradientClass
	ReshapeGradientClassOnce sync.Once
)

func getReshapeGradientClass() _ReshapeGradientClass {
	ReshapeGradientClassOnce.Do(func() {
		ReshapeGradientClass = _ReshapeGradientClass{objc.GetClass("MPSNNReshapeGradient")}
	})
	return ReshapeGradientClass
}

type _ReshapeGradientClass struct {
	class objc.Class
}

// An interface definition for the [ReshapeGradient] class.
type IReshapeGradient interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradient
type ReshapeGradient struct {
	objectivec.Object
}

// ReshapeGradientFrom constructs a [ReshapeGradient] from an unsafe.Pointer.
func ReshapeGradientFrom(ptr unsafe.Pointer) ReshapeGradient {
	return ReshapeGradient{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _ReshapeGradientClass) Alloc() ReshapeGradient {
	rv := objc.Send[ReshapeGradient](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReshapeGradientClass) New() ReshapeGradient {
	rv := objc.Send[ReshapeGradient](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReshapeGradient) Init() ReshapeGradient {
	rv := objc.Send[ReshapeGradient](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReshapeGradient) Autorelease() ReshapeGradient {
	rv := objc.Send[ReshapeGradient](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshapeGradient creates a new ReshapeGradient instance.
func NewReshapeGradient() ReshapeGradient {
	return getReshapeGradientClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradient/init(coder:device:)
func NewReshapeGradientWithCoderDevice(aDecoder foundation.ICoder, device objectivec.IObject) ReshapeGradient {
	instance := getReshapeGradientClass().Alloc()
	rv := objc.Send[ReshapeGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradient/init(device:)
func NewReshapeGradientWithDevice(device objectivec.IObject) ReshapeGradient {
	instance := getReshapeGradientClass().Alloc()
	rv := objc.Send[ReshapeGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



