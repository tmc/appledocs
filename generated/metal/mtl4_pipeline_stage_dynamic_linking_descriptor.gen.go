// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4PipelineStageDynamicLinkingDescriptor */


/* debug [class_header]: Header for MTL4PipelineStageDynamicLinkingDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4PipelineStageDynamicLinkingDescriptor */
// An interface definition for the [MTL4PipelineStageDynamicLinkingDescriptor] class.
type IMTL4PipelineStageDynamicLinkingDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4PipelineStageDynamicLinkingDescriptor */
	// properties:
	BinaryLinkedFunctions() []objc.ID
	SetBinaryLinkedFunctions(value []objc.ID)
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)
	PreloadedLibraries() []objc.ID
	SetPreloadedLibraries(value []objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4PipelineStageDynamicLinkingDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4PipelineStageDynamicLinkingDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4PipelineStageDynamicLinkingDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4PipelineStageDynamicLinkingDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4PipelineStageDynamicLinkingDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4PipelineStageDynamicLinkingDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4PipelineStageDynamicLinkingDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4PipelineStageDynamicLinkingDescriptor */

// Provides the array of binary functions to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/binaryLinkedFunctions
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) BinaryLinkedFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("binaryLinkedFunctions"))
	return rv
}/* debug [instance_properties/getter]: binaryLinkedFunctions */


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
}/* debug [instance_properties/setter]: binaryLinkedFunctions */


// Limits the maximum depth of the call stack for indirect function calls in the pipeline stage function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/maxCallStackDepth
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxCallStackDepth"))
	return rv
}/* debug [instance_properties/getter]: maxCallStackDepth */


// Limits the maximum depth of the call stack for indirect function calls in the pipeline stage function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/maxCallStackDepth
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxCallStackDepth:"), value)
}/* debug [instance_properties/setter]: maxCallStackDepth */


// Provides an array of dynamic libraries the compiler loads when it builds the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/preloadedLibraries
func (m_ MTL4PipelineStageDynamicLinkingDescriptor) PreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("preloadedLibraries"))
	return rv
}/* debug [instance_properties/getter]: preloadedLibraries */


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
}/* debug [instance_properties/setter]: preloadedLibraries */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4PipelineStageDynamicLinkingDescriptor */



