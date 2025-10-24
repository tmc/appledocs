// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureEventSound */


/* debug [class_header]: Header for AVCaptureEventSound */
// The class instance for the [CaptureEventSound] class.
var (
	CaptureEventSoundClass     _CaptureEventSoundClass
	CaptureEventSoundClassOnce sync.Once
)

func getCaptureEventSoundClass() _CaptureEventSoundClass {
	CaptureEventSoundClassOnce.Do(func() {
		CaptureEventSoundClass = _CaptureEventSoundClass{objc.GetClass("AVCaptureEventSound")}
	})
	return CaptureEventSoundClass
}

type _CaptureEventSoundClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureEventSound */
// An interface definition for the [CaptureEventSound] class.
type ICaptureEventSound interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureEventSound */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureEventSound */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureEventSound */
// Alloc allocates a new instance without initialization.
func (cc _CaptureEventSoundClass) Alloc() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureEventSoundClass) New() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureEventSound) Init() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureEventSound) Autorelease() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureEventSound creates a new CaptureEventSound instance.
func NewCaptureEventSound() CaptureEventSound {
	return getCaptureEventSoundClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureEventSound */
// A sound object for a capture event.


// A sound object for a capture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound
type CaptureEventSound struct {
	objectivec.Object
}

// CaptureEventSoundFrom constructs a [CaptureEventSound] from an unsafe.Pointer.
//
// A sound object for a capture event.
func CaptureEventSoundFrom(ptr unsafe.Pointer) CaptureEventSound {
	return CaptureEventSound{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureEventSound */

// Creates a sound object for a capture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/init(url:)
func NewCaptureEventSoundWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) CaptureEventSound {
	instance := getCaptureEventSoundClass().Alloc()
	rv := objc.Send[CaptureEventSound](instance.ID, objc.Sel("initWithURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureEventSoundWithURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureEventSound */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureEventSound */

// The default sound for starting a video recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/beginVideoRecording
func (cc _CaptureEventSoundClass) BeginVideoRecordingSound() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("beginVideoRecordingSound"))
	return rv
}/* debug [class_properties_class/property]: beginVideoRecordingSound */

// The default sound for photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/cameraShutter
func (cc _CaptureEventSoundClass) CameraShutterSound() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("cameraShutterSound"))
	return rv
}/* debug [class_properties_class/property]: cameraShutterSound */

// The default sound for ending a video recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/endVideoRecording
func (cc _CaptureEventSoundClass) EndVideoRecordingSound() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("endVideoRecordingSound"))
	return rv
}/* debug [class_properties_class/property]: endVideoRecordingSound */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureEventSound */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureEventSound */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureEventSound */


