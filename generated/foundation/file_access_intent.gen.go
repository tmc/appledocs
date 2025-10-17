// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FileAccessIntent] class.
var FileAccessIntentClass = _FileAccessIntentClass{objc.GetClass("NSFileAccessIntent")}

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




