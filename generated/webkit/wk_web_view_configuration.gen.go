// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	SetURLSchemeHandlerForURLScheme(urlSchemeHandler objc.ID, urlScheme string)
}

// A collection of properties that you use to initialize a web view.
//
// A object provides information about how to configure a object. Use your configuration object to specify: The initial cookies to make available to your web content Handlers for any custom URL schemes your web content uses Settings for how to handle media content Information about how to manage selections within the web view Custom scripts to inject into the webpage Custom rules that determine how to render content You create a object in your code, configure its properties, and pass it to the initializer of your object. The web view incorporates your configuration settings only at creation time; you cannot change those settings dynamically later.
//
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
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/setURLSchemeHandler(_:forURLScheme:)
func (w_ WebViewConfiguration) SetURLSchemeHandlerForURLScheme(urlSchemeHandler objc.ID, urlScheme string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setURLSchemeHandler:forURLScheme:"), urlSchemeHandler, objc.String(urlScheme))
}

// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/limitsNavigationsToAppBoundDomains
func (w_ WebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("limitsNavigationsToAppBoundDomains"))
	return rv
}


// SetLimitsNavigationsToAppBoundDomains sets the value of the limitsNavigationsToAppBoundDomains property.
// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/limitsNavigationsToAppBoundDomains
func (w_ WebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setLimitsNavigationsToAppBoundDomains:"), value)
}

// The object that manages the preference-related settings for the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/preferences
func (w_ WebViewConfiguration) Preferences() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("preferences"))
	return rv
}


// SetPreferences sets the value of the preferences property.
// The object that manages the preference-related settings for the web view.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/preferences
func (w_ WebViewConfiguration) SetPreferences(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferences:"), value)
}

// The object that coordinates the processes the web view uses to render its web content and execute scripts.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/processPool
func (w_ WebViewConfiguration) ProcessPool() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("processPool"))
	return rv
}


// SetProcessPool sets the value of the processPool property.
// The object that coordinates the processes the web view uses to render its web content and execute scripts.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/processPool
func (w_ WebViewConfiguration) SetProcessPool(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setProcessPool:"), value)
}

// The object you use to get and set the site’s cookies and to track the cached data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/websiteDataStore
func (w_ WebViewConfiguration) WebsiteDataStore() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("websiteDataStore"))
	return rv
}


// SetWebsiteDataStore sets the value of the websiteDataStore property.
// The object you use to get and set the site’s cookies and to track the cached data objects.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/websiteDataStore
func (w_ WebViewConfiguration) SetWebsiteDataStore(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebsiteDataStore:"), value)
}

// A Boolean value that indicates whether the web view allows media playback over AirPlay.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsairplayformediaplayback
func (w_ WebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsAirPlayForMediaPlayback"))
	return rv
}


// SetAllowsAirPlayForMediaPlayback sets the value of the allowsAirPlayForMediaPlayback property.
// A Boolean value that indicates whether the web view allows media playback over AirPlay.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsairplayformediaplayback
func (w_ WebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsAirPlayForMediaPlayback:"), value)
}

// A Boolean value that indicates whether HTML5 videos play inline or use the native full-screen controller.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsinlinemediaplayback
func (w_ WebViewConfiguration) AllowsInlineMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsInlineMediaPlayback"))
	return rv
}


// SetAllowsInlineMediaPlayback sets the value of the allowsInlineMediaPlayback property.
// A Boolean value that indicates whether HTML5 videos play inline or use the native full-screen controller.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsinlinemediaplayback
func (w_ WebViewConfiguration) SetAllowsInlineMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsInlineMediaPlayback:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsinlinepredictions
func (w_ WebViewConfiguration) AllowsInlinePredictions() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsInlinePredictions"))
	return rv
}


// SetAllowsInlinePredictions sets the value of the allowsInlinePredictions property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowsinlinepredictions
func (w_ WebViewConfiguration) SetAllowsInlinePredictions(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsInlinePredictions:"), value)
}

// A Boolean value that indicates whether HTML5 videos can play Picture in Picture.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowspictureinpicturemediaplayback
func (w_ WebViewConfiguration) AllowsPictureInPictureMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsPictureInPictureMediaPlayback"))
	return rv
}


// SetAllowsPictureInPictureMediaPlayback sets the value of the allowsPictureInPictureMediaPlayback property.
// A Boolean value that indicates whether HTML5 videos can play Picture in Picture.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/allowspictureinpicturemediaplayback
func (w_ WebViewConfiguration) SetAllowsPictureInPictureMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsPictureInPictureMediaPlayback:"), value)
}

// The app name that appears in the user agent string.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/applicationnameforuseragent
func (w_ WebViewConfiguration) ApplicationNameForUserAgent() string {
	rv := objc.Send[string](w_.ID, objc.Sel("applicationNameForUserAgent"))
	return rv
}


// SetApplicationNameForUserAgent sets the value of the applicationNameForUserAgent property.
// The app name that appears in the user agent string.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/applicationnameforuseragent
func (w_ WebViewConfiguration) SetApplicationNameForUserAgent(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setApplicationNameForUserAgent:"), objc.String(value))
}

// The types of data detectors to apply to the web view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/datadetectortypes
func (w_ WebViewConfiguration) DataDetectorTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("dataDetectorTypes"))
	return rv
}


// SetDataDetectorTypes sets the value of the dataDetectorTypes property.
// The types of data detectors to apply to the web view’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/datadetectortypes
func (w_ WebViewConfiguration) SetDataDetectorTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDataDetectorTypes:"), value)
}

// The default preferences to use when loading and rendering content.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/defaultwebpagepreferences
func (w_ WebViewConfiguration) DefaultWebpagePreferences() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("defaultWebpagePreferences"))
	return rv
}


// SetDefaultWebpagePreferences sets the value of the defaultWebpagePreferences property.
// The default preferences to use when loading and rendering content.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/defaultwebpagepreferences
func (w_ WebViewConfiguration) SetDefaultWebpagePreferences(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultWebpagePreferences:"), value)
}

// A Boolean value that determines whether a web view allows scaling of the webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/ignoresviewportscalelimits
func (w_ WebViewConfiguration) IgnoresViewportScaleLimits() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("ignoresViewportScaleLimits"))
	return rv
}


// SetIgnoresViewportScaleLimits sets the value of the ignoresViewportScaleLimits property.
// A Boolean value that determines whether a web view allows scaling of the webpage.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/ignoresviewportscalelimits
func (w_ WebViewConfiguration) SetIgnoresViewportScaleLimits(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIgnoresViewportScaleLimits:"), value)
}

// Deprecated property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediaplaybackallowsairplay
func (w_ WebViewConfiguration) MediaPlaybackAllowsAirPlay() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("mediaPlaybackAllowsAirPlay"))
	return rv
}


// SetMediaPlaybackAllowsAirPlay sets the value of the mediaPlaybackAllowsAirPlay property.
// Deprecated property.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediaplaybackallowsairplay
func (w_ WebViewConfiguration) SetMediaPlaybackAllowsAirPlay(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaPlaybackAllowsAirPlay:"), value)
}

// Deprecated property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediaplaybackrequiresuseraction
func (w_ WebViewConfiguration) MediaPlaybackRequiresUserAction() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("mediaPlaybackRequiresUserAction"))
	return rv
}


// SetMediaPlaybackRequiresUserAction sets the value of the mediaPlaybackRequiresUserAction property.
// Deprecated property.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediaplaybackrequiresuseraction
func (w_ WebViewConfiguration) SetMediaPlaybackRequiresUserAction(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaPlaybackRequiresUserAction:"), value)
}

// The media types that require a user gesture to begin playing.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediatypesrequiringuseractionforplayback
func (w_ WebViewConfiguration) MediaTypesRequiringUserActionForPlayback() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("mediaTypesRequiringUserActionForPlayback"))
	return rv
}


// SetMediaTypesRequiringUserActionForPlayback sets the value of the mediaTypesRequiringUserActionForPlayback property.
// The media types that require a user gesture to begin playing.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/mediatypesrequiringuseractionforplayback
func (w_ WebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaTypesRequiringUserActionForPlayback:"), value)
}

// A Boolean value that indicates whether HTML5 videos require the user to start playing them (
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/requiresuseractionformediaplayback
func (w_ WebViewConfiguration) RequiresUserActionForMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("requiresUserActionForMediaPlayback"))
	return rv
}


// SetRequiresUserActionForMediaPlayback sets the value of the requiresUserActionForMediaPlayback property.
// A Boolean value that indicates whether HTML5 videos require the user to start playing them (

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/requiresuseractionformediaplayback
func (w_ WebViewConfiguration) SetRequiresUserActionForMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequiresUserActionForMediaPlayback:"), value)
}

// The level of granularity with which the user can interactively select web view content.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/selectiongranularity
func (w_ WebViewConfiguration) SelectionGranularity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("selectionGranularity"))
	return rv
}


// SetSelectionGranularity sets the value of the selectionGranularity property.
// The level of granularity with which the user can interactively select web view content.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/selectiongranularity
func (w_ WebViewConfiguration) SetSelectionGranularity(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectionGranularity:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/showssystemscreentimeblockingview
func (w_ WebViewConfiguration) ShowsSystemScreenTimeBlockingView() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showsSystemScreenTimeBlockingView"))
	return rv
}


// SetShowsSystemScreenTimeBlockingView sets the value of the showsSystemScreenTimeBlockingView property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/showssystemscreentimeblockingview
func (w_ WebViewConfiguration) SetShowsSystemScreenTimeBlockingView(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShowsSystemScreenTimeBlockingView:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/supportsadaptiveimageglyph
func (w_ WebViewConfiguration) SupportsAdaptiveImageGlyph() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("supportsAdaptiveImageGlyph"))
	return rv
}


// SetSupportsAdaptiveImageGlyph sets the value of the supportsAdaptiveImageGlyph property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/supportsadaptiveimageglyph
func (w_ WebViewConfiguration) SetSupportsAdaptiveImageGlyph(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSupportsAdaptiveImageGlyph:"), value)
}

// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/suppressesincrementalrendering
func (w_ WebViewConfiguration) SuppressesIncrementalRendering() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("suppressesIncrementalRendering"))
	return rv
}


// SetSuppressesIncrementalRendering sets the value of the suppressesIncrementalRendering property.
// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/suppressesincrementalrendering
func (w_ WebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSuppressesIncrementalRendering:"), value)
}

// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/upgradeknownhoststohttps
func (w_ WebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("upgradeKnownHostsToHTTPS"))
	return rv
}


// SetUpgradeKnownHostsToHTTPS sets the value of the upgradeKnownHostsToHTTPS property.
// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/upgradeknownhoststohttps
func (w_ WebViewConfiguration) SetUpgradeKnownHostsToHTTPS(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUpgradeKnownHostsToHTTPS:"), value)
}

// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (w_ WebViewConfiguration) UserContentController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("userContentController"))
	return rv
}


// SetUserContentController sets the value of the userContentController property.
// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (w_ WebViewConfiguration) SetUserContentController(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUserContentController:"), value)
}

// The directionality of user interface elements.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/userinterfacedirectionpolicy
func (w_ WebViewConfiguration) UserInterfaceDirectionPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("userInterfaceDirectionPolicy"))
	return rv
}


// SetUserInterfaceDirectionPolicy sets the value of the userInterfaceDirectionPolicy property.
// The directionality of user interface elements.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/userinterfacedirectionpolicy
func (w_ WebViewConfiguration) SetUserInterfaceDirectionPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUserInterfaceDirectionPolicy:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w_ WebViewConfiguration) WebExtensionController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("webExtensionController"))
	return rv
}


// SetWebExtensionController sets the value of the webExtensionController property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w_ WebViewConfiguration) SetWebExtensionController(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtensionController:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/writingtoolsbehavior
func (w_ WebViewConfiguration) WritingToolsBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("writingToolsBehavior"))
	return rv
}


// SetWritingToolsBehavior sets the value of the writingToolsBehavior property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/writingtoolsbehavior
func (w_ WebViewConfiguration) SetWritingToolsBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWritingToolsBehavior:"), value)
}



