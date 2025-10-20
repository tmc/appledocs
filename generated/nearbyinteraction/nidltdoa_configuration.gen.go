// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [NIDLTDOAConfiguration] class.
type INIDLTDOAConfiguration interface {
	INIConfiguration
}

// A configuration that enables Downlink Time-Difference-of-Arrival ranging.
//
// Run an instance of this configuration to participate in a session that supports the Downlink Time-Difference-of-Arrival (DL-TDoA) feature. Before creating an instance of this class, call first to ensure device support. DL-TDoA is an Ultra Wideband (UWB) ranging strategy that can produce sub-meter (0.5 - 1 meter) location support for tracked devices in a well-defined area. The solution works by installing base stations, or , within the tracked area. The anchors send messages to receiver devices that support DL-TDoA, such as iPhone 12 and later, and the receivers use the messages to calculate their location.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NIDLTDOAConfigurationClass) Alloc() NIDLTDOAConfiguration {
	rv := objc.Send[NIDLTDOAConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




