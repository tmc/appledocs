// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKFrameInfo */


/* debug [class_header]: Header for WKFrameInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FrameInfo */
// An interface definition for the [FrameInfo] class.
type IFrameInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FrameInfo */
	// properties:
	MainFrame() bool
	Request() foundation.URLRequest
	SecurityOrigin() IWKSecurityOrigin
	WebView() IWKWebView
	IsMainFrame() bool
	SetIsMainFrame(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FrameInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FrameInfo */
// Alloc allocates a new instance without initialization.
func (fc _FrameInfoClass) Alloc() FrameInfo {
	rv := objc.Send[FrameInfo](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FrameInfo */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FrameInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FrameInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FrameInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FrameInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FrameInfo */

// A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo/isMainFrame
func (f_ FrameInfo) MainFrame() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("mainFrame"))
	return rv
}/* debug [instance_properties/getter]: mainFrame */


// The frame’s current request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo/request
func (f_ FrameInfo) Request() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](f_.ID, objc.Sel("request"))
	return rv
}/* debug [instance_properties/getter]: request */


// The frame’s security origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo/securityOrigin
func (f_ FrameInfo) SecurityOrigin() IWKSecurityOrigin {
	rv := objc.Send[SecurityOrigin](f_.ID, objc.Sel("securityOrigin"))
	return rv
}/* debug [instance_properties/getter]: securityOrigin */


// The web view that contains this frame and the containing webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFrameInfo/webView
func (f_ FrameInfo) WebView() IWKWebView {
	rv := objc.Send[WebView](f_.ID, objc.Sel("webView"))
	return rv
}/* debug [instance_properties/getter]: webView */


// A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/ismainframe
func (f_ FrameInfo) IsMainFrame() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isMainFrame"))
	return rv
}/* debug [instance_properties/getter]: isMainFrame */


// A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/ismainframe
func (f_ FrameInfo) SetIsMainFrame(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsMainFrame:"), value)
}/* debug [instance_properties/setter]: isMainFrame */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKFrameInfo */



