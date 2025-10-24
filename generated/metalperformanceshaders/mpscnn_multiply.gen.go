// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNMultiply] class.
var (
	CNNMultiplyClass     _CNNMultiplyClass
	CNNMultiplyClassOnce sync.Once
)

func getCNNMultiplyClass() _CNNMultiplyClass {
	CNNMultiplyClassOnce.Do(func() {
		CNNMultiplyClass = _CNNMultiplyClass{objc.GetClass("MPSCNNMultiply")}
	})
	return CNNMultiplyClass
}

type _CNNMultiplyClass struct {
	class objc.Class
}





// An interface definition for the [CNNMultiply] class.
type ICNNMultiply interface {
	ICNNArithmetic
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNMultiplyClass) Alloc() CNNMultiply {
	rv := objc.Send[CNNMultiply](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNMultiplyClass) New() CNNMultiply {
	rv := objc.Send[CNNMultiply](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNMultiply) Init() CNNMultiply {
	rv := objc.Send[CNNMultiply](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNMultiply) Autorelease() CNNMultiply {
	rv := objc.Send[CNNMultiply](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNMultiply creates a new CNNMultiply instance.
func NewCNNMultiply() CNNMultiply {
	return getCNNMultiplyClass().New()
}





// A multiply operator.


// A multiply operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiply
type CNNMultiply struct {
	CNNArithmetic
}

// CNNMultiplyFrom constructs a [CNNMultiply] from an unsafe.Pointer.
//
// A multiply operator.
func CNNMultiplyFrom(ptr unsafe.Pointer) CNNMultiply {
	return CNNMultiply{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiply/2942507-initwithdevice
func NewCNNMultiplyWithDevice(device unsafe.Pointer) CNNMultiply {
	instance := getCNNMultiplyClass().Alloc()
	rv := objc.Send[CNNMultiply](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























