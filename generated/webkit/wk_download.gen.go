// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Download] class.
var (
	DownloadClass     _DownloadClass
	DownloadClassOnce sync.Once
)

func getDownloadClass() _DownloadClass {
	DownloadClassOnce.Do(func() {
		DownloadClass = _DownloadClass{objc.GetClass("WKDownload")}
	})
	return DownloadClass
}

type _DownloadClass struct {
	class objc.Class
}

// An interface definition for the [Download] class.
type IDownload interface {
	objectivec.IObject
	Cancel(completionHandler unsafe.Pointer)
}

// An object that represents the download of a web resource.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload
type Download struct {
	objectivec.Object
}

// DownloadFrom constructs a [Download] from an unsafe.Pointer.
//
// An object that represents the download of a web resource.
func DownloadFrom(ptr unsafe.Pointer) Download {
	return Download{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DownloadClass) Alloc() Download {
	rv := objc.Send[Download](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DownloadClass) New() Download {
	rv := objc.Send[Download](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Download) Init() Download {
	rv := objc.Send[Download](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Download) Autorelease() Download {
	rv := objc.Send[Download](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDownload creates a new Download instance.
func NewDownload() Download {
	return getDownloadClass().New()
}


// Cancels the download, and optionally captures data so that you can resume the download later.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/cancel(_:)
func (d_ Download) Cancel(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("cancel:"), completionHandler)
}

// An object you use to track download progress and handle redirects, authentication challenges, and failures.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/delegate
func (d_ Download) Delegate() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object you use to track download progress and handle redirects, authentication challenges, and failures.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/delegate
func (d_ Download) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}

// The web view where the download initiated.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/webView
func (d_ Download) WebView() WKWebView {
	rv := objc.Send[WKWebView](d_.ID, objc.Sel("webView"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/isuserinitiated
func (d_ Download) IsUserInitiated() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isUserInitiated"))
	return rv
}


// SetIsUserInitiated sets the value of the isUserInitiated property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/isuserinitiated
func (d_ Download) SetIsUserInitiated(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsUserInitiated:"), value)
}

// An object that represents the request that initiated the download.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/originalrequest
func (d_ Download) OriginalRequest() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](d_.ID, objc.Sel("originalRequest"))
	return rv
}


// SetOriginalRequest sets the value of the originalRequest property.
// An object that represents the request that initiated the download.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/originalrequest
func (d_ Download) SetOriginalRequest(value foundation.IURLRequest) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOriginalRequest:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/originatingframe
func (d_ Download) OriginatingFrame() WKFrameInfo {
	rv := objc.Send[WKFrameInfo](d_.ID, objc.Sel("originatingFrame"))
	return rv
}


// SetOriginatingFrame sets the value of the originatingFrame property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/originatingframe
func (d_ Download) SetOriginatingFrame(value IWKFrameInfo) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOriginatingFrame:"), value)
}



