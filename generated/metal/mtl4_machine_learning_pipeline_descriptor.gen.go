// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTL4MachineLearningPipelineDescriptor */


/* debug [class_header]: Header for MTL4MachineLearningPipelineDescriptor */
// The class instance for the [MTL4MachineLearningPipelineDescriptor] class.
var (
	MTL4MachineLearningPipelineDescriptorClass     _MTL4MachineLearningPipelineDescriptorClass
	MTL4MachineLearningPipelineDescriptorClassOnce sync.Once
)

func getMTL4MachineLearningPipelineDescriptorClass() _MTL4MachineLearningPipelineDescriptorClass {
	MTL4MachineLearningPipelineDescriptorClassOnce.Do(func() {
		MTL4MachineLearningPipelineDescriptorClass = _MTL4MachineLearningPipelineDescriptorClass{objc.GetClass("MTL4MachineLearningPipelineDescriptor")}
	})
	return MTL4MachineLearningPipelineDescriptorClass
}

type _MTL4MachineLearningPipelineDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4MachineLearningPipelineDescriptor */
// An interface definition for the [MTL4MachineLearningPipelineDescriptor] class.
type IMTL4MachineLearningPipelineDescriptor interface {
	IMTL4PipelineDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4MachineLearningPipelineDescriptor */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	MachineLearningFunctionDescriptor() IMTL4FunctionDescriptor
	SetMachineLearningFunctionDescriptor(value IMTL4FunctionDescriptor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4MachineLearningPipelineDescriptor */
	// methods:
	InputDimensionsAtBufferIndex(bufferIndex int) ITensorExtents
	Reset()
	SetInputDimensionsAtBufferIndex(dimensions IMTLTensorExtents, bufferIndex int)
	SetInputDimensionsWithRange(dimensions []TensorExtents, range_ corefoundation.Range)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4MachineLearningPipelineDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4MachineLearningPipelineDescriptorClass) Alloc() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4MachineLearningPipelineDescriptorClass) New() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4MachineLearningPipelineDescriptor) Init() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4MachineLearningPipelineDescriptor) Autorelease() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4MachineLearningPipelineDescriptor creates a new MTL4MachineLearningPipelineDescriptor instance.
func NewMTL4MachineLearningPipelineDescriptor() MTL4MachineLearningPipelineDescriptor {
	return getMTL4MachineLearningPipelineDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4MachineLearningPipelineDescriptor */
// Description for a machine learning pipeline state.


// Description for a machine learning pipeline state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor
type MTL4MachineLearningPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4MachineLearningPipelineDescriptorFrom constructs a [MTL4MachineLearningPipelineDescriptor] from an unsafe.Pointer.
//
// Description for a machine learning pipeline state.
func MTL4MachineLearningPipelineDescriptorFrom(ptr unsafe.Pointer) MTL4MachineLearningPipelineDescriptor {
	return MTL4MachineLearningPipelineDescriptor{
		MTL4PipelineDescriptor: MTL4PipelineDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4MachineLearningPipelineDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4MachineLearningPipelineDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4MachineLearningPipelineDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4MachineLearningPipelineDescriptor */

// Obtains the dimensions of the input tensor at if set, otherwise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/inputDimensions(bufferIndex:)
func (m_ MTL4MachineLearningPipelineDescriptor) InputDimensionsAtBufferIndex(bufferIndex int) ITensorExtents {
	rv := objc.Send[TensorExtents](m_.ID, objc.Sel("inputDimensionsAtBufferIndex:"), bufferIndex)
	return rv
}/* debug [instance_methods/method]: InputDimensionsAtBufferIndex */


// Resets the descriptor to its default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/reset()
func (m_ MTL4MachineLearningPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// Sets the dimension of an input tensor at a buffer index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/setInputDimensions(_:bufferIndex:)-34gir
func (m_ MTL4MachineLearningPipelineDescriptor) SetInputDimensionsAtBufferIndex(dimensions IMTLTensorExtents, bufferIndex int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputDimensions:atBufferIndex:"), dimensions, bufferIndex)
}/* debug [instance_methods/method]: SetInputDimensionsAtBufferIndex */


// Sets the dimensions of multiple input tensors on a range of buffer bindings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/setInputDimensions:withRange:
func (m_ MTL4MachineLearningPipelineDescriptor) SetInputDimensionsWithRange(dimensions []TensorExtents, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputDimensions:withRange:"), dimensions, range_)
}/* debug [instance_methods/method]: SetInputDimensionsWithRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4MachineLearningPipelineDescriptor */

// Assigns an optional string that helps identify pipeline states you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/label
func (m_ MTL4MachineLearningPipelineDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Assigns an optional string that helps identify pipeline states you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/label
func (m_ MTL4MachineLearningPipelineDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Assigns the function that the machine learning pipeline you create from this descriptor executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/machineLearningFunctionDescriptor
func (m_ MTL4MachineLearningPipelineDescriptor) MachineLearningFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("machineLearningFunctionDescriptor"))
	return rv
}/* debug [instance_properties/getter]: machineLearningFunctionDescriptor */


// Assigns the function that the machine learning pipeline you create from this descriptor executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/machineLearningFunctionDescriptor
func (m_ MTL4MachineLearningPipelineDescriptor) SetMachineLearningFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMachineLearningFunctionDescriptor:"), value)
}/* debug [instance_properties/setter]: machineLearningFunctionDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4MachineLearningPipelineDescriptor */



