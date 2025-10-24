// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class WebView */

/* debug [class_header]: Header for WebView */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebView */
// An interface definition for the [WebView] class.
type IWebView interface {
	appkit.IView

	/* debug [class_interface_properties]: Properties for WebView */
	// properties:
	ApplicationNameForUserAgent() objc.IObject /* cross-framework: NSString */
	SetApplicationNameForUserAgent(value objc.IObject /* cross-framework: NSString */)
	BackForwardList() IWebBackForwardList
	CanGoBack() bool
	CanGoForward() bool
	CanMakeTextLarger() bool
	CanMakeTextSmaller() bool
	CanMakeTextStandardSize() bool
	CustomTextEncodingName() objc.IObject /* cross-framework: NSString */
	SetCustomTextEncodingName(value objc.IObject /* cross-framework: NSString */)
	CustomUserAgent() objc.IObject /* cross-framework: NSString */
	SetCustomUserAgent(value objc.IObject /* cross-framework: NSString */)
	DownloadDelegate() unsafe.Pointer
	SetDownloadDelegate(value unsafe.Pointer)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	EditingDelegate() unsafe.Pointer
	SetEditingDelegate(value unsafe.Pointer)
	EstimatedProgress() float64
	FrameLoadDelegate() unsafe.Pointer
	SetFrameLoadDelegate(value unsafe.Pointer)
	GroupName() objc.IObject /* cross-framework: NSString */
	SetGroupName(value objc.IObject /* cross-framework: NSString */)
	HostWindow() appkit.Window
	SetHostWindow(value appkit.Window)
	ContinuousSpellCheckingEnabled() bool
	SetContinuousSpellCheckingEnabled(value bool)
	Editable() bool
	SetEditable(value bool)
	Loading() bool
	MainFrame() IWebFrame
	MainFrameDocument() IDOMDocument
	MainFrameIcon() appkit.Image
	MainFrameTitle() objc.IObject /* cross-framework: NSString */
	MainFrameURL() objc.IObject   /* cross-framework: NSString */
	SetMainFrameURL(value objc.IObject /* cross-framework: NSString */)
	MaintainsInactiveSelection() bool
	MediaStyle() objc.IObject /* cross-framework: NSString */
	SetMediaStyle(value objc.IObject /* cross-framework: NSString */)
	PasteboardTypesForSelection() objc.IObject /* cross-framework: NSArray */
	PolicyDelegate() unsafe.Pointer
	SetPolicyDelegate(value unsafe.Pointer)
	Preferences() IWebPreferences
	SetPreferences(value IWebPreferences)
	PreferencesIdentifier() objc.IObject /* cross-framework: NSString */
	SetPreferencesIdentifier(value objc.IObject /* cross-framework: NSString */)
	ResourceLoadDelegate() unsafe.Pointer
	SetResourceLoadDelegate(value unsafe.Pointer)
	SelectedDOMRange() IDOMRange
	SelectedFrame() IWebFrame
	SelectionAffinity() SelectionAffinity /* not a class type */
	ShouldCloseWithWindow() bool
	SetShouldCloseWithWindow(value bool)
	ShouldUpdateWhileOffscreen() bool
	SetShouldUpdateWhileOffscreen(value bool)
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	SpellCheckerDocumentTag() int
	SupportsTextEncoding() bool
	TextSizeMultiplier() float32
	SetTextSizeMultiplier(value float32)
	TypingStyle() IDOMCSSStyleDeclaration
	SetTypingStyle(value IDOMCSSStyleDeclaration)
	UIDelegate() unsafe.Pointer
	SetUIDelegate(value unsafe.Pointer)
	UndoManager() foundation.UndoManager
	WindowScriptObject() IWebScriptObject
	WebActionNavigationTypeKey() objc.IObject /* cross-framework: NSString */
	Autosaves() bool
	SetAutosaves(value bool)
	IsContinuousSpellCheckingEnabled() bool
	SetIsContinuousSpellCheckingEnabled(value bool)
	IsEditable() bool
	SetIsEditable(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebView */
	// methods:
	AlignCenter(sender objc.IObject)
	AlignJustified(sender objc.IObject)
	AlignLeft(sender objc.IObject)
	AlignRight(sender objc.IObject)
	ApplyStyle(style IDOMCSSStyleDeclaration)
	ChangeAttributes(sender objc.IObject)
	ChangeColor(sender objc.IObject)
	ChangeDocumentBackgroundColor(sender objc.IObject)
	ChangeFont(sender objc.IObject)
	CheckSpelling(sender objc.IObject)
	ComputedStyleForElementPseudoElement(element IDOMElement, pseudoElement objc.IObject /* cross-framework: NSString */) IDOMCSSStyleDeclaration
	Copy(sender objc.IObject)
	CopyFont(sender objc.IObject)
	Cut(sender objc.IObject)
	Delete(sender objc.IObject)
	DeleteSelection()
	EditableDOMRangeForPoint(point objc.IObject /* cross-framework: Point */) IDOMRange
	GoBackWithSender(sender objc.IObject)
	GoForwardWithSender(sender objc.IObject)
	MakeTextLarger(sender objc.IObject)
	MakeTextSmaller(sender objc.IObject)
	MakeTextStandardSize(sender objc.IObject)
	MoveToBeginningOfSentence(sender objc.IObject)
	MoveToBeginningOfSentenceAndModifySelection(sender objc.IObject)
	MoveToEndOfSentence(sender objc.IObject)
	MoveToEndOfSentenceAndModifySelection(sender objc.IObject)
	OverWrite(sender objc.IObject)
	Paste(sender objc.IObject)
	PasteAsPlainText(sender objc.IObject)
	PasteAsRichText(sender objc.IObject)
	PasteFont(sender objc.IObject)
	PerformFindPanelAction(sender objc.IObject)
	Reload(sender objc.IObject)
	ReloadFromOrigin(sender objc.IObject)
	ReplaceSelectionWithArchive(archive IWebArchive)
	ReplaceSelectionWithNode(node IDOMNode)
	ReplaceSelectionWithMarkupString(markupString objc.IObject /* cross-framework: NSString */)
	ReplaceSelectionWithText(text objc.IObject /* cross-framework: NSString */)
	SelectSentence(sender objc.IObject)
	SetSelectedDOMRangeAffinity(range_ IDOMRange, selectionAffinity SelectionAffinity /* not a class type */)
	ShowGuessPanel(sender objc.IObject)
	StartSpeaking(sender objc.IObject)
	StopLoading(sender objc.IObject)
	StopSpeaking(sender objc.IObject)
	StyleDeclarationWithText(text objc.IObject /* cross-framework: NSString */) IDOMCSSStyleDeclaration
	TakeStringURLFrom(sender objc.IObject)
	ToggleContinuousSpellChecking(sender objc.IObject)
	ToggleSmartInsertDelete(sender objc.IObject)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebView */
// Alloc allocates a new instance without initialization.
func (wc _WebViewClass) Alloc() WebView {
	rv := objc.Send[WebView](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebView */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebView */

// Initializes the receiver with a frame rectangle, frame name, and group name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/init(frame:frameName:groupName:)
func NewWebViewWithFrameFrameNameGroupName(frame objc.IObject /* cross-framework: Rect */, frameName objc.IObject /* cross-framework: NSString */, groupName objc.IObject /* cross-framework: NSString */) WebView {
	instance := getWebViewClass().Alloc()
	rv := objc.Send[WebView](instance.ID, objc.Sel("initWithFrame:frameName:groupName:"), frame, frameName, groupName)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewWebViewWithFrameFrameNameGroupName */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebView */

// Returns whether the receiver can display content of a given MIME type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canShowMIMEType(_:)
func (wc _WebViewClass) CanShowMIMEType(MIMEType objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("canShowMIMEType:"), MIMEType)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=CanShowMIMEType) */

// Returns whether the receiver interprets a MIME type as HTML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canShowMIMEType(asHTML:)
func (wc _WebViewClass) CanShowMIMETypeAsHTML(MIMEType objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("canShowMIMETypeAsHTML:"), MIMEType)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=CanShowMIMETypeAsHTML) */

// Returns a list of MIME types that WebKit renders as HTML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mimeTypesShownAsHTML()
func (wc _WebViewClass) MIMETypesShownAsHTML() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(wc.class), objc.Sel("MIMETypesShownAsHTML"))
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=MIMETypesShownAsHTML) */

// Specifies the view and representation objects to be used for specific MIME types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/registerClass(_:representationClass:forMIMEType:)
func (wc _WebViewClass) RegisterViewClassRepresentationClassForMIMEType(viewClass objc.Class, representationClass objc.Class, MIMEType objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("registerViewClass:representationClass:forMIMEType:"), viewClass, representationClass, MIMEType)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterViewClassRepresentationClassForMIMEType) */

// Adds the specified URL scheme to the list of local schemes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/registerURLScheme(asLocal:)
func (wc _WebViewClass) RegisterURLSchemeAsLocal(scheme objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("registerURLSchemeAsLocal:"), scheme)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterURLSchemeAsLocal) */

// Sets the MIME types that WebKit attempts to render as HTML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/setMIMETypesShownAsHTML(_:)
func (wc _WebViewClass) SetMIMETypesShownAsHTML(MIMETypes objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("setMIMETypesShownAsHTML:"), MIMETypes)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=SetMIMETypesShownAsHTML) */

// Returns a URL from the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/url(from:)
func (wc _WebViewClass) URLFromPasteboard(pasteboard appkit.Pasteboard) foundation.URL {
	rv := objc.Send[foundation.URL](objc.ID(wc.class), objc.Sel("URLFromPasteboard:"), pasteboard)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=URLFromPasteboard) */

// Returns the title of a URL from the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/urlTitle(from:)
func (wc _WebViewClass) URLTitleFromPasteboard(pasteboard appkit.Pasteboard) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(wc.class), objc.Sel("URLTitleFromPasteboard:"), pasteboard)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=URLTitleFromPasteboard) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebView */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebView */

// An action method that applies center alignment to selected content or all content if there’s no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/alignCenter(_:)
func (w_ WebView) AlignCenter(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("alignCenter:"), sender)
} /* debug [instance_methods/method]: AlignCenter */

// An action method that applies full justification to selected content or all content if there’s no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/alignJustified(_:)
func (w_ WebView) AlignJustified(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("alignJustified:"), sender)
} /* debug [instance_methods/method]: AlignJustified */

// An action method that applies left justification to selected content or all content if there’s no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/alignLeft(_:)
func (w_ WebView) AlignLeft(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("alignLeft:"), sender)
} /* debug [instance_methods/method]: AlignLeft */

// An action method that applies right justification to selected content or all content if there is no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/alignRight(_:)
func (w_ WebView) AlignRight(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("alignRight:"), sender)
} /* debug [instance_methods/method]: AlignRight */

// Applies the CSS typing style to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/applyStyle(_:)
func (w_ WebView) ApplyStyle(style IDOMCSSStyleDeclaration) {
	objc.Send[objc.ID](w_.ID, objc.Sel("applyStyle:"), style)
} /* debug [instance_methods/method]: ApplyStyle */

// An action method that changes the attributes of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/changeAttributes(_:)
func (w_ WebView) ChangeAttributes(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("changeAttributes:"), sender)
} /* debug [instance_methods/method]: ChangeAttributes */

// Sets the color of the selected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/changeColor(_:)
func (w_ WebView) ChangeColor(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("changeColor:"), sender)
} /* debug [instance_methods/method]: ChangeColor */

// Sets the background color of the selected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/changeDocumentBackgroundColor(_:)
func (w_ WebView) ChangeDocumentBackgroundColor(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("changeDocumentBackgroundColor:"), sender)
} /* debug [instance_methods/method]: ChangeDocumentBackgroundColor */

// An action method that changes the font of the selection, or all content if there is no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/changeFont(_:)
func (w_ WebView) ChangeFont(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("changeFont:"), sender)
} /* debug [instance_methods/method]: ChangeFont */

// An action method that searches for a misspelled word in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/checkSpelling(_:)
func (w_ WebView) CheckSpelling(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("checkSpelling:"), sender)
} /* debug [instance_methods/method]: CheckSpelling */

// Returns the computed style of an element and its pseudo element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/computedStyle(for:pseudoElement:)
func (w_ WebView) ComputedStyleForElementPseudoElement(element IDOMElement, pseudoElement objc.IObject /* cross-framework: NSString */) IDOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](w_.ID, objc.Sel("computedStyleForElement:pseudoElement:"), element, pseudoElement)
	return rv
} /* debug [instance_methods/method]: ComputedStyleForElementPseudoElement */

// Action method that copies the selected content to the general pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/copy(_:)
func (w_ WebView) Copy(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("copy:"), sender)
} /* debug [instance_methods/method]: Copy */

// An action method that copies font information onto the font pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/copyFont(_:)
func (w_ WebView) CopyFont(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("copyFont:"), sender)
} /* debug [instance_methods/method]: CopyFont */

// An action method that deletes selected content and puts it on the general pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/cut(_:)
func (w_ WebView) Cut(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("cut:"), sender)
} /* debug [instance_methods/method]: Cut */

// An action method that deletes the selected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/delete(_:)
func (w_ WebView) Delete(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("delete:"), sender)
} /* debug [instance_methods/method]: Delete */

// Deletes the receiver’s current selection unless it’s collapsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/deleteSelection()
func (w_ WebView) DeleteSelection() {
	objc.Send[objc.ID](w_.ID, objc.Sel("deleteSelection"))
} /* debug [instance_methods/method]: DeleteSelection */

// Returns the editable DOM object located at a given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/editableDOMRange(for:)
func (w_ WebView) EditableDOMRangeForPoint(point objc.IObject /* cross-framework: Point */) IDOMRange {
	rv := objc.Send[DOMRange](w_.ID, objc.Sel("editableDOMRangeForPoint:"), point)
	return rv
} /* debug [instance_methods/method]: EditableDOMRangeForPoint */

// An action method that loads the previous location in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/goBack(_:)
func (w_ WebView) GoBackWithSender(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goBack:"), sender)
} /* debug [instance_methods/method]: GoBackWithSender */

// An action method that loads the next location in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/goForward(_:)
func (w_ WebView) GoForwardWithSender(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goForward:"), sender)
} /* debug [instance_methods/method]: GoForwardWithSender */

// Action method that increases the text size by one unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/makeTextLarger(_:)
func (w_ WebView) MakeTextLarger(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeTextLarger:"), sender)
} /* debug [instance_methods/method]: MakeTextLarger */

// Action method that reduces the text size by one unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/makeTextSmaller(_:)
func (w_ WebView) MakeTextSmaller(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeTextSmaller:"), sender)
} /* debug [instance_methods/method]: MakeTextSmaller */

// Resets the text size to a multiple of 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/makeTextStandardSize(_:)
func (w_ WebView) MakeTextStandardSize(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeTextStandardSize:"), sender)
} /* debug [instance_methods/method]: MakeTextStandardSize */

// Moves the insertion point to the beginning of the current sentence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/moveToBeginningOfSentence(_:)
func (w_ WebView) MoveToBeginningOfSentence(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("moveToBeginningOfSentence:"), sender)
} /* debug [instance_methods/method]: MoveToBeginningOfSentence */

// Moves the insertion point and extends the selection to the beginning of the current sentence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/moveToBeginningOfSentenceAndModifySelection(_:)
func (w_ WebView) MoveToBeginningOfSentenceAndModifySelection(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("moveToBeginningOfSentenceAndModifySelection:"), sender)
} /* debug [instance_methods/method]: MoveToBeginningOfSentenceAndModifySelection */

// Moves the insertion point to the end of the current sentence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/moveToEndOfSentence(_:)
func (w_ WebView) MoveToEndOfSentence(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("moveToEndOfSentence:"), sender)
} /* debug [instance_methods/method]: MoveToEndOfSentence */

// Moves the insertion point and extends the selection to the end of the current sentence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/moveToEndOfSentenceAndModifySelection(_:)
func (w_ WebView) MoveToEndOfSentenceAndModifySelection(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("moveToEndOfSentenceAndModifySelection:"), sender)
} /* debug [instance_methods/method]: MoveToEndOfSentenceAndModifySelection */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/overWrite(_:)
func (w_ WebView) OverWrite(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("overWrite:"), sender)
} /* debug [instance_methods/method]: OverWrite */

// An action method that pastes content from the pasteboard at the insertion point or over the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/paste(_:)
func (w_ WebView) Paste(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("paste:"), sender)
} /* debug [instance_methods/method]: Paste */

// An action method that pastes pasteboard content as plain text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/pasteAsPlainText(_:)
func (w_ WebView) PasteAsPlainText(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("pasteAsPlainText:"), sender)
} /* debug [instance_methods/method]: PasteAsPlainText */

// An action method that pastes pasteboard content into the receiver as rich text, maintaining its attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/pasteAsRichText(_:)
func (w_ WebView) PasteAsRichText(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("pasteAsRichText:"), sender)
} /* debug [instance_methods/method]: PasteAsRichText */

// An action method that pastes font information from the font pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/pasteFont(_:)
func (w_ WebView) PasteFont(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("pasteFont:"), sender)
} /* debug [instance_methods/method]: PasteFont */

// An action method that opens the Find menu and Find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/performFindPanelAction(_:)
func (w_ WebView) PerformFindPanelAction(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performFindPanelAction:"), sender)
} /* debug [instance_methods/method]: PerformFindPanelAction */

// An action method that reloads the current page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/reload(_:)
func (w_ WebView) Reload(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("reload:"), sender)
} /* debug [instance_methods/method]: Reload */

// Action method that performs an end-to-end revalidation using cache-validating conditionals if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/reloadFromOrigin(_:)
func (w_ WebView) ReloadFromOrigin(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("reloadFromOrigin:"), sender)
} /* debug [instance_methods/method]: ReloadFromOrigin */

// Replaces the current selection with an archive’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/replaceSelection(with:)-3vj8l
func (w_ WebView) ReplaceSelectionWithArchive(archive IWebArchive) {
	objc.Send[objc.ID](w_.ID, objc.Sel("replaceSelectionWithArchive:"), archive)
} /* debug [instance_methods/method]: ReplaceSelectionWithArchive */

// Replaces the receiver’s current selection with the specified DOM node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/replaceSelection(with:)-5px9m
func (w_ WebView) ReplaceSelectionWithNode(node IDOMNode) {
	objc.Send[objc.ID](w_.ID, objc.Sel("replaceSelectionWithNode:"), node)
} /* debug [instance_methods/method]: ReplaceSelectionWithNode */

// Replaces the current selection with mixed text and markup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/replaceSelection(withMarkupString:)
func (w_ WebView) ReplaceSelectionWithMarkupString(markupString objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("replaceSelectionWithMarkupString:"), markupString)
} /* debug [instance_methods/method]: ReplaceSelectionWithMarkupString */

// Replaces the current selection with a string of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/replaceSelection(withText:)
func (w_ WebView) ReplaceSelectionWithText(text objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("replaceSelectionWithText:"), text)
} /* debug [instance_methods/method]: ReplaceSelectionWithText */

// Selects the entire sentence around the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/selectSentence(_:)
func (w_ WebView) SelectSentence(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectSentence:"), sender)
} /* debug [instance_methods/method]: SelectSentence */

// Selects a range of nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/setSelectedDOMRange(_:affinity:)
func (w_ WebView) SetSelectedDOMRangeAffinity(range_ IDOMRange, selectionAffinity SelectionAffinity /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectedDOMRange:affinity:"), range_, selectionAffinity)
} /* debug [instance_methods/method]: SetSelectedDOMRangeAffinity */

// An action method that shows a spelling correction panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/showGuessPanel(_:)
func (w_ WebView) ShowGuessPanel(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("showGuessPanel:"), sender)
} /* debug [instance_methods/method]: ShowGuessPanel */

// An action method that starts speaking the selected text or all text if there’s no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/startSpeaking(_:)
func (w_ WebView) StartSpeaking(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("startSpeaking:"), sender)
} /* debug [instance_methods/method]: StartSpeaking */

// An action method that stops the loading of any web frame content managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/stopLoading(_:)
func (w_ WebView) StopLoading(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopLoading:"), sender)
} /* debug [instance_methods/method]: StopLoading */

// An action method that stops speaking that is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/stopSpeaking(_:)
func (w_ WebView) StopSpeaking(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopSpeaking:"), sender)
} /* debug [instance_methods/method]: StopSpeaking */

// Returns the CSS style declaration for the specified text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/styleDeclaration(withText:)
func (w_ WebView) StyleDeclarationWithText(text objc.IObject /* cross-framework: NSString */) IDOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](w_.ID, objc.Sel("styleDeclarationWithText:"), text)
	return rv
} /* debug [instance_methods/method]: StyleDeclarationWithText */

// Sets the receiver’s current location by obtaining a URL string from the sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/takeStringURLFrom(_:)
func (w_ WebView) TakeStringURLFrom(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("takeStringURLFrom:"), sender)
} /* debug [instance_methods/method]: TakeStringURLFrom */

// Toggles whether continuous spell checking is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/toggleContinuousSpellChecking(_:)
func (w_ WebView) ToggleContinuousSpellChecking(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleContinuousSpellChecking:"), sender)
} /* debug [instance_methods/method]: ToggleContinuousSpellChecking */

// Toggles whether spaces around selected words are inserted or deleted to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/toggleSmartInsertDelete(_:)
func (w_ WebView) ToggleSmartInsertDelete(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleSmartInsertDelete:"), sender)
} /* debug [instance_methods/method]: ToggleSmartInsertDelete */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebView */

// The receiver’s application name that is used in the user-agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/applicationNameForUserAgent
func (w_ WebView) ApplicationNameForUserAgent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("applicationNameForUserAgent"))
	return rv
} /* debug [instance_properties/getter]: applicationNameForUserAgent */

// The receiver’s application name that is used in the user-agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/applicationNameForUserAgent
func (w_ WebView) SetApplicationNameForUserAgent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setApplicationNameForUserAgent:"), value)
} /* debug [instance_properties/setter]: applicationNameForUserAgent */

// The receiver’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/backForwardList
func (w_ WebView) BackForwardList() IWebBackForwardList {
	rv := objc.Send[WebBackForwardList](w_.ID, objc.Sel("backForwardList"))
	return rv
} /* debug [instance_properties/getter]: backForwardList */

// A Boolean that indicates whether the previous location can be loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canGoBack
func (w_ WebView) CanGoBack() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoBack"))
	return rv
} /* debug [instance_properties/getter]: canGoBack */

// A Boolean that indicates whether the next location can be loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canGoForward
func (w_ WebView) CanGoForward() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoForward"))
	return rv
} /* debug [instance_properties/getter]: canGoForward */

// A Boolean that indicates whether the text can be made larger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canMakeTextLarger
func (w_ WebView) CanMakeTextLarger() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextLarger"))
	return rv
} /* debug [instance_properties/getter]: canMakeTextLarger */

// A Boolean that indicates whether the text can be made smaller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canMakeTextSmaller
func (w_ WebView) CanMakeTextSmaller() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextSmaller"))
	return rv
} /* debug [instance_properties/getter]: canMakeTextSmaller */

// A Boolean that indicates whether the current text size is a multiple of 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/canMakeTextStandardSize
func (w_ WebView) CanMakeTextStandardSize() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canMakeTextStandardSize"))
	return rv
} /* debug [instance_properties/getter]: canMakeTextStandardSize */

// The custom text encoding name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/customTextEncodingName
func (w_ WebView) CustomTextEncodingName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("customTextEncodingName"))
	return rv
} /* debug [instance_properties/getter]: customTextEncodingName */

// The custom text encoding name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/customTextEncodingName
func (w_ WebView) SetCustomTextEncodingName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomTextEncodingName:"), value)
} /* debug [instance_properties/setter]: customTextEncodingName */

// The receiver’s custom user-agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/customUserAgent
func (w_ WebView) CustomUserAgent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("customUserAgent"))
	return rv
} /* debug [instance_properties/getter]: customUserAgent */

// The receiver’s custom user-agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/customUserAgent
func (w_ WebView) SetCustomUserAgent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomUserAgent:"), value)
} /* debug [instance_properties/setter]: customUserAgent */

// The receiver’s download delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/downloadDelegate
func (w_ WebView) DownloadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("downloadDelegate"))
	return rv
} /* debug [instance_properties/getter]: downloadDelegate */

// The receiver’s download delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/downloadDelegate
func (w_ WebView) SetDownloadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDownloadDelegate:"), value)
} /* debug [instance_properties/setter]: downloadDelegate */

// A Boolean that indicates whether the web view draws a background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/drawsBackground
func (w_ WebView) DrawsBackground() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("drawsBackground"))
	return rv
} /* debug [instance_properties/getter]: drawsBackground */

// A Boolean that indicates whether the web view draws a background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/drawsBackground
func (w_ WebView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDrawsBackground:"), value)
} /* debug [instance_properties/setter]: drawsBackground */

// The receiver’s editing delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/editingDelegate
func (w_ WebView) EditingDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("editingDelegate"))
	return rv
} /* debug [instance_properties/getter]: editingDelegate */

// The receiver’s editing delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/editingDelegate
func (w_ WebView) SetEditingDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEditingDelegate:"), value)
} /* debug [instance_properties/setter]: editingDelegate */

// An estimate, as a percentage, of the amount of content that is currently loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/estimatedProgress
func (w_ WebView) EstimatedProgress() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("estimatedProgress"))
	return rv
} /* debug [instance_properties/getter]: estimatedProgress */

// The receiver’s frame load delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/frameLoadDelegate
func (w_ WebView) FrameLoadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("frameLoadDelegate"))
	return rv
} /* debug [instance_properties/getter]: frameLoadDelegate */

// The receiver’s frame load delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/frameLoadDelegate
func (w_ WebView) SetFrameLoadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameLoadDelegate:"), value)
} /* debug [instance_properties/setter]: frameLoadDelegate */

// The receiver’s group name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/groupName
func (w_ WebView) GroupName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("groupName"))
	return rv
} /* debug [instance_properties/getter]: groupName */

// The receiver’s group name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/groupName
func (w_ WebView) SetGroupName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGroupName:"), value)
} /* debug [instance_properties/setter]: groupName */

// The receiver’s host window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/hostWindow
func (w_ WebView) HostWindow() appkit.Window {
	rv := objc.Send[appkit.Window](w_.ID, objc.Sel("hostWindow"))
	return rv
} /* debug [instance_properties/getter]: hostWindow */

// The receiver’s host window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/hostWindow
func (w_ WebView) SetHostWindow(value appkit.Window) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHostWindow:"), value)
} /* debug [instance_properties/setter]: hostWindow */

// A Boolean that indicates whether the web view has continuous spell-checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isContinuousSpellCheckingEnabled
func (w_ WebView) ContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("continuousSpellCheckingEnabled"))
	return rv
} /* debug [instance_properties/getter]: continuousSpellCheckingEnabled */

// A Boolean that indicates whether the web view has continuous spell-checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isContinuousSpellCheckingEnabled
func (w_ WebView) SetContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContinuousSpellCheckingEnabled:"), value)
} /* debug [instance_properties/setter]: continuousSpellCheckingEnabled */

// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isEditable
func (w_ WebView) Editable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("editable"))
	return rv
} /* debug [instance_properties/getter]: editable */

// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isEditable
func (w_ WebView) SetEditable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEditable:"), value)
} /* debug [instance_properties/setter]: editable */

// A Boolean that indicates whether the web view is loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/isLoading
func (w_ WebView) Loading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("loading"))
	return rv
} /* debug [instance_properties/getter]: loading */

// The main frame, the root of the web frame hierarchy for this page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mainFrame
func (w_ WebView) MainFrame() IWebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("mainFrame"))
	return rv
} /* debug [instance_properties/getter]: mainFrame */

// The DOM document for the main frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mainFrameDocument
func (w_ WebView) MainFrameDocument() IDOMDocument {
	rv := objc.Send[DOMDocument](w_.ID, objc.Sel("mainFrameDocument"))
	return rv
} /* debug [instance_properties/getter]: mainFrameDocument */

// The site’s favicon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mainFrameIcon
func (w_ WebView) MainFrameIcon() appkit.Image {
	rv := objc.Send[appkit.Image](w_.ID, objc.Sel("mainFrameIcon"))
	return rv
} /* debug [instance_properties/getter]: mainFrameIcon */

// The HTML title of the loaded page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mainFrameTitle
func (w_ WebView) MainFrameTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("mainFrameTitle"))
	return rv
} /* debug [instance_properties/getter]: mainFrameTitle */

