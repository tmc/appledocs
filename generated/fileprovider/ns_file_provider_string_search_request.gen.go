// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderStringSearchRequest] class.
var (
	FileProviderStringSearchRequestClass     _FileProviderStringSearchRequestClass
	FileProviderStringSearchRequestClassOnce sync.Once
)

func getFileProviderStringSearchRequestClass() _FileProviderStringSearchRequestClass {
	FileProviderStringSearchRequestClassOnce.Do(func() {
		FileProviderStringSearchRequestClass = _FileProviderStringSearchRequestClass{objc.GetClass("NSFileProviderStringSearchRequest")}
	})
	return FileProviderStringSearchRequestClass
}

type _FileProviderStringSearchRequestClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderStringSearchRequest] class.
type IFileProviderStringSearchRequest interface {
	objectivec.IObject
}

// A type that contains details of a string-based search request.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderStringSearchRequest
type FileProviderStringSearchRequest struct {
	objectivec.Object
}

// FileProviderStringSearchRequestFrom constructs a [FileProviderStringSearchRequest] from an unsafe.Pointer.
//
// A type that contains details of a string-based search request.
func FileProviderStringSearchRequestFrom(ptr unsafe.Pointer) FileProviderStringSearchRequest {
	return FileProviderStringSearchRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderStringSearchRequestClass) Alloc() FileProviderStringSearchRequest {
	rv := objc.Send[FileProviderStringSearchRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderStringSearchRequestClass) New() FileProviderStringSearchRequest {
	rv := objc.Send[FileProviderStringSearchRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderStringSearchRequest) Init() FileProviderStringSearchRequest {
	rv := objc.Send[FileProviderStringSearchRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderStringSearchRequest) Autorelease() FileProviderStringSearchRequest {
	rv := objc.Send[FileProviderStringSearchRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderStringSearchRequest creates a new FileProviderStringSearchRequest instance.
func NewFileProviderStringSearchRequest() FileProviderStringSearchRequest {
	return getFileProviderStringSearchRequestClass().New()
}




