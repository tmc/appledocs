// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var PipeClass _PipeClass

func init() {
	PipeClass = _PipeClass{objc.GetClass("NSPipe")}
}

type _PipeClass struct {
	class objc.Class
}

type Pipe struct {
	objc.ID
}

func PipeFrom(ptr unsafe.Pointer) Pipe {
	return Pipe{
		ID: objc.ID(ptr),
	}
}




