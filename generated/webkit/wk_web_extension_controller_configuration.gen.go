// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionControllerConfiguration */


/* debug [class_header]: Header for WKWebExtensionControllerConfiguration */
// The class instance for the [WebExtensionControllerConfiguration] class.
var (
	WebExtensionControllerConfigurationClass     _WebExtensionControllerConfigurationClass
	WebExtensionControllerConfigurationClassOnce sync.Once
)

func getWebExtensionControllerConfigurationClass() _WebExtensionControllerConfigurationClass {
	WebExtensionControllerConfigurationClassOnce.Do(func() {
		WebExtensionControllerConfigurationClass = _WebExtensionControllerConfigurationClass{objc.GetClass("WKWebExtensionControllerConfiguration")}
	})
	return WebExtensionControllerConfigurationClass
}

type _WebExtensionControllerConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebExtensionControllerConfiguration */
// An interface definition for the [WebExtensionControllerConfiguration] class.
type IWebExtensionControllerConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebExtensionControllerConfiguration */
	// properties:
	DefaultWebsiteDataStore() IWKWebsiteDataStore
	SetDefaultWebsiteDataStore(value IWKWebsiteDataStore)
	Identifier() foundation.UUID
	Persistent() bool
	WebViewConfiguration() IWKWebViewConfiguration
	SetWebViewConfiguration(value IWKWebViewConfiguration)
	IsPersistent() bool
	SetIsPersistent(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebExtensionControllerConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebExtensionControllerConfiguration */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionControllerConfigurationClass) Alloc() WebExtensionControllerConfiguration {
	rv := objc.Send[WebExtensionControllerConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionControllerConfigurationClass) New() WebExtensionControllerConfiguration {
	rv := objc.Send[WebExtensionControllerConfiguration](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionControllerConfiguration) Init() WebExtensionControllerConfiguration {
	rv := objc.Send[WebExtensionControllerConfiguration](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionControllerConfiguration) Autorelease() WebExtensionControllerConfiguration {
	rv := objc.Send[WebExtensionControllerConfiguration](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionControllerConfiguration creates a new WebExtensionControllerConfiguration instance.
func NewWebExtensionControllerConfiguration() WebExtensionControllerConfiguration {
	return getWebExtensionControllerConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebExtensionControllerConfiguration */
// A object with which to initialize a web extension controller.
//
// Contains properties used to configure a .


// A object with which to initialize a web extension controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class
type WebExtensionControllerConfiguration struct {
	objectivec.Object
}

// WebExtensionControllerConfigurationFrom constructs a [WebExtensionControllerConfiguration] from an unsafe.Pointer.
//
// A object with which to initialize a web extension controller.
func WebExtensionControllerConfigurationFrom(ptr unsafe.Pointer) WebExtensionControllerConfiguration {
	return WebExtensionControllerConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebExtensionControllerConfiguration */

// Returns a new configuration that is persistent and unique for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/init(identifier:)
func NewWebExtensionControllerConfigurationWithIdentifier(identifier foundation.UUID) WebExtensionControllerConfiguration {
	rv := objc.Send[WebExtensionControllerConfiguration](objc.ID(getWebExtensionControllerConfigurationClass().class), objc.Sel("configurationWithIdentifier:"), identifier)
	return rv
}/* debug [class_init_methods/constructor]: NewWebExtensionControllerConfigurationWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebExtensionControllerConfiguration */

// Returns a new default configuration that is persistent and not unique.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/default()
func (wc _WebExtensionControllerConfigurationClass) DefaultConfiguration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("defaultConfiguration"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultConfiguration) */


// Returns a new configuration that is persistent and unique for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/init(identifier:)
func (wc _WebExtensionControllerConfigurationClass) ConfigurationWithIdentifier(identifier foundation.UUID) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("configurationWithIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithIdentifier) */


// Returns a new non-persistent configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/nonPersistent()
func (wc _WebExtensionControllerConfigurationClass) NonPersistentConfiguration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("nonPersistentConfiguration"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NonPersistentConfiguration) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebExtensionControllerConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebExtensionControllerConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebExtensionControllerConfiguration */

// The default data store for website data and cookie access in extension contexts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/defaultWebsiteDataStore
func (w_ WebExtensionControllerConfiguration) DefaultWebsiteDataStore() IWKWebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](w_.ID, objc.Sel("defaultWebsiteDataStore"))
	return rv
}/* debug [instance_properties/getter]: defaultWebsiteDataStore */


// The default data store for website data and cookie access in extension contexts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/defaultWebsiteDataStore
func (w_ WebExtensionControllerConfiguration) SetDefaultWebsiteDataStore(value IWKWebsiteDataStore) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultWebsiteDataStore:"), value)
}/* debug [instance_properties/setter]: defaultWebsiteDataStore */


// The unique identifier used for persistent configuration storage, or when it is the default or not persistent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/identifier
func (w_ WebExtensionControllerConfiguration) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](w_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value indicating if this context will write data to the the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/isPersistent
func (w_ WebExtensionControllerConfiguration) Persistent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("persistent"))
	return rv
}/* debug [instance_properties/getter]: persistent */


// The web view configuration to be used as a basis for configuring web views in extension contexts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/webViewConfiguration
func (w_ WebExtensionControllerConfiguration) WebViewConfiguration() IWKWebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](w_.ID, objc.Sel("webViewConfiguration"))
	return rv
}/* debug [instance_properties/getter]: webViewConfiguration */


// The web view configuration to be used as a basis for configuring web views in extension contexts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/webViewConfiguration
func (w_ WebExtensionControllerConfiguration) SetWebViewConfiguration(value IWKWebViewConfiguration) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebViewConfiguration:"), value)
}/* debug [instance_properties/setter]: webViewConfiguration */


// A Boolean value indicating if this context will write data to the the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/configuration-swift.class/ispersistent
func (w_ WebExtensionControllerConfiguration) IsPersistent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isPersistent"))
	return rv
}/* debug [instance_properties/getter]: isPersistent */


// A Boolean value indicating if this context will write data to the the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/configuration-swift.class/ispersistent
func (w_ WebExtensionControllerConfiguration) SetIsPersistent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsPersistent:"), value)
}/* debug [instance_properties/setter]: isPersistent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebExtensionControllerConfiguration */


