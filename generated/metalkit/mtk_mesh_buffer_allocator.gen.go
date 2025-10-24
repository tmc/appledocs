// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTKMeshBufferAllocator */


/* debug [class_header]: Header for MTKMeshBufferAllocator */
// The class instance for the [MeshBufferAllocator] class.
var (
	MeshBufferAllocatorClass     _MeshBufferAllocatorClass
	MeshBufferAllocatorClassOnce sync.Once
)

func getMeshBufferAllocatorClass() _MeshBufferAllocatorClass {
	MeshBufferAllocatorClassOnce.Do(func() {
		MeshBufferAllocatorClass = _MeshBufferAllocatorClass{objc.GetClass("MTKMeshBufferAllocator")}
	})
	return MeshBufferAllocatorClass
}

type _MeshBufferAllocatorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MeshBufferAllocator */
// An interface definition for the [MeshBufferAllocator] class.
type IMeshBufferAllocator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MeshBufferAllocator */
	// properties:
	Device() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MeshBufferAllocator */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MeshBufferAllocator */
// Alloc allocates a new instance without initialization.
func (mc _MeshBufferAllocatorClass) Alloc() MeshBufferAllocator {
	rv := objc.Send[MeshBufferAllocator](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MeshBufferAllocatorClass) New() MeshBufferAllocator {
	rv := objc.Send[MeshBufferAllocator](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MeshBufferAllocator) Init() MeshBufferAllocator {
	rv := objc.Send[MeshBufferAllocator](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MeshBufferAllocator) Autorelease() MeshBufferAllocator {
	rv := objc.Send[MeshBufferAllocator](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMeshBufferAllocator creates a new MeshBufferAllocator instance.
func NewMeshBufferAllocator() MeshBufferAllocator {
	return getMeshBufferAllocatorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MeshBufferAllocator */
// An interface for allocating a MetalKit buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.


// An interface for allocating a MetalKit buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator
type MeshBufferAllocator struct {
	objectivec.Object
}

// MeshBufferAllocatorFrom constructs a [MeshBufferAllocator] from an unsafe.Pointer.
//
// An interface for allocating a MetalKit buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
func MeshBufferAllocatorFrom(ptr unsafe.Pointer) MeshBufferAllocator {
	return MeshBufferAllocator{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MeshBufferAllocator */

// Initializes a new allocator object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator/init(device:)
func NewMeshBufferAllocatorWithDevice(device unsafe.Pointer) MeshBufferAllocator {
	instance := getMeshBufferAllocatorClass().Alloc()
	rv := objc.Send[MeshBufferAllocator](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMeshBufferAllocatorWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MeshBufferAllocator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MeshBufferAllocator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MeshBufferAllocator */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MeshBufferAllocator */

// The device used to create Metal objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator/device
func (m_ MeshBufferAllocator) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTKMeshBufferAllocator */


