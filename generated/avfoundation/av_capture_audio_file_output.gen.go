// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureAudioFileOutput] class.
var (
	CaptureAudioFileOutputClass     _CaptureAudioFileOutputClass
	CaptureAudioFileOutputClassOnce sync.Once
)

func getCaptureAudioFileOutputClass() _CaptureAudioFileOutputClass {
	CaptureAudioFileOutputClassOnce.Do(func() {
		CaptureAudioFileOutputClass = _CaptureAudioFileOutputClass{objc.GetClass("AVCaptureAudioFileOutput")}
	})
	return CaptureAudioFileOutputClass
}

type _CaptureAudioFileOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureAudioFileOutput] class.
type ICaptureAudioFileOutput interface {
	ICaptureFileOutput
	AudioSettings() string
	SetAudioSettings(value string)
	Metadata() AVMetadataItem
	SetMetadata(value IAVMetadataItem)
}

// A capture output that records audio and saves the recorded audio to a file.
//
// implements the complete file recording interface declared by for writing media data to audio files. In addition, you can configure options specific to the audio file formats, including writing metadata collections to each file and specifying audio encoding options. does not, however, support —use instead.


// A capture output that records audio and saves the recorded audio to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput

type CaptureAudioFileOutput struct {
	CaptureFileOutput
}

// CaptureAudioFileOutputFrom constructs a [CaptureAudioFileOutput] from an unsafe.Pointer.
//
// A capture output that records audio and saves the recorded audio to a file.
func CaptureAudioFileOutputFrom(ptr unsafe.Pointer) CaptureAudioFileOutput {
	return CaptureAudioFileOutput{
		CaptureFileOutput: CaptureFileOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureAudioFileOutputClass) Alloc() CaptureAudioFileOutput {
	rv := objc.Send[CaptureAudioFileOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureAudioFileOutputClass) New() CaptureAudioFileOutput {
	rv := objc.Send[CaptureAudioFileOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureAudioFileOutput) Init() CaptureAudioFileOutput {
	rv := objc.Send[CaptureAudioFileOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureAudioFileOutput) Autorelease() CaptureAudioFileOutput {
	rv := objc.Send[CaptureAudioFileOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureAudioFileOutput creates a new CaptureAudioFileOutput instance.
func NewCaptureAudioFileOutput() CaptureAudioFileOutput {
	return getCaptureAudioFileOutputClass().New()
}



// The settings used to decode or re-encode audio before it is output by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/audiosettings

func (c_ CaptureAudioFileOutput) AudioSettings() string {
	rv := objc.Send[string](c_.ID, objc.Sel("audioSettings"))
	return rv
}


// The settings used to decode or re-encode audio before it is output by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/audiosettings

func (c_ CaptureAudioFileOutput) SetAudioSettings(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSettings:"), objc.String(value))
}


// A collection of metadata to be written to the receiver’s output files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/metadata

func (c_ CaptureAudioFileOutput) Metadata() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}


// A collection of metadata to be written to the receiver’s output files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiofileoutput/metadata

func (c_ CaptureAudioFileOutput) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}



