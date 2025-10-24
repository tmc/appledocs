// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioRecorder */


/* debug [class_header]: Header for AVAudioRecorder */
// The class instance for the [AudioRecorder] class.
var (
	AudioRecorderClass     _AudioRecorderClass
	AudioRecorderClassOnce sync.Once
)

func getAudioRecorderClass() _AudioRecorderClass {
	AudioRecorderClassOnce.Do(func() {
		AudioRecorderClass = _AudioRecorderClass{objc.GetClass("AVAudioRecorder")}
	})
	return AudioRecorderClass
}

type _AudioRecorderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioRecorder */
// An interface definition for the [AudioRecorder] class.
type IAudioRecorder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioRecorder */
	// properties:
	CurrentTime() float64
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DeviceCurrentTime() float64
	Format() IAVAudioFormat
	MeteringEnabled() bool
	SetMeteringEnabled(value bool)
	Recording() bool
	Settings() foundation.IDictionary
	Url() objc.IObject /* cross-framework: NSURL */
	IsMeteringEnabled() bool
	SetIsMeteringEnabled(value bool)
	IsRecording() bool
	SetIsRecording(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioRecorder */
	// methods:
	AveragePowerForChannel(channelNumber uint) float32
	DeleteRecording() bool
	Pause()
	PeakPowerForChannel(channelNumber uint) float32
	PrepareToRecord() bool
	Record() bool
	RecordAtTime(time float64) bool
	RecordAtTimeForDuration(time float64, duration float64) bool
	RecordForDuration(duration float64) bool
	Stop()
	UpdateMeters()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioRecorder */
// Alloc allocates a new instance without initialization.
func (ac _AudioRecorderClass) Alloc() AudioRecorder {
	rv := objc.Send[AudioRecorder](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioRecorderClass) New() AudioRecorder {
	rv := objc.Send[AudioRecorder](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioRecorder) Init() AudioRecorder {
	rv := objc.Send[AudioRecorder](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioRecorder) Autorelease() AudioRecorder {
	rv := objc.Send[AudioRecorder](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioRecorder creates a new AudioRecorder instance.
func NewAudioRecorder() AudioRecorder {
	return getAudioRecorderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioRecorder */
// An object that records audio data to a file.
//
// Use an audio recorder to: Record audio from the system’s active input device Record for a specified duration or until the user stops recording Pause and resume a recording Access recording-level metering data To record audio in iOS or tvOS, configure your audio session to use the or category.


// An object that records audio data to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder
type AudioRecorder struct {
	objectivec.Object
}

// AudioRecorderFrom constructs a [AudioRecorder] from an unsafe.Pointer.
//
// An object that records audio data to a file.
func AudioRecorderFrom(ptr unsafe.Pointer) AudioRecorder {
	return AudioRecorder{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioRecorder */

// Creates an audio recorder with an audio format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/init(url:format:)
func NewAudioRecorderWithURLFormatError(url objc.IObject /* cross-framework: NSURL */, format IAVAudioFormat, outError objectivec.IObject) AudioRecorder {
	instance := getAudioRecorderClass().Alloc()
	rv := objc.Send[AudioRecorder](instance.ID, objc.Sel("initWithURL:format:error:"), url, format, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioRecorderWithURLFormatError */


// Creates an audio recorder with settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/init(url:settings:)
func NewAudioRecorderWithURLSettingsError(url objc.IObject /* cross-framework: NSURL */, settings foundation.IDictionary, outError objectivec.IObject) AudioRecorder {
	instance := getAudioRecorderClass().Alloc()
	rv := objc.Send[AudioRecorder](instance.ID, objc.Sel("initWithURL:settings:error:"), url, settings, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioRecorderWithURLSettingsError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioRecorder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioRecorder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioRecorder */

// Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/averagePower(forChannel:)
func (a_ AudioRecorder) AveragePowerForChannel(channelNumber uint) float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("averagePowerForChannel:"), channelNumber)
	return rv
}/* debug [instance_methods/method]: AveragePowerForChannel */


// Deletes a recorded audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/deleteRecording()
func (a_ AudioRecorder) DeleteRecording() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("deleteRecording"))
	return rv
}/* debug [instance_methods/method]: DeleteRecording */


// Pauses an audio recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/pause()
func (a_ AudioRecorder) Pause() {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause"))
}/* debug [instance_methods/method]: Pause */


// Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/peakPower(forChannel:)
func (a_ AudioRecorder) PeakPowerForChannel(channelNumber uint) float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("peakPowerForChannel:"), channelNumber)
	return rv
}/* debug [instance_methods/method]: PeakPowerForChannel */


// Creates an audio file and prepares the system for recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/prepareToRecord()
func (a_ AudioRecorder) PrepareToRecord() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prepareToRecord"))
	return rv
}/* debug [instance_methods/method]: PrepareToRecord */


// Starts or resumes audio recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record()
func (a_ AudioRecorder) Record() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("record"))
	return rv
}/* debug [instance_methods/method]: Record */


// Records audio starting at a specific time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(atTime:)
func (a_ AudioRecorder) RecordAtTime(time float64) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("recordAtTime:"), time)
	return rv
}/* debug [instance_methods/method]: RecordAtTime */


// Records audio starting at a specific time for the indicated duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(atTime:forDuration:)
func (a_ AudioRecorder) RecordAtTimeForDuration(time float64, duration float64) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("recordAtTime:forDuration:"), time, duration)
	return rv
}/* debug [instance_methods/method]: RecordAtTimeForDuration */


// Records audio for the indicated duration of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(forDuration:)
func (a_ AudioRecorder) RecordForDuration(duration float64) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("recordForDuration:"), duration)
	return rv
}/* debug [instance_methods/method]: RecordForDuration */


// Stops recording and closes the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/stop()
func (a_ AudioRecorder) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */


// Refreshes the average and peak power values for all channels of an audio recorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/updateMeters()
func (a_ AudioRecorder) UpdateMeters() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateMeters"))
}/* debug [instance_methods/method]: UpdateMeters */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioRecorder */

// The time, in seconds, since the beginning of the recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/currentTime
func (a_ AudioRecorder) CurrentTime() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("currentTime"))
	return rv
}/* debug [instance_properties/getter]: currentTime */


// The delegate object for the audio recorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/delegate
func (a_ AudioRecorder) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object for the audio recorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/delegate
func (a_ AudioRecorder) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The time, in seconds, of the host audio device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/deviceCurrentTime
func (a_ AudioRecorder) DeviceCurrentTime() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("deviceCurrentTime"))
	return rv
}/* debug [instance_properties/getter]: deviceCurrentTime */


// The format of the recorded audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/format
func (a_ AudioRecorder) Format() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/isMeteringEnabled
func (a_ AudioRecorder) MeteringEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("meteringEnabled"))
	return rv
}/* debug [instance_properties/getter]: meteringEnabled */


// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/isMeteringEnabled
func (a_ AudioRecorder) SetMeteringEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMeteringEnabled:"), value)
}/* debug [instance_properties/setter]: meteringEnabled */


// A Boolean value that indicates whether the audio recorder is recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/isRecording
func (a_ AudioRecorder) Recording() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("recording"))
	return rv
}/* debug [instance_properties/getter]: recording */


// The settings that describe the format of the recorded audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/settings
func (a_ AudioRecorder) Settings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("settings"))
	return rv
}/* debug [instance_properties/getter]: settings */


// The URL to which the recorder writes its data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/url
func (a_ AudioRecorder) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiorecorder/ismeteringenabled
func (a_ AudioRecorder) IsMeteringEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isMeteringEnabled"))
	return rv
}/* debug [instance_properties/getter]: isMeteringEnabled */


// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiorecorder/ismeteringenabled
func (a_ AudioRecorder) SetIsMeteringEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsMeteringEnabled:"), value)
}/* debug [instance_properties/setter]: isMeteringEnabled */


// A Boolean value that indicates whether the audio recorder is recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiorecorder/isrecording
func (a_ AudioRecorder) IsRecording() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRecording"))
	return rv
}/* debug [instance_properties/getter]: isRecording */


// A Boolean value that indicates whether the audio recorder is recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiorecorder/isrecording
func (a_ AudioRecorder) SetIsRecording(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRecording:"), value)
}/* debug [instance_properties/setter]: isRecording */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioRecorder */


