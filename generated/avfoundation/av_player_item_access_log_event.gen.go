// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerItemAccessLogEvent] class.
var (
	PlayerItemAccessLogEventClass     _PlayerItemAccessLogEventClass
	PlayerItemAccessLogEventClassOnce sync.Once
)

func getPlayerItemAccessLogEventClass() _PlayerItemAccessLogEventClass {
	PlayerItemAccessLogEventClassOnce.Do(func() {
		PlayerItemAccessLogEventClass = _PlayerItemAccessLogEventClass{objc.GetClass("AVPlayerItemAccessLogEvent")}
	})
	return PlayerItemAccessLogEventClass
}

type _PlayerItemAccessLogEventClass struct {
	class objc.Class
}





// An interface definition for the [PlayerItemAccessLogEvent] class.
type IPlayerItemAccessLogEvent interface {
	objectivec.IObject
	

	// properties:
	AverageAudioBitrate() float64
	AverageVideoBitrate() float64
	DownloadOverdue() int
	DurationWatched() float64
	IndicatedAverageBitrate() float64
	IndicatedBitrate() float64
	MediaRequestsWWAN() int
	NumberOfBytesTransferred() objectivec.IObject
	NumberOfDroppedVideoFrames() int
	NumberOfMediaRequests() int
	NumberOfServerAddressChanges() int
	NumberOfStalls() int
	ObservedBitrate() float64
	ObservedBitrateStandardDeviation() float64
	ObservedMaxBitrate() float64
	ObservedMinBitrate() float64
	PlaybackSessionID() objc.IObject /* cross-framework: NSString */
	PlaybackStartDate() objc.IObject /* cross-framework: NSDate */
	PlaybackStartOffset() float64
	PlaybackType() objc.IObject /* cross-framework: NSString */
	SegmentsDownloadedDuration() float64
	ServerAddress() objc.IObject /* cross-framework: NSString */
	StartupTime() float64
	SwitchBitrate() float64
	TransferDuration() float64
	URI() objc.IObject /* cross-framework: NSString */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemAccessLogEventClass) Alloc() PlayerItemAccessLogEvent {
	rv := objc.Send[PlayerItemAccessLogEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemAccessLogEventClass) New() PlayerItemAccessLogEvent {
	rv := objc.Send[PlayerItemAccessLogEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemAccessLogEvent) Init() PlayerItemAccessLogEvent {
	rv := objc.Send[PlayerItemAccessLogEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemAccessLogEvent) Autorelease() PlayerItemAccessLogEvent {
	rv := objc.Send[PlayerItemAccessLogEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemAccessLogEvent creates a new PlayerItemAccessLogEvent instance.
func NewPlayerItemAccessLogEvent() PlayerItemAccessLogEvent {
	return getPlayerItemAccessLogEventClass().New()
}





// A single entry in a player item’s access log.
//
// This object provides named properties for accessing the data fields of each log event. Each event is a single entry in an object’s access log. These properties aren’t observable. For more information about key-value observing, see .


// A single entry in a player item’s access log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent
type PlayerItemAccessLogEvent struct {
	objectivec.Object
}

// PlayerItemAccessLogEventFrom constructs a [PlayerItemAccessLogEvent] from an unsafe.Pointer.
//
// A single entry in a player item’s access log.
func PlayerItemAccessLogEventFrom(ptr unsafe.Pointer) PlayerItemAccessLogEvent {
	return PlayerItemAccessLogEvent{objectivec.Object{objc.ID(ptr)}}
}

























// The audio track’s average bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/averageAudioBitrate
func (p_ PlayerItemAccessLogEvent) AverageAudioBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("averageAudioBitrate"))
	return rv
}


// The video track’s average bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/averageVideoBitrate
func (p_ PlayerItemAccessLogEvent) AverageVideoBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("averageVideoBitrate"))
	return rv
}


// The total number of times that downloading the segments took too long.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/downloadOverdue
func (p_ PlayerItemAccessLogEvent) DownloadOverdue() int {
	rv := objc.Send[int](p_.ID, objc.Sel("downloadOverdue"))
	return rv
}


// The accumulated duration, in seconds, of the media played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/durationWatched
func (p_ PlayerItemAccessLogEvent) DurationWatched() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("durationWatched"))
	return rv
}


// The average throughput, in bits per second, required to play the stream, as advertised by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/indicatedAverageBitrate
func (p_ PlayerItemAccessLogEvent) IndicatedAverageBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("indicatedAverageBitrate"))
	return rv
}


// The throughput, in bits per second, required to play the stream, as advertised by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/indicatedBitrate
func (p_ PlayerItemAccessLogEvent) IndicatedBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("indicatedBitrate"))
	return rv
}


// The number of network read requests over a WWAN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/mediaRequestsWWAN
func (p_ PlayerItemAccessLogEvent) MediaRequestsWWAN() int {
	rv := objc.Send[int](p_.ID, objc.Sel("mediaRequestsWWAN"))
	return rv
}


// The accumulated number of bytes transferred by the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfBytesTransferred
func (p_ PlayerItemAccessLogEvent) NumberOfBytesTransferred() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("numberOfBytesTransferred"))
	return rv
}


// The total number of dropped video frames
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfDroppedVideoFrames
func (p_ PlayerItemAccessLogEvent) NumberOfDroppedVideoFrames() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}


// The number of media read requests from the server to this client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfMediaRequests
func (p_ PlayerItemAccessLogEvent) NumberOfMediaRequests() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfMediaRequests"))
	return rv
}


// A count of changes to the server address over the last uninterrupted period of playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfServerAddressChanges
func (p_ PlayerItemAccessLogEvent) NumberOfServerAddressChanges() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfServerAddressChanges"))
	return rv
}


// The total number of playback stalls encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfStalls
func (p_ PlayerItemAccessLogEvent) NumberOfStalls() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfStalls"))
	return rv
}


// The empirical throughput, in bits per second, across all media downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/observedBitrate
func (p_ PlayerItemAccessLogEvent) ObservedBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("observedBitrate"))
	return rv
}


// The standard deviation of the observed segment download bit rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/observedBitrateStandardDeviation
func (p_ PlayerItemAccessLogEvent) ObservedBitrateStandardDeviation() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("observedBitrateStandardDeviation"))
	return rv
}


// The maximum observed segment download bit rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/observedMaxBitrate
func (p_ PlayerItemAccessLogEvent) ObservedMaxBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("observedMaxBitrate"))
	return rv
}


// The minimum observed segment download bit rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/observedMinBitrate
func (p_ PlayerItemAccessLogEvent) ObservedMinBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("observedMinBitrate"))
	return rv
}


// A GUID that identifies the playback session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/playbackSessionID
func (p_ PlayerItemAccessLogEvent) PlaybackSessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playbackSessionID"))
	return rv
}


// The date and time at which playback began for this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/playbackStartDate
func (p_ PlayerItemAccessLogEvent) PlaybackStartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("playbackStartDate"))
	return rv
}


// The offset, in seconds, in the playlist where the last uninterrupted period of playback began.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/playbackStartOffset
func (p_ PlayerItemAccessLogEvent) PlaybackStartOffset() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("playbackStartOffset"))
	return rv
}


// The playback type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/playbackType
func (p_ PlayerItemAccessLogEvent) PlaybackType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playbackType"))
	return rv
}


// The accumulated duration, in seconds, of the media segments downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/segmentsDownloadedDuration
func (p_ PlayerItemAccessLogEvent) SegmentsDownloadedDuration() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("segmentsDownloadedDuration"))
	return rv
}


// The IP address of the server that was the source of the last delivered media segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/serverAddress
func (p_ PlayerItemAccessLogEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("serverAddress"))
	return rv
}


// The accumulated duration, in seconds, until the player item is ready to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/startupTime
func (p_ PlayerItemAccessLogEvent) StartupTime() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("startupTime"))
	return rv
}


// The bandwidth value that causes a switch, up or down, in the item’s quality being played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/switchBitrate
func (p_ PlayerItemAccessLogEvent) SwitchBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("switchBitrate"))
	return rv
}


// The accumulated duration, in seconds, of active network transfer of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/transferDuration
func (p_ PlayerItemAccessLogEvent) TransferDuration() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("transferDuration"))
	return rv
}


// The URI of the playback item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/uri
func (p_ PlayerItemAccessLogEvent) URI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("URI"))
	return rv
}







