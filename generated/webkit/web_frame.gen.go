// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/javascriptcore"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [WebFrame] class.
type IWebFrame interface {
	objectivec.IObject
	LoadRequest(request foundation.IURLRequest)
	LoadArchive(archive unsafe.Pointer)
	LoadDataMIMETypeTextEncodingNameBaseURL(data foundation.IData, MIMEType string, encodingName string, URL foundation.IURL)
	LoadAlternateHTMLStringBaseURLForUnreachableURL(string_ string, baseURL foundation.IURL, unreachableURL foundation.IURL)
	LoadHTMLStringBaseURL(string_ string, URL foundation.IURL)
	Reload()
	ReloadFromOrigin()
	StopLoading()
	DataSource() unsafe.Pointer
	ChildFrames() unsafe.Pointer
	SetChildFrames(value unsafe.Pointer)
	DomDocument() unsafe.Pointer
	SetDomDocument(value unsafe.Pointer)
	FrameElement() unsafe.Pointer
	SetFrameElement(value unsafe.Pointer)
	FrameView() unsafe.Pointer
	SetFrameView(value unsafe.Pointer)
	GlobalContext() unsafe.Pointer
	SetGlobalContext(value unsafe.Pointer)
	JavaScriptContext() javascriptcore.JSContext
	SetJavaScriptContext(value javascriptcore.IJSContext)
	Name() string
	SetName(value string)
	Parent() WebFrame
	SetParent(value IWebFrame)
	ProvisionalDataSource() unsafe.Pointer
	SetProvisionalDataSource(value unsafe.Pointer)
	WebView() WebView
	SetWebView(value IWebView)
	WindowObject() unsafe.Pointer
	SetWindowObject(value unsafe.Pointer)
}

// A object encapsulates the data displayed in a object. There is one object per frame displayed in a . An entire webpage is represented by a hierarchy of objects in which the root object is called the .
//
// Each also has a object that manages the loading of frame content. You use the method to initiate an asynchronous client request which will create a provisional data source. The provisional data source will transition to a committed data source once any data has been received. There are some special, predefined, frame names that you can use when referring to or finding a . Some of the predefined frame names are: “_self”, “_current”, “_parent”, and “_top.” See for a description of their meaning. Frame names may also be specified in the HTML source, or set by clients. However, the group name is an arbitrary identifier used to group related frames. For example, JavaScript running in a frame can access any other frame in the same group. It’s up to the application how it chooses to scope related frames.
//
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

// Alloc allocates a new instance without initialization.
func (wc _WebFrameClass) Alloc() WebFrame {
	rv := objc.Send[WebFrame](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Connects to a given URL by initiating an asynchronous client request.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/load(_:)-47p2s
func (w_ WebFrame) LoadRequest(request foundation.IURLRequest) {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadRequest:"), request)
}

// Loads an archive into the web frame.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/load(_:)-6wkx6
func (w_ WebFrame) LoadArchive(archive unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadArchive:"), archive)
}

// Sets the main page contents, MIME type, content encoding, and base URL.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/load(_:mimeType:textEncodingName:baseURL:)
func (w_ WebFrame) LoadDataMIMETypeTextEncodingNameBaseURL(data foundation.IData, MIMEType string, encodingName string, URL foundation.IURL) {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadData:MIMEType:textEncodingName:baseURL:"), data, objc.String(MIMEType), objc.String(encodingName), URL)
}

// Loads alternate content for a frame whose URL is unreachable.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/loadAlternateHTMLString(_:baseURL:forUnreachableURL:)
func (w_ WebFrame) LoadAlternateHTMLStringBaseURLForUnreachableURL(string_ string, baseURL foundation.IURL, unreachableURL foundation.IURL) {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadAlternateHTMLString:baseURL:forUnreachableURL:"), objc.String(string_), baseURL, unreachableURL)
}

// Sets the main page contents and base URL.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/loadHTMLString(_:baseURL:)
func (w_ WebFrame) LoadHTMLStringBaseURL(string_ string, URL foundation.IURL) {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadHTMLString:baseURL:"), objc.String(string_), URL)
}

// Reloads the initial request passed as an argument to .
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/reload()
func (w_ WebFrame) Reload() {
	objc.Send[objc.ID](w_.ID, objc.Sel("reload"))
}

// Performs an end-to-end revalidation using cache-validating conditionals if possible.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/reloadFromOrigin()
func (w_ WebFrame) ReloadFromOrigin() {
	objc.Send[objc.ID](w_.ID, objc.Sel("reloadFromOrigin"))
}

// Stops any pending loads on the receiver’s data source, and those of its children.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/stopLoading()
func (w_ WebFrame) StopLoading() {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopLoading"))
}

// The committed data source.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/dataSource
func (w_ WebFrame) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("dataSource"))
	return rv
}

// The frames of the web frame’s immediate children.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/childframes
func (w_ WebFrame) ChildFrames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("childFrames"))
	return rv
}


// SetChildFrames sets the value of the childFrames property.
// The frames of the web frame’s immediate children.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/childframes
func (w_ WebFrame) SetChildFrames(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setChildFrames:"), value)
}

// The web frame’s DOM document.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/domdocument
func (w_ WebFrame) DomDocument() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("domDocument"))
	return rv
}


// SetDomDocument sets the value of the domDocument property.
// The web frame’s DOM document.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/domdocument
func (w_ WebFrame) SetDomDocument(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDomDocument:"), value)
}

// The web view’s DOM frame element.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/frameelement
func (w_ WebFrame) FrameElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("frameElement"))
	return rv
}


// SetFrameElement sets the value of the frameElement property.
// The web view’s DOM frame element.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/frameelement
func (w_ WebFrame) SetFrameElement(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameElement:"), value)
}

// The web frame’s view object.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/frameview
func (w_ WebFrame) FrameView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("frameView"))
	return rv
}


// SetFrameView sets the value of the frameView property.
// The web frame’s view object.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/frameview
func (w_ WebFrame) SetFrameView(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameView:"), value)
}

// The global JavaScript execution context for bridging between the WebKit and JavaScriptCore C API.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/globalcontext
func (w_ WebFrame) GlobalContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("globalContext"))
	return rv
}


// SetGlobalContext sets the value of the globalContext property.
// The global JavaScript execution context for bridging between the WebKit and JavaScriptCore C API.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/globalcontext
func (w_ WebFrame) SetGlobalContext(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGlobalContext:"), value)
}

// The frame’s global JavaScript execution context.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/javascriptcontext
func (w_ WebFrame) JavaScriptContext() javascriptcore.JSContext {
	rv := objc.Send[javascriptcore.JSContext](w_.ID, objc.Sel("javaScriptContext"))
	return rv
}


// SetJavaScriptContext sets the value of the javaScriptContext property.
// The frame’s global JavaScript execution context.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/javascriptcontext
func (w_ WebFrame) SetJavaScriptContext(value javascriptcore.IJSContext) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setJavaScriptContext:"), value)
}

// The web frame’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/name
func (w_ WebFrame) Name() string {
	rv := objc.Send[string](w_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The web frame’s name.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/name
func (w_ WebFrame) SetName(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setName:"), objc.String(value))
}

// The web frame’s parent web frame.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/parent
func (w_ WebFrame) Parent() WebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
// The web frame’s parent web frame.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/parent
func (w_ WebFrame) SetParent(value IWebFrame) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setParent:"), value)
}

// The provisional data source, or
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/provisionaldatasource
func (w_ WebFrame) ProvisionalDataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("provisionalDataSource"))
	return rv
}


// SetProvisionalDataSource sets the value of the provisionalDataSource property.
// The provisional data source, or

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/provisionaldatasource
func (w_ WebFrame) SetProvisionalDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setProvisionalDataSource:"), value)
}

// The view object that manages the web frame.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/webview
func (w_ WebFrame) WebView() WebView {
	rv := objc.Send[WebView](w_.ID, objc.Sel("webView"))
	return rv
}


// SetWebView sets the value of the webView property.
// The view object that manages the web frame.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/webview
func (w_ WebFrame) SetWebView(value IWebView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebView:"), value)
}

// The JavaScript window object.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/windowobject
func (w_ WebFrame) WindowObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("windowObject"))
	return rv
}


// SetWindowObject sets the value of the windowObject property.
// The JavaScript window object.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe/windowobject
func (w_ WebFrame) SetWindowObject(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowObject:"), value)
}



