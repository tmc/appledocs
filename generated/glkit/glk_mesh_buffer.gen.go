// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GLKMeshBuffer */


/* debug [class_header]: Header for GLKMeshBuffer */
// The class instance for the [GLKMeshBuffer] class.
var (
	GLKMeshBufferClass     _GLKMeshBufferClass
	GLKMeshBufferClassOnce sync.Once
)

func getGLKMeshBufferClass() _GLKMeshBufferClass {
	GLKMeshBufferClassOnce.Do(func() {
		GLKMeshBufferClass = _GLKMeshBufferClass{objc.GetClass("GLKMeshBuffer")}
	})
	return GLKMeshBufferClass
}

type _GLKMeshBufferClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKMeshBuffer */
// An interface definition for the [GLKMeshBuffer] class.
type IGLKMeshBuffer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GLKMeshBuffer */
	// properties:
	Allocator() IGLKMeshBufferAllocator
	GlBufferName() unsafe.Pointer
	Length() uint
	Offset() uint
	Type() unsafe.Pointer
	Zone() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKMeshBuffer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKMeshBuffer */
// Alloc allocates a new instance without initialization.
func (gc _GLKMeshBufferClass) Alloc() GLKMeshBuffer {
	rv := objc.Send[GLKMeshBuffer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKMeshBufferClass) New() GLKMeshBuffer {
	rv := objc.Send[GLKMeshBuffer](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKMeshBuffer) Init() GLKMeshBuffer {
	rv := objc.Send[GLKMeshBuffer](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKMeshBuffer) Autorelease() GLKMeshBuffer {
	rv := objc.Send[GLKMeshBuffer](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKMeshBuffer creates a new GLKMeshBuffer instance.
func NewGLKMeshBuffer() GLKMeshBuffer {
	return getGLKMeshBufferClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKMeshBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer
type GLKMeshBuffer struct {
	objectivec.Object
}

// GLKMeshBufferFrom constructs a [GLKMeshBuffer] from an unsafe.Pointer.
func GLKMeshBufferFrom(ptr unsafe.Pointer) GLKMeshBuffer {
	return GLKMeshBuffer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKMeshBuffer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKMeshBuffer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKMeshBuffer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKMeshBuffer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKMeshBuffer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/allocator
func (g_ GLKMeshBuffer) Allocator() IGLKMeshBufferAllocator {
	rv := objc.Send[GLKMeshBufferAllocator](g_.ID, objc.Sel("allocator"))
	return rv
}/* debug [instance_properties/getter]: allocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/glBufferName
func (g_ GLKMeshBuffer) GlBufferName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("glBufferName"))
	return rv
}/* debug [instance_properties/getter]: glBufferName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/length
func (g_ GLKMeshBuffer) Length() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/offset
func (g_ GLKMeshBuffer) Offset() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/type
func (g_ GLKMeshBuffer) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/zone
func (g_ GLKMeshBuffer) Zone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("zone"))
	return rv
}/* debug [instance_properties/getter]: zone */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKMeshBuffer */



