// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTKMeshBufferAllocator] class.
var mTKMeshBufferAllocatorClass = _MTKMeshBufferAllocatorClass{objc.GetClass("MTKMeshBufferAllocator")}

type _MTKMeshBufferAllocatorClass struct {
	class objc.Class
}

// An interface definition for the [MTKMeshBufferAllocator] class.
type IMTKMeshBufferAllocator interface {
	objectivec.IObject
}

// An interface for allocating a MetalKit buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator

type MTKMeshBufferAllocator struct {
	objectivec.Object
}

// MTKMeshBufferAllocatorFrom constructs a [MTKMeshBufferAllocator] from an unsafe.Pointer.
//
// An interface for allocating a MetalKit buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
func MTKMeshBufferAllocatorFrom(ptr unsafe.Pointer) MTKMeshBufferAllocator {
	return MTKMeshBufferAllocator{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MTKMeshBufferAllocatorClass) Alloc() MTKMeshBufferAllocator {
	rv := objc.Send[MTKMeshBufferAllocator](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MTKMeshBufferAllocatorClass) New() MTKMeshBufferAllocator {
	rv := objc.Send[MTKMeshBufferAllocator](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTKMeshBufferAllocator) Init() MTKMeshBufferAllocator {
	rv := objc.Send[MTKMeshBufferAllocator](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTKMeshBufferAllocator) Autorelease() MTKMeshBufferAllocator {
	rv := objc.Send[MTKMeshBufferAllocator](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKMeshBufferAllocator creates a new MTKMeshBufferAllocator instance.
func NewMTKMeshBufferAllocator() MTKMeshBufferAllocator {
	return mTKMeshBufferAllocatorClass.New()
}


// Initializes a new allocator object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator/init(device:)
func NewMTKMeshBufferAllocatorWithDevice(device unsafe.Pointer) MTKMeshBufferAllocator {
	instance := mTKMeshBufferAllocatorClass.Alloc()
	rv := objc.Send[MTKMeshBufferAllocator](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



