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
}

// A class that consists of all the levers for compiling graphs.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphCompilationDescriptorClass) Alloc() GraphCompilationDescriptor {
	rv := objc.Send[GraphCompilationDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The optimization level for the graph execution, default is MPSGraphOptimizationLevel1.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/optimizationlevel
func (g_ GraphCompilationDescriptor) OptimizationLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("optimizationLevel"))
	return rv
}


// SetOptimizationLevel sets the value of the optimizationLevel property.
// The optimization level for the graph execution, default is MPSGraphOptimizationLevel1.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/optimizationlevel
func (g_ GraphCompilationDescriptor) SetOptimizationLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOptimizationLevel:"), value)
}

// The handler that the graph calls when the compilation completes.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/compilationcompletionhandler
func (g_ GraphCompilationDescriptor) CompilationCompletionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("compilationCompletionHandler"))
	return rv
}


// SetCompilationCompletionHandler sets the value of the compilationCompletionHandler property.
// The handler that the graph calls when the compilation completes.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/compilationcompletionhandler
func (g_ GraphCompilationDescriptor) SetCompilationCompletionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompilationCompletionHandler:"), value)
}

// Across the executable allow reduced precision fast math optimizations.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/reducedprecisionfastmath
func (g_ GraphCompilationDescriptor) ReducedPrecisionFastMath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("reducedPrecisionFastMath"))
	return rv
}


// SetReducedPrecisionFastMath sets the value of the reducedPrecisionFastMath property.
// Across the executable allow reduced precision fast math optimizations.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/reducedprecisionfastmath
func (g_ GraphCompilationDescriptor) SetReducedPrecisionFastMath(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReducedPrecisionFastMath:"), value)
}

// Flag that makes the compile or specialize call blocking till the entire compilation is complete, defaults to NO.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/waitforcompilationcompletion
func (g_ GraphCompilationDescriptor) WaitForCompilationCompletion() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("waitForCompilationCompletion"))
	return rv
}


// SetWaitForCompilationCompletion sets the value of the waitForCompilationCompletion property.
// Flag that makes the compile or specialize call blocking till the entire compilation is complete, defaults to NO.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/waitforcompilationcompletion
func (g_ GraphCompilationDescriptor) SetWaitForCompilationCompletion(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWaitForCompilationCompletion:"), value)
}

// The dictionary used during runtime to lookup the
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/callables
func (g_ GraphCompilationDescriptor) Callables() string {
	rv := objc.Send[string](g_.ID, objc.Sel("callables"))
	return rv
}


// SetCallables sets the value of the callables property.
// The dictionary used during runtime to lookup the

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/callables
func (g_ GraphCompilationDescriptor) SetCallables(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCallables:"), objc.String(value))
}

// The dispatch queue used for the compilation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/dispatchQueue
func (g_ GraphCompilationDescriptor) DispatchQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dispatchQueue"))
	return rv
}


// SetDispatchQueue sets the value of the dispatchQueue property.
// The dispatch queue used for the compilation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/dispatchQueue
func (g_ GraphCompilationDescriptor) SetDispatchQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDispatchQueue:"), value)
}

// The optimization profile for the graph optimization.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/optimizationProfile
func (g_ GraphCompilationDescriptor) OptimizationProfile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("optimizationProfile"))
	return rv
}


// SetOptimizationProfile sets the value of the optimizationProfile property.
// The optimization profile for the graph optimization.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/optimizationProfile
func (g_ GraphCompilationDescriptor) SetOptimizationProfile(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOptimizationProfile:"), value)
}



