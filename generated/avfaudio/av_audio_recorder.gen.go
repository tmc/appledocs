// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioRecorder] class.
type IAudioRecorder interface {
	objectivec.IObject
	AveragePowerForChannel(channelNumber uint) unsafe.Pointer
	DeleteRecording() bool
	Pause()
	PeakPowerForChannel(channelNumber uint) unsafe.Pointer
	PrepareToRecord() bool
	Record() bool
	RecordAtTime(time foundation.ITimeInterval) bool
	RecordAtTimeForDuration(time foundation.ITimeInterval, duration foundation.ITimeInterval) bool
	RecordForDuration(duration foundation.ITimeInterval) bool
	Stop()
	UpdateMeters()
}

// An object that records audio data to a file.
//
// Use an audio recorder to: Record audio from the system’s active input device Record for a specified duration or until the user stops recording Pause and resume a recording Access recording-level metering data To record audio in iOS or tvOS, configure your audio session to use the or category.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioRecorderClass) Alloc() AudioRecorder {
	rv := objc.Send[AudioRecorder](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an audio recorder with an audio format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/init(url:format:)
func NewAudioRecorderWithURLFormatError(url foundation.IURL, format AVAudioFormat, outError unsafe.Pointer) AudioRecorder {
	instance := getAudioRecorderClass().Alloc()
	rv := objc.Send[AudioRecorder](instance.ID, objc.Sel("initWithURL:format:error:"), url, format, outError)
	rv.Autorelease()
	return rv
}



// Creates an audio recorder with settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/init(url:settings:)
func NewAudioRecorderWithURLSettingsError(url foundation.IURL, settings unsafe.Pointer, outError unsafe.Pointer) AudioRecorder {
	instance := getAudioRecorderClass().Alloc()
	rv := objc.Send[AudioRecorder](instance.ID, objc.Sel("initWithURL:settings:error:"), url, settings, outError)
	rv.Autorelease()
	return rv
}


// Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/averagePower(forChannel:)
func (a_ AudioRecorder) AveragePowerForChannel(channelNumber uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("averagePowerForChannel:"), channelNumber)
	return rv
}

// Deletes a recorded audio file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/deleteRecording()
func (a_ AudioRecorder) DeleteRecording() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("deleteRecording"))
	return rv
}

// Pauses an audio recording.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/pause()
func (a_ AudioRecorder) Pause() {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause"))
}

// Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/peakPower(forChannel:)
func (a_ AudioRecorder) PeakPowerForChannel(channelNumber uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("peakPowerForChannel:"), channelNumber)
	return rv
}

// Creates an audio file and prepares the system for recording.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/prepareToRecord()
func (a_ AudioRecorder) PrepareToRecord() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prepareToRecord"))
	return rv
}

// Starts or resumes audio recording.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record()
func (a_ AudioRecorder) Record() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("record"))
	return rv
}

// Records audio starting at a specific time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(atTime:)
func (a_ AudioRecorder) RecordAtTime(time foundation.ITimeInterval) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("recordAtTime:"), time)
	return rv
}

// Records audio starting at a specific time for the indicated duration.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(atTime:forDuration:)
func (a_ AudioRecorder) RecordAtTimeForDuration(time foundation.ITimeInterval, duration foundation.ITimeInterval) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("recordAtTime:forDuration:"), time, duration)
	return rv
}

// Records audio for the indicated duration of time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(forDuration:)
func (a_ AudioRecorder) RecordForDuration(duration foundation.ITimeInterval) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("recordForDuration:"), duration)
	return rv
}

// Stops recording and closes the audio file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/stop()
func (a_ AudioRecorder) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}

// Refreshes the average and peak power values for all channels of an audio recorder.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/updateMeters()
func (a_ AudioRecorder) UpdateMeters() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateMeters"))
}

// An array of channel descriptions associated with the audio recorder.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/channelAssignments
func (a_ AudioRecorder) ChannelAssignments() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](a_.ID, objc.Sel("channelAssignments"))
	return rv
}


// SetChannelAssignments sets the value of the channelAssignments property.
// An array of channel descriptions associated with the audio recorder.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/channelAssignments
func (a_ AudioRecorder) SetChannelAssignments(value []unsafe.IPointer) {
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

// The time, in seconds, since the beginning of the recording.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/currentTime
func (a_ AudioRecorder) CurrentTime() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("currentTime"))
	return rv
}

// The delegate object for the audio recorder.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/delegate
func (a_ AudioRecorder) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for the audio recorder.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/delegate
func (a_ AudioRecorder) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// The time, in seconds, of the host audio device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/deviceCurrentTime
func (a_ AudioRecorder) DeviceCurrentTime() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("deviceCurrentTime"))
	return rv
}

// The format of the recorded audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/format
func (a_ AudioRecorder) Format() AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a_.ID, objc.Sel("format"))
	return rv
}

// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/isMeteringEnabled
func (a_ AudioRecorder) MeteringEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("meteringEnabled"))
	return rv
}


// SetMeteringEnabled sets the value of the meteringEnabled property.
// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/isMeteringEnabled
func (a_ AudioRecorder) SetMeteringEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMeteringEnabled:"), value)
}

// A Boolean value that indicates whether the audio recorder is recording.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/isRecording
func (a_ AudioRecorder) Recording() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("recording"))
	return rv
}

// The settings that describe the format of the recorded audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/settings
func (a_ AudioRecorder) Settings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("settings"))
	return rv
}

// The URL to which the recorder writes its data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/url
func (a_ AudioRecorder) Url() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("url"))
	return rv
}

// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiorecorder/ismeteringenabled
func (a_ AudioRecorder) IsMeteringEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isMeteringEnabled"))
	return rv
}


// SetIsMeteringEnabled sets the value of the isMeteringEnabled property.
// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiorecorder/ismeteringenabled
func (a_ AudioRecorder) SetIsMeteringEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsMeteringEnabled:"), value)
}

// A Boolean value that indicates whether the audio recorder is recording.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiorecorder/isrecording
func (a_ AudioRecorder) IsRecording() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRecording"))
	return rv
}


// SetIsRecording sets the value of the isRecording property.
// A Boolean value that indicates whether the audio recorder is recording.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiorecorder/isrecording
func (a_ AudioRecorder) SetIsRecording(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRecording:"), value)
}


