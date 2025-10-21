// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	AveragePowerForChannel(channelNumber uint) unsafe.Pointer
	Pause()
	PeakPowerForChannel(channelNumber uint) unsafe.Pointer
	Play() bool
	PlayAtTime(time TimeInterval) bool
	PrepareToPlay() bool
	SetVolumeFadeDuration(volume unsafe.Pointer, duration TimeInterval)
	Stop()
	UpdateMeters()
}

// An object that plays audio data from a file or buffer.
//
// Use an audio player to: Play audio of any duration from a file or buffer Control the volume, panning, rate, and looping behavior of the played audio Access playback-level metering data Play multiple sounds simultaneously by synchronizing the playback of multiple players For more information about preparing your app to play audio, see .
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:)
func NewAudioPlayerWithContentsOfURLError(url unsafe.Pointer, outError unsafe.Pointer) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, outError)
	rv.Autorelease()
	return rv
}



// Creates a player to play audio from a file of a particular type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:fileTypeHint:)
func NewAudioPlayerWithContentsOfURLFileTypeHintError(url unsafe.Pointer, utiString string, outError unsafe.Pointer) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithContentsOfURL:fileTypeHint:error:"), url, objc.String(utiString), outError)
	rv.Autorelease()
	return rv
}



// Creates a player to play in-memory audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:)
func NewAudioPlayerWithDataError(data unsafe.Pointer, outError unsafe.Pointer) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithData:error:"), data, outError)
	rv.Autorelease()
	return rv
}



// Creates a player to play in-memory audio data of a particular type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:fileTypeHint:)
func NewAudioPlayerWithDataFileTypeHintError(data unsafe.Pointer, utiString string, outError unsafe.Pointer) AudioPlayer {
	instance := getAudioPlayerClass().Alloc()
	rv := objc.Send[AudioPlayer](instance.ID, objc.Sel("initWithData:fileTypeHint:error:"), data, objc.String(utiString), outError)
	rv.Autorelease()
	return rv
}


// Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/averagePower(forChannel:)
func (a_ AudioPlayer) AveragePowerForChannel(channelNumber uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("averagePowerForChannel:"), channelNumber)
	return rv
}

// Pauses audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pause()
func (a_ AudioPlayer) Pause() {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause"))
}

// Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/peakPower(forChannel:)
func (a_ AudioPlayer) PeakPowerForChannel(channelNumber uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("peakPowerForChannel:"), channelNumber)
	return rv
}

// Plays audio asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/play()
func (a_ AudioPlayer) Play() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("play"))
	return rv
}

// Plays audio asynchronously, starting at a specified point in the audio output device’s timeline.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/play(atTime:)
func (a_ AudioPlayer) PlayAtTime(time TimeInterval) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playAtTime:"), time)
	return rv
}

// Prepares the player for audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/prepareToPlay()
func (a_ AudioPlayer) PrepareToPlay() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prepareToPlay"))
	return rv
}

// Changes the audio player’s volume over a duration of time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/setVolume(_:fadeDuration:)
func (a_ AudioPlayer) SetVolumeFadeDuration(volume unsafe.Pointer, duration TimeInterval) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVolume:fadeDuration:"), volume, duration)
}

// Stops playback and undoes the setup the system requires for playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/stop()
func (a_ AudioPlayer) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}

// Refreshes the average and peak power values for all channels of an audio player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/updateMeters()
func (a_ AudioPlayer) UpdateMeters() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateMeters"))
}

// An array of channel descriptions for the audio player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/channelAssignments
func (a_ AudioPlayer) ChannelAssignments() []AVAudioSessionChannelDescription {
	rv := objc.Send[[]AVAudioSessionChannelDescription](a_.ID, objc.Sel("channelAssignments"))
	return rv
}


// SetChannelAssignments sets the value of the channelAssignments property.
// An array of channel descriptions for the audio player.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/channelAssignments
func (a_ AudioPlayer) SetChannelAssignments(value []AVAudioSessionChannelDescription) {
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
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentDevice
func (a_ AudioPlayer) CurrentDevice() string {
	rv := objc.Send[string](a_.ID, objc.Sel("currentDevice"))
	return rv
}


// SetCurrentDevice sets the value of the currentDevice property.
// The unique identifier of the current audio player.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentDevice
func (a_ AudioPlayer) SetCurrentDevice(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentDevice:"), objc.String(value))
}

// The current playback time, in seconds, within the audio timeline.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentTime
func (a_ AudioPlayer) CurrentTime() TimeInterval {
	rv := objc.Send[TimeInterval](a_.ID, objc.Sel("currentTime"))
	return rv
}


// SetCurrentTime sets the value of the currentTime property.
// The current playback time, in seconds, within the audio timeline.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentTime
func (a_ AudioPlayer) SetCurrentTime(value TimeInterval) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentTime:"), value)
}

// The audio data associated with the player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/data
func (a_ AudioPlayer) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("data"))
	return rv
}

// The delegate object for the audio player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/delegate
func (a_ AudioPlayer) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for the audio player.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/delegate
func (a_ AudioPlayer) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// The time value, in seconds, of the audio output device’s clock.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/deviceCurrentTime
func (a_ AudioPlayer) DeviceCurrentTime() TimeInterval {
	rv := objc.Send[TimeInterval](a_.ID, objc.Sel("deviceCurrentTime"))
	return rv
}

// The total duration, in seconds, of the player’s audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/duration
func (a_ AudioPlayer) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](a_.ID, objc.Sel("duration"))
	return rv
}

// A Boolean value that indicates whether you can adjust the playback rate of the audio player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/enableRate
func (a_ AudioPlayer) EnableRate() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enableRate"))
	return rv
}


// SetEnableRate sets the value of the enableRate property.
// A Boolean value that indicates whether you can adjust the playback rate of the audio player.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/enableRate
func (a_ AudioPlayer) SetEnableRate(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnableRate:"), value)
}

// The format of the player’s audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/format
func (a_ AudioPlayer) Format() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("format"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/intendedSpatialExperience-6py9z
func (a_ AudioPlayer) IntendedSpatialExperience() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}


// SetIntendedSpatialExperience sets the value of the intendedSpatialExperience property.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/intendedSpatialExperience-6py9z
func (a_ AudioPlayer) SetIntendedSpatialExperience(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIntendedSpatialExperience:"), value)
}

// A Boolean value that indicates whether the player is able to generate audio-level metering data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isMeteringEnabled
func (a_ AudioPlayer) MeteringEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("meteringEnabled"))
	return rv
}


// SetMeteringEnabled sets the value of the meteringEnabled property.
// A Boolean value that indicates whether the player is able to generate audio-level metering data.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isMeteringEnabled
func (a_ AudioPlayer) SetMeteringEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMeteringEnabled:"), value)
}

// A Boolean value that indicates whether the player is currently playing audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isPlaying
func (a_ AudioPlayer) Playing() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playing"))
	return rv
}

// The number of audio channels in the player’s audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfChannels
func (a_ AudioPlayer) NumberOfChannels() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("numberOfChannels"))
	return rv
}

// The number of times the audio repeats playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfLoops
func (a_ AudioPlayer) NumberOfLoops() int {
	rv := objc.Send[int](a_.ID, objc.Sel("numberOfLoops"))
	return rv
}


// SetNumberOfLoops sets the value of the numberOfLoops property.
// The number of times the audio repeats playback.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfLoops
func (a_ AudioPlayer) SetNumberOfLoops(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumberOfLoops:"), value)
}

// The audio player’s stereo pan position.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pan
func (a_ AudioPlayer) Pan() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("pan"))
	return rv
}


// SetPan sets the value of the pan property.
// The audio player’s stereo pan position.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pan
func (a_ AudioPlayer) SetPan(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPan:"), value)
}

// The audio player’s playback rate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/rate
func (a_ AudioPlayer) Rate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
// The audio player’s playback rate.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/rate
func (a_ AudioPlayer) SetRate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}

// A dictionary that provides information about the player’s audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/settings
func (a_ AudioPlayer) Settings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("settings"))
	return rv
}

// The URL of the audio file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/url
func (a_ AudioPlayer) Url() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("url"))
	return rv
}

// The audio player’s volume relative to other audio output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/volume
func (a_ AudioPlayer) Volume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("volume"))
	return rv
}


// SetVolume sets the value of the volume property.
// The audio player’s volume relative to other audio output.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/volume
func (a_ AudioPlayer) SetVolume(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVolume:"), value)
}


