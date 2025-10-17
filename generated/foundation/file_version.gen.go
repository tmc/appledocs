// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FileVersion] class.
var FileVersionClass objc.Class

func init() {
	FileVersionClass = objc.GetClass("NSFileVersion")
}

type FileVersion struct {
	objc.ID
}

func FileVersionFrom(ptr unsafe.Pointer) FileVersion {
	return FileVersion{
		ID: objc.ID(ptr),
	}
}



