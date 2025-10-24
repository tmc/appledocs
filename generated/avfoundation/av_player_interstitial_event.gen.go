// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerInterstitialEvent */


/* debug [class_header]: Header for AVPlayerInterstitialEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerInterstitialEvent */
// An interface definition for the [PlayerInterstitialEvent] class.
type IPlayerInterstitialEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerInterstitialEvent */
	// properties:
	AlignsResumptionWithPrimarySegmentBoundary() bool
	AlignsStartWithPrimarySegmentBoundary() bool
	AssetListResponse() objc.IObject /* cross-framework: NSDictionary */
	ContentMayVary() bool
	Cue() PlayerInterstitialEventCue /* typedef */
	Date() objc.IObject /* cross-framework: NSDate */
	Identifier() objc.IObject /* cross-framework: NSString */
	PlannedDuration() objc.IObject /* cross-framework: Time */
	SetPlannedDuration(value objc.IObject /* cross-framework: Time */)
	PlayoutLimit() objc.IObject /* cross-framework: Time */
	PrimaryItem() IAVPlayerItem
	Restrictions() PlayerInterstitialEventRestrictions
	ResumptionOffset() objc.IObject /* cross-framework: Time */
	SkipControlLocalizedLabelBundleKey() objc.IObject /* cross-framework: NSString */
	SkipControlTimeRange() TimeRange /* not a class type */
	SupplementsPrimaryContent() bool
	TemplateItems() []PlayerItem
	Time() objc.IObject /* cross-framework: Time */
	TimelineOccupancy() PlayerInterstitialEventTimelineOccupancy
	UserDefinedAttributes() objc.IObject /* cross-framework: NSDictionary */
	WillPlayOnce() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerInterstitialEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerInterstitialEvent */
// Alloc allocates a new instance without initialization.
func (pc _PlayerInterstitialEventClass) Alloc() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerInterstitialEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerInterstitialEvent */

// Creates an interstitial event for the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/init(primaryItem:date:)
func NewPlayerInterstitialEventWithPrimaryItemDate(primaryItem IAVPlayerItem, date objc.IObject /* cross-framework: NSDate */) PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](objc.ID(getPlayerInterstitialEventClass().class), objc.Sel("interstitialEventWithPrimaryItem:date:"), primaryItem, date)
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerInterstitialEventWithPrimaryItemDate */


// Creates an interstitial event for the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/init(primaryItem:time:)
func NewPlayerInterstitialEventWithPrimaryItemTime(primaryItem IAVPlayerItem, time objc.IObject /* cross-framework: Time */) PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](objc.ID(getPlayerInterstitialEventClass().class), objc.Sel("interstitialEventWithPrimaryItem:time:"), primaryItem, time)
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerInterstitialEventWithPrimaryItemTime */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerInterstitialEvent */

// Creates an interstitial event for the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/init(primaryItem:date:)
func (pc _PlayerInterstitialEventClass) InterstitialEventWithPrimaryItemDate(primaryItem IAVPlayerItem, date objc.IObject /* cross-framework: NSDate */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("interstitialEventWithPrimaryItem:date:"), primaryItem, date)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterstitialEventWithPrimaryItemDate) */


// Creates an interstitial event for the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/init(primaryItem:time:)
func (pc _PlayerInterstitialEventClass) InterstitialEventWithPrimaryItemTime(primaryItem IAVPlayerItem, time objc.IObject /* cross-framework: Time */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("interstitialEventWithPrimaryItem:time:"), primaryItem, time)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterstitialEventWithPrimaryItemTime) */


// Creates an interstitial event, with user-defined attributes, for the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/interstitialEventWithPrimaryItem:identifier:date:templateItems:restrictions:resumptionOffset:playoutLimit:userDefinedAttributes:
func (pc _PlayerInterstitialEventClass) InterstitialEventWithPrimaryItemIdentifierDateTemplateItemsRestrictionsResumptionOffsetPlayoutLimitUserDefinedAttributes(primaryItem IAVPlayerItem, identifier objc.IObject /* cross-framework: NSString */, date objc.IObject /* cross-framework: NSDate */, templateItems []PlayerItem, restrictions PlayerInterstitialEventRestrictions, resumptionOffset objc.IObject /* cross-framework: Time */, playoutLimit objc.IObject /* cross-framework: Time */, userDefinedAttributes objc.IObject /* cross-framework: NSDictionary */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("interstitialEventWithPrimaryItem:identifier:date:templateItems:restrictions:resumptionOffset:playoutLimit:userDefinedAttributes:"), primaryItem, identifier, date, templateItems, restrictions, resumptionOffset, playoutLimit, userDefinedAttributes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterstitialEventWithPrimaryItemIdentifierDateTemplateItemsRestrictionsResumptionOffsetPlayoutLimitUserDefinedAttributes) */


// Creates an interstitial event, with user-defined attributes, for the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/interstitialEventWithPrimaryItem:identifier:time:templateItems:restrictions:resumptionOffset:playoutLimit:userDefinedAttributes:
func (pc _PlayerInterstitialEventClass) InterstitialEventWithPrimaryItemIdentifierTimeTemplateItemsRestrictionsResumptionOffsetPlayoutLimitUserDefinedAttributes(primaryItem IAVPlayerItem, identifier objc.IObject /* cross-framework: NSString */, time objc.IObject /* cross-framework: Time */, templateItems []PlayerItem, restrictions PlayerInterstitialEventRestrictions, resumptionOffset objc.IObject /* cross-framework: Time */, playoutLimit objc.IObject /* cross-framework: Time */, userDefinedAttributes objc.IObject /* cross-framework: NSDictionary */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("interstitialEventWithPrimaryItem:identifier:time:templateItems:restrictions:resumptionOffset:playoutLimit:userDefinedAttributes:"), primaryItem, identifier, time, templateItems, restrictions, resumptionOffset, playoutLimit, userDefinedAttributes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterstitialEventWithPrimaryItemIdentifierTimeTemplateItemsRestrictionsResumptionOffsetPlayoutLimitUserDefinedAttributes) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerInterstitialEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerInterstitialEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerInterstitialEvent */

// A Boolean value that indicates whether the resumption time of primary playback should snap to a segment boundary of the primary asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/alignsResumptionWithPrimarySegmentBoundary
func (p_ PlayerInterstitialEvent) AlignsResumptionWithPrimarySegmentBoundary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("alignsResumptionWithPrimarySegmentBoundary"))
	return rv
}/* debug [instance_properties/getter]: alignsResumptionWithPrimarySegmentBoundary */


// A Boolean value that indicates whether the start time of interstitial playback should snap to a segment boundary of the primary asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/alignsStartWithPrimarySegmentBoundary
func (p_ PlayerInterstitialEvent) AlignsStartWithPrimarySegmentBoundary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("alignsStartWithPrimarySegmentBoundary"))
	return rv
}/* debug [instance_properties/getter]: alignsStartWithPrimarySegmentBoundary */


// The asset list JSON response as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/assetListResponse
func (p_ PlayerInterstitialEvent) AssetListResponse() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("assetListResponse"))
	return rv
}/* debug [instance_properties/getter]: assetListResponse */


// A Boolean value that indicates whether an event’s content is dynamic and the server may respond with different interstitial assets for other participants in a coordinated playback session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/contentMayVary
func (p_ PlayerInterstitialEvent) ContentMayVary() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("contentMayVary"))
	return rv
}/* debug [instance_properties/getter]: contentMayVary */


// A cue to schedule interstitial event playback at a predefined position during primary playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/cue-swift.property
func (p_ PlayerInterstitialEvent) Cue() PlayerInterstitialEventCue /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("cue"))
	return rv
}/* debug [instance_properties/getter]: cue */


// A date within the date range of the primary content that playback of interstitial content begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/date
func (p_ PlayerInterstitialEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// An identifier for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/identifier
func (p_ PlayerInterstitialEvent) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The planned duration of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/plannedDuration
func (p_ PlayerInterstitialEvent) PlannedDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("plannedDuration"))
	return rv
}/* debug [instance_properties/getter]: plannedDuration */


// The planned duration of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/plannedDuration
func (p_ PlayerInterstitialEvent) SetPlannedDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlannedDuration:"), value)
}/* debug [instance_properties/setter]: plannedDuration */


// The time offset at which playback of the interstitial ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/playoutLimit
func (p_ PlayerInterstitialEvent) PlayoutLimit() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("playoutLimit"))
	return rv
}/* debug [instance_properties/getter]: playoutLimit */


// The player item that represents the primary content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/primaryItem
func (p_ PlayerInterstitialEvent) PrimaryItem() IAVPlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("primaryItem"))
	return rv
}/* debug [instance_properties/getter]: primaryItem */


// The restrictions the event imposes on the playback of interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/restrictions-swift.property
func (p_ PlayerInterstitialEvent) Restrictions() PlayerInterstitialEventRestrictions {
	rv := objc.Send[PlayerInterstitialEventRestrictions](p_.ID, objc.Sel("restrictions"))
	return rv
}/* debug [instance_properties/getter]: restrictions */


// A time offset at which playback of primary content resumes after interstitial content finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/resumptionOffset
func (p_ PlayerInterstitialEvent) ResumptionOffset() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("resumptionOffset"))
	return rv
}/* debug [instance_properties/getter]: resumptionOffset */


// The key defined in the AVPlayerInterstitialEventController’s localizedStringsBundle that points to the localized label for the skip button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/skipControlLocalizedLabelBundleKey
func (p_ PlayerInterstitialEvent) SkipControlLocalizedLabelBundleKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("skipControlLocalizedLabelBundleKey"))
	return rv
}/* debug [instance_properties/getter]: skipControlLocalizedLabelBundleKey */


// The time range within the duration of the interstitial event for which a skip button should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/skipControlTimeRange
func (p_ PlayerInterstitialEvent) SkipControlTimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](p_.ID, objc.Sel("skipControlTimeRange"))
	return rv
}/* debug [instance_properties/getter]: skipControlTimeRange */


// A Boolean value that indicates whether an event supplements the primary content and should present with the primary item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/supplementsPrimaryContent
func (p_ PlayerInterstitialEvent) SupplementsPrimaryContent() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("supplementsPrimaryContent"))
	return rv
}/* debug [instance_properties/getter]: supplementsPrimaryContent */


// An array of player item configurations to use as templates for player items that play interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/templateItems
func (p_ PlayerInterstitialEvent) TemplateItems() []PlayerItem {
	rv := objc.Send[[]PlayerItem](p_.ID, objc.Sel("templateItems"))
	return rv
}/* debug [instance_properties/getter]: templateItems */


// A time within the timeline of the primary content that playback of interstitial content begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/time
func (p_ PlayerInterstitialEvent) Time() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("time"))
	return rv
}/* debug [instance_properties/getter]: time */


// An event’s occupancy on the integrated timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/timelineOccupancy-swift.property
func (p_ PlayerInterstitialEvent) TimelineOccupancy() PlayerInterstitialEventTimelineOccupancy {
	rv := objc.Send[PlayerInterstitialEventTimelineOccupancy](p_.ID, objc.Sel("timelineOccupancy"))
	return rv
}/* debug [instance_properties/getter]: timelineOccupancy */


// Attributes of the event that the vendor or app defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/userDefinedAttributes
func (p_ PlayerInterstitialEvent) UserDefinedAttributes() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("userDefinedAttributes"))
	return rv
}/* debug [instance_properties/getter]: userDefinedAttributes */


// A Boolean value that indicates whether to schedule this event one time only and suppress subsequent replay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/willPlayOnce
func (p_ PlayerInterstitialEvent) WillPlayOnce() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("willPlayOnce"))
	return rv
}/* debug [instance_properties/getter]: willPlayOnce */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerInterstitialEvent */


