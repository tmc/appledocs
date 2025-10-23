// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerInterstitialEvent] class.
var (
	PlayerInterstitialEventClass     _PlayerInterstitialEventClass
	PlayerInterstitialEventClassOnce sync.Once
)

func getPlayerInterstitialEventClass() _PlayerInterstitialEventClass {
	PlayerInterstitialEventClassOnce.Do(func() {
		PlayerInterstitialEventClass = _PlayerInterstitialEventClass{objc.GetClass("AVPlayerInterstitialEvent")}
	})
	return PlayerInterstitialEventClass
}

type _PlayerInterstitialEventClass struct {
	class objc.Class
}

// An interface definition for the [PlayerInterstitialEvent] class.
type IPlayerInterstitialEvent interface {
	objectivec.IObject
	UserDefinedAttributes() objc.ID
	AlignsResumptionWithPrimarySegmentBoundary() bool
	SetAlignsResumptionWithPrimarySegmentBoundary(value bool)
	AlignsStartWithPrimarySegmentBoundary() bool
	SetAlignsStartWithPrimarySegmentBoundary(value bool)
	AssetListResponse() unsafe.Pointer
	SetAssetListResponse(value unsafe.Pointer)
	ContentMayVary() bool
	SetContentMayVary(value bool)
	Cue() unsafe.Pointer
	SetCue(value unsafe.Pointer)
	Date() foundation.Date
	SetDate(value foundation.IDate)
	Identifier() string
	SetIdentifier(value string)
	PlannedDuration() unsafe.Pointer
	SetPlannedDuration(value unsafe.Pointer)
	PlayoutLimit() unsafe.Pointer
	SetPlayoutLimit(value unsafe.Pointer)
	PrimaryItem() AVPlayerItem
	SetPrimaryItem(value IAVPlayerItem)
	Restrictions() unsafe.Pointer
	SetRestrictions(value unsafe.Pointer)
	ResumptionOffset() unsafe.Pointer
	SetResumptionOffset(value unsafe.Pointer)
	SkipControlLocalizedLabelBundleKey() string
	SetSkipControlLocalizedLabelBundleKey(value string)
	SkipControlTimeRange() unsafe.Pointer
	SetSkipControlTimeRange(value unsafe.Pointer)
	SupplementsPrimaryContent() bool
	SetSupplementsPrimaryContent(value bool)
	TemplateItems() AVPlayerItem
	SetTemplateItems(value IAVPlayerItem)
	Time() unsafe.Pointer
	SetTime(value unsafe.Pointer)
	TimelineOccupancy() unsafe.Pointer
	SetTimelineOccupancy(value unsafe.Pointer)
	WillPlayOnce() bool
	SetWillPlayOnce(value bool)
}

// An object that provides instructions for how a player presents interstitial content.
//
// An interstitial event defines a or , on the timeline of its , at which playback of interstitial content begins. It specifies the alternative interstitial content to play as an array of one or more template player items. The system uses the configuration of the event’s to build new player item instances to present the interstitial content. Use to observe the scheduling and progress of interstitial events. If your app requires specifying the schedule of interstitial events, use instead.


// An object that provides instructions for how a player presents interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent
type PlayerInterstitialEvent struct {
	objectivec.Object
}

// PlayerInterstitialEventFrom constructs a [PlayerInterstitialEvent] from an unsafe.Pointer.
//
// An object that provides instructions for how a player presents interstitial content.
func PlayerInterstitialEventFrom(ptr unsafe.Pointer) PlayerInterstitialEvent {
	return PlayerInterstitialEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerInterstitialEventClass) Alloc() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerInterstitialEventClass) New() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerInterstitialEvent) Init() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerInterstitialEvent) Autorelease() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerInterstitialEvent creates a new PlayerInterstitialEvent instance.
func NewPlayerInterstitialEvent() PlayerInterstitialEvent {
	return getPlayerInterstitialEventClass().New()
}



// Attributes of the event that the vendor or app defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/userDefinedAttributes
func (p_ PlayerInterstitialEvent) UserDefinedAttributes() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("userDefinedAttributes"))
	return rv
}


// A Boolean value that indicates whether the resumption time of primary playback should snap to a segment boundary of the primary asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/alignsresumptionwithprimarysegmentboundary
func (p_ PlayerInterstitialEvent) AlignsResumptionWithPrimarySegmentBoundary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("alignsResumptionWithPrimarySegmentBoundary"))
	return rv
}


// A Boolean value that indicates whether the resumption time of primary playback should snap to a segment boundary of the primary asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/alignsresumptionwithprimarysegmentboundary
func (p_ PlayerInterstitialEvent) SetAlignsResumptionWithPrimarySegmentBoundary(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignsResumptionWithPrimarySegmentBoundary:"), value)
}


// A Boolean value that indicates whether the start time of interstitial playback should snap to a segment boundary of the primary asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/alignsstartwithprimarysegmentboundary
func (p_ PlayerInterstitialEvent) AlignsStartWithPrimarySegmentBoundary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("alignsStartWithPrimarySegmentBoundary"))
	return rv
}


// A Boolean value that indicates whether the start time of interstitial playback should snap to a segment boundary of the primary asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/alignsstartwithprimarysegmentboundary
func (p_ PlayerInterstitialEvent) SetAlignsStartWithPrimarySegmentBoundary(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignsStartWithPrimarySegmentBoundary:"), value)
}


// The asset list JSON response as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/assetlistresponse
func (p_ PlayerInterstitialEvent) AssetListResponse() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("assetListResponse"))
	return rv
}


// The asset list JSON response as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/assetlistresponse
func (p_ PlayerInterstitialEvent) SetAssetListResponse(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAssetListResponse:"), value)
}


// A Boolean value that indicates whether an event’s content is dynamic and the server may respond with different interstitial assets for other participants in a coordinated playback session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/contentmayvary
func (p_ PlayerInterstitialEvent) ContentMayVary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("contentMayVary"))
	return rv
}


// A Boolean value that indicates whether an event’s content is dynamic and the server may respond with different interstitial assets for other participants in a coordinated playback session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/contentmayvary
func (p_ PlayerInterstitialEvent) SetContentMayVary(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentMayVary:"), value)
}


// A cue to schedule interstitial event playback at a predefined position during primary playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/cue-swift.property
func (p_ PlayerInterstitialEvent) Cue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cue"))
	return rv
}


// A cue to schedule interstitial event playback at a predefined position during primary playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/cue-swift.property
func (p_ PlayerInterstitialEvent) SetCue(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCue:"), value)
}


// A date within the date range of the primary content that playback of interstitial content begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/date
func (p_ PlayerInterstitialEvent) Date() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("date"))
	return rv
}


// A date within the date range of the primary content that playback of interstitial content begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/date
func (p_ PlayerInterstitialEvent) SetDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDate:"), value)
}


// An identifier for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/identifier
func (p_ PlayerInterstitialEvent) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}


// An identifier for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/identifier
func (p_ PlayerInterstitialEvent) SetIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The planned duration of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/plannedduration
func (p_ PlayerInterstitialEvent) PlannedDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("plannedDuration"))
	return rv
}


// The planned duration of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/plannedduration
func (p_ PlayerInterstitialEvent) SetPlannedDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlannedDuration:"), value)
}


// The time offset at which playback of the interstitial ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/playoutlimit
func (p_ PlayerInterstitialEvent) PlayoutLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playoutLimit"))
	return rv
}


// The time offset at which playback of the interstitial ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/playoutlimit
func (p_ PlayerInterstitialEvent) SetPlayoutLimit(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayoutLimit:"), value)
}


// The player item that represents the primary content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/primaryitem
func (p_ PlayerInterstitialEvent) PrimaryItem() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p_.ID, objc.Sel("primaryItem"))
	return rv
}


// The player item that represents the primary content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/primaryitem
func (p_ PlayerInterstitialEvent) SetPrimaryItem(value IAVPlayerItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrimaryItem:"), value)
}


// The restrictions the event imposes on the playback of interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/restrictions-swift.property
func (p_ PlayerInterstitialEvent) Restrictions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("restrictions"))
	return rv
}


// The restrictions the event imposes on the playback of interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/restrictions-swift.property
func (p_ PlayerInterstitialEvent) SetRestrictions(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRestrictions:"), value)
}


// A time offset at which playback of primary content resumes after interstitial content finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/resumptionoffset
func (p_ PlayerInterstitialEvent) ResumptionOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("resumptionOffset"))
	return rv
}


// A time offset at which playback of primary content resumes after interstitial content finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/resumptionoffset
func (p_ PlayerInterstitialEvent) SetResumptionOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setResumptionOffset:"), value)
}


// The key defined in the AVPlayerInterstitialEventController’s localizedStringsBundle that points to the localized label for the skip button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/skipcontrollocalizedlabelbundlekey
func (p_ PlayerInterstitialEvent) SkipControlLocalizedLabelBundleKey() string {
	rv := objc.Send[string](p_.ID, objc.Sel("skipControlLocalizedLabelBundleKey"))
	return rv
}


// The key defined in the AVPlayerInterstitialEventController’s localizedStringsBundle that points to the localized label for the skip button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/skipcontrollocalizedlabelbundlekey
func (p_ PlayerInterstitialEvent) SetSkipControlLocalizedLabelBundleKey(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkipControlLocalizedLabelBundleKey:"), objc.String(value))
}


// The time range within the duration of the interstitial event for which a skip button should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/skipcontroltimerange
func (p_ PlayerInterstitialEvent) SkipControlTimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("skipControlTimeRange"))
	return rv
}


// The time range within the duration of the interstitial event for which a skip button should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/skipcontroltimerange
func (p_ PlayerInterstitialEvent) SetSkipControlTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkipControlTimeRange:"), value)
}


// A Boolean value that indicates whether an event supplements the primary content and should present with the primary item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/supplementsprimarycontent
func (p_ PlayerInterstitialEvent) SupplementsPrimaryContent() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("supplementsPrimaryContent"))
	return rv
}


// A Boolean value that indicates whether an event supplements the primary content and should present with the primary item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/supplementsprimarycontent
func (p_ PlayerInterstitialEvent) SetSupplementsPrimaryContent(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSupplementsPrimaryContent:"), value)
}


// An array of player item configurations to use as templates for player items that play interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p_ PlayerInterstitialEvent) TemplateItems() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p_.ID, objc.Sel("templateItems"))
	return rv
}


// An array of player item configurations to use as templates for player items that play interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p_ PlayerInterstitialEvent) SetTemplateItems(value IAVPlayerItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTemplateItems:"), value)
}


// A time within the timeline of the primary content that playback of interstitial content begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/time
func (p_ PlayerInterstitialEvent) Time() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("time"))
	return rv
}


// A time within the timeline of the primary content that playback of interstitial content begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/time
func (p_ PlayerInterstitialEvent) SetTime(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTime:"), value)
}


// An event’s occupancy on the integrated timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/timelineoccupancy-swift.property
func (p_ PlayerInterstitialEvent) TimelineOccupancy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("timelineOccupancy"))
	return rv
}


// An event’s occupancy on the integrated timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/timelineoccupancy-swift.property
func (p_ PlayerInterstitialEvent) SetTimelineOccupancy(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimelineOccupancy:"), value)
}


// A Boolean value that indicates whether to schedule this event one time only and suppress subsequent replay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/willplayonce
func (p_ PlayerInterstitialEvent) WillPlayOnce() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("willPlayOnce"))
	return rv
}


// A Boolean value that indicates whether to schedule this event one time only and suppress subsequent replay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/willplayonce
func (p_ PlayerInterstitialEvent) SetWillPlayOnce(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWillPlayOnce:"), value)
}



