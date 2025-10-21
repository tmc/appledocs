// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [WebView] class.
var (
	WebViewClass     _WebViewClass
	WebViewClassOnce sync.Once
)

func getWebViewClass() _WebViewClass {
	WebViewClassOnce.Do(func() {
		WebViewClass = _WebViewClass{objc.GetClass("WebView")}
	})
	return WebViewClass
}

type _WebViewClass struct {
	class objc.Class
}

// An interface definition for the [WebView] class.
type IWebView interface {
	appkit.IView
	ChangeFont(sender objc.ID)
	Close()
	ElementAtPoint(point foundation.Point) unsafe.Pointer
	GoBack(sender objc.ID)
	GoForward(sender objc.ID)
	OverWrite(sender objc.ID)
}

// is the core view class in the WebKit framework that manages interactions between the and classes. To embed web content in your application, you just create a object, attach it to a window, and send a message to its main frame.
//
// Behind the scenes, objects encapsulate the content contained in a single frame element. A hierarchy of objects is used to model an entire webpage where the root is called the . There is a object per object used to display the frame content. Therefore, there is a parallel hierarchy of objects used to render an entire page. The object is also the parent view of this hierarchy. You do not need to create and objects directly. These objects are automatically created when the page loads, either programmatically or by the user clicking a link. You customize your embedded web content by implementing delegates to handle certain aspects of the process. objects have multiple delegates because the process of loading a webpage is asynchronous and complicated if errors occur. All the delegates use informal protocols so you only need to implement only the delegates and methods that define the behavior you wish to change—default implementations are already provided. For example, you might want to implement the frame load and resource load delegates to monitor the load progress and display status messages. Applications that use multiple windows may want to implement a user interface delegate. See the individual informal delegate protocols for more details: , , , and . Another way to monitor load progress with less control is to observe the , , and notifications. For example, you could observe these notifications to implement a simple progress indicator in your application. You update the progress indicator by invoking the method to get an estimate of the amount of content that is currently loaded. A object is intended to support most features you would expect in a web browser except that it doesn’t implement the specific user interface for those features. You are responsible for implementing the user interface objects such as status bars, toolbars, buttons, and text fields. For example, a object manages a back-forward list by default, and has and action methods. It is your responsibility to create the buttons that would send theses action messages. Note, there is some overhead in maintaining a back-forward list and page cache, so you should disable it if your application doesn’t use it. You use a object to encapsulate the preferences of a object, such as the font, text encoding, and image settings. You can modify the preferences for individual objects or specify a shared object using the method. Use the method to specify whether the preferences should be automatically saved to the user defaults database. You can also extend WebKit by implementing your own document view and representation classes for specific MIME types. Use the class method to register your custom classes with a object.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class
type WebView struct {
	appkit.View
}

// WebViewFrom constructs a [WebView] from an unsafe.Pointer.
//
// is the core view class in the WebKit framework that manages interactions between the and classes. To embed web content in your application, you just create a object, attach it to a window, and send a message to its main frame.
func WebViewFrom(ptr unsafe.Pointer) WebView {
	return WebView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (wc _WebViewClass) Alloc() WebView {
	rv := objc.Send[WebView](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebViewClass) New() WebView {
	rv := objc.Send[WebView](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebView) Init() WebView {
	rv := objc.Send[WebView](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebView) Autorelease() WebView {
	rv := objc.Send[WebView](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebView creates a new WebView instance.
func NewWebView() WebView {
	return getWebViewClass().New()
}


// Returns whether the receiver can display content of a given MIME type.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canShowMIMEType(_:)
func (wc _WebViewClass) CanShowMIMEType(MIMEType string) bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("canShowMIMEType:"), objc.String(MIMEType))
	return rv
}

// Specifies the view and representation objects to be used for specific MIME types.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/registerClass(_:representationClass:forMIMEType:)
func (wc _WebViewClass) RegisterViewClassRepresentationClassForMIMEType(viewClass objc.Class, representationClass objc.Class, MIMEType string) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("registerViewClass:representationClass:forMIMEType:"), viewClass, representationClass, objc.String(MIMEType))
}

// An action method that changes the font of the selection, or all content if there is no selection.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/changeFont(_:)
func (w_ WebView) ChangeFont(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("changeFont:"), sender)
}

// Closes the web view when it’s no longer needed.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/close()
func (w_ WebView) Close() {
	objc.Send[objc.ID](w_.ID, objc.Sel("close"))
}

// Returns a dictionary description of the element at a given point in the receiver’s coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/element(at:)
func (w_ WebView) ElementAtPoint(point foundation.Point) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("elementAtPoint:"), point)
	return rv
}

// An action method that loads the previous location in the back-forward list.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/goBack(_:)
func (w_ WebView) GoBack(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goBack:"), sender)
}

// An action method that loads the next location in the back-forward list.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/goForward(_:)
func (w_ WebView) GoForward(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goForward:"), sender)
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/overWrite(_:)
func (w_ WebView) OverWrite(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("overWrite:"), sender)
}

// The receiver’s back-forward list.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/backForwardList
func (w_ WebView) BackForwardList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("backForwardList"))
	return rv
}

// An estimate, as a percentage, of the amount of content that is currently loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/estimatedProgress
func (w_ WebView) EstimatedProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("estimatedProgress"))
	return rv
}

// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isEditable
func (w_ WebView) Editable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("editable"))
	return rv
}


// SetEditable sets the value of the editable property.
// A Boolean that indicates whether the user is allowed to edit the document.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isEditable
func (w_ WebView) SetEditable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEditable:"), value)
}

// The identifier of the receiver’s preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/preferencesIdentifier
func (w_ WebView) PreferencesIdentifier() string {
	rv := objc.Send[string](w_.ID, objc.Sel("preferencesIdentifier"))
	return rv
}


// SetPreferencesIdentifier sets the value of the preferencesIdentifier property.
// The identifier of the receiver’s preferences.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/preferencesIdentifier
func (w_ WebView) SetPreferencesIdentifier(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferencesIdentifier:"), objc.String(value))
}




