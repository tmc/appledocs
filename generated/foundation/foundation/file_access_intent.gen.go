// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FileAccessIntent] class.
var FileAccessIntentClass objc.Class

func init() {
	FileAccessIntentClass = objc.GetClass("NSFileAccessIntent")
}

type FileAccessIntent struct {
	objc.ID
}

func FileAccessIntentFrom(ptr unsafe.Pointer) FileAccessIntent {
	return FileAccessIntent{
		ID: objc.ID(ptr),
	}
}




