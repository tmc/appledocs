// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var fileSecurityClass _FileSecurityClass

func init() {
	fileSecurityClass = _FileSecurityClass{objc.GetClass("NSFileSecurity")}
}

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




