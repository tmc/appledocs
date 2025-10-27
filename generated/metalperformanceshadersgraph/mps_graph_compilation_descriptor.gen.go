// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [GraphCompilationDescriptor] class.
var (
	GraphCompilationDescriptorClass     _GraphCompilationDescriptorClass
	GraphCompilationDescriptorClassOnce sync.Once
)

func getGraphCompilationDescriptorClass() _GraphCompilationDescriptorClass {
	GraphCompilationDescriptorClassOnce.Do(func() {
		GraphCompilationDescriptorClass = _GraphCompilationDescriptorClass{objc.GetClass("MPSGraphCompilationDescriptor")}
	})
	return GraphCompilationDescriptorClass
}

type _GraphCompilationDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [GraphCompilationDescriptor] class.
type IGraphCompilationDescriptor interface {
	IGraphObject
	

	// properties:
	Callables() GraphCallableMap /* not a class type */
	SetCallables(value GraphCallableMap /* not a class type */)
	CompilationCompletionHandler() GraphCompilationCompletionHandler /* not a class type */
	SetCompilationCompletionHandler(value GraphCompilationCompletionHandler /* not a class type */)
	DispatchQueue() unsafe.Pointer
	SetDispatchQueue(value unsafe.Pointer)
	OptimizationLevel() GraphOptimization
	SetOptimizationLevel(value GraphOptimization)
	OptimizationProfile() GraphOptimizationProfile
	SetOptimizationProfile(value GraphOptimizationProfile)
	ReducedPrecisionFastMath() GraphReducedPrecisionFastMath
	SetReducedPrecisionFastMath(value GraphReducedPrecisionFastMath)
	WaitForCompilationCompletion() bool
	SetWaitForCompilationCompletion(value bool)


	

	// methods:
	DisableTypeInference()


}





// Alloc allocates a new instance without initialization.
func (gc _GraphCompilationDescriptorClass) Alloc() GraphCompilationDescriptor {
	rv := objc.Send[GraphCompilationDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphCompilationDescriptorClass) New() GraphCompilationDescriptor {
	rv := objc.Send[GraphCompilationDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphCompilationDescriptor) Init() GraphCompilationDescriptor {
	rv := objc.Send[GraphCompilationDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphCompilationDescriptor) Autorelease() GraphCompilationDescriptor {
	rv := objc.Send[GraphCompilationDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphCompilationDescriptor creates a new GraphCompilationDescriptor instance.
func NewGraphCompilationDescriptor() GraphCompilationDescriptor {
	return getGraphCompilationDescriptorClass().New()
}





// A class that consists of all the levers for compiling graphs.


// A class that consists of all the levers for compiling graphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor
type GraphCompilationDescriptor struct {
	GraphObject
}

// GraphCompilationDescriptorFrom constructs a [GraphCompilationDescriptor] from an unsafe.Pointer.
//
// A class that consists of all the levers for compiling graphs.
func GraphCompilationDescriptorFrom(ptr unsafe.Pointer) GraphCompilationDescriptor {
	return GraphCompilationDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}




















// Turns off type inference and relies on type inference during runtime.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/disableTypeInference()
func (g_ GraphCompilationDescriptor) DisableTypeInference() {
	objc.Send[objc.ID](g_.ID, objc.Sel("disableTypeInference"))
}







// The dictionary used during runtime to lookup the which correspond to the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/callables
func (g_ GraphCompilationDescriptor) Callables() GraphCallableMap /* not a class type */ {
	rv := objc.Send[GraphCallableMap](g_.ID, objc.Sel("callables"))
	return rv
}


// The dictionary used during runtime to lookup the which correspond to the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/callables
func (g_ GraphCompilationDescriptor) SetCallables(value GraphCallableMap /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCallables:"), value)
}


// The handler that the graph calls when the compilation completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/compilationCompletionHandler
func (g_ GraphCompilationDescriptor) CompilationCompletionHandler() GraphCompilationCompletionHandler /* not a class type */ {
	rv := objc.Send[GraphCompilationCompletionHandler](g_.ID, objc.Sel("compilationCompletionHandler"))
	return rv
}


// The handler that the graph calls when the compilation completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/compilationCompletionHandler
func (g_ GraphCompilationDescriptor) SetCompilationCompletionHandler(value GraphCompilationCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompilationCompletionHandler:"), value)
}


// The dispatch queue used for the compilation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/dispatchQueue
func (g_ GraphCompilationDescriptor) DispatchQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dispatchQueue"))
	return rv
}


// The dispatch queue used for the compilation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/dispatchQueue
func (g_ GraphCompilationDescriptor) SetDispatchQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDispatchQueue:"), value)
}


// The optimization level for the graph execution, default is MPSGraphOptimizationLevel1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/optimizationLevel
func (g_ GraphCompilationDescriptor) OptimizationLevel() GraphOptimization {
	rv := objc.Send[GraphOptimization](g_.ID, objc.Sel("optimizationLevel"))
	return rv
}


// The optimization level for the graph execution, default is MPSGraphOptimizationLevel1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/optimizationLevel
func (g_ GraphCompilationDescriptor) SetOptimizationLevel(value GraphOptimization) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOptimizationLevel:"), value)
}


// The optimization profile for the graph optimization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/optimizationProfile
func (g_ GraphCompilationDescriptor) OptimizationProfile() GraphOptimizationProfile {
	rv := objc.Send[GraphOptimizationProfile](g_.ID, objc.Sel("optimizationProfile"))
	return rv
}


// The optimization profile for the graph optimization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/optimizationProfile
func (g_ GraphCompilationDescriptor) SetOptimizationProfile(value GraphOptimizationProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOptimizationProfile:"), value)
}


// Across the executable allow reduced precision fast math optimizations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/reducedPrecisionFastMath
func (g_ GraphCompilationDescriptor) ReducedPrecisionFastMath() GraphReducedPrecisionFastMath {
	rv := objc.Send[GraphReducedPrecisionFastMath](g_.ID, objc.Sel("reducedPrecisionFastMath"))
	return rv
}


// Across the executable allow reduced precision fast math optimizations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/reducedPrecisionFastMath
func (g_ GraphCompilationDescriptor) SetReducedPrecisionFastMath(value GraphReducedPrecisionFastMath) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReducedPrecisionFastMath:"), value)
}


// Flag that makes the compile or specialize call blocking till the entire compilation is complete, defaults to NO.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/waitForCompilationCompletion
func (g_ GraphCompilationDescriptor) WaitForCompilationCompletion() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("waitForCompilationCompletion"))
	return rv
}


// Flag that makes the compile or specialize call blocking till the entire compilation is complete, defaults to NO.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/waitForCompilationCompletion
func (g_ GraphCompilationDescriptor) SetWaitForCompilationCompletion(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWaitForCompilationCompletion:"), value)
}








