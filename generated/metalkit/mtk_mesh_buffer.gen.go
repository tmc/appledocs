// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTKMeshBuffer] class.
var mTKMeshBufferClass = _MTKMeshBufferClass{objc.GetClass("MTKMeshBuffer")}

type _MTKMeshBufferClass struct {
	class objc.Class
}

// A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app. [Full Topic]
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

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/zone()
func (m_ MTKMeshBuffer) Zone() {
	objc.Send[objc.ID](m_.ID, objc.Sel("zone"))
}