// The URL that the main frame loads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mainFrameURL
func (w_ WebView) MainFrameURL() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("mainFrameURL"))
	return rv
} /* debug [instance_properties/getter]: mainFrameURL */

// The URL that the main frame loads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mainFrameURL
func (w_ WebView) SetMainFrameURL(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMainFrameURL:"), value)
} /* debug [instance_properties/setter]: mainFrameURL */

// A Boolean that indicates whether the selection is maintained when focus is lost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/maintainsInactiveSelection
func (w_ WebView) MaintainsInactiveSelection() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("maintainsInactiveSelection"))
	return rv
} /* debug [instance_properties/getter]: maintainsInactiveSelection */

// The receiver’s CSS media property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mediaStyle
func (w_ WebView) MediaStyle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("mediaStyle"))
	return rv
} /* debug [instance_properties/getter]: mediaStyle */

// The receiver’s CSS media property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/mediaStyle
func (w_ WebView) SetMediaStyle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaStyle:"), value)
} /* debug [instance_properties/setter]: mediaStyle */

// An array of pasteboard types that can be used for the current selection of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/pasteboardTypesForSelection
func (w_ WebView) PasteboardTypesForSelection() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](w_.ID, objc.Sel("pasteboardTypesForSelection"))
	return rv
} /* debug [instance_properties/getter]: pasteboardTypesForSelection */

// The receiver’s policy delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/policyDelegate
func (w_ WebView) PolicyDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("policyDelegate"))
	return rv
} /* debug [instance_properties/getter]: policyDelegate */

// The receiver’s policy delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/policyDelegate
func (w_ WebView) SetPolicyDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPolicyDelegate:"), value)
} /* debug [instance_properties/setter]: policyDelegate */

// The receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/preferences
func (w_ WebView) Preferences() IWebPreferences {
	rv := objc.Send[WebPreferences](w_.ID, objc.Sel("preferences"))
	return rv
} /* debug [instance_properties/getter]: preferences */

// The receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/preferences
func (w_ WebView) SetPreferences(value IWebPreferences) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferences:"), value)
} /* debug [instance_properties/setter]: preferences */

// The identifier of the receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/preferencesIdentifier
func (w_ WebView) PreferencesIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("preferencesIdentifier"))
	return rv
} /* debug [instance_properties/getter]: preferencesIdentifier */

// The identifier of the receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/preferencesIdentifier
func (w_ WebView) SetPreferencesIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferencesIdentifier:"), value)
} /* debug [instance_properties/setter]: preferencesIdentifier */

// The receiver’s resource load delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/resourceLoadDelegate
func (w_ WebView) ResourceLoadDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("resourceLoadDelegate"))
	return rv
} /* debug [instance_properties/getter]: resourceLoadDelegate */

// The receiver’s resource load delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/resourceLoadDelegate
func (w_ WebView) SetResourceLoadDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setResourceLoadDelegate:"), value)
} /* debug [instance_properties/setter]: resourceLoadDelegate */

// The range of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/selectedDOMRange
func (w_ WebView) SelectedDOMRange() IDOMRange {
	rv := objc.Send[DOMRange](w_.ID, objc.Sel("selectedDOMRange"))
	return rv
} /* debug [instance_properties/getter]: selectedDOMRange */

// The frame with the active selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/selectedFrame
func (w_ WebView) SelectedFrame() IWebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("selectedFrame"))
	return rv
} /* debug [instance_properties/getter]: selectedFrame */

// The current selection affinity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/selectionAffinity
func (w_ WebView) SelectionAffinity() SelectionAffinity /* not a class type */ {
	rv := objc.Send[SelectionAffinity](w_.ID, objc.Sel("selectionAffinity"))
	return rv
} /* debug [instance_properties/getter]: selectionAffinity */

// A Boolean that indicates whether the web view should close when its window or host window closes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/shouldCloseWithWindow
func (w_ WebView) ShouldCloseWithWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldCloseWithWindow"))
	return rv
} /* debug [instance_properties/getter]: shouldCloseWithWindow */

// A Boolean that indicates whether the web view should close when its window or host window closes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/shouldCloseWithWindow
func (w_ WebView) SetShouldCloseWithWindow(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldCloseWithWindow:"), value)
} /* debug [instance_properties/setter]: shouldCloseWithWindow */

// A Boolean that inidicates whether the web view should update even when it is not in a window that is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/shouldUpdateWhileOffscreen
func (w_ WebView) ShouldUpdateWhileOffscreen() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldUpdateWhileOffscreen"))
	return rv
} /* debug [instance_properties/getter]: shouldUpdateWhileOffscreen */

// A Boolean that inidicates whether the web view should update even when it is not in a window that is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/shouldUpdateWhileOffscreen
func (w_ WebView) SetShouldUpdateWhileOffscreen(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldUpdateWhileOffscreen:"), value)
} /* debug [instance_properties/setter]: shouldUpdateWhileOffscreen */

// A Boolean that indicates whether smart-space insertion and deletion is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/smartInsertDeleteEnabled
func (w_ WebView) SmartInsertDeleteEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
} /* debug [instance_properties/getter]: smartInsertDeleteEnabled */

// A Boolean that indicates whether smart-space insertion and deletion is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/smartInsertDeleteEnabled
func (w_ WebView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
} /* debug [instance_properties/setter]: smartInsertDeleteEnabled */

// The spell-checker document tag for this document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/spellCheckerDocumentTag
func (w_ WebView) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](w_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
} /* debug [instance_properties/getter]: spellCheckerDocumentTag */

// A Boolean that indicates whether the document view supports different text encodings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/supportsTextEncoding
func (w_ WebView) SupportsTextEncoding() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("supportsTextEncoding"))
	return rv
} /* debug [instance_properties/getter]: supportsTextEncoding */

// The font size multiplier for text displayed in web frame view objects managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/textSizeMultiplier
func (w_ WebView) TextSizeMultiplier() float32 {
	rv := objc.Send[float32](w_.ID, objc.Sel("textSizeMultiplier"))
	return rv
} /* debug [instance_properties/getter]: textSizeMultiplier */

// The font size multiplier for text displayed in web frame view objects managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/textSizeMultiplier
func (w_ WebView) SetTextSizeMultiplier(value float32) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTextSizeMultiplier:"), value)
} /* debug [instance_properties/setter]: textSizeMultiplier */

// The receiver’s CSS typing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/typingStyle
func (w_ WebView) TypingStyle() IDOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](w_.ID, objc.Sel("typingStyle"))
	return rv
} /* debug [instance_properties/getter]: typingStyle */

// The receiver’s CSS typing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/typingStyle
func (w_ WebView) SetTypingStyle(value IDOMCSSStyleDeclaration) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTypingStyle:"), value)
} /* debug [instance_properties/setter]: typingStyle */

// The receiver’s user interface delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/uiDelegate
func (w_ WebView) UIDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("UIDelegate"))
	return rv
} /* debug [instance_properties/getter]: UIDelegate */

// The receiver’s user interface delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/uiDelegate
func (w_ WebView) SetUIDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUIDelegate:"), value)
} /* debug [instance_properties/setter]: UIDelegate */

// The receiver’s undo manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/undoManager
func (w_ WebView) UndoManager() foundation.UndoManager {
	rv := objc.Send[foundation.UndoManager](w_.ID, objc.Sel("undoManager"))
	return rv
} /* debug [instance_properties/getter]: undoManager */

// The receiver’s window object from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebView-swift.class/windowScriptObject
func (w_ WebView) WindowScriptObject() IWebScriptObject {
	rv := objc.Send[WebScriptObject](w_.ID, objc.Sel("windowScriptObject"))
	return rv
} /* debug [instance_properties/getter]: windowScriptObject */

// The navigation type of the action. Can be any of the values defined in
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webactionnavigationtypekey
func (w_ WebView) WebActionNavigationTypeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("WebActionNavigationTypeKey"))
	return rv
} /* debug [instance_properties/getter]: WebActionNavigationTypeKey */

// A Boolean that indicates whether or not the receiver’s attributes are automatically stored in the user defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/autosaves
func (w_ WebView) Autosaves() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autosaves"))
	return rv
} /* debug [instance_properties/getter]: autosaves */

// A Boolean that indicates whether or not the receiver’s attributes are automatically stored in the user defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/autosaves
func (w_ WebView) SetAutosaves(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutosaves:"), value)
} /* debug [instance_properties/setter]: autosaves */

// A Boolean that indicates whether the web view has continuous spell-checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iscontinuousspellcheckingenabled
func (w_ WebView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
} /* debug [instance_properties/getter]: isContinuousSpellCheckingEnabled */

// A Boolean that indicates whether the web view has continuous spell-checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iscontinuousspellcheckingenabled
func (w_ WebView) SetIsContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsContinuousSpellCheckingEnabled:"), value)
} /* debug [instance_properties/setter]: isContinuousSpellCheckingEnabled */

// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iseditable
func (w_ WebView) IsEditable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isEditable"))
	return rv
} /* debug [instance_properties/getter]: isEditable */

// A Boolean that indicates whether the user is allowed to edit the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/iseditable
func (w_ WebView) SetIsEditable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsEditable:"), value)
} /* debug [instance_properties/setter]: isEditable */

// A Boolean that indicates whether the web view is loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/isloading
func (w_ WebView) IsLoading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoading"))
	return rv
} /* debug [instance_properties/getter]: isLoading */

// A Boolean that indicates whether the web view is loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/isloading
func (w_ WebView) SetIsLoading(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoading:"), value)
} /* debug [instance_properties/setter]: isLoading */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WebView */
