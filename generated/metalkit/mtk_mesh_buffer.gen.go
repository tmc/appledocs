// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MeshBuffer] class.
type IMeshBuffer interface {
	objectivec.IObject
}

// A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MeshBufferClass) Alloc() MeshBuffer {
	rv := objc.Send[MeshBuffer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The allocator object used to create this mesh buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/allocator
func (m_ MeshBuffer) Allocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("allocator"))
	return rv
}

// The Metal buffer backing all vertex and index data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/buffer
func (m_ MeshBuffer) Buffer() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("buffer"))
	return rv
}

// The logical size of the Metal buffer, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/length
func (m_ MeshBuffer) Length() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("length"))
	return rv
}

// The byte offset of the data within the Metal buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/offset
func (m_ MeshBuffer) Offset() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("offset"))
	return rv
}

// The type of data contained in the originating Model I/O buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/type
func (m_ MeshBuffer) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("type"))
	return rv
}

// The zone, if any, from which this mesh buffer was created.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/zone
func (m_ MeshBuffer) Zone() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("zone"))
	return rv
}



