// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTKMeshBuffer */


/* debug [class_header]: Header for MTKMeshBuffer */
// The class instance for the [MeshBuffer] class.
var (
	MeshBufferClass     _MeshBufferClass
	MeshBufferClassOnce sync.Once
)

func getMeshBufferClass() _MeshBufferClass {
	MeshBufferClassOnce.Do(func() {
		MeshBufferClass = _MeshBufferClass{objc.GetClass("MTKMeshBuffer")}
	})
	return MeshBufferClass
}

type _MeshBufferClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MeshBuffer */
// An interface definition for the [MeshBuffer] class.
type IMeshBuffer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MeshBuffer */
	// properties:
	Allocator() IMTKMeshBufferAllocator
	Buffer() unsafe.Pointer
	Length() uint
	Offset() uint
	Type() objectivec.IObject
	Zone() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MeshBuffer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MeshBuffer */
// Alloc allocates a new instance without initialization.
func (mc _MeshBufferClass) Alloc() MeshBuffer {
	rv := objc.Send[MeshBuffer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MeshBufferClass) New() MeshBuffer {
	rv := objc.Send[MeshBuffer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MeshBuffer) Init() MeshBuffer {
	rv := objc.Send[MeshBuffer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MeshBuffer) Autorelease() MeshBuffer {
	rv := objc.Send[MeshBuffer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMeshBuffer creates a new MeshBuffer instance.
func NewMeshBuffer() MeshBuffer {
	return getMeshBufferClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MeshBuffer */
// A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.


// A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer
type MeshBuffer struct {
	objectivec.Object
}

// MeshBufferFrom constructs a [MeshBuffer] from an unsafe.Pointer.
//
// A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
func MeshBufferFrom(ptr unsafe.Pointer) MeshBuffer {
	return MeshBuffer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MeshBuffer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MeshBuffer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MeshBuffer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MeshBuffer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MeshBuffer */

// The allocator object used to create this mesh buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/allocator
func (m_ MeshBuffer) Allocator() IMTKMeshBufferAllocator {
	rv := objc.Send[MeshBufferAllocator](m_.ID, objc.Sel("allocator"))
	return rv
}/* debug [instance_properties/getter]: allocator */


// The Metal buffer backing all vertex and index data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/buffer
func (m_ MeshBuffer) Buffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("buffer"))
	return rv
}/* debug [instance_properties/getter]: buffer */


// The logical size of the Metal buffer, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/length
func (m_ MeshBuffer) Length() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// The byte offset of the data within the Metal buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/offset
func (m_ MeshBuffer) Offset() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The type of data contained in the originating Model I/O buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/type
func (m_ MeshBuffer) Type() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The zone, if any, from which this mesh buffer was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/zone
func (m_ MeshBuffer) Zone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("zone"))
	return rv
}/* debug [instance_properties/getter]: zone */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTKMeshBuffer */



