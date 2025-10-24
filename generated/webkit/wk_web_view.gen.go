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
	// properties:
	AllowsBackForwardNavigationGestures() bool
	SetAllowsBackForwardNavigationGestures(value bool)
	AllowsLinkPreview() bool
	SetAllowsLinkPreview(value bool)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	BackForwardList() IWKBackForwardList
	SetBackForwardList(value IWKBackForwardList)
	CameraCaptureState() MediaCaptureState
	SetCameraCaptureState(value MediaCaptureState)
	CanGoBack() bool
	SetCanGoBack(value bool)
	CanGoForward() bool
	SetCanGoForward(value bool)
	Configuration() IWKWebViewConfiguration
	SetConfiguration(value IWKWebViewConfiguration)
	CustomUserAgent() objc.IObject /* cross-framework: NSString */
	SetCustomUserAgent(value objc.IObject /* cross-framework: NSString */)
	EstimatedProgress() float64
	SetEstimatedProgress(value float64)
	FindInteraction() FindInteraction /* not a class type */
	SetFindInteraction(value FindInteraction /* not a class type */)
	FullscreenState() unsafe.Pointer
	SetFullscreenState(value unsafe.Pointer)
	HasOnlySecureContent() bool
	SetHasOnlySecureContent(value bool)
	InteractionState() unsafe.Pointer
	SetInteractionState(value unsafe.Pointer)
	IsBlockedByScreenTime() bool
	SetIsBlockedByScreenTime(value bool)
	IsFindInteractionEnabled() bool
	SetIsFindInteractionEnabled(value bool)
	IsInspectable() bool
	SetIsInspectable(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
	IsWritingToolsActive() bool
	SetIsWritingToolsActive(value bool)
	Magnification() float64
	SetMagnification(value float64)
	MaximumViewportInset() objc.IObject /* cross-framework: EdgeInsets */
	SetMaximumViewportInset(value objc.IObject /* cross-framework: EdgeInsets */)
	MediaType() objc.IObject /* cross-framework: NSString */
	SetMediaType(value objc.IObject /* cross-framework: NSString */)
	MicrophoneCaptureState() MediaCaptureState
	SetMicrophoneCaptureState(value MediaCaptureState)
	MinimumViewportInset() objc.IObject /* cross-framework: EdgeInsets */
	SetMinimumViewportInset(value objc.IObject /* cross-framework: EdgeInsets */)
	NavigationDelegate() NavigationDelegate /* not a class type */
	SetNavigationDelegate(value NavigationDelegate /* not a class type */)
	ObscuredContentInsets() objc.IObject /* cross-framework: EdgeInsets */
	SetObscuredContentInsets(value objc.IObject /* cross-framework: EdgeInsets */)
	PageZoom() float64
	SetPageZoom(value float64)
	ScrollView() objc.IObject /* cross-framework: ScrollView */
	SetScrollView(value objc.IObject /* cross-framework: ScrollView */)
	ServerTrust() unsafe.Pointer
	SetServerTrust(value unsafe.Pointer)
	ThemeColor() objc.IObject /* cross-framework: Color */
	SetThemeColor(value objc.IObject /* cross-framework: Color */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	UiDelegate() UIDelegate /* not a class type */
	SetUiDelegate(value UIDelegate /* not a class type */)
	UnderPageBackgroundColor() objc.IObject /* cross-framework: Color */
	SetUnderPageBackgroundColor(value objc.IObject /* cross-framework: Color */)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	// methods:
	GoBack(sender objectivec.IObject)
	LoadRequest(request objc.IObject /* cross-framework: URLRequest */) INavigation
	LoadFileURLAllowingReadAccessToURL(URL objc.IObject /* cross-framework: NSURL */, readAccessURL objc.IObject /* cross-framework: NSURL */) INavigation
	LoadSimulatedRequestResponseHTMLString(request objc.IObject /* cross-framework: URLRequest */, string_ objc.IObject /* cross-framework: NSString */) INavigation
	SetMagnificationCenteredAtPoint(magnification float64, point objc.IObject /* cross-framework: Point */)
}

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



// Navigates to the back item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/goBack(_:)
func (w_ WebView) GoBack(sender objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("goBack:"), sender)
}


// Loads the web content that the specified URL request object references and navigates to that content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/load(_:)
func (w_ WebView) LoadRequest(request objc.IObject /* cross-framework: URLRequest */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadRequest:"), request)
	return rv
}


// Loads the web content from the specified file and navigates to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadFileURL(_:allowingReadAccessTo:)
func (w_ WebView) LoadFileURLAllowingReadAccessToURL(URL objc.IObject /* cross-framework: NSURL */, readAccessURL objc.IObject /* cross-framework: NSURL */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadFileURL:allowingReadAccessToURL:"), URL, readAccessURL)
	return rv
}


// Loads the web content from the HTML you provide as if the HTML were the response to the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:responseHTML:)
func (w_ WebView) LoadSimulatedRequestResponseHTMLString(request objc.IObject /* cross-framework: URLRequest */, string_ objc.IObject /* cross-framework: NSString */) INavigation {
	rv := objc.Send[Navigation](w_.ID, objc.Sel("loadSimulatedRequest:responseHTMLString:"), request, string_)
	return rv
}


// Scales the page content and centers the result on the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/setMagnification(_:centeredAt:)
func (w_ WebView) SetMagnificationCenteredAtPoint(magnification float64, point objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}


// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/allowsbackforwardnavigationgestures
func (w_ WebView) AllowsBackForwardNavigationGestures() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsBackForwardNavigationGestures"))
	return rv
}


// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/allowsbackforwardnavigationgestures
func (w_ WebView) SetAllowsBackForwardNavigationGestures(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsBackForwardNavigationGestures:"), value)
}


// A Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/allowslinkpreview
func (w_ WebView) AllowsLinkPreview() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsLinkPreview"))
	return rv
}


// A Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/allowslinkpreview
func (w_ WebView) SetAllowsLinkPreview(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsLinkPreview:"), value)
}


// A Boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/allowsmagnification
func (w_ WebView) AllowsMagnification() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsMagnification"))
	return rv
}


// A Boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/allowsmagnification
func (w_ WebView) SetAllowsMagnification(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsMagnification:"), value)
}


// The web view’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/backforwardlist
func (w_ WebView) BackForwardList() IWKBackForwardList {
	rv := objc.Send[BackForwardList](w_.ID, objc.Sel("backForwardList"))
	return rv
}


// The web view’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/backforwardlist
func (w_ WebView) SetBackForwardList(value IWKBackForwardList) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBackForwardList:"), value)
}


// An enumeration case that indicates whether the webpage is using the camera to capture images or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/cameracapturestate
func (w_ WebView) CameraCaptureState() MediaCaptureState {
	rv := objc.Send[MediaCaptureState](w_.ID, objc.Sel("cameraCaptureState"))
	return rv
}


// An enumeration case that indicates whether the webpage is using the camera to capture images or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/cameracapturestate
func (w_ WebView) SetCameraCaptureState(value MediaCaptureState) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCameraCaptureState:"), value)
}


// A Boolean value that indicates whether there is a valid back item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/cangoback
func (w_ WebView) CanGoBack() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoBack"))
	return rv
}


// A Boolean value that indicates whether there is a valid back item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/cangoback
func (w_ WebView) SetCanGoBack(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanGoBack:"), value)
}


// A Boolean value that indicates whether there is a valid forward item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/cangoforward
func (w_ WebView) CanGoForward() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canGoForward"))
	return rv
}


// A Boolean value that indicates whether there is a valid forward item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/cangoforward
func (w_ WebView) SetCanGoForward(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanGoForward:"), value)
}


// The object that contains the configuration details for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/configuration
func (w_ WebView) Configuration() IWKWebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](w_.ID, objc.Sel("configuration"))
	return rv
}


// The object that contains the configuration details for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/configuration
func (w_ WebView) SetConfiguration(value IWKWebViewConfiguration) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setConfiguration:"), value)
}


// The custom user agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/customuseragent
func (w_ WebView) CustomUserAgent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("customUserAgent"))
	return rv
}


// The custom user agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/customuseragent
func (w_ WebView) SetCustomUserAgent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCustomUserAgent:"), value)
}


// An estimate of what fraction of the current navigation has been loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/estimatedprogress
func (w_ WebView) EstimatedProgress() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("estimatedProgress"))
	return rv
}


// An estimate of what fraction of the current navigation has been loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/estimatedprogress
func (w_ WebView) SetEstimatedProgress(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEstimatedProgress:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/findinteraction
func (w_ WebView) FindInteraction() FindInteraction /* not a class type */ {
	rv := objc.Send[FindInteraction](w_.ID, objc.Sel("findInteraction"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/findinteraction
func (w_ WebView) SetFindInteraction(value FindInteraction /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFindInteraction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/fullscreenstate-swift.property
func (w_ WebView) FullscreenState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("fullscreenState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/fullscreenstate-swift.property
func (w_ WebView) SetFullscreenState(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFullscreenState:"), value)
}


// A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/hasonlysecurecontent
func (w_ WebView) HasOnlySecureContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOnlySecureContent"))
	return rv
}


// A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/hasonlysecurecontent
func (w_ WebView) SetHasOnlySecureContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasOnlySecureContent:"), value)
}


// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/interactionstate
func (w_ WebView) InteractionState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("interactionState"))
	return rv
}


// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/interactionstate
func (w_ WebView) SetInteractionState(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInteractionState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isblockedbyscreentime
func (w_ WebView) IsBlockedByScreenTime() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isBlockedByScreenTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isblockedbyscreentime
func (w_ WebView) SetIsBlockedByScreenTime(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsBlockedByScreenTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isfindinteractionenabled
func (w_ WebView) IsFindInteractionEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isFindInteractionEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isfindinteractionenabled
func (w_ WebView) SetIsFindInteractionEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsFindInteractionEnabled:"), value)
}


// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isinspectable
func (w_ WebView) IsInspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isInspectable"))
	return rv
}


// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isinspectable
func (w_ WebView) SetIsInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsInspectable:"), value)
}


// A Boolean value that indicates whether the view is currently loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isloading
func (w_ WebView) IsLoading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoading"))
	return rv
}


// A Boolean value that indicates whether the view is currently loading content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/isloading
func (w_ WebView) SetIsLoading(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoading:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/iswritingtoolsactive
func (w_ WebView) IsWritingToolsActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isWritingToolsActive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/iswritingtoolsactive
func (w_ WebView) SetIsWritingToolsActive(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsWritingToolsActive:"), value)
}


// The factor by which the page content is currently scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/magnification
func (w_ WebView) Magnification() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("magnification"))
	return rv
}


// The factor by which the page content is currently scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/magnification
func (w_ WebView) SetMagnification(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMagnification:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/maximumviewportinset
func (w_ WebView) MaximumViewportInset() objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[foundation.EdgeInsets](w_.ID, objc.Sel("maximumViewportInset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/maximumviewportinset
func (w_ WebView) SetMaximumViewportInset(value objc.IObject /* cross-framework: EdgeInsets */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMaximumViewportInset:"), value)
}


// The media type for the contents of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/mediatype
func (w_ WebView) MediaType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("mediaType"))
	return rv
}


// The media type for the contents of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/mediatype
func (w_ WebView) SetMediaType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaType:"), value)
}


// An enumeration case that indicates whether the webpage is using the microphone to capture audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/microphonecapturestate
func (w_ WebView) MicrophoneCaptureState() MediaCaptureState {
	rv := objc.Send[MediaCaptureState](w_.ID, objc.Sel("microphoneCaptureState"))
	return rv
}


// An enumeration case that indicates whether the webpage is using the microphone to capture audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/microphonecapturestate
func (w_ WebView) SetMicrophoneCaptureState(value MediaCaptureState) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMicrophoneCaptureState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/minimumviewportinset
func (w_ WebView) MinimumViewportInset() objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[foundation.EdgeInsets](w_.ID, objc.Sel("minimumViewportInset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/minimumviewportinset
func (w_ WebView) SetMinimumViewportInset(value objc.IObject /* cross-framework: EdgeInsets */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinimumViewportInset:"), value)
}


// The object you use to manage navigation behavior for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/navigationdelegate
func (w_ WebView) NavigationDelegate() NavigationDelegate /* not a class type */ {
	rv := objc.Send[NavigationDelegate](w_.ID, objc.Sel("navigationDelegate"))
	return rv
}


// The object you use to manage navigation behavior for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/navigationdelegate
func (w_ WebView) SetNavigationDelegate(value NavigationDelegate /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setNavigationDelegate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/obscuredcontentinsets
func (w_ WebView) ObscuredContentInsets() objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[foundation.EdgeInsets](w_.ID, objc.Sel("obscuredContentInsets"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/obscuredcontentinsets
func (w_ WebView) SetObscuredContentInsets(value objc.IObject /* cross-framework: EdgeInsets */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setObscuredContentInsets:"), value)
}


// The scale factor by which the web view scales content relative to its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/pagezoom
func (w_ WebView) PageZoom() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("pageZoom"))
	return rv
}


// The scale factor by which the web view scales content relative to its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/pagezoom
func (w_ WebView) SetPageZoom(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPageZoom:"), value)
}


// The scroll view associated with the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/scrollview
func (w_ WebView) ScrollView() objc.IObject /* cross-framework: ScrollView */ {
	rv := objc.Send[appkit.ScrollView](w_.ID, objc.Sel("scrollView"))
	return rv
}


// The scroll view associated with the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/scrollview
func (w_ WebView) SetScrollView(value objc.IObject /* cross-framework: ScrollView */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setScrollView:"), value)
}


// The trust management object you use to evaluate trust for the current webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/servertrust
func (w_ WebView) ServerTrust() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("serverTrust"))
	return rv
}


// The trust management object you use to evaluate trust for the current webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/servertrust
func (w_ WebView) SetServerTrust(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setServerTrust:"), value)
}


// The theme color that the system gets from the first valid meta tag in the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/themecolor
func (w_ WebView) ThemeColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](w_.ID, objc.Sel("themeColor"))
	return rv
}


// The theme color that the system gets from the first valid meta tag in the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/themecolor
func (w_ WebView) SetThemeColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setThemeColor:"), value)
}


// The page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/title
func (w_ WebView) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
}


// The page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/title
func (w_ WebView) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitle:"), value)
}


// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/uidelegate
func (w_ WebView) UiDelegate() UIDelegate /* not a class type */ {
	rv := objc.Send[Delegate](w_.ID, objc.Sel("uiDelegate"))
	return rv
}


// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/uidelegate
func (w_ WebView) SetUiDelegate(value UIDelegate /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUiDelegate:"), value)
}


// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/underpagebackgroundcolor
func (w_ WebView) UnderPageBackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](w_.ID, objc.Sel("underPageBackgroundColor"))
	return rv
}


// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/underpagebackgroundcolor
func (w_ WebView) SetUnderPageBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUnderPageBackgroundColor:"), value)
}


// The URL for the current webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/url
func (w_ WebView) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("url"))
	return rv
}


// The URL for the current webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/url
func (w_ WebView) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUrl:"), value)
}



