// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
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
func (p_ PlayerItem) InterstitialTimeRanges() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("interstitialTimeRanges"))
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
func (p_ PlayerItem) NavigationMarkerGroups() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("navigationMarkerGroups"))
	return rv
}


// SetNavigationMarkerGroups sets the value of the navigationMarkerGroups property.
// The time marker groups that provide ways to navigate the player item’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/navigationMarkerGroups
func (p_ PlayerItem) SetNavigationMarkerGroups(value []unsafe.Pointer) {
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
func (p_ PlayerItem) PreferredCustomMediaSelectionSchemes() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("preferredCustomMediaSelectionSchemes"))
	return rv
}


// SetPreferredCustomMediaSelectionSchemes sets the value of the preferredCustomMediaSelectionSchemes property.
// Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of the receiver’s asset with which an associated UI implementation should configure its interface for media selection.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredCustomMediaSelectionSchemes
func (p_ PlayerItem) SetPreferredCustomMediaSelectionSchemes(value []unsafe.Pointer) {
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

// A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/appliesperframehdrdisplaymetadata
func (p_ PlayerItem) AppliesPerFrameHDRDisplayMetadata() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("appliesPerFrameHDRDisplayMetadata"))
	return rv
}


// SetAppliesPerFrameHDRDisplayMetadata sets the value of the appliesPerFrameHDRDisplayMetadata property.
// A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/appliesperframehdrdisplaymetadata
func (p_ PlayerItem) SetAppliesPerFrameHDRDisplayMetadata(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppliesPerFrameHDRDisplayMetadata:"), value)
}

// The asset provided during initialization.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/asset
func (p_ PlayerItem) Asset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("asset"))
	return rv
}


// SetAsset sets the value of the asset property.
// The asset provided during initialization.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/asset
func (p_ PlayerItem) SetAsset(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAsset:"), value)
}

// The audio mix parameters to be applied during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/audiomix
func (p_ PlayerItem) AudioMix() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("audioMix"))
	return rv
}


// SetAudioMix sets the value of the audioMix property.
// The audio mix parameters to be applied during playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/audiomix
func (p_ PlayerItem) SetAudioMix(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioMix:"), value)
}

// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/audiotimepitchalgorithm
func (p_ PlayerItem) AudioTimePitchAlgorithm() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// SetAudioTimePitchAlgorithm sets the value of the audioTimePitchAlgorithm property.
// The processing algorithm used to manage audio pitch for scaled audio edits.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/audiotimepitchalgorithm
func (p_ PlayerItem) SetAudioTimePitchAlgorithm(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/automaticallyhandlesinterstitialevents
func (p_ PlayerItem) AutomaticallyHandlesInterstitialEvents() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("automaticallyHandlesInterstitialEvents"))
	return rv
}


// SetAutomaticallyHandlesInterstitialEvents sets the value of the automaticallyHandlesInterstitialEvents property.
// A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/automaticallyhandlesinterstitialevents
func (p_ PlayerItem) SetAutomaticallyHandlesInterstitialEvents(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticallyHandlesInterstitialEvents:"), value)
}

// The array of asset keys to be automatically loaded before the player item is ready to play.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/automaticallyloadedassetkeys
func (p_ PlayerItem) AutomaticallyLoadedAssetKeys() string {
	rv := objc.Send[string](p_.ID, objc.Sel("automaticallyLoadedAssetKeys"))
	return rv
}


// SetAutomaticallyLoadedAssetKeys sets the value of the automaticallyLoadedAssetKeys property.
// The array of asset keys to be automatically loaded before the player item is ready to play.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/automaticallyloadedassetkeys
func (p_ PlayerItem) SetAutomaticallyLoadedAssetKeys(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticallyLoadedAssetKeys:"), objc.String(value))
}

// A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/automaticallypreservestimeoffsetfromlive
func (p_ PlayerItem) AutomaticallyPreservesTimeOffsetFromLive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("automaticallyPreservesTimeOffsetFromLive"))
	return rv
}


// SetAutomaticallyPreservesTimeOffsetFromLive sets the value of the automaticallyPreservesTimeOffsetFromLive property.
// A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/automaticallypreservestimeoffsetfromlive
func (p_ PlayerItem) SetAutomaticallyPreservesTimeOffsetFromLive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticallyPreservesTimeOffsetFromLive:"), value)
}

// A Boolean value that indicates whether the item can be fast forwarded.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canplayfastforward
func (p_ PlayerItem) CanPlayFastForward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlayFastForward"))
	return rv
}


// SetCanPlayFastForward sets the value of the canPlayFastForward property.
// A Boolean value that indicates whether the item can be fast forwarded.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canplayfastforward
func (p_ PlayerItem) SetCanPlayFastForward(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanPlayFastForward:"), value)
}

// A Boolean value that indicates whether the item can be quickly reversed.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canplayfastreverse
func (p_ PlayerItem) CanPlayFastReverse() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlayFastReverse"))
	return rv
}


// SetCanPlayFastReverse sets the value of the canPlayFastReverse property.
// A Boolean value that indicates whether the item can be quickly reversed.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canplayfastreverse
func (p_ PlayerItem) SetCanPlayFastReverse(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanPlayFastReverse:"), value)
}

// A Boolean value that indicates whether the item can play in reverse.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canplayreverse
func (p_ PlayerItem) CanPlayReverse() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlayReverse"))
	return rv
}


// SetCanPlayReverse sets the value of the canPlayReverse property.
// A Boolean value that indicates whether the item can play in reverse.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canplayreverse
func (p_ PlayerItem) SetCanPlayReverse(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanPlayReverse:"), value)
}

// A Boolean value that indicates whether the item can play slowly backward.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canplayslowreverse
func (p_ PlayerItem) CanPlaySlowReverse() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlaySlowReverse"))
	return rv
}


// SetCanPlaySlowReverse sets the value of the canPlaySlowReverse property.
// A Boolean value that indicates whether the item can play slowly backward.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canplayslowreverse
func (p_ PlayerItem) SetCanPlaySlowReverse(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanPlaySlowReverse:"), value)
}

// A Boolean value that indicates whether the item supports stepping backward.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canstepbackward
func (p_ PlayerItem) CanStepBackward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStepBackward"))
	return rv
}


// SetCanStepBackward sets the value of the canStepBackward property.
// A Boolean value that indicates whether the item supports stepping backward.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canstepbackward
func (p_ PlayerItem) SetCanStepBackward(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanStepBackward:"), value)
}

// A Boolean value that indicates whether the item supports stepping forward.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canstepforward
func (p_ PlayerItem) CanStepForward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStepForward"))
	return rv
}


// SetCanStepForward sets the value of the canStepForward property.
// A Boolean value that indicates whether the item supports stepping forward.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canstepforward
func (p_ PlayerItem) SetCanStepForward(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanStepForward:"), value)
}

// A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canusenetworkresourcesforlivestreamingwhilepaused
func (p_ PlayerItem) CanUseNetworkResourcesForLiveStreamingWhilePaused() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canUseNetworkResourcesForLiveStreamingWhilePaused"))
	return rv
}


// SetCanUseNetworkResourcesForLiveStreamingWhilePaused sets the value of the canUseNetworkResourcesForLiveStreamingWhilePaused property.
// A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/canusenetworkresourcesforlivestreamingwhilepaused
func (p_ PlayerItem) SetCanUseNetworkResourcesForLiveStreamingWhilePaused(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanUseNetworkResourcesForLiveStreamingWhilePaused:"), value)
}

// A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/configuredtimeoffsetfromlive
func (p_ PlayerItem) ConfiguredTimeOffsetFromLive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("configuredTimeOffsetFromLive"))
	return rv
}


// SetConfiguredTimeOffsetFromLive sets the value of the configuredTimeOffsetFromLive property.
// A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/configuredtimeoffsetfromlive
func (p_ PlayerItem) SetConfiguredTimeOffsetFromLive(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConfiguredTimeOffsetFromLive:"), value)
}

// The status of the most recent content authorization request.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/contentauthorizationrequeststatus
func (p_ PlayerItem) ContentAuthorizationRequestStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentAuthorizationRequestStatus"))
	return rv
}


// SetContentAuthorizationRequestStatus sets the value of the contentAuthorizationRequestStatus property.
// The status of the most recent content authorization request.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/contentauthorizationrequeststatus
func (p_ PlayerItem) SetContentAuthorizationRequestStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentAuthorizationRequestStatus:"), value)
}

// The custom video compositor.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/customvideocompositor
func (p_ PlayerItem) CustomVideoCompositor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("customVideoCompositor"))
	return rv
}


// SetCustomVideoCompositor sets the value of the customVideoCompositor property.
// The custom video compositor.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/customvideocompositor
func (p_ PlayerItem) SetCustomVideoCompositor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomVideoCompositor:"), value)
}

// The duration of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/duration
func (p_ PlayerItem) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// The duration of the item.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/duration
func (p_ PlayerItem) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDuration:"), value)
}

// An integrated timeline that represents the player item timing including its scheduled interstitial events.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/integratedtimeline
func (p_ PlayerItem) IntegratedTimeline() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("integratedTimeline"))
	return rv
}


// SetIntegratedTimeline sets the value of the integratedTimeline property.
// An integrated timeline that represents the player item timing including its scheduled interstitial events.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/integratedtimeline
func (p_ PlayerItem) SetIntegratedTimeline(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIntegratedTimeline:"), value)
}

// A Boolean value that indicates whether the application can be used to play the content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isapplicationauthorizedforplayback
func (p_ PlayerItem) IsApplicationAuthorizedForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isApplicationAuthorizedForPlayback"))
	return rv
}


// SetIsApplicationAuthorizedForPlayback sets the value of the isApplicationAuthorizedForPlayback property.
// A Boolean value that indicates whether the application can be used to play the content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isapplicationauthorizedforplayback
func (p_ PlayerItem) SetIsApplicationAuthorizedForPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsApplicationAuthorizedForPlayback:"), value)
}

// A Boolean value that indicates whether the player item allows spatialized audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isaudiospatializationallowed
func (p_ PlayerItem) IsAudioSpatializationAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isAudioSpatializationAllowed"))
	return rv
}


// SetIsAudioSpatializationAllowed sets the value of the isAudioSpatializationAllowed property.
// A Boolean value that indicates whether the player item allows spatialized audio playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isaudiospatializationallowed
func (p_ PlayerItem) SetIsAudioSpatializationAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsAudioSpatializationAllowed:"), value)
}

// A Boolean value that indicates whether authorization is required to play the content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isauthorizationrequiredforplayback
func (p_ PlayerItem) IsAuthorizationRequiredForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isAuthorizationRequiredForPlayback"))
	return rv
}


// SetIsAuthorizationRequiredForPlayback sets the value of the isAuthorizationRequiredForPlayback property.
// A Boolean value that indicates whether authorization is required to play the content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isauthorizationrequiredforplayback
func (p_ PlayerItem) SetIsAuthorizationRequiredForPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsAuthorizationRequiredForPlayback:"), value)
}

// A Boolean value that indicates whether the content has been authorized by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/iscontentauthorizedforplayback
func (p_ PlayerItem) IsContentAuthorizedForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isContentAuthorizedForPlayback"))
	return rv
}


// SetIsContentAuthorizedForPlayback sets the value of the isContentAuthorizedForPlayback property.
// A Boolean value that indicates whether the content has been authorized by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/iscontentauthorizedforplayback
func (p_ PlayerItem) SetIsContentAuthorizedForPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsContentAuthorizedForPlayback:"), value)
}

// A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybackbufferempty
func (p_ PlayerItem) IsPlaybackBufferEmpty() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPlaybackBufferEmpty"))
	return rv
}


// SetIsPlaybackBufferEmpty sets the value of the isPlaybackBufferEmpty property.
// A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybackbufferempty
func (p_ PlayerItem) SetIsPlaybackBufferEmpty(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPlaybackBufferEmpty:"), value)
}

// A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybackbufferfull
func (p_ PlayerItem) IsPlaybackBufferFull() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPlaybackBufferFull"))
	return rv
}


// SetIsPlaybackBufferFull sets the value of the isPlaybackBufferFull property.
// A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybackbufferfull
func (p_ PlayerItem) SetIsPlaybackBufferFull(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPlaybackBufferFull:"), value)
}

// A Boolean value that indicates whether the item will likely play through without stalling.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybacklikelytokeepup
func (p_ PlayerItem) IsPlaybackLikelyToKeepUp() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPlaybackLikelyToKeepUp"))
	return rv
}


// SetIsPlaybackLikelyToKeepUp sets the value of the isPlaybackLikelyToKeepUp property.
// A Boolean value that indicates whether the item will likely play through without stalling.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybacklikelytokeepup
func (p_ PlayerItem) SetIsPlaybackLikelyToKeepUp(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPlaybackLikelyToKeepUp:"), value)
}

// An array of time ranges indicating media data that is readily available.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/loadedtimeranges
func (p_ PlayerItem) LoadedTimeRanges() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}


// SetLoadedTimeRanges sets the value of the loadedTimeRanges property.
// An array of time ranges indicating media data that is readily available.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/loadedtimeranges
func (p_ PlayerItem) SetLoadedTimeRanges(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLoadedTimeRanges:"), value)
}

// The collection of associated media data collectors.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/mediadatacollectors
func (p_ PlayerItem) MediaDataCollectors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mediaDataCollectors"))
	return rv
}


// SetMediaDataCollectors sets the value of the mediaDataCollectors property.
// The collection of associated media data collectors.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/mediadatacollectors
func (p_ PlayerItem) SetMediaDataCollectors(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaDataCollectors:"), value)
}

// An array of outputs associated with the player item.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/outputs
func (p_ PlayerItem) Outputs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("outputs"))
	return rv
}


// SetOutputs sets the value of the outputs property.
// An array of outputs associated with the player item.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/outputs
func (p_ PlayerItem) SetOutputs(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOutputs:"), value)
}

// The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/preferredforwardbufferduration
func (p_ PlayerItem) PreferredForwardBufferDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("preferredForwardBufferDuration"))
	return rv
}


// SetPreferredForwardBufferDuration sets the value of the preferredForwardBufferDuration property.
// The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/preferredforwardbufferduration
func (p_ PlayerItem) SetPreferredForwardBufferDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredForwardBufferDuration:"), value)
}

// The desired maximum resolution of a video that is to be downloaded.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/preferredmaximumresolution
func (p_ PlayerItem) PreferredMaximumResolution() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("preferredMaximumResolution"))
	return rv
}


// SetPreferredMaximumResolution sets the value of the preferredMaximumResolution property.
// The desired maximum resolution of a video that is to be downloaded.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/preferredmaximumresolution
func (p_ PlayerItem) SetPreferredMaximumResolution(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredMaximumResolution:"), value)
}

// An upper limit on the resolution of video to download when connecting over expensive networks.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/preferredmaximumresolutionforexpensivenetworks
func (p_ PlayerItem) PreferredMaximumResolutionForExpensiveNetworks() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("preferredMaximumResolutionForExpensiveNetworks"))
	return rv
}


// SetPreferredMaximumResolutionForExpensiveNetworks sets the value of the preferredMaximumResolutionForExpensiveNetworks property.
// An upper limit on the resolution of video to download when connecting over expensive networks.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/preferredmaximumresolutionforexpensivenetworks
func (p_ PlayerItem) SetPreferredMaximumResolutionForExpensiveNetworks(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredMaximumResolutionForExpensiveNetworks:"), value)
}

// A limit of network bandwidth consumption by the item when connecting over expensive networks.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/preferredpeakbitrateforexpensivenetworks
func (p_ PlayerItem) PreferredPeakBitRateForExpensiveNetworks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("preferredPeakBitRateForExpensiveNetworks"))
	return rv
}


// SetPreferredPeakBitRateForExpensiveNetworks sets the value of the preferredPeakBitRateForExpensiveNetworks property.
// A limit of network bandwidth consumption by the item when connecting over expensive networks.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/preferredpeakbitrateforexpensivenetworks
func (p_ PlayerItem) SetPreferredPeakBitRateForExpensiveNetworks(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredPeakBitRateForExpensiveNetworks:"), value)
}

// The size at which the visual portion of the item is presented by the player.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/presentationsize
func (p_ PlayerItem) PresentationSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("presentationSize"))
	return rv
}


// SetPresentationSize sets the value of the presentationSize property.
// The size at which the visual portion of the item is presented by the player.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/presentationsize
func (p_ PlayerItem) SetPresentationSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPresentationSize:"), value)
}

// A recommended time offset from the live time based on observed network conditions.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/recommendedtimeoffsetfromlive
func (p_ PlayerItem) RecommendedTimeOffsetFromLive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("recommendedTimeOffsetFromLive"))
	return rv
}


// SetRecommendedTimeOffsetFromLive sets the value of the recommendedTimeOffsetFromLive property.
// A recommended time offset from the live time based on observed network conditions.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/recommendedtimeoffsetfromlive
func (p_ PlayerItem) SetRecommendedTimeOffsetFromLive(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRecommendedTimeOffsetFromLive:"), value)
}

// An array of time ranges within which it is possible to seek.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/seekabletimeranges
func (p_ PlayerItem) SeekableTimeRanges() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("seekableTimeRanges"))
	return rv
}


// SetSeekableTimeRanges sets the value of the seekableTimeRanges property.
// An array of time ranges within which it is possible to seek.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/seekabletimeranges
func (p_ PlayerItem) SetSeekableTimeRanges(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSeekableTimeRanges:"), value)
}

// A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/seekingwaitsforvideocompositionrendering
func (p_ PlayerItem) SeekingWaitsForVideoCompositionRendering() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("seekingWaitsForVideoCompositionRendering"))
	return rv
}


// SetSeekingWaitsForVideoCompositionRendering sets the value of the seekingWaitsForVideoCompositionRendering property.
// A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/seekingwaitsforvideocompositionrendering
func (p_ PlayerItem) SetSeekingWaitsForVideoCompositionRendering(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSeekingWaitsForVideoCompositionRendering:"), value)
}

// A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/startsonfirsteligiblevariant
func (p_ PlayerItem) StartsOnFirstEligibleVariant() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("startsOnFirstEligibleVariant"))
	return rv
}


// SetStartsOnFirstEligibleVariant sets the value of the startsOnFirstEligibleVariant property.
// A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/startsonfirsteligiblevariant
func (p_ PlayerItem) SetStartsOnFirstEligibleVariant(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartsOnFirstEligibleVariant:"), value)
}

// The template player item that initializes this instance.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/template
func (p_ PlayerItem) Template() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("template"))
	return rv
}


// SetTemplate sets the value of the template property.
// The template player item that initializes this instance.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/template
func (p_ PlayerItem) SetTemplate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTemplate:"), value)
}

// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/textstylerules
func (p_ PlayerItem) TextStyleRules() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("textStyleRules"))
	return rv
}


// SetTextStyleRules sets the value of the textStyleRules property.
// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/textstylerules
func (p_ PlayerItem) SetTextStyleRules(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTextStyleRules:"), value)
}

// The timebase information for the item.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/timebase
func (p_ PlayerItem) Timebase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("timebase"))
	return rv
}


// SetTimebase sets the value of the timebase property.
// The timebase information for the item.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/timebase
func (p_ PlayerItem) SetTimebase(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimebase:"), value)
}

// An array of player item track objects.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/tracks
func (p_ PlayerItem) Tracks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("tracks"))
	return rv
}


// SetTracks sets the value of the tracks property.
// An array of player item track objects.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/tracks
func (p_ PlayerItem) SetTracks(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTracks:"), value)
}

// The preferences the player item uses when selecting variant playlists.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/variantpreferences
func (p_ PlayerItem) VariantPreferences() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("variantPreferences"))
	return rv
}


// SetVariantPreferences sets the value of the variantPreferences property.
// The preferences the player item uses when selecting variant playlists.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/variantpreferences
func (p_ PlayerItem) SetVariantPreferences(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVariantPreferences:"), value)
}

// The video aperture mode to apply during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/videoaperturemode
func (p_ PlayerItem) VideoApertureMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("videoApertureMode"))
	return rv
}


// SetVideoApertureMode sets the value of the videoApertureMode property.
// The video aperture mode to apply during playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/videoaperturemode
func (p_ PlayerItem) SetVideoApertureMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoApertureMode:"), value)
}

// The video composition settings to be applied during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/videocomposition
func (p_ PlayerItem) VideoComposition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("videoComposition"))
	return rv
}


// SetVideoComposition sets the value of the videoComposition property.
// The video composition settings to be applied during playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/videocomposition
func (p_ PlayerItem) SetVideoComposition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoComposition:"), value)
}



