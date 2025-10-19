// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OSLogEntryLog] class.
var (
	oSLogEntryLogClass     _OSLogEntryLogClass
	oSLogEntryLogClassOnce sync.Once
)

func getOSLogEntryLogClass() _OSLogEntryLogClass {
	oSLogEntryLogClassOnce.Do(func() {
		oSLogEntryLogClass = _OSLogEntryLogClass{objc.GetClass("OSLogEntryLog")}
	})
	return oSLogEntryLogClass
}

type _OSLogEntryLogClass struct {
	class objc.Class
}

// An interface definition for the [OSLogEntryLog] class.
type IOSLogEntryLog interface {
	IOSLogEntry
}

// A log entry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryLog
type OSLogEntryLog struct {
	OSLogEntry
}

// OSLogEntryLogFrom constructs a [OSLogEntryLog] from an unsafe.Pointer.
//
// A log entry.
func OSLogEntryLogFrom(ptr unsafe.Pointer) OSLogEntryLog {
	return OSLogEntryLog{
		OSLogEntry: OSLogEntryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _OSLogEntryLogClass) Alloc() OSLogEntryLog {
	rv := objc.Send[OSLogEntryLog](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSLogEntryLogClass) New() OSLogEntryLog {
	rv := objc.Send[OSLogEntryLog](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogEntryLog) Init() OSLogEntryLog {
	rv := objc.Send[OSLogEntryLog](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogEntryLog) Autorelease() OSLogEntryLog {
	rv := objc.Send[OSLogEntryLog](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntryLog creates a new OSLogEntryLog instance.
func NewOSLogEntryLog() OSLogEntryLog {
	return getOSLogEntryLogClass().New()
}




