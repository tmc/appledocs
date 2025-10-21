// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LogItem] class.
var (
	LogItemClass     _LogItemClass
	LogItemClassOnce sync.Once
)

func getLogItemClass() _LogItemClass {
	LogItemClassOnce.Do(func() {
		LogItemClass = _LogItemClass{objc.GetClass("CMLogItem")}
	})
	return LogItemClass
}

type _LogItemClass struct {
	class objc.Class
}

// An interface definition for the [LogItem] class.
type ILogItem interface {
	objectivec.IObject
}

// The base class for all motion-related data objects.
//
// The class defines a read-only property that records the time a motion-event measurement was taken.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMLogItem
type LogItem struct {
	objectivec.Object
}

// LogItemFrom constructs a [LogItem] from an unsafe.Pointer.
//
// The base class for all motion-related data objects.
func LogItemFrom(ptr unsafe.Pointer) LogItem {
	return LogItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LogItemClass) Alloc() LogItem {
	rv := objc.Send[LogItem](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LogItemClass) New() LogItem {
	rv := objc.Send[LogItem](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LogItem) Init() LogItem {
	rv := objc.Send[LogItem](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LogItem) Autorelease() LogItem {
	rv := objc.Send[LogItem](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLogItem creates a new LogItem instance.
func NewLogItem() LogItem {
	return getLogItemClass().New()
}


// The time when the logged item is valid.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMLogItem/timestamp
func (l_ LogItem) Timestamp() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](l_.ID, objc.Sel("timestamp"))
	return rv
}



