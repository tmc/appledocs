// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InitialGradient] class.
var (
	InitialGradientClass     _InitialGradientClass
	InitialGradientClassOnce sync.Once
)

func getInitialGradientClass() _InitialGradientClass {
	InitialGradientClassOnce.Do(func() {
		InitialGradientClass = _InitialGradientClass{objc.GetClass("MPSNNInitialGradient")}
	})
	return InitialGradientClass
}

type _InitialGradientClass struct {
	class objc.Class
}

// An interface definition for the [InitialGradient] class.
type IInitialGradient interface {
	ICNNKernel
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradient
type InitialGradient struct {
	CNNKernel
}

// InitialGradientFrom constructs a [InitialGradient] from an unsafe.Pointer.
func InitialGradientFrom(ptr unsafe.Pointer) InitialGradient {
	return InitialGradient{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InitialGradientClass) Alloc() InitialGradient {
	rv := objc.Send[InitialGradient](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InitialGradientClass) New() InitialGradient {
	rv := objc.Send[InitialGradient](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InitialGradient) Init() InitialGradient {
	rv := objc.Send[InitialGradient](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InitialGradient) Autorelease() InitialGradient {
	rv := objc.Send[InitialGradient](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInitialGradient creates a new InitialGradient instance.
func NewInitialGradient() InitialGradient {
	return getInitialGradientClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradient/init(device:)
func NewInitialGradientWithDevice(device objectivec.IObject) InitialGradient {
	instance := getInitialGradientClass().Alloc()
	rv := objc.Send[InitialGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



