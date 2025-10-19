// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TimeZone] class.
var (
	timeZoneClass     _TimeZoneClass
	timeZoneClassOnce sync.Once
)

func getTimeZoneClass() _TimeZoneClass {
	timeZoneClassOnce.Do(func() {
		timeZoneClass = _TimeZoneClass{objc.GetClass("NSTimeZone")}
	})
	return timeZoneClass
}

type _TimeZoneClass struct {
	class objc.Class
}

// An interface definition for the [TimeZone] class.
type ITimeZone interface {
	objectivec.IObject
}

// Information about standard time conventions associated with a specific geopolitical region.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone
type TimeZone struct {
	objectivec.Object
}

// TimeZoneFrom constructs a [TimeZone] from an unsafe.Pointer.
//
// Information about standard time conventions associated with a specific geopolitical region.
func TimeZoneFrom(ptr unsafe.Pointer) TimeZone {
	return TimeZone{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TimeZoneClass) Alloc() TimeZone {
	rv := objc.Send[TimeZone](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TimeZoneClass) New() TimeZone {
	rv := objc.Send[TimeZone](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TimeZone) Init() TimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TimeZone) Autorelease() TimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTimeZone creates a new TimeZone instance.
func NewTimeZone() TimeZone {
	return getTimeZoneClass().New()
}




