// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderService] class.
var fileProviderServiceClass = _FileProviderServiceClass{objc.GetClass("NSFileProviderService")}

type _FileProviderServiceClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderService] class.
type IFileProviderService interface {
	objectivec.IObject
}

// A service that provides a custom communication channel between your app and a File Provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileProviderService

type FileProviderService struct {
	objectivec.Object
}

// FileProviderServiceFrom constructs a [FileProviderService] from an unsafe.Pointer.
//
// A service that provides a custom communication channel between your app and a File Provider extension.
func FileProviderServiceFrom(ptr unsafe.Pointer) FileProviderService {
	return FileProviderService{objectivec.Object{objc.ID(ptr)}}
}



