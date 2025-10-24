// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphExecutable */


/* debug [class_header]: Header for MPSGraphExecutable */
// The class instance for the [GraphExecutable] class.
var (
	GraphExecutableClass     _GraphExecutableClass
	GraphExecutableClassOnce sync.Once
)

func getGraphExecutableClass() _GraphExecutableClass {
	GraphExecutableClassOnce.Do(func() {
		GraphExecutableClass = _GraphExecutableClass{objc.GetClass("MPSGraphExecutable")}
	})
	return GraphExecutableClass
}

type _GraphExecutableClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphExecutable */
// An interface definition for the [GraphExecutable] class.
type IGraphExecutable interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphExecutable */
	// properties:
	FeedTensors() []GraphTensor
	Options() GraphOptions
	SetOptions(value GraphOptions)
	TargetTensors() []GraphTensor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphExecutable */
	// methods:
	EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor(commandBuffer objc.IObject /* cross-framework: MPSCommandBuffer */, inputsArray []GraphTensorData, resultsArray []GraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []GraphTensorData
	GetOutputTypesWithDeviceInputTypesCompilationDescriptor(device IMPSGraphDevice, inputTypes []GraphType, compilationDescriptor IMPSGraphCompilationDescriptor) []GraphShapedType
	RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue unsafe.Pointer, inputsArray []GraphTensorData, resultsArray []GraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []GraphTensorData
	RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue unsafe.Pointer, inputsArray []GraphTensorData, resultsArray []GraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []GraphTensorData
	SerializeToMPSGraphPackageAtURLDescriptor(url objc.IObject /* cross-framework: NSURL */, descriptor IMPSGraphExecutableSerializationDescriptor)
	SpecializeWithDeviceInputTypesCompilationDescriptor(device IMPSGraphDevice, inputTypes []GraphType, compilationDescriptor IMPSGraphCompilationDescriptor)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphExecutable */
// Alloc allocates a new instance without initialization.
func (gc _GraphExecutableClass) Alloc() GraphExecutable {
	rv := objc.Send[GraphExecutable](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphExecutableClass) New() GraphExecutable {
	rv := objc.Send[GraphExecutable](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphExecutable) Init() GraphExecutable {
	rv := objc.Send[GraphExecutable](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphExecutable) Autorelease() GraphExecutable {
	rv := objc.Send[GraphExecutable](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphExecutable creates a new GraphExecutable instance.
func NewGraphExecutable() GraphExecutable {
	return getGraphExecutableClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphExecutable */
// The compiled representation of a compute graph executable.
//
// An is a compiled graph for specific feeds for specific target tensors and target operations.


// The compiled representation of a compute graph executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable
type GraphExecutable struct {
	GraphObject
}

// GraphExecutableFrom constructs a [GraphExecutable] from an unsafe.Pointer.
//
// The compiled representation of a compute graph executable.
func GraphExecutableFrom(ptr unsafe.Pointer) GraphExecutable {
	return GraphExecutable{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphExecutable */

// Initialize the executable with the Core ML model package at the provided URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/init(coreMLPackageAtURL:descriptor:)
func NewGraphExecutableWithCoreMLPackageAtURLCompilationDescriptor(coreMLPackageURL objc.IObject /* cross-framework: NSURL */, compilationDescriptor IMPSGraphCompilationDescriptor) GraphExecutable {
	instance := getGraphExecutableClass().Alloc()
	rv := objc.Send[GraphExecutable](instance.ID, objc.Sel("initWithCoreMLPackageAtURL:compilationDescriptor:"), coreMLPackageURL, compilationDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphExecutableWithCoreMLPackageAtURLCompilationDescriptor */


// Initialize the executable with the Metal Performance Shaders Graph package at the provided URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/init(package:descriptor:)
func NewGraphExecutableWithMPSGraphPackageAtURLCompilationDescriptor(mpsgraphPackageURL objc.IObject /* cross-framework: NSURL */, compilationDescriptor IMPSGraphCompilationDescriptor) GraphExecutable {
	instance := getGraphExecutableClass().Alloc()
	rv := objc.Send[GraphExecutable](instance.ID, objc.Sel("initWithMPSGraphPackageAtURL:compilationDescriptor:"), mpsgraphPackageURL, compilationDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphExecutableWithMPSGraphPackageAtURLCompilationDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphExecutable */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphExecutable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphExecutable */

// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed. This call is asynchronous and will return immediately after finishing encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/encode(to:inputs:results:executionDescriptor:)
func (g_ GraphExecutable) EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor(commandBuffer objc.IObject /* cross-framework: MPSCommandBuffer */, inputsArray []GraphTensorData, resultsArray []GraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []GraphTensorData {
	rv := objc.Send[[]GraphTensorData](g_.ID, objc.Sel("encodeToCommandBuffer:inputsArray:resultsArray:executionDescriptor:"), commandBuffer, inputsArray, resultsArray, executionDescriptor)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor */


// Get output shapes for a specialized executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/getOutputTypes(with:inputTypes:compilationDescriptor:)
func (g_ GraphExecutable) GetOutputTypesWithDeviceInputTypesCompilationDescriptor(device IMPSGraphDevice, inputTypes []GraphType, compilationDescriptor IMPSGraphCompilationDescriptor) []GraphShapedType {
	rv := objc.Send[[]GraphShapedType](g_.ID, objc.Sel("getOutputTypesWithDevice:inputTypes:compilationDescriptor:"), device, inputTypes, compilationDescriptor)
	return rv
}/* debug [instance_methods/method]: GetOutputTypesWithDeviceInputTypesCompilationDescriptor */


// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/run(with:inputs:results:executionDescriptor:)
func (g_ GraphExecutable) RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue unsafe.Pointer, inputsArray []GraphTensorData, resultsArray []GraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []GraphTensorData {
	rv := objc.Send[[]GraphTensorData](g_.ID, objc.Sel("runWithMTLCommandQueue:inputsArray:resultsArray:executionDescriptor:"), commandQueue, inputsArray, resultsArray, executionDescriptor)
	return rv
}/* debug [instance_methods/method]: RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor */


// Runs the graph for the given feeds and returns the target tensor values, ensuring all target operations also executed. This call is asynchronous and will return immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/runAsync(with:inputs:results:executionDescriptor:)
func (g_ GraphExecutable) RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue unsafe.Pointer, inputsArray []GraphTensorData, resultsArray []GraphTensorData, executionDescriptor IMPSGraphExecutableExecutionDescriptor) []GraphTensorData {
	rv := objc.Send[[]GraphTensorData](g_.ID, objc.Sel("runAsyncWithMTLCommandQueue:inputsArray:resultsArray:executionDescriptor:"), commandQueue, inputsArray, resultsArray, executionDescriptor)
	return rv
}/* debug [instance_methods/method]: RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor */


// Serialize the MPSGraph executable at the provided url.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/serialize(package:descriptor:)
func (g_ GraphExecutable) SerializeToMPSGraphPackageAtURLDescriptor(url objc.IObject /* cross-framework: NSURL */, descriptor IMPSGraphExecutableSerializationDescriptor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("serializeToMPSGraphPackageAtURL:descriptor:"), url, descriptor)
}/* debug [instance_methods/method]: SerializeToMPSGraphPackageAtURLDescriptor */


// Specialize the executable and optimize it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/specialize(with:inputTypes:compilationDescriptor:)
func (g_ GraphExecutable) SpecializeWithDeviceInputTypesCompilationDescriptor(device IMPSGraphDevice, inputTypes []GraphType, compilationDescriptor IMPSGraphCompilationDescriptor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("specializeWithDevice:inputTypes:compilationDescriptor:"), device, inputTypes, compilationDescriptor)
}/* debug [instance_methods/method]: SpecializeWithDeviceInputTypesCompilationDescriptor */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphExecutable */

// Tensors fed to the graph, can be used to order the inputs when executable is created with a graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/feedTensors
func (g_ GraphExecutable) FeedTensors() []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("feedTensors"))
	return rv
}/* debug [instance_properties/getter]: feedTensors */


// Options for the graph executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/options
func (g_ GraphExecutable) Options() GraphOptions {
	rv := objc.Send[GraphOptions](g_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// Options for the graph executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/options
func (g_ GraphExecutable) SetOptions(value GraphOptions) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */


// Tensors targeted by the graph, can be used to order the outputs when executable was created with a graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/targetTensors
func (g_ GraphExecutable) TargetTensors() []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("targetTensors"))
	return rv
}/* debug [instance_properties/getter]: targetTensors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphExecutable */


