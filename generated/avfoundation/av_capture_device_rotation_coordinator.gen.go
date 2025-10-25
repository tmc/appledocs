// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureDeviceRotationCoordinator */


/* debug [class_header]: Header for AVCaptureDeviceRotationCoordinator */
// The class instance for the [CaptureDeviceRotationCoordinator] class.
var (
	CaptureDeviceRotationCoordinatorClass     _CaptureDeviceRotationCoordinatorClass
	CaptureDeviceRotationCoordinatorClassOnce sync.Once
)

func getCaptureDeviceRotationCoordinatorClass() _CaptureDeviceRotationCoordinatorClass {
	CaptureDeviceRotationCoordinatorClassOnce.Do(func() {
		CaptureDeviceRotationCoordinatorClass = _CaptureDeviceRotationCoordinatorClass{objc.GetClass("AVCaptureDeviceRotationCoordinator")}
	})
	return CaptureDeviceRotationCoordinatorClass
}

type _CaptureDeviceRotationCoordinatorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureDeviceRotationCoordinator */
// An interface definition for the [CaptureDeviceRotationCoordinator] class.
type ICaptureDeviceRotationCoordinator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureDeviceRotationCoordinator */
	// properties:
	Device() IAVCaptureDevice
	PreviewLayer() objc.IObject /* cross-framework: Layer */
	VideoRotationAngleForHorizonLevelCapture() float64
	VideoRotationAngleForHorizonLevelPreview() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureDeviceRotationCoordinator */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureDeviceRotationCoordinator */
// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceRotationCoordinatorClass) Alloc() CaptureDeviceRotationCoordinator {
	rv := objc.Send[CaptureDeviceRotationCoordinator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDeviceRotationCoordinatorClass) New() CaptureDeviceRotationCoordinator {
	rv := objc.Send[CaptureDeviceRotationCoordinator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeviceRotationCoordinator) Init() CaptureDeviceRotationCoordinator {
	rv := objc.Send[CaptureDeviceRotationCoordinator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeviceRotationCoordinator) Autorelease() CaptureDeviceRotationCoordinator {
	rv := objc.Send[CaptureDeviceRotationCoordinator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeviceRotationCoordinator creates a new CaptureDeviceRotationCoordinator instance.
func NewCaptureDeviceRotationCoordinator() CaptureDeviceRotationCoordinator {
	return getCaptureDeviceRotationCoordinatorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureDeviceRotationCoordinator */
// A class that monitors the physical orientation of a capture device and provides adjustment angles to keep images level, relative to gravity.
//
// Correctly rotate the photos and movies your app captures, and optionally, a live camera preview, by applying a coordinator’s and properties, respectively. Each rotation coordinator instance updates its properties so that your app can observe them and immediately apply them to the relevant components.


// A class that monitors the physical orientation of a capture device and provides adjustment angles to keep images level, relative to gravity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator
type CaptureDeviceRotationCoordinator struct {
	objectivec.Object
}

// CaptureDeviceRotationCoordinatorFrom constructs a [CaptureDeviceRotationCoordinator] from an unsafe.Pointer.
//
// A class that monitors the physical orientation of a capture device and provides adjustment angles to keep images level, relative to gravity.
func CaptureDeviceRotationCoordinatorFrom(ptr unsafe.Pointer) CaptureDeviceRotationCoordinator {
	return CaptureDeviceRotationCoordinator{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureDeviceRotationCoordinator */

// Creates a coordinator that provides separate compensation angles for content your app takes with a capture device, and for your app’s camera preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/init(device:previewLayer:)
func NewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer(device IAVCaptureDevice, previewLayer objc.IObject /* cross-framework: Layer */) CaptureDeviceRotationCoordinator {
	instance := getCaptureDeviceRotationCoordinatorClass().Alloc()
	rv := objc.Send[CaptureDeviceRotationCoordinator](instance.ID, objc.Sel("initWithDevice:previewLayer:"), device, previewLayer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureDeviceRotationCoordinator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureDeviceRotationCoordinator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureDeviceRotationCoordinator */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureDeviceRotationCoordinator */

// The capture device the coordinator monitors to track its physical rotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/device
func (c_ CaptureDeviceRotationCoordinator) Device() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The layer that displays a camera preview the coordinator calculates a video rotation angle for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/previewLayer
func (c_ CaptureDeviceRotationCoordinator) PreviewLayer() objc.IObject /* cross-framework: Layer */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("previewLayer"))
	return rv
}/* debug [instance_properties/getter]: previewLayer */


// An angle the coordinator provides your app to apply to photos or videos it captures with the device so that they’re level relative to gravity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/videoRotationAngleForHorizonLevelCapture
func (c_ CaptureDeviceRotationCoordinator) VideoRotationAngleForHorizonLevelCapture() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngleForHorizonLevelCapture"))
	return rv
}/* debug [instance_properties/getter]: videoRotationAngleForHorizonLevelCapture */


// An angle the coordinator provides your app to apply to the preview layer so that it’s level relative to gravity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/videoRotationAngleForHorizonLevelPreview
func (c_ CaptureDeviceRotationCoordinator) VideoRotationAngleForHorizonLevelPreview() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngleForHorizonLevelPreview"))
	return rv
}/* debug [instance_properties/getter]: videoRotationAngleForHorizonLevelPreview */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureDeviceRotationCoordinator */


