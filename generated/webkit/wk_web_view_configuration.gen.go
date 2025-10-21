// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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



