// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [FileThumbnailRequest] class.
var (
	FileThumbnailRequestClass     _FileThumbnailRequestClass
	FileThumbnailRequestClassOnce sync.Once
)

func getFileThumbnailRequestClass() _FileThumbnailRequestClass {
	FileThumbnailRequestClassOnce.Do(func() {
		FileThumbnailRequestClass = _FileThumbnailRequestClass{objc.GetClass("QLFileThumbnailRequest")}
	})
	return FileThumbnailRequestClass
}

type _FileThumbnailRequestClass struct {
	class objc.Class
}

// An interface definition for the [FileThumbnailRequest] class.
type IFileThumbnailRequest interface {
	objectivec.IObject
}

// A request to generate a thumbnail for a custom file type.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLFileThumbnailRequest
type FileThumbnailRequest struct {
	objectivec.Object
}

// FileThumbnailRequestFrom constructs a [FileThumbnailRequest] from an unsafe.Pointer.
//
// A request to generate a thumbnail for a custom file type.
func FileThumbnailRequestFrom(ptr unsafe.Pointer) FileThumbnailRequest {
	return FileThumbnailRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileThumbnailRequestClass) Alloc() FileThumbnailRequest {
	rv := objc.Send[FileThumbnailRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileThumbnailRequestClass) New() FileThumbnailRequest {
	rv := objc.Send[FileThumbnailRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileThumbnailRequest) Init() FileThumbnailRequest {
	rv := objc.Send[FileThumbnailRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileThumbnailRequest) Autorelease() FileThumbnailRequest {
	rv := objc.Send[FileThumbnailRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileThumbnailRequest creates a new FileThumbnailRequest instance.
func NewFileThumbnailRequest() FileThumbnailRequest {
	return getFileThumbnailRequestClass().New()
}


// The maximum accepted size of a thumbnail.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLFileThumbnailRequest/maximumSize
func (f_ FileThumbnailRequest) MaximumSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](f_.ID, objc.Sel("maximumSize"))
	return rv
}



