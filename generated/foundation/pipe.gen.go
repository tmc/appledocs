// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Pipe] class.
var PipeClass objc.Class

func init() {
	PipeClass = objc.GetClass("NSPipe")
}

type Pipe struct {
	objc.ID
}

func PipeFrom(ptr unsafe.Pointer) Pipe {
	return Pipe{
		ID: objc.ID(ptr),
	}
}



