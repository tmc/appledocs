// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MonitoringRecord] class.
var (
	MonitoringRecordClass     _MonitoringRecordClass
	MonitoringRecordClassOnce sync.Once
)

func getMonitoringRecordClass() _MonitoringRecordClass {
	MonitoringRecordClassOnce.Do(func() {
		MonitoringRecordClass = _MonitoringRecordClass{objc.GetClass("CLMonitoringRecord")}
	})
	return MonitoringRecordClass
}

type _MonitoringRecordClass struct {
	class objc.Class
}

// An interface definition for the [MonitoringRecord] class.
type IMonitoringRecord interface {
	objectivec.IObject
}

// An object that represents a condition and its associated information that a location monitor is monitoring.
//
// When handling a new , the available for the indicated identifier from the contains the prior event. The updates with the new event when the handling is complete.


// An object that represents a condition and its associated information that a location monitor is monitoring.
//
// [Full Topic]
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

// Alloc allocates a new instance without initialization.
func (mc _MonitoringRecordClass) Alloc() MonitoringRecord {
	rv := objc.Send[MonitoringRecord](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MonitoringRecordClass) New() MonitoringRecord {
	rv := objc.Send[MonitoringRecord](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MonitoringRecord) Init() MonitoringRecord {
	rv := objc.Send[MonitoringRecord](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MonitoringRecord) Autorelease() MonitoringRecord {
	rv := objc.Send[MonitoringRecord](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMonitoringRecord creates a new MonitoringRecord instance.
func NewMonitoringRecord() MonitoringRecord {
	return getMonitoringRecordClass().New()
}




