// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLossLabels */


/* debug [class_header]: Header for MPSCNNLossLabels */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLossLabels */
// An interface definition for the [CNNLossLabels] class.
type ICNNLossLabels interface {
	IState
	
/* debug [class_interface_properties]: Properties for CNNLossLabels */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLossLabels */
	// methods:
	LossImage()
	LabelsImage()
	WeightsImage()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLossLabels */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLossLabels */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLossLabels */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2951850-initwithdevice
func NewCNNLossLabelsWithDeviceLabelsDescriptor(device unsafe.Pointer, labelsDescriptor ICNNLossDataDescriptor) CNNLossLabels {
	instance := getCNNLossLabelsClass().Alloc()
	rv := objc.Send[CNNLossLabels](instance.ID, objc.Sel("initWithDevice:labelsDescriptor:"), device, labelsDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLossLabelsWithDeviceLabelsDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2951841-initwithdevice
func NewCNNLossLabelsWithDeviceLossImageSizeLabelsDescriptorWeightsDescriptor(device unsafe.Pointer, lossImageSize Size /* not a class type */, labelsDescriptor ICNNLossDataDescriptor, weightsDescriptor ICNNLossDataDescriptor) CNNLossLabels {
	instance := getCNNLossLabelsClass().Alloc()
	rv := objc.Send[CNNLossLabels](instance.ID, objc.Sel("initWithDevice:lossImageSize:labelsDescriptor:weightsDescriptor:"), device, lossImageSize, labelsDescriptor, weightsDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLossLabelsWithDeviceLossImageSizeLabelsDescriptorWeightsDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/3114086-initwithdevice
func NewCNNLossLabelsWithDeviceLossImageSizeLabelsImageWeightsImage(device unsafe.Pointer, lossImageSize Size /* not a class type */, labelsImage IImage, weightsImage IImage) CNNLossLabels {
	instance := getCNNLossLabelsClass().Alloc()
	rv := objc.Send[CNNLossLabels](instance.ID, objc.Sel("initWithDevice:lossImageSize:labelsImage:weightsImage:"), device, lossImageSize, labelsImage, weightsImage)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLossLabelsWithDeviceLossImageSizeLabelsImageWeightsImage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLossLabels */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLossLabels */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLossLabels */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2951845-lossimage
func (c_ CNNLossLabels) LossImage() {
	objc.Send[objc.ID](c_.ID, objc.Sel("lossImage"))
}/* debug [instance_methods/method]: LossImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2976472-labelsimage
func (c_ CNNLossLabels) LabelsImage() {
	objc.Send[objc.ID](c_.ID, objc.Sel("labelsImage"))
}/* debug [instance_methods/method]: LabelsImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2976473-weightsimage
func (c_ CNNLossLabels) WeightsImage() {
	objc.Send[objc.ID](c_.ID, objc.Sel("weightsImage"))
}/* debug [instance_methods/method]: WeightsImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLossLabels */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLossLabels */


