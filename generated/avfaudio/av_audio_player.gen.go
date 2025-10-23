// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioPlayer] class.
type IAudioPlayer interface {
	objectivec.IObject
	// properties:
	ChannelAssignments() []AudioSessionChannelDescription /* primitive/slice/pointer. */
	SetChannelAssignments(value []AudioSessionChannelDescription /* primitive/slice/pointer. */)
	CurrentDevice() string /* primitive/slice/pointer. */
	SetCurrentDevice(value string /* primitive/slice/pointer. */)
	CurrentTime() foundation.TimeInterval /* not a class type */
	SetCurrentTime(value foundation.TimeInterval /* not a class type */)
	Data() foundation.objc.IObject /* cross-framework: NSData */
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DeviceCurrentTime() foundation.TimeInterval /* not a class type */
	Duration() foundation.TimeInterval /* not a class type */
	EnableRate() bool /* primitive/slice/pointer. */
	SetEnableRate(value bool /* primitive/slice/pointer. */)
	Format() IAVAudioFormat
	IntendedSpatialExperience() objc.IObject /* cross-framework: SpatialAudioExperience */
	SetIntendedSpatialExperience(value objc.IObject /* cross-framework: SpatialAudioExperience */)
	MeteringEnabled() bool /* primitive/slice/pointer. */
	SetMeteringEnabled(value bool /* primitive/slice/pointer. */)
	Playing() bool /* primitive/slice/pointer. */
	NumberOfChannels() uint /* primitive/slice/pointer. */
	NumberOfLoops() int /* primitive/slice/pointer. */
	SetNumberOfLoops(value int /* primitive/slice/pointer. */)
	Pan() float32 /* primitive/slice/pointer. */
	SetPan(value float32 /* primitive/slice/pointer. */)
	Rate() float32 /* primitive/slice/pointer. */
	SetRate(value float32 /* primitive/slice/pointer. */)
	Settings() foundation.IDictionary /* already interface */
	Url() foundation.objc.IObject /* cross-framework: URL */
	Volume() float32 /* primitive/slice/pointer. */
	SetVolume(value float32 /* primitive/slice/pointer. */)
	IsMeteringEnabled() bool /* primitive/slice/pointer. */
	SetIsMeteringEnabled(value bool /* primitive/slice/pointer. */)
	IsPlaying() bool /* primitive/slice/pointer. */
	SetIsPlaying(value bool /* primitive/slice/pointer. */)
	// methods:
	AveragePowerForChannel(channelNumber uint /* primitive/slice/pointer. */) float32 /* primitive/slice/pointer. */
	Pause()
	PeakPowerForChannel(channelNumber uint /* primitive/slice/pointer. */) float32 /* primitive/slice/pointer. */
	Play() bool /* primitive/slice/pointer. */
	PlayAtTime(time foundation.TimeInterval /* not a class type */) bool /* primitive/slice/pointer. */
	PrepareToPlay() bool /* primitive/slice/pointer. */
	SetVolumeFadeDuration(volume float32 /* primitive/slice/pointer. */, duration foundation.TimeInterval /* not a class type */)
	Stop()
	UpdateMeters()
}

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

// Alloc allocates a new instance without initialization.
func (ac _AudioPlayerClass) Alloc() AudioPlayer {
	rv := objc.Send[AudioPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a player to play audio from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:)
func NewAudioPlayerWithContentsOfURLError(url foundation.objc.IObject /* cross-framework URL */, outError unsafe.Pointer) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, outError)
	rv.Autorelease()
	return rv
}


// Creates a player to play audio from a file of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:fileTypeHint:)
func NewAudioPlayerWithContentsOfURLFileTypeHintError(url foundation.objc.IObject /* cross-framework URL */, utiString string /* primitive/slice/pointer. */, outError unsafe.Pointer) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithContentsOfURL:fileTypeHint:error:"), url, objc.String(utiString), outError)
	rv.Autorelease()
	return rv
}


// Creates a player to play in-memory audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:)
func NewAudioPlayerWithDataError(data foundation.objc.IObject /* cross-framework NSData */, outError unsafe.Pointer) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithData:error:"), data, outError)
	rv.Autorelease()
	return rv
}


// Creates a player to play in-memory audio data of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:fileTypeHint:)
func NewAudioPlayerWithDataFileTypeHintError(data foundation.objc.IObject /* cross-framework NSData */, utiString string /* primitive/slice/pointer. */, outError unsafe.Pointer) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithData:fileTypeHint:error:"), data, objc.String(utiString), outError)
	rv.Autorelease()
	return rv
}



// Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/averagePower(forChannel:)
func (a_ AudioPlayer) AveragePowerForChannel(channelNumber uint /* primitive/slice/pointer. */) float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("averagePowerForChannel:"), channelNumber)
	return rv
}


// Pauses audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pause()
func (a_ AudioPlayer) Pause() {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause"))
}


// Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/peakPower(forChannel:)
func (a_ AudioPlayer) PeakPowerForChannel(channelNumber uint /* primitive/slice/pointer. */) float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("peakPowerForChannel:"), channelNumber)
	return rv
}


// Plays audio asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/play()
func (a_ AudioPlayer) Play() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("play"))
	return rv
}


// Plays audio asynchronously, starting at a specified point in the audio output device’s timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/play(atTime:)
func (a_ AudioPlayer) PlayAtTime(time foundation.TimeInterval /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("playAtTime:"), time)
	return rv
}


// Prepares the player for audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/prepareToPlay()
func (a_ AudioPlayer) PrepareToPlay() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("prepareToPlay"))
	return rv
}


// Changes the audio player’s volume over a duration of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/setVolume(_:fadeDuration:)
func (a_ AudioPlayer) SetVolumeFadeDuration(volume float32 /* primitive/slice/pointer. */, duration foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVolume:fadeDuration:"), volume, duration)
}


// Stops playback and undoes the setup the system requires for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/stop()
func (a_ AudioPlayer) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}


// Refreshes the average and peak power values for all channels of an audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/updateMeters()
func (a_ AudioPlayer) UpdateMeters() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateMeters"))
}


// An array of channel descriptions for the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/channelAssignments
func (a_ AudioPlayer) ChannelAssignments() []AudioSessionChannelDescription /* primitive/slice/pointer. */ {
	rv := objc.Send[[]AudioSessionChannelDescription](a_.ID, objc.Sel("channelAssignments"))
	return rv
}


// An array of channel descriptions for the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/channelAssignments
func (a_ AudioPlayer) SetChannelAssignments(value []AudioSessionChannelDescription /* primitive/slice/pointer. */) {
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
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelAssignments:"), nsArray)
}


// The unique identifier of the current audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentDevice
func (a_ AudioPlayer) CurrentDevice() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("currentDevice"))
	return rv
}


// The unique identifier of the current audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentDevice
func (a_ AudioPlayer) SetCurrentDevice(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentDevice:"), objc.String(value))
}


// The current playback time, in seconds, within the audio timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentTime
func (a_ AudioPlayer) CurrentTime() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("currentTime"))
	return rv
}


// The current playback time, in seconds, within the audio timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentTime
func (a_ AudioPlayer) SetCurrentTime(value foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentTime:"), value)
}


// The audio data associated with the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/data
func (a_ AudioPlayer) Data() foundation.objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("data"))
	return rv
}


// The delegate object for the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/delegate
func (a_ AudioPlayer) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object for the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/delegate
func (a_ AudioPlayer) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// The time value, in seconds, of the audio output device’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/deviceCurrentTime
func (a_ AudioPlayer) DeviceCurrentTime() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("deviceCurrentTime"))
	return rv
}


// The total duration, in seconds, of the player’s audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/duration
func (a_ AudioPlayer) Duration() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("duration"))
	return rv
}


// A Boolean value that indicates whether you can adjust the playback rate of the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/enableRate
func (a_ AudioPlayer) EnableRate() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("enableRate"))
	return rv
}


// A Boolean value that indicates whether you can adjust the playback rate of the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/enableRate
func (a_ AudioPlayer) SetEnableRate(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnableRate:"), value)
}


// The format of the player’s audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/format
func (a_ AudioPlayer) Format() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("format"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/intendedSpatialExperience-6py9z
func (a_ AudioPlayer) IntendedSpatialExperience() objc.IObject /* cross-framework: SpatialAudioExperience */ {
	rv := objc.Send[SpatialAudioExperience](a_.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/intendedSpatialExperience-6py9z
func (a_ AudioPlayer) SetIntendedSpatialExperience(value objc.IObject /* cross-framework: SpatialAudioExperience */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIntendedSpatialExperience:"), value)
}


// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isMeteringEnabled
func (a_ AudioPlayer) MeteringEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("meteringEnabled"))
	return rv
}


// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isMeteringEnabled
func (a_ AudioPlayer) SetMeteringEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMeteringEnabled:"), value)
}


// A Boolean value that indicates whether the player is currently playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isPlaying
func (a_ AudioPlayer) Playing() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("playing"))
	return rv
}


// The number of audio channels in the player’s audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfChannels
func (a_ AudioPlayer) NumberOfChannels() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](a_.ID, objc.Sel("numberOfChannels"))
	return rv
}


// The number of times the audio repeats playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfLoops
func (a_ AudioPlayer) NumberOfLoops() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](a_.ID, objc.Sel("numberOfLoops"))
	return rv
}


// The number of times the audio repeats playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfLoops
func (a_ AudioPlayer) SetNumberOfLoops(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumberOfLoops:"), value)
}


// The audio player’s stereo pan position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pan
func (a_ AudioPlayer) Pan() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("pan"))
	return rv
}


// The audio player’s stereo pan position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pan
func (a_ AudioPlayer) SetPan(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPan:"), value)
}


// The audio player’s playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/rate
func (a_ AudioPlayer) Rate() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("rate"))
	return rv
}


// The audio player’s playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/rate
func (a_ AudioPlayer) SetRate(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}


// A dictionary that provides information about the player’s audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/settings
func (a_ AudioPlayer) Settings() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("settings"))
	return rv
}


// The URL of the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/url
func (a_ AudioPlayer) Url() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("url"))
	return rv
}


// The audio player’s volume relative to other audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/volume
func (a_ AudioPlayer) Volume() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("volume"))
	return rv
}


// The audio player’s volume relative to other audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/volume
func (a_ AudioPlayer) SetVolume(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVolume:"), value)
}


// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayer/ismeteringenabled
func (a_ AudioPlayer) IsMeteringEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isMeteringEnabled"))
	return rv
}


// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayer/ismeteringenabled
func (a_ AudioPlayer) SetIsMeteringEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsMeteringEnabled:"), value)
}


// A Boolean value that indicates whether the player is currently playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayer/isplaying
func (a_ AudioPlayer) IsPlaying() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlaying"))
	return rv
}


// A Boolean value that indicates whether the player is currently playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayer/isplaying
func (a_ AudioPlayer) SetIsPlaying(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlaying:"), value)
}


