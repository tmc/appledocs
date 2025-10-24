// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureSystemPressureState */


/* debug [class_header]: Header for AVCaptureSystemPressureState */
// The class instance for the [CaptureSystemPressureState] class.
var (
	CaptureSystemPressureStateClass     _CaptureSystemPressureStateClass
	CaptureSystemPressureStateClassOnce sync.Once
)

func getCaptureSystemPressureStateClass() _CaptureSystemPressureStateClass {
	CaptureSystemPressureStateClassOnce.Do(func() {
		CaptureSystemPressureStateClass = _CaptureSystemPressureStateClass{objc.GetClass("AVCaptureSystemPressureState")}
	})
	return CaptureSystemPressureStateClass
}

type _CaptureSystemPressureStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSystemPressureState */
// An interface definition for the [CaptureSystemPressureState] class.
type ICaptureSystemPressureState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureSystemPressureState */
	// properties:
	SystemPressureState() IAVCaptureSystemPressureState
	SetSystemPressureState(value IAVCaptureSystemPressureState)
	AVCaptureSessionInterruptionSystemPressureStateKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSystemPressureState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSystemPressureState */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSystemPressureStateClass) Alloc() CaptureSystemPressureState {
	rv := objc.Send[CaptureSystemPressureState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSystemPressureStateClass) New() CaptureSystemPressureState {
	rv := objc.Send[CaptureSystemPressureState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSystemPressureState) Init() CaptureSystemPressureState {
	rv := objc.Send[CaptureSystemPressureState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSystemPressureState) Autorelease() CaptureSystemPressureState {
	rv := objc.Send[CaptureSystemPressureState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSystemPressureState creates a new CaptureSystemPressureState instance.
func NewCaptureSystemPressureState() CaptureSystemPressureState {
	return getCaptureSystemPressureStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSystemPressureState */
// An object that provides information about OS and hardware status affecting capture system performance and availability.
//
// The performance and availability of the camera capture system on an iOS device is subject to several external factors, such as power usage and device temperature. If during a capture session the total system pressure reaches excessive levels, the capture system automatically shuts down, causing a session interruption (see ). Under less heavy pressure, the system may automatically reduce capture quality. Key-value observe the capture device’s property to monitor its state, and take action to reduce the performance impact of your capture session when system pressure increases—for example, by reducing the capture frame rate.


// An object that provides information about OS and hardware status affecting capture system performance and availability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class
type CaptureSystemPressureState struct {
	objectivec.Object
}

// CaptureSystemPressureStateFrom constructs a [CaptureSystemPressureState] from an unsafe.Pointer.
//
// An object that provides information about OS and hardware status affecting capture system performance and availability.
func CaptureSystemPressureStateFrom(ptr unsafe.Pointer) CaptureSystemPressureState {
	return CaptureSystemPressureState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSystemPressureState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSystemPressureState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSystemPressureState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSystemPressureState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSystemPressureState */

// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.property
func (c_ CaptureSystemPressureState) SystemPressureState() IAVCaptureSystemPressureState {
	rv := objc.Send[CaptureSystemPressureState](c_.ID, objc.Sel("systemPressureState"))
	return rv
}/* debug [instance_properties/getter]: systemPressureState */


// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.property
func (c_ CaptureSystemPressureState) SetSystemPressureState(value IAVCaptureSystemPressureState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemPressureState:"), value)
}/* debug [instance_properties/setter]: systemPressureState */


// A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesessioninterruptionsystempressurestatekey
func (c_ CaptureSystemPressureState) AVCaptureSessionInterruptionSystemPressureStateKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVCaptureSessionInterruptionSystemPressureStateKey"))
	return rv
}/* debug [instance_properties/getter]: AVCaptureSessionInterruptionSystemPressureStateKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSystemPressureState */


