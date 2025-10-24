// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NIDLTDOAConfiguration */


/* debug [class_header]: Header for NIDLTDOAConfiguration */
// The class instance for the [NIDLTDOAConfiguration] class.
var (
	NIDLTDOAConfigurationClass     _NIDLTDOAConfigurationClass
	NIDLTDOAConfigurationClassOnce sync.Once
)

func getNIDLTDOAConfigurationClass() _NIDLTDOAConfigurationClass {
	NIDLTDOAConfigurationClassOnce.Do(func() {
		NIDLTDOAConfigurationClass = _NIDLTDOAConfigurationClass{objc.GetClass("NIDLTDOAConfiguration")}
	})
	return NIDLTDOAConfigurationClass
}

type _NIDLTDOAConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NIDLTDOAConfiguration */
// An interface definition for the [NIDLTDOAConfiguration] class.
type INIDLTDOAConfiguration interface {
	INIConfiguration
	
/* debug [class_interface_properties]: Properties for NIDLTDOAConfiguration */
	// properties:
	SupportsDLTDOAMeasurement() bool
	SetSupportsDLTDOAMeasurement(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NIDLTDOAConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NIDLTDOAConfiguration */
// Alloc allocates a new instance without initialization.
func (nc _NIDLTDOAConfigurationClass) Alloc() NIDLTDOAConfiguration {
	rv := objc.Send[NIDLTDOAConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NIDLTDOAConfigurationClass) New() NIDLTDOAConfiguration {
	rv := objc.Send[NIDLTDOAConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NIDLTDOAConfiguration) Init() NIDLTDOAConfiguration {
	rv := objc.Send[NIDLTDOAConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NIDLTDOAConfiguration) Autorelease() NIDLTDOAConfiguration {
	rv := objc.Send[NIDLTDOAConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNIDLTDOAConfiguration creates a new NIDLTDOAConfiguration instance.
func NewNIDLTDOAConfiguration() NIDLTDOAConfiguration {
	return getNIDLTDOAConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NIDLTDOAConfiguration */
// A configuration that enables Downlink Time-Difference-of-Arrival ranging.
//
// Run an instance of this configuration to participate in a session that supports the Downlink Time-Difference-of-Arrival (DL-TDoA) feature. Before creating an instance of this class, call first to ensure device support. DL-TDoA is an Ultra Wideband (UWB) ranging strategy that can produce sub-meter (0.5 - 1 meter) location support for tracked devices in a well-defined area. The solution works by installing base stations, or , within the tracked area. The anchors send messages to receiver devices that support DL-TDoA, such as iPhone 12 and later, and the receivers use the messages to calculate their location.


// A configuration that enables Downlink Time-Difference-of-Arrival ranging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAConfiguration
type NIDLTDOAConfiguration struct {
	NIConfiguration
}

// NIDLTDOAConfigurationFrom constructs a [NIDLTDOAConfiguration] from an unsafe.Pointer.
//
// A configuration that enables Downlink Time-Difference-of-Arrival ranging.
func NIDLTDOAConfigurationFrom(ptr unsafe.Pointer) NIDLTDOAConfiguration {
	return NIDLTDOAConfiguration{
		NIConfiguration: NIConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NIDLTDOAConfiguration */

// Initializes a Downlink Time-Difference-of-Arrival (DL-TDoA) configuration for a specific tracked area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAConfiguration/init(networkIdentifier:)
func NewNIDLTDOAConfigurationWithNetworkIdentifier(networkIdentifier int) NIDLTDOAConfiguration {
	instance := getNIDLTDOAConfigurationClass().Alloc()
	rv := objc.Send[NIDLTDOAConfiguration](instance.ID, objc.Sel("initWithNetworkIdentifier:"), networkIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNIDLTDOAConfigurationWithNetworkIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NIDLTDOAConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NIDLTDOAConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NIDLTDOAConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NIDLTDOAConfiguration */

// A property that indicates if the device supports Downlink Time-Difference-of-Arrival ranging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidevicecapability/supportsdltdoameasurement
func (n_ NIDLTDOAConfiguration) SupportsDLTDOAMeasurement() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("supportsDLTDOAMeasurement"))
	return rv
}/* debug [instance_properties/getter]: supportsDLTDOAMeasurement */


// A property that indicates if the device supports Downlink Time-Difference-of-Arrival ranging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nidevicecapability/supportsdltdoameasurement
func (n_ NIDLTDOAConfiguration) SetSupportsDLTDOAMeasurement(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSupportsDLTDOAMeasurement:"), value)
}/* debug [instance_properties/setter]: supportsDLTDOAMeasurement */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NIDLTDOAConfiguration */


