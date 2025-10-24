// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKDownload */


/* debug [class_header]: Header for WKDownload */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Download */
// An interface definition for the [Download] class.
type IDownload interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Download */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	UserInitiated() bool
	OriginalRequest() foundation.URLRequest
	OriginatingFrame() IWKFrameInfo
	WebView() IWKWebView
	IsUserInitiated() bool
	SetIsUserInitiated(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Download */
	// methods:
	Cancel(completionHandler func(unsafe.Pointer))
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Download */
// Alloc allocates a new instance without initialization.
func (dc _DownloadClass) Alloc() Download {
	rv := objc.Send[Download](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Download */
// An object that represents the download of a web resource.


// An object that represents the download of a web resource.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Download *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Download */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Download */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Download */

// Cancels the download, and optionally captures data so that you can resume the download later.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/cancel(_:)
func (d_ Download) Cancel(completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](d_.ID, objc.Sel("cancel:"), completionHandler)
}/* debug [instance_methods/method]: Cancel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Download */

// An object you use to track download progress and handle redirects, authentication challenges, and failures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/delegate
func (d_ Download) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// An object you use to track download progress and handle redirects, authentication challenges, and failures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/delegate
func (d_ Download) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/isUserInitiated
func (d_ Download) UserInitiated() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("userInitiated"))
	return rv
}/* debug [instance_properties/getter]: userInitiated */


// An object that represents the request that initiated the download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/originalRequest
func (d_ Download) OriginalRequest() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](d_.ID, objc.Sel("originalRequest"))
	return rv
}/* debug [instance_properties/getter]: originalRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/originatingFrame
func (d_ Download) OriginatingFrame() IWKFrameInfo {
	rv := objc.Send[FrameInfo](d_.ID, objc.Sel("originatingFrame"))
	return rv
}/* debug [instance_properties/getter]: originatingFrame */


// The web view where the download initiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKDownload/webView
func (d_ Download) WebView() IWKWebView {
	rv := objc.Send[WebView](d_.ID, objc.Sel("webView"))
	return rv
}/* debug [instance_properties/getter]: webView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/isuserinitiated
func (d_ Download) IsUserInitiated() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isUserInitiated"))
	return rv
}/* debug [instance_properties/getter]: isUserInitiated */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/isuserinitiated
func (d_ Download) SetIsUserInitiated(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsUserInitiated:"), value)
}/* debug [instance_properties/setter]: isUserInitiated */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKDownload */



