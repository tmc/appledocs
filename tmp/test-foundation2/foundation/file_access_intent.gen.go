// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var fileAccessIntentClass _FileAccessIntentClass

func init() {
	fileAccessIntentClass = _FileAccessIntentClass{objc.GetClass("NSFileAccessIntent")}
}

type _FileAccessIntentClass struct {
	class objc.Class
}

type FileAccessIntent struct {
	objc.ID
}

func FileAccessIntentFrom(ptr unsafe.Pointer) FileAccessIntent {
	return FileAccessIntent{
		ID: objc.ID(ptr),
	}
}




