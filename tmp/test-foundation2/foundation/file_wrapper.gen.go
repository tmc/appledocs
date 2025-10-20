// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var fileWrapperClass _FileWrapperClass

func init() {
	fileWrapperClass = _FileWrapperClass{objc.GetClass("NSFileWrapper")}
}

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




