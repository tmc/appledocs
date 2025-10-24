// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebViewConfiguration */

/* debug [class_header]: Header for WKWebViewConfiguration */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebViewConfiguration */
// An interface definition for the [WebViewConfiguration] class.
type IWebViewConfiguration interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebViewConfiguration */
	// properties:
	AllowsAirPlayForMediaPlayback() bool
	SetAllowsAirPlayForMediaPlayback(value bool)
	AllowsInlinePredictions() bool
	SetAllowsInlinePredictions(value bool)
	ApplicationNameForUserAgent() objc.IObject /* cross-framework: NSString */
	SetApplicationNameForUserAgent(value objc.IObject /* cross-framework: NSString */)
	DefaultWebpagePreferences() IWKWebpagePreferences
	SetDefaultWebpagePreferences(value IWKWebpagePreferences)
	LimitsNavigationsToAppBoundDomains() bool
	SetLimitsNavigationsToAppBoundDomains(value bool)
	MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes
	SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes)
	Preferences() IWKPreferences
	SetPreferences(value IWKPreferences)
	ProcessPool() IWKProcessPool
	SetProcessPool(value IWKProcessPool)
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
	UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy
	SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy)
	WebExtensionController() IWKWebExtensionController
	SetWebExtensionController(value IWKWebExtensionController)
	WebsiteDataStore() IWKWebsiteDataStore
	SetWebsiteDataStore(value IWKWebsiteDataStore)
	WritingToolsBehavior() WritingToolsBehavior /* not a class type */
	SetWritingToolsBehavior(value WritingToolsBehavior /* not a class type */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebViewConfiguration */
	// methods:
	SetURLSchemeHandlerForURLScheme(urlSchemeHandler unsafe.Pointer, urlScheme objc.IObject /* cross-framework: NSString */)
	UrlSchemeHandlerForURLScheme(urlScheme objc.IObject /* cross-framework: NSString */) unsafe.Pointer
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebViewConfiguration */
// Alloc allocates a new instance without initialization.
func (wc _WebViewConfigurationClass) Alloc() WebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebViewConfiguration */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebViewConfiguration */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebViewConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebViewConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebViewConfiguration */

// Registers an object to load resources associated with the specified URL scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/setURLSchemeHandler(_:forURLScheme:)
func (w_ WebViewConfiguration) SetURLSchemeHandlerForURLScheme(urlSchemeHandler unsafe.Pointer, urlScheme objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setURLSchemeHandler:forURLScheme:"), urlSchemeHandler, urlScheme)
} /* debug [instance_methods/method]: SetURLSchemeHandlerForURLScheme */

// Returns the currently registered handler object for the specified URL scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/urlSchemeHandler(forURLScheme:)
func (w_ WebViewConfiguration) UrlSchemeHandlerForURLScheme(urlScheme objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("urlSchemeHandlerForURLScheme:"), urlScheme)
	return rv
} /* debug [instance_methods/method]: UrlSchemeHandlerForURLScheme */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebViewConfiguration */

// A Boolean value that indicates whether the web view allows media playback over AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsAirPlayForMediaPlayback
func (w_ WebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsAirPlayForMediaPlayback"))
	return rv
} /* debug [instance_properties/getter]: allowsAirPlayForMediaPlayback */

// A Boolean value that indicates whether the web view allows media playback over AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsAirPlayForMediaPlayback
func (w_ WebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsAirPlayForMediaPlayback:"), value)
} /* debug [instance_properties/setter]: allowsAirPlayForMediaPlayback */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsInlinePredictions
func (w_ WebViewConfiguration) AllowsInlinePredictions() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsInlinePredictions"))
	return rv
} /* debug [instance_properties/getter]: allowsInlinePredictions */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsInlinePredictions
func (w_ WebViewConfiguration) SetAllowsInlinePredictions(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsInlinePredictions:"), value)
} /* debug [instance_properties/setter]: allowsInlinePredictions */

// The app name that appears in the user agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/applicationNameForUserAgent
func (w_ WebViewConfiguration) ApplicationNameForUserAgent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("applicationNameForUserAgent"))
	return rv
} /* debug [instance_properties/getter]: applicationNameForUserAgent */

// The app name that appears in the user agent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/applicationNameForUserAgent
func (w_ WebViewConfiguration) SetApplicationNameForUserAgent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setApplicationNameForUserAgent:"), value)
} /* debug [instance_properties/setter]: applicationNameForUserAgent */

// The default preferences to use when loading and rendering content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/defaultWebpagePreferences
func (w_ WebViewConfiguration) DefaultWebpagePreferences() IWKWebpagePreferences {
	rv := objc.Send[WebpagePreferences](w_.ID, objc.Sel("defaultWebpagePreferences"))
	return rv
} /* debug [instance_properties/getter]: defaultWebpagePreferences */

// The default preferences to use when loading and rendering content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/defaultWebpagePreferences
func (w_ WebViewConfiguration) SetDefaultWebpagePreferences(value IWKWebpagePreferences) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultWebpagePreferences:"), value)
} /* debug [instance_properties/setter]: defaultWebpagePreferences */

// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/limitsNavigationsToAppBoundDomains
func (w_ WebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("limitsNavigationsToAppBoundDomains"))
	return rv
} /* debug [instance_properties/getter]: limitsNavigationsToAppBoundDomains */

// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/limitsNavigationsToAppBoundDomains
func (w_ WebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setLimitsNavigationsToAppBoundDomains:"), value)
} /* debug [instance_properties/setter]: limitsNavigationsToAppBoundDomains */

// The media types that require a user gesture to begin playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/mediaTypesRequiringUserActionForPlayback
func (w_ WebViewConfiguration) MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes {
	rv := objc.Send[AudiovisualMediaTypes](w_.ID, objc.Sel("mediaTypesRequiringUserActionForPlayback"))
	return rv
} /* debug [instance_properties/getter]: mediaTypesRequiringUserActionForPlayback */

// The media types that require a user gesture to begin playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/mediaTypesRequiringUserActionForPlayback
func (w_ WebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMediaTypesRequiringUserActionForPlayback:"), value)
} /* debug [instance_properties/setter]: mediaTypesRequiringUserActionForPlayback */

// The object that manages the preference-related settings for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/preferences
func (w_ WebViewConfiguration) Preferences() IWKPreferences {
	rv := objc.Send[Preferences](w_.ID, objc.Sel("preferences"))
	return rv
} /* debug [instance_properties/getter]: preferences */

// The object that manages the preference-related settings for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/preferences
func (w_ WebViewConfiguration) SetPreferences(value IWKPreferences) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferences:"), value)
} /* debug [instance_properties/setter]: preferences */

// The object that coordinates the processes the web view uses to render its web content and execute scripts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/processPool
func (w_ WebViewConfiguration) ProcessPool() IWKProcessPool {
	rv := objc.Send[ProcessPool](w_.ID, objc.Sel("processPool"))
	return rv
} /* debug [instance_properties/getter]: processPool */

// The object that coordinates the processes the web view uses to render its web content and execute scripts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/processPool
func (w_ WebViewConfiguration) SetProcessPool(value IWKProcessPool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setProcessPool:"), value)
} /* debug [instance_properties/setter]: processPool */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/showsSystemScreenTimeBlockingView
func (w_ WebViewConfiguration) ShowsSystemScreenTimeBlockingView() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showsSystemScreenTimeBlockingView"))
	return rv
} /* debug [instance_properties/getter]: showsSystemScreenTimeBlockingView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/showsSystemScreenTimeBlockingView
func (w_ WebViewConfiguration) SetShowsSystemScreenTimeBlockingView(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShowsSystemScreenTimeBlockingView:"), value)
} /* debug [instance_properties/setter]: showsSystemScreenTimeBlockingView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/supportsAdaptiveImageGlyph
func (w_ WebViewConfiguration) SupportsAdaptiveImageGlyph() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("supportsAdaptiveImageGlyph"))
	return rv
} /* debug [instance_properties/getter]: supportsAdaptiveImageGlyph */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/supportsAdaptiveImageGlyph
func (w_ WebViewConfiguration) SetSupportsAdaptiveImageGlyph(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSupportsAdaptiveImageGlyph:"), value)
} /* debug [instance_properties/setter]: supportsAdaptiveImageGlyph */

// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/suppressesIncrementalRendering
func (w_ WebViewConfiguration) SuppressesIncrementalRendering() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("suppressesIncrementalRendering"))
	return rv
} /* debug [instance_properties/getter]: suppressesIncrementalRendering */

// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/suppressesIncrementalRendering
func (w_ WebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSuppressesIncrementalRendering:"), value)
} /* debug [instance_properties/setter]: suppressesIncrementalRendering */

// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/upgradeKnownHostsToHTTPS
func (w_ WebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("upgradeKnownHostsToHTTPS"))
	return rv
} /* debug [instance_properties/getter]: upgradeKnownHostsToHTTPS */

// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/upgradeKnownHostsToHTTPS
func (w_ WebViewConfiguration) SetUpgradeKnownHostsToHTTPS(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUpgradeKnownHostsToHTTPS:"), value)
} /* debug [instance_properties/setter]: upgradeKnownHostsToHTTPS */

// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/userContentController
func (w_ WebViewConfiguration) UserContentController() IWKUserContentController {
	rv := objc.Send[UserContentController](w_.ID, objc.Sel("userContentController"))
	return rv
} /* debug [instance_properties/getter]: userContentController */

// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/userContentController
func (w_ WebViewConfiguration) SetUserContentController(value IWKUserContentController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUserContentController:"), value)
} /* debug [instance_properties/setter]: userContentController */

// The directionality of user interface elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/userInterfaceDirectionPolicy
func (w_ WebViewConfiguration) UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy {
	rv := objc.Send[UserInterfaceDirectionPolicy](w_.ID, objc.Sel("userInterfaceDirectionPolicy"))
	return rv
} /* debug [instance_properties/getter]: userInterfaceDirectionPolicy */

// The directionality of user interface elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/userInterfaceDirectionPolicy
func (w_ WebViewConfiguration) SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUserInterfaceDirectionPolicy:"), value)
} /* debug [instance_properties/setter]: userInterfaceDirectionPolicy */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/webExtensionController
func (w_ WebViewConfiguration) WebExtensionController() IWKWebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("webExtensionController"))
	return rv
} /* debug [instance_properties/getter]: webExtensionController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/webExtensionController
func (w_ WebViewConfiguration) SetWebExtensionController(value IWKWebExtensionController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtensionController:"), value)
} /* debug [instance_properties/setter]: webExtensionController */

// The object you use to get and set the site’s cookies and to track the cached data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/websiteDataStore
func (w_ WebViewConfiguration) WebsiteDataStore() IWKWebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](w_.ID, objc.Sel("websiteDataStore"))
	return rv
} /* debug [instance_properties/getter]: websiteDataStore */

// The object you use to get and set the site’s cookies and to track the cached data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/websiteDataStore
func (w_ WebViewConfiguration) SetWebsiteDataStore(value IWKWebsiteDataStore) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebsiteDataStore:"), value)
} /* debug [instance_properties/setter]: websiteDataStore */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/writingToolsBehavior
func (w_ WebViewConfiguration) WritingToolsBehavior() WritingToolsBehavior /* not a class type */ {
	rv := objc.Send[WritingToolsBehavior](w_.ID, objc.Sel("writingToolsBehavior"))
	return rv
} /* debug [instance_properties/getter]: writingToolsBehavior */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/writingToolsBehavior
func (w_ WebViewConfiguration) SetWritingToolsBehavior(value WritingToolsBehavior /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWritingToolsBehavior:"), value)
} /* debug [instance_properties/setter]: writingToolsBehavior */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKWebViewConfiguration */
