// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CaptureSystemPressureState] class.
type ICaptureSystemPressureState interface {
	objectivec.IObject
	Factors() unsafe.Pointer
	SetFactors(value unsafe.Pointer)
	Level() unsafe.Pointer
	SetLevel(value unsafe.Pointer)
	SystemPressureState() AVCaptureSystemPressureState
	SetSystemPressureState(value IAVCaptureSystemPressureState)
	AVCaptureSessionInterruptionSystemPressureStateKey() string
}

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

// Alloc allocates a new instance without initialization.
func (cc _CaptureSystemPressureStateClass) Alloc() CaptureSystemPressureState {
	rv := objc.Send[CaptureSystemPressureState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The set of underlying causes for the system pressure level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/factors-swift.property

func (c_ CaptureSystemPressureState) Factors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("factors"))
	return rv
}


// The set of underlying causes for the system pressure level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/factors-swift.property

func (c_ CaptureSystemPressureState) SetFactors(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFactors:"), value)
}


// The overall level of performance constraints on the capture system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/level-swift.property

func (c_ CaptureSystemPressureState) Level() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("level"))
	return rv
}


// The overall level of performance constraints on the capture system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/level-swift.property

func (c_ CaptureSystemPressureState) SetLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLevel:"), value)
}


// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.property

func (c_ CaptureSystemPressureState) SystemPressureState() AVCaptureSystemPressureState {
	rv := objc.Send[AVCaptureSystemPressureState](c_.ID, objc.Sel("systemPressureState"))
	return rv
}


// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.property

func (c_ CaptureSystemPressureState) SetSystemPressureState(value IAVCaptureSystemPressureState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemPressureState:"), value)
}


// A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesessioninterruptionsystempressurestatekey

func (c_ CaptureSystemPressureState) AVCaptureSessionInterruptionSystemPressureStateKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("AVCaptureSessionInterruptionSystemPressureStateKey"))
	return rv
}



