// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GLKMeshBufferAllocator] class.
type IGLKMeshBufferAllocator interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBufferAllocator
type GLKMeshBufferAllocator struct {
	objectivec.Object
}

// GLKMeshBufferAllocatorFrom constructs a [GLKMeshBufferAllocator] from an unsafe.Pointer.
func GLKMeshBufferAllocatorFrom(ptr unsafe.Pointer) GLKMeshBufferAllocator {
	return GLKMeshBufferAllocator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKMeshBufferAllocatorClass) Alloc() GLKMeshBufferAllocator {
	rv := objc.Send[GLKMeshBufferAllocator](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




