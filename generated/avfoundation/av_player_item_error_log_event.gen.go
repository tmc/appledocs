// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerItemErrorLogEvent] class.
var (
	PlayerItemErrorLogEventClass     _PlayerItemErrorLogEventClass
	PlayerItemErrorLogEventClassOnce sync.Once
)

func getPlayerItemErrorLogEventClass() _PlayerItemErrorLogEventClass {
	PlayerItemErrorLogEventClassOnce.Do(func() {
		PlayerItemErrorLogEventClass = _PlayerItemErrorLogEventClass{objc.GetClass("AVPlayerItemErrorLogEvent")}
	})
	return PlayerItemErrorLogEventClass
}

type _PlayerItemErrorLogEventClass struct {
	class objc.Class
}





// An interface definition for the [PlayerItemErrorLogEvent] class.
type IPlayerItemErrorLogEvent interface {
	objectivec.IObject
	

	// properties:
	AllHTTPResponseHeaderFields() foundation.IDictionary
	Date() foundation.foundation.INSDate
	ErrorComment() foundation.foundation.INSString
	ErrorDomain() foundation.foundation.INSString
	ErrorStatusCode() int
	PlaybackSessionID() foundation.foundation.INSString
	ServerAddress() foundation.foundation.INSString
	URI() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemErrorLogEventClass) Alloc() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemErrorLogEventClass) New() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemErrorLogEvent) Init() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemErrorLogEvent) Autorelease() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemErrorLogEvent creates a new PlayerItemErrorLogEvent instance.
func NewPlayerItemErrorLogEvent() PlayerItemErrorLogEvent {
	return getPlayerItemErrorLogEventClass().New()
}





// A single item in a player item’s error log.
//
// This object provides properties for accessing the data fields of each log event. Each event is a single entry in an object’s error log. These properties aren’t observable. For more information about key-value observing, see .


// A single item in a player item’s error log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent
type PlayerItemErrorLogEvent struct {
	objectivec.Object
}

// PlayerItemErrorLogEventFrom constructs a [PlayerItemErrorLogEvent] from an unsafe.Pointer.
//
// A single item in a player item’s error log.
func PlayerItemErrorLogEventFrom(ptr unsafe.Pointer) PlayerItemErrorLogEvent {
	return PlayerItemErrorLogEvent{objectivec.Object{objc.ID(ptr)}}
}

























// The HTTP header fields the server returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/allHTTPResponseHeaderFields
func (p_ PlayerItemErrorLogEvent) AllHTTPResponseHeaderFields() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("allHTTPResponseHeaderFields"))
	return rv
}


// The date and time when the error occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/date
func (p_ PlayerItemErrorLogEvent) Date() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("date"))
	return rv
}


// A description of the error encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorComment
func (p_ PlayerItemErrorLogEvent) ErrorComment() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("errorComment"))
	return rv
}


// The domain of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorDomain
func (p_ PlayerItemErrorLogEvent) ErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("errorDomain"))
	return rv
}


// A unique error code identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorStatusCode
func (p_ PlayerItemErrorLogEvent) ErrorStatusCode() int {
	rv := objc.Send[int](p_.ID, objc.Sel("errorStatusCode"))
	return rv
}


// A GUID that identifies the playback session that had an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/playbackSessionID
func (p_ PlayerItemErrorLogEvent) PlaybackSessionID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playbackSessionID"))
	return rv
}


// The IP address of the server that was the source of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/serverAddress
func (p_ PlayerItemErrorLogEvent) ServerAddress() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("serverAddress"))
	return rv
}


// The URI of the playback item that had an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/uri
func (p_ PlayerItemErrorLogEvent) URI() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("URI"))
	return rv
}








