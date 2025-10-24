// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterThermostatUserInterfaceConfiguration] class.
var (
	MTRClusterThermostatUserInterfaceConfigurationClass     _MTRClusterThermostatUserInterfaceConfigurationClass
	MTRClusterThermostatUserInterfaceConfigurationClassOnce sync.Once
)

func getMTRClusterThermostatUserInterfaceConfigurationClass() _MTRClusterThermostatUserInterfaceConfigurationClass {
	MTRClusterThermostatUserInterfaceConfigurationClassOnce.Do(func() {
		MTRClusterThermostatUserInterfaceConfigurationClass = _MTRClusterThermostatUserInterfaceConfigurationClass{objc.GetClass("MTRClusterThermostatUserInterfaceConfiguration")}
	})
	return MTRClusterThermostatUserInterfaceConfigurationClass
}

type _MTRClusterThermostatUserInterfaceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterThermostatUserInterfaceConfiguration] class.
type IMTRClusterThermostatUserInterfaceConfiguration interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThermostatUserInterfaceConfiguration
type MTRClusterThermostatUserInterfaceConfiguration struct {
	MTRGenericCluster
}

// MTRClusterThermostatUserInterfaceConfigurationFrom constructs a [MTRClusterThermostatUserInterfaceConfiguration] from an unsafe.Pointer.
func MTRClusterThermostatUserInterfaceConfigurationFrom(ptr unsafe.Pointer) MTRClusterThermostatUserInterfaceConfiguration {
	return MTRClusterThermostatUserInterfaceConfiguration{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterThermostatUserInterfaceConfigurationClass) Alloc() MTRClusterThermostatUserInterfaceConfiguration {
	rv := objc.Send[MTRClusterThermostatUserInterfaceConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterThermostatUserInterfaceConfigurationClass) New() MTRClusterThermostatUserInterfaceConfiguration {
	rv := objc.Send[MTRClusterThermostatUserInterfaceConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterThermostatUserInterfaceConfiguration) Init() MTRClusterThermostatUserInterfaceConfiguration {
	rv := objc.Send[MTRClusterThermostatUserInterfaceConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterThermostatUserInterfaceConfiguration) Autorelease() MTRClusterThermostatUserInterfaceConfiguration {
	rv := objc.Send[MTRClusterThermostatUserInterfaceConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterThermostatUserInterfaceConfiguration creates a new MTRClusterThermostatUserInterfaceConfiguration instance.
func NewMTRClusterThermostatUserInterfaceConfiguration() MTRClusterThermostatUserInterfaceConfiguration {
	return getMTRClusterThermostatUserInterfaceConfigurationClass().New()
}
