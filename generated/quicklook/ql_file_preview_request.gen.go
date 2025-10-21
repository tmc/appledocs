// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [FilePreviewRequest] class.
var (
	FilePreviewRequestClass     _FilePreviewRequestClass
	FilePreviewRequestClassOnce sync.Once
)

func getFilePreviewRequestClass() _FilePreviewRequestClass {
	FilePreviewRequestClassOnce.Do(func() {
		FilePreviewRequestClass = _FilePreviewRequestClass{objc.GetClass("QLFilePreviewRequest")}
	})
	return FilePreviewRequestClass
}

type _FilePreviewRequestClass struct {
	class objc.Class
}

// An interface definition for the [FilePreviewRequest] class.
type IFilePreviewRequest interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLFilePreviewRequest
type FilePreviewRequest struct {
	objectivec.Object
}

// FilePreviewRequestFrom constructs a [FilePreviewRequest] from an unsafe.Pointer.
func FilePreviewRequestFrom(ptr unsafe.Pointer) FilePreviewRequest {
	return FilePreviewRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilePreviewRequestClass) Alloc() FilePreviewRequest {
	rv := objc.Send[FilePreviewRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilePreviewRequestClass) New() FilePreviewRequest {
	rv := objc.Send[FilePreviewRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilePreviewRequest) Init() FilePreviewRequest {
	rv := objc.Send[FilePreviewRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilePreviewRequest) Autorelease() FilePreviewRequest {
	rv := objc.Send[FilePreviewRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilePreviewRequest creates a new FilePreviewRequest instance.
func NewFilePreviewRequest() FilePreviewRequest {
	return getFilePreviewRequestClass().New()
}




