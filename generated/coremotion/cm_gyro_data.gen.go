// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GyroData] class.
var (
	GyroDataClass     _GyroDataClass
	GyroDataClassOnce sync.Once
)

func getGyroDataClass() _GyroDataClass {
	GyroDataClassOnce.Do(func() {
		GyroDataClass = _GyroDataClass{objc.GetClass("CMGyroData")}
	})
	return GyroDataClass
}

type _GyroDataClass struct {
	class objc.Class
}

// An interface definition for the [GyroData] class.
type IGyroData interface {
	ILogItem
	// properties:
	RotationRate() RotationRate /* not a class type */
	// methods:
}

// A single measurement of the device’s rotation rate.
//
// An application receives or samples objects at regular intervals after calling the method or the method of the class.


// A single measurement of the device’s rotation rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMGyroData
type GyroData struct {
	LogItem
}

// GyroDataFrom constructs a [GyroData] from an unsafe.Pointer.
//
// A single measurement of the device’s rotation rate.
func GyroDataFrom(ptr unsafe.Pointer) GyroData {
	return GyroData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GyroDataClass) Alloc() GyroData {
	rv := objc.Send[GyroData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GyroDataClass) New() GyroData {
	rv := objc.Send[GyroData](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GyroData) Init() GyroData {
	rv := objc.Send[GyroData](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GyroData) Autorelease() GyroData {
	rv := objc.Send[GyroData](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGyroData creates a new GyroData instance.
func NewGyroData() GyroData {
	return getGyroDataClass().New()
}



// The rotation rate as measured by the device’s gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMGyroData/rotationRate
func (g_ GyroData) RotationRate() RotationRate /* not a class type */ {
	rv := objc.Send[RotationRate](g_.ID, objc.Sel("rotationRate"))
	return rv
}



