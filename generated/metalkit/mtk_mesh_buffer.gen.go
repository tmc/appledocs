// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTKMeshBuffer] class.
var (
	mTKMeshBufferClass     _MTKMeshBufferClass
	mTKMeshBufferClassOnce sync.Once
)

func getMTKMeshBufferClass() _MTKMeshBufferClass {
	mTKMeshBufferClassOnce.Do(func() {
		mTKMeshBufferClass = _MTKMeshBufferClass{objc.GetClass("MTKMeshBuffer")}
	})
	return mTKMeshBufferClass
}

type _MTKMeshBufferClass struct {
	class objc.Class
}

// An interface definition for the [MTKMeshBuffer] class.
type IMTKMeshBuffer interface {
	objectivec.IObject
	Zone()
}

// A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer
type MTKMeshBuffer struct {
	objectivec.Object
}

// MTKMeshBufferFrom constructs a [MTKMeshBuffer] from an unsafe.Pointer.
//
// A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
func MTKMeshBufferFrom(ptr unsafe.Pointer) MTKMeshBuffer {
	return MTKMeshBuffer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTKMeshBufferClass) Alloc() MTKMeshBuffer {
	rv := objc.Send[MTKMeshBuffer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTKMeshBufferClass) New() MTKMeshBuffer {
	rv := objc.Send[MTKMeshBuffer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTKMeshBuffer) Init() MTKMeshBuffer {
	rv := objc.Send[MTKMeshBuffer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTKMeshBuffer) Autorelease() MTKMeshBuffer {
	rv := objc.Send[MTKMeshBuffer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKMeshBuffer creates a new MTKMeshBuffer instance.
func NewMTKMeshBuffer() MTKMeshBuffer {
	return getMTKMeshBufferClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/zone()
func (m_ MTKMeshBuffer) Zone() {
	objc.Send[objc.ID](m_.ID, objc.Sel("zone"))
}


