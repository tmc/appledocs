// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SensorDataList] class.
var (
	SensorDataListClass     _SensorDataListClass
	SensorDataListClassOnce sync.Once
)

func getSensorDataListClass() _SensorDataListClass {
	SensorDataListClassOnce.Do(func() {
		SensorDataListClass = _SensorDataListClass{objc.GetClass("CMSensorDataList")}
	})
	return SensorDataListClass
}

type _SensorDataListClass struct {
	class objc.Class
}

// An interface definition for the [SensorDataList] class.
type ISensorDataList interface {
	objectivec.IObject
}

// A list of the accelerometer data recorded by the system.
//
// You do not create instances of this class directly. Instead, you receive one as the result of a query for accelerometer data from a object. You use a sensor data list object to enumerate over the accelerometer data as shown in the following example:


// A list of the accelerometer data recorded by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorDataList

type SensorDataList struct {
	objectivec.Object
}

// SensorDataListFrom constructs a [SensorDataList] from an unsafe.Pointer.
//
// A list of the accelerometer data recorded by the system.
func SensorDataListFrom(ptr unsafe.Pointer) SensorDataList {
	return SensorDataList{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SensorDataListClass) Alloc() SensorDataList {
	rv := objc.Send[SensorDataList](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SensorDataListClass) New() SensorDataList {
	rv := objc.Send[SensorDataList](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SensorDataList) Init() SensorDataList {
	rv := objc.Send[SensorDataList](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SensorDataList) Autorelease() SensorDataList {
	rv := objc.Send[SensorDataList](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSensorDataList creates a new SensorDataList instance.
func NewSensorDataList() SensorDataList {
	return getSensorDataListClass().New()
}




