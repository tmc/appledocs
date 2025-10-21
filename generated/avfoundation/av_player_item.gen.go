// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PlayerItem] class.
var (
	PlayerItemClass     _PlayerItemClass
	PlayerItemClassOnce sync.Once
)

func getPlayerItemClass() _PlayerItemClass {
	PlayerItemClassOnce.Do(func() {
		PlayerItemClass = _PlayerItemClass{objc.GetClass("AVPlayerItem")}
	})
	return PlayerItemClass
}

type _PlayerItemClass struct {
	class objc.Class
}

// An interface definition for the [PlayerItem] class.
type IPlayerItem interface {
	objectivec.IObject
	CancelPlaybackRestrictionsAuthorizationRequest()
	CurrentTime() unsafe.Pointer
	EffectiveMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup unsafe.Pointer) unsafe.Pointer
	RequestPlaybackRestrictionsAuthorization(completion unsafe.Pointer)
	SelectMediaPresentationSettingForMediaSelectionGroup(mediaPresentationSetting unsafe.Pointer, mediaSelectionGroup unsafe.Pointer)
	SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption unsafe.Pointer, mediaSelectionGroup unsafe.Pointer)
	SelectMediaOptionAutomaticallyInMediaSelectionGroup(mediaSelectionGroup unsafe.Pointer)
	SelectMediaPresentationLanguageForMediaSelectionGroup(language string, mediaSelectionGroup unsafe.Pointer)
	SelectedMediaPresentationLanguageForMediaSelectionGroup(mediaSelectionGroup unsafe.Pointer) string
	SelectedMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup unsafe.Pointer) unsafe.Pointer
}

// An object that models the timing and presentation state of an asset during playback.
//
// A player item stores a reference to an object, which represents the media to play. If you require inspecting an asset before you enqueue it for playback, call its method to retrieve the values of one or more properties. Alternatively, you can tell the player item to automatically load the required properties by passing them to its initializer. When the player item is ready to play, the asset properties you request are ready to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem
type PlayerItem struct {
	objectivec.Object
}

// PlayerItemFrom constructs a [PlayerItem] from an unsafe.Pointer.
//
// An object that models the timing and presentation state of an asset during playback.
func PlayerItemFrom(ptr unsafe.Pointer) PlayerItem {
	return PlayerItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemClass) Alloc() PlayerItem {
	rv := objc.Send[PlayerItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerItemClass) New() PlayerItem {
	rv := objc.Send[PlayerItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItem) Init() PlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItem) Autorelease() PlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItem creates a new PlayerItem instance.
func NewPlayerItem() PlayerItem {
	return getPlayerItemClass().New()
}


// Cancels a pending authorization request and dismisses the passcode entry, if displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/cancelPlaybackRestrictionsAuthorizationRequest()
func (p_ PlayerItem) CancelPlaybackRestrictionsAuthorizationRequest() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelPlaybackRestrictionsAuthorizationRequest"))
}

// Returns the current time of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentTime()
func (p_ PlayerItem) CurrentTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentTime"))
	return rv
}

// Indicates the media presentation settings with media characteristics that are possessed by the currently selected AVMediaSelectionOption in the specified AVMediaSelectionGroup.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/effectiveMediaPresentationSettings(for:)
func (p_ PlayerItem) EffectiveMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("effectiveMediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}

// Determines whether this item is subject to parental restrictions, and, if so, prompts the user to enter the restrictions passcode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/requestPlaybackRestrictionsAuthorization(_:)
func (p_ PlayerItem) RequestPlaybackRestrictionsAuthorization(completion unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("requestPlaybackRestrictionsAuthorization:"), completion)
}

// When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular presentation setting, replacing any previous preference for settings of the same media presentation selector.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/select(_:for:)
func (p_ PlayerItem) SelectMediaPresentationSettingForMediaSelectionGroup(mediaPresentationSetting unsafe.Pointer, mediaSelectionGroup unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectMediaPresentationSetting:forMediaSelectionGroup:"), mediaPresentationSetting, mediaSelectionGroup)
}

// Selects a media option in a given media selection group and deselects all other options in that group.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/select(_:in:)
func (p_ PlayerItem) SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption unsafe.Pointer, mediaSelectionGroup unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectMediaOption:inMediaSelectionGroup:"), mediaSelectionOption, mediaSelectionGroup)
}

// Selects the media option in the specified media selection group that best matches the receiver’s automatic selection criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectMediaOptionAutomatically(in:)
func (p_ PlayerItem) SelectMediaOptionAutomaticallyInMediaSelectionGroup(mediaSelectionGroup unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectMediaOptionAutomaticallyInMediaSelectionGroup:"), mediaSelectionGroup)
}

// When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular language, replacing any previous preference for available languages of the specified group’s custom media selection scheme.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectMediaPresentationLanguage(_:for:)
func (p_ PlayerItem) SelectMediaPresentationLanguageForMediaSelectionGroup(language string, mediaSelectionGroup unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectMediaPresentationLanguage:forMediaSelectionGroup:"), objc.String(language), mediaSelectionGroup)
}

// Returns the selected media presentation language for the specified media selection group, if any language has previously been selected via use of -selectMediaPresentationLanguages:forMediaSelectionGroup:.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectedMediaPresentationLanguage(for:)
func (p_ PlayerItem) SelectedMediaPresentationLanguageForMediaSelectionGroup(mediaSelectionGroup unsafe.Pointer) string {
	rv := objc.Send[string](p_.ID, objc.Sel("selectedMediaPresentationLanguageForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}

// Indicates the media presentation settings that have most recently been selected for each AVMediaPresentationSelector of the AVCustomMediaSelectionScheme of the specified AVMediaSelectionGroup.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectedMediaPresentationSettings(for:)
func (p_ PlayerItem) SelectedMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectedMediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}

// The source audio channel layouts the player item supports for spatialization.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/allowedAudioSpatializationFormats
func (p_ PlayerItem) AllowedAudioSpatializationFormats() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}


// SetAllowedAudioSpatializationFormats sets the value of the allowedAudioSpatializationFormats property.
// The source audio channel layouts the player item supports for spatialization.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/allowedAudioSpatializationFormats
func (p_ PlayerItem) SetAllowedAudioSpatializationFormats(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}

// A Boolean value that indicates whether the item can play slower than normal.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlaySlowForward
func (p_ PlayerItem) CanPlaySlowForward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlaySlowForward"))
	return rv
}

// The current media selections for each of the receiver’s media selection groups.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentMediaSelection
func (p_ PlayerItem) CurrentMediaSelection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentMediaSelection"))
	return rv
}

// The error that caused the player item to fail.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/error
func (p_ PlayerItem) Error() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("error"))
	return rv
}

// An array of additional metadata for the player item to supplement or replace an asset’s embedded metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/externalMetadata
func (p_ PlayerItem) ExternalMetadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](p_.ID, objc.Sel("externalMetadata"))
	return rv
}


// SetExternalMetadata sets the value of the externalMetadata property.
// An array of additional metadata for the player item to supplement or replace an asset’s embedded metadata.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/externalMetadata
func (p_ PlayerItem) SetExternalMetadata(value []MetadataItem) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setExternalMetadata:"), nsArray)
}

// The time at which forward playback ends.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/forwardPlaybackEndTime
func (p_ PlayerItem) ForwardPlaybackEndTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("forwardPlaybackEndTime"))
	return rv
}


// SetForwardPlaybackEndTime sets the value of the forwardPlaybackEndTime property.
// The time at which forward playback ends.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/forwardPlaybackEndTime
func (p_ PlayerItem) SetForwardPlaybackEndTime(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setForwardPlaybackEndTime:"), value)
}

// An array of time ranges that identify interstitial content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/interstitialTimeRanges
func (p_ PlayerItem) InterstitialTimeRanges() []AVInterstitialTimeRange {
	rv := objc.Send[[]AVInterstitialTimeRange](p_.ID, objc.Sel("interstitialTimeRanges"))
	return rv
}

// A Boolean value that indicates whether the player item allows spatialized audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isAudioSpatializationAllowed
func (p_ PlayerItem) AudioSpatializationAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("audioSpatializationAllowed"))
	return rv
}


// SetAudioSpatializationAllowed sets the value of the audioSpatializationAllowed property.
// A Boolean value that indicates whether the player item allows spatialized audio playback.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isAudioSpatializationAllowed
func (p_ PlayerItem) SetAudioSpatializationAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioSpatializationAllowed:"), value)
}

// The time marker groups that provide ways to navigate the player item’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/navigationMarkerGroups
func (p_ PlayerItem) NavigationMarkerGroups() []AVNavigationMarkersGroup {
	rv := objc.Send[[]AVNavigationMarkersGroup](p_.ID, objc.Sel("navigationMarkerGroups"))
	return rv
}


// SetNavigationMarkerGroups sets the value of the navigationMarkerGroups property.
// The time marker groups that provide ways to navigate the player item’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/navigationMarkerGroups
func (p_ PlayerItem) SetNavigationMarkerGroups(value []AVNavigationMarkersGroup) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setNavigationMarkerGroups:"), nsArray)
}

// The item proposed to follow the current content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/nextContentProposal
func (p_ PlayerItem) NextContentProposal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("nextContentProposal"))
	return rv
}


// SetNextContentProposal sets the value of the nextContentProposal property.
// The item proposed to follow the current content.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/nextContentProposal
func (p_ PlayerItem) SetNextContentProposal(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNextContentProposal:"), value)
}

// The current now playing information for the player item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/nowPlayingInfo
func (p_ PlayerItem) NowPlayingInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("nowPlayingInfo"))
	return rv
}


// SetNowPlayingInfo sets the value of the nowPlayingInfo property.
// The current now playing information for the player item.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/nowPlayingInfo
func (p_ PlayerItem) SetNowPlayingInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNowPlayingInfo:"), value)
}

// Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of the receiver’s asset with which an associated UI implementation should configure its interface for media selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredCustomMediaSelectionSchemes
func (p_ PlayerItem) PreferredCustomMediaSelectionSchemes() []AVCustomMediaSelectionScheme {
	rv := objc.Send[[]AVCustomMediaSelectionScheme](p_.ID, objc.Sel("preferredCustomMediaSelectionSchemes"))
	return rv
}


// SetPreferredCustomMediaSelectionSchemes sets the value of the preferredCustomMediaSelectionSchemes property.
// Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of the receiver’s asset with which an associated UI implementation should configure its interface for media selection.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredCustomMediaSelectionSchemes
func (p_ PlayerItem) SetPreferredCustomMediaSelectionSchemes(value []AVCustomMediaSelectionScheme) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredCustomMediaSelectionSchemes:"), nsArray)
}

// The desired limit, in bits per second, of network bandwidth consumption for this item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredPeakBitRate
func (p_ PlayerItem) PreferredPeakBitRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("preferredPeakBitRate"))
	return rv
}


// SetPreferredPeakBitRate sets the value of the preferredPeakBitRate property.
// The desired limit, in bits per second, of network bandwidth consumption for this item.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredPeakBitRate
func (p_ PlayerItem) SetPreferredPeakBitRate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredPeakBitRate:"), value)
}

// The time at which reverse playback ends.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/reversePlaybackEndTime
func (p_ PlayerItem) ReversePlaybackEndTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("reversePlaybackEndTime"))
	return rv
}


// SetReversePlaybackEndTime sets the value of the reversePlaybackEndTime property.
// The time at which reverse playback ends.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/reversePlaybackEndTime
func (p_ PlayerItem) SetReversePlaybackEndTime(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReversePlaybackEndTime:"), value)
}

// The status of the player item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/status-swift.property
func (p_ PlayerItem) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("status"))
	return rv
}

// A Boolean value that indicates whether the player translates interstitial events to interstitial time ranges.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/translatesPlayerInterstitialEvents
func (p_ PlayerItem) TranslatesPlayerInterstitialEvents() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("translatesPlayerInterstitialEvents"))
	return rv
}


// SetTranslatesPlayerInterstitialEvents sets the value of the translatesPlayerInterstitialEvents property.
// A Boolean value that indicates whether the player translates interstitial events to interstitial time ranges.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/translatesPlayerInterstitialEvents
func (p_ PlayerItem) SetTranslatesPlayerInterstitialEvents(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTranslatesPlayerInterstitialEvents:"), value)
}



