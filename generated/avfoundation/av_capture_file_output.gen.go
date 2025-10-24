// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CaptureFileOutput] class.
var (
	CaptureFileOutputClass     _CaptureFileOutputClass
	CaptureFileOutputClassOnce sync.Once
)

func getCaptureFileOutputClass() _CaptureFileOutputClass {
	CaptureFileOutputClassOnce.Do(func() {
		CaptureFileOutputClass = _CaptureFileOutputClass{objc.GetClass("AVCaptureFileOutput")}
	})
	return CaptureFileOutputClass
}

type _CaptureFileOutputClass struct {
	class objc.Class
}





// An interface definition for the [CaptureFileOutput] class.
type ICaptureFileOutput interface {
	ICaptureOutput
	

	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Recording() bool
	RecordingPaused() bool
	MaxRecordedDuration() objc.IObject /* cross-framework: Time */
	SetMaxRecordedDuration(value objc.IObject /* cross-framework: Time */)
	MaxRecordedFileSize() int64
	SetMaxRecordedFileSize(value int64)
	MinFreeDiskSpaceLimit() int64
	SetMinFreeDiskSpaceLimit(value int64)
	OutputFileURL() objc.IObject /* cross-framework: NSURL */
	RecordedDuration() objc.IObject /* cross-framework: Time */
	RecordedFileSize() int64
	IsRecording() bool
	SetIsRecording(value bool)
	IsRecordingPaused() bool
	SetIsRecordingPaused(value bool)


	

	// methods:
	PauseRecording()
	ResumeRecording()
	StartRecordingToOutputFileURLRecordingDelegate(outputFileURL objc.IObject /* cross-framework: NSURL */, delegate unsafe.Pointer)
	StopRecording()


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureFileOutputClass) Alloc() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureFileOutputClass) New() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureFileOutput) Init() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureFileOutput) Autorelease() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureFileOutput creates a new CaptureFileOutput instance.
func NewCaptureFileOutput() CaptureFileOutput {
	return getCaptureFileOutputClass().New()
}





// The abstract superclass for capture outputs that can record captured data to a file.


// The abstract superclass for capture outputs that can record captured data to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput
type CaptureFileOutput struct {
	CaptureOutput
}

// CaptureFileOutputFrom constructs a [CaptureFileOutput] from an unsafe.Pointer.
//
// The abstract superclass for capture outputs that can record captured data to a file.
func CaptureFileOutputFrom(ptr unsafe.Pointer) CaptureFileOutput {
	return CaptureFileOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}




















// Pauses recording to the current output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/pauseRecording()
func (c_ CaptureFileOutput) PauseRecording() {
	objc.Send[objc.ID](c_.ID, objc.Sel("pauseRecording"))
}


// Resumes recording to the current output file after it was previously paused using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/resumeRecording()
func (c_ CaptureFileOutput) ResumeRecording() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resumeRecording"))
}


// Starts recording media to the specified output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/startRecording(to:recordingDelegate:)
func (c_ CaptureFileOutput) StartRecordingToOutputFileURLRecordingDelegate(outputFileURL objc.IObject /* cross-framework: NSURL */, delegate unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startRecordingToOutputFileURL:recordingDelegate:"), outputFileURL, delegate)
}


// Tells the receiver to stop recording to the current file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/stopRecording()
func (c_ CaptureFileOutput) StopRecording() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopRecording"))
}







// The delegate object for the capture file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/delegate
func (c_ CaptureFileOutput) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object for the capture file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/delegate
func (c_ CaptureFileOutput) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// Indicates whether recording is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/isRecording
func (c_ CaptureFileOutput) Recording() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("recording"))
	return rv
}


// Indicates whether recording to the current output file is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/isRecordingPaused
func (c_ CaptureFileOutput) RecordingPaused() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("recordingPaused"))
	return rv
}


// The longest duration allowed for the recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/maxRecordedDuration
func (c_ CaptureFileOutput) MaxRecordedDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("maxRecordedDuration"))
	return rv
}


// The longest duration allowed for the recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/maxRecordedDuration
func (c_ CaptureFileOutput) SetMaxRecordedDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxRecordedDuration:"), value)
}


// The maximum size, in bytes, of the data that should be recorded by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/maxRecordedFileSize
func (c_ CaptureFileOutput) MaxRecordedFileSize() int64 {
	rv := objc.Send[int64](c_.ID, objc.Sel("maxRecordedFileSize"))
	return rv
}


// The maximum size, in bytes, of the data that should be recorded by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/maxRecordedFileSize
func (c_ CaptureFileOutput) SetMaxRecordedFileSize(value int64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxRecordedFileSize:"), value)
}


// The minimum amount of free space, in bytes, required for recording to continue on a given volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/minFreeDiskSpaceLimit
func (c_ CaptureFileOutput) MinFreeDiskSpaceLimit() int64 {
	rv := objc.Send[int64](c_.ID, objc.Sel("minFreeDiskSpaceLimit"))
	return rv
}


// The minimum amount of free space, in bytes, required for recording to continue on a given volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/minFreeDiskSpaceLimit
func (c_ CaptureFileOutput) SetMinFreeDiskSpaceLimit(value int64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFreeDiskSpaceLimit:"), value)
}


// The URL to which output is directed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/outputFileURL
func (c_ CaptureFileOutput) OutputFileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("outputFileURL"))
	return rv
}


// Indicates the duration of the media recorded to the current output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/recordedDuration
func (c_ CaptureFileOutput) RecordedDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("recordedDuration"))
	return rv
}


// Indicates the size, in bytes, of the data recorded to the current output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/recordedFileSize
func (c_ CaptureFileOutput) RecordedFileSize() int64 {
	rv := objc.Send[int64](c_.ID, objc.Sel("recordedFileSize"))
	return rv
}


// Indicates whether recording is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/isrecording
func (c_ CaptureFileOutput) IsRecording() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRecording"))
	return rv
}


// Indicates whether recording is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/isrecording
func (c_ CaptureFileOutput) SetIsRecording(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRecording:"), value)
}


// Indicates whether recording to the current output file is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/isrecordingpaused
func (c_ CaptureFileOutput) IsRecordingPaused() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRecordingPaused"))
	return rv
}


// Indicates whether recording to the current output file is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/isrecordingpaused
func (c_ CaptureFileOutput) SetIsRecordingPaused(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRecordingPaused:"), value)
}








