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
}

// A capture output that records audio and saves the recorded audio to a file.
//
// implements the complete file recording interface declared by for writing media data to audio files. In addition, you can configure options specific to the audio file formats, including writing metadata collections to each file and specifying audio encoding options. does not, however, support —use instead.
//
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




