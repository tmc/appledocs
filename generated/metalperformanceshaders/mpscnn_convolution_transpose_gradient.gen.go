// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNConvolutionTransposeGradient] class.
var (
	CNNConvolutionTransposeGradientClass     _CNNConvolutionTransposeGradientClass
	CNNConvolutionTransposeGradientClassOnce sync.Once
)

func getCNNConvolutionTransposeGradientClass() _CNNConvolutionTransposeGradientClass {
	CNNConvolutionTransposeGradientClassOnce.Do(func() {
		CNNConvolutionTransposeGradientClass = _CNNConvolutionTransposeGradientClass{objc.GetClass("MPSCNNConvolutionTransposeGradient")}
	})
	return CNNConvolutionTransposeGradientClass
}

type _CNNConvolutionTransposeGradientClass struct {
	class objc.Class
}

// An interface definition for the [CNNConvolutionTransposeGradient] class.
type ICNNConvolutionTransposeGradient interface {
	objectivec.IObject
	ReloadWeightsAndBiasesFromDataSource()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient
type CNNConvolutionTransposeGradient struct {
	objectivec.Object
}

// CNNConvolutionTransposeGradientFrom constructs a [CNNConvolutionTransposeGradient] from an unsafe.Pointer.
func CNNConvolutionTransposeGradientFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradient {
	return CNNConvolutionTransposeGradient{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeGradientClass) Alloc() CNNConvolutionTransposeGradient {
	rv := objc.Send[CNNConvolutionTransposeGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNConvolutionTransposeGradientClass) New() CNNConvolutionTransposeGradient {
	rv := objc.Send[CNNConvolutionTransposeGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTransposeGradient) Init() CNNConvolutionTransposeGradient {
	rv := objc.Send[CNNConvolutionTransposeGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTransposeGradient) Autorelease() CNNConvolutionTransposeGradient {
	rv := objc.Send[CNNConvolutionTransposeGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTransposeGradient creates a new CNNConvolutionTransposeGradient instance.
func NewCNNConvolutionTransposeGradient() CNNConvolutionTransposeGradient {
	return getCNNConvolutionTransposeGradientClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/reloadWeightsAndBiasesFromDataSource()
func (c_ CNNConvolutionTransposeGradient) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/gradientoption
func (c_ CNNConvolutionTransposeGradient) GradientOption() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("gradientOption"))
	return rv
}


// SetGradientOption sets the value of the gradientOption property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/gradientoption
func (c_ CNNConvolutionTransposeGradient) SetGradientOption(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientOption:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/sourcegradientfeaturechannels
func (c_ CNNConvolutionTransposeGradient) SourceGradientFeatureChannels() int {
	rv := objc.Send[int](c_.ID, objc.Sel("sourceGradientFeatureChannels"))
	return rv
}


// SetSourceGradientFeatureChannels sets the value of the sourceGradientFeatureChannels property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/sourcegradientfeaturechannels
func (c_ CNNConvolutionTransposeGradient) SetSourceGradientFeatureChannels(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceGradientFeatureChannels:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/datasource
func (c_ CNNConvolutionTransposeGradient) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dataSource"))
	return rv
}


// SetDataSource sets the value of the dataSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/datasource
func (c_ CNNConvolutionTransposeGradient) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/groups
func (c_ CNNConvolutionTransposeGradient) Groups() int {
	rv := objc.Send[int](c_.ID, objc.Sel("groups"))
	return rv
}


// SetGroups sets the value of the groups property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/groups
func (c_ CNNConvolutionTransposeGradient) SetGroups(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/sourceimagefeaturechannels
func (c_ CNNConvolutionTransposeGradient) SourceImageFeatureChannels() int {
	rv := objc.Send[int](c_.ID, objc.Sel("sourceImageFeatureChannels"))
	return rv
}


// SetSourceImageFeatureChannels sets the value of the sourceImageFeatureChannels property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/sourceimagefeaturechannels
func (c_ CNNConvolutionTransposeGradient) SetSourceImageFeatureChannels(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceImageFeatureChannels:"), value)
}



