// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontAssetRequest] class.
var (
	FontAssetRequestClass     _FontAssetRequestClass
	FontAssetRequestClassOnce sync.Once
)

func getFontAssetRequestClass() _FontAssetRequestClass {
	FontAssetRequestClassOnce.Do(func() {
		FontAssetRequestClass = _FontAssetRequestClass{objc.GetClass("NSFontAssetRequest")}
	})
	return FontAssetRequestClass
}

type _FontAssetRequestClass struct {
	class objc.Class
}

// An interface definition for the [FontAssetRequest] class.
type IFontAssetRequest interface {
	objectivec.IObject
	// properties:
	DownloadedFontDescriptors() []FontDescriptor /* primitive/slice/pointer. */
	Progress() Progress /* not a class type */
	// methods:
	DownloadFontAssetsWithCompletionHandler(completionHandler unsafe.Pointer)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest
type FontAssetRequest struct {
	objectivec.Object
}

// FontAssetRequestFrom constructs a [FontAssetRequest] from an unsafe.Pointer.
func FontAssetRequestFrom(ptr unsafe.Pointer) FontAssetRequest {
	return FontAssetRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FontAssetRequestClass) Alloc() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FontAssetRequestClass) New() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontAssetRequest) Init() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontAssetRequest) Autorelease() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontAssetRequest creates a new FontAssetRequest instance.
func NewFontAssetRequest() FontAssetRequest {
	return getFontAssetRequestClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/init(fontDescriptors:options:)
func NewFontAssetRequestWithFontDescriptorsOptions(fontDescriptors []FontDescriptor /* primitive/slice/pointer. */, options FontAssetRequestOptions) FontAssetRequest {
	instance := getFontAssetRequestClass().Alloc()
	rv := objc.Send[FontAssetRequest](instance.ID, objc.Sel("initWithFontDescriptors:options:"), fontDescriptors, options)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/download(withCompletionHandler:)
func (f_ FontAssetRequest) DownloadFontAssetsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("downloadFontAssetsWithCompletionHandler:"), completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/downloadedFontDescriptors
func (f_ FontAssetRequest) DownloadedFontDescriptors() []FontDescriptor /* primitive/slice/pointer. */ {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("downloadedFontDescriptors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/progress
func (f_ FontAssetRequest) Progress() Progress /* not a class type */ {
	rv := objc.Send[Progress](f_.ID, objc.Sel("progress"))
	return rv
}


