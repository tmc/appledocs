// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PadGradient] class.
var (
	PadGradientClass     _PadGradientClass
	PadGradientClassOnce sync.Once
)

func getPadGradientClass() _PadGradientClass {
	PadGradientClassOnce.Do(func() {
		PadGradientClass = _PadGradientClass{objc.GetClass("MPSNNPadGradient")}
	})
	return PadGradientClass
}

type _PadGradientClass struct {
	class objc.Class
}

// An interface definition for the [PadGradient] class.
type IPadGradient interface {
	ICNNGradientKernel
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradient
type PadGradient struct {
	CNNGradientKernel
}

// PadGradientFrom constructs a [PadGradient] from an unsafe.Pointer.
func PadGradientFrom(ptr unsafe.Pointer) PadGradient {
	return PadGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PadGradientClass) Alloc() PadGradient {
	rv := objc.Send[PadGradient](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PadGradientClass) New() PadGradient {
	rv := objc.Send[PadGradient](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PadGradient) Init() PadGradient {
	rv := objc.Send[PadGradient](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PadGradient) Autorelease() PadGradient {
	rv := objc.Send[PadGradient](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPadGradient creates a new PadGradient instance.
func NewPadGradient() PadGradient {
	return getPadGradientClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradient/init(device:)
func NewPadGradientWithDevice(device objectivec.IObject) PadGradient {
	instance := getPadGradientClass().Alloc()
	rv := objc.Send[PadGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



