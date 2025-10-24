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

/* debug [class.gen.go]: Generating class WKWebView */


/* debug [class_header]: Header for WKWebView */
// The class instance for the [WebView] class.
var (
	WebViewClass     _WebViewClass
	WebViewClassOnce sync.Once
)

func getWebViewClass() _WebViewClass {
	WebViewClassOnce.Do(func() {
		WebViewClass = _WebViewClass{objc.GetClass("WKWebView")}
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
	AllowsBackForwardNavigationGestures() bool
	SetAllowsBackForwardNavigationGestures(value bool)
	AllowsLinkPreview() bool
	SetAllowsLinkPreview(value bool)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	BackForwardList() IWKBackForwardList
	CameraCaptureState() MediaCaptureState
	CanGoBack() bool
	CanGoForward() bool
	CertificateChain() objc.IObject /* cross-framework: NSArray */
	Configuration() IWKWebViewConfiguration
	CustomUserAgent() objc.IObject /* cross-framework: NSString */
	SetCustomUserAgent(value objc.IObject /* cross-framework: NSString */)
	EstimatedProgress() float64
	FullscreenState() FullscreenState
	HasOnlySecureContent() bool
	InteractionState() objc.ID
	SetInteractionState(value objc.ID)
	IsBlockedByScreenTime() bool
	Inspectable() bool
	SetInspectable(value bool)
	Loading() bool
	WritingToolsActive() bool
	Magnification() float64
	SetMagnification(value float64)
	MaximumViewportInset() foundation.EdgeInsets
	MediaType() objc.IObject /* cross-framework: NSString */
	SetMediaType(value objc.IObject /* cross-framework: NSString */)
	MicrophoneCaptureState() MediaCaptureState
	MinimumViewportInset() foundation.EdgeInsets
	NavigationDelegate() unsafe.Pointer
	SetNavigationDelegate(value unsafe.Pointer)
	ObscuredContentInsets() foundation.EdgeInsets
	SetObscuredContentInsets(value foundation.EdgeInsets)
	PageZoom() float64
	SetPageZoom(value float64)
	ServerTrust() objectivec.IObject
	ThemeColor() appkit.Color
	Title() objc.IObject /* cross-framework: NSString */
	UIDelegate() unsafe.Pointer
	SetUIDelegate(value unsafe.Pointer)
	UnderPageBackgroundColor() appkit.Color
	SetUnderPageBackgroundColor(value appkit.Color)
	URL() objc.IObject /* cross-framework: NSURL */
	IsFindInteractionEnabled() bool
	SetIsFindInteractionEnabled(value bool)
	IsInspectable() bool
	SetIsInspectable(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
	IsWritingToolsActive() bool
	SetIsWritingToolsActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebView */
	// methods:
	CloseAllMediaPresentationsWithCompletionHandler(completionHandler func())
	CreatePDFWithConfigurationCompletionHandler(pdfConfiguration IWKPDFConfiguration, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	CreateWebArchiveDataWithCompletionHandler(completionHandler func(unsafe.Pointer, unsafe.Pointer))
	EvaluateJavaScriptCompletionHandler(javaScriptString objc.IObject /* cross-framework: NSString */, completionHandler func(objc.ID, unsafe.Pointer))
	FetchDataOfTypesCompletionHandler(dataTypes WebViewDataType, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	FindStringWithConfigurationCompletionHandler(string_ objc.IObject /* cross-framework: NSString */, configuration IWKFindConfiguration, completionHandler func(unsafe.Pointer))
	GoToBackForwardListItem(item IWKBackForwardListItem) INavigation
	GoBack() INavigation
	GoBackWithSender(sender objc.IObject)
	GoForward() INavigation
	GoForwardWithSender(sender objc.IObject)
	LoadRequest(request foundation.URLRequest) INavigation
	LoadDataMIMETypeCharacterEncodingNameBaseURL(data objc.IObject /* cross-framework: NSData */, MIMEType objc.IObject /* cross-framework: NSString */, characterEncodingName objc.IObject /* cross-framework: NSString */, baseURL objc.IObject /* cross-framework: NSURL */) INavigation
	LoadFileRequestAllowingReadAccessToURL(request foundation.URLRequest, readAccessURL objc.IObject /* cross-framework: NSURL */) INavigation
	LoadFileURLAllowingReadAccessToURL(URL objc.IObject /* cross-framework: NSURL */, readAccessURL objc.IObject /* cross-framework: NSURL */) INavigation
	LoadHTMLStringBaseURL(string_ objc.IObject /* cross-framework: NSString */, baseURL objc.IObject /* cross-framework: NSURL */) INavigation
	LoadSimulatedRequestResponseResponseData(request foundation.URLRequest, response foundation.URLResponse, data objc.IObject /* cross-framework: NSData */) INavigation
	LoadSimulatedRequestResponseHTMLString(request foundation.URLRequest, string_ objc.IObject /* cross-framework: NSString */) INavigation
	PauseAllMediaPlaybackWithCompletionHandler(completionHandler func())
	PrintOperationWithPrintInfo(printInfo appkit.PrintInfo) appkit.PrintOperation
	Reload() INavigation
	ReloadWithSender(sender objc.IObject)
	ReloadFromOrigin() INavigation
	ReloadFromOriginWithSender(sender objc.IObject)
	RequestMediaPlaybackStateWithCompletionHandler(completionHandler func(unsafe.Pointer))
	RestoreDataCompletionHandler(data objc.IObject /* cross-framework: NSData */, completionHandler func(unsafe.Pointer))
	ResumeDownloadFromResumeDataCompletionHandler(resumeData objc.IObject /* cross-framework: NSData */, completionHandler func(unsafe.Pointer))
	SetAllMediaPlaybackSuspendedCompletionHandler(suspended bool, completionHandler func())
	SetCameraCaptureStateCompletionHandler(state MediaCaptureState, completionHandler func())
	SetMagnificationCenteredAtPoint(magnification float64, point corefoundation.CGPoint)
	SetMicrophoneCaptureStateCompletionHandler(state MediaCaptureState, completionHandler func())
	SetMinimumViewportInsetMaximumViewportInset(minimumViewportInset foundation.EdgeInsets, maximumViewportInset foundation.EdgeInsets)
	StartDownloadUsingRequestCompletionHandler(request foundation.URLRequest, completionHandler func(unsafe.Pointer))
	StopLoading()
	StopLoadingWithSender(sender objc.IObject)
	TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration IWKSnapshotConfiguration, completionHandler func(unsafe.Pointer, unsafe.Pointer))
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
// An object that displays interactive web content, such as for an in-app browser.
//
// A object is a platform-native view that you use to incorporate web content seamlessly into your app’s UI. A web view supports a full web-browsing experience, and presents HTML, CSS, and JavaScript content alongside your app’s native views. Use it when web technologies satisfy your app’s layout and styling requirements more readily than native views. For example, you might use it when your app’s content changes frequently. A web view offers control over the navigation and user experience through delegate objects. Use the navigation delegate to react when the user clicks links in your web content, or interacts with the content in a way that affects navigation. For example, you might prevent the user from navigating to new content unless specific conditions are met. Use the UI delegate to present native UI elements, such as alerts or contextual menus, in response to interactions with your web content. Embed a object programmatically into your view hierarchy, or add it using Interface Builder. Interface Builder supports many customizations, such as configuring data detectors, media playback, and interaction behaviors. For more extensive customizations, create your web view programmatically using a object. For example, use a web view configuration object to specify handlers for custom URL schemes, manage cookies, and customize preferences for your web content. Before your web view appears onscreen, load content from a web server using a structure or load content directly from a local file or HTML string. The web view automatically loads embedded resources such as images or videos as part of the initial load request. It then renders your content and displays the results inside the view’s bounds rectangle. The following code example shows a view controller that replaces its default view with a custom object. A web view automatically converts telephone numbers that appear in web content to Phone links. When the user taps a Phone link, the Phone app launches and dials the number. Use the object to change the default data detector behavior. You can also use to programmatically set the scale of web content the first time it appears in a web view. Thereafter, the user can change the scale using gestures.


// An object that displays interactive web content, such as for an in-app browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView
type WebView struct {
	appkit.View
}

// WebViewFrom constructs a [WebView] from an unsafe.Pointer.
//
// An object that displays interactive web content, such as for an in-app browser.
func WebViewFrom(ptr unsafe.Pointer) WebView {
	return WebView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebView */

// Returns an object initialized from data in the specified coder object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/init(coder:)
func NewWebViewWithCoder(coder foundation.Coder) WebView {
	instance := getWebViewClass().Alloc()
	rv := objc.Send[WebView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebViewWithCoder */


// Creates a web view and initializes it with the specified frame and configuration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/init(frame:configuration:)
func NewWebViewWithFrameConfiguration(frame corefoundation.CGRect, configuration IWKWebViewConfiguration) WebView {
	instance := getWebViewClass().Alloc()
	rv := objc.Send[WebView](instance.ID, objc.Sel("initWithFrame:configuration:"), frame, configuration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebViewWithFrameConfiguration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebView */

// Returns a Boolean value that indicates whether WebKit natively supports resources with the specified URL scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/handlesURLScheme(_:)
func (wc _WebViewClass) HandlesURLScheme(urlScheme objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("handlesURLScheme:"), urlScheme)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HandlesURLScheme) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebView */

// Closes all media the web view is presenting, including picture-in-picture video and fullscreen video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/closeAllMediaPresentations(completionHandler:)
func (w_ WebView) CloseAllMediaPresentationsWithCompletionHandler(completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("closeAllMediaPresentationsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: CloseAllMediaPresentationsWithCompletionHandler */


// Generates PDF data from the web view’s contents asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/createPDFWithConfiguration:completionHandler:
func (w_ WebView) CreatePDFWithConfigurationCompletionHandler(pdfConfiguration IWKPDFConfiguration, completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("createPDFWithConfiguration:completionHandler:"), pdfConfiguration, completionHandler)
}/* debug [instance_methods/method]: CreatePDFWithConfigurationCompletionHandler */


// Creates a web archive of the web view’s contents asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/createWebArchiveDataWithCompletionHandler:
func (w_ WebView) CreateWebArchiveDataWithCompletionHandler(completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("createWebArchiveDataWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: CreateWebArchiveDataWithCompletionHandler */


// Evaluates the specified JavaScript string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/evaluateJavaScript(_:completionHandler:)
func (w_ WebView) EvaluateJavaScriptCompletionHandler(javaScriptString objc.IObject /* cross-framework: NSString */, completionHandler func(objc.ID, unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("evaluateJavaScript:completionHandler:"), javaScriptString, completionHandler)
}/* debug [instance_methods/method]: EvaluateJavaScriptCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/fetchData(of:completionHandler:)
func (w_ WebView) FetchDataOfTypesCompletionHandler(dataTypes WebViewDataType, completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("fetchDataOfTypes:completionHandler:"), dataTypes, completionHandler)
}/* debug [instance_methods/method]: FetchDataOfTypesCompletionHandler */


// Searches for the specified string in the web view’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/findString:withConfiguration:completionHandler:
func (w_ WebView) FindStringWithConfigurationCompletionHandler(string_ objc.IObject /* cross-framework: NSString */, configuration IWKFindConfiguration, completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("findString:withConfiguration:completionHandler:"), string_, configuration, completionHandler)
}/* debug [instance_methods/method]: FindStringWithConfigurationCompletionHandler */


// Navigates to an item from the back-forward list and sets it as the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/go(to:)
func (w_ WebView) GoToBackForwardListItem(item IWKBackForwardListItem) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("goToBackForwardListItem:"), item)
	return rv
}/* debug [instance_methods/method]: GoToBackForwardListItem */


// Navigates to the back item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/goBack()
func (w_ WebView) GoBack() INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("goBack"))
	return rv
}/* debug [instance_methods/method]: GoBack */


// Navigates to the back item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/goBack(_:)
func (w_ WebView) GoBackWithSender(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goBack:"), sender)
}/* debug [instance_methods/method]: GoBackWithSender */


// Navigates to the forward item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/goForward()
func (w_ WebView) GoForward() INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("goForward"))
	return rv
}/* debug [instance_methods/method]: GoForward */


// Navigates to the forward item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/goForward(_:)
func (w_ WebView) GoForwardWithSender(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goForward:"), sender)
}/* debug [instance_methods/method]: GoForwardWithSender */


// Loads the web content that the specified URL request object references and navigates to that content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/load(_:)
func (w_ WebView) LoadRequest(request foundation.URLRequest) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadRequest:"), request)
	return rv
}/* debug [instance_methods/method]: LoadRequest */


// Loads the content of the specified data object and navigates to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/load(_:mimeType:characterEncodingName:baseURL:)
func (w_ WebView) LoadDataMIMETypeCharacterEncodingNameBaseURL(data objc.IObject /* cross-framework: NSData */, MIMEType objc.IObject /* cross-framework: NSString */, characterEncodingName objc.IObject /* cross-framework: NSString */, baseURL objc.IObject /* cross-framework: NSURL */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadData:MIMEType:characterEncodingName:baseURL:"), data, MIMEType, characterEncodingName, baseURL)
	return rv
}/* debug [instance_methods/method]: LoadDataMIMETypeCharacterEncodingNameBaseURL */


// Loads the web content from the file the URL request object specifies and navigates to that content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadFileRequest(_:allowingReadAccessTo:)
func (w_ WebView) LoadFileRequestAllowingReadAccessToURL(request foundation.URLRequest, readAccessURL objc.IObject /* cross-framework: NSURL */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadFileRequest:allowingReadAccessToURL:"), request, readAccessURL)
	return rv
}/* debug [instance_methods/method]: LoadFileRequestAllowingReadAccessToURL */


// Loads the web content from the specified file and navigates to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadFileURL(_:allowingReadAccessTo:)
func (w_ WebView) LoadFileURLAllowingReadAccessToURL(URL objc.IObject /* cross-framework: NSURL */, readAccessURL objc.IObject /* cross-framework: NSURL */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadFileURL:allowingReadAccessToURL:"), URL, readAccessURL)
	return rv
}/* debug [instance_methods/method]: LoadFileURLAllowingReadAccessToURL */


// Loads the contents of the specified HTML string and navigates to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadHTMLString(_:baseURL:)
func (w_ WebView) LoadHTMLStringBaseURL(string_ objc.IObject /* cross-framework: NSString */, baseURL objc.IObject /* cross-framework: NSURL */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadHTMLString:baseURL:"), string_, baseURL)
	return rv
}/* debug [instance_methods/method]: LoadHTMLStringBaseURL */


// Loads the web content from the data you provide as if the data were the response to the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:response:responseData:)
func (w_ WebView) LoadSimulatedRequestResponseResponseData(request foundation.URLRequest, response foundation.URLResponse, data objc.IObject /* cross-framework: NSData */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadSimulatedRequest:response:responseData:"), request, response, data)
	return rv
}/* debug [instance_methods/method]: LoadSimulatedRequestResponseResponseData */


// Loads the web content from the HTML you provide as if the HTML were the response to the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:responseHTML:)
func (w_ WebView) LoadSimulatedRequestResponseHTMLString(request foundation.URLRequest, string_ objc.IObject /* cross-framework: NSString */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadSimulatedRequest:responseHTMLString:"), request, string_)
	return rv
}/* debug [instance_methods/method]: LoadSimulatedRequestResponseHTMLString */


// Pauses playback of all media in the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/pauseAllMediaPlayback(completionHandler:)
func (w_ WebView) PauseAllMediaPlaybackWithCompletionHandler(completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("pauseAllMediaPlaybackWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: PauseAllMediaPlaybackWithCompletionHandler */


// Returns the print operation object to use when printing the contents of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/printOperation(with:)
func (w_ WebView) PrintOperationWithPrintInfo(printInfo appkit.PrintInfo) appkit.PrintOperation {
	rv := objc.Send[appkit.PrintOperation](w_.ID, objc.Sel("printOperationWithPrintInfo:"), printInfo)
	return rv
}/* debug [instance_methods/method]: PrintOperationWithPrintInfo */


// Reloads the current webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/reload()
func (w_ WebView) Reload() INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("reload"))
	return rv
}/* debug [instance_methods/method]: Reload */


// Reloads the current webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/reload(_:)
func (w_ WebView) ReloadWithSender(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("reload:"), sender)
}/* debug [instance_methods/method]: ReloadWithSender */


// Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/reloadFromOrigin()
func (w_ WebView) ReloadFromOrigin() INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("reloadFromOrigin"))
	return rv
}/* debug [instance_methods/method]: ReloadFromOrigin */


// Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/reloadFromOrigin(_:)
func (w_ WebView) ReloadFromOriginWithSender(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("reloadFromOrigin:"), sender)
}/* debug [instance_methods/method]: ReloadFromOriginWithSender */


// Requests the playback status of media in the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/requestMediaPlaybackState(completionHandler:)
func (w_ WebView) RequestMediaPlaybackStateWithCompletionHandler(completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("requestMediaPlaybackStateWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: RequestMediaPlaybackStateWithCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/restoreData(_:completionHandler:)
func (w_ WebView) RestoreDataCompletionHandler(data objc.IObject /* cross-framework: NSData */, completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("restoreData:completionHandler:"), data, completionHandler)
}/* debug [instance_methods/method]: RestoreDataCompletionHandler */


// Resumes a failed or canceled download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/resumeDownload(fromResumeData:completionHandler:)
func (w_ WebView) ResumeDownloadFromResumeDataCompletionHandler(resumeData objc.IObject /* cross-framework: NSData */, completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("resumeDownloadFromResumeData:completionHandler:"), resumeData, completionHandler)
}/* debug [instance_methods/method]: ResumeDownloadFromResumeDataCompletionHandler */


// Changes whether the webpage is suspending playback of all media in the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setAllMediaPlaybackSuspended(_:completionHandler:)
func (w_ WebView) SetAllMediaPlaybackSuspendedCompletionHandler(suspended bool, completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllMediaPlaybackSuspended:completionHandler:"), suspended, completionHandler)
}/* debug [instance_methods/method]: SetAllMediaPlaybackSuspendedCompletionHandler */


// Changes whether the webpage is using the camera to capture images or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setCameraCaptureState(_:completionHandler:)
func (w_ WebView) SetCameraCaptureStateCompletionHandler(state MediaCaptureState, completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCameraCaptureState:completionHandler:"), state, completionHandler)
}/* debug [instance_methods/method]: SetCameraCaptureStateCompletionHandler */


// Scales the page content and centers the result on the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setMagnification(_:centeredAt:)
func (w_ WebView) SetMagnificationCenteredAtPoint(magnification float64, point corefoundation.CGPoint) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}/* debug [instance_methods/method]: SetMagnificationCenteredAtPoint */


// Changes whether the webpage is using the microphone to capture audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setMicrophoneCaptureState(_:completionHandler:)
func (w_ WebView) SetMicrophoneCaptureStateCompletionHandler(state MediaCaptureState, completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMicrophoneCaptureState:completionHandler:"), state, completionHandler)
}/* debug [instance_methods/method]: SetMicrophoneCaptureStateCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setMinimumViewportInset(_:maximumViewportInset:)
func (w_ WebView) SetMinimumViewportInsetMaximumViewportInset(minimumViewportInset foundation.EdgeInsets, maximumViewportInset foundation.EdgeInsets) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinimumViewportInset:maximumViewportInset:"), minimumViewportInset, maximumViewportInset)
}/* debug [instance_methods/method]: SetMinimumViewportInsetMaximumViewportInset */


// Starts to download the resource at the URL in the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/startDownload(using:completionHandler:)
func (w_ WebView) StartDownloadUsingRequestCompletionHandler(request foundation.URLRequest, completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("startDownloadUsingRequest:completionHandler:"), request, completionHandler)
}/* debug [instance_methods/method]: StartDownloadUsingRequestCompletionHandler */


// Stops loading all resources on the current page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/stopLoading()
func (w_ WebView) StopLoading() {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopLoading"))
}/* debug [instance_methods/method]: StopLoading */


// Stops loading all resources on the current page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/stopLoading(_:)
func (w_ WebView) StopLoadingWithSender(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopLoading:"), sender)
}/* debug [instance_methods/method]: StopLoadingWithSender */


// Generates a platform-native image from the web view’s contents asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/takeSnapshot(with:completionHandler:)
func (w_ WebView) TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration IWKSnapshotConfiguration, completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("takeSnapshotWithConfiguration:completionHandler:"), snapshotConfiguration, completionHandler)
}/* debug [instance_methods/method]: TakeSnapshotWithConfigurationCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebView */

// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsBackForwardNavigationGestures
func (w_ WebView) AllowsBackForwardNavigationGestures() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsBackForwardNavigationGestures"))
	return rv
}/* debug [instance_properties/getter]: allowsBackForwardNavigationGestures */


// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsBackForwardNavigationGestures
func (w_ WebView) SetAllowsBackForwardNavigationGestures(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsBackForwardNavigationGestures:"), value)
}/* debug [instance_properties/setter]: allowsBackForwardNavigationGestures */


// A Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsLinkPreview
func (w_ WebView) AllowsLinkPreview() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsLinkPreview"))
	return rv
}/* debug [instance_properties/getter]: allowsLinkPreview */


// A Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsLinkPreview
func (w_ WebView) SetAllowsLinkPreview(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsLinkPreview:"), value)
}/* debug [instance_properties/setter]: allowsLinkPreview */


// A Boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsMagnification
func (w_ WebView) AllowsMagnification() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsMagnification"))
	return rv
}/* debug [instance_properties/getter]: allowsMagnification */


// A Boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsMagnification
func (w_ WebView) SetAllowsMagnification(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsMagnification:"), value)
}/* debug [instance_properties/setter]: allowsMagnification */


// The web view’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/backForwardList
func (w_ WebView) BackForwardList() IWKBackForwardList {
	rv := objc.Send[BackForwardList](w_.ID, objc.Sel("backForwardList"))
	return rv
}/* debug [instance_properties/getter]: backForwardList */


// An enumeration case that indicates whether the webpage is using the camera to capture images or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/cameraCaptureState
func (w_ WebView) CameraCaptureState() MediaCaptureState {
	rv := objc.Send[MediaCaptureState](w_.ID, objc.Sel("cameraCaptureState"))
	return rv
}/* debug [instance_properties/getter]: cameraCaptureState */


// A Boolean value that indicates whether there is a valid back item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/canGoBack
func (w_ WebView) CanGoBack() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoBack"))
	return rv
}/* debug [instance_properties/getter]: canGoBack */


// A Boolean value that indicates whether there is a valid forward item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/canGoForward
func (w_ WebView) CanGoForward() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoForward"))
	return rv
}/* debug [instance_properties/getter]: canGoForward */


// An array of objects forming the certificate chain for the currently committed navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/certificateChain
func (w_ WebView) CertificateChain() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](w_.ID, objc.Sel("certificateChain"))
	return rv
}/* debug [instance_properties/getter]: certificateChain */


// The object that contains the configuration details for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/configuration
func (w_ WebView) Configuration() IWKWebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](w_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The custom user agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/customUserAgent
func (w_ WebView) CustomUserAgent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("customUserAgent"))
	return rv
}/* debug [instance_properties/getter]: customUserAgent */


// The custom user agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/customUserAgent
func (w_ WebView) SetCustomUserAgent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomUserAgent:"), value)
}/* debug [instance_properties/setter]: customUserAgent */


// An estimate of what fraction of the current navigation has been loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/estimatedProgress
func (w_ WebView) EstimatedProgress() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("estimatedProgress"))
	return rv
}/* debug [instance_properties/getter]: estimatedProgress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/fullscreenState-swift.property
func (w_ WebView) FullscreenState() FullscreenState {
	rv := objc.Send[FullscreenState](w_.ID, objc.Sel("fullscreenState"))
	return rv
}/* debug [instance_properties/getter]: fullscreenState */


// A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/hasOnlySecureContent
func (w_ WebView) HasOnlySecureContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOnlySecureContent"))
	return rv
}/* debug [instance_properties/getter]: hasOnlySecureContent */


// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/interactionState
func (w_ WebView) InteractionState() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("interactionState"))
	return rv
}/* debug [instance_properties/getter]: interactionState */


// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/interactionState
func (w_ WebView) SetInteractionState(value objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInteractionState:"), value)
}/* debug [instance_properties/setter]: interactionState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isBlockedByScreenTime
func (w_ WebView) IsBlockedByScreenTime() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isBlockedByScreenTime"))
	return rv
}/* debug [instance_properties/getter]: isBlockedByScreenTime */


// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isInspectable
func (w_ WebView) Inspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("inspectable"))
	return rv
}/* debug [instance_properties/getter]: inspectable */


// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isInspectable
func (w_ WebView) SetInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInspectable:"), value)
}/* debug [instance_properties/setter]: inspectable */


// A Boolean value that indicates whether the view is currently loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isLoading
func (w_ WebView) Loading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isWritingToolsActive
func (w_ WebView) WritingToolsActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("writingToolsActive"))
	return rv
}/* debug [instance_properties/getter]: writingToolsActive */


// The factor by which the page content is currently scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/magnification
func (w_ WebView) Magnification() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("magnification"))
	return rv
}/* debug [instance_properties/getter]: magnification */


// The factor by which the page content is currently scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/magnification
func (w_ WebView) SetMagnification(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMagnification:"), value)
}/* debug [instance_properties/setter]: magnification */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/maximumViewportInset
func (w_ WebView) MaximumViewportInset() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](w_.ID, objc.Sel("maximumViewportInset"))
	return rv
}/* debug [instance_properties/getter]: maximumViewportInset */


// The media type for the contents of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/mediaType
func (w_ WebView) MediaType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("mediaType"))
	return rv
}/* debug [instance_properties/getter]: mediaType */


// The media type for the contents of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/mediaType
func (w_ WebView) SetMediaType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaType:"), value)
}/* debug [instance_properties/setter]: mediaType */


// An enumeration case that indicates whether the webpage is using the microphone to capture audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/microphoneCaptureState
func (w_ WebView) MicrophoneCaptureState() MediaCaptureState {
	rv := objc.Send[MediaCaptureState](w_.ID, objc.Sel("microphoneCaptureState"))
	return rv
}/* debug [instance_properties/getter]: microphoneCaptureState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/minimumViewportInset
func (w_ WebView) MinimumViewportInset() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](w_.ID, objc.Sel("minimumViewportInset"))
	return rv
}/* debug [instance_properties/getter]: minimumViewportInset */


// The object you use to manage navigation behavior for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/navigationDelegate
func (w_ WebView) NavigationDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("navigationDelegate"))
	return rv
}/* debug [instance_properties/getter]: navigationDelegate */


// The object you use to manage navigation behavior for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/navigationDelegate
func (w_ WebView) SetNavigationDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setNavigationDelegate:"), value)
}/* debug [instance_properties/setter]: navigationDelegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/obscuredContentInsets
func (w_ WebView) ObscuredContentInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](w_.ID, objc.Sel("obscuredContentInsets"))
	return rv
}/* debug [instance_properties/getter]: obscuredContentInsets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/obscuredContentInsets
func (w_ WebView) SetObscuredContentInsets(value foundation.EdgeInsets) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setObscuredContentInsets:"), value)
}/* debug [instance_properties/setter]: obscuredContentInsets */


// The scale factor by which the web view scales content relative to its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/pageZoom
func (w_ WebView) PageZoom() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("pageZoom"))
	return rv
}/* debug [instance_properties/getter]: pageZoom */


// The scale factor by which the web view scales content relative to its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/pageZoom
func (w_ WebView) SetPageZoom(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPageZoom:"), value)
}/* debug [instance_properties/setter]: pageZoom */


// The trust management object you use to evaluate trust for the current webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/serverTrust
func (w_ WebView) ServerTrust() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("serverTrust"))
	return rv
}/* debug [instance_properties/getter]: serverTrust */


// The theme color that the system gets from the first valid meta tag in the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/themeColor
func (w_ WebView) ThemeColor() appkit.Color {
	rv := objc.Send[appkit.Color](w_.ID, objc.Sel("themeColor"))
	return rv
}/* debug [instance_properties/getter]: themeColor */


// The page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/title
func (w_ WebView) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/uiDelegate
func (w_ WebView) UIDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("UIDelegate"))
	return rv
}/* debug [instance_properties/getter]: UIDelegate */


// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/uiDelegate
func (w_ WebView) SetUIDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUIDelegate:"), value)
}/* debug [instance_properties/setter]: UIDelegate */


// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/underPageBackgroundColor
func (w_ WebView) UnderPageBackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](w_.ID, objc.Sel("underPageBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: underPageBackgroundColor */


// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/underPageBackgroundColor
func (w_ WebView) SetUnderPageBackgroundColor(value appkit.Color) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUnderPageBackgroundColor:"), value)
}/* debug [instance_properties/setter]: underPageBackgroundColor */


// The URL for the current webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/url
func (w_ WebView) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isfindinteractionenabled
func (w_ WebView) IsFindInteractionEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isFindInteractionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isFindInteractionEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isfindinteractionenabled
func (w_ WebView) SetIsFindInteractionEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsFindInteractionEnabled:"), value)
}/* debug [instance_properties/setter]: isFindInteractionEnabled */


// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isinspectable
func (w_ WebView) IsInspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isInspectable"))
	return rv
}/* debug [instance_properties/getter]: isInspectable */


// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isinspectable
func (w_ WebView) SetIsInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsInspectable:"), value)
}/* debug [instance_properties/setter]: isInspectable */


// A Boolean value that indicates whether the view is currently loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isloading
func (w_ WebView) IsLoading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoading"))
	return rv
}/* debug [instance_properties/getter]: isLoading */


// A Boolean value that indicates whether the view is currently loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isloading
func (w_ WebView) SetIsLoading(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoading:"), value)
}/* debug [instance_properties/setter]: isLoading */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/iswritingtoolsactive
func (w_ WebView) IsWritingToolsActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isWritingToolsActive"))
	return rv
}/* debug [instance_properties/getter]: isWritingToolsActive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/iswritingtoolsactive
func (w_ WebView) SetIsWritingToolsActive(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsWritingToolsActive:"), value)
}/* debug [instance_properties/setter]: isWritingToolsActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebView */


