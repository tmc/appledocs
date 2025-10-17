// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FileWrapper] class.
var FileWrapperClass objc.Class

func init() {
	FileWrapperClass = objc.GetClass("NSFileWrapper")
}

type FileWrapper struct {
	objc.ID
}

func FileWrapperFrom(ptr unsafe.Pointer) FileWrapper {
	return FileWrapper{
		ID: objc.ID(ptr),
	}
}



