// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphStencilOpDescriptor */


/* debug [class_header]: Header for MPSGraphStencilOpDescriptor */
// The class instance for the [GraphStencilOpDescriptor] class.
var (
	GraphStencilOpDescriptorClass     _GraphStencilOpDescriptorClass
	GraphStencilOpDescriptorClassOnce sync.Once
)

func getGraphStencilOpDescriptorClass() _GraphStencilOpDescriptorClass {
	GraphStencilOpDescriptorClassOnce.Do(func() {
		GraphStencilOpDescriptorClass = _GraphStencilOpDescriptorClass{objc.GetClass("MPSGraphStencilOpDescriptor")}
	})
	return GraphStencilOpDescriptorClass
}

type _GraphStencilOpDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphStencilOpDescriptor */
// An interface definition for the [GraphStencilOpDescriptor] class.
type IGraphStencilOpDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphStencilOpDescriptor */
	// properties:
	BoundaryMode() GraphPaddingMode
	SetBoundaryMode(value GraphPaddingMode)
	DilationRates() Shape /* not a class type */
	SetDilationRates(value Shape /* not a class type */)
	ExplicitPadding() Shape /* not a class type */
	SetExplicitPadding(value Shape /* not a class type */)
	Offsets() Shape /* not a class type */
	SetOffsets(value Shape /* not a class type */)
	PaddingConstant() float32
	SetPaddingConstant(value float32)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	ReductionMode() GraphReductionMode
	SetReductionMode(value GraphReductionMode)
	Strides() Shape /* not a class type */
	SetStrides(value Shape /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphStencilOpDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphStencilOpDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphStencilOpDescriptorClass) Alloc() GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphStencilOpDescriptorClass) New() GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphStencilOpDescriptor) Init() GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphStencilOpDescriptor) Autorelease() GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphStencilOpDescriptor creates a new GraphStencilOpDescriptor instance.
func NewGraphStencilOpDescriptor() GraphStencilOpDescriptor {
	return getGraphStencilOpDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphStencilOpDescriptor */
// The class that defines the parameters for a stencil operation.
//
// Use this descriptor with the following method:


// The class that defines the parameters for a stencil operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor
type GraphStencilOpDescriptor struct {
	GraphObject
}

// GraphStencilOpDescriptorFrom constructs a [GraphStencilOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a stencil operation.
func GraphStencilOpDescriptorFrom(ptr unsafe.Pointer) GraphStencilOpDescriptor {
	return GraphStencilOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphStencilOpDescriptor */

// Creates a stencil operation descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(explicitPadding:)
func NewGraphStencilOpDescriptorWithExplicitPadding(explicitPadding Shape /* not a class type */) GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](objc.ID(getGraphStencilOpDescriptorClass().class), objc.Sel("descriptorWithExplicitPadding:"), explicitPadding)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphStencilOpDescriptorWithExplicitPadding */


// Creates a stencil operation descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(offsets:explicitPadding:)
func NewGraphStencilOpDescriptorWithOffsetsExplicitPadding(offsets Shape /* not a class type */, explicitPadding Shape /* not a class type */) GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](objc.ID(getGraphStencilOpDescriptorClass().class), objc.Sel("descriptorWithOffsets:explicitPadding:"), offsets, explicitPadding)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphStencilOpDescriptorWithOffsetsExplicitPadding */


// Creates a stencil operation descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(paddingStyle:)
func NewGraphStencilOpDescriptorWithPaddingStyle(paddingStyle GraphPaddingStyle) GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](objc.ID(getGraphStencilOpDescriptorClass().class), objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphStencilOpDescriptorWithPaddingStyle */


// Creates a stencil operation descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(reductionMode:offsets:strides:dilationRates:explicitPadding:boundaryMode:paddingStyle:paddingConstant:)
func NewGraphStencilOpDescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant(reductionMode GraphReductionMode, offsets Shape /* not a class type */, strides Shape /* not a class type */, dilationRates Shape /* not a class type */, explicitPadding Shape /* not a class type */, boundaryMode GraphPaddingMode, paddingStyle GraphPaddingStyle, paddingConstant float32) GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](objc.ID(getGraphStencilOpDescriptorClass().class), objc.Sel("descriptorWithReductionMode:offsets:strides:dilationRates:explicitPadding:boundaryMode:paddingStyle:paddingConstant:"), reductionMode, offsets, strides, dilationRates, explicitPadding, boundaryMode, paddingStyle, paddingConstant)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphStencilOpDescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphStencilOpDescriptor */

