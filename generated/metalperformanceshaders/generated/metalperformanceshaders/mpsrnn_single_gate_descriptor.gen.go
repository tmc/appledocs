// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSRNNSingleGateDescriptor */


/* debug [class_header]: Header for MPSRNNSingleGateDescriptor */
// The class instance for the [RNNSingleGateDescriptor] class.
var (
	RNNSingleGateDescriptorClass     _RNNSingleGateDescriptorClass
	RNNSingleGateDescriptorClassOnce sync.Once
)

func getRNNSingleGateDescriptorClass() _RNNSingleGateDescriptorClass {
	RNNSingleGateDescriptorClassOnce.Do(func() {
		RNNSingleGateDescriptorClass = _RNNSingleGateDescriptorClass{objc.GetClass("MPSRNNSingleGateDescriptor")}
	})
	return RNNSingleGateDescriptorClass
}

type _RNNSingleGateDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RNNSingleGateDescriptor */
// An interface definition for the [RNNSingleGateDescriptor] class.
type IRNNSingleGateDescriptor interface {
	IRNNDescriptor
	
/* debug [class_interface_properties]: Properties for RNNSingleGateDescriptor */
	// properties:
	RecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	InputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	InputFeatureChannels() int
	SetInputFeatureChannels(value int)
	OutputFeatureChannels() int
	SetOutputFeatureChannels(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RNNSingleGateDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RNNSingleGateDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _RNNSingleGateDescriptorClass) Alloc() RNNSingleGateDescriptor {
	rv := objc.Send[RNNSingleGateDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNSingleGateDescriptorClass) New() RNNSingleGateDescriptor {
	rv := objc.Send[RNNSingleGateDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNSingleGateDescriptor) Init() RNNSingleGateDescriptor {
	rv := objc.Send[RNNSingleGateDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNSingleGateDescriptor) Autorelease() RNNSingleGateDescriptor {
	rv := objc.Send[RNNSingleGateDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNSingleGateDescriptor creates a new RNNSingleGateDescriptor instance.
func NewRNNSingleGateDescriptor() RNNSingleGateDescriptor {
	return getRNNSingleGateDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RNNSingleGateDescriptor */
// A description of a simple recurrent block or layer.
//
// The recurrent neural network (RNN) layer initialized with a transforms the input data (image or matrix) and previous output with a set of filters. Each produces one feature map in the new output data. You may provide the RNN unit with a single input or a sequence of inputs.


// A description of a simple recurrent block or layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSingleGateDescriptor
type RNNSingleGateDescriptor struct {
	RNNDescriptor
}

// RNNSingleGateDescriptorFrom constructs a [RNNSingleGateDescriptor] from an unsafe.Pointer.
//
// A description of a simple recurrent block or layer.
func RNNSingleGateDescriptorFrom(ptr unsafe.Pointer) RNNSingleGateDescriptor {
	return RNNSingleGateDescriptor{
		RNNDescriptor: RNNDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RNNSingleGateDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RNNSingleGateDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865703-creaternnsinglegatedescriptor
func (rc _RNNSingleGateDescriptorClass) CreateRNNSingleGateDescriptor() {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("createRNNSingleGateDescriptor"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateRNNSingleGateDescriptor) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865703-creaternnsinglegatedescriptorwit
func (rc _RNNSingleGateDescriptorClass) CreateRNNSingleGateDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("createRNNSingleGateDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateRNNSingleGateDescriptorWithInputFeatureChannelsOutputFeatureChannels) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RNNSingleGateDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RNNSingleGateDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RNNSingleGateDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865686-recurrentweights
func (r_ RNNSingleGateDescriptor) RecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("recurrentWeights"))
	return rv
}/* debug [instance_properties/getter]: recurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865686-recurrentweights
func (r_ RNNSingleGateDescriptor) SetRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecurrentWeights:"), value)
}/* debug [instance_properties/setter]: recurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865723-inputweights
func (r_ RNNSingleGateDescriptor) InputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("inputWeights"))
	return rv
}/* debug [instance_properties/getter]: inputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865723-inputweights
func (r_ RNNSingleGateDescriptor) SetInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputWeights:"), value)
}/* debug [instance_properties/setter]: inputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (r_ RNNSingleGateDescriptor) InputFeatureChannels() int {
	rv := objc.Send[int](r_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (r_ RNNSingleGateDescriptor) SetInputFeatureChannels(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (r_ RNNSingleGateDescriptor) OutputFeatureChannels() int {
	rv := objc.Send[int](r_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (r_ RNNSingleGateDescriptor) SetOutputFeatureChannels(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSRNNSingleGateDescriptor */



