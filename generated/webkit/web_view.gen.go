// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	BackForwardList() unsafe.Pointer
	CustomTextEncodingName() objc.IObject /* cross-framework: NSString */
	SetCustomTextEncodingName(value objc.IObject /* cross-framework: NSString */)
	EstimatedProgress() float64
	Editable() bool
	SetEditable(value bool)
	MainFrameURL() objc.IObject /* cross-framework: NSString */
	SetMainFrameURL(value objc.IObject /* cross-framework: NSString */)
	PreferencesIdentifier() objc.IObject /* cross-framework: NSString */
	SetPreferencesIdentifier(value objc.IObject /* cross-framework: NSString */)
	SelectedFrame() IWebFrame
	TypingStyle() unsafe.Pointer
	SetTypingStyle(value unsafe.Pointer)
	WebActionNavigationTypeKey() objc.IObject /* cross-framework: NSString */
	Autosaves() bool
	SetAutosaves(value bool)
	ApplicationNameForUserAgent() objc.IObject /* cross-framework: NSString */
	SetApplicationNameForUserAgent(value objc.IObject /* cross-framework: NSString */)
	CanGoBack() bool
	SetCanGoBack(value bool)
	CanGoForward() bool
	SetCanGoForward(value bool)
	CanMakeTextLarger() bool
	SetCanMakeTextLarger(value bool)
	CanMakeTextSmaller() bool
	SetCanMakeTextSmaller(value bool)
	CanMakeTextStandardSize() bool
	SetCanMakeTextStandardSize(value bool)
	CustomUserAgent() objc.IObject /* cross-framework: NSString */
	SetCustomUserAgent(value objc.IObject /* cross-framework: NSString */)
	DownloadDelegate() unsafe.Pointer
	SetDownloadDelegate(value unsafe.Pointer)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	EditingDelegate() unsafe.Pointer
	SetEditingDelegate(value unsafe.Pointer)
	FrameLoadDelegate() unsafe.Pointer
	SetFrameLoadDelegate(value unsafe.Pointer)
	GroupName() objc.IObject /* cross-framework: NSString */
	SetGroupName(value objc.IObject /* cross-framework: NSString */)
	HostWindow() objc.IObject /* cross-framework: Window */
	SetHostWindow(value objc.IObject /* cross-framework: Window */)
	IsContinuousSpellCheckingEnabled() bool
	SetIsContinuousSpellCheckingEnabled(value bool)
	IsEditable() bool
	SetIsEditable(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
	MainFrame() IWebFrame
	SetMainFrame(value IWebFrame)
	MainFrameDocument() unsafe.Pointer
	SetMainFrameDocument(value unsafe.Pointer)
	MainFrameIcon() objc.IObject /* cross-framework: Image */
	SetMainFrameIcon(value objc.IObject /* cross-framework: Image */)
	MainFrameTitle() objc.IObject /* cross-framework: NSString */
	SetMainFrameTitle(value objc.IObject /* cross-framework: NSString */)
	MaintainsInactiveSelection() bool
	SetMaintainsInactiveSelection(value bool)
	MediaStyle() objc.IObject /* cross-framework: NSString */
	SetMediaStyle(value objc.IObject /* cross-framework: NSString */)
	PasteboardTypesForSelection() unsafe.Pointer
	SetPasteboardTypesForSelection(value unsafe.Pointer)
	PolicyDelegate() unsafe.Pointer
	SetPolicyDelegate(value unsafe.Pointer)
	Preferences() unsafe.Pointer
	SetPreferences(value unsafe.Pointer)
	ResourceLoadDelegate() unsafe.Pointer
	SetResourceLoadDelegate(value unsafe.Pointer)
	SelectedDOMRange() unsafe.Pointer
	SetSelectedDOMRange(value unsafe.Pointer)
	SelectionAffinity() SelectionAffinity /* not a class type */
	SetSelectionAffinity(value SelectionAffinity /* not a class type */)
	ShouldCloseWithWindow() bool
	SetShouldCloseWithWindow(value bool)
	ShouldUpdateWhileOffscreen() bool
	SetShouldUpdateWhileOffscreen(value bool)
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	SpellCheckerDocumentTag() int
	SetSpellCheckerDocumentTag(value int)
	SupportsTextEncoding() bool
	SetSupportsTextEncoding(value bool)
	TextSizeMultiplier() float32
	SetTextSizeMultiplier(value float32)
	UiDelegate() unsafe.Pointer
	SetUiDelegate(value unsafe.Pointer)
	UndoManager() objc.IObject /* cross-framework: UndoManager */
	SetUndoManager(value objc.IObject /* cross-framework: UndoManager */)
	WindowScriptObject() unsafe.Pointer
	SetWindowScriptObject(value unsafe.Pointer)
	// methods:
	ChangeDocumentBackgroundColor(sender objectivec.IObject)
	ChangeFont(sender objectivec.IObject)
	GoBack(sender objectivec.IObject)
	GoForward(sender objectivec.IObject)
	MoveToBeginningOfSentence(sender objectivec.IObject)
	ReloadFromOrigin(sender objectivec.IObject)
	ReplaceSelectionWithNode(node unsafe.Pointer)
	StartSpeaking(sender objectivec.IObject)
	StopLoading(sender objectivec.IObject)
}

// is the core view class in the WebKit framework that manages interactions between the and classes. To embed web content in your application, you just create a object, attach it to a window, and send a message to its main frame.
//
// Behind the scenes, objects encapsulate the content contained in a single frame element. A hierarchy of objects is used to model an entire webpage where the root is called the . There is a object per object used to display the frame content. Therefore, there is a parallel hierarchy of objects used to render an entire page. The object is also the parent view of this hierarchy. You do not need to create and objects directly. These objects are automatically created when the page loads, either programmatically or by the user clicking a link. You customize your embedded web content by implementing delegates to handle certain aspects of the process. objects have multiple delegates because the process of loading a webpage is asynchronous and complicated if errors occur. All the delegates use informal protocols so you only need to implement only the delegates and methods that define the behavior you wish to change—default implementations are already provided. For example, you might want to implement the frame load and resource load delegates to monitor the load progress and display status messages. Applications that use multiple windows may want to implement a user interface delegate. See the individual informal delegate protocols for more details: , , , and . Another way to monitor load progress with less control is to observe the , , and notifications. For example, you could observe these notifications to implement a simple progress indicator in your application. You update the progress indicator by invoking the method to get an estimate of the amount of content that is currently loaded. A object is intended to support most features you would expect in a web browser except that it doesn’t implement the specific user interface for those features. You are responsible for implementing the user interface objects such as status bars, toolbars, buttons, and text fields. For example, a object manages a back-forward list by default, and has and action methods. It is your responsibility to create the buttons that would send theses action messages. Note, there is some overhead in maintaining a back-forward list and page cache, so you should disable it if your application doesn’t use it. You use a object to encapsulate the preferences of a object, such as the font, text encoding, and image settings. You can modify the preferences for individual objects or specify a shared object using the method. Use the method to specify whether the preferences should be automatically saved to the user defaults database. You can also extend WebKit by implementing your own document view and representation classes for specific MIME types. Use the class method to register your custom classes with a object.


// is the core view class in the WebKit framework that manages interactions between the and classes. To embed web content in your application, you just create a object, attach it to a window, and send a message to its main frame.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canShowMIMEType(_:)
func (wc _WebViewClass) CanShowMIMEType(MIMEType objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("canShowMIMEType:"), MIMEType)
	return rv
}


// Specifies the view and representation objects to be used for specific MIME types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/registerClass(_:representationClass:forMIMEType:)
func (wc _WebViewClass) RegisterViewClassRepresentationClassForMIMEType(viewClass objc.Class, representationClass objc.Class, MIMEType objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("registerViewClass:representationClass:forMIMEType:"), viewClass, representationClass, MIMEType)
}


// Sets the background color of the selected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/changeDocumentBackgroundColor(_:)
func (w_ WebView) ChangeDocumentBackgroundColor(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("changeDocumentBackgroundColor:"), sender)
}


// An action method that changes the font of the selection, or all content if there is no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/changeFont(_:)
func (w_ WebView) ChangeFont(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("changeFont:"), sender)
}


// An action method that loads the previous location in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/goBack(_:)
func (w_ WebView) GoBack(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goBack:"), sender)
}


// An action method that loads the next location in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/goForward(_:)
func (w_ WebView) GoForward(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goForward:"), sender)
}


// Moves the insertion point to the beginning of the current sentence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/moveToBeginningOfSentence(_:)
func (w_ WebView) MoveToBeginningOfSentence(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("moveToBeginningOfSentence:"), sender)
}


// Action method that performs an end-to-end revalidation using cache-validating conditionals if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/reloadFromOrigin(_:)
func (w_ WebView) ReloadFromOrigin(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("reloadFromOrigin:"), sender)
}


// Replaces the receiver’s current selection with the specified DOM node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/replaceSelection(with:)-5px9m
func (w_ WebView) ReplaceSelectionWithNode(node unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("replaceSelectionWithNode:"), node)
}


// An action method that starts speaking the selected text or all text if there’s no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/startSpeaking(_:)
func (w_ WebView) StartSpeaking(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("startSpeaking:"), sender)
}


// An action method that stops the loading of any web frame content managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/stopLoading(_:)
func (w_ WebView) StopLoading(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopLoading:"), sender)
}


// The receiver’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/backForwardList
func (w_ WebView) BackForwardList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("backForwardList"))
	return rv
}


// The custom text encoding name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/customTextEncodingName
func (w_ WebView) CustomTextEncodingName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("customTextEncodingName"))
	return rv
}


// The custom text encoding name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/customTextEncodingName
func (w_ WebView) SetCustomTextEncodingName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomTextEncodingName:"), value)
}


// An estimate, as a percentage, of the amount of content that is currently loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/estimatedProgress
func (w_ WebView) EstimatedProgress() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("estimatedProgress"))
	return rv
}


// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isEditable
func (w_ WebView) Editable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("editable"))
	return rv
}


// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isEditable
func (w_ WebView) SetEditable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEditable:"), value)
}


// The URL that the main frame loads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mainFrameURL
func (w_ WebView) MainFrameURL() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("mainFrameURL"))
	return rv
}


// The URL that the main frame loads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mainFrameURL
func (w_ WebView) SetMainFrameURL(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameURL:"), value)
}


// The identifier of the receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/preferencesIdentifier
func (w_ WebView) PreferencesIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("preferencesIdentifier"))
	return rv
}


// The identifier of the receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/preferencesIdentifier
func (w_ WebView) SetPreferencesIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferencesIdentifier:"), value)
}


// The frame with the active selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/selectedFrame
func (w_ WebView) SelectedFrame() IWebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("selectedFrame"))
	return rv
}


// The receiver’s CSS typing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/typingStyle
func (w_ WebView) TypingStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("typingStyle"))
	return rv
}


// The receiver’s CSS typing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/typingStyle
func (w_ WebView) SetTypingStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTypingStyle:"), value)
}


// The navigation type of the action. Can be any of the values defined in
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webactionnavigationtypekey
func (w_ WebView) WebActionNavigationTypeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("WebActionNavigationTypeKey"))
	return rv
}


// A Boolean that indicates whether or not the receiver’s attributes are automatically stored in the user defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/autosaves
func (w_ WebView) Autosaves() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autosaves"))
	return rv
}


// A Boolean that indicates whether or not the receiver’s attributes are automatically stored in the user defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/autosaves
func (w_ WebView) SetAutosaves(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutosaves:"), value)
}


// The receiver’s application name that is used in the user-agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/applicationnameforuseragent
func (w_ WebView) ApplicationNameForUserAgent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("applicationNameForUserAgent"))
	return rv
}


// The receiver’s application name that is used in the user-agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/applicationnameforuseragent
func (w_ WebView) SetApplicationNameForUserAgent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setApplicationNameForUserAgent:"), value)
}


// A Boolean that indicates whether the previous location can be loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/cangoback
func (w_ WebView) CanGoBack() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoBack"))
	return rv
}


// A Boolean that indicates whether the previous location can be loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/cangoback
func (w_ WebView) SetCanGoBack(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanGoBack:"), value)
}


// A Boolean that indicates whether the next location can be loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/cangoforward
func (w_ WebView) CanGoForward() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoForward"))
	return rv
}


// A Boolean that indicates whether the next location can be loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/cangoforward
func (w_ WebView) SetCanGoForward(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanGoForward:"), value)
}


// A Boolean that indicates whether the text can be made larger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextlarger
func (w_ WebView) CanMakeTextLarger() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextLarger"))
	return rv
}


// A Boolean that indicates whether the text can be made larger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextlarger
func (w_ WebView) SetCanMakeTextLarger(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanMakeTextLarger:"), value)
}


// A Boolean that indicates whether the text can be made smaller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextsmaller
func (w_ WebView) CanMakeTextSmaller() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextSmaller"))
	return rv
}


// A Boolean that indicates whether the text can be made smaller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextsmaller
func (w_ WebView) SetCanMakeTextSmaller(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanMakeTextSmaller:"), value)
}


// A Boolean that indicates whether the current text size is a multiple of 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextstandardsize
func (w_ WebView) CanMakeTextStandardSize() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextStandardSize"))
	return rv
}


// A Boolean that indicates whether the current text size is a multiple of 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/canmaketextstandardsize
func (w_ WebView) SetCanMakeTextStandardSize(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanMakeTextStandardSize:"), value)
}


// The receiver’s custom user-agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/customuseragent
func (w_ WebView) CustomUserAgent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("customUserAgent"))
	return rv
}


// The receiver’s custom user-agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/customuseragent
func (w_ WebView) SetCustomUserAgent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomUserAgent:"), value)
}


// The receiver’s download delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/downloaddelegate
func (w_ WebView) DownloadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("downloadDelegate"))
	return rv
}


// The receiver’s download delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/downloaddelegate
func (w_ WebView) SetDownloadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDownloadDelegate:"), value)
}


// A Boolean that indicates whether the web view draws a background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/drawsbackground
func (w_ WebView) DrawsBackground() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean that indicates whether the web view draws a background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/drawsbackground
func (w_ WebView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The receiver’s editing delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/editingdelegate
func (w_ WebView) EditingDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("editingDelegate"))
	return rv
}


// The receiver’s editing delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/editingdelegate
func (w_ WebView) SetEditingDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEditingDelegate:"), value)
}


// The receiver’s frame load delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/frameloaddelegate
func (w_ WebView) FrameLoadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("frameLoadDelegate"))
	return rv
}


// The receiver’s frame load delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/frameloaddelegate
func (w_ WebView) SetFrameLoadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameLoadDelegate:"), value)
}


// The receiver’s group name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/groupname
func (w_ WebView) GroupName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("groupName"))
	return rv
}


// The receiver’s group name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/groupname
func (w_ WebView) SetGroupName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGroupName:"), value)
}


// The receiver’s host window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/hostwindow
func (w_ WebView) HostWindow() objc.IObject /* cross-framework: Window */ {
	rv := objc.Send[appkit.Window](w_.ID, objc.Sel("hostWindow"))
	return rv
}


// The receiver’s host window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/hostwindow
func (w_ WebView) SetHostWindow(value objc.IObject /* cross-framework: Window */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHostWindow:"), value)
}


// A Boolean that indicates whether the web view has continuous spell-checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iscontinuousspellcheckingenabled
func (w_ WebView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}


// A Boolean that indicates whether the web view has continuous spell-checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iscontinuousspellcheckingenabled
func (w_ WebView) SetIsContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsContinuousSpellCheckingEnabled:"), value)
}


// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iseditable
func (w_ WebView) IsEditable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iseditable
func (w_ WebView) SetIsEditable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean that indicates whether the web view is loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/isloading
func (w_ WebView) IsLoading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoading"))
	return rv
}


// A Boolean that indicates whether the web view is loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/isloading
func (w_ WebView) SetIsLoading(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoading:"), value)
}


// The main frame, the root of the web frame hierarchy for this page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframe
func (w_ WebView) MainFrame() IWebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("mainFrame"))
	return rv
}


// The main frame, the root of the web frame hierarchy for this page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframe
func (w_ WebView) SetMainFrame(value IWebFrame) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrame:"), value)
}


// The DOM document for the main frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframedocument
func (w_ WebView) MainFrameDocument() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("mainFrameDocument"))
	return rv
}


// The DOM document for the main frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframedocument
func (w_ WebView) SetMainFrameDocument(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameDocument:"), value)
}


// The site’s favicon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframeicon
func (w_ WebView) MainFrameIcon() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](w_.ID, objc.Sel("mainFrameIcon"))
	return rv
}


// The site’s favicon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframeicon
func (w_ WebView) SetMainFrameIcon(value objc.IObject /* cross-framework: Image */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameIcon:"), value)
}


// The HTML title of the loaded page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframetitle
func (w_ WebView) MainFrameTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("mainFrameTitle"))
	return rv
}


// The HTML title of the loaded page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mainframetitle
func (w_ WebView) SetMainFrameTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameTitle:"), value)
}


// A Boolean that indicates whether the selection is maintained when focus is lost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/maintainsinactiveselection
func (w_ WebView) MaintainsInactiveSelection() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("maintainsInactiveSelection"))
	return rv
}


// A Boolean that indicates whether the selection is maintained when focus is lost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/maintainsinactiveselection
func (w_ WebView) SetMaintainsInactiveSelection(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMaintainsInactiveSelection:"), value)
}


// The receiver’s CSS media property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mediastyle
func (w_ WebView) MediaStyle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("mediaStyle"))
	return rv
}


// The receiver’s CSS media property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/mediastyle
func (w_ WebView) SetMediaStyle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaStyle:"), value)
}


// An array of pasteboard types that can be used for the current selection of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/pasteboardtypesforselection
func (w_ WebView) PasteboardTypesForSelection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("pasteboardTypesForSelection"))
	return rv
}


// An array of pasteboard types that can be used for the current selection of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/pasteboardtypesforselection
func (w_ WebView) SetPasteboardTypesForSelection(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPasteboardTypesForSelection:"), value)
}


// The receiver’s policy delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/policydelegate
func (w_ WebView) PolicyDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("policyDelegate"))
	return rv
}


// The receiver’s policy delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/policydelegate
func (w_ WebView) SetPolicyDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPolicyDelegate:"), value)
}


// The receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/preferences
func (w_ WebView) Preferences() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("preferences"))
	return rv
}


// The receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/preferences
func (w_ WebView) SetPreferences(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferences:"), value)
}


// The receiver’s resource load delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/resourceloaddelegate
func (w_ WebView) ResourceLoadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("resourceLoadDelegate"))
	return rv
}


// The receiver’s resource load delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/resourceloaddelegate
func (w_ WebView) SetResourceLoadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setResourceLoadDelegate:"), value)
}


// The range of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selecteddomrange
func (w_ WebView) SelectedDOMRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("selectedDOMRange"))
	return rv
}


// The range of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selecteddomrange
func (w_ WebView) SetSelectedDOMRange(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectedDOMRange:"), value)
}


// The current selection affinity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selectionaffinity
func (w_ WebView) SelectionAffinity() SelectionAffinity /* not a class type */ {
	rv := objc.Send[SelectionAffinity](w_.ID, objc.Sel("selectionAffinity"))
	return rv
}


// The current selection affinity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/selectionaffinity
func (w_ WebView) SetSelectionAffinity(value SelectionAffinity /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectionAffinity:"), value)
}


// A Boolean that indicates whether the web view should close when its window or host window closes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/shouldclosewithwindow
func (w_ WebView) ShouldCloseWithWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldCloseWithWindow"))
	return rv
}


// A Boolean that indicates whether the web view should close when its window or host window closes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/shouldclosewithwindow
func (w_ WebView) SetShouldCloseWithWindow(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldCloseWithWindow:"), value)
}


// A Boolean that inidicates whether the web view should update even when it is not in a window that is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/shouldupdatewhileoffscreen
func (w_ WebView) ShouldUpdateWhileOffscreen() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldUpdateWhileOffscreen"))
	return rv
}


// A Boolean that inidicates whether the web view should update even when it is not in a window that is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/shouldupdatewhileoffscreen
func (w_ WebView) SetShouldUpdateWhileOffscreen(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldUpdateWhileOffscreen:"), value)
}


// A Boolean that indicates whether smart-space insertion and deletion is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/smartinsertdeleteenabled
func (w_ WebView) SmartInsertDeleteEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}


// A Boolean that indicates whether smart-space insertion and deletion is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/smartinsertdeleteenabled
func (w_ WebView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}


// The spell-checker document tag for this document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/spellcheckerdocumenttag
func (w_ WebView) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](w_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}


// The spell-checker document tag for this document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/spellcheckerdocumenttag
func (w_ WebView) SetSpellCheckerDocumentTag(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSpellCheckerDocumentTag:"), value)
}


// A Boolean that indicates whether the document view supports different text encodings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/supportstextencoding
func (w_ WebView) SupportsTextEncoding() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("supportsTextEncoding"))
	return rv
}


// A Boolean that indicates whether the document view supports different text encodings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/supportstextencoding
func (w_ WebView) SetSupportsTextEncoding(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSupportsTextEncoding:"), value)
}


// The font size multiplier for text displayed in web frame view objects managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/textsizemultiplier
func (w_ WebView) TextSizeMultiplier() float32 {
	rv := objc.Send[float32](w_.ID, objc.Sel("textSizeMultiplier"))
	return rv
}


// The font size multiplier for text displayed in web frame view objects managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/textsizemultiplier
func (w_ WebView) SetTextSizeMultiplier(value float32) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTextSizeMultiplier:"), value)
}


// The receiver’s user interface delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/uidelegate
func (w_ WebView) UiDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("uiDelegate"))
	return rv
}


// The receiver’s user interface delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/uidelegate
func (w_ WebView) SetUiDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUiDelegate:"), value)
}


// The receiver’s undo manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/undomanager
func (w_ WebView) UndoManager() objc.IObject /* cross-framework: UndoManager */ {
	rv := objc.Send[foundation.UndoManager](w_.ID, objc.Sel("undoManager"))
	return rv
}


// The receiver’s undo manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/undomanager
func (w_ WebView) SetUndoManager(value objc.IObject /* cross-framework: UndoManager */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUndoManager:"), value)
}


// The receiver’s window object from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/windowscriptobject
func (w_ WebView) WindowScriptObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("windowScriptObject"))
	return rv
}


// The receiver’s window object from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/windowscriptobject
func (w_ WebView) SetWindowScriptObject(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowScriptObject:"), value)
}




