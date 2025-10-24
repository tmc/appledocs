// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNDropoutGradientState] class.
var (
	CNNDropoutGradientStateClass     _CNNDropoutGradientStateClass
	CNNDropoutGradientStateClassOnce sync.Once
)

func getCNNDropoutGradientStateClass() _CNNDropoutGradientStateClass {
	CNNDropoutGradientStateClassOnce.Do(func() {
		CNNDropoutGradientStateClass = _CNNDropoutGradientStateClass{objc.GetClass("MPSCNNDropoutGradientState")}
	})
	return CNNDropoutGradientStateClass
}

type _CNNDropoutGradientStateClass struct {
	class objc.Class
}





// An interface definition for the [CNNDropoutGradientState] class.
type ICNNDropoutGradientState interface {
	IGradientState
	

	// properties:


	

	// methods:
	MaskData()


}





// Alloc allocates a new instance without initialization.
func (cc _CNNDropoutGradientStateClass) Alloc() CNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDropoutGradientStateClass) New() CNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDropoutGradientState) Init() CNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDropoutGradientState) Autorelease() CNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDropoutGradientState creates a new CNNDropoutGradientState instance.
func NewCNNDropoutGradientState() CNNDropoutGradientState {
	return getCNNDropoutGradientStateClass().New()
}





// A class that stores the mask used by dropout and gradient dropout filters.


// A class that stores the mask used by dropout and gradient dropout filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientState
type CNNDropoutGradientState struct {
	GradientState
}

// CNNDropoutGradientStateFrom constructs a [CNNDropoutGradientState] from an unsafe.Pointer.
//
// A class that stores the mask used by dropout and gradient dropout filters.
func CNNDropoutGradientStateFrom(ptr unsafe.Pointer) CNNDropoutGradientState {
	return CNNDropoutGradientState{
		GradientState: GradientStateFrom(ptr),
	}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientstate/2942527-maskdata
func (c_ CNNDropoutGradientState) MaskData() {
	objc.Send[objc.ID](c_.ID, objc.Sel("maskData"))
}













