// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureAudioFileOutput */


/* debug [class_header]: Header for AVCaptureAudioFileOutput */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureAudioFileOutput */
// An interface definition for the [CaptureAudioFileOutput] class.
type ICaptureAudioFileOutput interface {
	ICaptureFileOutput
	
/* debug [class_interface_properties]: Properties for CaptureAudioFileOutput */
	// properties:
	AudioSettings() foundation.IDictionary
	SetAudioSettings(value foundation.IDictionary)
	Metadata() []MetadataItem
	SetMetadata(value []MetadataItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureAudioFileOutput */
	// methods:
	StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate(outputFileURL objc.IObject /* cross-framework: NSURL */, fileType FileType /* typedef */, delegate unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureAudioFileOutput */
// Alloc allocates a new instance without initialization.
func (cc _CaptureAudioFileOutputClass) Alloc() CaptureAudioFileOutput {
	rv := objc.Send[CaptureAudioFileOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureAudioFileOutput */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureAudioFileOutput */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureAudioFileOutput */

// Returns an array containing UTIs identifying the file types can write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/availableOutputFileTypes()
func (cc _CaptureAudioFileOutputClass) AvailableOutputFileTypes() []string {
	rv := objc.Send[[]string](objc.ID(cc.class), objc.Sel("availableOutputFileTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AvailableOutputFileTypes) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureAudioFileOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureAudioFileOutput */

// Tells the receiver to start recording to a new file of the specified format, and specifies a delegate that will be notified when recording is finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/startRecording(to:outputFileType:recordingDelegate:)
func (c_ CaptureAudioFileOutput) StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate(outputFileURL objc.IObject /* cross-framework: NSURL */, fileType FileType /* typedef */, delegate unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startRecordingToOutputFileURL:outputFileType:recordingDelegate:"), outputFileURL, fileType, delegate)
}/* debug [instance_methods/method]: StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureAudioFileOutput */

// The settings used to decode or re-encode audio before it is output by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/audioSettings
func (c_ CaptureAudioFileOutput) AudioSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("audioSettings"))
	return rv
}/* debug [instance_properties/getter]: audioSettings */


// The settings used to decode or re-encode audio before it is output by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/audioSettings
func (c_ CaptureAudioFileOutput) SetAudioSettings(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSettings:"), value)
}/* debug [instance_properties/setter]: audioSettings */


// A collection of metadata to be written to the receiver’s output files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/metadata
func (c_ CaptureAudioFileOutput) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// A collection of metadata to be written to the receiver’s output files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/metadata
func (c_ CaptureAudioFileOutput) SetMetadata(value []MetadataItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), nsArray)
}/* debug [instance_properties/setter]: metadata */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureAudioFileOutput */


