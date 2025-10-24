// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSRNNDescriptor */


/* debug [class_header]: Header for MPSRNNDescriptor */
// The class instance for the [RNNDescriptor] class.
var (
	RNNDescriptorClass     _RNNDescriptorClass
	RNNDescriptorClassOnce sync.Once
)

func getRNNDescriptorClass() _RNNDescriptorClass {
	RNNDescriptorClassOnce.Do(func() {
		RNNDescriptorClass = _RNNDescriptorClass{objc.GetClass("MPSRNNDescriptor")}
	})
	return RNNDescriptorClass
}

type _RNNDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RNNDescriptor */
// An interface definition for the [RNNDescriptor] class.
type IRNNDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RNNDescriptor */
	// properties:
	UseLayerInputUnitTransformMode() objectivec.IObject
	SetUseLayerInputUnitTransformMode(value objectivec.IObject)
	OutputFeatureChannels() objectivec.IObject
	SetOutputFeatureChannels(value objectivec.IObject)
	InputFeatureChannels() objectivec.IObject
	SetInputFeatureChannels(value objectivec.IObject)
	LayerSequenceDirection() RNNSequenceDirection get set /* not a class type */
	SetLayerSequenceDirection(value RNNSequenceDirection get set /* not a class type */)
	UseFloat32Weights() objectivec.IObject
	SetUseFloat32Weights(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RNNDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RNNDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _RNNDescriptorClass) Alloc() RNNDescriptor {
	rv := objc.Send[RNNDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNDescriptorClass) New() RNNDescriptor {
	rv := objc.Send[RNNDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNDescriptor) Init() RNNDescriptor {
	rv := objc.Send[RNNDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNDescriptor) Autorelease() RNNDescriptor {
	rv := objc.Send[RNNDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNDescriptor creates a new RNNDescriptor instance.
func NewRNNDescriptor() RNNDescriptor {
	return getRNNDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RNNDescriptor */
// A description of a recursive neural network block or layer.


// A description of a recursive neural network block or layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNDescriptor
type RNNDescriptor struct {
	objectivec.Object
}

// RNNDescriptorFrom constructs a [RNNDescriptor] from an unsafe.Pointer.
//
// A description of a recursive neural network block or layer.
func RNNDescriptorFrom(ptr unsafe.Pointer) RNNDescriptor {
	return RNNDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RNNDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RNNDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RNNDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RNNDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RNNDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865687-uselayerinputunittransformmode
func (r_ RNNDescriptor) UseLayerInputUnitTransformMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("useLayerInputUnitTransformMode"))
	return rv
}/* debug [instance_properties/getter]: useLayerInputUnitTransformMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865687-uselayerinputunittransformmode
func (r_ RNNDescriptor) SetUseLayerInputUnitTransformMode(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUseLayerInputUnitTransformMode:"), value)
}/* debug [instance_properties/setter]: useLayerInputUnitTransformMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865702-outputfeaturechannels
func (r_ RNNDescriptor) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865702-outputfeaturechannels
func (r_ RNNDescriptor) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865707-inputfeaturechannels
func (r_ RNNDescriptor) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865707-inputfeaturechannels
func (r_ RNNDescriptor) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865730-layersequencedirection
func (r_ RNNDescriptor) LayerSequenceDirection() RNNSequenceDirection get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("layerSequenceDirection"))
	return rv
}/* debug [instance_properties/getter]: layerSequenceDirection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865730-layersequencedirection
func (r_ RNNDescriptor) SetLayerSequenceDirection(value RNNSequenceDirection get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayerSequenceDirection:"), value)
}/* debug [instance_properties/setter]: layerSequenceDirection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2881202-usefloat32weights
func (r_ RNNDescriptor) UseFloat32Weights() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("useFloat32Weights"))
	return rv
}/* debug [instance_properties/getter]: useFloat32Weights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2881202-usefloat32weights
func (r_ RNNDescriptor) SetUseFloat32Weights(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUseFloat32Weights:"), value)
}/* debug [instance_properties/setter]: useFloat32Weights */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSRNNDescriptor */



