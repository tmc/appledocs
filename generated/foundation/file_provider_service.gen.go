// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FileProviderService] class.
var FileProviderServiceClass = _FileProviderServiceClass{objc.GetClass("NSFileProviderService")}

type _FileProviderServiceClass struct {
	class objc.Class
}

type FileProviderService struct {
	objc.ID
}

func FileProviderServiceFrom(ptr unsafe.Pointer) FileProviderService {
	return FileProviderService{
		ID: objc.ID(ptr),
	}
}




