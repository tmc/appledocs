// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CWConfiguration */


/* debug [class_header]: Header for CWConfiguration */
// The class instance for the [CWConfiguration] class.
var (
	CWConfigurationClass     _CWConfigurationClass
	CWConfigurationClassOnce sync.Once
)

func getCWConfigurationClass() _CWConfigurationClass {
	CWConfigurationClassOnce.Do(func() {
		CWConfigurationClass = _CWConfigurationClass{objc.GetClass("CWConfiguration")}
	})
	return CWConfigurationClass
}

type _CWConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CWConfiguration */
// An interface definition for the [CWConfiguration] class.
type ICWConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CWConfiguration */
	// properties:
	NetworkProfiles() unsafe.Pointer
	RememberJoinedNetworks() bool
	RequireAdministratorForAssociation() bool
	RequireAdministratorForIBSSMode() bool
	RequireAdministratorForPower() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CWConfiguration */
	// methods:
	IsEqualToConfiguration(configuration ICWConfiguration) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CWConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _CWConfigurationClass) Alloc() CWConfiguration {
	rv := objc.Send[CWConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CWConfigurationClass) New() CWConfiguration {
	rv := objc.Send[CWConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWConfiguration) Init() CWConfiguration {
	rv := objc.Send[CWConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWConfiguration) Autorelease() CWConfiguration {
	rv := objc.Send[CWConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWConfiguration creates a new CWConfiguration instance.
func NewCWConfiguration() CWConfiguration {
	return getCWConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CWConfiguration */
// Encapsulates an immutable configuration for an AirPort WLAN interface.


// Encapsulates an immutable configuration for an AirPort WLAN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration
type CWConfiguration struct {
	objectivec.Object
}

// CWConfigurationFrom constructs a [CWConfiguration] from an unsafe.Pointer.
//
// Encapsulates an immutable configuration for an AirPort WLAN interface.
func CWConfigurationFrom(ptr unsafe.Pointer) CWConfiguration {
	return CWConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CWConfiguration */

// Creates and returns a CWConfiguration object initialized with the given CWConfiguration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/init(configuration:)
func NewCWConfigurationWithConfiguration(configuration ICWConfiguration) CWConfiguration {
	instance := getCWConfigurationClass().Alloc()
	rv := objc.Send[CWConfiguration](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCWConfigurationWithConfiguration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CWConfiguration */

// Convenience method for getting an empty CWConfiguration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/configuration
func (cc _CWConfigurationClass) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("configuration"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Configuration) */


// Convenience method for getting a CWConfiguration object initialized with the given CWConfiguration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/configurationWithConfiguration:
func (cc _CWConfigurationClass) ConfigurationWithConfiguration(configuration ICWConfiguration) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("configurationWithConfiguration:"), configuration)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithConfiguration) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CWConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CWConfiguration */

// Determine CWConfiguration object equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/isEqual(to:)
func (c_ CWConfiguration) IsEqualToConfiguration(configuration ICWConfiguration) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToConfiguration:"), configuration)
	return rv
}/* debug [instance_methods/method]: IsEqualToConfiguration */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CWConfiguration */

// An array of remembered CWNetworkProfile objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/networkProfiles
func (c_ CWConfiguration) NetworkProfiles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("networkProfiles"))
	return rv
}/* debug [instance_properties/getter]: networkProfiles */


// AirPort client will remember all joined networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/rememberJoinedNetworks
func (c_ CWConfiguration) RememberJoinedNetworks() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rememberJoinedNetworks"))
	return rv
}/* debug [instance_properties/getter]: rememberJoinedNetworks */


// Require an administrator password to change networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForAssociation
func (c_ CWConfiguration) RequireAdministratorForAssociation() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForAssociation"))
	return rv
}/* debug [instance_properties/getter]: requireAdministratorForAssociation */


// Require an administrator password to create a computer-to-computer network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForIBSSMode
func (c_ CWConfiguration) RequireAdministratorForIBSSMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForIBSSMode"))
	return rv
}/* debug [instance_properties/getter]: requireAdministratorForIBSSMode */


// Require an administrator password to change the interface power state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForPower
func (c_ CWConfiguration) RequireAdministratorForPower() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForPower"))
	return rv
}/* debug [instance_properties/getter]: requireAdministratorForPower */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CWConfiguration */


