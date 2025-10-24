// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRenderPipelineFunctionsDescriptor */


/* debug [class_header]: Header for MTLRenderPipelineFunctionsDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RenderPipelineFunctionsDescriptor */
// An interface definition for the [RenderPipelineFunctionsDescriptor] class.
type IRenderPipelineFunctionsDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RenderPipelineFunctionsDescriptor */
	// properties:
	FragmentAdditionalBinaryFunctions() []objc.ID
	SetFragmentAdditionalBinaryFunctions(value []objc.ID)
	TileAdditionalBinaryFunctions() []objc.ID
	SetTileAdditionalBinaryFunctions(value []objc.ID)
	VertexAdditionalBinaryFunctions() []objc.ID
	SetVertexAdditionalBinaryFunctions(value []objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RenderPipelineFunctionsDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RenderPipelineFunctionsDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _RenderPipelineFunctionsDescriptorClass) Alloc() RenderPipelineFunctionsDescriptor {
	rv := objc.Send[RenderPipelineFunctionsDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RenderPipelineFunctionsDescriptor */
// A collection of functions for updating a render pipeline.
//
// When you create a render pipeline that takes visible functions as parameters, you must specify all possible functions that the render pipeline can call. If you already have a pipeline, you can create a new render pipeline with the same configuration but additional callable functions. To create the new pipeline state, configure an instance with the additional callable functions to add, and then call the pipeline state’s method, passing the descriptor.


// A collection of functions for updating a render pipeline.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RenderPipelineFunctionsDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RenderPipelineFunctionsDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RenderPipelineFunctionsDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RenderPipelineFunctionsDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RenderPipelineFunctionsDescriptor */

// The fragment functions to add to the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) FragmentAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("fragmentAdditionalBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: fragmentAdditionalBinaryFunctions */


// The fragment functions to add to the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions(value []objc.ID) {
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
}/* debug [instance_properties/setter]: fragmentAdditionalBinaryFunctions */


// The tile functions to add to the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/tileAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) TileAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("tileAdditionalBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: tileAdditionalBinaryFunctions */


// The tile functions to add to the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/tileAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) SetTileAdditionalBinaryFunctions(value []objc.ID) {
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
}/* debug [instance_properties/setter]: tileAdditionalBinaryFunctions */


// The vertex functions to add to the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) VertexAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("vertexAdditionalBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: vertexAdditionalBinaryFunctions */


// The vertex functions to add to the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (r_ RenderPipelineFunctionsDescriptor) SetVertexAdditionalBinaryFunctions(value []objc.ID) {
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
}/* debug [instance_properties/setter]: vertexAdditionalBinaryFunctions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRenderPipelineFunctionsDescriptor */



