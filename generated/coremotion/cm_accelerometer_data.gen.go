// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerometerData] class.
var (
	AccelerometerDataClass     _AccelerometerDataClass
	AccelerometerDataClassOnce sync.Once
)

func getAccelerometerDataClass() _AccelerometerDataClass {
	AccelerometerDataClassOnce.Do(func() {
		AccelerometerDataClass = _AccelerometerDataClass{objc.GetClass("CMAccelerometerData")}
	})
	return AccelerometerDataClass
}

type _AccelerometerDataClass struct {
	class objc.Class
}

// An interface definition for the [AccelerometerData] class.
type IAccelerometerData interface {
	ILogItem
	Acceleration() unsafe.Pointer
	Timestamp() unsafe.Pointer
	SetTimestamp(value unsafe.Pointer)
	AccelerometerData() ICMAccelerometerData
	SetAccelerometerData(value ICMAccelerometerData)
}

// A data sample from the device’s three accelerometers.
//
// An application accesses objects through the block handler specified as the last parameter of the method and through the property, both declared by the class. The superclass of , , defines a property that records when the acceleration measurement was taken.


// A data sample from the device’s three accelerometers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAccelerometerData
type AccelerometerData struct {
	LogItem
}

// AccelerometerDataFrom constructs a [AccelerometerData] from an unsafe.Pointer.
//
// A data sample from the device’s three accelerometers.
func AccelerometerDataFrom(ptr unsafe.Pointer) AccelerometerData {
	return AccelerometerData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerometerDataClass) Alloc() AccelerometerData {
	rv := objc.Send[AccelerometerData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerometerDataClass) New() AccelerometerData {
	rv := objc.Send[AccelerometerData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerometerData) Init() AccelerometerData {
	rv := objc.Send[AccelerometerData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerometerData) Autorelease() AccelerometerData {
	rv := objc.Send[AccelerometerData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerometerData creates a new AccelerometerData instance.
func NewAccelerometerData() AccelerometerData {
	return getAccelerometerDataClass().New()
}



// The acceleration measured by the accelerometer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAccelerometerData/acceleration
func (a_ AccelerometerData) Acceleration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("acceleration"))
	return rv
}


// The time when the logged item is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmlogitem/timestamp
func (a_ AccelerometerData) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("timestamp"))
	return rv
}


// The time when the logged item is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmlogitem/timestamp
func (a_ AccelerometerData) SetTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimestamp:"), value)
}


// The latest sample of accelerometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/accelerometerdata
func (a_ AccelerometerData) AccelerometerData() ICMAccelerometerData {
	rv := objc.Send[AccelerometerData](a_.ID, objc.Sel("accelerometerData"))
	return rv
}


// The latest sample of accelerometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/accelerometerdata
func (a_ AccelerometerData) SetAccelerometerData(value ICMAccelerometerData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccelerometerData:"), value)
}



