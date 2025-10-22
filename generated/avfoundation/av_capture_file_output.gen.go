// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	StartRecordingToOutputFileURLRecordingDelegate(outputFileURL foundation.IURL, delegate objectivec.IObject)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsRecording() bool
	SetIsRecording(value bool)
	IsRecordingPaused() bool
	SetIsRecordingPaused(value bool)
	MaxRecordedDuration() unsafe.Pointer
	SetMaxRecordedDuration(value unsafe.Pointer)
	MaxRecordedFileSize() unsafe.Pointer
	SetMaxRecordedFileSize(value unsafe.Pointer)
	MinFreeDiskSpaceLimit() unsafe.Pointer
	SetMinFreeDiskSpaceLimit(value unsafe.Pointer)
	OutputFileURL() foundation.URL
	SetOutputFileURL(value foundation.IURL)
	RecordedDuration() unsafe.Pointer
	SetRecordedDuration(value unsafe.Pointer)
	RecordedFileSize() unsafe.Pointer
	SetRecordedFileSize(value unsafe.Pointer)
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

// Alloc allocates a new instance without initialization.
func (cc _CaptureFileOutputClass) Alloc() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Starts recording media to the specified output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/startRecording(to:recordingDelegate:)

func (c_ CaptureFileOutput) StartRecordingToOutputFileURLRecordingDelegate(outputFileURL foundation.IURL, delegate objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startRecordingToOutputFileURL:recordingDelegate:"), outputFileURL, delegate)
}


// The delegate object for the capture file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/delegate

func (c_ CaptureFileOutput) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object for the capture file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/delegate

func (c_ CaptureFileOutput) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
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


// The longest duration allowed for the recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/maxrecordedduration

func (c_ CaptureFileOutput) MaxRecordedDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maxRecordedDuration"))
	return rv
}


// The longest duration allowed for the recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/maxrecordedduration

func (c_ CaptureFileOutput) SetMaxRecordedDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxRecordedDuration:"), value)
}


// The maximum size, in bytes, of the data that should be recorded by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/maxrecordedfilesize

func (c_ CaptureFileOutput) MaxRecordedFileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maxRecordedFileSize"))
	return rv
}


// The maximum size, in bytes, of the data that should be recorded by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/maxrecordedfilesize

func (c_ CaptureFileOutput) SetMaxRecordedFileSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxRecordedFileSize:"), value)
}


// The minimum amount of free space, in bytes, required for recording to continue on a given volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/minfreediskspacelimit

func (c_ CaptureFileOutput) MinFreeDiskSpaceLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("minFreeDiskSpaceLimit"))
	return rv
}


// The minimum amount of free space, in bytes, required for recording to continue on a given volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/minfreediskspacelimit

func (c_ CaptureFileOutput) SetMinFreeDiskSpaceLimit(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFreeDiskSpaceLimit:"), value)
}


// The URL to which output is directed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/outputfileurl

func (c_ CaptureFileOutput) OutputFileURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("outputFileURL"))
	return rv
}


// The URL to which output is directed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/outputfileurl

func (c_ CaptureFileOutput) SetOutputFileURL(value foundation.IURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputFileURL:"), value)
}


// Indicates the duration of the media recorded to the current output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/recordedduration

func (c_ CaptureFileOutput) RecordedDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordedDuration"))
	return rv
}


// Indicates the duration of the media recorded to the current output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/recordedduration

func (c_ CaptureFileOutput) SetRecordedDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordedDuration:"), value)
}


// Indicates the size, in bytes, of the data recorded to the current output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/recordedfilesize

func (c_ CaptureFileOutput) RecordedFileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordedFileSize"))
	return rv
}


// Indicates the size, in bytes, of the data recorded to the current output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/recordedfilesize

func (c_ CaptureFileOutput) SetRecordedFileSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordedFileSize:"), value)
}



