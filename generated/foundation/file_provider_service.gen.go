// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderService] class.
var (
	fileProviderServiceClass     _FileProviderServiceClass
	fileProviderServiceClassOnce sync.Once
)

func getFileProviderServiceClass() _FileProviderServiceClass {
	fileProviderServiceClassOnce.Do(func() {
		fileProviderServiceClass = _FileProviderServiceClass{objc.GetClass("NSFileProviderService")}
	})
	return fileProviderServiceClass
}

type _FileProviderServiceClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderService] class.
type IFileProviderService interface {
	objectivec.IObject
}

// A service that provides a custom communication channel between your app and a File Provider extension.
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

// Alloc allocates a new instance without initialization.
func (fc _FileProviderServiceClass) Alloc() FileProviderService {
	rv := objc.Send[FileProviderService](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderServiceClass) New() FileProviderService {
	rv := objc.Send[FileProviderService](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderService) Init() FileProviderService {
	rv := objc.Send[FileProviderService](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderService) Autorelease() FileProviderService {
	rv := objc.Send[FileProviderService](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderService creates a new FileProviderService instance.
func NewFileProviderService() FileProviderService {
	return getFileProviderServiceClass().New()
}




