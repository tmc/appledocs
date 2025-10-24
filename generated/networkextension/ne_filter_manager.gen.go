// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterManager */


/* debug [class_header]: Header for NEFilterManager */
// The class instance for the [NEFilterManager] class.
var (
	NEFilterManagerClass     _NEFilterManagerClass
	NEFilterManagerClassOnce sync.Once
)

func getNEFilterManagerClass() _NEFilterManagerClass {
	NEFilterManagerClassOnce.Do(func() {
		NEFilterManagerClass = _NEFilterManagerClass{objc.GetClass("NEFilterManager")}
	})
	return NEFilterManagerClass
}

type _NEFilterManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterManager */
// An interface definition for the [NEFilterManager] class.
type INEFilterManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFilterManager */
	// properties:
	DisableEncryptedDNSSettings() bool
	SetDisableEncryptedDNSSettings(value bool)
	Grade() NEFilterManagerGrade
	SetGrade(value NEFilterManagerGrade)
	Enabled() bool
	SetEnabled(value bool)
	LocalizedDescription() objc.IObject /* cross-framework: NSString */
	SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */)
	ProviderConfiguration() INEFilterProviderConfiguration
	SetProviderConfiguration(value INEFilterProviderConfiguration)
	NEFilterErrorDomain() objc.IObject /* cross-framework: NSString */
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterManager */
	// methods:
	LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterManager */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterManagerClass) Alloc() NEFilterManager {
	rv := objc.Send[NEFilterManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterManagerClass) New() NEFilterManager {
	rv := objc.Send[NEFilterManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterManager) Init() NEFilterManager {
	rv := objc.Send[NEFilterManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterManager) Autorelease() NEFilterManager {
	rv := objc.Send[NEFilterManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterManager creates a new NEFilterManager instance.
func NewNEFilterManager() NEFilterManager {
	return getNEFilterManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterManager */
// An object to create and manage a content filter’s configuration.
//
// Each app is allowed to create a single filter configuration. The class has a class method ( ) that provides access to a single instance. This single instance corresponds to a single filter configuration. The filter configuration is stored in the Network Extension preferences which are managed by the Network Extension framework. The filter configuration must be explicitly loaded into memory from the Network Extension preferences before it can be used, and any changes must be explicitly saved to the Network Extension preferences before taking effect on the system.


// An object to create and manage a content filter’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager
type NEFilterManager struct {
	objectivec.Object
}

// NEFilterManagerFrom constructs a [NEFilterManager] from an unsafe.Pointer.
//
// An object to create and manage a content filter’s configuration.
func NEFilterManagerFrom(ptr unsafe.Pointer) NEFilterManager {
	return NEFilterManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterManager */

// Access the single instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/shared()
func (nc _NEFilterManagerClass) SharedManager() NEFilterManager {
	rv := objc.Send[NEFilterManager](objc.ID(nc.class), objc.Sel("sharedManager"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedManager) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterManager */

// Load the filter configuration from the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/loadFromPreferences(completionHandler:)
func (n_ NEFilterManager) LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadFromPreferencesWithCompletionHandler */


// Remove the filter configuration from the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/removeFromPreferences(completionHandler:)
func (n_ NEFilterManager) RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: RemoveFromPreferencesWithCompletionHandler */


// Save the filter configuration in the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/saveToPreferences(completionHandler:)
func (n_ NEFilterManager) SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: SaveToPreferencesWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/disableEncryptedDNSSettings
func (n_ NEFilterManager) DisableEncryptedDNSSettings() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disableEncryptedDNSSettings"))
	return rv
}/* debug [instance_properties/getter]: disableEncryptedDNSSettings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/disableEncryptedDNSSettings
func (n_ NEFilterManager) SetDisableEncryptedDNSSettings(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisableEncryptedDNSSettings:"), value)
}/* debug [instance_properties/setter]: disableEncryptedDNSSettings */


// The grade of the filter, which determines when it acts relative to other filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/grade-swift.property
func (n_ NEFilterManager) Grade() NEFilterManagerGrade {
	rv := objc.Send[NEFilterManagerGrade](n_.ID, objc.Sel("grade"))
	return rv
}/* debug [instance_properties/getter]: grade */


// The grade of the filter, which determines when it acts relative to other filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/grade-swift.property
func (n_ NEFilterManager) SetGrade(value NEFilterManagerGrade) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGrade:"), value)
}/* debug [instance_properties/setter]: grade */


// A Boolean used to toggle the enabled state of the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/isEnabled
func (n_ NEFilterManager) Enabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean used to toggle the enabled state of the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/isEnabled
func (n_ NEFilterManager) SetEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A string containing a description of the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/localizedDescription
func (n_ NEFilterManager) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}/* debug [instance_properties/getter]: localizedDescription */


// A string containing a description of the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/localizedDescription
func (n_ NEFilterManager) SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}/* debug [instance_properties/setter]: localizedDescription */


// A object containing the filter configuration settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/providerConfiguration
func (n_ NEFilterManager) ProviderConfiguration() INEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}/* debug [instance_properties/getter]: providerConfiguration */


// A object containing the filter configuration settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/providerConfiguration
func (n_ NEFilterManager) SetProviderConfiguration(value INEFilterProviderConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}/* debug [instance_properties/setter]: providerConfiguration */


// The domain for errors resulting from calls to the filter manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltererrordomain
func (n_ NEFilterManager) NEFilterErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: NEFilterErrorDomain */


// A Boolean used to toggle the enabled state of the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/isenabled
func (n_ NEFilterManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean used to toggle the enabled state of the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/isenabled
func (n_ NEFilterManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterManager */



