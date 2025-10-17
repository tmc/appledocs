// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Pipe] class.
var pipeClass = _PipeClass{objc.GetClass("NSPipe")}

type _PipeClass struct {
	class objc.Class
}

// An interface definition for the [Pipe] class.
type IPipe interface {
	objectivec.IObject
}

// A one-way communications channel between related processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Pipe

type Pipe struct {
	objectivec.Object
}

// PipeFrom constructs a [Pipe] from an unsafe.Pointer.
//
// A one-way communications channel between related processes.
func PipeFrom(ptr unsafe.Pointer) Pipe {
	return Pipe{objectivec.Object{objc.ID(ptr)}}
}



