// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RenderPipelineFunctionsDescriptor] class.
var (
	RenderPipelineFunctionsDescriptorClass     _RenderPipelineFunctionsDescriptorClass
	RenderPipelineFunctionsDescriptorClassOnce sync.Once
)

func getRenderPipelineFunctionsDescriptorClass() _RenderPipelineFunctionsDescriptorClass {
	RenderPipelineFunctionsDescriptorClassOnce.Do(func() {
		RenderPipelineFunctionsDescriptorClass = _RenderPipelineFunctionsDescriptorClass{objc.GetClass("MTLRenderPipelineFunctionsDescriptor")}
	})
	return RenderPipelineFunctionsDescriptorClass
}

type _RenderPipelineFunctionsDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [RenderPipelineFunctionsDescriptor] class.
type IRenderPipelineFunctionsDescriptor interface {
	objectivec.IObject
	FragmentAdditionalBinaryFunctions() []objc.ID
	SetFragmentAdditionalBinaryFunctions(value []objc.ID)
	TileAdditionalBinaryFunctions() []objc.ID
	SetTileAdditionalBinaryFunctions(value []objc.ID)
	VertexAdditionalBinaryFunctions() []objc.ID
	SetVertexAdditionalBinaryFunctions(value []objc.ID)
}

// A collection of functions for updating a render pipeline.
//
// When you create a render pipeline that takes visible functions as parameters, you must specify all possible functions that the render pipeline can call. If you already have a pipeline, you can create a new render pipeline with the same configuration but additional callable functions. To create the new pipeline state, configure an instance with the additional callable functions to add, and then call the pipeline state’s method, passing the descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor
type RenderPipelineFunctionsDescriptor struct {
	objectivec.Object
}

// RenderPipelineFunctionsDescriptorFrom constructs a [RenderPipelineFunctionsDescriptor] from an unsafe.Pointer.
//
// A collection of functions for updating a render pipeline.
func RenderPipelineFunctionsDescriptorFrom(ptr unsafe.Pointer) RenderPipelineFunctionsDescriptor {
	return RenderPipelineFunctionsDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RenderPipelineFunctionsDescriptorClass) Alloc() RenderPipelineFunctionsDescriptor {
	rv := objc.Send[RenderPipelineFunctionsDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RenderPipelineFunctionsDescriptorClass) New() RenderPipelineFunctionsDescriptor {
	rv := objc.Send[RenderPipelineFunctionsDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPipelineFunctionsDescriptor) Init() RenderPipelineFunctionsDescriptor {
	rv := objc.Send[RenderPipelineFunctionsDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPipelineFunctionsDescriptor) Autorelease() RenderPipelineFunctionsDescriptor {
	rv := objc.Send[RenderPipelineFunctionsDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPipelineFunctionsDescriptor creates a new RenderPipelineFunctionsDescriptor instance.
func NewRenderPipelineFunctionsDescriptor() RenderPipelineFunctionsDescriptor {
	return getRenderPipelineFunctionsDescriptorClass().New()
}


// The fragment functions to add to the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) FragmentAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("fragmentAdditionalBinaryFunctions"))
	return rv
}


// SetFragmentAdditionalBinaryFunctions sets the value of the fragmentAdditionalBinaryFunctions property.
// The fragment functions to add to the render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setFragmentAdditionalBinaryFunctions:"), nsArray)
}

// The tile functions to add to the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/tileAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) TileAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("tileAdditionalBinaryFunctions"))
	return rv
}


// SetTileAdditionalBinaryFunctions sets the value of the tileAdditionalBinaryFunctions property.
// The tile functions to add to the render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/tileAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) SetTileAdditionalBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setTileAdditionalBinaryFunctions:"), nsArray)
}

// The vertex functions to add to the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) VertexAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("vertexAdditionalBinaryFunctions"))
	return rv
}


// SetVertexAdditionalBinaryFunctions sets the value of the vertexAdditionalBinaryFunctions property.
// The vertex functions to add to the render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) SetVertexAdditionalBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertexAdditionalBinaryFunctions:"), nsArray)
}



