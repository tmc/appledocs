// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNMultiplyGradient] class.
var (
	CNNMultiplyGradientClass     _CNNMultiplyGradientClass
	CNNMultiplyGradientClassOnce sync.Once
)

func getCNNMultiplyGradientClass() _CNNMultiplyGradientClass {
	CNNMultiplyGradientClassOnce.Do(func() {
		CNNMultiplyGradientClass = _CNNMultiplyGradientClass{objc.GetClass("MPSCNNMultiplyGradient")}
	})
	return CNNMultiplyGradientClass
}

type _CNNMultiplyGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNMultiplyGradient] class.
type ICNNMultiplyGradient interface {
	ICNNArithmeticGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNMultiplyGradientClass) Alloc() CNNMultiplyGradient {
	rv := objc.Send[CNNMultiplyGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNMultiplyGradientClass) New() CNNMultiplyGradient {
	rv := objc.Send[CNNMultiplyGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNMultiplyGradient) Init() CNNMultiplyGradient {
	rv := objc.Send[CNNMultiplyGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNMultiplyGradient) Autorelease() CNNMultiplyGradient {
	rv := objc.Send[CNNMultiplyGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNMultiplyGradient creates a new CNNMultiplyGradient instance.
func NewCNNMultiplyGradient() CNNMultiplyGradient {
	return getCNNMultiplyGradientClass().New()
}





// A gradient multiply operator.


// A gradient multiply operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiplyGradient
type CNNMultiplyGradient struct {
	CNNArithmeticGradient
}

// CNNMultiplyGradientFrom constructs a [CNNMultiplyGradient] from an unsafe.Pointer.
//
// A gradient multiply operator.
func CNNMultiplyGradientFrom(ptr unsafe.Pointer) CNNMultiplyGradient {
	return CNNMultiplyGradient{
		CNNArithmeticGradient: CNNArithmeticGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiplygradient/2956164-initwithdevice
func NewCNNMultiplyGradientWithDeviceIsSecondarySourceFilter(device unsafe.Pointer, isSecondarySourceFilter bool) CNNMultiplyGradient {
	instance := getCNNMultiplyGradientClass().Alloc()
	rv := objc.Send[CNNMultiplyGradient](instance.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	rv.Autorelease()
	return rv
}



























