// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FrameInfo] class.
var (
	FrameInfoClass     _FrameInfoClass
	FrameInfoClassOnce sync.Once
)

func getFrameInfoClass() _FrameInfoClass {
	FrameInfoClassOnce.Do(func() {
		FrameInfoClass = _FrameInfoClass{objc.GetClass("WKFrameInfo")}
	})
	return FrameInfoClass
}

type _FrameInfoClass struct {
	class objc.Class
}

// An interface definition for the [FrameInfo] class.
type IFrameInfo interface {
	objectivec.IObject
	// properties:
	MainFrame() bool
	IsMainFrame() bool
	SetIsMainFrame(value bool)
	Request() objc.IObject /* cross-framework: URLRequest */
	SetRequest(value objc.IObject /* cross-framework: URLRequest */)
	SecurityOrigin() IWKSecurityOrigin
	SetSecurityOrigin(value IWKSecurityOrigin)
	WebView() IWKWebView
	SetWebView(value IWKWebView)
	// methods:
}

// An object that contains information about a frame on a webpage.
//
// An instance of this class is a transient, data-only object; it does not uniquely identify a frame across multiple delegate method calls.


// An object that contains information about a frame on a webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo
type FrameInfo struct {
	objectivec.Object
}

// FrameInfoFrom constructs a [FrameInfo] from an unsafe.Pointer.
//
// An object that contains information about a frame on a webpage.
func FrameInfoFrom(ptr unsafe.Pointer) FrameInfo {
	return FrameInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FrameInfoClass) Alloc() FrameInfo {
	rv := objc.Send[FrameInfo](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FrameInfoClass) New() FrameInfo {
	rv := objc.Send[FrameInfo](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FrameInfo) Init() FrameInfo {
	rv := objc.Send[FrameInfo](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FrameInfo) Autorelease() FrameInfo {
	rv := objc.Send[FrameInfo](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFrameInfo creates a new FrameInfo instance.
func NewFrameInfo() FrameInfo {
	return getFrameInfoClass().New()
}



// A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo/isMainFrame
func (f_ FrameInfo) MainFrame() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("mainFrame"))
	return rv
}


// A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/ismainframe
func (f_ FrameInfo) IsMainFrame() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isMainFrame"))
	return rv
}


// A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/ismainframe
func (f_ FrameInfo) SetIsMainFrame(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsMainFrame:"), value)
}


// The frame’s current request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/request
func (f_ FrameInfo) Request() objc.IObject /* cross-framework: URLRequest */ {
	rv := objc.Send[foundation.URLRequest](f_.ID, objc.Sel("request"))
	return rv
}


// The frame’s current request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/request
func (f_ FrameInfo) SetRequest(value objc.IObject /* cross-framework: URLRequest */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRequest:"), value)
}


// The frame’s security origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/securityorigin
func (f_ FrameInfo) SecurityOrigin() IWKSecurityOrigin {
	rv := objc.Send[SecurityOrigin](f_.ID, objc.Sel("securityOrigin"))
	return rv
}


// The frame’s security origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/securityorigin
func (f_ FrameInfo) SetSecurityOrigin(value IWKSecurityOrigin) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSecurityOrigin:"), value)
}


// The web view that contains this frame and the containing webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/webview
func (f_ FrameInfo) WebView() IWKWebView {
	rv := objc.Send[WebView](f_.ID, objc.Sel("webView"))
	return rv
}


// The web view that contains this frame and the containing webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/webview
func (f_ FrameInfo) SetWebView(value IWKWebView) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWebView:"), value)
}



