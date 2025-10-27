// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	DisableEncryptedDNSSettings() bool
	SetDisableEncryptedDNSSettings(value bool)
	Grade() NEFilterManagerGrade
	SetGrade(value NEFilterManagerGrade)
	Enabled() bool
	SetEnabled(value bool)
	LocalizedDescription() foundation.foundation.INSString
	SetLocalizedDescription(value foundation.foundation.INSString)
	ProviderConfiguration() INEFilterProviderConfiguration
	SetProviderConfiguration(value INEFilterProviderConfiguration)
	NEFilterErrorDomain() foundation.foundation.INSString
	IsEnabled() bool
	SetIsEnabled(value bool)


	

	// methods:
	LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)


}





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










// Access the single instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/shared()
func (nc _NEFilterManagerClass) SharedManager() NEFilterManager {
	rv := objc.Send[NEFilterManager](objc.ID(nc.class), objc.Sel("sharedManager"))
	return rv
}












// Load the filter configuration from the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/loadFromPreferences(completionHandler:)
func (n_ NEFilterManager) LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Remove the filter configuration from the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/removeFromPreferences(completionHandler:)
func (n_ NEFilterManager) RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Save the filter configuration in the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/saveToPreferences(completionHandler:)
func (n_ NEFilterManager) SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/disableEncryptedDNSSettings
func (n_ NEFilterManager) DisableEncryptedDNSSettings() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disableEncryptedDNSSettings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/disableEncryptedDNSSettings
func (n_ NEFilterManager) SetDisableEncryptedDNSSettings(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisableEncryptedDNSSettings:"), value)
}


// The grade of the filter, which determines when it acts relative to other filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/grade-swift.property
func (n_ NEFilterManager) Grade() NEFilterManagerGrade {
	rv := objc.Send[NEFilterManagerGrade](n_.ID, objc.Sel("grade"))
	return rv
}


// The grade of the filter, which determines when it acts relative to other filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/grade-swift.property
func (n_ NEFilterManager) SetGrade(value NEFilterManagerGrade) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGrade:"), value)
}


// A Boolean used to toggle the enabled state of the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/isEnabled
func (n_ NEFilterManager) Enabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean used to toggle the enabled state of the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/isEnabled
func (n_ NEFilterManager) SetEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnabled:"), value)
}


// A string containing a description of the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/localizedDescription
func (n_ NEFilterManager) LocalizedDescription() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string containing a description of the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/localizedDescription
func (n_ NEFilterManager) SetLocalizedDescription(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// A object containing the filter configuration settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/providerConfiguration
func (n_ NEFilterManager) ProviderConfiguration() INEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}


// A object containing the filter configuration settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/providerConfiguration
func (n_ NEFilterManager) SetProviderConfiguration(value INEFilterProviderConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}


// The domain for errors resulting from calls to the filter manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltererrordomain
func (n_ NEFilterManager) NEFilterErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterErrorDomain"))
	return rv
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








