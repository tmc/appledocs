// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNLossLabels] class.
var (
	CNNLossLabelsClass     _CNNLossLabelsClass
	CNNLossLabelsClassOnce sync.Once
)

func getCNNLossLabelsClass() _CNNLossLabelsClass {
	CNNLossLabelsClassOnce.Do(func() {
		CNNLossLabelsClass = _CNNLossLabelsClass{objc.GetClass("MPSCNNLossLabels")}
	})
	return CNNLossLabelsClass
}

type _CNNLossLabelsClass struct {
	class objc.Class
}





// An interface definition for the [CNNLossLabels] class.
type ICNNLossLabels interface {
	IState
	

	// properties:


	

	// methods:
	LossImage()
	LabelsImage()
	WeightsImage()


}





// Alloc allocates a new instance without initialization.
func (cc _CNNLossLabelsClass) Alloc() CNNLossLabels {
	rv := objc.Send[CNNLossLabels](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLossLabelsClass) New() CNNLossLabels {
	rv := objc.Send[CNNLossLabels](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLossLabels) Init() CNNLossLabels {
	rv := objc.Send[CNNLossLabels](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLossLabels) Autorelease() CNNLossLabels {
	rv := objc.Send[CNNLossLabels](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLossLabels creates a new CNNLossLabels instance.
func NewCNNLossLabels() CNNLossLabels {
	return getCNNLossLabelsClass().New()
}





// A class that stores the per-element weight buffer used by loss and gradient loss kernels.


// A class that stores the per-element weight buffer used by loss and gradient loss kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels
type CNNLossLabels struct {
	State
}

// CNNLossLabelsFrom constructs a [CNNLossLabels] from an unsafe.Pointer.
//
// A class that stores the per-element weight buffer used by loss and gradient loss kernels.
func CNNLossLabelsFrom(ptr unsafe.Pointer) CNNLossLabels {
	return CNNLossLabels{
		State: StateFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2951850-initwithdevice
func NewCNNLossLabelsWithDeviceLabelsDescriptor(device unsafe.Pointer, labelsDescriptor ICNNLossDataDescriptor) CNNLossLabels {
	instance := getCNNLossLabelsClass().Alloc()
	rv := objc.Send[CNNLossLabels](instance.ID, objc.Sel("initWithDevice:labelsDescriptor:"), device, labelsDescriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2951841-initwithdevice
func NewCNNLossLabelsWithDeviceLossImageSizeLabelsDescriptorWeightsDescriptor(device unsafe.Pointer, lossImageSize objc.IObject /* cross-framework: MTLSize */, labelsDescriptor ICNNLossDataDescriptor, weightsDescriptor ICNNLossDataDescriptor) CNNLossLabels {
	instance := getCNNLossLabelsClass().Alloc()
	rv := objc.Send[CNNLossLabels](instance.ID, objc.Sel("initWithDevice:lossImageSize:labelsDescriptor:weightsDescriptor:"), device, lossImageSize, labelsDescriptor, weightsDescriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/3114086-initwithdevice
func NewCNNLossLabelsWithDeviceLossImageSizeLabelsImageWeightsImage(device unsafe.Pointer, lossImageSize objc.IObject /* cross-framework: MTLSize */, labelsImage IImage, weightsImage IImage) CNNLossLabels {
	instance := getCNNLossLabelsClass().Alloc()
	rv := objc.Send[CNNLossLabels](instance.ID, objc.Sel("initWithDevice:lossImageSize:labelsImage:weightsImage:"), device, lossImageSize, labelsImage, weightsImage)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2951845-lossimage
func (c_ CNNLossLabels) LossImage() {
	objc.Send[objc.ID](c_.ID, objc.Sel("lossImage"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2976472-labelsimage
func (c_ CNNLossLabels) LabelsImage() {
	objc.Send[objc.ID](c_.ID, objc.Sel("labelsImage"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2976473-weightsimage
func (c_ CNNLossLabels) WeightsImage() {
	objc.Send[objc.ID](c_.ID, objc.Sel("weightsImage"))
}












