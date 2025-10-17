// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FileWrapper] class.
var FileWrapperClass = _FileWrapperClass{objc.GetClass("NSFileWrapper")}

type _FileWrapperClass struct {
	class objc.Class
}

type FileWrapper struct {
	objc.ID
}

func FileWrapperFrom(ptr unsafe.Pointer) FileWrapper {
	return FileWrapper{
		ID: objc.ID(ptr),
	}
}




