// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderRequest] class.
var (
	FileProviderRequestClass     _FileProviderRequestClass
	FileProviderRequestClassOnce sync.Once
)

func getFileProviderRequestClass() _FileProviderRequestClass {
	FileProviderRequestClassOnce.Do(func() {
		FileProviderRequestClass = _FileProviderRequestClass{objc.GetClass("NSFileProviderRequest")}
	})
	return FileProviderRequestClass
}

type _FileProviderRequestClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderRequest] class.
type IFileProviderRequest interface {
	objectivec.IObject
	// properties:
	DomainVersion() IFileProviderDomainVersion
	SetDomainVersion(value IFileProviderDomainVersion)
	IsFileViewerRequest() bool
	SetIsFileViewerRequest(value bool)
	IsSystemRequest() bool
	SetIsSystemRequest(value bool)
	RequestingExecutable() objc.IObject /* cross-framework: URL */
	SetRequestingExecutable(value objc.IObject /* cross-framework: URL */)
	// methods:
}

// An object that provides information about the application requesting data from the File Provider extension.


// An object that provides information about the application requesting data from the File Provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest
type FileProviderRequest struct {
	objectivec.Object
}

// FileProviderRequestFrom constructs a [FileProviderRequest] from an unsafe.Pointer.
//
// An object that provides information about the application requesting data from the File Provider extension.
func FileProviderRequestFrom(ptr unsafe.Pointer) FileProviderRequest {
	return FileProviderRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderRequestClass) Alloc() FileProviderRequest {
	rv := objc.Send[FileProviderRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderRequestClass) New() FileProviderRequest {
	rv := objc.Send[FileProviderRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderRequest) Init() FileProviderRequest {
	rv := objc.Send[FileProviderRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderRequest) Autorelease() FileProviderRequest {
	rv := objc.Send[FileProviderRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderRequest creates a new FileProviderRequest instance.
func NewFileProviderRequest() FileProviderRequest {
	return getFileProviderRequestClass().New()
}



// The version of the domain for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/domainversion
func (f_ FileProviderRequest) DomainVersion() IFileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](f_.ID, objc.Sel("domainVersion"))
	return rv
}


// The version of the domain for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/domainversion
func (f_ FileProviderRequest) SetDomainVersion(value IFileProviderDomainVersion) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDomainVersion:"), value)
}


// A Boolean value that indicates whether the request came from Finder or related system file browsers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/isfileviewerrequest
func (f_ FileProviderRequest) IsFileViewerRequest() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isFileViewerRequest"))
	return rv
}


// A Boolean value that indicates whether the request came from Finder or related system file browsers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/isfileviewerrequest
func (f_ FileProviderRequest) SetIsFileViewerRequest(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsFileViewerRequest:"), value)
}


// A Boolean value that indicates whether the request came from a system process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/issystemrequest
func (f_ FileProviderRequest) IsSystemRequest() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isSystemRequest"))
	return rv
}


// A Boolean value that indicates whether the request came from a system process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/issystemrequest
func (f_ FileProviderRequest) SetIsSystemRequest(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsSystemRequest:"), value)
}


// The URL of the requesting executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/requestingexecutable
func (f_ FileProviderRequest) RequestingExecutable() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("requestingExecutable"))
	return rv
}


// The URL of the requesting executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderrequest/requestingexecutable
func (f_ FileProviderRequest) SetRequestingExecutable(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRequestingExecutable:"), value)
}



