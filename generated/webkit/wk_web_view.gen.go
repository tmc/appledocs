// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [WebView] class.
type IWebView interface {
	appkit.IView
	CloseAllMediaPresentations()
	CloseAllMediaPresentationsWithCompletionHandler(completionHandler func())
	CreatePDFWithConfigurationCompletionHandler(pdfConfiguration unsafe.Pointer, completionHandler unsafe.Pointer)
	CreateWebArchiveDataWithCompletionHandler(completionHandler unsafe.Pointer)
	EvaluateJavaScriptCompletionHandler(javaScriptString string, completionHandler unsafe.Pointer)
	FetchDataOfTypesCompletionHandler(dataTypes unsafe.Pointer, completionHandler unsafe.Pointer)
	FindStringWithConfigurationCompletionHandler(string_ string, configuration unsafe.Pointer, completionHandler unsafe.Pointer)
	GoToBackForwardListItem(item unsafe.Pointer) unsafe.Pointer
	GoBack() unsafe.Pointer
	GoForward() unsafe.Pointer
	LoadRequest(request unsafe.Pointer) unsafe.Pointer
	LoadDataMIMETypeCharacterEncodingNameBaseURL(data unsafe.Pointer, MIMEType string, characterEncodingName string, baseURL foundation.URL) unsafe.Pointer
	LoadFileRequestAllowingReadAccessToURL(request unsafe.Pointer, readAccessURL foundation.URL) unsafe.Pointer
	LoadFileURLAllowingReadAccessToURL(URL foundation.URL, readAccessURL foundation.URL) unsafe.Pointer
	LoadHTMLStringBaseURL(string_ string, baseURL foundation.URL) unsafe.Pointer
	LoadSimulatedRequestResponseResponseData(request unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer
	LoadSimulatedRequestResponseHTMLString(request unsafe.Pointer, string_ string) unsafe.Pointer
	LoadSimulatedRequestWithResponseResponseData(request unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer
	LoadSimulatedRequestWithResponseHTMLString(request unsafe.Pointer, string_ string) unsafe.Pointer
	PauseAllMediaPlaybackWithCompletionHandler(completionHandler func())
	PauseAllMediaPlayback(completionHandler func())
	PrintOperationWithPrintInfo(printInfo unsafe.Pointer) unsafe.Pointer
	Reload() unsafe.Pointer
	ReloadFromOrigin() unsafe.Pointer
	RequestMediaPlaybackStateWithCompletionHandler(completionHandler unsafe.Pointer)
	RequestMediaPlaybackState(completionHandler unsafe.Pointer)
	RestoreDataCompletionHandler(data unsafe.Pointer, completionHandler func(error objc.ID))
	ResumeAllMediaPlayback(completionHandler func())
	ResumeDownloadFromResumeDataCompletionHandler(resumeData unsafe.Pointer, completionHandler unsafe.Pointer)
	SetAllMediaPlaybackSuspendedCompletionHandler(suspended bool, completionHandler func())
	SetCameraCaptureStateCompletionHandler(state unsafe.Pointer, completionHandler func())
	SetMagnificationCenteredAtPoint(magnification float64, point coregraphics.CGPoint)
	SetMicrophoneCaptureStateCompletionHandler(state unsafe.Pointer, completionHandler func())
	SetMinimumViewportInsetMaximumViewportInset(minimumViewportInset unsafe.Pointer, maximumViewportInset unsafe.Pointer)
	StartDownloadUsingRequestCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer)
	StopLoading()
	SuspendAllMediaPlayback(completionHandler func())
	TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration unsafe.Pointer, completionHandler unsafe.Pointer)
}

// An object that displays interactive web content, such as for an in-app browser.
//
// A object is a platform-native view that you use to incorporate web content seamlessly into your app’s UI. A web view supports a full web-browsing experience, and presents HTML, CSS, and JavaScript content alongside your app’s native views. Use it when web technologies satisfy your app’s layout and styling requirements more readily than native views. For example, you might use it when your app’s content changes frequently. A web view offers control over the navigation and user experience through delegate objects. Use the navigation delegate to react when the user clicks links in your web content, or interacts with the content in a way that affects navigation. For example, you might prevent the user from navigating to new content unless specific conditions are met. Use the UI delegate to present native UI elements, such as alerts or contextual menus, in response to interactions with your web content. Embed a object programmatically into your view hierarchy, or add it using Interface Builder. Interface Builder supports many customizations, such as configuring data detectors, media playback, and interaction behaviors. For more extensive customizations, create your web view programmatically using a object. For example, use a web view configuration object to specify handlers for custom URL schemes, manage cookies, and customize preferences for your web content. Before your web view appears onscreen, load content from a web server using a structure or load content directly from a local file or HTML string. The web view automatically loads embedded resources such as images or videos as part of the initial load request. It then renders your content and displays the results inside the view’s bounds rectangle. The following code example shows a view controller that replaces its default view with a custom object. A web view automatically converts telephone numbers that appear in web content to Phone links. When the user taps a Phone link, the Phone app launches and dials the number. Use the object to change the default data detector behavior. You can also use to programmatically set the scale of web content the first time it appears in a web view. Thereafter, the user can change the scale using gestures.
//
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




// Returns an object initialized from data in the specified coder object.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/init(coder:)
func NewWebViewWithCoder(coder unsafe.Pointer) WebView {
	instance := getWebViewClass().Alloc()
	rv := objc.Send[WebView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Creates a web view and initializes it with the specified frame and configuration data.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/init(frame:configuration:)
func NewWebViewWithFrameConfiguration(frame coregraphics.CGRect, configuration unsafe.Pointer) WebView {
	instance := getWebViewClass().Alloc()
	rv := objc.Send[WebView](instance.ID, objc.Sel("initWithFrame:configuration:"), frame, configuration)
	rv.Autorelease()
	return rv
}


// Returns a Boolean value that indicates whether WebKit natively supports resources with the specified URL scheme.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/handlesURLScheme(_:)
func (wc _WebViewClass) HandlesURLScheme(urlScheme string) bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("handlesURLScheme:"), objc.String(urlScheme))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/closeAllMediaPresentations()
func (w_ WebView) CloseAllMediaPresentations() {
	objc.Send[objc.ID](w_.ID, objc.Sel("closeAllMediaPresentations"))
}

// Closes all media the web view is presenting, including picture-in-picture video and fullscreen video.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/closeAllMediaPresentations(completionHandler:)
func (w_ WebView) CloseAllMediaPresentationsWithCompletionHandler(completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("closeAllMediaPresentationsWithCompletionHandler:"), completionHandler)
}

// Generates PDF data from the web view’s contents asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/createPDFWithConfiguration:completionHandler:
func (w_ WebView) CreatePDFWithConfigurationCompletionHandler(pdfConfiguration unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("createPDFWithConfiguration:completionHandler:"), pdfConfiguration, completionHandler)
}

// Creates a web archive of the web view’s contents asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/createWebArchiveDataWithCompletionHandler:
func (w_ WebView) CreateWebArchiveDataWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("createWebArchiveDataWithCompletionHandler:"), completionHandler)
}

// Evaluates the specified JavaScript string.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/evaluateJavaScript(_:completionHandler:)
func (w_ WebView) EvaluateJavaScriptCompletionHandler(javaScriptString string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("evaluateJavaScript:completionHandler:"), objc.String(javaScriptString), completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/fetchData(of:completionHandler:)
func (w_ WebView) FetchDataOfTypesCompletionHandler(dataTypes unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("fetchDataOfTypes:completionHandler:"), dataTypes, completionHandler)
}

// Searches for the specified string in the web view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/findString:withConfiguration:completionHandler:
func (w_ WebView) FindStringWithConfigurationCompletionHandler(string_ string, configuration unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("findString:withConfiguration:completionHandler:"), objc.String(string_), configuration, completionHandler)
}

// Navigates to an item from the back-forward list and sets it as the current item.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/go(to:)
func (w_ WebView) GoToBackForwardListItem(item unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("goToBackForwardListItem:"), item)
	return rv
}

// Navigates to the back item in the back-forward list.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/goBack()
func (w_ WebView) GoBack() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("goBack"))
	return rv
}

// Navigates to the forward item in the back-forward list.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/goForward()
func (w_ WebView) GoForward() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("goForward"))
	return rv
}

// Loads the web content that the specified URL request object references and navigates to that content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/load(_:)
func (w_ WebView) LoadRequest(request unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadRequest:"), request)
	return rv
}

// Loads the content of the specified data object and navigates to it.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/load(_:mimeType:characterEncodingName:baseURL:)
func (w_ WebView) LoadDataMIMETypeCharacterEncodingNameBaseURL(data unsafe.Pointer, MIMEType string, characterEncodingName string, baseURL foundation.URL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadData:MIMEType:characterEncodingName:baseURL:"), data, objc.String(MIMEType), objc.String(characterEncodingName), baseURL)
	return rv
}

// Loads the web content from the file the URL request object specifies and navigates to that content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadFileRequest(_:allowingReadAccessTo:)
func (w_ WebView) LoadFileRequestAllowingReadAccessToURL(request unsafe.Pointer, readAccessURL foundation.URL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadFileRequest:allowingReadAccessToURL:"), request, readAccessURL)
	return rv
}

// Loads the web content from the specified file and navigates to it.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadFileURL(_:allowingReadAccessTo:)
func (w_ WebView) LoadFileURLAllowingReadAccessToURL(URL foundation.URL, readAccessURL foundation.URL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadFileURL:allowingReadAccessToURL:"), URL, readAccessURL)
	return rv
}

// Loads the contents of the specified HTML string and navigates to it.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadHTMLString(_:baseURL:)
func (w_ WebView) LoadHTMLStringBaseURL(string_ string, baseURL foundation.URL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadHTMLString:baseURL:"), objc.String(string_), baseURL)
	return rv
}

// Loads the web content from the data you provide as if the data were the response to the request.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:response:responseData:)
func (w_ WebView) LoadSimulatedRequestResponseResponseData(request unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadSimulatedRequest:response:responseData:"), request, response, data)
	return rv
}

// Loads the web content from the HTML you provide as if the HTML were the response to the request.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:responseHTML:)
func (w_ WebView) LoadSimulatedRequestResponseHTMLString(request unsafe.Pointer, string_ string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadSimulatedRequest:responseHTMLString:"), request, objc.String(string_))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:with:responseData:)
func (w_ WebView) LoadSimulatedRequestWithResponseResponseData(request unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadSimulatedRequest:withResponse:responseData:"), request, response, data)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:withResponseHTML:)
func (w_ WebView) LoadSimulatedRequestWithResponseHTMLString(request unsafe.Pointer, string_ string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("loadSimulatedRequest:withResponseHTMLString:"), request, objc.String(string_))
	return rv
}

// Pauses playback of all media in the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/pauseAllMediaPlayback(completionHandler:)
func (w_ WebView) PauseAllMediaPlaybackWithCompletionHandler(completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("pauseAllMediaPlaybackWithCompletionHandler:"), completionHandler)
}

// Pauses playback of all media in the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/pauseAllMediaPlayback:
func (w_ WebView) PauseAllMediaPlayback(completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("pauseAllMediaPlayback:"), completionHandler)
}

// Returns the print operation object to use when printing the contents of the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/printOperation(with:)
func (w_ WebView) PrintOperationWithPrintInfo(printInfo unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("printOperationWithPrintInfo:"), printInfo)
	return rv
}

// Reloads the current webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/reload()
func (w_ WebView) Reload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("reload"))
	return rv
}

// Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/reloadFromOrigin()
func (w_ WebView) ReloadFromOrigin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("reloadFromOrigin"))
	return rv
}

// Requests the playback status of media in the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/requestMediaPlaybackState(completionHandler:)
func (w_ WebView) RequestMediaPlaybackStateWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("requestMediaPlaybackStateWithCompletionHandler:"), completionHandler)
}

// Requests the playback status of media in the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/requestMediaPlaybackState:
func (w_ WebView) RequestMediaPlaybackState(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("requestMediaPlaybackState:"), completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/restoreData(_:completionHandler:)
func (w_ WebView) RestoreDataCompletionHandler(data unsafe.Pointer, completionHandler func(error objc.ID)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("restoreData:completionHandler:"), data, completionHandler)
}

// Resumes playback of all media in a web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/resumeAllMediaPlayback:
func (w_ WebView) ResumeAllMediaPlayback(completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("resumeAllMediaPlayback:"), completionHandler)
}

// Resumes a failed or canceled download.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/resumeDownload(fromResumeData:completionHandler:)
func (w_ WebView) ResumeDownloadFromResumeDataCompletionHandler(resumeData unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("resumeDownloadFromResumeData:completionHandler:"), resumeData, completionHandler)
}

// Changes whether the webpage is suspending playback of all media in the page.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setAllMediaPlaybackSuspended(_:completionHandler:)
func (w_ WebView) SetAllMediaPlaybackSuspendedCompletionHandler(suspended bool, completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllMediaPlaybackSuspended:completionHandler:"), suspended, completionHandler)
}

// Changes whether the webpage is using the camera to capture images or video.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setCameraCaptureState(_:completionHandler:)
func (w_ WebView) SetCameraCaptureStateCompletionHandler(state unsafe.Pointer, completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCameraCaptureState:completionHandler:"), state, completionHandler)
}

// Scales the page content and centers the result on the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setMagnification(_:centeredAt:)
func (w_ WebView) SetMagnificationCenteredAtPoint(magnification float64, point coregraphics.CGPoint) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}

// Changes whether the webpage is using the microphone to capture audio.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setMicrophoneCaptureState(_:completionHandler:)
func (w_ WebView) SetMicrophoneCaptureStateCompletionHandler(state unsafe.Pointer, completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMicrophoneCaptureState:completionHandler:"), state, completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setMinimumViewportInset(_:maximumViewportInset:)
func (w_ WebView) SetMinimumViewportInsetMaximumViewportInset(minimumViewportInset unsafe.Pointer, maximumViewportInset unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinimumViewportInset:maximumViewportInset:"), minimumViewportInset, maximumViewportInset)
}

// Starts to download the resource at the URL in the request.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/startDownload(using:completionHandler:)
func (w_ WebView) StartDownloadUsingRequestCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("startDownloadUsingRequest:completionHandler:"), request, completionHandler)
}

// Stops loading all resources on the current page.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/stopLoading()
func (w_ WebView) StopLoading() {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopLoading"))
}

// Changes whether the webpage is suspending playback of all media in the page.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/suspendAllMediaPlayback:
func (w_ WebView) SuspendAllMediaPlayback(completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("suspendAllMediaPlayback:"), completionHandler)
}

// Generates a platform-native image from the web view’s contents asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/takeSnapshot(with:completionHandler:)
func (w_ WebView) TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("takeSnapshotWithConfiguration:completionHandler:"), snapshotConfiguration, completionHandler)
}

// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsBackForwardNavigationGestures
func (w_ WebView) AllowsBackForwardNavigationGestures() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsBackForwardNavigationGestures"))
	return rv
}


// SetAllowsBackForwardNavigationGestures sets the value of the allowsBackForwardNavigationGestures property.
// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsBackForwardNavigationGestures
func (w_ WebView) SetAllowsBackForwardNavigationGestures(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsBackForwardNavigationGestures:"), value)
}

// A Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsLinkPreview
func (w_ WebView) AllowsLinkPreview() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsLinkPreview"))
	return rv
}


// SetAllowsLinkPreview sets the value of the allowsLinkPreview property.
// A Boolean value that determines whether pressing a link displays a preview of the destination for the link.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsLinkPreview
func (w_ WebView) SetAllowsLinkPreview(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsLinkPreview:"), value)
}

// A Boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsMagnification
func (w_ WebView) AllowsMagnification() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsMagnification"))
	return rv
}


// SetAllowsMagnification sets the value of the allowsMagnification property.
// A Boolean value that indicates whether magnify gestures change the web view’s magnification.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/allowsMagnification
func (w_ WebView) SetAllowsMagnification(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsMagnification:"), value)
}

// The web view’s back-forward list.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/backForwardList
func (w_ WebView) BackForwardList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("backForwardList"))
	return rv
}

// An enumeration case that indicates whether the webpage is using the camera to capture images or video.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/cameraCaptureState
func (w_ WebView) CameraCaptureState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("cameraCaptureState"))
	return rv
}

// A Boolean value that indicates whether there is a valid back item in the back-forward list.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/canGoBack
func (w_ WebView) CanGoBack() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoBack"))
	return rv
}

// A Boolean value that indicates whether there is a valid forward item in the back-forward list.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/canGoForward
func (w_ WebView) CanGoForward() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoForward"))
	return rv
}

// An array of objects forming the certificate chain for the currently committed navigation.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/certificateChain
func (w_ WebView) CertificateChain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("certificateChain"))
	return rv
}

// The object that contains the configuration details for the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/configuration
func (w_ WebView) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("configuration"))
	return rv
}

// The custom user agent string.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/customUserAgent
func (w_ WebView) CustomUserAgent() string {
	rv := objc.Send[string](w_.ID, objc.Sel("customUserAgent"))
	return rv
}


// SetCustomUserAgent sets the value of the customUserAgent property.
// The custom user agent string.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/customUserAgent
func (w_ WebView) SetCustomUserAgent(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomUserAgent:"), objc.String(value))
}

// An estimate of what fraction of the current navigation has been loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/estimatedProgress
func (w_ WebView) EstimatedProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("estimatedProgress"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/findInteraction
func (w_ WebView) FindInteraction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("findInteraction"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/fullscreenState-swift.property
func (w_ WebView) FullscreenState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("fullscreenState"))
	return rv
}

// A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/hasOnlySecureContent
func (w_ WebView) HasOnlySecureContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOnlySecureContent"))
	return rv
}

// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/interactionState
func (w_ WebView) InteractionState() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("interactionState"))
	return rv
}


// SetInteractionState sets the value of the interactionState property.
// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/interactionState
func (w_ WebView) SetInteractionState(value objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInteractionState:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isBlockedByScreenTime
func (w_ WebView) IsBlockedByScreenTime() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isBlockedByScreenTime"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isFindInteractionEnabled
func (w_ WebView) FindInteractionEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("findInteractionEnabled"))
	return rv
}


// SetFindInteractionEnabled sets the value of the findInteractionEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isFindInteractionEnabled
func (w_ WebView) SetFindInteractionEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFindInteractionEnabled:"), value)
}

// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isInspectable
func (w_ WebView) Inspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("inspectable"))
	return rv
}


// SetInspectable sets the value of the inspectable property.
// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isInspectable
func (w_ WebView) SetInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInspectable:"), value)
}

// A Boolean value that indicates whether the view is currently loading content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isLoading
func (w_ WebView) Loading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("loading"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isWritingToolsActive
func (w_ WebView) WritingToolsActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("writingToolsActive"))
	return rv
}

// The factor by which the page content is currently scaled.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/magnification
func (w_ WebView) Magnification() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("magnification"))
	return rv
}


// SetMagnification sets the value of the magnification property.
// The factor by which the page content is currently scaled.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/magnification
func (w_ WebView) SetMagnification(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMagnification:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/maximumViewportInset
func (w_ WebView) MaximumViewportInset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("maximumViewportInset"))
	return rv
}

// The media type for the contents of the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/mediaType
func (w_ WebView) MediaType() string {
	rv := objc.Send[string](w_.ID, objc.Sel("mediaType"))
	return rv
}


// SetMediaType sets the value of the mediaType property.
// The media type for the contents of the web view.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/mediaType
func (w_ WebView) SetMediaType(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaType:"), objc.String(value))
}

// An enumeration case that indicates whether the webpage is using the microphone to capture audio.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/microphoneCaptureState
func (w_ WebView) MicrophoneCaptureState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("microphoneCaptureState"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/minimumViewportInset
func (w_ WebView) MinimumViewportInset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("minimumViewportInset"))
	return rv
}

// The object you use to manage navigation behavior for the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/navigationDelegate
func (w_ WebView) NavigationDelegate() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("navigationDelegate"))
	return rv
}


// SetNavigationDelegate sets the value of the navigationDelegate property.
// The object you use to manage navigation behavior for the web view.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/navigationDelegate
func (w_ WebView) SetNavigationDelegate(value objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setNavigationDelegate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/obscuredContentInsets
func (w_ WebView) ObscuredContentInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("obscuredContentInsets"))
	return rv
}


// SetObscuredContentInsets sets the value of the obscuredContentInsets property.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/obscuredContentInsets
func (w_ WebView) SetObscuredContentInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setObscuredContentInsets:"), value)
}

// The scale factor by which the web view scales content relative to its bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/pageZoom
func (w_ WebView) PageZoom() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("pageZoom"))
	return rv
}


// SetPageZoom sets the value of the pageZoom property.
// The scale factor by which the web view scales content relative to its bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/pageZoom
func (w_ WebView) SetPageZoom(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPageZoom:"), value)
}

// The scroll view associated with the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/scrollView
func (w_ WebView) ScrollView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("scrollView"))
	return rv
}

// The trust management object you use to evaluate trust for the current webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/serverTrust
func (w_ WebView) ServerTrust() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("serverTrust"))
	return rv
}

// The theme color that the system gets from the first valid meta tag in the webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/themeColor
func (w_ WebView) ThemeColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("themeColor"))
	return rv
}

// The page title.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/title
func (w_ WebView) Title() string {
	rv := objc.Send[string](w_.ID, objc.Sel("title"))
	return rv
}

// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/uiDelegate
func (w_ WebView) UIDelegate() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("UIDelegate"))
	return rv
}


// SetUIDelegate sets the value of the UIDelegate property.
// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/uiDelegate
func (w_ WebView) SetUIDelegate(value objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUIDelegate:"), value)
}

// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/underPageBackgroundColor
func (w_ WebView) UnderPageBackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("underPageBackgroundColor"))
	return rv
}


// SetUnderPageBackgroundColor sets the value of the underPageBackgroundColor property.
// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/underPageBackgroundColor
func (w_ WebView) SetUnderPageBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUnderPageBackgroundColor:"), value)
}

// The URL for the current webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/url
func (w_ WebView) URL() foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URL"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isfindinteractionenabled
func (w_ WebView) IsFindInteractionEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isFindInteractionEnabled"))
	return rv
}


// SetIsFindInteractionEnabled sets the value of the isFindInteractionEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isfindinteractionenabled
func (w_ WebView) SetIsFindInteractionEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsFindInteractionEnabled:"), value)
}

// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isinspectable
func (w_ WebView) IsInspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isInspectable"))
	return rv
}


// SetIsInspectable sets the value of the isInspectable property.
// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isinspectable
func (w_ WebView) SetIsInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsInspectable:"), value)
}

// A Boolean value that indicates whether the view is currently loading content.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isloading
func (w_ WebView) IsLoading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoading"))
	return rv
}


// SetIsLoading sets the value of the isLoading property.
// A Boolean value that indicates whether the view is currently loading content.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isloading
func (w_ WebView) SetIsLoading(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoading:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/iswritingtoolsactive
func (w_ WebView) IsWritingToolsActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isWritingToolsActive"))
	return rv
}


// SetIsWritingToolsActive sets the value of the isWritingToolsActive property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/iswritingtoolsactive
func (w_ WebView) SetIsWritingToolsActive(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsWritingToolsActive:"), value)
}


