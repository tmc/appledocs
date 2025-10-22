// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	FileURL() foundation.URL
}

// A Quick Look preview request that indicates the content to preview.
//
// The system provides a to the method of your data-based Quick Look extension.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLFilePreviewRequest
type FilePreviewRequest struct {
	objectivec.Object
}

// FilePreviewRequestFrom constructs a [FilePreviewRequest] from an unsafe.Pointer.
//
// A Quick Look preview request that indicates the content to preview.
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


// The URL that indicates the content to preview.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLFilePreviewRequest/fileURL
func (f_ FilePreviewRequest) FileURL() foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("fileURL"))
	return rv
}



