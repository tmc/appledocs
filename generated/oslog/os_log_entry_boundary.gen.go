// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OSLogEntryBoundary] class.
var (
	OSLogEntryBoundaryClass     _OSLogEntryBoundaryClass
	OSLogEntryBoundaryClassOnce sync.Once
)

func getOSLogEntryBoundaryClass() _OSLogEntryBoundaryClass {
	OSLogEntryBoundaryClassOnce.Do(func() {
		OSLogEntryBoundaryClass = _OSLogEntryBoundaryClass{objc.GetClass("OSLogEntryBoundary")}
	})
	return OSLogEntryBoundaryClass
}

type _OSLogEntryBoundaryClass struct {
	class objc.Class
}

// An interface definition for the [OSLogEntryBoundary] class.
type IOSLogEntryBoundary interface {
	IOSLogEntry
}

// The metadata that partitions sequences of other entries.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryBoundary
type OSLogEntryBoundary struct {
	OSLogEntry
}

// OSLogEntryBoundaryFrom constructs a [OSLogEntryBoundary] from an unsafe.Pointer.
//
// The metadata that partitions sequences of other entries.
func OSLogEntryBoundaryFrom(ptr unsafe.Pointer) OSLogEntryBoundary {
	return OSLogEntryBoundary{
		OSLogEntry: OSLogEntryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _OSLogEntryBoundaryClass) Alloc() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSLogEntryBoundaryClass) New() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogEntryBoundary) Init() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogEntryBoundary) Autorelease() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntryBoundary creates a new OSLogEntryBoundary instance.
func NewOSLogEntryBoundary() OSLogEntryBoundary {
	return getOSLogEntryBoundaryClass().New()
}




