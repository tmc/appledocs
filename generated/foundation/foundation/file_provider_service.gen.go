// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FileProviderService] class.
var FileProviderServiceClass objc.Class

func init() {
	FileProviderServiceClass = objc.GetClass("NSFileProviderService")
}

type FileProviderService struct {
	objc.ID
}

func FileProviderServiceFrom(ptr unsafe.Pointer) FileProviderService {
	return FileProviderService{
		ID: objc.ID(ptr),
	}
}




