// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ComputePipelineReflection] class.
var (
	ComputePipelineReflectionClass     _ComputePipelineReflectionClass
	ComputePipelineReflectionClassOnce sync.Once
)

func getComputePipelineReflectionClass() _ComputePipelineReflectionClass {
	ComputePipelineReflectionClassOnce.Do(func() {
		ComputePipelineReflectionClass = _ComputePipelineReflectionClass{objc.GetClass("MTLComputePipelineReflection")}
	})
	return ComputePipelineReflectionClass
}

type _ComputePipelineReflectionClass struct {
	class objc.Class
}





// An interface definition for the [ComputePipelineReflection] class.
type IComputePipelineReflection interface {
	objectivec.IObject
	

	// properties:
	Arguments() []Argument
	Bindings() []objc.ID


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _ComputePipelineReflectionClass) Alloc() ComputePipelineReflection {
	rv := objc.Send[ComputePipelineReflection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComputePipelineReflectionClass) New() ComputePipelineReflection {
	rv := objc.Send[ComputePipelineReflection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePipelineReflection) Init() ComputePipelineReflection {
	rv := objc.Send[ComputePipelineReflection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePipelineReflection) Autorelease() ComputePipelineReflection {
	rv := objc.Send[ComputePipelineReflection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePipelineReflection creates a new ComputePipelineReflection instance.
func NewComputePipelineReflection() ComputePipelineReflection {
	return getComputePipelineReflectionClass().New()
}





// Information about the arguments of a compute function.
//
// An object provides access to the arguments of the compute function used in an object. An object can be created along with an object. Don’t create an object directly. Instead, call either the or method of to create both an object and an object. objects can use a significant amount of memory; release any strong references to them after you finish creating pipeline objects.


// Information about the arguments of a compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection
type ComputePipelineReflection struct {
	objectivec.Object
}

// ComputePipelineReflectionFrom constructs a [ComputePipelineReflection] from an unsafe.Pointer.
//
// Information about the arguments of a compute function.
func ComputePipelineReflectionFrom(ptr unsafe.Pointer) ComputePipelineReflection {
	return ComputePipelineReflection{objectivec.Object{objc.ID(ptr)}}
}

























// An array of instances that describe the arguments of a compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection/arguments
func (c_ ComputePipelineReflection) Arguments() []Argument {
	rv := objc.Send[[]Argument](c_.ID, objc.Sel("arguments"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection/bindings
func (c_ ComputePipelineReflection) Bindings() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("bindings"))
	return rv
}








