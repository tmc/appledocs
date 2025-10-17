// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FileSecurity] class.
var FileSecurityClass objc.Class

func init() {
	FileSecurityClass = objc.GetClass("NSFileSecurity")
}

type FileSecurity struct {
	objc.ID
}

func FileSecurityFrom(ptr unsafe.Pointer) FileSecurity {
	return FileSecurity{
		ID: objc.ID(ptr),
	}
}



