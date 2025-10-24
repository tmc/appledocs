// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AppliesPerFrameHDRDisplayMetadata() bool
	SetAppliesPerFrameHDRDisplayMetadata(value bool)
	Asset() IAVAsset
	AudioMix() IAVAudioMix
	SetAudioMix(value IAVAudioMix)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* not a class type */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* not a class type */)
	AutomaticallyHandlesInterstitialEvents() bool
	SetAutomaticallyHandlesInterstitialEvents(value bool)
	AutomaticallyLoadedAssetKeys() []string
	AutomaticallyPreservesTimeOffsetFromLive() bool
	SetAutomaticallyPreservesTimeOffsetFromLive(value bool)
	CanPlayFastForward() bool
	CanPlayFastReverse() bool
	CanPlayReverse() bool
	CanPlaySlowForward() bool
	CanPlaySlowReverse() bool
	CanStepBackward() bool
	CanStepForward() bool
	CanUseNetworkResourcesForLiveStreamingWhilePaused() bool
	SetCanUseNetworkResourcesForLiveStreamingWhilePaused(value bool)
	ConfiguredTimeOffsetFromLive() objc.IObject /* cross-framework: Time */
	SetConfiguredTimeOffsetFromLive(value objc.IObject /* cross-framework: Time */)
	ContentAuthorizationRequestStatus() ContentAuthorizationStatus
	CurrentMediaSelection() IAVMediaSelection
	CustomVideoCompositor() objc.ID
	Duration() objc.IObject /* cross-framework: Time */
	Error() Error
	ForwardPlaybackEndTime() objc.IObject /* cross-framework: Time */
	SetForwardPlaybackEndTime(value objc.IObject /* cross-framework: Time */)
	IntegratedTimeline() IAVPlayerItemIntegratedTimeline
	ApplicationAuthorizedForPlayback() bool
	AuthorizationRequiredForPlayback() bool
	ContentAuthorizedForPlayback() bool
	PlaybackBufferEmpty() bool
	PlaybackBufferFull() bool
	PlaybackLikelyToKeepUp() bool
	LoadedTimeRanges() []objc.IObject /* cross-framework: Value */
	MediaDataCollectors() []PlayerItemMediaDataCollector /* not a class type */
	NavigationMarkerGroups() []objc.IObject /* cross-framework: NavigationMarkersGroup */
	SetNavigationMarkerGroups(value []objc.IObject /* cross-framework: NavigationMarkersGroup */)
	Outputs() []objc.IObject /* cross-framework: PlayerItemOutput */
	PreferredCustomMediaSelectionSchemes() []ICustomMediaSelectionScheme
	SetPreferredCustomMediaSelectionSchemes(value []ICustomMediaSelectionScheme)
	PreferredForwardBufferDuration() float64
	SetPreferredForwardBufferDuration(value float64)
	PreferredMaximumResolution() objc.IObject /* cross-framework: Size */
	SetPreferredMaximumResolution(value objc.IObject /* cross-framework: Size */)
	PreferredMaximumResolutionForExpensiveNetworks() objc.IObject /* cross-framework: Size */
	SetPreferredMaximumResolutionForExpensiveNetworks(value objc.IObject /* cross-framework: Size */)
	PreferredPeakBitRate() float64
	SetPreferredPeakBitRate(value float64)
	PreferredPeakBitRateForExpensiveNetworks() float64
	SetPreferredPeakBitRateForExpensiveNetworks(value float64)
	PresentationSize() objc.IObject /* cross-framework: Size */
	RecommendedTimeOffsetFromLive() objc.IObject /* cross-framework: Time */
	ReversePlaybackEndTime() objc.IObject /* cross-framework: Time */
	SetReversePlaybackEndTime(value objc.IObject /* cross-framework: Time */)
	SeekableTimeRanges() []objc.IObject /* cross-framework: Value */
	SeekingWaitsForVideoCompositionRendering() bool
	SetSeekingWaitsForVideoCompositionRendering(value bool)
	StartsOnFirstEligibleVariant() bool
	SetStartsOnFirstEligibleVariant(value bool)
	Status() PlayerItemStatus
	TemplatePlayerItem() IAVPlayerItem
	TextStyleRules() []ITextStyleRule
	SetTextStyleRules(value []ITextStyleRule)
	Timebase() TimebaseRef /* not a class type */
	Tracks() []IPlayerItemTrack
	VariantPreferences() VariantPreferences
	SetVariantPreferences(value VariantPreferences)
	VideoApertureMode() objc.IObject /* cross-framework: VideoApertureMode */
	SetVideoApertureMode(value objc.IObject /* cross-framework: VideoApertureMode */)
	VideoComposition() objc.IObject /* cross-framework: VideoComposition */
	SetVideoComposition(value objc.IObject /* cross-framework: VideoComposition */)
	AllowedAudioSpatializationFormats() AudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats)
	IsApplicationAuthorizedForPlayback() bool
	SetIsApplicationAuthorizedForPlayback(value bool)
	IsAudioSpatializationAllowed() bool
	SetIsAudioSpatializationAllowed(value bool)
	IsAuthorizationRequiredForPlayback() bool
	SetIsAuthorizationRequiredForPlayback(value bool)
	IsContentAuthorizedForPlayback() bool
	SetIsContentAuthorizedForPlayback(value bool)
	IsPlaybackBufferEmpty() bool
	SetIsPlaybackBufferEmpty(value bool)
	IsPlaybackBufferFull() bool
	SetIsPlaybackBufferFull(value bool)
	IsPlaybackLikelyToKeepUp() bool
	SetIsPlaybackLikelyToKeepUp(value bool)
	Template() IAVPlayerItem
	SetTemplate(value IAVPlayerItem)
	// methods:
	AccessLog() IPlayerItemAccessLog
	AddOutput(output objc.IObject /* cross-framework: PlayerItemOutput */)
	AddMediaDataCollector(collector PlayerItemMediaDataCollector /* not a class type */)
	CancelContentAuthorizationRequest()
	CancelPendingSeeks()
	CurrentDate() objc.IObject /* cross-framework: Date */
	CurrentTime() objc.IObject /* cross-framework: Time */
	EffectiveMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.IDictionary
	ErrorLog() IPlayerItemErrorLog
	RemoveMediaDataCollector(collector PlayerItemMediaDataCollector /* not a class type */)
	RemoveOutput(output objc.IObject /* cross-framework: PlayerItemOutput */)
	RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler(timeoutInterval float64, handler unsafe.Pointer)
	SeekToDateCompletionHandler(date objc.IObject /* cross-framework: NSDate */, completionHandler unsafe.Pointer) bool
	SeekToTimeCompletionHandler(time objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
	SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objc.IObject /* cross-framework: Time */, toleranceBefore objc.IObject /* cross-framework: Time */, toleranceAfter objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
	SelectMediaPresentationSettingForMediaSelectionGroup(mediaPresentationSetting IAVMediaPresentationSetting, mediaSelectionGroup IAVMediaSelectionGroup)
	SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IAVMediaSelectionOption, mediaSelectionGroup IAVMediaSelectionGroup)
	SelectMediaOptionAutomaticallyInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup)
	SelectMediaPresentationLanguageForMediaSelectionGroup(language objc.IObject /* cross-framework: NSString */, mediaSelectionGroup IAVMediaSelectionGroup)
	SelectedMediaPresentationLanguageForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) objc.IObject /* cross-framework: String */
	SelectedMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.IDictionary
	StepByCount(stepCount int)
}

// An object that models the timing and presentation state of an asset during playback.
//
// A player item stores a reference to an object, which represents the media to play. If you require inspecting an asset before you enqueue it for playback, call its method to retrieve the values of one or more properties. Alternatively, you can tell the player item to automatically load the required properties by passing them to its initializer. When the player item is ready to play, the asset properties you request are ready to use.


// An object that models the timing and presentation state of an asset during playback.
//
// [Full Topic]
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



// Creates a player item for a specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(asset:)-87rjl
func NewPlayerItemWithAsset(asset IAVAsset) PlayerItem {
	instance := getPlayerItemClass().Alloc()
	rv := objc.Send[PlayerItem](instance.ID, objc.Sel("initWithAsset:"), asset)
	rv.Autorelease()
	return rv
}


// Creates a player item with the specified asset and the asset keys to automatically load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(asset:automaticallyLoadedAssetKeys:)-8x4
func NewPlayerItemWithAssetAutomaticallyLoadedAssetKeys(asset IAVAsset, automaticallyLoadedAssetKeys []string) PlayerItem {
	instance := getPlayerItemClass().Alloc()
	rv := objc.Send[PlayerItem](instance.ID, objc.Sel("initWithAsset:automaticallyLoadedAssetKeys:"), asset, automaticallyLoadedAssetKeys)
	rv.Autorelease()
	return rv
}


// Creates a player item with a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(url:)
func NewPlayerItemWithURL(URL objc.IObject /* cross-framework: NSURL */) PlayerItem {
	instance := getPlayerItemClass().Alloc()
	rv := objc.Send[PlayerItem](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}



// Returns a new player item for a specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/playerItemWithAsset:
func (pc _PlayerItemClass) PlayerItemWithAsset(asset IAVAsset) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("playerItemWithAsset:"), asset)
	return rv
}


// Creates a player item with the specified asset and the asset keys to automatically load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/playerItemWithAsset:automaticallyLoadedAssetKeys:
func (pc _PlayerItemClass) PlayerItemWithAssetAutomaticallyLoadedAssetKeys(asset IAVAsset, automaticallyLoadedAssetKeys []string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("playerItemWithAsset:automaticallyLoadedAssetKeys:"), asset, automaticallyLoadedAssetKeys)
	return rv
}


// Returns a new player item with a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/playerItemWithURL:
func (pc _PlayerItemClass) PlayerItemWithURL(URL objc.IObject /* cross-framework: NSURL */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("playerItemWithURL:"), URL)
	return rv
}


// Returns an object that represents a snapshot of the network access log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/accessLog()
func (p_ PlayerItem) AccessLog() IPlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](p_.ID, objc.Sel("accessLog"))
	return rv
}


// Adds the specified player item output object to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/add(_:)-16ctk
func (p_ PlayerItem) AddOutput(output objc.IObject /* cross-framework: PlayerItemOutput */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addOutput:"), output)
}


// Adds the specified media data collector to the player item’s collection of media collectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/add(_:)-9l3to
func (p_ PlayerItem) AddMediaDataCollector(collector PlayerItemMediaDataCollector /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addMediaDataCollector:"), collector)
}


// Cancels the currently outstanding content authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/cancelContentAuthorizationRequest()
func (p_ PlayerItem) CancelContentAuthorizationRequest() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelContentAuthorizationRequest"))
}


// Cancels any pending seek requests and invokes the corresponding completion handlers if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/cancelPendingSeeks()
func (p_ PlayerItem) CancelPendingSeeks() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelPendingSeeks"))
}


// Returns the current time of the item as a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentDate()
func (p_ PlayerItem) CurrentDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("currentDate"))
	return rv
}


// Returns the current time of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentTime()
func (p_ PlayerItem) CurrentTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](p_.ID, objc.Sel("currentTime"))
	return rv
}


// Indicates the media presentation settings with media characteristics that are possessed by the currently selected AVMediaSelectionOption in the specified AVMediaSelectionGroup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/effectiveMediaPresentationSettings(for:)
func (p_ PlayerItem) EffectiveMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("effectiveMediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}


// Returns an object that represents a snapshot of the error log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/errorLog()
func (p_ PlayerItem) ErrorLog() IPlayerItemErrorLog {
	rv := objc.Send[PlayerItemErrorLog](p_.ID, objc.Sel("errorLog"))
	return rv
}


// Removes the specified media data collector from the player item’s collection of media collectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/remove(_:)-29iuz
func (p_ PlayerItem) RemoveMediaDataCollector(collector PlayerItemMediaDataCollector /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeMediaDataCollector:"), collector)
}


// Removes the specified player item output object from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/remove(_:)-46b1r
func (p_ PlayerItem) RemoveOutput(output objc.IObject /* cross-framework: PlayerItemOutput */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeOutput:"), output)
}


// Presents the user the opportunity to authorize the content for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/requestContentAuthorizationAsynchronously(withTimeoutInterval:completionHandler:)
func (p_ PlayerItem) RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler(timeoutInterval float64, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("requestContentAuthorizationAsynchronouslyWithTimeoutInterval:completionHandler:"), timeoutInterval, handler)
}


// Sets the current playback time to the time specified by the date object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seek(to:completionHandler:)-1dibq
func (p_ PlayerItem) SeekToDateCompletionHandler(date objc.IObject /* cross-framework: NSDate */, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("seekToDate:completionHandler:"), date, completionHandler)
	return rv
}


// Sets the current playback time to the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seek(to:completionHandler:)-91gnw
func (p_ PlayerItem) SeekToTimeCompletionHandler(time objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:completionHandler:"), time, completionHandler)
}


// Sets the current playback time within a specified time bound and invokes the specified block when the seek operation completes or is interrupted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seek(to:toleranceBefore:toleranceAfter:completionHandler:)
func (p_ PlayerItem) SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objc.IObject /* cross-framework: Time */, toleranceBefore objc.IObject /* cross-framework: Time */, toleranceAfter objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:completionHandler:"), time, toleranceBefore, toleranceAfter, completionHandler)
}


// When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular presentation setting, replacing any previous preference for settings of the same media presentation selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/select(_:for:)
func (p_ PlayerItem) SelectMediaPresentationSettingForMediaSelectionGroup(mediaPresentationSetting IAVMediaPresentationSetting, mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectMediaPresentationSetting:forMediaSelectionGroup:"), mediaPresentationSetting, mediaSelectionGroup)
}


// Selects a media option in a given media selection group and deselects all other options in that group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/select(_:in:)
func (p_ PlayerItem) SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IAVMediaSelectionOption, mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectMediaOption:inMediaSelectionGroup:"), mediaSelectionOption, mediaSelectionGroup)
}


// Selects the media option in the specified media selection group that best matches the receiver’s automatic selection criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectMediaOptionAutomatically(in:)
func (p_ PlayerItem) SelectMediaOptionAutomaticallyInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectMediaOptionAutomaticallyInMediaSelectionGroup:"), mediaSelectionGroup)
}


// When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular language, replacing any previous preference for available languages of the specified group’s custom media selection scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectMediaPresentationLanguage(_:for:)
func (p_ PlayerItem) SelectMediaPresentationLanguageForMediaSelectionGroup(language objc.IObject /* cross-framework: NSString */, mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectMediaPresentationLanguage:forMediaSelectionGroup:"), language, mediaSelectionGroup)
}


// Returns the selected media presentation language for the specified media selection group, if any language has previously been selected via use of -selectMediaPresentationLanguages:forMediaSelectionGroup:.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectedMediaPresentationLanguage(for:)
func (p_ PlayerItem) SelectedMediaPresentationLanguageForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("selectedMediaPresentationLanguageForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}


// Indicates the media presentation settings that have most recently been selected for each AVMediaPresentationSelector of the AVCustomMediaSelectionScheme of the specified AVMediaSelectionGroup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectedMediaPresentationSettings(for:)
func (p_ PlayerItem) SelectedMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("selectedMediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}


// Moves the player item’s current time forward or backward by a specified number of steps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/step(byCount:)
func (p_ PlayerItem) StepByCount(stepCount int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("stepByCount:"), stepCount)
}


// A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/appliesPerFrameHDRDisplayMetadata
func (p_ PlayerItem) AppliesPerFrameHDRDisplayMetadata() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("appliesPerFrameHDRDisplayMetadata"))
	return rv
}


// A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/appliesPerFrameHDRDisplayMetadata
func (p_ PlayerItem) SetAppliesPerFrameHDRDisplayMetadata(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppliesPerFrameHDRDisplayMetadata:"), value)
}


// The asset provided during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/asset
func (p_ PlayerItem) Asset() IAVAsset {
	rv := objc.Send[Asset](p_.ID, objc.Sel("asset"))
	return rv
}


// The audio mix parameters to be applied during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/audioMix
func (p_ PlayerItem) AudioMix() IAVAudioMix {
	rv := objc.Send[AudioMix](p_.ID, objc.Sel("audioMix"))
	return rv
}


// The audio mix parameters to be applied during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/audioMix
func (p_ PlayerItem) SetAudioMix(value IAVAudioMix) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioMix:"), value)
}


// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/audioTimePitchAlgorithm
func (p_ PlayerItem) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* not a class type */ {
	rv := objc.Send[AudioTimePitchAlgorithm](p_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/audioTimePitchAlgorithm
func (p_ PlayerItem) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}


// A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/automaticallyHandlesInterstitialEvents
func (p_ PlayerItem) AutomaticallyHandlesInterstitialEvents() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("automaticallyHandlesInterstitialEvents"))
	return rv
}


// A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/automaticallyHandlesInterstitialEvents
func (p_ PlayerItem) SetAutomaticallyHandlesInterstitialEvents(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticallyHandlesInterstitialEvents:"), value)
}


// The array of asset keys to be automatically loaded before the player item is ready to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/automaticallyLoadedAssetKeys
func (p_ PlayerItem) AutomaticallyLoadedAssetKeys() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("automaticallyLoadedAssetKeys"))
	return rv
}


// A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/automaticallyPreservesTimeOffsetFromLive
func (p_ PlayerItem) AutomaticallyPreservesTimeOffsetFromLive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("automaticallyPreservesTimeOffsetFromLive"))
	return rv
}


// A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/automaticallyPreservesTimeOffsetFromLive
func (p_ PlayerItem) SetAutomaticallyPreservesTimeOffsetFromLive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticallyPreservesTimeOffsetFromLive:"), value)
}


// A Boolean value that indicates whether the item can be fast forwarded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlayFastForward
func (p_ PlayerItem) CanPlayFastForward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlayFastForward"))
	return rv
}


// A Boolean value that indicates whether the item can be quickly reversed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlayFastReverse
func (p_ PlayerItem) CanPlayFastReverse() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlayFastReverse"))
	return rv
}


// A Boolean value that indicates whether the item can play in reverse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlayReverse
func (p_ PlayerItem) CanPlayReverse() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlayReverse"))
	return rv
}


// A Boolean value that indicates whether the item can play slower than normal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlaySlowForward
func (p_ PlayerItem) CanPlaySlowForward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlaySlowForward"))
	return rv
}


// A Boolean value that indicates whether the item can play slowly backward.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlaySlowReverse
func (p_ PlayerItem) CanPlaySlowReverse() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPlaySlowReverse"))
	return rv
}


// A Boolean value that indicates whether the item supports stepping backward.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canStepBackward
func (p_ PlayerItem) CanStepBackward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStepBackward"))
	return rv
}


// A Boolean value that indicates whether the item supports stepping forward.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canStepForward
func (p_ PlayerItem) CanStepForward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStepForward"))
	return rv
}


// A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canUseNetworkResourcesForLiveStreamingWhilePaused
func (p_ PlayerItem) CanUseNetworkResourcesForLiveStreamingWhilePaused() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canUseNetworkResourcesForLiveStreamingWhilePaused"))
	return rv
}


// A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canUseNetworkResourcesForLiveStreamingWhilePaused
func (p_ PlayerItem) SetCanUseNetworkResourcesForLiveStreamingWhilePaused(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanUseNetworkResourcesForLiveStreamingWhilePaused:"), value)
}


// A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/configuredTimeOffsetFromLive
func (p_ PlayerItem) ConfiguredTimeOffsetFromLive() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](p_.ID, objc.Sel("configuredTimeOffsetFromLive"))
	return rv
}


// A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/configuredTimeOffsetFromLive
func (p_ PlayerItem) SetConfiguredTimeOffsetFromLive(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConfiguredTimeOffsetFromLive:"), value)
}


// The status of the most recent content authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/contentAuthorizationRequestStatus
func (p_ PlayerItem) ContentAuthorizationRequestStatus() ContentAuthorizationStatus {
	rv := objc.Send[ContentAuthorizationStatus](p_.ID, objc.Sel("contentAuthorizationRequestStatus"))
	return rv
}


// The current media selections for each of the receiver’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentMediaSelection
func (p_ PlayerItem) CurrentMediaSelection() IAVMediaSelection {
	rv := objc.Send[MediaSelection](p_.ID, objc.Sel("currentMediaSelection"))
	return rv
}


// The custom video compositor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/customVideoCompositor
func (p_ PlayerItem) CustomVideoCompositor() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("customVideoCompositor"))
	return rv
}


// The duration of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/duration
func (p_ PlayerItem) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](p_.ID, objc.Sel("duration"))
	return rv
}


// The error that caused the player item to fail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/error
func (p_ PlayerItem) Error() Error {
	rv := objc.Send[Error](p_.ID, objc.Sel("error"))
	return rv
}


// The time at which forward playback ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/forwardPlaybackEndTime
func (p_ PlayerItem) ForwardPlaybackEndTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](p_.ID, objc.Sel("forwardPlaybackEndTime"))
	return rv
}


// The time at which forward playback ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/forwardPlaybackEndTime
func (p_ PlayerItem) SetForwardPlaybackEndTime(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setForwardPlaybackEndTime:"), value)
}


// An integrated timeline that represents the player item timing including its scheduled interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/integratedTimeline
func (p_ PlayerItem) IntegratedTimeline() IAVPlayerItemIntegratedTimeline {
	rv := objc.Send[PlayerItemIntegratedTimeline](p_.ID, objc.Sel("integratedTimeline"))
	return rv
}


// A Boolean value that indicates whether the application can be used to play the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isApplicationAuthorizedForPlayback
func (p_ PlayerItem) ApplicationAuthorizedForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("applicationAuthorizedForPlayback"))
	return rv
}


// A Boolean value that indicates whether authorization is required to play the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isAuthorizationRequiredForPlayback
func (p_ PlayerItem) AuthorizationRequiredForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("authorizationRequiredForPlayback"))
	return rv
}


// A Boolean value that indicates whether the content has been authorized by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isContentAuthorizedForPlayback
func (p_ PlayerItem) ContentAuthorizedForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("contentAuthorizedForPlayback"))
	return rv
}


// A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isPlaybackBufferEmpty
func (p_ PlayerItem) PlaybackBufferEmpty() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("playbackBufferEmpty"))
	return rv
}


// A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isPlaybackBufferFull
func (p_ PlayerItem) PlaybackBufferFull() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("playbackBufferFull"))
	return rv
}


// A Boolean value that indicates whether the item will likely play through without stalling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isPlaybackLikelyToKeepUp
func (p_ PlayerItem) PlaybackLikelyToKeepUp() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("playbackLikelyToKeepUp"))
	return rv
}


// An array of time ranges indicating media data that is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/loadedTimeRanges
func (p_ PlayerItem) LoadedTimeRanges() []objc.IObject /* cross-framework: Value */ {
	rv := objc.Send[[]foundation.Value](p_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}


// The collection of associated media data collectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/mediaDataCollectors
func (p_ PlayerItem) MediaDataCollectors() []PlayerItemMediaDataCollector /* not a class type */ {
	rv := objc.Send[[]PlayerItemMediaDataCollector](p_.ID, objc.Sel("mediaDataCollectors"))
	return rv
}


// The time marker groups that provide ways to navigate the player item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/navigationMarkerGroups
func (p_ PlayerItem) NavigationMarkerGroups() []objc.IObject /* cross-framework: NavigationMarkersGroup */ {
	rv := objc.Send[[]avkit.NavigationMarkersGroup](p_.ID, objc.Sel("navigationMarkerGroups"))
	return rv
}


// The time marker groups that provide ways to navigate the player item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/navigationMarkerGroups
func (p_ PlayerItem) SetNavigationMarkerGroups(value []objc.IObject /* cross-framework: NavigationMarkersGroup */) {
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


// An array of outputs associated with the player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/outputs
func (p_ PlayerItem) Outputs() []objc.IObject /* cross-framework: PlayerItemOutput */ {
	rv := objc.Send[[]PlayerItemOutput](p_.ID, objc.Sel("outputs"))
	return rv
}


// Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of the receiver’s asset with which an associated UI implementation should configure its interface for media selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredCustomMediaSelectionSchemes
func (p_ PlayerItem) PreferredCustomMediaSelectionSchemes() []ICustomMediaSelectionScheme {
	rv := objc.Send[[]CustomMediaSelectionScheme](p_.ID, objc.Sel("preferredCustomMediaSelectionSchemes"))
	return rv
}


// Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of the receiver’s asset with which an associated UI implementation should configure its interface for media selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredCustomMediaSelectionSchemes
func (p_ PlayerItem) SetPreferredCustomMediaSelectionSchemes(value []ICustomMediaSelectionScheme) {
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


// The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredForwardBufferDuration
func (p_ PlayerItem) PreferredForwardBufferDuration() float64 {
	rv := objc.Send[TimeInterval](p_.ID, objc.Sel("preferredForwardBufferDuration"))
	return rv
}


// The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredForwardBufferDuration
func (p_ PlayerItem) SetPreferredForwardBufferDuration(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredForwardBufferDuration:"), value)
}


// The desired maximum resolution of a video that is to be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredMaximumResolution
func (p_ PlayerItem) PreferredMaximumResolution() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](p_.ID, objc.Sel("preferredMaximumResolution"))
	return rv
}


// The desired maximum resolution of a video that is to be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredMaximumResolution
func (p_ PlayerItem) SetPreferredMaximumResolution(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredMaximumResolution:"), value)
}


// An upper limit on the resolution of video to download when connecting over expensive networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredMaximumResolutionForExpensiveNetworks
func (p_ PlayerItem) PreferredMaximumResolutionForExpensiveNetworks() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](p_.ID, objc.Sel("preferredMaximumResolutionForExpensiveNetworks"))
	return rv
}


// An upper limit on the resolution of video to download when connecting over expensive networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredMaximumResolutionForExpensiveNetworks
func (p_ PlayerItem) SetPreferredMaximumResolutionForExpensiveNetworks(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredMaximumResolutionForExpensiveNetworks:"), value)
}


// The desired limit, in bits per second, of network bandwidth consumption for this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredPeakBitRate
func (p_ PlayerItem) PreferredPeakBitRate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("preferredPeakBitRate"))
	return rv
}


// The desired limit, in bits per second, of network bandwidth consumption for this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredPeakBitRate
func (p_ PlayerItem) SetPreferredPeakBitRate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredPeakBitRate:"), value)
}


// A limit of network bandwidth consumption by the item when connecting over expensive networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredPeakBitRateForExpensiveNetworks
func (p_ PlayerItem) PreferredPeakBitRateForExpensiveNetworks() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("preferredPeakBitRateForExpensiveNetworks"))
	return rv
}


// A limit of network bandwidth consumption by the item when connecting over expensive networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredPeakBitRateForExpensiveNetworks
func (p_ PlayerItem) SetPreferredPeakBitRateForExpensiveNetworks(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredPeakBitRateForExpensiveNetworks:"), value)
}


// The size at which the visual portion of the item is presented by the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/presentationSize
func (p_ PlayerItem) PresentationSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](p_.ID, objc.Sel("presentationSize"))
	return rv
}


// A recommended time offset from the live time based on observed network conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/recommendedTimeOffsetFromLive
func (p_ PlayerItem) RecommendedTimeOffsetFromLive() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](p_.ID, objc.Sel("recommendedTimeOffsetFromLive"))
	return rv
}


// The time at which reverse playback ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/reversePlaybackEndTime
func (p_ PlayerItem) ReversePlaybackEndTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](p_.ID, objc.Sel("reversePlaybackEndTime"))
	return rv
}


// The time at which reverse playback ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/reversePlaybackEndTime
func (p_ PlayerItem) SetReversePlaybackEndTime(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReversePlaybackEndTime:"), value)
}


// An array of time ranges within which it is possible to seek.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seekableTimeRanges
func (p_ PlayerItem) SeekableTimeRanges() []objc.IObject /* cross-framework: Value */ {
	rv := objc.Send[[]foundation.Value](p_.ID, objc.Sel("seekableTimeRanges"))
	return rv
}


// A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seekingWaitsForVideoCompositionRendering
func (p_ PlayerItem) SeekingWaitsForVideoCompositionRendering() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("seekingWaitsForVideoCompositionRendering"))
	return rv
}


// A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seekingWaitsForVideoCompositionRendering
func (p_ PlayerItem) SetSeekingWaitsForVideoCompositionRendering(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSeekingWaitsForVideoCompositionRendering:"), value)
}


// A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/startsOnFirstEligibleVariant
func (p_ PlayerItem) StartsOnFirstEligibleVariant() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("startsOnFirstEligibleVariant"))
	return rv
}


// A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/startsOnFirstEligibleVariant
func (p_ PlayerItem) SetStartsOnFirstEligibleVariant(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartsOnFirstEligibleVariant:"), value)
}


// The status of the player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/status-swift.property
func (p_ PlayerItem) Status() PlayerItemStatus {
	rv := objc.Send[PlayerItemStatus](p_.ID, objc.Sel("status"))
	return rv
}


// The template player item that initializes this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/template
func (p_ PlayerItem) TemplatePlayerItem() IAVPlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("templatePlayerItem"))
	return rv
}


// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/textStyleRules
func (p_ PlayerItem) TextStyleRules() []ITextStyleRule {
	rv := objc.Send[[]TextStyleRule](p_.ID, objc.Sel("textStyleRules"))
	return rv
}


// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/textStyleRules
func (p_ PlayerItem) SetTextStyleRules(value []ITextStyleRule) {
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
	objc.Send[objc.ID](p_.ID, objc.Sel("setTextStyleRules:"), nsArray)
}


// The timebase information for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/timebase
func (p_ PlayerItem) Timebase() TimebaseRef /* not a class type */ {
	rv := objc.Send[TimebaseRef](p_.ID, objc.Sel("timebase"))
	return rv
}


// An array of player item track objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/tracks
func (p_ PlayerItem) Tracks() []IPlayerItemTrack {
	rv := objc.Send[[]PlayerItemTrack](p_.ID, objc.Sel("tracks"))
	return rv
}


// The preferences the player item uses when selecting variant playlists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/variantPreferences
func (p_ PlayerItem) VariantPreferences() VariantPreferences {
	rv := objc.Send[VariantPreferences](p_.ID, objc.Sel("variantPreferences"))
	return rv
}


// The preferences the player item uses when selecting variant playlists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/variantPreferences
func (p_ PlayerItem) SetVariantPreferences(value VariantPreferences) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVariantPreferences:"), value)
}


// The video aperture mode to apply during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/videoApertureMode
func (p_ PlayerItem) VideoApertureMode() objc.IObject /* cross-framework: VideoApertureMode */ {
	rv := objc.Send[VideoApertureMode](p_.ID, objc.Sel("videoApertureMode"))
	return rv
}


// The video aperture mode to apply during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/videoApertureMode
func (p_ PlayerItem) SetVideoApertureMode(value objc.IObject /* cross-framework: VideoApertureMode */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoApertureMode:"), value)
}


// The video composition settings to be applied during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/videoComposition
func (p_ PlayerItem) VideoComposition() objc.IObject /* cross-framework: VideoComposition */ {
	rv := objc.Send[VideoComposition](p_.ID, objc.Sel("videoComposition"))
	return rv
}


// The video composition settings to be applied during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/videoComposition
func (p_ PlayerItem) SetVideoComposition(value objc.IObject /* cross-framework: VideoComposition */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoComposition:"), value)
}


// The source audio channel layouts the player item supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/allowedaudiospatializationformats
func (p_ PlayerItem) AllowedAudioSpatializationFormats() AudioSpatializationFormats {
	rv := objc.Send[AudioSpatializationFormats](p_.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}


// The source audio channel layouts the player item supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/allowedaudiospatializationformats
func (p_ PlayerItem) SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}


// A Boolean value that indicates whether the application can be used to play the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isapplicationauthorizedforplayback
func (p_ PlayerItem) IsApplicationAuthorizedForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isApplicationAuthorizedForPlayback"))
	return rv
}


// A Boolean value that indicates whether the application can be used to play the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isapplicationauthorizedforplayback
func (p_ PlayerItem) SetIsApplicationAuthorizedForPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsApplicationAuthorizedForPlayback:"), value)
}


// A Boolean value that indicates whether the player item allows spatialized audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isaudiospatializationallowed
func (p_ PlayerItem) IsAudioSpatializationAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isAudioSpatializationAllowed"))
	return rv
}


// A Boolean value that indicates whether the player item allows spatialized audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isaudiospatializationallowed
func (p_ PlayerItem) SetIsAudioSpatializationAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsAudioSpatializationAllowed:"), value)
}


// A Boolean value that indicates whether authorization is required to play the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isauthorizationrequiredforplayback
func (p_ PlayerItem) IsAuthorizationRequiredForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isAuthorizationRequiredForPlayback"))
	return rv
}


// A Boolean value that indicates whether authorization is required to play the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isauthorizationrequiredforplayback
func (p_ PlayerItem) SetIsAuthorizationRequiredForPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsAuthorizationRequiredForPlayback:"), value)
}


// A Boolean value that indicates whether the content has been authorized by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/iscontentauthorizedforplayback
func (p_ PlayerItem) IsContentAuthorizedForPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isContentAuthorizedForPlayback"))
	return rv
}


// A Boolean value that indicates whether the content has been authorized by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/iscontentauthorizedforplayback
func (p_ PlayerItem) SetIsContentAuthorizedForPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsContentAuthorizedForPlayback:"), value)
}


// A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybackbufferempty
func (p_ PlayerItem) IsPlaybackBufferEmpty() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPlaybackBufferEmpty"))
	return rv
}


// A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybackbufferempty
func (p_ PlayerItem) SetIsPlaybackBufferEmpty(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPlaybackBufferEmpty:"), value)
}


// A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybackbufferfull
func (p_ PlayerItem) IsPlaybackBufferFull() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPlaybackBufferFull"))
	return rv
}


// A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybackbufferfull
func (p_ PlayerItem) SetIsPlaybackBufferFull(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPlaybackBufferFull:"), value)
}


// A Boolean value that indicates whether the item will likely play through without stalling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybacklikelytokeepup
func (p_ PlayerItem) IsPlaybackLikelyToKeepUp() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPlaybackLikelyToKeepUp"))
	return rv
}


// A Boolean value that indicates whether the item will likely play through without stalling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isplaybacklikelytokeepup
func (p_ PlayerItem) SetIsPlaybackLikelyToKeepUp(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPlaybackLikelyToKeepUp:"), value)
}


// The template player item that initializes this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/template
func (p_ PlayerItem) Template() IAVPlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("template"))
	return rv
}


// The template player item that initializes this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/template
func (p_ PlayerItem) SetTemplate(value IAVPlayerItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTemplate:"), value)
}


