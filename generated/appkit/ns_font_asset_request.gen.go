// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	DownloadedFontDescriptors() NSFontDescriptor
	SetDownloadedFontDescriptors(value IFontDescriptor)
	Progress() foundation.Progress
	SetProgress(value foundation.IProgress)
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
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest/downloadedfontdescriptors

func (f_ FontAssetRequest) DownloadedFontDescriptors() NSFontDescriptor {
	rv := objc.Send[NSFontDescriptor](f_.ID, objc.Sel("downloadedFontDescriptors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest/downloadedfontdescriptors

func (f_ FontAssetRequest) SetDownloadedFontDescriptors(value IFontDescriptor) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDownloadedFontDescriptors:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest/progress

func (f_ FontAssetRequest) Progress() foundation.Progress {
	rv := objc.Send[foundation.Progress](f_.ID, objc.Sel("progress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest/progress

func (f_ FontAssetRequest) SetProgress(value foundation.IProgress) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setProgress:"), value)
}



