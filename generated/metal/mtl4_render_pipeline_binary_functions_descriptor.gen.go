// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4RenderPipelineBinaryFunctionsDescriptor */


/* debug [class_header]: Header for MTL4RenderPipelineBinaryFunctionsDescriptor */
// The class instance for the [MTL4RenderPipelineBinaryFunctionsDescriptor] class.
var (
	MTL4RenderPipelineBinaryFunctionsDescriptorClass     _MTL4RenderPipelineBinaryFunctionsDescriptorClass
	MTL4RenderPipelineBinaryFunctionsDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineBinaryFunctionsDescriptorClass() _MTL4RenderPipelineBinaryFunctionsDescriptorClass {
	MTL4RenderPipelineBinaryFunctionsDescriptorClassOnce.Do(func() {
		MTL4RenderPipelineBinaryFunctionsDescriptorClass = _MTL4RenderPipelineBinaryFunctionsDescriptorClass{objc.GetClass("MTL4RenderPipelineBinaryFunctionsDescriptor")}
	})
	return MTL4RenderPipelineBinaryFunctionsDescriptorClass
}

type _MTL4RenderPipelineBinaryFunctionsDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4RenderPipelineBinaryFunctionsDescriptor */
// An interface definition for the [MTL4RenderPipelineBinaryFunctionsDescriptor] class.
type IMTL4RenderPipelineBinaryFunctionsDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4RenderPipelineBinaryFunctionsDescriptor */
	// properties:
	FragmentAdditionalBinaryFunctions() []objc.ID
	SetFragmentAdditionalBinaryFunctions(value []objc.ID)
	MeshAdditionalBinaryFunctions() []objc.ID
	SetMeshAdditionalBinaryFunctions(value []objc.ID)
	ObjectAdditionalBinaryFunctions() []objc.ID
	SetObjectAdditionalBinaryFunctions(value []objc.ID)
	TileAdditionalBinaryFunctions() []objc.ID
	SetTileAdditionalBinaryFunctions(value []objc.ID)
	VertexAdditionalBinaryFunctions() []objc.ID
	SetVertexAdditionalBinaryFunctions(value []objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4RenderPipelineBinaryFunctionsDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4RenderPipelineBinaryFunctionsDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineBinaryFunctionsDescriptorClass) Alloc() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4RenderPipelineBinaryFunctionsDescriptorClass) New() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) Init() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) Autorelease() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineBinaryFunctionsDescriptor creates a new MTL4RenderPipelineBinaryFunctionsDescriptor instance.
func NewMTL4RenderPipelineBinaryFunctionsDescriptor() MTL4RenderPipelineBinaryFunctionsDescriptor {
	return getMTL4RenderPipelineBinaryFunctionsDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4RenderPipelineBinaryFunctionsDescriptor */
// Allows you to specify additional binary functions to link to each stage of a render pipeline.


// Allows you to specify additional binary functions to link to each stage of a render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor
type MTL4RenderPipelineBinaryFunctionsDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineBinaryFunctionsDescriptorFrom constructs a [MTL4RenderPipelineBinaryFunctionsDescriptor] from an unsafe.Pointer.
//
// Allows you to specify additional binary functions to link to each stage of a render pipeline.
func MTL4RenderPipelineBinaryFunctionsDescriptorFrom(ptr unsafe.Pointer) MTL4RenderPipelineBinaryFunctionsDescriptor {
	return MTL4RenderPipelineBinaryFunctionsDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4RenderPipelineBinaryFunctionsDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4RenderPipelineBinaryFunctionsDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4RenderPipelineBinaryFunctionsDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4RenderPipelineBinaryFunctionsDescriptor */

// Resets this descriptor to its default state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/reset()
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4RenderPipelineBinaryFunctionsDescriptor */

// Provides an array of binary functions representing additional binary fragment shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) FragmentAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("fragmentAdditionalBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: fragmentAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary fragment shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentAdditionalBinaryFunctions:"), nsArray)
}/* debug [instance_properties/setter]: fragmentAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary mesh shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/meshAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) MeshAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("meshAdditionalBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: meshAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary mesh shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/meshAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetMeshAdditionalBinaryFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshAdditionalBinaryFunctions:"), nsArray)
}/* debug [instance_properties/setter]: meshAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary object shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/objectAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) ObjectAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("objectAdditionalBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: objectAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary object shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/objectAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetObjectAdditionalBinaryFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectAdditionalBinaryFunctions:"), nsArray)
}/* debug [instance_properties/setter]: objectAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary tile shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/tileAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) TileAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("tileAdditionalBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: tileAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary tile shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/tileAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetTileAdditionalBinaryFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileAdditionalBinaryFunctions:"), nsArray)
}/* debug [instance_properties/setter]: tileAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary vertex shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) VertexAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("vertexAdditionalBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: vertexAdditionalBinaryFunctions */


// Provides an array of binary functions representing additional binary vertex shader functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetVertexAdditionalBinaryFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexAdditionalBinaryFunctions:"), nsArray)
}/* debug [instance_properties/setter]: vertexAdditionalBinaryFunctions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4RenderPipelineBinaryFunctionsDescriptor */