// Creates a stencil operation descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(explicitPadding:)
func (gc _GraphStencilOpDescriptorClass) DescriptorWithExplicitPadding(explicitPadding Shape /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithExplicitPadding:"), explicitPadding)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithExplicitPadding) */


// Creates a stencil operation descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(offsets:explicitPadding:)
func (gc _GraphStencilOpDescriptorClass) DescriptorWithOffsetsExplicitPadding(offsets Shape /* not a class type */, explicitPadding Shape /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithOffsets:explicitPadding:"), offsets, explicitPadding)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithOffsetsExplicitPadding) */


// Creates a stencil operation descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(paddingStyle:)
func (gc _GraphStencilOpDescriptorClass) DescriptorWithPaddingStyle(paddingStyle GraphPaddingStyle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithPaddingStyle) */


// Creates a stencil operation descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(reductionMode:offsets:strides:dilationRates:explicitPadding:boundaryMode:paddingStyle:paddingConstant:)
func (gc _GraphStencilOpDescriptorClass) DescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant(reductionMode GraphReductionMode, offsets Shape /* not a class type */, strides Shape /* not a class type */, dilationRates Shape /* not a class type */, explicitPadding Shape /* not a class type */, boundaryMode GraphPaddingMode, paddingStyle GraphPaddingStyle, paddingConstant float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithReductionMode:offsets:strides:dilationRates:explicitPadding:boundaryMode:paddingStyle:paddingConstant:"), reductionMode, offsets, strides, dilationRates, explicitPadding, boundaryMode, paddingStyle, paddingConstant)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphStencilOpDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphStencilOpDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphStencilOpDescriptor */

// The property that determines which values to use for padding the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/boundaryMode
func (g_ GraphStencilOpDescriptor) BoundaryMode() GraphPaddingMode {
	rv := objc.Send[GraphPaddingMode](g_.ID, objc.Sel("boundaryMode"))
	return rv
}/* debug [instance_properties/getter]: boundaryMode */


// The property that determines which values to use for padding the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/boundaryMode
func (g_ GraphStencilOpDescriptor) SetBoundaryMode(value GraphPaddingMode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBoundaryMode:"), value)
}/* debug [instance_properties/setter]: boundaryMode */


// The property that defines dilation rates for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/dilationRates
func (g_ GraphStencilOpDescriptor) DilationRates() Shape /* not a class type */ {
	rv := objc.Send[Shape](g_.ID, objc.Sel("dilationRates"))
	return rv
}/* debug [instance_properties/getter]: dilationRates */


// The property that defines dilation rates for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/dilationRates
func (g_ GraphStencilOpDescriptor) SetDilationRates(value Shape /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRates:"), value)
}/* debug [instance_properties/setter]: dilationRates */


// The property that defines padding values for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/explicitPadding
func (g_ GraphStencilOpDescriptor) ExplicitPadding() Shape /* not a class type */ {
	rv := objc.Send[Shape](g_.ID, objc.Sel("explicitPadding"))
	return rv
}/* debug [instance_properties/getter]: explicitPadding */


// The property that defines padding values for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/explicitPadding
func (g_ GraphStencilOpDescriptor) SetExplicitPadding(value Shape /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExplicitPadding:"), value)
}/* debug [instance_properties/setter]: explicitPadding */


// An array of length four that determines from which offset to start reading the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/offsets
func (g_ GraphStencilOpDescriptor) Offsets() Shape /* not a class type */ {
	rv := objc.Send[Shape](g_.ID, objc.Sel("offsets"))
	return rv
}/* debug [instance_properties/getter]: offsets */


// An array of length four that determines from which offset to start reading the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/offsets
func (g_ GraphStencilOpDescriptor) SetOffsets(value Shape /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOffsets:"), value)
}/* debug [instance_properties/setter]: offsets */


// The padding value for .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/paddingConstant
func (g_ GraphStencilOpDescriptor) PaddingConstant() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("paddingConstant"))
	return rv
}/* debug [instance_properties/getter]: paddingConstant */


// The padding value for .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/paddingConstant
func (g_ GraphStencilOpDescriptor) SetPaddingConstant(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingConstant:"), value)
}/* debug [instance_properties/setter]: paddingConstant */


// The property that defines what kind of padding to apply to the stencil operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/paddingStyle
func (g_ GraphStencilOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}/* debug [instance_properties/getter]: paddingStyle */


// The property that defines what kind of padding to apply to the stencil operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/paddingStyle
func (g_ GraphStencilOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}/* debug [instance_properties/setter]: paddingStyle */


// The reduction mode to use within the stencil window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/reductionMode
func (g_ GraphStencilOpDescriptor) ReductionMode() GraphReductionMode {
	rv := objc.Send[GraphReductionMode](g_.ID, objc.Sel("reductionMode"))
	return rv
}/* debug [instance_properties/getter]: reductionMode */


// The reduction mode to use within the stencil window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/reductionMode
func (g_ GraphStencilOpDescriptor) SetReductionMode(value GraphReductionMode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReductionMode:"), value)
}/* debug [instance_properties/setter]: reductionMode */


// The property that defines strides for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/strides
func (g_ GraphStencilOpDescriptor) Strides() Shape /* not a class type */ {
	rv := objc.Send[Shape](g_.ID, objc.Sel("strides"))
	return rv
}/* debug [instance_properties/getter]: strides */


// The property that defines strides for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/strides
func (g_ GraphStencilOpDescriptor) SetStrides(value Shape /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrides:"), value)
}/* debug [instance_properties/setter]: strides */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphStencilOpDescriptor */


