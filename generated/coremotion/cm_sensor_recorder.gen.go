// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SensorRecorder] class.
var (
	SensorRecorderClass     _SensorRecorderClass
	SensorRecorderClassOnce sync.Once
)

func getSensorRecorderClass() _SensorRecorderClass {
	SensorRecorderClassOnce.Do(func() {
		SensorRecorderClass = _SensorRecorderClass{objc.GetClass("CMSensorRecorder")}
	})
	return SensorRecorderClass
}

type _SensorRecorderClass struct {
	class objc.Class
}

// An interface definition for the [SensorRecorder] class.
type ISensorRecorder interface {
	objectivec.IObject
	AccelerometerDataFromDateToDate(fromDate foundation.NSDate, toDate foundation.NSDate) ISensorDataList
	RecordAccelerometerForDuration(duration foundation.TimeInterval)
}

// An object that gathers and retrieves accelerometer data from a device.
//
// Use a sensor recorder to initiate the gathering of accelerometer data. Later, use the sensor recorder to fetch the recorded data so you can analyze it. You might use the recorded data to assess specific types of motion and incorporate the results into your app. To use a sensor recorder, create an instance of this class and call the method to begin recording data. You do not need to stop the recording process explicitly. The system stops recording automatically when the specified time expires and no other apps extend the recording time. The following example shows how to record 20 minutes worth of accelerometer data:


// An object that gathers and retrieves accelerometer data from a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorRecorder
type SensorRecorder struct {
	objectivec.Object
}

// SensorRecorderFrom constructs a [SensorRecorder] from an unsafe.Pointer.
//
// An object that gathers and retrieves accelerometer data from a device.
func SensorRecorderFrom(ptr unsafe.Pointer) SensorRecorder {
	return SensorRecorder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SensorRecorderClass) Alloc() SensorRecorder {
	rv := objc.Send[SensorRecorder](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SensorRecorderClass) New() SensorRecorder {
	rv := objc.Send[SensorRecorder](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SensorRecorder) Init() SensorRecorder {
	rv := objc.Send[SensorRecorder](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SensorRecorder) Autorelease() SensorRecorder {
	rv := objc.Send[SensorRecorder](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSensorRecorder creates a new SensorRecorder instance.
func NewSensorRecorder() SensorRecorder {
	return getSensorRecorderClass().New()
}



// Returns a value indicating whether the app is authorized to record sensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorRecorder/authorizationStatus()
func (sc _SensorRecorderClass) AuthorizationStatus() CMAuthorizationStatus {
	rv := objc.Send[CMAuthorizationStatus](objc.ID(sc.class), objc.Sel("authorizationStatus"))
	return rv
}


// Returns a Boolean value indicating whether accelerometer recording is supported on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorRecorder/isAccelerometerRecordingAvailable()
func (sc _SensorRecorderClass) IsAccelerometerRecordingAvailable() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("isAccelerometerRecordingAvailable"))
	return rv
}


// Returns a Boolean value indicating whether the app is authorized to record sensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorRecorder/isAuthorizedForRecording()
func (sc _SensorRecorderClass) IsAuthorizedForRecording() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("isAuthorizedForRecording"))
	return rv
}


// Retrieves the accelerometer data collected between the specified dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorRecorder/accelerometerData(from:to:)
func (s_ SensorRecorder) AccelerometerDataFromDateToDate(fromDate foundation.NSDate, toDate foundation.NSDate) ISensorDataList {
	rv := objc.Send[SensorDataList](s_.ID, objc.Sel("accelerometerDataFromDate:toDate:"), fromDate, toDate)
	return rv
}


// Begins recording accelerometer data for the specified period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorRecorder/recordAccelerometer(forDuration:)
func (s_ SensorRecorder) RecordAccelerometerForDuration(duration foundation.TimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("recordAccelerometerForDuration:"), duration)
}



