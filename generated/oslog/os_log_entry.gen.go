// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OSLogEntry] class.
var (
	OSLogEntryClass     _OSLogEntryClass
	OSLogEntryClassOnce sync.Once
)

func getOSLogEntryClass() _OSLogEntryClass {
	OSLogEntryClassOnce.Do(func() {
		OSLogEntryClass = _OSLogEntryClass{objc.GetClass("OSLogEntry")}
	})
	return OSLogEntryClass
}

type _OSLogEntryClass struct {
	class objc.Class
}

// An interface definition for the [OSLogEntry] class.
type IOSLogEntry interface {
	objectivec.IObject
}

// A single entry from the unified logging system.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry
type OSLogEntry struct {
	objectivec.Object
}

// OSLogEntryFrom constructs a [OSLogEntry] from an unsafe.Pointer.
//
// A single entry from the unified logging system.
func OSLogEntryFrom(ptr unsafe.Pointer) OSLogEntry {
	return OSLogEntry{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OSLogEntryClass) Alloc() OSLogEntry {
	rv := objc.Send[OSLogEntry](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSLogEntryClass) New() OSLogEntry {
	rv := objc.Send[OSLogEntry](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogEntry) Init() OSLogEntry {
	rv := objc.Send[OSLogEntry](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogEntry) Autorelease() OSLogEntry {
	rv := objc.Send[OSLogEntry](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntry creates a new OSLogEntry instance.
func NewOSLogEntry() OSLogEntry {
	return getOSLogEntryClass().New()
}


// The fully formatted message for the entry.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/composedMessage
func (o_ OSLogEntry) ComposedMessage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("composedMessage"))
	return rv
}

// The timestamp of the entry.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/date
func (o_ OSLogEntry) Date() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("date"))
	return rv
}

// The current log entry’s storage tag.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntry/storeCategory-swift.property
func (o_ OSLogEntry) StoreCategory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("storeCategory"))
	return rv
}



