// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MonitorConfiguration] class.
var monitorConfigurationClass = _MonitorConfigurationClass{objc.GetClass("CLMonitorConfiguration")}

type _MonitorConfigurationClass struct {
	class objc.Class
}

// An object for configuring a location monitor instance. [Full Topic]
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



