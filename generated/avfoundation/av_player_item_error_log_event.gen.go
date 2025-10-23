// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AllHTTPResponseHeaderFields() objc.IObject /* cross-framework: NSString */
	SetAllHTTPResponseHeaderFields(value objc.IObject /* cross-framework: NSString */)
	Date() objc.IObject /* cross-framework: Date */
	SetDate(value objc.IObject /* cross-framework: Date */)
	ErrorComment() objc.IObject /* cross-framework: NSString */
	SetErrorComment(value objc.IObject /* cross-framework: NSString */)
	ErrorDomain() objc.IObject /* cross-framework: NSString */
	SetErrorDomain(value objc.IObject /* cross-framework: NSString */)
	ErrorStatusCode() int /* primitive/slice/pointer. */
	SetErrorStatusCode(value int /* primitive/slice/pointer. */)
	PlaybackSessionID() objc.IObject /* cross-framework: NSString */
	SetPlaybackSessionID(value objc.IObject /* cross-framework: NSString */)
	ServerAddress() objc.IObject /* cross-framework: NSString */
	SetServerAddress(value objc.IObject /* cross-framework: NSString */)
	Uri() objc.IObject /* cross-framework: NSString */
	SetUri(value objc.IObject /* cross-framework: NSString */)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemErrorLogEventClass) Alloc() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The HTTP header fields the server returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/allhttpresponseheaderfields
func (p_ PlayerItemErrorLogEvent) AllHTTPResponseHeaderFields() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("allHTTPResponseHeaderFields"))
	return rv
}


// The HTTP header fields the server returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/allhttpresponseheaderfields
func (p_ PlayerItemErrorLogEvent) SetAllHTTPResponseHeaderFields(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllHTTPResponseHeaderFields:"), value)
}


// The date and time when the error occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/date
func (p_ PlayerItemErrorLogEvent) Date() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("date"))
	return rv
}


// The date and time when the error occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/date
func (p_ PlayerItemErrorLogEvent) SetDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDate:"), value)
}


// A description of the error encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/errorcomment
func (p_ PlayerItemErrorLogEvent) ErrorComment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("errorComment"))
	return rv
}


// A description of the error encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/errorcomment
func (p_ PlayerItemErrorLogEvent) SetErrorComment(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setErrorComment:"), value)
}


// The domain of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/errordomain
func (p_ PlayerItemErrorLogEvent) ErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("errorDomain"))
	return rv
}


// The domain of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/errordomain
func (p_ PlayerItemErrorLogEvent) SetErrorDomain(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setErrorDomain:"), value)
}


// A unique error code identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/errorstatuscode
func (p_ PlayerItemErrorLogEvent) ErrorStatusCode() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("errorStatusCode"))
	return rv
}


// A unique error code identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/errorstatuscode
func (p_ PlayerItemErrorLogEvent) SetErrorStatusCode(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setErrorStatusCode:"), value)
}


// A GUID that identifies the playback session that had an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/playbacksessionid
func (p_ PlayerItemErrorLogEvent) PlaybackSessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playbackSessionID"))
	return rv
}


// A GUID that identifies the playback session that had an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/playbacksessionid
func (p_ PlayerItemErrorLogEvent) SetPlaybackSessionID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackSessionID:"), value)
}


// The IP address of the server that was the source of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/serveraddress
func (p_ PlayerItemErrorLogEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("serverAddress"))
	return rv
}


// The IP address of the server that was the source of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/serveraddress
func (p_ PlayerItemErrorLogEvent) SetServerAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setServerAddress:"), value)
}


// The URI of the playback item that had an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/uri
func (p_ PlayerItemErrorLogEvent) Uri() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("uri"))
	return rv
}


// The URI of the playback item that had an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/uri
func (p_ PlayerItemErrorLogEvent) SetUri(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUri:"), value)
}



