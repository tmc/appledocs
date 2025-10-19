// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OSLogPosition] class.
var (
	oSLogPositionClass     _OSLogPositionClass
	oSLogPositionClassOnce sync.Once
)

func getOSLogPositionClass() _OSLogPositionClass {
	oSLogPositionClassOnce.Do(func() {
		oSLogPositionClass = _OSLogPositionClass{objc.GetClass("OSLogPosition")}
	})
	return oSLogPositionClass
}

type _OSLogPositionClass struct {
	class objc.Class
}

// An interface definition for the [OSLogPosition] class.
type IOSLogPosition interface {
	objectivec.IObject
}

// A representation of a point in a sequence of entries in the unified logging system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogPosition
type OSLogPosition struct {
	objectivec.Object
}

// OSLogPositionFrom constructs a [OSLogPosition] from an unsafe.Pointer.
//
// A representation of a point in a sequence of entries in the unified logging system.
func OSLogPositionFrom(ptr unsafe.Pointer) OSLogPosition {
	return OSLogPosition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OSLogPositionClass) Alloc() OSLogPosition {
	rv := objc.Send[OSLogPosition](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSLogPositionClass) New() OSLogPosition {
	rv := objc.Send[OSLogPosition](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogPosition) Init() OSLogPosition {
	rv := objc.Send[OSLogPosition](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogPosition) Autorelease() OSLogPosition {
	rv := objc.Send[OSLogPosition](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogPosition creates a new OSLogPosition instance.
func NewOSLogPosition() OSLogPosition {
	return getOSLogPositionClass().New()
}




