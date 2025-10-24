// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionTransposeGradient */


/* debug [class_header]: Header for MPSCNNConvolutionTransposeGradient */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionTransposeGradient */
// An interface definition for the [CNNConvolutionTransposeGradient] class.
type ICNNConvolutionTransposeGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNConvolutionTransposeGradient */
	// properties:
	DataSource() CNNConvolutionDataSource get /* not a class type */
	SetDataSource(value CNNConvolutionDataSource get /* not a class type */)
	GradientOption() CNNConvolutionGradientOption get set /* not a class type */
	SetGradientOption(value CNNConvolutionGradientOption get set /* not a class type */)
	Groups() objectivec.IObject
	SetGroups(value objectivec.IObject)
	SourceGradientFeatureChannels() objectivec.IObject
	SetSourceGradientFeatureChannels(value objectivec.IObject)
	SourceImageFeatureChannels() objectivec.IObject
	SetSourceImageFeatureChannels(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionTransposeGradient */
	// methods:
	ReloadWeightsAndBiasesFromDataSource()
	ReloadWeightsAndBiases()
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionTransposeGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeGradientClass) Alloc() CNNConvolutionTransposeGradient {
	rv := objc.Send[CNNConvolutionTransposeGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionTransposeGradient */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient
type CNNConvolutionTransposeGradient struct {
	CNNGradientKernel
}

// CNNConvolutionTransposeGradientFrom constructs a [CNNConvolutionTransposeGradient] from an unsafe.Pointer.
func CNNConvolutionTransposeGradientFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradient {
	return CNNConvolutionTransposeGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionTransposeGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131783-initwithcoder
func NewCNNConvolutionTransposeGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNConvolutionTransposeGradient {
	instance := getCNNConvolutionTransposeGradientClass().Alloc()
	rv := objc.Send[CNNConvolutionTransposeGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionTransposeGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131784-initwithdevice
func NewCNNConvolutionTransposeGradientWithDeviceWeights(device unsafe.Pointer, weights unsafe.Pointer) CNNConvolutionTransposeGradient {
	instance := getCNNConvolutionTransposeGradientClass().Alloc()
	rv := objc.Send[CNNConvolutionTransposeGradient](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionTransposeGradientWithDeviceWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionTransposeGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionTransposeGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionTransposeGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131785-reloadweightsandbiasesfromdataso
func (c_ CNNConvolutionTransposeGradient) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}/* debug [instance_methods/method]: ReloadWeightsAndBiasesFromDataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131786-reloadweightsandbiases
func (c_ CNNConvolutionTransposeGradient) ReloadWeightsAndBiases() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiases"))
}/* debug [instance_methods/method]: ReloadWeightsAndBiases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131786-reloadweightsandbiaseswithcomman
func (c_ CNNConvolutionTransposeGradient) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}/* debug [instance_methods/method]: ReloadWeightsAndBiasesWithCommandBufferState */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionTransposeGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131780-datasource
func (c_ CNNConvolutionTransposeGradient) DataSource() CNNConvolutionDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131780-datasource
func (c_ CNNConvolutionTransposeGradient) SetDataSource(value CNNConvolutionDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131781-gradientoption
func (c_ CNNConvolutionTransposeGradient) GradientOption() CNNConvolutionGradientOption get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientOption"))
	return rv
}/* debug [instance_properties/getter]: gradientOption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131781-gradientoption
func (c_ CNNConvolutionTransposeGradient) SetGradientOption(value CNNConvolutionGradientOption get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientOption:"), value)
}/* debug [instance_properties/setter]: gradientOption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131782-groups
func (c_ CNNConvolutionTransposeGradient) Groups() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131782-groups
func (c_ CNNConvolutionTransposeGradient) SetGroups(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131787-sourcegradientfeaturechannels
func (c_ CNNConvolutionTransposeGradient) SourceGradientFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sourceGradientFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceGradientFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131787-sourcegradientfeaturechannels
func (c_ CNNConvolutionTransposeGradient) SetSourceGradientFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceGradientFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceGradientFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131788-sourceimagefeaturechannels
func (c_ CNNConvolutionTransposeGradient) SourceImageFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sourceImageFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: sourceImageFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131788-sourceimagefeaturechannels
func (c_ CNNConvolutionTransposeGradient) SetSourceImageFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceImageFeatureChannels:"), value)
}/* debug [instance_properties/setter]: sourceImageFeatureChannels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionTransposeGradient */


