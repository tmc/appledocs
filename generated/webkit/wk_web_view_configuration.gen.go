// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WebViewConfiguration] class.
var (
	WebViewConfigurationClass     _WebViewConfigurationClass
	WebViewConfigurationClassOnce sync.Once
)

func getWebViewConfigurationClass() _WebViewConfigurationClass {
	WebViewConfigurationClassOnce.Do(func() {
		WebViewConfigurationClass = _WebViewConfigurationClass{objc.GetClass("WKWebViewConfiguration")}
	})
	return WebViewConfigurationClass
}

type _WebViewConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [WebViewConfiguration] class.
type IWebViewConfiguration interface {
	objectivec.IObject
	// properties:
	LimitsNavigationsToAppBoundDomains() bool
	SetLimitsNavigationsToAppBoundDomains(value bool)
	Preferences() IWKPreferences
	SetPreferences(value IWKPreferences)
	ProcessPool() IWKProcessPool
	SetProcessPool(value IWKProcessPool)
	WebsiteDataStore() IWKWebsiteDataStore
	SetWebsiteDataStore(value IWKWebsiteDataStore)
	AllowsAirPlayForMediaPlayback() bool
	SetAllowsAirPlayForMediaPlayback(value bool)
	AllowsInlineMediaPlayback() bool
	SetAllowsInlineMediaPlayback(value bool)
	AllowsInlinePredictions() bool
	SetAllowsInlinePredictions(value bool)
	AllowsPictureInPictureMediaPlayback() bool
	SetAllowsPictureInPictureMediaPlayback(value bool)
	ApplicationNameForUserAgent() objc.IObject /* cross-framework: NSString */
	SetApplicationNameForUserAgent(value objc.IObject /* cross-framework: NSString */)
	DataDetectorTypes() DataDetectorTypes /* not a class type */
	SetDataDetectorTypes(value DataDetectorTypes /* not a class type */)
	DefaultWebpagePreferences() IWKWebpagePreferences
	SetDefaultWebpagePreferences(value IWKWebpagePreferences)
	MediaPlaybackAllowsAirPlay() bool
	SetMediaPlaybackAllowsAirPlay(value bool)
	MediaPlaybackRequiresUserAction() bool
	SetMediaPlaybackRequiresUserAction(value bool)
	MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes /* not a class type */
	SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes /* not a class type */)
	RequiresUserActionForMediaPlayback() bool
	SetRequiresUserActionForMediaPlayback(value bool)
	SelectionGranularity() SelectionGranularity /* not a class type */
	SetSelectionGranularity(value SelectionGranularity /* not a class type */)
	ShowsSystemScreenTimeBlockingView() bool
	SetShowsSystemScreenTimeBlockingView(value bool)
	SupportsAdaptiveImageGlyph() bool
	SetSupportsAdaptiveImageGlyph(value bool)
	SuppressesIncrementalRendering() bool
	SetSuppressesIncrementalRendering(value bool)
	UpgradeKnownHostsToHTTPS() bool
	SetUpgradeKnownHostsToHTTPS(value bool)
	UserContentController() IWKUserContentController
	SetUserContentController(value IWKUserContentController)
	UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy /* not a class type */
	SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy /* not a class type */)
	WebExtensionController() IWKWebExtensionController
	SetWebExtensionController(value IWKWebExtensionController)
	WritingToolsBehavior() WritingToolsBehavior /* not a class type */
	SetWritingToolsBehavior(value WritingToolsBehavior /* not a class type */)
	// methods:
	SetURLSchemeHandlerForURLScheme(urlSchemeHandler objectivec.IObject, urlScheme objc.IObject /* cross-framework: NSString */)
}

// A collection of properties that you use to initialize a web view.
//
// A object provides information about how to configure a object. Use your configuration object to specify: The initial cookies to make available to your web content Handlers for any custom URL schemes your web content uses Settings for how to handle media content Information about how to manage selections within the web view Custom scripts to inject into the webpage Custom rules that determine how to render content You create a object in your code, configure its properties, and pass it to the initializer of your object. The web view incorporates your configuration settings only at creation time; you cannot change those settings dynamically later.


// A collection of properties that you use to initialize a web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration
type WebViewConfiguration struct {
	objectivec.Object
}

// WebViewConfigurationFrom constructs a [WebViewConfiguration] from an unsafe.Pointer.
//
// A collection of properties that you use to initialize a web view.
func WebViewConfigurationFrom(ptr unsafe.Pointer) WebViewConfiguration {
	return WebViewConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebViewConfigurationClass) Alloc() WebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebViewConfigurationClass) New() WebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebViewConfiguration) Init() WebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebViewConfiguration) Autorelease() WebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebViewConfiguration creates a new WebViewConfiguration instance.
func NewWebViewConfiguration() WebViewConfiguration {
	return getWebViewConfigurationClass().New()
}



// Registers an object to load resources associated with the specified URL scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/setURLSchemeHandler(_:forURLScheme:)
func (w_ WebViewConfiguration) SetURLSchemeHandlerForURLScheme(urlSchemeHandler objectivec.IObject, urlScheme objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setURLSchemeHandler:forURLScheme:"), urlSchemeHandler, urlScheme)
}


// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/limitsNavigationsToAppBoundDomains
func (w_ WebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("limitsNavigationsToAppBoundDomains"))
	return rv
}


// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/limitsNavigationsToAppBoundDomains
func (w_ WebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setLimitsNavigationsToAppBoundDomains:"), value)
}


// The object that manages the preference-related settings for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/preferences
func (w_ WebViewConfiguration) Preferences() IWKPreferences {
	rv := objc.Send[Preferences](w_.ID, objc.Sel("preferences"))
	return rv
}


// The object that manages the preference-related settings for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/preferences
func (w_ WebViewConfiguration) SetPreferences(value IWKPreferences) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferences:"), value)
}


// The object that coordinates the processes the web view uses to render its web content and execute scripts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/processPool
func (w_ WebViewConfiguration) ProcessPool() IWKProcessPool {
	rv := objc.Send[ProcessPool](w_.ID, objc.Sel("processPool"))
	return rv
}


// The object that coordinates the processes the web view uses to render its web content and execute scripts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/processPool
func (w_ WebViewConfiguration) SetProcessPool(value IWKProcessPool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setProcessPool:"), value)
}


// The object you use to get and set the site’s cookies and to track the cached data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/websiteDataStore
func (w_ WebViewConfiguration) WebsiteDataStore() IWKWebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](w_.ID, objc.Sel("websiteDataStore"))
	return rv
}


// The object you use to get and set the site’s cookies and to track the cached data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/websiteDataStore
func (w_ WebViewConfiguration) SetWebsiteDataStore(value IWKWebsiteDataStore) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebsiteDataStore:"), value)
}


// A Boolean value that indicates whether the web view allows media playback over AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsairplayformediaplayback
func (w_ WebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsAirPlayForMediaPlayback"))
	return rv
}


// A Boolean value that indicates whether the web view allows media playback over AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsairplayformediaplayback
func (w_ WebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsAirPlayForMediaPlayback:"), value)
}


// A Boolean value that indicates whether HTML5 videos play inline or use the native full-screen controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsinlinemediaplayback
func (w_ WebViewConfiguration) AllowsInlineMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsInlineMediaPlayback"))
	return rv
}


// A Boolean value that indicates whether HTML5 videos play inline or use the native full-screen controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsinlinemediaplayback
func (w_ WebViewConfiguration) SetAllowsInlineMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsInlineMediaPlayback:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsinlinepredictions
func (w_ WebViewConfiguration) AllowsInlinePredictions() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsInlinePredictions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsinlinepredictions
func (w_ WebViewConfiguration) SetAllowsInlinePredictions(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsInlinePredictions:"), value)
}


// A Boolean value that indicates whether HTML5 videos can play Picture in Picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowspictureinpicturemediaplayback
func (w_ WebViewConfiguration) AllowsPictureInPictureMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsPictureInPictureMediaPlayback"))
	return rv
}


// A Boolean value that indicates whether HTML5 videos can play Picture in Picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowspictureinpicturemediaplayback
func (w_ WebViewConfiguration) SetAllowsPictureInPictureMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsPictureInPictureMediaPlayback:"), value)
}


// The app name that appears in the user agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/applicationnameforuseragent
func (w_ WebViewConfiguration) ApplicationNameForUserAgent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("applicationNameForUserAgent"))
	return rv
}


// The app name that appears in the user agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/applicationnameforuseragent
func (w_ WebViewConfiguration) SetApplicationNameForUserAgent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setApplicationNameForUserAgent:"), value)
}


// The types of data detectors to apply to the web view’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/datadetectortypes
func (w_ WebViewConfiguration) DataDetectorTypes() DataDetectorTypes /* not a class type */ {
	rv := objc.Send[DataDetectorTypes](w_.ID, objc.Sel("dataDetectorTypes"))
	return rv
}


// The types of data detectors to apply to the web view’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/datadetectortypes
func (w_ WebViewConfiguration) SetDataDetectorTypes(value DataDetectorTypes /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDataDetectorTypes:"), value)
}


// The default preferences to use when loading and rendering content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/defaultwebpagepreferences
func (w_ WebViewConfiguration) DefaultWebpagePreferences() IWKWebpagePreferences {
	rv := objc.Send[WebpagePreferences](w_.ID, objc.Sel("defaultWebpagePreferences"))
	return rv
}


// The default preferences to use when loading and rendering content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/defaultwebpagepreferences
func (w_ WebViewConfiguration) SetDefaultWebpagePreferences(value IWKWebpagePreferences) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultWebpagePreferences:"), value)
}


// Deprecated property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediaplaybackallowsairplay
func (w_ WebViewConfiguration) MediaPlaybackAllowsAirPlay() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("mediaPlaybackAllowsAirPlay"))
	return rv
}


// Deprecated property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediaplaybackallowsairplay
func (w_ WebViewConfiguration) SetMediaPlaybackAllowsAirPlay(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaPlaybackAllowsAirPlay:"), value)
}


// Deprecated property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediaplaybackrequiresuseraction
func (w_ WebViewConfiguration) MediaPlaybackRequiresUserAction() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("mediaPlaybackRequiresUserAction"))
	return rv
}


// Deprecated property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediaplaybackrequiresuseraction
func (w_ WebViewConfiguration) SetMediaPlaybackRequiresUserAction(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaPlaybackRequiresUserAction:"), value)
}


// The media types that require a user gesture to begin playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediatypesrequiringuseractionforplayback
func (w_ WebViewConfiguration) MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes /* not a class type */ {
	rv := objc.Send[AudiovisualMediaTypes](w_.ID, objc.Sel("mediaTypesRequiringUserActionForPlayback"))
	return rv
}


// The media types that require a user gesture to begin playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediatypesrequiringuseractionforplayback
func (w_ WebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaTypesRequiringUserActionForPlayback:"), value)
}


// A Boolean value that indicates whether HTML5 videos require the user to start playing them (
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/requiresuseractionformediaplayback
func (w_ WebViewConfiguration) RequiresUserActionForMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("requiresUserActionForMediaPlayback"))
	return rv
}


// A Boolean value that indicates whether HTML5 videos require the user to start playing them (
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/requiresuseractionformediaplayback
func (w_ WebViewConfiguration) SetRequiresUserActionForMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequiresUserActionForMediaPlayback:"), value)
}


// The level of granularity with which the user can interactively select web view content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/selectiongranularity
func (w_ WebViewConfiguration) SelectionGranularity() SelectionGranularity /* not a class type */ {
	rv := objc.Send[SelectionGranularity](w_.ID, objc.Sel("selectionGranularity"))
	return rv
}


// The level of granularity with which the user can interactively select web view content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/selectiongranularity
func (w_ WebViewConfiguration) SetSelectionGranularity(value SelectionGranularity /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectionGranularity:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/showssystemscreentimeblockingview
func (w_ WebViewConfiguration) ShowsSystemScreenTimeBlockingView() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showsSystemScreenTimeBlockingView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/showssystemscreentimeblockingview
func (w_ WebViewConfiguration) SetShowsSystemScreenTimeBlockingView(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShowsSystemScreenTimeBlockingView:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/supportsadaptiveimageglyph
func (w_ WebViewConfiguration) SupportsAdaptiveImageGlyph() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("supportsAdaptiveImageGlyph"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/supportsadaptiveimageglyph
func (w_ WebViewConfiguration) SetSupportsAdaptiveImageGlyph(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSupportsAdaptiveImageGlyph:"), value)
}


// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/suppressesincrementalrendering
func (w_ WebViewConfiguration) SuppressesIncrementalRendering() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("suppressesIncrementalRendering"))
	return rv
}


// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/suppressesincrementalrendering
func (w_ WebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSuppressesIncrementalRendering:"), value)
}


// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/upgradeknownhoststohttps
func (w_ WebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("upgradeKnownHostsToHTTPS"))
	return rv
}


// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/upgradeknownhoststohttps
func (w_ WebViewConfiguration) SetUpgradeKnownHostsToHTTPS(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUpgradeKnownHostsToHTTPS:"), value)
}


// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (w_ WebViewConfiguration) UserContentController() IWKUserContentController {
	rv := objc.Send[UserContentController](w_.ID, objc.Sel("userContentController"))
	return rv
}


// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (w_ WebViewConfiguration) SetUserContentController(value IWKUserContentController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUserContentController:"), value)
}


// The directionality of user interface elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/userinterfacedirectionpolicy
func (w_ WebViewConfiguration) UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy /* not a class type */ {
	rv := objc.Send[UserInterfaceDirectionPolicy](w_.ID, objc.Sel("userInterfaceDirectionPolicy"))
	return rv
}


// The directionality of user interface elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/userinterfacedirectionpolicy
func (w_ WebViewConfiguration) SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUserInterfaceDirectionPolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w_ WebViewConfiguration) WebExtensionController() IWKWebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("webExtensionController"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w_ WebViewConfiguration) SetWebExtensionController(value IWKWebExtensionController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtensionController:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/writingtoolsbehavior
func (w_ WebViewConfiguration) WritingToolsBehavior() WritingToolsBehavior /* not a class type */ {
	rv := objc.Send[WritingToolsBehavior](w_.ID, objc.Sel("writingToolsBehavior"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/writingtoolsbehavior
func (w_ WebViewConfiguration) SetWritingToolsBehavior(value WritingToolsBehavior /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWritingToolsBehavior:"), value)
}


