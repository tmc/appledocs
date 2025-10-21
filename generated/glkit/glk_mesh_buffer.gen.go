// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GLKMeshBuffer] class.
type IGLKMeshBuffer interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer
type GLKMeshBuffer struct {
	objectivec.Object
}

// GLKMeshBufferFrom constructs a [GLKMeshBuffer] from an unsafe.Pointer.
func GLKMeshBufferFrom(ptr unsafe.Pointer) GLKMeshBuffer {
	return GLKMeshBuffer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKMeshBufferClass) Alloc() GLKMeshBuffer {
	rv := objc.Send[GLKMeshBuffer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/allocator
func (g_ GLKMeshBuffer) Allocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allocator"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/glBufferName
func (g_ GLKMeshBuffer) GlBufferName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("glBufferName"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/length
func (g_ GLKMeshBuffer) Length() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("length"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/offset
func (g_ GLKMeshBuffer) Offset() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("offset"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/type
func (g_ GLKMeshBuffer) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("type"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMeshBuffer/zone
func (g_ GLKMeshBuffer) Zone() objc.ID {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("zone"))
	return rv
}



