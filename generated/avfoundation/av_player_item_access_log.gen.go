// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerItemAccessLog] class.
var (
	PlayerItemAccessLogClass     _PlayerItemAccessLogClass
	PlayerItemAccessLogClassOnce sync.Once
)

func getPlayerItemAccessLogClass() _PlayerItemAccessLogClass {
	PlayerItemAccessLogClassOnce.Do(func() {
		PlayerItemAccessLogClass = _PlayerItemAccessLogClass{objc.GetClass("AVPlayerItemAccessLog")}
	})
	return PlayerItemAccessLogClass
}

type _PlayerItemAccessLogClass struct {
	class objc.Class
}

// An interface definition for the [PlayerItemAccessLog] class.
type IPlayerItemAccessLog interface {
	objectivec.IObject
	// properties:
	Events() IAVPlayerItemAccessLogEvent
	SetEvents(value IAVPlayerItemAccessLogEvent)
	ExtendedLogDataStringEncoding() uint
	SetExtendedLogDataStringEncoding(value uint)
	// methods:
}

// An object used to retrieve the access log associated with a player item.
//
// An object accumulates key metrics about network playback and presents them as a collection of instances. Each event instance collates the data that relates to each uninterrupted period of playback.


// An object used to retrieve the access log associated with a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog
type PlayerItemAccessLog struct {
	objectivec.Object
}

// PlayerItemAccessLogFrom constructs a [PlayerItemAccessLog] from an unsafe.Pointer.
//
// An object used to retrieve the access log associated with a player item.
func PlayerItemAccessLogFrom(ptr unsafe.Pointer) PlayerItemAccessLog {
	return PlayerItemAccessLog{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemAccessLogClass) Alloc() PlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerItemAccessLogClass) New() PlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemAccessLog) Init() PlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemAccessLog) Autorelease() PlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemAccessLog creates a new PlayerItemAccessLog instance.
func NewPlayerItemAccessLog() PlayerItemAccessLog {
	return getPlayerItemAccessLogClass().New()
}



// A chronologically ordered array of player item access log events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslog/events
func (p_ PlayerItemAccessLog) Events() IAVPlayerItemAccessLogEvent {
	rv := objc.Send[PlayerItemAccessLogEvent](p_.ID, objc.Sel("events"))
	return rv
}


// A chronologically ordered array of player item access log events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslog/events
func (p_ PlayerItemAccessLog) SetEvents(value IAVPlayerItemAccessLogEvent) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEvents:"), value)
}


// The string encoding of the extended log data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslog/extendedlogdatastringencoding
func (p_ PlayerItemAccessLog) ExtendedLogDataStringEncoding() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}


// The string encoding of the extended log data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslog/extendedlogdatastringencoding
func (p_ PlayerItemAccessLog) SetExtendedLogDataStringEncoding(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExtendedLogDataStringEncoding:"), value)
}



