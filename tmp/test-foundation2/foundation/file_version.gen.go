// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var fileVersionClass _FileVersionClass

func init() {
	fileVersionClass = _FileVersionClass{objc.GetClass("NSFileVersion")}
}

type _FileVersionClass struct {
	class objc.Class
}

type FileVersion struct {
	objc.ID
}

func FileVersionFrom(ptr unsafe.Pointer) FileVersion {
	return FileVersion{
		ID: objc.ID(ptr),
	}
}




