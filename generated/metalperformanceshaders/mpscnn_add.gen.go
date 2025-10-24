// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNAdd] class.
var (
	CNNAddClass     _CNNAddClass
	CNNAddClassOnce sync.Once
)

func getCNNAddClass() _CNNAddClass {
	CNNAddClassOnce.Do(func() {
		CNNAddClass = _CNNAddClass{objc.GetClass("MPSCNNAdd")}
	})
	return CNNAddClass
}

type _CNNAddClass struct {
	class objc.Class
}





// An interface definition for the [CNNAdd] class.
type ICNNAdd interface {
	ICNNArithmetic
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNAddClass) Alloc() CNNAdd {
	rv := objc.Send[CNNAdd](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNAddClass) New() CNNAdd {
	rv := objc.Send[CNNAdd](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNAdd) Init() CNNAdd {
	rv := objc.Send[CNNAdd](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNAdd) Autorelease() CNNAdd {
	rv := objc.Send[CNNAdd](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNAdd creates a new CNNAdd instance.
func NewCNNAdd() CNNAdd {
	return getCNNAddClass().New()
}





// An addition operator.


// An addition operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAdd
type CNNAdd struct {
	CNNArithmetic
}

// CNNAddFrom constructs a [CNNAdd] from an unsafe.Pointer.
//
// An addition operator.
func CNNAddFrom(ptr unsafe.Pointer) CNNAdd {
	return CNNAdd{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnadd/2942501-initwithdevice
func NewCNNAddWithDevice(device unsafe.Pointer) CNNAdd {
	instance := getCNNAddClass().Alloc()
	rv := objc.Send[CNNAdd](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























