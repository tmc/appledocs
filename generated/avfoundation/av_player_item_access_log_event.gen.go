// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	SetAverageAudioBitrate(value float64)
	AverageVideoBitrate() float64
	SetAverageVideoBitrate(value float64)
	DownloadOverdue() int
	SetDownloadOverdue(value int)
	DurationWatched() float64
	SetDurationWatched(value float64)
	IndicatedAverageBitrate() float64
	SetIndicatedAverageBitrate(value float64)
	IndicatedBitrate() float64
	SetIndicatedBitrate(value float64)
	MediaRequestsWWAN() int
	SetMediaRequestsWWAN(value int)
	NumberOfBytesTransferred() unsafe.Pointer
	SetNumberOfBytesTransferred(value unsafe.Pointer)
	NumberOfDroppedVideoFrames() int
	SetNumberOfDroppedVideoFrames(value int)
	NumberOfMediaRequests() int
	SetNumberOfMediaRequests(value int)
	NumberOfSegmentsDownloaded() int
	SetNumberOfSegmentsDownloaded(value int)
	NumberOfServerAddressChanges() int
	SetNumberOfServerAddressChanges(value int)
	NumberOfStalls() int
	SetNumberOfStalls(value int)
	ObservedBitrate() float64
	SetObservedBitrate(value float64)
	ObservedBitrateStandardDeviation() float64
	SetObservedBitrateStandardDeviation(value float64)
	ObservedMaxBitrate() float64
	SetObservedMaxBitrate(value float64)
	ObservedMinBitrate() float64
	SetObservedMinBitrate(value float64)
	PlaybackSessionID() objc.IObject /* cross-framework: NSString */
	SetPlaybackSessionID(value objc.IObject /* cross-framework: NSString */)
	PlaybackStartDate() objc.IObject /* cross-framework: Date */
	SetPlaybackStartDate(value objc.IObject /* cross-framework: Date */)
	PlaybackStartOffset() float64
	SetPlaybackStartOffset(value float64)
	PlaybackType() objc.IObject /* cross-framework: NSString */
	SetPlaybackType(value objc.IObject /* cross-framework: NSString */)
	SegmentsDownloadedDuration() float64
	SetSegmentsDownloadedDuration(value float64)
	ServerAddress() objc.IObject /* cross-framework: NSString */
	SetServerAddress(value objc.IObject /* cross-framework: NSString */)
	StartupTime() float64
	SetStartupTime(value float64)
	SwitchBitrate() float64
	SetSwitchBitrate(value float64)
	TransferDuration() float64
	SetTransferDuration(value float64)
	Uri() objc.IObject /* cross-framework: NSString */
	SetUri(value objc.IObject /* cross-framework: NSString */)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemAccessLogEventClass) Alloc() PlayerItemAccessLogEvent {
	rv := objc.Send[PlayerItemAccessLogEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The audio track’s average bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/averageaudiobitrate
func (p_ PlayerItemAccessLogEvent) AverageAudioBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("averageAudioBitrate"))
	return rv
}


// The audio track’s average bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/averageaudiobitrate
func (p_ PlayerItemAccessLogEvent) SetAverageAudioBitrate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAverageAudioBitrate:"), value)
}


// The video track’s average bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/averagevideobitrate
func (p_ PlayerItemAccessLogEvent) AverageVideoBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("averageVideoBitrate"))
	return rv
}


// The video track’s average bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/averagevideobitrate
func (p_ PlayerItemAccessLogEvent) SetAverageVideoBitrate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAverageVideoBitrate:"), value)
}


// The total number of times that downloading the segments took too long.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/downloadoverdue
func (p_ PlayerItemAccessLogEvent) DownloadOverdue() int {
	rv := objc.Send[int](p_.ID, objc.Sel("downloadOverdue"))
	return rv
}


// The total number of times that downloading the segments took too long.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/downloadoverdue
func (p_ PlayerItemAccessLogEvent) SetDownloadOverdue(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDownloadOverdue:"), value)
}


// The accumulated duration, in seconds, of the media played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/durationwatched
func (p_ PlayerItemAccessLogEvent) DurationWatched() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("durationWatched"))
	return rv
}


// The accumulated duration, in seconds, of the media played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/durationwatched
func (p_ PlayerItemAccessLogEvent) SetDurationWatched(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDurationWatched:"), value)
}


// The average throughput, in bits per second, required to play the stream, as advertised by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/indicatedaveragebitrate
func (p_ PlayerItemAccessLogEvent) IndicatedAverageBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("indicatedAverageBitrate"))
	return rv
}


// The average throughput, in bits per second, required to play the stream, as advertised by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/indicatedaveragebitrate
func (p_ PlayerItemAccessLogEvent) SetIndicatedAverageBitrate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndicatedAverageBitrate:"), value)
}


// The throughput, in bits per second, required to play the stream, as advertised by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/indicatedbitrate
func (p_ PlayerItemAccessLogEvent) IndicatedBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("indicatedBitrate"))
	return rv
}


// The throughput, in bits per second, required to play the stream, as advertised by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/indicatedbitrate
func (p_ PlayerItemAccessLogEvent) SetIndicatedBitrate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndicatedBitrate:"), value)
}


// The number of network read requests over a WWAN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/mediarequestswwan
func (p_ PlayerItemAccessLogEvent) MediaRequestsWWAN() int {
	rv := objc.Send[int](p_.ID, objc.Sel("mediaRequestsWWAN"))
	return rv
}


// The number of network read requests over a WWAN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/mediarequestswwan
func (p_ PlayerItemAccessLogEvent) SetMediaRequestsWWAN(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaRequestsWWAN:"), value)
}


// The accumulated number of bytes transferred by the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofbytestransferred
func (p_ PlayerItemAccessLogEvent) NumberOfBytesTransferred() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("numberOfBytesTransferred"))
	return rv
}


// The accumulated number of bytes transferred by the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofbytestransferred
func (p_ PlayerItemAccessLogEvent) SetNumberOfBytesTransferred(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfBytesTransferred:"), value)
}


// The total number of dropped video frames
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofdroppedvideoframes
func (p_ PlayerItemAccessLogEvent) NumberOfDroppedVideoFrames() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}


// The total number of dropped video frames
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofdroppedvideoframes
func (p_ PlayerItemAccessLogEvent) SetNumberOfDroppedVideoFrames(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfDroppedVideoFrames:"), value)
}


// The number of media read requests from the server to this client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofmediarequests
func (p_ PlayerItemAccessLogEvent) NumberOfMediaRequests() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfMediaRequests"))
	return rv
}


// The number of media read requests from the server to this client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofmediarequests
func (p_ PlayerItemAccessLogEvent) SetNumberOfMediaRequests(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfMediaRequests:"), value)
}


// A count of the media segments downloaded from the server to this client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofsegmentsdownloaded
func (p_ PlayerItemAccessLogEvent) NumberOfSegmentsDownloaded() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfSegmentsDownloaded"))
	return rv
}


// A count of the media segments downloaded from the server to this client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofsegmentsdownloaded
func (p_ PlayerItemAccessLogEvent) SetNumberOfSegmentsDownloaded(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfSegmentsDownloaded:"), value)
}


// A count of changes to the server address over the last uninterrupted period of playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofserveraddresschanges
func (p_ PlayerItemAccessLogEvent) NumberOfServerAddressChanges() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfServerAddressChanges"))
	return rv
}


// A count of changes to the server address over the last uninterrupted period of playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofserveraddresschanges
func (p_ PlayerItemAccessLogEvent) SetNumberOfServerAddressChanges(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfServerAddressChanges:"), value)
}


// The total number of playback stalls encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofstalls
func (p_ PlayerItemAccessLogEvent) NumberOfStalls() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfStalls"))
	return rv
}


// The total number of playback stalls encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/numberofstalls
func (p_ PlayerItemAccessLogEvent) SetNumberOfStalls(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfStalls:"), value)
}


// The empirical throughput, in bits per second, across all media downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/observedbitrate
func (p_ PlayerItemAccessLogEvent) ObservedBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("observedBitrate"))
	return rv
}


// The empirical throughput, in bits per second, across all media downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/observedbitrate
func (p_ PlayerItemAccessLogEvent) SetObservedBitrate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObservedBitrate:"), value)
}


// The standard deviation of the observed segment download bit rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/observedbitratestandarddeviation
func (p_ PlayerItemAccessLogEvent) ObservedBitrateStandardDeviation() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("observedBitrateStandardDeviation"))
	return rv
}


// The standard deviation of the observed segment download bit rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/observedbitratestandarddeviation
func (p_ PlayerItemAccessLogEvent) SetObservedBitrateStandardDeviation(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObservedBitrateStandardDeviation:"), value)
}


// The maximum observed segment download bit rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/observedmaxbitrate
func (p_ PlayerItemAccessLogEvent) ObservedMaxBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("observedMaxBitrate"))
	return rv
}


// The maximum observed segment download bit rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/observedmaxbitrate
func (p_ PlayerItemAccessLogEvent) SetObservedMaxBitrate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObservedMaxBitrate:"), value)
}


// The minimum observed segment download bit rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/observedminbitrate
func (p_ PlayerItemAccessLogEvent) ObservedMinBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("observedMinBitrate"))
	return rv
}


// The minimum observed segment download bit rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/observedminbitrate
func (p_ PlayerItemAccessLogEvent) SetObservedMinBitrate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObservedMinBitrate:"), value)
}


// A GUID that identifies the playback session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/playbacksessionid
func (p_ PlayerItemAccessLogEvent) PlaybackSessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playbackSessionID"))
	return rv
}


// A GUID that identifies the playback session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/playbacksessionid
func (p_ PlayerItemAccessLogEvent) SetPlaybackSessionID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackSessionID:"), value)
}


// The date and time at which playback began for this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/playbackstartdate
func (p_ PlayerItemAccessLogEvent) PlaybackStartDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("playbackStartDate"))
	return rv
}


// The date and time at which playback began for this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/playbackstartdate
func (p_ PlayerItemAccessLogEvent) SetPlaybackStartDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackStartDate:"), value)
}


// The offset, in seconds, in the playlist where the last uninterrupted period of playback began.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/playbackstartoffset
func (p_ PlayerItemAccessLogEvent) PlaybackStartOffset() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("playbackStartOffset"))
	return rv
}


// The offset, in seconds, in the playlist where the last uninterrupted period of playback began.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/playbackstartoffset
func (p_ PlayerItemAccessLogEvent) SetPlaybackStartOffset(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackStartOffset:"), value)
}


// The playback type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/playbacktype
func (p_ PlayerItemAccessLogEvent) PlaybackType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playbackType"))
	return rv
}


// The playback type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/playbacktype
func (p_ PlayerItemAccessLogEvent) SetPlaybackType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackType:"), value)
}


// The accumulated duration, in seconds, of the media segments downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/segmentsdownloadedduration
func (p_ PlayerItemAccessLogEvent) SegmentsDownloadedDuration() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("segmentsDownloadedDuration"))
	return rv
}


// The accumulated duration, in seconds, of the media segments downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/segmentsdownloadedduration
func (p_ PlayerItemAccessLogEvent) SetSegmentsDownloadedDuration(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSegmentsDownloadedDuration:"), value)
}


// The IP address of the server that was the source of the last delivered media segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/serveraddress
func (p_ PlayerItemAccessLogEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("serverAddress"))
	return rv
}


// The IP address of the server that was the source of the last delivered media segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/serveraddress
func (p_ PlayerItemAccessLogEvent) SetServerAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setServerAddress:"), value)
}


// The accumulated duration, in seconds, until the player item is ready to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/startuptime
func (p_ PlayerItemAccessLogEvent) StartupTime() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("startupTime"))
	return rv
}


// The accumulated duration, in seconds, until the player item is ready to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/startuptime
func (p_ PlayerItemAccessLogEvent) SetStartupTime(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartupTime:"), value)
}


// The bandwidth value that causes a switch, up or down, in the item’s quality being played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/switchbitrate
func (p_ PlayerItemAccessLogEvent) SwitchBitrate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("switchBitrate"))
	return rv
}


// The bandwidth value that causes a switch, up or down, in the item’s quality being played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/switchbitrate
func (p_ PlayerItemAccessLogEvent) SetSwitchBitrate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSwitchBitrate:"), value)
}


// The accumulated duration, in seconds, of active network transfer of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/transferduration
func (p_ PlayerItemAccessLogEvent) TransferDuration() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("transferDuration"))
	return rv
}


// The accumulated duration, in seconds, of active network transfer of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/transferduration
func (p_ PlayerItemAccessLogEvent) SetTransferDuration(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransferDuration:"), value)
}


// The URI of the playback item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/uri
func (p_ PlayerItemAccessLogEvent) Uri() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("uri"))
	return rv
}


// The URI of the playback item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/uri
func (p_ PlayerItemAccessLogEvent) SetUri(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUri:"), value)
}



