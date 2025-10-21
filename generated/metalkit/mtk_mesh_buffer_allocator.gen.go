// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MeshBufferAllocator] class.
type IMeshBufferAllocator interface {
	objectivec.IObject
}

// An interface for allocating a MetalKit buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MeshBufferAllocatorClass) Alloc() MeshBufferAllocator {
	rv := objc.Send[MeshBufferAllocator](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a new allocator object.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator/init(device:)
func NewMeshBufferAllocatorWithDevice(device objc.ID) MeshBufferAllocator {
	instance := getMeshBufferAllocatorClass().Alloc()
	rv := objc.Send[MeshBufferAllocator](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// The device used to create Metal objects.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator/device
func (m_ MeshBufferAllocator) Device() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("device"))
	return rv
}


