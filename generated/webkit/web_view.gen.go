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

// A Boolean that indicates whether or not the receiver’s attributes are automatically stored in the user defaults database.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/autosaves
func (w_ WebView) Autosaves() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autosaves"))
	return rv
}


// SetAutosaves sets the value of the autosaves property.
// A Boolean that indicates whether or not the receiver’s attributes are automatically stored in the user defaults database.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/autosaves
func (w_ WebView) SetAutosaves(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutosaves:"), value)
}

// The receiver’s CSS media property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mediastyle
func (w_ WebView) MediaStyle() string {
	rv := objc.Send[string](w_.ID, objc.Sel("mediaStyle"))
	return rv
}


// SetMediaStyle sets the value of the mediaStyle property.
// The receiver’s CSS media property.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mediastyle
func (w_ WebView) SetMediaStyle(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaStyle:"), objc.String(value))
}

// The receiver’s custom user-agent string.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/customuseragent
func (w_ WebView) CustomUserAgent() string {
	rv := objc.Send[string](w_.ID, objc.Sel("customUserAgent"))
	return rv
}


// SetCustomUserAgent sets the value of the customUserAgent property.
// The receiver’s custom user-agent string.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/customuseragent
func (w_ WebView) SetCustomUserAgent(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomUserAgent:"), objc.String(value))
}

// The receiver’s frame load delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/frameloaddelegate
func (w_ WebView) FrameLoadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("frameLoadDelegate"))
	return rv
}


// SetFrameLoadDelegate sets the value of the frameLoadDelegate property.
// The receiver’s frame load delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/frameloaddelegate
func (w_ WebView) SetFrameLoadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameLoadDelegate:"), value)
}

// A Boolean that indicates whether the text can be made larger.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextlarger
func (w_ WebView) CanMakeTextLarger() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextLarger"))
	return rv
}


// SetCanMakeTextLarger sets the value of the canMakeTextLarger property.
// A Boolean that indicates whether the text can be made larger.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextlarger
func (w_ WebView) SetCanMakeTextLarger(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanMakeTextLarger:"), value)
}

// An array of pasteboard types that can be used for the current selection of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/pasteboardtypesforselection
func (w_ WebView) PasteboardTypesForSelection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("pasteboardTypesForSelection"))
	return rv
}


// SetPasteboardTypesForSelection sets the value of the pasteboardTypesForSelection property.
// An array of pasteboard types that can be used for the current selection of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/pasteboardtypesforselection
func (w_ WebView) SetPasteboardTypesForSelection(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPasteboardTypesForSelection:"), value)
}

// The receiver’s policy delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/policydelegate
func (w_ WebView) PolicyDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("policyDelegate"))
	return rv
}


// SetPolicyDelegate sets the value of the policyDelegate property.
// The receiver’s policy delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/policydelegate
func (w_ WebView) SetPolicyDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPolicyDelegate:"), value)
}

// The receiver’s download delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/downloaddelegate
func (w_ WebView) DownloadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("downloadDelegate"))
	return rv
}


// SetDownloadDelegate sets the value of the downloadDelegate property.
// The receiver’s download delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/downloaddelegate
func (w_ WebView) SetDownloadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDownloadDelegate:"), value)
}

// A Boolean that indicates whether the selection is maintained when focus is lost.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/maintainsinactiveselection
func (w_ WebView) MaintainsInactiveSelection() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("maintainsInactiveSelection"))
	return rv
}


// SetMaintainsInactiveSelection sets the value of the maintainsInactiveSelection property.
// A Boolean that indicates whether the selection is maintained when focus is lost.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/maintainsinactiveselection
func (w_ WebView) SetMaintainsInactiveSelection(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMaintainsInactiveSelection:"), value)
}

// The receiver’s window object from the scripting environment.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/windowscriptobject
func (w_ WebView) WindowScriptObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("windowScriptObject"))
	return rv
}


// SetWindowScriptObject sets the value of the windowScriptObject property.
// The receiver’s window object from the scripting environment.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/windowscriptobject
func (w_ WebView) SetWindowScriptObject(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowScriptObject:"), value)
}

// A Boolean that indicates whether the web view is loading content.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/isloading
func (w_ WebView) IsLoading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoading"))
	return rv
}


// SetIsLoading sets the value of the isLoading property.
// A Boolean that indicates whether the web view is loading content.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/isloading
func (w_ WebView) SetIsLoading(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoading:"), value)
}

// The custom text encoding name.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/customtextencodingname
func (w_ WebView) CustomTextEncodingName() string {
	rv := objc.Send[string](w_.ID, objc.Sel("customTextEncodingName"))
	return rv
}


// SetCustomTextEncodingName sets the value of the customTextEncodingName property.
// The custom text encoding name.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/customtextencodingname
func (w_ WebView) SetCustomTextEncodingName(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomTextEncodingName:"), objc.String(value))
}

// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iseditable
func (w_ WebView) IsEditable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isEditable"))
	return rv
}


// SetIsEditable sets the value of the isEditable property.
// A Boolean that indicates whether the user is allowed to edit the document.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iseditable
func (w_ WebView) SetIsEditable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsEditable:"), value)
}

// The receiver’s application name that is used in the user-agent string.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/applicationnameforuseragent
func (w_ WebView) ApplicationNameForUserAgent() string {
	rv := objc.Send[string](w_.ID, objc.Sel("applicationNameForUserAgent"))
	return rv
}


// SetApplicationNameForUserAgent sets the value of the applicationNameForUserAgent property.
// The receiver’s application name that is used in the user-agent string.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/applicationnameforuseragent
func (w_ WebView) SetApplicationNameForUserAgent(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setApplicationNameForUserAgent:"), objc.String(value))
}

// The URL that the main frame loads.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframeurl
func (w_ WebView) MainFrameURL() string {
	rv := objc.Send[string](w_.ID, objc.Sel("mainFrameURL"))
	return rv
}


// SetMainFrameURL sets the value of the mainFrameURL property.
// The URL that the main frame loads.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframeurl
func (w_ WebView) SetMainFrameURL(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameURL:"), objc.String(value))
}

// The receiver’s editing delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/editingdelegate
func (w_ WebView) EditingDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("editingDelegate"))
	return rv
}


// SetEditingDelegate sets the value of the editingDelegate property.
// The receiver’s editing delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/editingdelegate
func (w_ WebView) SetEditingDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEditingDelegate:"), value)
}

// A Boolean that inidicates whether the web view should update even when it is not in a window that is currently visible.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/shouldupdatewhileoffscreen
func (w_ WebView) ShouldUpdateWhileOffscreen() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldUpdateWhileOffscreen"))
	return rv
}


// SetShouldUpdateWhileOffscreen sets the value of the shouldUpdateWhileOffscreen property.
// A Boolean that inidicates whether the web view should update even when it is not in a window that is currently visible.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/shouldupdatewhileoffscreen
func (w_ WebView) SetShouldUpdateWhileOffscreen(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldUpdateWhileOffscreen:"), value)
}

// A Boolean that indicates whether the current text size is a multiple of 1.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextstandardsize
func (w_ WebView) CanMakeTextStandardSize() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextStandardSize"))
	return rv
}


// SetCanMakeTextStandardSize sets the value of the canMakeTextStandardSize property.
// A Boolean that indicates whether the current text size is a multiple of 1.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextstandardsize
func (w_ WebView) SetCanMakeTextStandardSize(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanMakeTextStandardSize:"), value)
}

// A Boolean that indicates whether the web view has continuous spell-checking enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iscontinuousspellcheckingenabled
func (w_ WebView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}


// SetIsContinuousSpellCheckingEnabled sets the value of the isContinuousSpellCheckingEnabled property.
// A Boolean that indicates whether the web view has continuous spell-checking enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iscontinuousspellcheckingenabled
func (w_ WebView) SetIsContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsContinuousSpellCheckingEnabled:"), value)
}

// The frame with the active selection.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selectedframe
func (w_ WebView) SelectedFrame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("selectedFrame"))
	return rv
}


// SetSelectedFrame sets the value of the selectedFrame property.
// The frame with the active selection.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selectedframe
func (w_ WebView) SetSelectedFrame(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectedFrame:"), value)
}

// The main frame, the root of the web frame hierarchy for this page.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframe
func (w_ WebView) MainFrame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("mainFrame"))
	return rv
}


// SetMainFrame sets the value of the mainFrame property.
// The main frame, the root of the web frame hierarchy for this page.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframe
func (w_ WebView) SetMainFrame(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrame:"), value)
}

// A Boolean that indicates whether the previous location can be loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/cangoback
func (w_ WebView) CanGoBack() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoBack"))
	return rv
}


// SetCanGoBack sets the value of the canGoBack property.
// A Boolean that indicates whether the previous location can be loaded.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/cangoback
func (w_ WebView) SetCanGoBack(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanGoBack:"), value)
}

// The current selection affinity.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selectionaffinity
func (w_ WebView) SelectionAffinity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("selectionAffinity"))
	return rv
}


// SetSelectionAffinity sets the value of the selectionAffinity property.
// The current selection affinity.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selectionaffinity
func (w_ WebView) SetSelectionAffinity(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectionAffinity:"), value)
}

// The receiver’s undo manager.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/undomanager
func (w_ WebView) UndoManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("undoManager"))
	return rv
}


// SetUndoManager sets the value of the undoManager property.
// The receiver’s undo manager.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/undomanager
func (w_ WebView) SetUndoManager(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUndoManager:"), value)
}

// A Boolean that indicates whether the text can be made smaller.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextsmaller
func (w_ WebView) CanMakeTextSmaller() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextSmaller"))
	return rv
}


// SetCanMakeTextSmaller sets the value of the canMakeTextSmaller property.
// A Boolean that indicates whether the text can be made smaller.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextsmaller
func (w_ WebView) SetCanMakeTextSmaller(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanMakeTextSmaller:"), value)
}

// The receiver’s CSS typing style.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/typingstyle
func (w_ WebView) TypingStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("typingStyle"))
	return rv
}


// SetTypingStyle sets the value of the typingStyle property.
// The receiver’s CSS typing style.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/typingstyle
func (w_ WebView) SetTypingStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTypingStyle:"), value)
}

// A Boolean that indicates whether the document view supports different text encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/supportstextencoding
func (w_ WebView) SupportsTextEncoding() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("supportsTextEncoding"))
	return rv
}


// SetSupportsTextEncoding sets the value of the supportsTextEncoding property.
// A Boolean that indicates whether the document view supports different text encodings.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/supportstextencoding
func (w_ WebView) SetSupportsTextEncoding(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSupportsTextEncoding:"), value)
}

// The receiver’s user interface delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/uidelegate
func (w_ WebView) UiDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("uiDelegate"))
	return rv
}


// SetUiDelegate sets the value of the uiDelegate property.
// The receiver’s user interface delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/uidelegate
func (w_ WebView) SetUiDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUiDelegate:"), value)
}

// The font size multiplier for text displayed in web frame view objects managed by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/textsizemultiplier
func (w_ WebView) TextSizeMultiplier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("textSizeMultiplier"))
	return rv
}


// SetTextSizeMultiplier sets the value of the textSizeMultiplier property.
// The font size multiplier for text displayed in web frame view objects managed by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/textsizemultiplier
func (w_ WebView) SetTextSizeMultiplier(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTextSizeMultiplier:"), value)
}

// A Boolean that indicates whether smart-space insertion and deletion is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/smartinsertdeleteenabled
func (w_ WebView) SmartInsertDeleteEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}


// SetSmartInsertDeleteEnabled sets the value of the smartInsertDeleteEnabled property.
// A Boolean that indicates whether smart-space insertion and deletion is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/smartinsertdeleteenabled
func (w_ WebView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}

// The navigation type of the action. Can be any of the values defined in
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webactionnavigationtypekey
func (w_ WebView) WebActionNavigationTypeKey() string {
	rv := objc.Send[string](w_.ID, objc.Sel("WebActionNavigationTypeKey"))
	return rv
}

// The HTML title of the loaded page.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframetitle
func (w_ WebView) MainFrameTitle() string {
	rv := objc.Send[string](w_.ID, objc.Sel("mainFrameTitle"))
	return rv
}


// SetMainFrameTitle sets the value of the mainFrameTitle property.
// The HTML title of the loaded page.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframetitle
func (w_ WebView) SetMainFrameTitle(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameTitle:"), objc.String(value))
}

// The DOM document for the main frame.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframedocument
func (w_ WebView) MainFrameDocument() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("mainFrameDocument"))
	return rv
}


// SetMainFrameDocument sets the value of the mainFrameDocument property.
// The DOM document for the main frame.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframedocument
func (w_ WebView) SetMainFrameDocument(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameDocument:"), value)
}

// The receiver’s group name.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/groupname
func (w_ WebView) GroupName() string {
	rv := objc.Send[string](w_.ID, objc.Sel("groupName"))
	return rv
}


// SetGroupName sets the value of the groupName property.
// The receiver’s group name.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/groupname
func (w_ WebView) SetGroupName(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGroupName:"), objc.String(value))
}

// A Boolean that indicates whether the web view should close when its window or host window closes.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/shouldclosewithwindow
func (w_ WebView) ShouldCloseWithWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldCloseWithWindow"))
	return rv
}


// SetShouldCloseWithWindow sets the value of the shouldCloseWithWindow property.
// A Boolean that indicates whether the web view should close when its window or host window closes.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/shouldclosewithwindow
func (w_ WebView) SetShouldCloseWithWindow(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldCloseWithWindow:"), value)
}

// The receiver’s host window.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/hostwindow
func (w_ WebView) HostWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("hostWindow"))
	return rv
}


// SetHostWindow sets the value of the hostWindow property.
// The receiver’s host window.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/hostwindow
func (w_ WebView) SetHostWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHostWindow:"), value)
}

// The range of the current selection.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selecteddomrange
func (w_ WebView) SelectedDOMRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("selectedDOMRange"))
	return rv
}


// SetSelectedDOMRange sets the value of the selectedDOMRange property.
// The range of the current selection.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selecteddomrange
func (w_ WebView) SetSelectedDOMRange(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectedDOMRange:"), value)
}

// The spell-checker document tag for this document.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/spellcheckerdocumenttag
func (w_ WebView) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](w_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}


// SetSpellCheckerDocumentTag sets the value of the spellCheckerDocumentTag property.
// The spell-checker document tag for this document.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/spellcheckerdocumenttag
func (w_ WebView) SetSpellCheckerDocumentTag(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSpellCheckerDocumentTag:"), value)
}

// A Boolean that indicates whether the web view draws a background.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/drawsbackground
func (w_ WebView) DrawsBackground() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("drawsBackground"))
	return rv
}


// SetDrawsBackground sets the value of the drawsBackground property.
// A Boolean that indicates whether the web view draws a background.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/drawsbackground
func (w_ WebView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDrawsBackground:"), value)
}

// The site’s favicon.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframeicon
func (w_ WebView) MainFrameIcon() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("mainFrameIcon"))
	return rv
}


// SetMainFrameIcon sets the value of the mainFrameIcon property.
// The site’s favicon.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframeicon
func (w_ WebView) SetMainFrameIcon(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameIcon:"), value)
}

// The receiver’s resource load delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/resourceloaddelegate
func (w_ WebView) ResourceLoadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("resourceLoadDelegate"))
	return rv
}


// SetResourceLoadDelegate sets the value of the resourceLoadDelegate property.
// The receiver’s resource load delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/resourceloaddelegate
func (w_ WebView) SetResourceLoadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setResourceLoadDelegate:"), value)
}

// A Boolean that indicates whether the next location can be loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/cangoforward
func (w_ WebView) CanGoForward() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoForward"))
	return rv
}


// SetCanGoForward sets the value of the canGoForward property.
// A Boolean that indicates whether the next location can be loaded.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/cangoforward
func (w_ WebView) SetCanGoForward(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanGoForward:"), value)
}

// The receiver’s preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/preferences
func (w_ WebView) Preferences() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("preferences"))
	return rv
}


// SetPreferences sets the value of the preferences property.
// The receiver’s preferences.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/preferences
func (w_ WebView) SetPreferences(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferences:"), value)
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




