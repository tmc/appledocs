// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRenderPipelineReflection */


/* debug [class_header]: Header for MTLRenderPipelineReflection */
// The class instance for the [RenderPipelineReflection] class.
var (
	RenderPipelineReflectionClass     _RenderPipelineReflectionClass
	RenderPipelineReflectionClassOnce sync.Once
)

func getRenderPipelineReflectionClass() _RenderPipelineReflectionClass {
	RenderPipelineReflectionClassOnce.Do(func() {
		RenderPipelineReflectionClass = _RenderPipelineReflectionClass{objc.GetClass("MTLRenderPipelineReflection")}
	})
	return RenderPipelineReflectionClass
}

type _RenderPipelineReflectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RenderPipelineReflection */
// An interface definition for the [RenderPipelineReflection] class.
type IRenderPipelineReflection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RenderPipelineReflection */
	// properties:
	FragmentArguments() []Argument
	FragmentBindings() []objc.ID
	MeshBindings() []objc.ID
	ObjectBindings() []objc.ID
	TileArguments() []Argument
	TileBindings() []objc.ID
	VertexArguments() []Argument
	VertexBindings() []objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RenderPipelineReflection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RenderPipelineReflection */
// Alloc allocates a new instance without initialization.
func (rc _RenderPipelineReflectionClass) Alloc() RenderPipelineReflection {
	rv := objc.Send[RenderPipelineReflection](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderPipelineReflectionClass) New() RenderPipelineReflection {
	rv := objc.Send[RenderPipelineReflection](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPipelineReflection) Init() RenderPipelineReflection {
	rv := objc.Send[RenderPipelineReflection](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPipelineReflection) Autorelease() RenderPipelineReflection {
	rv := objc.Send[RenderPipelineReflection](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPipelineReflection creates a new RenderPipelineReflection instance.
func NewRenderPipelineReflection() RenderPipelineReflection {
	return getRenderPipelineReflectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RenderPipelineReflection */
// Information about the arguments of a graphics function.
//
// The class is an interface that represents the parameters for the shaders in a render pipeline state (see ). Each pipeline state can include object, mesh, vertex, fragment, and tile shaders. You create a reflection instance at the same time as the pipeline state that it represents by calling the appropriate method. For example, the and methods create the pipeline state and the reflection instances at the same time. For more information, see .


// Information about the arguments of a graphics function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection
type RenderPipelineReflection struct {
	objectivec.Object
}

// RenderPipelineReflectionFrom constructs a [RenderPipelineReflection] from an unsafe.Pointer.
//
// Information about the arguments of a graphics function.
func RenderPipelineReflectionFrom(ptr unsafe.Pointer) RenderPipelineReflection {
	return RenderPipelineReflection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RenderPipelineReflection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RenderPipelineReflection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RenderPipelineReflection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RenderPipelineReflection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RenderPipelineReflection */

// An array of argument instances, each of which represent a parameter of the pipeline state’s fragment shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/fragmentArguments
func (r_ RenderPipelineReflection) FragmentArguments() []Argument {
	rv := objc.Send[[]Argument](r_.ID, objc.Sel("fragmentArguments"))
	return rv
}/* debug [instance_properties/getter]: fragmentArguments */


// An array of binding instances, each of which represents a parameter of the pipeline state’s fragment shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/fragmentBindings
func (r_ RenderPipelineReflection) FragmentBindings() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("fragmentBindings"))
	return rv
}/* debug [instance_properties/getter]: fragmentBindings */


// An array of binding instances, each of which represents a parameter of the pipeline state’s mesh shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/meshBindings
func (r_ RenderPipelineReflection) MeshBindings() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("meshBindings"))
	return rv
}/* debug [instance_properties/getter]: meshBindings */


// An array of binding instances, each of which represents a parameter of the pipeline state’s object shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/objectBindings
func (r_ RenderPipelineReflection) ObjectBindings() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("objectBindings"))
	return rv
}/* debug [instance_properties/getter]: objectBindings */


// An array of argument instances, each of which represent a parameter of the pipeline state’s tile shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/tileArguments
func (r_ RenderPipelineReflection) TileArguments() []Argument {
	rv := objc.Send[[]Argument](r_.ID, objc.Sel("tileArguments"))
	return rv
}/* debug [instance_properties/getter]: tileArguments */


// An array of binding instances, each of which represents a parameter of the pipeline state’s tile shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/tileBindings
func (r_ RenderPipelineReflection) TileBindings() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("tileBindings"))
	return rv
}/* debug [instance_properties/getter]: tileBindings */


// An array of argument instances, each of which represent a parameter of the pipeline state’s vertex shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/vertexArguments
func (r_ RenderPipelineReflection) VertexArguments() []Argument {
	rv := objc.Send[[]Argument](r_.ID, objc.Sel("vertexArguments"))
	return rv
}/* debug [instance_properties/getter]: vertexArguments */


// An array of binding instances, each of which represents a parameter of the pipeline state’s vertex shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/vertexBindings
func (r_ RenderPipelineReflection) VertexBindings() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("vertexBindings"))
	return rv
}/* debug [instance_properties/getter]: vertexBindings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRenderPipelineReflection */



