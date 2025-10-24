// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerItemErrorLog] class.
var (
	PlayerItemErrorLogClass     _PlayerItemErrorLogClass
	PlayerItemErrorLogClassOnce sync.Once
)

func getPlayerItemErrorLogClass() _PlayerItemErrorLogClass {
	PlayerItemErrorLogClassOnce.Do(func() {
		PlayerItemErrorLogClass = _PlayerItemErrorLogClass{objc.GetClass("AVPlayerItemErrorLog")}
	})
	return PlayerItemErrorLogClass
}

type _PlayerItemErrorLogClass struct {
	class objc.Class
}





// An interface definition for the [PlayerItemErrorLog] class.
type IPlayerItemErrorLog interface {
	objectivec.IObject
	

	// properties:
	Events() []PlayerItemErrorLogEvent
	ExtendedLogDataStringEncoding() StringEncoding /* not a class type */


	

	// methods:
	ExtendedLogData() foundation.Data


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemErrorLogClass) Alloc() PlayerItemErrorLog {
	rv := objc.Send[PlayerItemErrorLog](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemErrorLogClass) New() PlayerItemErrorLog {
	rv := objc.Send[PlayerItemErrorLog](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemErrorLog) Init() PlayerItemErrorLog {
	rv := objc.Send[PlayerItemErrorLog](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemErrorLog) Autorelease() PlayerItemErrorLog {
	rv := objc.Send[PlayerItemErrorLog](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemErrorLog creates a new PlayerItemErrorLog instance.
func NewPlayerItemErrorLog() PlayerItemErrorLog {
	return getPlayerItemErrorLogClass().New()
}





// The error log associated with a player item.


// The error log associated with a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog
type PlayerItemErrorLog struct {
	objectivec.Object
}

// PlayerItemErrorLogFrom constructs a [PlayerItemErrorLog] from an unsafe.Pointer.
//
// The error log associated with a player item.
func PlayerItemErrorLogFrom(ptr unsafe.Pointer) PlayerItemErrorLog {
	return PlayerItemErrorLog{objectivec.Object{objc.ID(ptr)}}
}




















// Returns a serialized representation of the error log in the Extended Log File Format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog/extendedLogData()
func (p_ PlayerItemErrorLog) ExtendedLogData() foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("extendedLogData"))
	return rv
}







// A chronologically ordered array of player item error log event objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog/events
func (p_ PlayerItemErrorLog) Events() []PlayerItemErrorLogEvent {
	rv := objc.Send[[]PlayerItemErrorLogEvent](p_.ID, objc.Sel("events"))
	return rv
}


// The string encoding of the extended log data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog/extendedLogDataStringEncoding
func (p_ PlayerItemErrorLog) ExtendedLogDataStringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](p_.ID, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}








