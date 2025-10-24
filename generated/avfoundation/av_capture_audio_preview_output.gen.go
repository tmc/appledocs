// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureAudioPreviewOutput */


/* debug [class_header]: Header for AVCaptureAudioPreviewOutput */
// The class instance for the [CaptureAudioPreviewOutput] class.
var (
	CaptureAudioPreviewOutputClass     _CaptureAudioPreviewOutputClass
	CaptureAudioPreviewOutputClassOnce sync.Once
)

func getCaptureAudioPreviewOutputClass() _CaptureAudioPreviewOutputClass {
	CaptureAudioPreviewOutputClassOnce.Do(func() {
		CaptureAudioPreviewOutputClass = _CaptureAudioPreviewOutputClass{objc.GetClass("AVCaptureAudioPreviewOutput")}
	})
	return CaptureAudioPreviewOutputClass
}

type _CaptureAudioPreviewOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureAudioPreviewOutput */
// An interface definition for the [CaptureAudioPreviewOutput] class.
type ICaptureAudioPreviewOutput interface {
	ICaptureOutput
	
/* debug [class_interface_properties]: Properties for CaptureAudioPreviewOutput */
	// properties:
	OutputDeviceUniqueID() objc.IObject /* cross-framework: NSString */
	SetOutputDeviceUniqueID(value objc.IObject /* cross-framework: NSString */)
	Volume() float32
	SetVolume(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureAudioPreviewOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureAudioPreviewOutput */
// Alloc allocates a new instance without initialization.
func (cc _CaptureAudioPreviewOutputClass) Alloc() CaptureAudioPreviewOutput {
	rv := objc.Send[CaptureAudioPreviewOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureAudioPreviewOutputClass) New() CaptureAudioPreviewOutput {
	rv := objc.Send[CaptureAudioPreviewOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureAudioPreviewOutput) Init() CaptureAudioPreviewOutput {
	rv := objc.Send[CaptureAudioPreviewOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureAudioPreviewOutput) Autorelease() CaptureAudioPreviewOutput {
	rv := objc.Send[CaptureAudioPreviewOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureAudioPreviewOutput creates a new CaptureAudioPreviewOutput instance.
func NewCaptureAudioPreviewOutput() CaptureAudioPreviewOutput {
	return getCaptureAudioPreviewOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureAudioPreviewOutput */
// A capture output that provides a preview of the captured audio.


// A capture output that provides a preview of the captured audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput
type CaptureAudioPreviewOutput struct {
	CaptureOutput
}

// CaptureAudioPreviewOutputFrom constructs a [CaptureAudioPreviewOutput] from an unsafe.Pointer.
//
// A capture output that provides a preview of the captured audio.
func CaptureAudioPreviewOutputFrom(ptr unsafe.Pointer) CaptureAudioPreviewOutput {
	return CaptureAudioPreviewOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureAudioPreviewOutput */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureAudioPreviewOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureAudioPreviewOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureAudioPreviewOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureAudioPreviewOutput */

// The unique identifier of the Core Audio output device to use for audio preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput/outputDeviceUniqueID
func (c_ CaptureAudioPreviewOutput) OutputDeviceUniqueID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("outputDeviceUniqueID"))
	return rv
}/* debug [instance_properties/getter]: outputDeviceUniqueID */


// The unique identifier of the Core Audio output device to use for audio preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput/outputDeviceUniqueID
func (c_ CaptureAudioPreviewOutput) SetOutputDeviceUniqueID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputDeviceUniqueID:"), value)
}/* debug [instance_properties/setter]: outputDeviceUniqueID */


// The output volume of the audio preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput/volume
func (c_ CaptureAudioPreviewOutput) Volume() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("volume"))
	return rv
}/* debug [instance_properties/getter]: volume */


// The output volume of the audio preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput/volume
func (c_ CaptureAudioPreviewOutput) SetVolume(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVolume:"), value)
}/* debug [instance_properties/setter]: volume */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureAudioPreviewOutput */


