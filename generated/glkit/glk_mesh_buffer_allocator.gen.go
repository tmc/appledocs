// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GLKMeshBufferAllocator */


/* debug [class_header]: Header for GLKMeshBufferAllocator */
// The class instance for the [GLKMeshBufferAllocator] class.
var (
	GLKMeshBufferAllocatorClass     _GLKMeshBufferAllocatorClass
	GLKMeshBufferAllocatorClassOnce sync.Once
)

func getGLKMeshBufferAllocatorClass() _GLKMeshBufferAllocatorClass {
	GLKMeshBufferAllocatorClassOnce.Do(func() {
		GLKMeshBufferAllocatorClass = _GLKMeshBufferAllocatorClass{objc.GetClass("GLKMeshBufferAllocator")}
	})
	return GLKMeshBufferAllocatorClass
}

type _GLKMeshBufferAllocatorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKMeshBufferAllocator */
// An interface definition for the [GLKMeshBufferAllocator] class.
type IGLKMeshBufferAllocator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GLKMeshBufferAllocator */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKMeshBufferAllocator */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKMeshBufferAllocator */
// Alloc allocates a new instance without initialization.
func (gc _GLKMeshBufferAllocatorClass) Alloc() GLKMeshBufferAllocator {
	rv := objc.Send[GLKMeshBufferAllocator](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKMeshBufferAllocatorClass) New() GLKMeshBufferAllocator {
	rv := objc.Send[GLKMeshBufferAllocator](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKMeshBufferAllocator) Init() GLKMeshBufferAllocator {
	rv := objc.Send[GLKMeshBufferAllocator](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKMeshBufferAllocator) Autorelease() GLKMeshBufferAllocator {
	rv := objc.Send[GLKMeshBufferAllocator](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKMeshBufferAllocator creates a new GLKMeshBufferAllocator instance.
func NewGLKMeshBufferAllocator() GLKMeshBufferAllocator {
	return getGLKMeshBufferAllocatorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKMeshBufferAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBufferAllocator
type GLKMeshBufferAllocator struct {
	objectivec.Object
}

// GLKMeshBufferAllocatorFrom constructs a [GLKMeshBufferAllocator] from an unsafe.Pointer.
func GLKMeshBufferAllocatorFrom(ptr unsafe.Pointer) GLKMeshBufferAllocator {
	return GLKMeshBufferAllocator{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKMeshBufferAllocator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKMeshBufferAllocator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKMeshBufferAllocator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKMeshBufferAllocator */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKMeshBufferAllocator */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKMeshBufferAllocator */



