// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioPlayer */


/* debug [class_header]: Header for AVAudioPlayer */
// The class instance for the [AudioPlayer] class.
var (
	AudioPlayerClass     _AudioPlayerClass
	AudioPlayerClassOnce sync.Once
)

func getAudioPlayerClass() _AudioPlayerClass {
	AudioPlayerClassOnce.Do(func() {
		AudioPlayerClass = _AudioPlayerClass{objc.GetClass("AVAudioPlayer")}
	})
	return AudioPlayerClass
}

type _AudioPlayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioPlayer */
// An interface definition for the [AudioPlayer] class.
type IAudioPlayer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioPlayer */
	// properties:
	CurrentDevice() objc.IObject /* cross-framework: NSString */
	SetCurrentDevice(value objc.IObject /* cross-framework: NSString */)
	CurrentTime() float64
	SetCurrentTime(value float64)
	Data() objc.IObject /* cross-framework: NSData */
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DeviceCurrentTime() float64
	Duration() float64
	EnableRate() bool
	SetEnableRate(value bool)
	Format() IAVAudioFormat
	MeteringEnabled() bool
	SetMeteringEnabled(value bool)
	Playing() bool
	NumberOfChannels() uint
	NumberOfLoops() int
	SetNumberOfLoops(value int)
	Pan() float32
	SetPan(value float32)
	Rate() float32
	SetRate(value float32)
	Settings() foundation.IDictionary
	Url() objc.IObject /* cross-framework: NSURL */
	Volume() float32
	SetVolume(value float32)
	IsMeteringEnabled() bool
	SetIsMeteringEnabled(value bool)
	IsPlaying() bool
	SetIsPlaying(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioPlayer */
	// methods:
	AveragePowerForChannel(channelNumber uint) float32
	Pause()
	PeakPowerForChannel(channelNumber uint) float32
	Play() bool
	PlayAtTime(time float64) bool
	PrepareToPlay() bool
	SetVolumeFadeDuration(volume float32, duration float64)
	Stop()
	UpdateMeters()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioPlayer */
// Alloc allocates a new instance without initialization.
func (ac _AudioPlayerClass) Alloc() AudioPlayer {
	rv := objc.Send[AudioPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioPlayerClass) New() AudioPlayer {
	rv := objc.Send[AudioPlayer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioPlayer) Init() AudioPlayer {
	rv := objc.Send[AudioPlayer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioPlayer) Autorelease() AudioPlayer {
	rv := objc.Send[AudioPlayer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioPlayer creates a new AudioPlayer instance.
func NewAudioPlayer() AudioPlayer {
	return getAudioPlayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioPlayer */
// An object that plays audio data from a file or buffer.
//
// Use an audio player to: Play audio of any duration from a file or buffer Control the volume, panning, rate, and looping behavior of the played audio Access playback-level metering data Play multiple sounds simultaneously by synchronizing the playback of multiple players For more information about preparing your app to play audio, see .


// An object that plays audio data from a file or buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer
type AudioPlayer struct {
	objectivec.Object
}

// AudioPlayerFrom constructs a [AudioPlayer] from an unsafe.Pointer.
//
// An object that plays audio data from a file or buffer.
func AudioPlayerFrom(ptr unsafe.Pointer) AudioPlayer {
	return AudioPlayer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioPlayer */

// Creates a player to play audio from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:)
func NewAudioPlayerWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, outError objectivec.IObject) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioPlayerWithContentsOfURLError */


// Creates a player to play audio from a file of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:fileTypeHint:)
func NewAudioPlayerWithContentsOfURLFileTypeHintError(url objc.IObject /* cross-framework: NSURL */, utiString objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithContentsOfURL:fileTypeHint:error:"), url, utiString, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioPlayerWithContentsOfURLFileTypeHintError */


// Creates a player to play in-memory audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:)
func NewAudioPlayerWithDataError(data objc.IObject /* cross-framework: NSData */, outError objectivec.IObject) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithData:error:"), data, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioPlayerWithDataError */


// Creates a player to play in-memory audio data of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:fileTypeHint:)
func NewAudioPlayerWithDataFileTypeHintError(data objc.IObject /* cross-framework: NSData */, utiString objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithData:fileTypeHint:error:"), data, utiString, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioPlayerWithDataFileTypeHintError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioPlayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioPlayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioPlayer */

// Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/averagePower(forChannel:)
func (a_ AudioPlayer) AveragePowerForChannel(channelNumber uint) float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("averagePowerForChannel:"), channelNumber)
	return rv
}/* debug [instance_methods/method]: AveragePowerForChannel */


// Pauses audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pause()
func (a_ AudioPlayer) Pause() {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause"))
}/* debug [instance_methods/method]: Pause */


// Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/peakPower(forChannel:)
func (a_ AudioPlayer) PeakPowerForChannel(channelNumber uint) float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("peakPowerForChannel:"), channelNumber)
	return rv
}/* debug [instance_methods/method]: PeakPowerForChannel */


// Plays audio asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/play()
func (a_ AudioPlayer) Play() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("play"))
	return rv
}/* debug [instance_methods/method]: Play */


// Plays audio asynchronously, starting at a specified point in the audio output device’s timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/play(atTime:)
func (a_ AudioPlayer) PlayAtTime(time float64) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playAtTime:"), time)
	return rv
}/* debug [instance_methods/method]: PlayAtTime */


// Prepares the player for audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/prepareToPlay()
func (a_ AudioPlayer) PrepareToPlay() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prepareToPlay"))
	return rv
}/* debug [instance_methods/method]: PrepareToPlay */


// Changes the audio player’s volume over a duration of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/setVolume(_:fadeDuration:)
func (a_ AudioPlayer) SetVolumeFadeDuration(volume float32, duration float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVolume:fadeDuration:"), volume, duration)
}/* debug [instance_methods/method]: SetVolumeFadeDuration */


// Stops playback and undoes the setup the system requires for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/stop()
func (a_ AudioPlayer) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */


// Refreshes the average and peak power values for all channels of an audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/updateMeters()
func (a_ AudioPlayer) UpdateMeters() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateMeters"))
}/* debug [instance_methods/method]: UpdateMeters */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioPlayer */

// The unique identifier of the current audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentDevice
func (a_ AudioPlayer) CurrentDevice() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("currentDevice"))
	return rv
}/* debug [instance_properties/getter]: currentDevice */


// The unique identifier of the current audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentDevice
func (a_ AudioPlayer) SetCurrentDevice(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentDevice:"), value)
}/* debug [instance_properties/setter]: currentDevice */


// The current playback time, in seconds, within the audio timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentTime
func (a_ AudioPlayer) CurrentTime() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("currentTime"))
	return rv
}/* debug [instance_properties/getter]: currentTime */


// The current playback time, in seconds, within the audio timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentTime
func (a_ AudioPlayer) SetCurrentTime(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentTime:"), value)
}/* debug [instance_properties/setter]: currentTime */


// The audio data associated with the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/data
func (a_ AudioPlayer) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The delegate object for the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/delegate
func (a_ AudioPlayer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object for the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/delegate
func (a_ AudioPlayer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The time value, in seconds, of the audio output device’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/deviceCurrentTime
func (a_ AudioPlayer) DeviceCurrentTime() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("deviceCurrentTime"))
	return rv
}/* debug [instance_properties/getter]: deviceCurrentTime */


// The total duration, in seconds, of the player’s audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/duration
func (a_ AudioPlayer) Duration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// A Boolean value that indicates whether you can adjust the playback rate of the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/enableRate
func (a_ AudioPlayer) EnableRate() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enableRate"))
	return rv
}/* debug [instance_properties/getter]: enableRate */


// A Boolean value that indicates whether you can adjust the playback rate of the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/enableRate
func (a_ AudioPlayer) SetEnableRate(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnableRate:"), value)
}/* debug [instance_properties/setter]: enableRate */


// The format of the player’s audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/format
func (a_ AudioPlayer) Format() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isMeteringEnabled
func (a_ AudioPlayer) MeteringEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("meteringEnabled"))
	return rv
}/* debug [instance_properties/getter]: meteringEnabled */


// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isMeteringEnabled
func (a_ AudioPlayer) SetMeteringEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMeteringEnabled:"), value)
}/* debug [instance_properties/setter]: meteringEnabled */


// A Boolean value that indicates whether the player is currently playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isPlaying
func (a_ AudioPlayer) Playing() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playing"))
	return rv
}/* debug [instance_properties/getter]: playing */


// The number of audio channels in the player’s audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfChannels
func (a_ AudioPlayer) NumberOfChannels() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("numberOfChannels"))
	return rv
}/* debug [instance_properties/getter]: numberOfChannels */


// The number of times the audio repeats playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfLoops
func (a_ AudioPlayer) NumberOfLoops() int {
	rv := objc.Send[int](a_.ID, objc.Sel("numberOfLoops"))
	return rv
}/* debug [instance_properties/getter]: numberOfLoops */


// The number of times the audio repeats playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfLoops
func (a_ AudioPlayer) SetNumberOfLoops(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumberOfLoops:"), value)
}/* debug [instance_properties/setter]: numberOfLoops */


// The audio player’s stereo pan position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pan
func (a_ AudioPlayer) Pan() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("pan"))
	return rv
}/* debug [instance_properties/getter]: pan */


// The audio player’s stereo pan position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pan
func (a_ AudioPlayer) SetPan(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPan:"), value)
}/* debug [instance_properties/setter]: pan */


// The audio player’s playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/rate
func (a_ AudioPlayer) Rate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The audio player’s playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/rate
func (a_ AudioPlayer) SetRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// A dictionary that provides information about the player’s audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/settings
func (a_ AudioPlayer) Settings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("settings"))
	return rv
}/* debug [instance_properties/getter]: settings */


// The URL of the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/url
func (a_ AudioPlayer) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The audio player’s volume relative to other audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/volume
func (a_ AudioPlayer) Volume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("volume"))
	return rv
}/* debug [instance_properties/getter]: volume */


// The audio player’s volume relative to other audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/volume
func (a_ AudioPlayer) SetVolume(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVolume:"), value)
}/* debug [instance_properties/setter]: volume */


// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayer/ismeteringenabled
func (a_ AudioPlayer) IsMeteringEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isMeteringEnabled"))
	return rv
}/* debug [instance_properties/getter]: isMeteringEnabled */


// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayer/ismeteringenabled
func (a_ AudioPlayer) SetIsMeteringEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsMeteringEnabled:"), value)
}/* debug [instance_properties/setter]: isMeteringEnabled */


// A Boolean value that indicates whether the player is currently playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayer/isplaying
func (a_ AudioPlayer) IsPlaying() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlaying"))
	return rv
}/* debug [instance_properties/getter]: isPlaying */


// A Boolean value that indicates whether the player is currently playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayer/isplaying
func (a_ AudioPlayer) SetIsPlaying(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlaying:"), value)
}/* debug [instance_properties/setter]: isPlaying */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioPlayer */


