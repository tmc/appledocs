// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTKSubmesh] class.
var mTKSubmeshClass = _MTKSubmeshClass{objc.GetClass("MTKSubmesh")}

type _MTKSubmeshClass struct {
	class objc.Class
}

// A container for the index data of a Model I/O submesh, suitable for use in a Metal app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh

type MTKSubmesh struct {
	objectivec.Object
}

// MTKSubmeshFrom constructs a [MTKSubmesh] from an unsafe.Pointer.
//
// A container for the index data of a Model I/O submesh, suitable for use in a Metal app.
func MTKSubmeshFrom(ptr unsafe.Pointer) MTKSubmesh {
	return MTKSubmesh{objectivec.Object{objc.ID(ptr)}}
}



