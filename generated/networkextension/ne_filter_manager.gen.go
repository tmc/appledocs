// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NEFilterManager] class.
type INEFilterManager interface {
	objectivec.IObject
	// properties:
	NEFilterErrorDomain() objc.IObject /* cross-framework: NSString */
	DisableEncryptedDNSSettings() bool
	SetDisableEncryptedDNSSettings(value bool)
	Grade() unsafe.Pointer
	SetGrade(value unsafe.Pointer)
	IsEnabled() bool
	SetIsEnabled(value bool)
	LocalizedDescription() objc.IObject /* cross-framework: NSString */
	SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */)
	ProviderConfiguration() INEFilterProviderConfiguration
	SetProviderConfiguration(value INEFilterProviderConfiguration)
	// methods:
	RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (nc _NEFilterManagerClass) Alloc() NEFilterManager {
	rv := objc.Send[NEFilterManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Remove the filter configuration from the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/removeFromPreferences(completionHandler:)
func (n_ NEFilterManager) RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), completionHandler)
}


// The domain for errors resulting from calls to the filter manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltererrordomain
func (n_ NEFilterManager) NEFilterErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterErrorDomain"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/disableencrypteddnssettings
func (n_ NEFilterManager) DisableEncryptedDNSSettings() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disableEncryptedDNSSettings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/disableencrypteddnssettings
func (n_ NEFilterManager) SetDisableEncryptedDNSSettings(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisableEncryptedDNSSettings:"), value)
}


// The grade of the filter, which determines when it acts relative to other filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/grade-swift.property
func (n_ NEFilterManager) Grade() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("grade"))
	return rv
}


// The grade of the filter, which determines when it acts relative to other filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/grade-swift.property
func (n_ NEFilterManager) SetGrade(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGrade:"), value)
}


// A Boolean used to toggle the enabled state of the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/isenabled
func (n_ NEFilterManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean used to toggle the enabled state of the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/isenabled
func (n_ NEFilterManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}


// A string containing a description of the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/localizeddescription
func (n_ NEFilterManager) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string containing a description of the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/localizeddescription
func (n_ NEFilterManager) SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/providerconfiguration
func (n_ NEFilterManager) ProviderConfiguration() INEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltermanager/providerconfiguration
func (n_ NEFilterManager) SetProviderConfiguration(value INEFilterProviderConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}



