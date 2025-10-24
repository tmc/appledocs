// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterThermostatUserInterfaceConfiguration] class.
var (
	MTRBaseClusterThermostatUserInterfaceConfigurationClass     _MTRBaseClusterThermostatUserInterfaceConfigurationClass
	MTRBaseClusterThermostatUserInterfaceConfigurationClassOnce sync.Once
)

func getMTRBaseClusterThermostatUserInterfaceConfigurationClass() _MTRBaseClusterThermostatUserInterfaceConfigurationClass {
	MTRBaseClusterThermostatUserInterfaceConfigurationClassOnce.Do(func() {
		MTRBaseClusterThermostatUserInterfaceConfigurationClass = _MTRBaseClusterThermostatUserInterfaceConfigurationClass{objc.GetClass("MTRBaseClusterThermostatUserInterfaceConfiguration")}
	})
	return MTRBaseClusterThermostatUserInterfaceConfigurationClass
}

type _MTRBaseClusterThermostatUserInterfaceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterThermostatUserInterfaceConfiguration] class.
type IMTRBaseClusterThermostatUserInterfaceConfiguration interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostatUserInterfaceConfiguration
type MTRBaseClusterThermostatUserInterfaceConfiguration struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterThermostatUserInterfaceConfigurationFrom constructs a [MTRBaseClusterThermostatUserInterfaceConfiguration] from an unsafe.Pointer.
func MTRBaseClusterThermostatUserInterfaceConfigurationFrom(ptr unsafe.Pointer) MTRBaseClusterThermostatUserInterfaceConfiguration {
	return MTRBaseClusterThermostatUserInterfaceConfiguration{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterThermostatUserInterfaceConfigurationClass) Alloc() MTRBaseClusterThermostatUserInterfaceConfiguration {
	rv := objc.Send[MTRBaseClusterThermostatUserInterfaceConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterThermostatUserInterfaceConfigurationClass) New() MTRBaseClusterThermostatUserInterfaceConfiguration {
	rv := objc.Send[MTRBaseClusterThermostatUserInterfaceConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterThermostatUserInterfaceConfiguration) Init() MTRBaseClusterThermostatUserInterfaceConfiguration {
	rv := objc.Send[MTRBaseClusterThermostatUserInterfaceConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterThermostatUserInterfaceConfiguration) Autorelease() MTRBaseClusterThermostatUserInterfaceConfiguration {
	rv := objc.Send[MTRBaseClusterThermostatUserInterfaceConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterThermostatUserInterfaceConfiguration creates a new MTRBaseClusterThermostatUserInterfaceConfiguration instance.
func NewMTRBaseClusterThermostatUserInterfaceConfiguration() MTRBaseClusterThermostatUserInterfaceConfiguration {
	return getMTRBaseClusterThermostatUserInterfaceConfigurationClass().New()
}




