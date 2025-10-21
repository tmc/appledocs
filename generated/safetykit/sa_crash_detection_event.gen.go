// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SACrashDetectionEvent] class.
var (
	SACrashDetectionEventClass     _SACrashDetectionEventClass
	SACrashDetectionEventClassOnce sync.Once
)

func getSACrashDetectionEventClass() _SACrashDetectionEventClass {
	SACrashDetectionEventClassOnce.Do(func() {
		SACrashDetectionEventClass = _SACrashDetectionEventClass{objc.GetClass("SACrashDetectionEvent")}
	})
	return SACrashDetectionEventClass
}

type _SACrashDetectionEventClass struct {
	class objc.Class
}

// An interface definition for the [SACrashDetectionEvent] class.
type ISACrashDetectionEvent interface {
	objectivec.IObject
}

// Describes the information about a vehicular crash.
//
// When a vehicular crash occurs, SafetyKit calls your delegate’s method with an object. Inspect this object to determine information about the crash, including the date and time, location, and if the system attempted to contact emergency services.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent
type SACrashDetectionEvent struct {
	objectivec.Object
}

// SACrashDetectionEventFrom constructs a [SACrashDetectionEvent] from an unsafe.Pointer.
//
// Describes the information about a vehicular crash.
func SACrashDetectionEventFrom(ptr unsafe.Pointer) SACrashDetectionEvent {
	return SACrashDetectionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SACrashDetectionEventClass) Alloc() SACrashDetectionEvent {
	rv := objc.Send[SACrashDetectionEvent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SACrashDetectionEventClass) New() SACrashDetectionEvent {
	rv := objc.Send[SACrashDetectionEvent](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SACrashDetectionEvent) Init() SACrashDetectionEvent {
	rv := objc.Send[SACrashDetectionEvent](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SACrashDetectionEvent) Autorelease() SACrashDetectionEvent {
	rv := objc.Send[SACrashDetectionEvent](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSACrashDetectionEvent creates a new SACrashDetectionEvent instance.
func NewSACrashDetectionEvent() SACrashDetectionEvent {
	return getSACrashDetectionEventClass().New()
}


// The date and time the crash occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/date
func (s_ SACrashDetectionEvent) Date() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("date"))
	return rv
}

// The longitude and latitude where the crash detection occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/location
func (s_ SACrashDetectionEvent) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("location"))
	return rv
}

// An indication of whether the system attempted to call an Emergency SOS provider.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/response-swift.property
func (s_ SACrashDetectionEvent) Response() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("response"))
	return rv
}



