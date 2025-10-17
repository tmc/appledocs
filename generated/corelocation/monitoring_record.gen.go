// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MonitoringRecord] class.
var monitoringRecordClass = _MonitoringRecordClass{objc.GetClass("CLMonitoringRecord")}

type _MonitoringRecordClass struct {
	class objc.Class
}

// An object that represents a condition and its associated information that a location monitor is monitoring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringRecord

type MonitoringRecord struct {
	objectivec.Object
}

// MonitoringRecordFrom constructs a [MonitoringRecord] from an unsafe.Pointer.
//
// An object that represents a condition and its associated information that a location monitor is monitoring.
func MonitoringRecordFrom(ptr unsafe.Pointer) MonitoringRecord {
	return MonitoringRecord{objectivec.Object{objc.ID(ptr)}}
}



