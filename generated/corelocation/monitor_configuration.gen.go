// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MonitorConfiguration] class.
var (
	monitorConfigurationClass     _MonitorConfigurationClass
	monitorConfigurationClassOnce sync.Once
)

func getMonitorConfigurationClass() _MonitorConfigurationClass {
	monitorConfigurationClassOnce.Do(func() {
		monitorConfigurationClass = _MonitorConfigurationClass{objc.GetClass("CLMonitorConfiguration")}
	})
	return monitorConfigurationClass
}

type _MonitorConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MonitorConfiguration] class.
type IMonitorConfiguration interface {
	objectivec.IObject
}

// An object for configuring a location monitor instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration
type MonitorConfiguration struct {
	objectivec.Object
}

// MonitorConfigurationFrom constructs a [MonitorConfiguration] from an unsafe.Pointer.
//
// An object for configuring a location monitor instance.
func MonitorConfigurationFrom(ptr unsafe.Pointer) MonitorConfiguration {
	return MonitorConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MonitorConfigurationClass) Alloc() MonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MonitorConfigurationClass) New() MonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MonitorConfiguration) Init() MonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MonitorConfiguration) Autorelease() MonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMonitorConfiguration creates a new MonitorConfiguration instance.
func NewMonitorConfiguration() MonitorConfiguration {
	return getMonitorConfigurationClass().New()
}




