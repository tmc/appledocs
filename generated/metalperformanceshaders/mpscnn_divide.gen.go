// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNDivide] class.
var (
	CNNDivideClass     _CNNDivideClass
	CNNDivideClassOnce sync.Once
)

func getCNNDivideClass() _CNNDivideClass {
	CNNDivideClassOnce.Do(func() {
		CNNDivideClass = _CNNDivideClass{objc.GetClass("MPSCNNDivide")}
	})
	return CNNDivideClass
}

type _CNNDivideClass struct {
	class objc.Class
}





// An interface definition for the [CNNDivide] class.
type ICNNDivide interface {
	ICNNArithmetic
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNDivideClass) Alloc() CNNDivide {
	rv := objc.Send[CNNDivide](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDivideClass) New() CNNDivide {
	rv := objc.Send[CNNDivide](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDivide) Init() CNNDivide {
	rv := objc.Send[CNNDivide](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDivide) Autorelease() CNNDivide {
	rv := objc.Send[CNNDivide](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDivide creates a new CNNDivide instance.
func NewCNNDivide() CNNDivide {
	return getCNNDivideClass().New()
}





// A division operator.


// A division operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDivide
type CNNDivide struct {
	CNNArithmetic
}

// CNNDivideFrom constructs a [CNNDivide] from an unsafe.Pointer.
//
// A division operator.
func CNNDivideFrom(ptr unsafe.Pointer) CNNDivide {
	return CNNDivide{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndivide/2942508-initwithdevice
func NewCNNDivideWithDevice(device unsafe.Pointer) CNNDivide {
	instance := getCNNDivideClass().Alloc()
	rv := objc.Send[CNNDivide](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























