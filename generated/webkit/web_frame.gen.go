// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebFrame */


/* debug [class_header]: Header for WebFrame */
// The class instance for the [WebFrame] class.
var (
	WebFrameClass     _WebFrameClass
	WebFrameClassOnce sync.Once
)

func getWebFrameClass() _WebFrameClass {
	WebFrameClassOnce.Do(func() {
		WebFrameClass = _WebFrameClass{objc.GetClass("WebFrame")}
	})
	return WebFrameClass
}

type _WebFrameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebFrame */
// An interface definition for the [WebFrame] class.
type IWebFrame interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebFrame */
	// properties:
	ChildFrames() objc.IObject /* cross-framework: NSArray */
	DataSource() IWebDataSource
	DOMDocument() IDOMDocument
	FrameElement() IDOMHTMLElement
	FrameView() IWebFrameView
	GlobalContext() objectivec.IObject
	JavaScriptContext() objc.IObject
	Name() objc.IObject /* cross-framework: NSString */
	ParentFrame() IWebFrame
	ProvisionalDataSource() IWebDataSource
	WebView() IWebView
	WindowObject() IWebScriptObject
	Parent() IWebFrame
	SetParent(value IWebFrame)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebFrame */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebFrame */
// Alloc allocates a new instance without initialization.
func (wc _WebFrameClass) Alloc() WebFrame {
	rv := objc.Send[WebFrame](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebFrameClass) New() WebFrame {
	rv := objc.Send[WebFrame](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebFrame) Init() WebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebFrame) Autorelease() WebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebFrame creates a new WebFrame instance.
func NewWebFrame() WebFrame {
	return getWebFrameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebFrame */
// A object encapsulates the data displayed in a object. There is one object per frame displayed in a . An entire webpage is represented by a hierarchy of objects in which the root object is called the .
//
// Each also has a object that manages the loading of frame content. You use the method to initiate an asynchronous client request which will create a provisional data source. The provisional data source will transition to a committed data source once any data has been received. There are some special, predefined, frame names that you can use when referring to or finding a . Some of the predefined frame names are: “_self”, “_current”, “_parent”, and “_top.” See for a description of their meaning. Frame names may also be specified in the HTML source, or set by clients. However, the group name is an arbitrary identifier used to group related frames. For example, JavaScript running in a frame can access any other frame in the same group. It’s up to the application how it chooses to scope related frames.


// A object encapsulates the data displayed in a object. There is one object per frame displayed in a . An entire webpage is represented by a hierarchy of objects in which the root object is called the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame
type WebFrame struct {
	objectivec.Object
}

// WebFrameFrom constructs a [WebFrame] from an unsafe.Pointer.
//
// A object encapsulates the data displayed in a object. There is one object per frame displayed in a . An entire webpage is represented by a hierarchy of objects in which the root object is called the .
func WebFrameFrom(ptr unsafe.Pointer) WebFrame {
	return WebFrame{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebFrame */

// Initializes the receiver with a frame name, web frame view, and controlling web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/init(name:webFrameView:webView:)
func NewWebFrameWithNameWebFrameViewWebView(name objc.IObject /* cross-framework: NSString */, view IWebFrameView, webView IWebView) WebFrame {
	instance := getWebFrameClass().Alloc()
	rv := objc.Send[WebFrame](instance.ID, objc.Sel("initWithName:webFrameView:webView:"), name, view, webView)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebFrameWithNameWebFrameViewWebView */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebFrame */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebFrame */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebFrame */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebFrame */

// The frames of the web frame’s immediate children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/childFrames
func (w_ WebFrame) ChildFrames() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](w_.ID, objc.Sel("childFrames"))
	return rv
}/* debug [instance_properties/getter]: childFrames */


// The committed data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/dataSource
func (w_ WebFrame) DataSource() IWebDataSource {
	rv := objc.Send[WebDataSource](w_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// The web frame’s DOM document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/domDocument
func (w_ WebFrame) DOMDocument() IDOMDocument {
	rv := objc.Send[DOMDocument](w_.ID, objc.Sel("DOMDocument"))
	return rv
}/* debug [instance_properties/getter]: DOMDocument */


// The web view’s DOM frame element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/frameElement
func (w_ WebFrame) FrameElement() IDOMHTMLElement {
	rv := objc.Send[DOMHTMLElement](w_.ID, objc.Sel("frameElement"))
	return rv
}/* debug [instance_properties/getter]: frameElement */


// The web frame’s view object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/frameView
func (w_ WebFrame) FrameView() IWebFrameView {
	rv := objc.Send[WebFrameView](w_.ID, objc.Sel("frameView"))
	return rv
}/* debug [instance_properties/getter]: frameView */


// The global JavaScript execution context for bridging between the WebKit and JavaScriptCore C API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/globalContext
func (w_ WebFrame) GlobalContext() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("globalContext"))
	return rv
}/* debug [instance_properties/getter]: globalContext */


// The frame’s global JavaScript execution context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/javaScriptContext
func (w_ WebFrame) JavaScriptContext() objc.IObject {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("javaScriptContext"))
	return rv
}/* debug [instance_properties/getter]: javaScriptContext */


// The web frame’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/name
func (w_ WebFrame) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The web frame’s parent web frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/parent
func (w_ WebFrame) ParentFrame() IWebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("parentFrame"))
	return rv
}/* debug [instance_properties/getter]: parentFrame */


// The provisional data source, or if either a load request is not in progress or a load request has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/provisionalDataSource
func (w_ WebFrame) ProvisionalDataSource() IWebDataSource {
	rv := objc.Send[WebDataSource](w_.ID, objc.Sel("provisionalDataSource"))
	return rv
}/* debug [instance_properties/getter]: provisionalDataSource */


// The view object that manages the web frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/webView
func (w_ WebFrame) WebView() IWebView {
	rv := objc.Send[WebView](w_.ID, objc.Sel("webView"))
	return rv
}/* debug [instance_properties/getter]: webView */


// The JavaScript window object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/windowObject
func (w_ WebFrame) WindowObject() IWebScriptObject {
	rv := objc.Send[WebScriptObject](w_.ID, objc.Sel("windowObject"))
	return rv
}/* debug [instance_properties/getter]: windowObject */


// The web frame’s parent web frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/parent
func (w_ WebFrame) Parent() IWebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// The web frame’s parent web frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/parent
func (w_ WebFrame) SetParent(value IWebFrame) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WebFrame */


