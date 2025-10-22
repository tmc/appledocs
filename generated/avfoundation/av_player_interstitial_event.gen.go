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
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/userDefinedAttributes
func (p_ PlayerInterstitialEvent) UserDefinedAttributes() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("userDefinedAttributes"))
	return rv
}

// A Boolean value that indicates whether the resumption time of primary playback should snap to a segment boundary of the primary asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/alignsresumptionwithprimarysegmentboundary
func (p_ PlayerInterstitialEvent) AlignsResumptionWithPrimarySegmentBoundary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("alignsResumptionWithPrimarySegmentBoundary"))
	return rv
}


// SetAlignsResumptionWithPrimarySegmentBoundary sets the value of the alignsResumptionWithPrimarySegmentBoundary property.
// A Boolean value that indicates whether the resumption time of primary playback should snap to a segment boundary of the primary asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/alignsresumptionwithprimarysegmentboundary
func (p_ PlayerInterstitialEvent) SetAlignsResumptionWithPrimarySegmentBoundary(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignsResumptionWithPrimarySegmentBoundary:"), value)
}

// A Boolean value that indicates whether the start time of interstitial playback should snap to a segment boundary of the primary asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/alignsstartwithprimarysegmentboundary
func (p_ PlayerInterstitialEvent) AlignsStartWithPrimarySegmentBoundary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("alignsStartWithPrimarySegmentBoundary"))
	return rv
}


// SetAlignsStartWithPrimarySegmentBoundary sets the value of the alignsStartWithPrimarySegmentBoundary property.
// A Boolean value that indicates whether the start time of interstitial playback should snap to a segment boundary of the primary asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/alignsstartwithprimarysegmentboundary
func (p_ PlayerInterstitialEvent) SetAlignsStartWithPrimarySegmentBoundary(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignsStartWithPrimarySegmentBoundary:"), value)
}

// The asset list JSON response as a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/assetlistresponse
func (p_ PlayerInterstitialEvent) AssetListResponse() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("assetListResponse"))
	return rv
}


// SetAssetListResponse sets the value of the assetListResponse property.
// The asset list JSON response as a dictionary.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/assetlistresponse
func (p_ PlayerInterstitialEvent) SetAssetListResponse(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAssetListResponse:"), value)
}

// A Boolean value that indicates whether an event’s content is dynamic and the server may respond with different interstitial assets for other participants in a coordinated playback session.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/contentmayvary
func (p_ PlayerInterstitialEvent) ContentMayVary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("contentMayVary"))
	return rv
}


// SetContentMayVary sets the value of the contentMayVary property.
// A Boolean value that indicates whether an event’s content is dynamic and the server may respond with different interstitial assets for other participants in a coordinated playback session.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/contentmayvary
func (p_ PlayerInterstitialEvent) SetContentMayVary(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentMayVary:"), value)
}

// A cue to schedule interstitial event playback at a predefined position during primary playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/cue-swift.property
func (p_ PlayerInterstitialEvent) Cue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cue"))
	return rv
}


// SetCue sets the value of the cue property.
// A cue to schedule interstitial event playback at a predefined position during primary playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/cue-swift.property
func (p_ PlayerInterstitialEvent) SetCue(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCue:"), value)
}

// A date within the date range of the primary content that playback of interstitial content begins.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/date
func (p_ PlayerInterstitialEvent) Date() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("date"))
	return rv
}


// SetDate sets the value of the date property.
// A date within the date range of the primary content that playback of interstitial content begins.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/date
func (p_ PlayerInterstitialEvent) SetDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDate:"), value)
}

// An identifier for the event.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/identifier
func (p_ PlayerInterstitialEvent) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// An identifier for the event.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/identifier
func (p_ PlayerInterstitialEvent) SetIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// The planned duration of the event.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/plannedduration
func (p_ PlayerInterstitialEvent) PlannedDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("plannedDuration"))
	return rv
}


// SetPlannedDuration sets the value of the plannedDuration property.
// The planned duration of the event.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/plannedduration
func (p_ PlayerInterstitialEvent) SetPlannedDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlannedDuration:"), value)
}

// The time offset at which playback of the interstitial ends.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/playoutlimit
func (p_ PlayerInterstitialEvent) PlayoutLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playoutLimit"))
	return rv
}


// SetPlayoutLimit sets the value of the playoutLimit property.
// The time offset at which playback of the interstitial ends.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/playoutlimit
func (p_ PlayerInterstitialEvent) SetPlayoutLimit(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayoutLimit:"), value)
}

// The player item that represents the primary content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/primaryitem
func (p_ PlayerInterstitialEvent) PrimaryItem() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p_.ID, objc.Sel("primaryItem"))
	return rv
}


// SetPrimaryItem sets the value of the primaryItem property.
// The player item that represents the primary content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/primaryitem
func (p_ PlayerInterstitialEvent) SetPrimaryItem(value IAVPlayerItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrimaryItem:"), value)
}

// The restrictions the event imposes on the playback of interstitial content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/restrictions-swift.property
func (p_ PlayerInterstitialEvent) Restrictions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("restrictions"))
	return rv
}


// SetRestrictions sets the value of the restrictions property.
// The restrictions the event imposes on the playback of interstitial content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/restrictions-swift.property
func (p_ PlayerInterstitialEvent) SetRestrictions(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRestrictions:"), value)
}

// A time offset at which playback of primary content resumes after interstitial content finishes.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/resumptionoffset
func (p_ PlayerInterstitialEvent) ResumptionOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("resumptionOffset"))
	return rv
}


// SetResumptionOffset sets the value of the resumptionOffset property.
// A time offset at which playback of primary content resumes after interstitial content finishes.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/resumptionoffset
func (p_ PlayerInterstitialEvent) SetResumptionOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setResumptionOffset:"), value)
}

// The key defined in the AVPlayerInterstitialEventController’s localizedStringsBundle that points to the localized label for the skip button.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/skipcontrollocalizedlabelbundlekey
func (p_ PlayerInterstitialEvent) SkipControlLocalizedLabelBundleKey() string {
	rv := objc.Send[string](p_.ID, objc.Sel("skipControlLocalizedLabelBundleKey"))
	return rv
}


// SetSkipControlLocalizedLabelBundleKey sets the value of the skipControlLocalizedLabelBundleKey property.
// The key defined in the AVPlayerInterstitialEventController’s localizedStringsBundle that points to the localized label for the skip button.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/skipcontrollocalizedlabelbundlekey
func (p_ PlayerInterstitialEvent) SetSkipControlLocalizedLabelBundleKey(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkipControlLocalizedLabelBundleKey:"), objc.String(value))
}

// The time range within the duration of the interstitial event for which a skip button should be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/skipcontroltimerange
func (p_ PlayerInterstitialEvent) SkipControlTimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("skipControlTimeRange"))
	return rv
}


// SetSkipControlTimeRange sets the value of the skipControlTimeRange property.
// The time range within the duration of the interstitial event for which a skip button should be displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/skipcontroltimerange
func (p_ PlayerInterstitialEvent) SetSkipControlTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkipControlTimeRange:"), value)
}

// A Boolean value that indicates whether an event supplements the primary content and should present with the primary item.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/supplementsprimarycontent
func (p_ PlayerInterstitialEvent) SupplementsPrimaryContent() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("supplementsPrimaryContent"))
	return rv
}


// SetSupplementsPrimaryContent sets the value of the supplementsPrimaryContent property.
// A Boolean value that indicates whether an event supplements the primary content and should present with the primary item.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/supplementsprimarycontent
func (p_ PlayerInterstitialEvent) SetSupplementsPrimaryContent(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSupplementsPrimaryContent:"), value)
}

// An array of player item configurations to use as templates for player items that play interstitial content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p_ PlayerInterstitialEvent) TemplateItems() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p_.ID, objc.Sel("templateItems"))
	return rv
}


// SetTemplateItems sets the value of the templateItems property.
// An array of player item configurations to use as templates for player items that play interstitial content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p_ PlayerInterstitialEvent) SetTemplateItems(value IAVPlayerItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTemplateItems:"), value)
}

// A time within the timeline of the primary content that playback of interstitial content begins.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/time
func (p_ PlayerInterstitialEvent) Time() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("time"))
	return rv
}


// SetTime sets the value of the time property.
// A time within the timeline of the primary content that playback of interstitial content begins.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/time
func (p_ PlayerInterstitialEvent) SetTime(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTime:"), value)
}

// An event’s occupancy on the integrated timeline.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/timelineoccupancy-swift.property
func (p_ PlayerInterstitialEvent) TimelineOccupancy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("timelineOccupancy"))
	return rv
}


// SetTimelineOccupancy sets the value of the timelineOccupancy property.
// An event’s occupancy on the integrated timeline.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/timelineoccupancy-swift.property
func (p_ PlayerInterstitialEvent) SetTimelineOccupancy(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimelineOccupancy:"), value)
}

// A Boolean value that indicates whether to schedule this event one time only and suppress subsequent replay.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/willplayonce
func (p_ PlayerInterstitialEvent) WillPlayOnce() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("willPlayOnce"))
	return rv
}


// SetWillPlayOnce sets the value of the willPlayOnce property.
// A Boolean value that indicates whether to schedule this event one time only and suppress subsequent replay.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/willplayonce
func (p_ PlayerInterstitialEvent) SetWillPlayOnce(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWillPlayOnce:"), value)
}



