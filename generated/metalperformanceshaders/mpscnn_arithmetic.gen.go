// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNArithmetic] class.
var (
	CNNArithmeticClass     _CNNArithmeticClass
	CNNArithmeticClassOnce sync.Once
)

func getCNNArithmeticClass() _CNNArithmeticClass {
	CNNArithmeticClassOnce.Do(func() {
		CNNArithmeticClass = _CNNArithmeticClass{objc.GetClass("MPSCNNArithmetic")}
	})
	return CNNArithmeticClass
}

type _CNNArithmeticClass struct {
	class objc.Class
}

// An interface definition for the [CNNArithmetic] class.
type ICNNArithmetic interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type CNNArithmetic struct {
	objectivec.Object
}

// CNNArithmeticFrom constructs a [CNNArithmetic] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func CNNArithmeticFrom(ptr unsafe.Pointer) CNNArithmetic {
	return CNNArithmetic{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNArithmeticClass) Alloc() CNNArithmetic {
	rv := objc.Send[CNNArithmetic](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNArithmeticClass) New() CNNArithmetic {
	rv := objc.Send[CNNArithmetic](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNArithmetic) Init() CNNArithmetic {
	rv := objc.Send[CNNArithmetic](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNArithmetic) Autorelease() CNNArithmetic {
	rv := objc.Send[CNNArithmetic](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNArithmetic creates a new CNNArithmetic instance.
func NewCNNArithmetic() CNNArithmetic {
	return getCNNArithmeticClass().New()
}




