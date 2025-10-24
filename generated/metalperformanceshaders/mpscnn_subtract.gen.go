// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNSubtract] class.
var (
	CNNSubtractClass     _CNNSubtractClass
	CNNSubtractClassOnce sync.Once
)

func getCNNSubtractClass() _CNNSubtractClass {
	CNNSubtractClassOnce.Do(func() {
		CNNSubtractClass = _CNNSubtractClass{objc.GetClass("MPSCNNSubtract")}
	})
	return CNNSubtractClass
}

type _CNNSubtractClass struct {
	class objc.Class
}





// An interface definition for the [CNNSubtract] class.
type ICNNSubtract interface {
	ICNNArithmetic
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNSubtractClass) Alloc() CNNSubtract {
	rv := objc.Send[CNNSubtract](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSubtractClass) New() CNNSubtract {
	rv := objc.Send[CNNSubtract](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSubtract) Init() CNNSubtract {
	rv := objc.Send[CNNSubtract](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSubtract) Autorelease() CNNSubtract {
	rv := objc.Send[CNNSubtract](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSubtract creates a new CNNSubtract instance.
func NewCNNSubtract() CNNSubtract {
	return getCNNSubtractClass().New()
}





// A subtraction operator.


// A subtraction operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtract
type CNNSubtract struct {
	CNNArithmetic
}

// CNNSubtractFrom constructs a [CNNSubtract] from an unsafe.Pointer.
//
// A subtraction operator.
func CNNSubtractFrom(ptr unsafe.Pointer) CNNSubtract {
	return CNNSubtract{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubtract/2942503-initwithdevice
func NewCNNSubtractWithDevice(device unsafe.Pointer) CNNSubtract {
	instance := getCNNSubtractClass().Alloc()
	rv := objc.Send[CNNSubtract](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























