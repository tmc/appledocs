// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayer] class.
var aVPlayerClass = _AVPlayerClass{objc.GetClass("AVPlayer")}

type _AVPlayerClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayer] class.
type IAVPlayer interface {
	objectivec.IObject
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) objc.ID
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) objc.ID
	CancelPendingPrerolls()
	CurrentTime() unsafe.Pointer
	MediaSelectionCriteriaForMediaCharacteristic(mediaCharacteristic unsafe.Pointer) unsafe.Pointer
	Pause()
	Play()
	PlayImmediatelyAtRate(rate float32)
	PrerollAtRateCompletionHandler(rate float32, completionHandler unsafe.Pointer)
	RemoveTimeObserver(observer objc.ID)
	ReplaceCurrentItemWithPlayerItem(item unsafe.Pointer)
	SeekToTime(time unsafe.Pointer)
	SeekToDate(date unsafe.Pointer)
	SeekToTimeCompletionHandler(time unsafe.Pointer, completionHandler unsafe.Pointer)
	SeekToDateCompletionHandler(date unsafe.Pointer, completionHandler unsafe.Pointer)
	SeekToTimeToleranceBeforeToleranceAfter(time unsafe.Pointer, toleranceBefore unsafe.Pointer, toleranceAfter unsafe.Pointer)
	SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time unsafe.Pointer, toleranceBefore unsafe.Pointer, toleranceAfter unsafe.Pointer, completionHandler unsafe.Pointer)
	SetMediaSelectionCriteriaForMediaCharacteristic(criteria unsafe.Pointer, mediaCharacteristic unsafe.Pointer)
	SetRateTimeAtHostTime(rate float32, itemTime unsafe.Pointer, hostClockTime unsafe.Pointer)
}

// An object that provides the interface to control the player’s transport behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer

type AVPlayer struct {
	objectivec.Object
}

// AVPlayerFrom constructs a [AVPlayer] from an unsafe.Pointer.
//
// An object that provides the interface to control the player’s transport behavior.
func AVPlayerFrom(ptr unsafe.Pointer) AVPlayer {
	return AVPlayer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVPlayerClass) Alloc() AVPlayer {
	rv := objc.Send[AVPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVPlayerClass) New() AVPlayer {
	rv := objc.Send[AVPlayer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayer) Init() AVPlayer {
	rv := objc.Send[AVPlayer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayer) Autorelease() AVPlayer {
	rv := objc.Send[AVPlayer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayer creates a new AVPlayer instance.
func NewAVPlayer() AVPlayer {
	return aVPlayerClass.New()
}


// Creates a new player to play the specified player item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(playerItem:)
func NewAVPlayerWithPlayerItem(item unsafe.Pointer) AVPlayer {
	instance := aVPlayerClass.Alloc()
	rv := objc.Send[AVPlayer](instance.ID, objc.Sel("initWithPlayerItem:"), item)
	rv.Autorelease()
	return rv
}
// Creates a new player to play a single audiovisual resource referenced by a given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(url:)
func NewAVPlayerWithURL(URL unsafe.Pointer) AVPlayer {
	instance := aVPlayerClass.Alloc()
	rv := objc.Send[AVPlayer](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}


// Returns a new player initialized to play the specified player item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playerWithPlayerItem:
func (ac _AVPlayerClass) PlayerWithPlayerItem(item unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("playerWithPlayerItem:"), item)
	return rv
}
// Returns a new player to play a single audiovisual resource referenced by a given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playerWithURL:
func (ac _AVPlayerClass) PlayerWithURL(URL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("playerWithURL:"), URL)
	return rv
}
// Requests the invocation of a block when specified times are traversed during normal playback. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/addBoundaryTimeObserver(forTimes:queue:using:)
func (a_ AVPlayer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}
// Requests the periodic invocation of a given block during playback to report changing time. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/addPeriodicTimeObserver(forInterval:queue:using:)
func (a_ AVPlayer) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, queue, block)
	return rv
}
// Cancels any pending preroll requests and invokes the corresponding completion handlers, if present. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/cancelPendingPrerolls()
func (a_ AVPlayer) CancelPendingPrerolls() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelPendingPrerolls"))
}
// Returns the current time of the current player item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/currentTime()
func (a_ AVPlayer) CurrentTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentTime"))
	return rv
}
// Returns the automatic selection criteria for media items with the specified media characteristic. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/mediaSelectionCriteria(forMediaCharacteristic:)
func (a_ AVPlayer) MediaSelectionCriteriaForMediaCharacteristic(mediaCharacteristic unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mediaSelectionCriteriaForMediaCharacteristic:"), mediaCharacteristic)
	return rv
}
// Pauses playback of the current item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/pause()
func (a_ AVPlayer) Pause() {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause"))
}
// Begins playback of the current item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/play()
func (a_ AVPlayer) Play() {
	objc.Send[objc.ID](a_.ID, objc.Sel("play"))
}
// Plays the available media data immediately, at the specified rate. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playImmediately(atRate:)
func (a_ AVPlayer) PlayImmediatelyAtRate(rate float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("playImmediatelyAtRate:"), rate)
}
// Begins loading media data to prime the media pipelines for playback. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preroll(atRate:completionHandler:)
func (a_ AVPlayer) PrerollAtRateCompletionHandler(rate float32, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prerollAtRate:completionHandler:"), rate, completionHandler)
}
// Cancels a previously registered periodic or boundary time observer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/removeTimeObserver(_:)
func (a_ AVPlayer) RemoveTimeObserver(observer objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeTimeObserver:"), observer)
}
// Replaces the current item with a new item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/replaceCurrentItem(with:)
func (a_ AVPlayer) ReplaceCurrentItemWithPlayerItem(item unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("replaceCurrentItemWithPlayerItem:"), item)
}
// Requests that the player seek to a specified time. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:)-87h2r
func (a_ AVPlayer) SeekToTime(time unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("seekToTime:"), time)
}
// Requests that the player seek to a specified date. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:)-9h9qr
func (a_ AVPlayer) SeekToDate(date unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("seekToDate:"), date)
}
// Requests that the player seek to a specified time, and to notify you when the seek is complete. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:completionHandler:)-75bls
func (a_ AVPlayer) SeekToTimeCompletionHandler(time unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("seekToTime:completionHandler:"), time, completionHandler)
}
// Requests that the player seek to a specified date, and to notify you when the seek is complete. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:completionHandler:)-wr1l
func (a_ AVPlayer) SeekToDateCompletionHandler(date unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("seekToDate:completionHandler:"), date, completionHandler)
}
// Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:toleranceBefore:toleranceAfter:)
func (a_ AVPlayer) SeekToTimeToleranceBeforeToleranceAfter(time unsafe.Pointer, toleranceBefore unsafe.Pointer, toleranceAfter unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:"), time, toleranceBefore, toleranceAfter)
}
// Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values, and to notify you when the seek is complete. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:toleranceBefore:toleranceAfter:completionHandler:)
func (a_ AVPlayer) SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time unsafe.Pointer, toleranceBefore unsafe.Pointer, toleranceAfter unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:completionHandler:"), time, toleranceBefore, toleranceAfter, completionHandler)
}
// Applies automatic selection criteria for media that has the specified media characteristic. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/setMediaSelectionCriteria(_:forMediaCharacteristic:)
func (a_ AVPlayer) SetMediaSelectionCriteriaForMediaCharacteristic(criteria unsafe.Pointer, mediaCharacteristic unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaSelectionCriteria:forMediaCharacteristic:"), criteria, mediaCharacteristic)
}
// Synchronizes the playback rate and time of the current item with an external source. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/setRate(_:time:atHostTime:)
func (a_ AVPlayer) SetRateTimeAtHostTime(rate float32, itemTime unsafe.Pointer, hostClockTime unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:time:atHostTime:"), rate, itemTime, hostClockTime)
}

