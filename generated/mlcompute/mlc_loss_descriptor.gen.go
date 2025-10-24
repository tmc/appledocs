// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCLossDescriptor */


/* debug [class_header]: Header for MLCLossDescriptor */
// The class instance for the [CLossDescriptor] class.
var (
	CLossDescriptorClass     _CLossDescriptorClass
	CLossDescriptorClassOnce sync.Once
)

func getCLossDescriptorClass() _CLossDescriptorClass {
	CLossDescriptorClassOnce.Do(func() {
		CLossDescriptorClass = _CLossDescriptorClass{objc.GetClass("MLCLossDescriptor")}
	})
	return CLossDescriptorClass
}

type _CLossDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CLossDescriptor */
// An interface definition for the [CLossDescriptor] class.
type ICLossDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CLossDescriptor */
	// properties:
	ClassCount() uint
	Delta() float32
	Epsilon() float32
	LabelSmoothing() float32
	LossType() CLossType
	ReductionType() CReductionType
	Weight() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CLossDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CLossDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CLossDescriptorClass) Alloc() CLossDescriptor {
	rv := objc.Send[CLossDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CLossDescriptorClass) New() CLossDescriptor {
	rv := objc.Send[CLossDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CLossDescriptor) Init() CLossDescriptor {
	rv := objc.Send[CLossDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CLossDescriptor) Autorelease() CLossDescriptor {
	rv := objc.Send[CLossDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLossDescriptor creates a new CLossDescriptor instance.
func NewCLossDescriptor() CLossDescriptor {
	return getCLossDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CLossDescriptor */
// A configuration object you use to create a loss layer.


// A configuration object you use to create a loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor
type CLossDescriptor struct {
	objectivec.Object
}

// CLossDescriptorFrom constructs a [CLossDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create a loss layer.
func CLossDescriptorFrom(ptr unsafe.Pointer) CLossDescriptor {
	return CLossDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CLossDescriptor */

// Creates a loss descriptor with the loss function and reduction type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/init(type:reductionType:)
func NewCLossDescriptorWithTypeReductionType(lossType CLossType, reductionType CReductionType) CLossDescriptor {
	rv := objc.Send[CLossDescriptor](objc.ID(getCLossDescriptorClass().class), objc.Sel("descriptorWithType:reductionType:"), lossType, reductionType)
	return rv
}/* debug [class_init_methods/constructor]: NewCLossDescriptorWithTypeReductionType */


// Creates a loss descriptor with the loss function, reduction type, and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/init(type:reductionType:weight:)
func NewCLossDescriptorWithTypeReductionTypeWeight(lossType CLossType, reductionType CReductionType, weight float32) CLossDescriptor {
	rv := objc.Send[CLossDescriptor](objc.ID(getCLossDescriptorClass().class), objc.Sel("descriptorWithType:reductionType:weight:"), lossType, reductionType, weight)
	return rv
}/* debug [class_init_methods/constructor]: NewCLossDescriptorWithTypeReductionTypeWeight */


// Creates a loss descriptor with the loss function, reduction type, weight, label smoothing, and number of classes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/init(type:reductionType:weight:labelSmoothing:classCount:)
func NewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCount(lossType CLossType, reductionType CReductionType, weight float32, labelSmoothing float32, classCount uint) CLossDescriptor {
	rv := objc.Send[CLossDescriptor](objc.ID(getCLossDescriptorClass().class), objc.Sel("descriptorWithType:reductionType:weight:labelSmoothing:classCount:"), lossType, reductionType, weight, labelSmoothing, classCount)
	return rv
}/* debug [class_init_methods/constructor]: NewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCount */


// Creates a loss descriptor with the loss function, reduction type, weight, label smoothing, and number of classes, epsilon, and delta that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/init(type:reductionType:weight:labelSmoothing:classCount:epsilon:delta:)
func NewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCountEpsilonDelta(lossType CLossType, reductionType CReductionType, weight float32, labelSmoothing float32, classCount uint, epsilon float32, delta float32) CLossDescriptor {
	rv := objc.Send[CLossDescriptor](objc.ID(getCLossDescriptorClass().class), objc.Sel("descriptorWithType:reductionType:weight:labelSmoothing:classCount:epsilon:delta:"), lossType, reductionType, weight, labelSmoothing, classCount, epsilon, delta)
	return rv
}/* debug [class_init_methods/constructor]: NewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCountEpsilonDelta */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CLossDescriptor */

// Creates a loss descriptor with the loss function and reduction type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/init(type:reductionType:)
func (cc _CLossDescriptorClass) DescriptorWithTypeReductionType(lossType CLossType, reductionType CReductionType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:reductionType:"), lossType, reductionType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithTypeReductionType) */


// Creates a loss descriptor with the loss function, reduction type, and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/init(type:reductionType:weight:)
func (cc _CLossDescriptorClass) DescriptorWithTypeReductionTypeWeight(lossType CLossType, reductionType CReductionType, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:reductionType:weight:"), lossType, reductionType, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithTypeReductionTypeWeight) */


// Creates a loss descriptor with the loss function, reduction type, weight, label smoothing, and number of classes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/init(type:reductionType:weight:labelSmoothing:classCount:)
func (cc _CLossDescriptorClass) DescriptorWithTypeReductionTypeWeightLabelSmoothingClassCount(lossType CLossType, reductionType CReductionType, weight float32, labelSmoothing float32, classCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:reductionType:weight:labelSmoothing:classCount:"), lossType, reductionType, weight, labelSmoothing, classCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithTypeReductionTypeWeightLabelSmoothingClassCount) */


// Creates a loss descriptor with the loss function, reduction type, weight, label smoothing, and number of classes, epsilon, and delta that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/init(type:reductionType:weight:labelSmoothing:classCount:epsilon:delta:)
func (cc _CLossDescriptorClass) DescriptorWithTypeReductionTypeWeightLabelSmoothingClassCountEpsilonDelta(lossType CLossType, reductionType CReductionType, weight float32, labelSmoothing float32, classCount uint, epsilon float32, delta float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:reductionType:weight:labelSmoothing:classCount:epsilon:delta:"), lossType, reductionType, weight, labelSmoothing, classCount, epsilon, delta)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithTypeReductionTypeWeightLabelSmoothingClassCountEpsilonDelta) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CLossDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CLossDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CLossDescriptor */

// The number of classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/classCount
func (c_ CLossDescriptor) ClassCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("classCount"))
	return rv
}/* debug [instance_properties/getter]: classCount */


// The delta value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/delta
func (c_ CLossDescriptor) Delta() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// The epsilon value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/epsilon
func (c_ CLossDescriptor) Epsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// The value for label smoothing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/labelSmoothing
func (c_ CLossDescriptor) LabelSmoothing() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("labelSmoothing"))
	return rv
}/* debug [instance_properties/getter]: labelSmoothing */


// The loss function type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/lossType
func (c_ CLossDescriptor) LossType() CLossType {
	rv := objc.Send[CLossType](c_.ID, objc.Sel("lossType"))
	return rv
}/* debug [instance_properties/getter]: lossType */


// The reduction operation performed by the loss function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/reductionType
func (c_ CLossDescriptor) ReductionType() CReductionType {
	rv := objc.Send[CReductionType](c_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// The scale factor you apply to each element of a result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossDescriptor/weight
func (c_ CLossDescriptor) Weight() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCLossDescriptor */


