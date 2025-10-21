// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	LoadRequest(request unsafe.Pointer)
	LoadArchive(archive unsafe.Pointer)
	LoadDataMIMETypeTextEncodingNameBaseURL(data unsafe.Pointer, MIMEType string, encodingName string, URL unsafe.Pointer)
	LoadAlternateHTMLStringBaseURLForUnreachableURL(string_ string, baseURL unsafe.Pointer, unreachableURL unsafe.Pointer)
	LoadHTMLStringBaseURL(string_ string, URL unsafe.Pointer)
	Reload()
	ReloadFromOrigin()
	StopLoading()
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
func (w_ WebFrame) LoadRequest(request unsafe.Pointer) {
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
func (w_ WebFrame) LoadDataMIMETypeTextEncodingNameBaseURL(data unsafe.Pointer, MIMEType string, encodingName string, URL unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadData:MIMEType:textEncodingName:baseURL:"), data, objc.String(MIMEType), objc.String(encodingName), URL)
}

// Loads alternate content for a frame whose URL is unreachable.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/loadAlternateHTMLString(_:baseURL:forUnreachableURL:)
func (w_ WebFrame) LoadAlternateHTMLStringBaseURLForUnreachableURL(string_ string, baseURL unsafe.Pointer, unreachableURL unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadAlternateHTMLString:baseURL:forUnreachableURL:"), objc.String(string_), baseURL, unreachableURL)
}

// Sets the main page contents and base URL.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrame/loadHTMLString(_:baseURL:)
func (w_ WebFrame) LoadHTMLStringBaseURL(string_ string, URL unsafe.Pointer) {
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



