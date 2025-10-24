// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionGradient */


/* debug [class_header]: Header for MPSCNNConvolutionGradient */
// The class instance for the [CNNConvolutionGradient] class.
var (
	CNNConvolutionGradientClass     _CNNConvolutionGradientClass
	CNNConvolutionGradientClassOnce sync.Once
)

func getCNNConvolutionGradientClass() _CNNConvolutionGradientClass {
	CNNConvolutionGradientClassOnce.Do(func() {
		CNNConvolutionGradientClass = _CNNConvolutionGradientClass{objc.GetClass("MPSCNNConvolutionGradient")}
	})
	return CNNConvolutionGradientClass
}

type _CNNConvolutionGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionGradient */
// An interface definition for the [CNNConvolutionGradient] class.
type ICNNConvolutionGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNConvolutionGradient */
	// properties:
	Groups() objectivec.IObject
	SetGroups(value objectivec.IObject)
	GradientOption() CNNConvolutionGradientOption get set /* not a class type */
	SetGradientOption(value CNNConvolutionGradientOption get set /* not a class type */)
	SourceGradientFeatureChannels() objectivec.IObject
	SetSourceGradientFeatureChannels(value objectivec.IObject)
	SourceImageFeatureChannels() objectivec.IObject
	SetSourceImageFeatureChannels(value objectivec.IObject)
	SerializeWeightsAndBiases() objectivec.IObject
	SetSerializeWeightsAndBiases(value objectivec.IObject)
	DataSource() CNNConvolutionDataSource get /* not a class type */
	SetDataSource(value CNNConvolutionDataSource get /* not a class type */)
	ChannelMultiplier() objectivec.IObject
	SetChannelMultiplier(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionGradient */
	// methods:
	ReloadWeightsAndBiases()
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesFromDataSource()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionGradientClass) Alloc() CNNConvolutionGradient {
	rv := objc.Send[CNNConvolutionGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionGradientClass) New() CNNConvolutionGradient {
	rv := objc.Send[CNNConvolutionGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionGradient) Init() CNNConvolutionGradient {
	rv := objc.Send[CNNConvolutionGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionGradient) Autorelease() CNNConvolutionGradient {
	rv := objc.Send[CNNConvolutionGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionGradient creates a new CNNConvolutionGradient instance.
func NewCNNConvolutionGradient() CNNConvolutionGradient {
	return getCNNConvolutionGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionGradient */
// A gradient convolution kernel.


// A gradient convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient
type CNNConvolutionGradient struct {
	CNNGradientKernel
}

// CNNConvolutionGradientFrom constructs a [CNNConvolutionGradient] from an unsafe.Pointer.
//
// A gradient convolution kernel.
func CNNConvolutionGradientFrom(ptr unsafe.Pointer) CNNConvolutionGradient {
	return CNNConvolutionGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942425-initwithcoder
func NewCNNConvolutionGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNConvolutionGradient {
	instance := getCNNConvolutionGradientClass().Alloc()
	rv := objc.Send[CNNConvolutionGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942414-initwithdevice
func NewCNNConvolutionGradientWithDeviceWeights(device unsafe.Pointer, weights unsafe.Pointer) CNNConvolutionGradient {
	instance := getCNNConvolutionGradientClass().Alloc()
	rv := objc.Send[CNNConvolutionGradient](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionGradientWithDeviceWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2953960-reloadweightsandbiases
func (c_ CNNConvolutionGradient) ReloadWeightsAndBiases() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiases"))
}/* debug [instance_methods/method]: ReloadWeightsAndBiases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2953960-reloadweightsandbiaseswithcomman
func (c_ CNNConvolutionGradient) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}/* debug [instance_methods/method]: ReloadWeightsAndBiasesWithCommandBufferState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2966659-reloadweightsandbiasesfromdataso
func (c_ CNNConvolutionGradient) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}/* debug [instance_methods/method]: ReloadWeightsAndBiasesFromDataSource */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942430-groups
func (c_ CNNConvolutionGradient) Groups() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942430-groups
func (c_ CNNConvolutionGradient) SetGroups(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942432-gradientoption
func (c_ CNNConvolutionGradient) GradientOption() CNNConvolutionGradientOption get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientOption"))
	return rv
}/* debug [instance_properties/getter]: gradientOption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942432-gradientoption
func (c_ CNNConvolutionGradient) SetGradientOption(value CNNConvolutionGradientOption get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientOption:"), value)
}/* debug [instance_properties/setter]: gradientOption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2947880-sourcegradientfeaturechannels
func (c_ CNNConvolutionGradient) SourceGradientFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sourceGradientFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceGradientFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2947880-sourcegradientfeaturechannels
func (c_ CNNConvolutionGradient) SetSourceGradientFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceGradientFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceGradientFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2947882-sourceimagefeaturechannels
func (c_ CNNConvolutionGradient) SourceImageFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sourceImageFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceImageFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2947882-sourceimagefeaturechannels
func (c_ CNNConvolutionGradient) SetSourceImageFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceImageFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceImageFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2951926-serializeweightsandbiases
func (c_ CNNConvolutionGradient) SerializeWeightsAndBiases() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("serializeWeightsAndBiases"))
	return rv
}/* debug [instance_properties/getter]: serializeWeightsAndBiases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2951926-serializeweightsandbiases
func (c_ CNNConvolutionGradient) SetSerializeWeightsAndBiases(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSerializeWeightsAndBiases:"), value)
}/* debug [instance_properties/setter]: serializeWeightsAndBiases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2953959-datasource
func (c_ CNNConvolutionGradient) DataSource() CNNConvolutionDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2953959-datasource
func (c_ CNNConvolutionGradient) SetDataSource(value CNNConvolutionDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2966658-channelmultiplier
func (c_ CNNConvolutionGradient) ChannelMultiplier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("channelMultiplier"))
	return rv
}/* debug [instance_properties/getter]: channelMultiplier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2966658-channelmultiplier
func (c_ CNNConvolutionGradient) SetChannelMultiplier(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChannelMultiplier:"), value)
}/* debug [instance_properties/setter]: channelMultiplier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionGradient */


