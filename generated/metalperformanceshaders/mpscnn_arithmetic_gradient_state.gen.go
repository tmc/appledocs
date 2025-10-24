// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNArithmeticGradientState] class.
var (
	CNNArithmeticGradientStateClass     _CNNArithmeticGradientStateClass
	CNNArithmeticGradientStateClassOnce sync.Once
)

func getCNNArithmeticGradientStateClass() _CNNArithmeticGradientStateClass {
	CNNArithmeticGradientStateClassOnce.Do(func() {
		CNNArithmeticGradientStateClass = _CNNArithmeticGradientStateClass{objc.GetClass("MPSCNNArithmeticGradientState")}
	})
	return CNNArithmeticGradientStateClass
}

type _CNNArithmeticGradientStateClass struct {
	class objc.Class
}





// An interface definition for the [CNNArithmeticGradientState] class.
type ICNNArithmeticGradientState interface {
	IBinaryGradientState
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNArithmeticGradientStateClass) Alloc() CNNArithmeticGradientState {
	rv := objc.Send[CNNArithmeticGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNArithmeticGradientStateClass) New() CNNArithmeticGradientState {
	rv := objc.Send[CNNArithmeticGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNArithmeticGradientState) Init() CNNArithmeticGradientState {
	rv := objc.Send[CNNArithmeticGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNArithmeticGradientState) Autorelease() CNNArithmeticGradientState {
	rv := objc.Send[CNNArithmeticGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNArithmeticGradientState creates a new CNNArithmeticGradientState instance.
func NewCNNArithmeticGradientState() CNNArithmeticGradientState {
	return getCNNArithmeticGradientStateClass().New()
}





// An object that stores the clamp mask used by gradient arithmetic operators.


// An object that stores the clamp mask used by gradient arithmetic operators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradientState
type CNNArithmeticGradientState struct {
	BinaryGradientState
}

// CNNArithmeticGradientStateFrom constructs a [CNNArithmeticGradientState] from an unsafe.Pointer.
//
// An object that stores the clamp mask used by gradient arithmetic operators.
func CNNArithmeticGradientStateFrom(ptr unsafe.Pointer) CNNArithmeticGradientState {
	return CNNArithmeticGradientState{
		BinaryGradientState: BinaryGradientStateFrom(ptr),
	}
}































