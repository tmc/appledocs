// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4PipelineStageDynamicLinkingDescriptor] class.
var (
	MTL4PipelineStageDynamicLinkingDescriptorClass     _MTL4PipelineStageDynamicLinkingDescriptorClass
	MTL4PipelineStageDynamicLinkingDescriptorClassOnce sync.Once
)

func getMTL4PipelineStageDynamicLinkingDescriptorClass() _MTL4PipelineStageDynamicLinkingDescriptorClass {
	MTL4PipelineStageDynamicLinkingDescriptorClassOnce.Do(func() {
		MTL4PipelineStageDynamicLinkingDescriptorClass = _MTL4PipelineStageDynamicLinkingDescriptorClass{objc.GetClass("MTL4PipelineStageDynamicLinkingDescriptor")}
	})
	return MTL4PipelineStageDynamicLinkingDescriptorClass
}

type _MTL4PipelineStageDynamicLinkingDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4PipelineStageDynamicLinkingDescriptor] class.
type IMTL4PipelineStageDynamicLinkingDescriptor interface {
	objectivec.IObject
	

	// properties:
	BinaryLinkedFunctions() []objc.ID
	SetBinaryLinkedFunctions(value []objc.ID)
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)
	PreloadedLibraries() []objc.ID
	SetPreloadedLibraries(value []objc.ID)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4PipelineStageDynamicLinkingDescriptorClass) Alloc() MTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4PipelineStageDynamicLinkingDescriptorClass) New() MTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) Init() MTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) Autorelease() MTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineStageDynamicLinkingDescriptor creates a new MTL4PipelineStageDynamicLinkingDescriptor instance.
func NewMTL4PipelineStageDynamicLinkingDescriptor() MTL4PipelineStageDynamicLinkingDescriptor {
	return getMTL4PipelineStageDynamicLinkingDescriptorClass().New()
}





// Groups together properties to drive the dynamic linking process of a pipeline stage.


// Groups together properties to drive the dynamic linking process of a pipeline stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor
type MTL4PipelineStageDynamicLinkingDescriptor struct {
	objectivec.Object
}

// MTL4PipelineStageDynamicLinkingDescriptorFrom constructs a [MTL4PipelineStageDynamicLinkingDescriptor] from an unsafe.Pointer.
//
// Groups together properties to drive the dynamic linking process of a pipeline stage.
func MTL4PipelineStageDynamicLinkingDescriptorFrom(ptr unsafe.Pointer) MTL4PipelineStageDynamicLinkingDescriptor {
	return MTL4PipelineStageDynamicLinkingDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// Provides the array of binary functions to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/binaryLinkedFunctions
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) BinaryLinkedFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("binaryLinkedFunctions"))
	return rv
}


// Provides the array of binary functions to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/binaryLinkedFunctions
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) SetBinaryLinkedFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setBinaryLinkedFunctions:"), nsArray)
}


// Limits the maximum depth of the call stack for indirect function calls in the pipeline stage function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/maxCallStackDepth
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxCallStackDepth"))
	return rv
}


// Limits the maximum depth of the call stack for indirect function calls in the pipeline stage function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/maxCallStackDepth
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxCallStackDepth:"), value)
}


// Provides an array of dynamic libraries the compiler loads when it builds the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/preloadedLibraries
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) PreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("preloadedLibraries"))
	return rv
}


// Provides an array of dynamic libraries the compiler loads when it builds the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/preloadedLibraries
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) SetPreloadedLibraries(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreloadedLibraries:"), nsArray)
}








