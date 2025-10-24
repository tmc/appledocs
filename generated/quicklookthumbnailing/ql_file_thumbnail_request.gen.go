// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	MaximumSize() objc.IObject /* cross-framework: Size */
	FileURL() objc.IObject /* cross-framework: URL */
	SetFileURL(value objc.IObject /* cross-framework: URL */)
	MinimumSize() objc.IObject /* cross-framework: Size */
	SetMinimumSize(value objc.IObject /* cross-framework: Size */)
	Scale() float64
	SetScale(value float64)
	// methods:
}

// A request to generate a thumbnail for a custom file type.


// A request to generate a thumbnail for a custom file type.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLFileThumbnailRequest/maximumSize
func (f_ FileThumbnailRequest) MaximumSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](f_.ID, objc.Sel("maximumSize"))
	return rv
}


// The URL of the image file to use for the thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlfilethumbnailrequest/fileurl
func (f_ FileThumbnailRequest) FileURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("fileURL"))
	return rv
}


// The URL of the image file to use for the thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlfilethumbnailrequest/fileurl
func (f_ FileThumbnailRequest) SetFileURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileURL:"), value)
}


// The minimum accepted size of a thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlfilethumbnailrequest/minimumsize
func (f_ FileThumbnailRequest) MinimumSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](f_.ID, objc.Sel("minimumSize"))
	return rv
}


// The minimum accepted size of a thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlfilethumbnailrequest/minimumsize
func (f_ FileThumbnailRequest) SetMinimumSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMinimumSize:"), value)
}


// The scale of the requested thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlfilethumbnailrequest/scale
func (f_ FileThumbnailRequest) Scale() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("scale"))
	return rv
}


// The scale of the requested thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlfilethumbnailrequest/scale
func (f_ FileThumbnailRequest) SetScale(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setScale:"), value)
}



