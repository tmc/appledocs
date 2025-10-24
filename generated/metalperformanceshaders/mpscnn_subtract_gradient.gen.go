// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNSubtractGradient] class.
var (
	CNNSubtractGradientClass     _CNNSubtractGradientClass
	CNNSubtractGradientClassOnce sync.Once
)

func getCNNSubtractGradientClass() _CNNSubtractGradientClass {
	CNNSubtractGradientClassOnce.Do(func() {
		CNNSubtractGradientClass = _CNNSubtractGradientClass{objc.GetClass("MPSCNNSubtractGradient")}
	})
	return CNNSubtractGradientClass
}

type _CNNSubtractGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNSubtractGradient] class.
type ICNNSubtractGradient interface {
	ICNNArithmeticGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNSubtractGradientClass) Alloc() CNNSubtractGradient {
	rv := objc.Send[CNNSubtractGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSubtractGradientClass) New() CNNSubtractGradient {
	rv := objc.Send[CNNSubtractGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSubtractGradient) Init() CNNSubtractGradient {
	rv := objc.Send[CNNSubtractGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSubtractGradient) Autorelease() CNNSubtractGradient {
	rv := objc.Send[CNNSubtractGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSubtractGradient creates a new CNNSubtractGradient instance.
func NewCNNSubtractGradient() CNNSubtractGradient {
	return getCNNSubtractGradientClass().New()
}





// A gradient subtraction operator.


// A gradient subtraction operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtractGradient
type CNNSubtractGradient struct {
	CNNArithmeticGradient
}

// CNNSubtractGradientFrom constructs a [CNNSubtractGradient] from an unsafe.Pointer.
//
// A gradient subtraction operator.
func CNNSubtractGradientFrom(ptr unsafe.Pointer) CNNSubtractGradient {
	return CNNSubtractGradient{
		CNNArithmeticGradient: CNNArithmeticGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubtractgradient/2956165-initwithdevice
func NewCNNSubtractGradientWithDeviceIsSecondarySourceFilter(device unsafe.Pointer, isSecondarySourceFilter bool) CNNSubtractGradient {
	instance := getCNNSubtractGradientClass().Alloc()
	rv := objc.Send[CNNSubtractGradient](instance.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	rv.Autorelease()
	return rv
}



























