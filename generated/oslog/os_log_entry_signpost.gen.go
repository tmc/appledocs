// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OSLogEntrySignpost] class.
var (
	oSLogEntrySignpostClass     _OSLogEntrySignpostClass
	oSLogEntrySignpostClassOnce sync.Once
)

func getOSLogEntrySignpostClass() _OSLogEntrySignpostClass {
	oSLogEntrySignpostClassOnce.Do(func() {
		oSLogEntrySignpostClass = _OSLogEntrySignpostClass{objc.GetClass("OSLogEntrySignpost")}
	})
	return oSLogEntrySignpostClass
}

type _OSLogEntrySignpostClass struct {
	class objc.Class
}

// An interface definition for the [OSLogEntrySignpost] class.
type IOSLogEntrySignpost interface {
	IOSLogEntry
}

// An entry containing a signpost. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost
type OSLogEntrySignpost struct {
	OSLogEntry
}

// OSLogEntrySignpostFrom constructs a [OSLogEntrySignpost] from an unsafe.Pointer.
//
// An entry containing a signpost.
func OSLogEntrySignpostFrom(ptr unsafe.Pointer) OSLogEntrySignpost {
	return OSLogEntrySignpost{
		OSLogEntry: OSLogEntryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _OSLogEntrySignpostClass) Alloc() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSLogEntrySignpostClass) New() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogEntrySignpost) Init() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogEntrySignpost) Autorelease() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntrySignpost creates a new OSLogEntrySignpost instance.
func NewOSLogEntrySignpost() OSLogEntrySignpost {
	return getOSLogEntrySignpostClass().New()
}




