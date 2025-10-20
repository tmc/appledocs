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
// In Swift, this type bridges to ; use when you need reference semantics or other Foundation-specific behavior. Time zones represent the standard time policies for a geopolitical region. Time zones have identifiers like “America/Los_Angeles” and can also be identified by abbreviations, such as PST for Pacific Standard Time. You can create time zone objects by ID with and by abbreviation with . Time zones can also represent a temporal offset—either plus or minus—from Greenwich Mean Time (GMT). For example, the temporal offset of Pacific Standard Time is 8 hours behind Greenwich Mean Time (GMT-8). You can create time zone objects with a temporal offset by using . You typically work with system time zones rather than creating time zones by identifier or by offset. The class property returns the time zone currently used by the system, if known. This value is cached once the property is accessed and doesn’t reflect any system time zone changes until you call the method. The class property returns an autoupdating proxy object that always returns the current time zone used by the system. You can also set the class property to make your app run as if it were in a different time zone than the system. is with its Core Foundation counterpart, . See for more information on toll-free bridging.
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




