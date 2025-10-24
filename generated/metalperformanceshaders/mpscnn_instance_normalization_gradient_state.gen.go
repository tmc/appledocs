// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNInstanceNormalizationGradientState] class.
var (
	CNNInstanceNormalizationGradientStateClass     _CNNInstanceNormalizationGradientStateClass
	CNNInstanceNormalizationGradientStateClassOnce sync.Once
)

func getCNNInstanceNormalizationGradientStateClass() _CNNInstanceNormalizationGradientStateClass {
	CNNInstanceNormalizationGradientStateClassOnce.Do(func() {
		CNNInstanceNormalizationGradientStateClass = _CNNInstanceNormalizationGradientStateClass{objc.GetClass("MPSCNNInstanceNormalizationGradientState")}
	})
	return CNNInstanceNormalizationGradientStateClass
}

type _CNNInstanceNormalizationGradientStateClass struct {
	class objc.Class
}





// An interface definition for the [CNNInstanceNormalizationGradientState] class.
type ICNNInstanceNormalizationGradientState interface {
	IGradientState
	

	// properties:
	InstanceNormalization() IMPSCNNInstanceNormalization
	SetInstanceNormalization(value IMPSCNNInstanceNormalization)
	GradientForGamma() Buffer get /* not a class type */
	SetGradientForGamma(value Buffer get /* not a class type */)
	GradientForBeta() Buffer get /* not a class type */
	SetGradientForBeta(value Buffer get /* not a class type */)
	Gamma() Buffer get /* not a class type */
	SetGamma(value Buffer get /* not a class type */)
	Beta() Buffer get /* not a class type */
	SetBeta(value Buffer get /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNInstanceNormalizationGradientStateClass) Alloc() CNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNInstanceNormalizationGradientStateClass) New() CNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNInstanceNormalizationGradientState) Init() CNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNInstanceNormalizationGradientState) Autorelease() CNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNInstanceNormalizationGradientState creates a new CNNInstanceNormalizationGradientState instance.
func NewCNNInstanceNormalizationGradientState() CNNInstanceNormalizationGradientState {
	return getCNNInstanceNormalizationGradientStateClass().New()
}





// An object that stores information required to execute a gradient pass for instance normalization.


// An object that stores information required to execute a gradient pass for instance normalization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState
type CNNInstanceNormalizationGradientState struct {
	GradientState
}

// CNNInstanceNormalizationGradientStateFrom constructs a [CNNInstanceNormalizationGradientState] from an unsafe.Pointer.
//
// An object that stores information required to execute a gradient pass for instance normalization.
func CNNInstanceNormalizationGradientStateFrom(ptr unsafe.Pointer) CNNInstanceNormalizationGradientState {
	return CNNInstanceNormalizationGradientState{
		GradientState: GradientStateFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953924-instancenormalization
func (c_ CNNInstanceNormalizationGradientState) InstanceNormalization() IMPSCNNInstanceNormalization {
	rv := objc.Send[CNNInstanceNormalization](c_.ID, objc.Sel("instanceNormalization"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953924-instancenormalization
func (c_ CNNInstanceNormalizationGradientState) SetInstanceNormalization(value IMPSCNNInstanceNormalization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstanceNormalization:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953928-gradientforgamma
func (c_ CNNInstanceNormalizationGradientState) GradientForGamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForGamma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953928-gradientforgamma
func (c_ CNNInstanceNormalizationGradientState) SetGradientForGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForGamma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953930-gradientforbeta
func (c_ CNNInstanceNormalizationGradientState) GradientForBeta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForBeta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953930-gradientforbeta
func (c_ CNNInstanceNormalizationGradientState) SetGradientForBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForBeta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956161-gamma
func (c_ CNNInstanceNormalizationGradientState) Gamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gamma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956161-gamma
func (c_ CNNInstanceNormalizationGradientState) SetGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956162-beta
func (c_ CNNInstanceNormalizationGradientState) Beta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956162-beta
func (c_ CNNInstanceNormalizationGradientState) SetBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}








