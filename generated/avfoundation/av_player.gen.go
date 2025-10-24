// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Player] class.
var (
	PlayerClass     _PlayerClass
	PlayerClassOnce sync.Once
)

func getPlayerClass() _PlayerClass {
	PlayerClassOnce.Do(func() {
		PlayerClass = _PlayerClass{objc.GetClass("AVPlayer")}
	})
	return PlayerClass
}

type _PlayerClass struct {
	class objc.Class
}





// An interface definition for the [Player] class.
type IPlayer interface {
	objectivec.IObject
	

	// properties:
	ActionAtItemEnd() PlayerActionAtItemEnd
	SetActionAtItemEnd(value PlayerActionAtItemEnd)
	AllowsExternalPlayback() bool
	SetAllowsExternalPlayback(value bool)
	AppliesMediaSelectionCriteriaAutomatically() bool
	SetAppliesMediaSelectionCriteriaAutomatically(value bool)
	AudioOutputDeviceUniqueID() objc.IObject /* cross-framework: NSString */
	SetAudioOutputDeviceUniqueID(value objc.IObject /* cross-framework: NSString */)
	AudiovisualBackgroundPlaybackPolicy() PlayerAudiovisualBackgroundPlaybackPolicy
	SetAudiovisualBackgroundPlaybackPolicy(value PlayerAudiovisualBackgroundPlaybackPolicy)
	AutomaticallyWaitsToMinimizeStalling() bool
	SetAutomaticallyWaitsToMinimizeStalling(value bool)
	CurrentItem() IAVPlayerItem
	DefaultRate() float32
	SetDefaultRate(value float32)
	Error() Error
	ClosedCaptionDisplayEnabled() bool
	SetClosedCaptionDisplayEnabled(value bool)
	ExternalPlaybackActive() bool
	Muted() bool
	SetMuted(value bool)
	OutputObscuredDueToInsufficientExternalProtection() bool
	MasterClock() ClockRef /* not a class type */
	SetMasterClock(value ClockRef /* not a class type */)
	NetworkResourcePriority() PlayerNetworkResourcePriority
	SetNetworkResourcePriority(value PlayerNetworkResourcePriority)
	PlaybackCoordinator() IAVPlayerPlaybackCoordinator
	PreferredVideoDecoderGPURegistryID() uint64
	SetPreferredVideoDecoderGPURegistryID(value uint64)
	PreventsDisplaySleepDuringVideoPlayback() bool
	SetPreventsDisplaySleepDuringVideoPlayback(value bool)
	Rate() float32
	SetRate(value float32)
	ReasonForWaitingToPlay() PlayerWaitingReason /* typedef */
	SourceClock() ClockRef /* not a class type */
	SetSourceClock(value ClockRef /* not a class type */)
	Status() PlayerStatus
	TimeControlStatus() PlayerTimeControlStatus
	VideoOutput() IAVPlayerVideoOutput
	SetVideoOutput(value IAVPlayerVideoOutput)
	Volume() float32
	SetVolume(value float32)
	IsAirPlayVideoActive() bool
	SetIsAirPlayVideoActive(value bool)
	IsClosedCaptionDisplayEnabled() bool
	SetIsClosedCaptionDisplayEnabled(value bool)
	IsExternalPlaybackActive() bool
	SetIsExternalPlaybackActive(value bool)
	IsMuted() bool
	SetIsMuted(value bool)
	IsOutputObscuredDueToInsufficientExternalProtection() bool
	SetIsOutputObscuredDueToInsufficientExternalProtection(value bool)
	AllowedAudioSpatializationFormats() AudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats)
	IsAudioSpatializationAllowed() bool
	SetIsAudioSpatializationAllowed(value bool)


	

	// methods:
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.Value, queue objectivec.IObject, block unsafe.Pointer) objc.ID
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objc.IObject /* cross-framework: Time */, queue objectivec.IObject, block unsafe.Pointer) objc.ID
	CancelPendingPrerolls()
	CurrentTime() objc.IObject /* cross-framework: Time */
	MediaSelectionCriteriaForMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) IPlayerMediaSelectionCriteria
	Pause()
	Play()
	PlayImmediatelyAtRate(rate float32)
	PrerollAtRateCompletionHandler(rate float32, completionHandler unsafe.Pointer)
	RemoveTimeObserver(observer objc.IObject)
	ReplaceCurrentItemWithPlayerItem(item IAVPlayerItem)
	SeekToTime(time objc.IObject /* cross-framework: Time */)
	SeekToDate(date objc.IObject /* cross-framework: NSDate */)
	SeekToTimeCompletionHandler(time objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
	SeekToDateCompletionHandler(date objc.IObject /* cross-framework: NSDate */, completionHandler unsafe.Pointer)
	SeekToTimeToleranceBeforeToleranceAfter(time objc.IObject /* cross-framework: Time */, toleranceBefore objc.IObject /* cross-framework: Time */, toleranceAfter objc.IObject /* cross-framework: Time */)
	SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objc.IObject /* cross-framework: Time */, toleranceBefore objc.IObject /* cross-framework: Time */, toleranceAfter objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
	SetMediaSelectionCriteriaForMediaCharacteristic(criteria IAVPlayerMediaSelectionCriteria, mediaCharacteristic MediaCharacteristic /* typedef */)
	SetRateTimeAtHostTime(rate float32, itemTime objc.IObject /* cross-framework: Time */, hostClockTime objc.IObject /* cross-framework: Time */)


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerClass) Alloc() Player {
	rv := objc.Send[Player](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerClass) New() Player {
	rv := objc.Send[Player](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Player) Init() Player {
	rv := objc.Send[Player](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Player) Autorelease() Player {
	rv := objc.Send[Player](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayer creates a new Player instance.
func NewPlayer() Player {
	return getPlayerClass().New()
}





// An object that provides the interface to control the player’s transport behavior.
//
// A player is a controller object that manages the playback and timing of a media asset. Use an instance of to play local and remote file-based media, such as QuickTime movies and MP3 audio files, as well as audiovisual media served using HTTP Live Streaming. Use a player object to play a single media asset. You can reuse the player instance to play additional media assets using its method, but it manages the playback of only a single media asset at a time. The framework also provides a subclass called that you can use to manage the playback of a queue of media assets. You use an to play media assets, which AVFoundation represents using the class. only models the aspects of the media, such as its duration or creation date, and on its own, isn’t suitable for playback with an . To play an asset, you create an instance of its counterpart found in . This object models the timing and presentation state of an asset played by an instance of . See the reference for more details. is a dynamic object whose state continuously changes. There are two approaches you can use to observe a player’s state: You can use key-value observing (KVO) to observe state changes to many of the player’s dynamic properties, such as its or its playback . KVO works well for general state observations, but isn’t intended for observing continuously changing state like the player’s time. provides two methods to observe time changes: These methods let you observe time changes either periodically or by boundary, respectively. As changes occur, invoke the callback block or closure you supply to these methods to give you the opportunity to take some action such as updating the state of your player’s user interface. and are nonvisual objects, meaning that on their own they’re unable to present an asset’s video onscreen. There are two primary approaches you use to present your video content onscreen: The best way to present your video content is with the AVKit framework’s class in iOS and tvOS, or the class in macOS. These classes present the video content, along with playback controls and other media features giving you a full-featured playback experience. When building a custom interface for your player, use . You can set this layer a view’s backing layer or add it directly to the layer hierarchy. Unlike and , a player layer doesn’t present any playback controls—it only presents the visual content onscreen. It’s up to you to build the playback transport controls to play, pause, and seek through the media. Alongside the visual content presented with AVKit or , you can also present animated content synchronized with the player’s timing using . Use a synchronized layer pass along player timing to its layer subtree. You can use to build custom effects in Core Animation, such as animated lower thirds or video transitions, and have them play in sync with the timing of the player’s current .


// An object that provides the interface to control the player’s transport behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer
type Player struct {
	objectivec.Object
}

// PlayerFrom constructs a [Player] from an unsafe.Pointer.
//
// An object that provides the interface to control the player’s transport behavior.
func PlayerFrom(ptr unsafe.Pointer) Player {
	return Player{objectivec.Object{objc.ID(ptr)}}
}






// Creates a new player to play the specified player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(playerItem:)
func NewPlayerWithPlayerItem(item IAVPlayerItem) Player {
	instance := getPlayerClass().Alloc()
	rv := objc.Send[Player](instance.ID, objc.Sel("initWithPlayerItem:"), item)
	rv.Autorelease()
	return rv
}


// Creates a new player to play a single audiovisual resource referenced by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(url:)
func NewPlayerWithURL(URL objc.IObject /* cross-framework: NSURL */) Player {
	instance := getPlayerClass().Alloc()
	rv := objc.Send[Player](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}







// Returns a new player initialized to play the specified player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playerWithPlayerItem:
func (pc _PlayerClass) PlayerWithPlayerItem(item IAVPlayerItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("playerWithPlayerItem:"), item)
	return rv
}


// Returns a new player to play a single audiovisual resource referenced by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playerWithURL:
func (pc _PlayerClass) PlayerWithURL(URL objc.IObject /* cross-framework: NSURL */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("playerWithURL:"), URL)
	return rv
}







// The HDR modes that are available for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/availableHDRModes
func (pc _PlayerClass) AvailableHDRModes() PlayerHDRMode {
	rv := objc.Send[PlayerHDRMode](objc.ID(pc.class), objc.Sel("availableHDRModes"))
	return rv
}

// A Boolean value that indicates whether the current device can present content to an HDR display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/eligibleForHDRPlayback
func (pc _PlayerClass) EligibleForHDRPlayback() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("eligibleForHDRPlayback"))
	return rv
}

// AVPlayer and other AVFoundation types can optionally be observed using Swift Observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isObservationEnabled
func (pc _PlayerClass) ObservationEnabled() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("observationEnabled"))
	return rv
}






// Requests the invocation of a block when specified times are traversed during normal playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/addBoundaryTimeObserver(forTimes:queue:using:)
func (p_ Player) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.Value, queue objectivec.IObject, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}


// Requests the periodic invocation of a given block during playback to report changing time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/addPeriodicTimeObserver(forInterval:queue:using:)
func (p_ Player) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objc.IObject /* cross-framework: Time */, queue objectivec.IObject, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, queue, block)
	return rv
}


// Cancels any pending preroll requests and invokes the corresponding completion handlers, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/cancelPendingPrerolls()
func (p_ Player) CancelPendingPrerolls() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelPendingPrerolls"))
}


// Returns the current time of the current player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/currentTime()
func (p_ Player) CurrentTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("currentTime"))
	return rv
}


// Returns the automatic selection criteria for media items with the specified media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/mediaSelectionCriteria(forMediaCharacteristic:)
func (p_ Player) MediaSelectionCriteriaForMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) IPlayerMediaSelectionCriteria {
	rv := objc.Send[PlayerMediaSelectionCriteria](p_.ID, objc.Sel("mediaSelectionCriteriaForMediaCharacteristic:"), mediaCharacteristic)
	return rv
}


// Pauses playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/pause()
func (p_ Player) Pause() {
	objc.Send[objc.ID](p_.ID, objc.Sel("pause"))
}


// Begins playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/play()
func (p_ Player) Play() {
	objc.Send[objc.ID](p_.ID, objc.Sel("play"))
}


// Plays the available media data immediately, at the specified rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playImmediately(atRate:)
func (p_ Player) PlayImmediatelyAtRate(rate float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("playImmediatelyAtRate:"), rate)
}


// Begins loading media data to prime the media pipelines for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preroll(atRate:completionHandler:)
func (p_ Player) PrerollAtRateCompletionHandler(rate float32, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("prerollAtRate:completionHandler:"), rate, completionHandler)
}


// Cancels a previously registered periodic or boundary time observer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/removeTimeObserver(_:)
func (p_ Player) RemoveTimeObserver(observer objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeTimeObserver:"), observer)
}


// Replaces the current item with a new item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/replaceCurrentItem(with:)
func (p_ Player) ReplaceCurrentItemWithPlayerItem(item IAVPlayerItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("replaceCurrentItemWithPlayerItem:"), item)
}


// Requests that the player seek to a specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:)-87h2r
func (p_ Player) SeekToTime(time objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:"), time)
}


// Requests that the player seek to a specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:)-9h9qr
func (p_ Player) SeekToDate(date objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToDate:"), date)
}


// Requests that the player seek to a specified time, and to notify you when the seek is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:completionHandler:)-75bls
func (p_ Player) SeekToTimeCompletionHandler(time objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:completionHandler:"), time, completionHandler)
}


// Requests that the player seek to a specified date, and to notify you when the seek is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:completionHandler:)-wr1l
func (p_ Player) SeekToDateCompletionHandler(date objc.IObject /* cross-framework: NSDate */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToDate:completionHandler:"), date, completionHandler)
}


// Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:toleranceBefore:toleranceAfter:)
func (p_ Player) SeekToTimeToleranceBeforeToleranceAfter(time objc.IObject /* cross-framework: Time */, toleranceBefore objc.IObject /* cross-framework: Time */, toleranceAfter objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:"), time, toleranceBefore, toleranceAfter)
}


// Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values, and to notify you when the seek is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:toleranceBefore:toleranceAfter:completionHandler:)
func (p_ Player) SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objc.IObject /* cross-framework: Time */, toleranceBefore objc.IObject /* cross-framework: Time */, toleranceAfter objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:completionHandler:"), time, toleranceBefore, toleranceAfter, completionHandler)
}


// Applies automatic selection criteria for media that has the specified media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/setMediaSelectionCriteria(_:forMediaCharacteristic:)
func (p_ Player) SetMediaSelectionCriteriaForMediaCharacteristic(criteria IAVPlayerMediaSelectionCriteria, mediaCharacteristic MediaCharacteristic /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaSelectionCriteria:forMediaCharacteristic:"), criteria, mediaCharacteristic)
}


// Synchronizes the playback rate and time of the current item with an external source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/setRate(_:time:atHostTime:)
func (p_ Player) SetRateTimeAtHostTime(rate float32, itemTime objc.IObject /* cross-framework: Time */, hostClockTime objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRate:time:atHostTime:"), rate, itemTime, hostClockTime)
}







// The action to perform when the current player item has finished playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/actionAtItemEnd-swift.property
func (p_ Player) ActionAtItemEnd() PlayerActionAtItemEnd {
	rv := objc.Send[PlayerActionAtItemEnd](p_.ID, objc.Sel("actionAtItemEnd"))
	return rv
}


// The action to perform when the current player item has finished playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/actionAtItemEnd-swift.property
func (p_ Player) SetActionAtItemEnd(value PlayerActionAtItemEnd) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setActionAtItemEnd:"), value)
}


// A Boolean value that indicates whether the player allows switching to external playback mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/allowsExternalPlayback
func (p_ Player) AllowsExternalPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsExternalPlayback"))
	return rv
}


// A Boolean value that indicates whether the player allows switching to external playback mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/allowsExternalPlayback
func (p_ Player) SetAllowsExternalPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsExternalPlayback:"), value)
}


// A Boolean value that indicates whether the receiver should apply the current selection criteria automatically to player items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/appliesMediaSelectionCriteriaAutomatically
func (p_ Player) AppliesMediaSelectionCriteriaAutomatically() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("appliesMediaSelectionCriteriaAutomatically"))
	return rv
}


// A Boolean value that indicates whether the receiver should apply the current selection criteria automatically to player items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/appliesMediaSelectionCriteriaAutomatically
func (p_ Player) SetAppliesMediaSelectionCriteriaAutomatically(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppliesMediaSelectionCriteriaAutomatically:"), value)
}


// Specifies the unique ID of the Core Audio output device used to play audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/audioOutputDeviceUniqueID
func (p_ Player) AudioOutputDeviceUniqueID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("audioOutputDeviceUniqueID"))
	return rv
}


// Specifies the unique ID of the Core Audio output device used to play audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/audioOutputDeviceUniqueID
func (p_ Player) SetAudioOutputDeviceUniqueID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioOutputDeviceUniqueID:"), value)
}


// A policy that determines how playback of audiovisual media continues when the app transitions to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/audiovisualBackgroundPlaybackPolicy
func (p_ Player) AudiovisualBackgroundPlaybackPolicy() PlayerAudiovisualBackgroundPlaybackPolicy {
	rv := objc.Send[PlayerAudiovisualBackgroundPlaybackPolicy](p_.ID, objc.Sel("audiovisualBackgroundPlaybackPolicy"))
	return rv
}


// A policy that determines how playback of audiovisual media continues when the app transitions to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/audiovisualBackgroundPlaybackPolicy
func (p_ Player) SetAudiovisualBackgroundPlaybackPolicy(value PlayerAudiovisualBackgroundPlaybackPolicy) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudiovisualBackgroundPlaybackPolicy:"), value)
}


// A Boolean value that indicates whether the player should automatically delay playback in order to minimize stalling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/automaticallyWaitsToMinimizeStalling
func (p_ Player) AutomaticallyWaitsToMinimizeStalling() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("automaticallyWaitsToMinimizeStalling"))
	return rv
}


// A Boolean value that indicates whether the player should automatically delay playback in order to minimize stalling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/automaticallyWaitsToMinimizeStalling
func (p_ Player) SetAutomaticallyWaitsToMinimizeStalling(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticallyWaitsToMinimizeStalling:"), value)
}


// The item for which the player is currently controlling playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/currentItem
func (p_ Player) CurrentItem() IAVPlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("currentItem"))
	return rv
}


// A default rate at which to begin playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/defaultRate
func (p_ Player) DefaultRate() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("defaultRate"))
	return rv
}


// A default rate at which to begin playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/defaultRate
func (p_ Player) SetDefaultRate(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultRate:"), value)
}


// A Boolean value that indicates whether the current device can present content to an HDR display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/eligibleForHDRPlayback
func (p_ Player) EligibleForHDRPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("eligibleForHDRPlayback"))
	return rv
}


// An error that caused a failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/error
func (p_ Player) Error() Error {
	rv := objc.Send[Error](p_.ID, objc.Sel("error"))
	return rv
}


// A Boolean value that indicates whether the player uses closed captioning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isClosedCaptionDisplayEnabled
func (p_ Player) ClosedCaptionDisplayEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("closedCaptionDisplayEnabled"))
	return rv
}


// A Boolean value that indicates whether the player uses closed captioning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isClosedCaptionDisplayEnabled
func (p_ Player) SetClosedCaptionDisplayEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setClosedCaptionDisplayEnabled:"), value)
}


// A Boolean value that indicates whether the player is currently playing video in external playback mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isExternalPlaybackActive
func (p_ Player) ExternalPlaybackActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("externalPlaybackActive"))
	return rv
}


// A Boolean value that indicates whether the audio output of the player is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isMuted
func (p_ Player) Muted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("muted"))
	return rv
}


// A Boolean value that indicates whether the audio output of the player is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isMuted
func (p_ Player) SetMuted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMuted:"), value)
}


// AVPlayer and other AVFoundation types can optionally be observed using Swift Observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isObservationEnabled
func (p_ Player) ObservationEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("observationEnabled"))
	return rv
}


// AVPlayer and other AVFoundation types can optionally be observed using Swift Observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isObservationEnabled
func (p_ Player) SetObservationEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObservationEnabled:"), value)
}


// A Boolean value that indicates whether output is being obscured because of insufficient external protection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isOutputObscuredDueToInsufficientExternalProtection
func (p_ Player) OutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("outputObscuredDueToInsufficientExternalProtection"))
	return rv
}


// The host clock for item time bases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/masterClock
func (p_ Player) MasterClock() ClockRef /* not a class type */ {
	rv := objc.Send[ClockRef](p_.ID, objc.Sel("masterClock"))
	return rv
}


// The host clock for item time bases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/masterClock
func (p_ Player) SetMasterClock(value ClockRef /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMasterClock:"), value)
}


// Indicates the priority of this player for network bandwidth resource distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/networkResourcePriority-swift.property
func (p_ Player) NetworkResourcePriority() PlayerNetworkResourcePriority {
	rv := objc.Send[PlayerNetworkResourcePriority](p_.ID, objc.Sel("networkResourcePriority"))
	return rv
}


// Indicates the priority of this player for network bandwidth resource distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/networkResourcePriority-swift.property
func (p_ Player) SetNetworkResourcePriority(value PlayerNetworkResourcePriority) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNetworkResourcePriority:"), value)
}


// The playback coordinator for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playbackCoordinator
func (p_ Player) PlaybackCoordinator() IAVPlayerPlaybackCoordinator {
	rv := objc.Send[PlayerPlaybackCoordinator](p_.ID, objc.Sel("playbackCoordinator"))
	return rv
}


// The registry identifier for the GPU used for video decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preferredVideoDecoderGPURegistryID
func (p_ Player) PreferredVideoDecoderGPURegistryID() uint64 {
	rv := objc.Send[uint64](p_.ID, objc.Sel("preferredVideoDecoderGPURegistryID"))
	return rv
}


// The registry identifier for the GPU used for video decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preferredVideoDecoderGPURegistryID
func (p_ Player) SetPreferredVideoDecoderGPURegistryID(value uint64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredVideoDecoderGPURegistryID:"), value)
}


// A Boolean value that indicates whether video playback prevents display and device sleep.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preventsDisplaySleepDuringVideoPlayback
func (p_ Player) PreventsDisplaySleepDuringVideoPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}


// A Boolean value that indicates whether video playback prevents display and device sleep.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preventsDisplaySleepDuringVideoPlayback
func (p_ Player) SetPreventsDisplaySleepDuringVideoPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}


// The current playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/rate
func (p_ Player) Rate() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("rate"))
	return rv
}


// The current playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/rate
func (p_ Player) SetRate(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRate:"), value)
}


// The reason the player is currently waiting for playback to begin or resume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/reasonForWaitingToPlay
func (p_ Player) ReasonForWaitingToPlay() PlayerWaitingReason /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("reasonForWaitingToPlay"))
	return rv
}


// A clock the player uses for item time bases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/sourceClock
func (p_ Player) SourceClock() ClockRef /* not a class type */ {
	rv := objc.Send[ClockRef](p_.ID, objc.Sel("sourceClock"))
	return rv
}


// A clock the player uses for item time bases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/sourceClock
func (p_ Player) SetSourceClock(value ClockRef /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSourceClock:"), value)
}


// A value that indicates the readiness of a player object for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/status-swift.property
func (p_ Player) Status() PlayerStatus {
	rv := objc.Send[PlayerStatus](p_.ID, objc.Sel("status"))
	return rv
}


// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/timeControlStatus-swift.property
func (p_ Player) TimeControlStatus() PlayerTimeControlStatus {
	rv := objc.Send[PlayerTimeControlStatus](p_.ID, objc.Sel("timeControlStatus"))
	return rv
}


// The video output for this player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/videoOutput
func (p_ Player) VideoOutput() IAVPlayerVideoOutput {
	rv := objc.Send[PlayerVideoOutput](p_.ID, objc.Sel("videoOutput"))
	return rv
}


// The video output for this player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/videoOutput
func (p_ Player) SetVideoOutput(value IAVPlayerVideoOutput) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoOutput:"), value)
}


// The audio playback volume for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/volume
func (p_ Player) Volume() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("volume"))
	return rv
}


// The audio playback volume for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/volume
func (p_ Player) SetVolume(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVolume:"), value)
}


// A Boolean value that indicates whether the player is playing video through AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/isairplayvideoactive
func (p_ Player) IsAirPlayVideoActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isAirPlayVideoActive"))
	return rv
}


// A Boolean value that indicates whether the player is playing video through AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/isairplayvideoactive
func (p_ Player) SetIsAirPlayVideoActive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsAirPlayVideoActive:"), value)
}


// A Boolean value that indicates whether the player uses closed captioning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/isclosedcaptiondisplayenabled
func (p_ Player) IsClosedCaptionDisplayEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isClosedCaptionDisplayEnabled"))
	return rv
}


// A Boolean value that indicates whether the player uses closed captioning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/isclosedcaptiondisplayenabled
func (p_ Player) SetIsClosedCaptionDisplayEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsClosedCaptionDisplayEnabled:"), value)
}


// A Boolean value that indicates whether the player is currently playing video in external playback mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/isexternalplaybackactive
func (p_ Player) IsExternalPlaybackActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isExternalPlaybackActive"))
	return rv
}


// A Boolean value that indicates whether the player is currently playing video in external playback mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/isexternalplaybackactive
func (p_ Player) SetIsExternalPlaybackActive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsExternalPlaybackActive:"), value)
}


// A Boolean value that indicates whether the audio output of the player is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/ismuted
func (p_ Player) IsMuted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isMuted"))
	return rv
}


// A Boolean value that indicates whether the audio output of the player is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/ismuted
func (p_ Player) SetIsMuted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsMuted:"), value)
}


// A Boolean value that indicates whether output is being obscured because of insufficient external protection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/isoutputobscuredduetoinsufficientexternalprotection
func (p_ Player) IsOutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOutputObscuredDueToInsufficientExternalProtection"))
	return rv
}


// A Boolean value that indicates whether output is being obscured because of insufficient external protection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/isoutputobscuredduetoinsufficientexternalprotection
func (p_ Player) SetIsOutputObscuredDueToInsufficientExternalProtection(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOutputObscuredDueToInsufficientExternalProtection:"), value)
}


// The source audio channel layouts the player item supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/allowedaudiospatializationformats
func (p_ Player) AllowedAudioSpatializationFormats() AudioSpatializationFormats {
	rv := objc.Send[AudioSpatializationFormats](p_.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}


// The source audio channel layouts the player item supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/allowedaudiospatializationformats
func (p_ Player) SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}


// A Boolean value that indicates whether the player item allows spatialized audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isaudiospatializationallowed
func (p_ Player) IsAudioSpatializationAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isAudioSpatializationAllowed"))
	return rv
}


// A Boolean value that indicates whether the player item allows spatialized audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/isaudiospatializationallowed
func (p_ Player) SetIsAudioSpatializationAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsAudioSpatializationAllowed:"), value)
}







