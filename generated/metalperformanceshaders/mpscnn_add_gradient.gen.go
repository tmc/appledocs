// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNAddGradient] class.
var (
	CNNAddGradientClass     _CNNAddGradientClass
	CNNAddGradientClassOnce sync.Once
)

func getCNNAddGradientClass() _CNNAddGradientClass {
	CNNAddGradientClassOnce.Do(func() {
		CNNAddGradientClass = _CNNAddGradientClass{objc.GetClass("MPSCNNAddGradient")}
	})
	return CNNAddGradientClass
}

type _CNNAddGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNAddGradient] class.
type ICNNAddGradient interface {
	ICNNArithmeticGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNAddGradientClass) Alloc() CNNAddGradient {
	rv := objc.Send[CNNAddGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNAddGradientClass) New() CNNAddGradient {
	rv := objc.Send[CNNAddGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNAddGradient) Init() CNNAddGradient {
	rv := objc.Send[CNNAddGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNAddGradient) Autorelease() CNNAddGradient {
	rv := objc.Send[CNNAddGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNAddGradient creates a new CNNAddGradient instance.
func NewCNNAddGradient() CNNAddGradient {
	return getCNNAddGradientClass().New()
}





// A gradient addition operator.


// A gradient addition operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAddGradient
type CNNAddGradient struct {
	CNNArithmeticGradient
}

// CNNAddGradientFrom constructs a [CNNAddGradient] from an unsafe.Pointer.
//
// A gradient addition operator.
func CNNAddGradientFrom(ptr unsafe.Pointer) CNNAddGradient {
	return CNNAddGradient{
		CNNArithmeticGradient: CNNArithmeticGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnaddgradient/2956163-initwithdevice
func NewCNNAddGradientWithDeviceIsSecondarySourceFilter(device unsafe.Pointer, isSecondarySourceFilter bool) CNNAddGradient {
	instance := getCNNAddGradientClass().Alloc()
	rv := objc.Send[CNNAddGradient](instance.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	rv.Autorelease()
	return rv
}



























