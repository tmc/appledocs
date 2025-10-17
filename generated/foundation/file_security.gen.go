// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FileSecurity] class.
var FileSecurityClass = _FileSecurityClass{objc.GetClass("NSFileSecurity")}

type _FileSecurityClass struct {
	class objc.Class
}

type FileSecurity struct {
	objc.ID
}

func FileSecurityFrom(ptr unsafe.Pointer) FileSecurity {
	return FileSecurity{
		ID: objc.ID(ptr),
	}
}




